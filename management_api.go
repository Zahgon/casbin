package casbin

import (
	"github.com/casbin/govaluate"
)

func (e *Enforcer) GetAllSubjects() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetAllNamedSubjects(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetAllObjects() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetAllNamedObjects(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetAllActions() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetAllNamedActions(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetAllRoles() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetAllNamedRoles(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetAllUsers() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetPolicy() ([][]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetFilteredPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetNamedPolicy(ptype string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetGroupingPolicy() ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetNamedGroupingPolicy(ptype string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetFilteredNamedPolicyWithMatcher(ptype string, matcher string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) HasPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) HasNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddPoliciesEx(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddNamedPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddNamedPoliciesEx(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemovePolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdatePolicy(oldPolicy []string, newPolicy []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateNamedPolicy(ptype string, p1 []string, p2 []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdatePolicies(oldPolices [][]string, newPolicies [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateNamedPolicies(ptype string, p1 [][]string, p2 [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateFilteredPolicies(newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateFilteredNamedPolicies(ptype string, newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemovePolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveNamedPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) HasGroupingPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) HasNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddGroupingPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddGroupingPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddGroupingPoliciesEx(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddNamedGroupingPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddNamedGroupingPoliciesEx(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveGroupingPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveGroupingPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveNamedGroupingPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateGroupingPolicy(oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateGroupingPolicies(oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateNamedGroupingPolicy(ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) UpdateNamedGroupingPolicies(ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddFunction(name string, function govaluate.ExpressionFunction) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) SelfAddPolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfAddPolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfAddPoliciesEx(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfRemovePolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfRemovePolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfRemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfUpdatePolicy(sec string, ptype string, oldRule, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) SelfUpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
