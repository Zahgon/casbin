package persist

import "context"

type ContextUpdatableAdapter interface {
	ContextAdapter

	UpdatePolicyCtx(ctx context.Context, sec string, ptype string, oldRule, newRule []string) error

	UpdatePoliciesCtx(ctx context.Context, sec string, ptype string, oldRules, newRules [][]string) error

	UpdateFilteredPoliciesCtx(ctx context.Context, sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error)
}
