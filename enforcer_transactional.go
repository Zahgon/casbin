package casbin

import (
	"context"
	"sync"
)

type TransactionalEnforcer struct {
	*Enforcer
	activeTransactions sync.Map
	modelVersion       int64
	commitLock         sync.Mutex
}

func NewTransactionalEnforcer(params ...interface{}) (*TransactionalEnforcer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (te *TransactionalEnforcer) BeginTransaction(ctx context.Context) (*Transaction, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (te *TransactionalEnforcer) GetTransaction(id string) *Transaction {
	_ = "STUB: not implemented"
	return nil
}

func (te *TransactionalEnforcer) IsTransactionActive(id string) bool {
	_ = "STUB: not implemented"
	return false
}

func (te *TransactionalEnforcer) WithTransaction(ctx context.Context, fn func(*Transaction) error) error {
	_ = "STUB: not implemented"
	return nil
}
