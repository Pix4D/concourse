package testflight_test

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("A pipeline containing idtoken var sources", Ordered, func() {

	var testPipelineName string
	var jwks *jose.JSONWebKeySet
	var outputText []byte

	type claimStruct struct {
		jwt.Claims
		Team     string `json:"team"`
		Pipeline string `json:"pipeline"`
	}

	BeforeAll(func() {
		testPipelineName = pipelineName
		var err error
		jwks, err = getJWKS(config.ATCURL)
		Expect(err).ToNot(HaveOccurred())

		setAndUnpausePipeline("fixtures/idtoken.yml")
		watch := fly("trigger-job", "-j", inPipeline("print-creds"), "-w")
		Expect(watch).To(gexec.Exit(0))
		outputText = watch.Buffer().Contents()
	})

	It("creates valid default idtoken", func() {
		token := extractIDtokenFromBuffer(outputText, "default-token")
		Expect(token).ToNot(BeEmpty())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())
		var claims claimStruct
		err = parsed.Claims(jwks, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Audience).To(ContainElement("sts.amazonaws.com"))
		Expect(claims.Team).To(Equal(teamName))
		Expect(claims.Pipeline).To(Equal(testPipelineName))
		Expect(claims.Subject).To(Equal(teamName + "/" + testPipelineName))
	})

	It("creates valid custom idtoken", func() {
		token := extractIDtokenFromBuffer(outputText, "custom-token")
		Expect(token).ToNot(BeEmpty())

		parsed, err := jwt.ParseSigned(token)
		Expect(err).ToNot(HaveOccurred())
		var claims claimStruct
		err = parsed.Claims(jwks, &claims)
		Expect(err).To(Succeed())

		Expect(claims.Audience).To(ContainElement("sts.amazonaws.com"))
		Expect(claims.Team).To(Equal(teamName))
		Expect(claims.Pipeline).To(Equal(testPipelineName))

		Expect(parsed.Headers[0].Algorithm).To(Equal("ES256"))
		Expect(claims.Subject).To(Equal(teamName))
	})
})

func extractIDtokenFromBuffer(buffer []byte, whichToken string) string {
	tokenMatcher := regexp.MustCompile("(?m)" + whichToken + ": (.*)$")
	tokenMatches := tokenMatcher.FindSubmatch(buffer)
	if len(tokenMatches) != 2 {
		return ""
	}
	return string(tokenMatches[1])
}

func getJWKS(atcurl string) (*jose.JSONWebKeySet, error) {
	url, err := getJWKSURL(atcurl)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var jwks jose.JSONWebKeySet

	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return nil, err
	}

	return &jwks, nil
}

func getJWKSURL(atcurl string) (string, error) {
	url := atcurl + "/.well-known/openid-configuration"
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	type oidcConfig struct {
		JWKSURI string `json:"jwks_uri"`
	}

	var conf oidcConfig

	err = json.NewDecoder(resp.Body).Decode(&conf)
	if err != nil {
		return "", err
	}

	return conf.JWKSURI, nil
}
