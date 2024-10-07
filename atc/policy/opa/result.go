package opa

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/concourse/concourse/atc/policy"
	"github.com/concourse/concourse/vars"
)

func ParseOpaResult(bytesResult []byte, opaConfig OpaConfig) (policy.PolicyCheckResult, error) {
	var results vars.StaticVariables
	err := json.Unmarshal(bytesResult, &results)
	if err != nil {
		return policy.PolicyCheckResult{}, err
	}

	var allowed, shouldBlock, ok bool
	var messages []string

	parts := strings.Split(opaConfig.ResultAllowedKey, ".")
	v, found, err := results.Get(vars.Reference{Path: parts[0], Fields: parts[1:]})
	if err != nil {
		return policy.PolicyCheckResult{}, fmt.Errorf("allowed: %w", err)
	}
	if !found {
		return policy.PolicyCheckResult{}, fmt.Errorf("allowed: key '%s' not found", opaConfig.ResultAllowedKey)
	}
	if allowed, ok = v.(bool); !ok {
		return policy.PolicyCheckResult{}, fmt.Errorf("allowed: key '%s' must have a boolean value", opaConfig.ResultAllowedKey)
	}

	parts = strings.Split(opaConfig.ResultShouldBlockKey, ".")
	v, found, err = results.Get(vars.Reference{Path: parts[0], Fields: parts[1:]})
	if err != nil || !found {
		shouldBlock = !allowed
	} else if shouldBlock, ok = v.(bool); !ok {
		return policy.PolicyCheckResult{}, fmt.Errorf("shouldBlock: key '%s' must have a boolean value", opaConfig.ResultShouldBlockKey)
	}

	parts = strings.Split(opaConfig.ResultMessagesKey, ".")
	v, found, err = results.Get(vars.Reference{Path: parts[0], Fields: parts[1:]})
	if err != nil || !found {
		messages = []string{}
	} else if arr, ok := v.([]interface{}); v != nil && !ok {
		return policy.PolicyCheckResult{}, fmt.Errorf("messages: key '%s' must have a list of strings", opaConfig.ResultMessagesKey)
	} else {
		for _, item := range arr {
			switch v := item.(type) {
			case string:
				messages = append(messages, v)
			default:
				return policy.PolicyCheckResult{}, fmt.Errorf("invalid messages")
			}
		}
	}

	return policy.PolicyCheckResult{Allowed: allowed, ShouldBlock: shouldBlock, Messages: messages}, nil
}
