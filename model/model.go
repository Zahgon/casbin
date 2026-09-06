package model

import (
	"regexp"

	"github.com/casbin/casbin/v3/config"
)

type Model map[string]AssertionMap

type AssertionMap map[string]*Assertion

const (
	defaultDomain    string = ""
	defaultSeparator string = "::"
)

var sectionNameMap = map[string]string{
	"r": "request_definition",
	"p": "policy_definition",
	"g": "role_definition",
	"e": "policy_effect",
	"m": "matchers",
	"c": "constraint_definition",
}

var requiredSections = []string{"r", "p", "e", "m"}

func loadAssertion(model Model, cfg config.ConfigInterface, sec string, key string) bool {
	_ = "STUB: not implemented"
	return false
}

var paramsRegex = regexp.MustCompile(`\((.*?)\)`)

func getParamsToken(value string) []string { _ = "STUB: not implemented"; return nil }

func (model Model) AddDef(sec string, key string, value string) bool {
	_ = "STUB: not implemented"
	return false
}

func getKeySuffix(i int) string { _ = "STUB: not implemented"; return "" }

func loadSection(model Model, cfg config.ConfigInterface, sec string) {
	_ = "STUB: not implemented"
	return
}

func NewModel() Model { _ = "STUB: not implemented"; return *new(Model) }

func NewModelFromFile(path string) (Model, error) {
	_ = "STUB: not implemented"
	return *new(Model), nil
}

func NewModelFromString(text string) (Model, error) {
	_ = "STUB: not implemented"
	return *new(Model), nil
}

func (model Model) LoadModel(path string) error { _ = "STUB: not implemented"; return nil }

func (model Model) LoadModelFromText(text string) error { _ = "STUB: not implemented"; return nil }

func (model Model) loadModelFromConfig(cfg config.ConfigInterface) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) hasSection(sec string) bool { _ = "STUB: not implemented"; return false }

func (model Model) GetAssertion(sec string, ptype string) (*Assertion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) PrintModel() { _ = "STUB: not implemented"; return }

func (model Model) SortPoliciesBySubjectHierarchy() error { _ = "STUB: not implemented"; return nil }

func getSubjectHierarchyMap(policies [][]string) (map[string]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNameWithDomain(domain string, name string) string { _ = "STUB: not implemented"; return "" }

func (model Model) SortPoliciesByPriority() error { _ = "STUB: not implemented"; return nil }

var (
	pPattern = regexp.MustCompile("^p_")
	rPattern = regexp.MustCompile("^r_")
)

func (model Model) ToText() string { _ = "STUB: not implemented"; return "" }

func (model Model) Copy() Model { _ = "STUB: not implemented"; return *new(Model) }

func (model Model) GetFieldIndex(ptype string, field string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
