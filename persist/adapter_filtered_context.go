package persist

import (
	"context"

	"github.com/casbin/casbin/v3/model"
)

type ContextFilteredAdapter interface {
	ContextAdapter

	LoadFilteredPolicyCtx(ctx context.Context, model model.Model, filter interface{}) error

	IsFilteredCtx(ctx context.Context) bool
}
