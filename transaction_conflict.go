package casbin

import (
	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
)

type ConflictError struct {
	Operation persist.PolicyOperation
	Reason    string
}

func (e *ConflictError) Error() string { _ = "STUB: not implemented"; return "" }

type ConflictDetector struct {
	baseModel    model.Model
	currentModel model.Model
	operations   []persist.PolicyOperation
}

func NewConflictDetector(baseModel, currentModel model.Model, operations []persist.PolicyOperation) *ConflictDetector {
	_ = "STUB: not implemented"
	return nil
}

func (cd *ConflictDetector) DetectConflicts() error { _ = "STUB: not implemented"; return nil }

func (cd *ConflictDetector) detectRemoveConflict(op persist.PolicyOperation) error {
	_ = "STUB: not implemented"
	return nil
}

func (cd *ConflictDetector) detectUpdateConflict(op persist.PolicyOperation) error {
	_ = "STUB: not implemented"
	return nil
}
