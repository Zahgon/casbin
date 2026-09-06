package fileadapter

import (
	"context"

	"github.com/casbin/casbin/v3/model"
)

func (a *FilteredAdapter) LoadPolicyCtx(ctx context.Context, model model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *FilteredAdapter) LoadFilteredPolicyCtx(ctx context.Context, model model.Model, filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *FilteredAdapter) SavePolicyCtx(ctx context.Context, model model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *FilteredAdapter) IsFilteredCtx(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}
