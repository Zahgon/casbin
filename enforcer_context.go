package casbin

import (
	"context"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
)

type ContextEnforcer struct {
	*Enforcer
	adapterCtx persist.ContextAdapter
}

func NewContextEnforcer(params ...interface{}) (IEnforcerContext, error) {
	_ = "STUB: not implemented"
	return *new(IEnforcerContext), nil
}

func (e *ContextEnforcer) LoadPolicyCtx(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ContextEnforcer) loadPolicyFromAdapterCtx(ctx context.Context, baseModel model.Model) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func (e *Enforcer) LoadFilteredPolicyCtx(ctx context.Context, filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) LoadIncrementalFilteredPolicyCtx(ctx context.Context, filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) loadFilteredPolicyCtx(ctx context.Context, filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ContextEnforcer) IsFilteredCtx(ctx context.Context) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *ContextEnforcer) SavePolicyCtx(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *ContextEnforcer) AddPolicyCtx(ctx context.Context, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddPoliciesCtx(ctx context.Context, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddNamedPolicyCtx(ctx context.Context, ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddNamedPoliciesCtx(ctx context.Context, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddPoliciesExCtx(ctx context.Context, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddNamedPoliciesExCtx(ctx context.Context, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemovePolicyCtx(ctx context.Context, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveNamedPolicyCtx(ctx context.Context, ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemovePoliciesCtx(ctx context.Context, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveNamedPoliciesCtx(ctx context.Context, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveFilteredPolicyCtx(ctx context.Context, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveFilteredNamedPolicyCtx(ctx context.Context, ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdatePolicyCtx(ctx context.Context, oldPolicy []string, newPolicy []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateNamedPolicyCtx(ctx context.Context, ptype string, p1 []string, p2 []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdatePoliciesCtx(ctx context.Context, oldPolicies [][]string, newPolicies [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateNamedPoliciesCtx(ctx context.Context, ptype string, p1 [][]string, p2 [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateFilteredPoliciesCtx(ctx context.Context, newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateFilteredNamedPoliciesCtx(ctx context.Context, ptype string, newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddGroupingPolicyCtx(ctx context.Context, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddGroupingPoliciesCtx(ctx context.Context, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddGroupingPoliciesExCtx(ctx context.Context, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddNamedGroupingPolicyCtx(ctx context.Context, ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddNamedGroupingPoliciesCtx(ctx context.Context, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddNamedGroupingPoliciesExCtx(ctx context.Context, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveGroupingPolicyCtx(ctx context.Context, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveNamedGroupingPolicyCtx(ctx context.Context, ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveGroupingPoliciesCtx(ctx context.Context, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveNamedGroupingPoliciesCtx(ctx context.Context, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveFilteredGroupingPolicyCtx(ctx context.Context, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) RemoveFilteredNamedGroupingPolicyCtx(ctx context.Context, ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateGroupingPolicyCtx(ctx context.Context, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateNamedGroupingPolicyCtx(ctx context.Context, ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateGroupingPoliciesCtx(ctx context.Context, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) UpdateNamedGroupingPoliciesCtx(ctx context.Context, ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfAddPolicyCtx(ctx context.Context, sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfAddPoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfAddPoliciesExCtx(ctx context.Context, sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfRemovePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfRemovePoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfRemoveFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfUpdatePolicyCtx(ctx context.Context, sec string, ptype string, oldRule, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) SelfUpdatePoliciesCtx(ctx context.Context, sec string, ptype string, oldRules, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) addPolicyWithoutNotifyCtx(ctx context.Context, sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) addPoliciesWithoutNotifyCtx(ctx context.Context, sec string, ptype string, rules [][]string, autoRemoveRepeat bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) removePolicyWithoutNotifyCtx(ctx context.Context, sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) removePoliciesWithoutNotifyCtx(ctx context.Context, sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) removeFilteredPolicyWithoutNotifyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) updatePolicyWithoutNotifyCtx(ctx context.Context, sec string, ptype string, oldRule, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) updatePoliciesWithoutNotifyCtx(ctx context.Context, sec string, ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) addPolicyCtx(ctx context.Context, sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) addPoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string, autoRemoveRepeat bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) updatePolicyCtx(ctx context.Context, sec string, ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) updatePoliciesCtx(ctx context.Context, sec string, ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) removePolicyCtx(ctx context.Context, sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) removePoliciesCtx(ctx context.Context, sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) removeFilteredPolicyCtx(ctx context.Context, sec string, ptype string, fieldIndex int, fieldValues []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) updateFilteredPoliciesCtx(ctx context.Context, sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) updateFilteredPoliciesWithoutNotifyCtx(ctx context.Context, sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
