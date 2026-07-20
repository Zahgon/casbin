package persist

import "context"

type ContextBatchAdapter interface {
	ContextAdapter

	AddPoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error

	RemovePoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error
}
