package stringadapter

import (
	"context"

	"github.com/casbin/casbin/v3/model"
)

func (a *Adapter) LoadPolicyCtx(ctx context.Context, model model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) SavePolicyCtx(ctx context.Context, model model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) AddPolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemovePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemoveFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkCtx(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
