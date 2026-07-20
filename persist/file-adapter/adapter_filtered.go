package fileadapter

import (
	"github.com/casbin/casbin/v3/model"
)

type FilteredAdapter struct {
	*Adapter
	filtered bool
}

type Filter struct {
	P  []string
	G  []string
	G1 []string
	G2 []string
	G3 []string
	G4 []string
	G5 []string
}

func NewFilteredAdapter(filePath string) *FilteredAdapter { _ = "STUB: not implemented"; return nil }

func (a *FilteredAdapter) LoadPolicy(model model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *FilteredAdapter) LoadFilteredPolicy(model model.Model, filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *FilteredAdapter) loadFilteredPolicyFile(model model.Model, filter *Filter, handler func(string, model.Model) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *FilteredAdapter) IsFiltered() bool { _ = "STUB: not implemented"; return false }

func (a *FilteredAdapter) SavePolicy(model model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func filterLine(line string, filter *Filter) bool { _ = "STUB: not implemented"; return false }

func filterWords(line []string, filter []string) bool { _ = "STUB: not implemented"; return false }
