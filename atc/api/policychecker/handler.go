package policychecker

import (
	"fmt"
	"net/http"

	"code.cloudfoundry.org/lager/v3"

	"github.com/concourse/concourse/atc/api/accessor"
	"github.com/concourse/concourse/atc/policy"
)

func NewHandler(
	logger lager.Logger,
	handler http.Handler,
	action string,
	policyChecker PolicyChecker,
) http.Handler {
	return policyCheckingHandler{
		logger:        logger,
		handler:       handler,
		action:        action,
		policyChecker: policyChecker,
	}
}

type policyCheckingHandler struct {
	logger        lager.Logger
	handler       http.Handler
	action        string
	policyChecker PolicyChecker
}

func (h policyCheckingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	acc := accessor.GetAccessor(r)

	logger := h.logger.Session("policy-check-handler")
	result, err := h.policyChecker.Check(h.action, acc, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "policy check error: %s", err.Error())
		// logged error should be meaningful to the Concourse operator/maintainer
		logger.Error("policy-checker", err)
		fmt.Fprintf(w, "policy checker: <meaningful-reason-for-the-(concourse-user>/client)>")
		return
	}

	// log only if we policy check return block or warn
	if result.Status != policy.Allow {
		logger.Info("policy-checker", lager.Data{"result": result.Status.String(), "reason": result.Reasons})
	}

	if result.Status == policy.Block {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, result.Reasons)
		return
	}

	if result.Status == policy.Warn {
		w.Header().Add("X-Concourse-Policy-Check-Warning", result.Reasons)
	}

	h.handler.ServeHTTP(w, r)
}
