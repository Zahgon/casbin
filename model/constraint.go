package model

import (
	"regexp"
)

type ConstraintType int

const (
	ConstraintTypeSOD ConstraintType = iota
	ConstraintTypeSODMax
	ConstraintTypeRoleMax
	ConstraintTypeRolePre
)

type Constraint struct {
	Key        string
	Type       ConstraintType
	Roles      []string
	Role       string
	MaxCount   int
	PreReqRole string
}

var (
	sodPattern     = regexp.MustCompile(`^sod\s*\(\s*"([^"]+)"\s*,\s*"([^"]+)"\s*\)$`)
	sodMaxPattern  = regexp.MustCompile(`^sodMax\s*\(\s*\[([^\]]+)\]\s*,\s*(\d+)\s*\)$`)
	roleMaxPattern = regexp.MustCompile(`^roleMax\s*\(\s*"([^"]+)"\s*,\s*(\d+)\s*\)$`)
	rolePrePattern = regexp.MustCompile(`^rolePre\s*\(\s*"([^"]+)"\s*,\s*"([^"]+)"\s*\)$`)
)

func parseRolesArray(rolesStr string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func parseConstraint(key, value string) (*Constraint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (model Model) ValidateConstraints() error { _ = "STUB: not implemented"; return nil }

func (model Model) validateConstraint(constraint *Constraint, groupingPolicy [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func buildUserRoleMap(groupingPolicy [][]string) map[string]map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) validateSOD(constraint *Constraint, groupingPolicy [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) validateSODMax(constraint *Constraint, groupingPolicy [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) validateRoleMax(constraint *Constraint, groupingPolicy [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (model Model) validateRolePre(constraint *Constraint, groupingPolicy [][]string) error {
	_ = "STUB: not implemented"
	return nil
}
