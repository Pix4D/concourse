package idtoken_test

import (
	"fmt"
	"time"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/creds/idtoken"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/dbfakes"
	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	testIssuer     = "https://concourse.test"
	tokenExpiresIn = 15 * time.Minute
)

var _ = Describe("IDToken TokenGenerator", func() {
	var rsaSigningKey db.SigningKey
	var rsaVerificationKey jose.JSONWebKey
	var ecSigningKey db.SigningKey
	var ecVerificationKey jose.JSONWebKey
	var signingKeyFactory db.SigningKeyFactory
	var tokenGenerator *idtoken.TokenGenerator
	var params creds.SecretLookupParams

	BeforeEach(func() {
		rsaSigningKey = createFakeSigningKey(*rsaJWK, time.Now())
		ecSigningKey = createFakeSigningKey(*ecJWK, time.Now())

		rsaVerificationKey = rsaJWK.Public()
		ecVerificationKey = ecJWK.Public()

		signingKeyFactoryFake := &dbfakes.FakeSigningKeyFactory{}
		signingKeyFactoryFake.GetAllKeysReturns([]db.SigningKey{
			rsaSigningKey,
			ecSigningKey,
		}, nil)

		signingKeyFactoryFake.GetNewestKeyStub = func(skt db.SigningKeyType) (db.SigningKey, error) {
			switch skt {
			case db.SigningKeyTypeRSA:
				return rsaSigningKey, nil
			case db.SigningKeyTypeEC:
				return ecSigningKey, nil
			}
			return nil, fmt.Errorf("not found")
		}
		signingKeyFactory = signingKeyFactoryFake

		tokenGenerator = &idtoken.TokenGenerator{
			Issuer:            testIssuer,
			SigningKeyFactory: signingKeyFactory,
			ExpiresIn:         tokenExpiresIn,
		}

		params = creds.SecretLookupParams{
			Team:     "main",
			Pipeline: "idtoken",
			InstanceVars: atc.InstanceVars{
				"foo": "bar",
			},
			Job: "testjob",
		}
	})

	It("generates a valid token", func() {
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(rsaVerificationKey, &claims)
		Expect(err).To(Succeed())
		Expect(claims.Subject).To(Equal(params.Team + "/" + params.Pipeline))
	})

	It("respects subject scope team", func() {
		tokenGenerator.SubjectScope = idtoken.SubjectScopeTeam
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(rsaVerificationKey, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Subject).To(Equal(params.Team))
	})

	It("respects subject scope instance", func() {
		tokenGenerator.SubjectScope = idtoken.SubjectScopeInstance
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(rsaVerificationKey, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Subject).To(Equal(params.Team + "/" + params.Pipeline + "/" + params.InstanceVars.String()))
	})

	It("respects subject scope job", func() {
		tokenGenerator.SubjectScope = idtoken.SubjectScopeJob
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(rsaVerificationKey, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Subject).To(Equal(params.Team + "/" + params.Pipeline + "/" + params.InstanceVars.String() + "/" + params.Job))
	})

	It("escapes sub parts safely", func() {
		params = creds.SecretLookupParams{
			Team:     "fake/team",
			Pipeline: "fake/pipeline",
			InstanceVars: atc.InstanceVars{
				"fake/foo": "fake/bar",
			},
			Job: "fake/job",
		}
		tokenGenerator.SubjectScope = idtoken.SubjectScopeJob
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(rsaVerificationKey, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Subject).To(Equal("fake%2Fteam/fake%2Fpipeline/\"fake%2Ffoo\":\"fake%2Fbar\"/fake%2Fjob"))
	})

	It("adds aud claim when requested", func() {
		tokenGenerator.Audience = []string{"testaud"}
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(rsaVerificationKey, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Audience).To(ContainElement("testaud"))
	})

	It("uses ES256 when requested", func() {
		tokenGenerator.Algorithm = jose.ES256
		token, _, err := tokenGenerator.GenerateToken(params)
		Expect(err).ToNot(HaveOccurred())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())

		claims := jwt.Claims{}
		err = parsed.Claims(ecVerificationKey, &claims)
		Expect(err).To(Succeed())

		Expect(parsed.Headers[0].Algorithm).To(Equal("ES256"))
	})

	Context("Generated Token", func() {
		type claimStruct struct {
			jwt.Claims
			Team         string           `json:"team"`
			Pipeline     string           `json:"pipeline"`
			InstanceVars atc.InstanceVars `json:"instance_vars"`
			Job          string           `json:"job"`
		}

		var generatedToken string
		var claims claimStruct
		var generatedAt time.Time

		BeforeEach(func() {
			var err error
			generatedToken, _, err = tokenGenerator.GenerateToken(params)
			Expect(err).ToNot(HaveOccurred())

			generatedAt = time.Now()

			parsed, err := jwt.ParseSigned(generatedToken)
			Expect(err).ToNot(HaveOccurred())

			err = parsed.Claims(rsaVerificationKey, &claims)
			Expect(err).To(Succeed())
		})

		It("contains the correct subject", func() {
			Expect(claims.Subject).To(Equal(params.Team + "/" + params.Pipeline))
		})

		It("contains the correct issuer", func() {
			Expect(claims.Issuer).To(Equal(testIssuer))
		})

		It("has the correct expiration time", func() {
			exp := claims.Expiry.Time()
			expected := generatedAt.Add(tokenExpiresIn)
			difference := exp.Sub(expected)
			Expect(difference < 10*time.Second).To(BeTrue())
		})

		It("contains the correct team", func() {
			Expect(claims.Team).To(Equal(params.Team))
		})

		It("contains the correct pipeline", func() {
			Expect(claims.Pipeline).To(Equal(params.Pipeline))
		})

		It("contains the correct instance vars", func() {
			Expect(claims.InstanceVars["foo"]).ToNot(BeEmpty())
			Expect(claims.InstanceVars["foo"]).To(Equal(params.InstanceVars["foo"]))
		})

		It("contains the correct job", func() {
			Expect(claims.Job).To(Equal(params.Job))
		})

		It("has no default audience", func() {
			Expect(claims.Audience).To(HaveLen(0))
		})
	})
})
