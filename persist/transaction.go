package persist

import "context"

type TransactionalAdapter interface {
	Adapter

	BeginTransaction(ctx context.Context) (TransactionContext, error)
}

type TransactionContext interface {
	Commit() error

	Rollback() error

	GetAdapter() Adapter
}

type PolicyOperation struct {
	Type       OperationType
	Section    string
	PolicyType string
	Rules      [][]string
	OldRules   [][]string
}

type OperationType int

const (
	OperationAdd OperationType = iota

	OperationRemove

	OperationUpdate
)
