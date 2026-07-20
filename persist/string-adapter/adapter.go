package stringadapter

import (
	"github.com/casbin/casbin/v3/model"
)

type Adapter struct {
	Line string
}

func NewAdapter(line string) *Adapter { _ = "STUB: not implemented"; return nil }

func (a *Adapter) LoadPolicy(model model.Model) error { _ = "STUB: not implemented"; return nil }

func (a *Adapter) SavePolicy(model model.Model) error { _ = "STUB: not implemented"; return nil }

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	_ = "STUB: not implemented"
	return nil
}
