package fileadapter

import (
	"context"

	"github.com/casbin/casbin/v3/model"
)

func (a *Adapter) UpdatePolicyCtx(ctx context.Context, sec string, ptype string, oldRule, newRule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) UpdatePoliciesCtx(ctx context.Context, sec string, ptype string, oldRules, newRules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) UpdateFilteredPoliciesCtx(ctx context.Context, sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

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

func (a *Adapter) AddPoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemovePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemovePoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemoveFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkCtx(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
