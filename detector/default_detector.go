package detector

import (
	"github.com/casbin/casbin/v3/rbac"
)

type rangeableRM interface {
	Range(func(name1, name2 string, domain ...string) bool)
}

type DefaultDetector struct{}

func NewDefaultDetector() *DefaultDetector { _ = "STUB: not implemented"; return nil }

func (d *DefaultDetector) Check(rm rbac.RoleManager) error { _ = "STUB: not implemented"; return nil }

func (d *DefaultDetector) buildGraph(rm rbac.RoleManager) (graph map[string][]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DefaultDetector) detectCycle(
	role string,
	graph map[string][]string,
	visited map[string]bool,
	recursionStack map[string]bool,
	path []string,
) []string {
	_ = "STUB: not implemented"
	return nil
}
