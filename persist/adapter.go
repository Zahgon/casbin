package persist

import (
	"github.com/casbin/casbin/v3/model"
)

func LoadPolicyLine(line string, m model.Model) error { _ = "STUB: not implemented"; return nil }

func PolicyLineToCsv(ptype string, rule []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func LoadPolicyArray(rule []string, m model.Model) error { _ = "STUB: not implemented"; return nil }

type Adapter interface {
	LoadPolicy(model model.Model) error

	SavePolicy(model model.Model) error

	AddPolicy(sec string, ptype string, rule []string) error

	RemovePolicy(sec string, ptype string, rule []string) error

	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error
}
