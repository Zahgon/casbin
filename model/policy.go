package model

import (
	"github.com/casbin/casbin/v3/rbac"
)

type (
	PolicyOp int
)

const (
	PolicyAdd PolicyOp = iota
	PolicyRemove
)

const DefaultSep = ","

func (model Model) BuildIncrementalRoleLinks(rmMap map[string]rbac.RoleManager, op PolicyOp, sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) BuildRoleLinks(rmMap map[string]rbac.RoleManager) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) BuildIncrementalConditionalRoleLinks(condRmMap map[string]rbac.ConditionalRoleManager, op PolicyOp, sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) BuildConditionalRoleLinks(condRmMap map[string]rbac.ConditionalRoleManager) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) PrintPolicy() { _ = "STUB: not implemented"; return }

func (model Model) ClearPolicy() { _ = "STUB: not implemented"; return }

func (model Model) GetPolicy(sec string, ptype string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) GetFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) HasPolicyEx(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) HasPolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) HasPolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) AddPolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) AddPolicies(sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) AddPoliciesWithAffected(sec string, ptype string, rules [][]string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) RemovePolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) UpdatePolicy(sec string, ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) RemovePolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (model Model) RemovePoliciesWithAffected(sec string, ptype string, rules [][]string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (bool, [][]string, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (model Model) GetValuesForFieldInPolicy(sec string, ptype string, fieldIndex int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) GetValuesForFieldInPolicyAllTypes(sec string, fieldIndex int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) GetValuesForFieldInPolicyAllTypesByName(sec string, field string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
