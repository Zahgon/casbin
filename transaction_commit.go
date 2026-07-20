package casbin

import (
	"github.com/casbin/casbin/v3/persist"
)

func (tx *Transaction) Commit() error { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) Rollback() error { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) applyOperationsToDatabase() error { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) applyAddOperationToDatabase(adapter persist.Adapter, op persist.PolicyOperation) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) applyRemoveOperationToDatabase(adapter persist.Adapter, op persist.PolicyOperation) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) applyUpdateOperationToDatabase(adapter persist.Adapter, op persist.PolicyOperation) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Transaction) applyOperationsToModel() error { _ = "STUB: not implemented"; return nil }

func (tx *Transaction) IsCommitted() bool { _ = "STUB: not implemented"; return false }

func (tx *Transaction) IsRolledBack() bool { _ = "STUB: not implemented"; return false }

func (tx *Transaction) IsActive() bool { _ = "STUB: not implemented"; return false }
