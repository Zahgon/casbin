package fileadapter

import (
	"github.com/casbin/casbin/v3/model"
)

type Adapter struct {
	filePath string
}

func (a *Adapter) UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) UpdateFilteredPolicies(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewAdapter(filePath string) *Adapter { _ = "STUB: not implemented"; return nil }

func (a *Adapter) LoadPolicy(model model.Model) error { _ = "STUB: not implemented"; return nil }

func (a *Adapter) SavePolicy(model model.Model) error { _ = "STUB: not implemented"; return nil }

func (a *Adapter) loadPolicyFile(model model.Model, handler func(string, model.Model) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) savePolicyFile(text string) error { _ = "STUB: not implemented"; return nil }

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) AddPolicies(sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemovePolicies(sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	_ = "STUB: not implemented"
	return nil
}
