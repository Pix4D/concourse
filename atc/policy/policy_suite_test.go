package policy_test

import (
	"code.cloudfoundry.org/lager/v3/lagertest"
	"github.com/concourse/concourse/atc/policy"
	"github.com/concourse/concourse/atc/policy/policyfakes"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPolicy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Policy Suite")
}

var (
	testLogger = lagertest.NewTestLogger("test")

	fakeAgent        *policyfakes.FakeAgent
	fakeAgentFactory *policyfakes.FakeAgentFactory
)

var _ = BeforeSuite(func() {
	fakeAgentFactory = new(policyfakes.FakeAgentFactory)
	fakeAgentFactory.IsConfiguredReturns(true)
	fakeAgentFactory.DescriptionReturns("fakeAgent")
	policy.RegisterAgent(fakeAgentFactory)
})
