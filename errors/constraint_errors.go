package errors

import (
	"errors"
)

var (
	ErrConstraintViolation         = errors.New("constraint violation")
	ErrConstraintParsingError      = errors.New("constraint parsing error")
	ErrConstraintRequiresRBAC      = errors.New("constraints require RBAC to be enabled (role_definition section must exist)")
	ErrInvalidConstraintDefinition = errors.New("invalid constraint definition")
)

type ConstraintViolationError struct {
	ConstraintName string
	Message        string
}

func (e *ConstraintViolationError) Error() string { _ = "STUB: not implemented"; return "" }

func NewConstraintViolationError(constraintName, message string) error {
	_ = "STUB: not implemented"
	return nil
}
