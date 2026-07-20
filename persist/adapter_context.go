package persist

import (
	"context"

	"github.com/casbin/casbin/v3/model"
)

type ContextAdapter interface {
	LoadPolicyCtx(ctx context.Context, model model.Model) error

	SavePolicyCtx(ctx context.Context, model model.Model) error

	AddPolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error

	RemovePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error

	RemoveFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues ...string) error
}
