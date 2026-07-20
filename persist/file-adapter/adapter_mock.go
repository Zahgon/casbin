package fileadapter

import (
	"github.com/casbin/casbin/v3/model"
)

type AdapterMock struct {
	filePath   string
	errorValue string
}

func NewAdapterMock(filePath string) *AdapterMock { _ = "STUB: not implemented"; return nil }

func (a *AdapterMock) LoadPolicy(model model.Model) error { _ = "STUB: not implemented"; return nil }

func (a *AdapterMock) SavePolicy(model model.Model) error { _ = "STUB: not implemented"; return nil }

func (a *AdapterMock) loadPolicyFile(model model.Model, handler func(string, model.Model) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) SetMockErr(errorToSet string) { _ = "STUB: not implemented"; return }

func (a *AdapterMock) GetMockErr() error { _ = "STUB: not implemented"; return nil }

func (a *AdapterMock) AddPolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) AddPolicies(sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) RemovePolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) RemovePolicies(sec string, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) UpdatePolicy(sec string, ptype string, oldRule, newPolicy []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *AdapterMock) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	_ = "STUB: not implemented"
	return nil
}
