package casbin

import (
	"sync"

	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
)

type TransactionBuffer struct {
	operations    []persist.PolicyOperation
	modelSnapshot model.Model
	mutex         sync.RWMutex
}

func NewTransactionBuffer(baseModel model.Model) *TransactionBuffer {
	_ = "STUB: not implemented"
	return nil
}

func (tb *TransactionBuffer) AddOperation(op persist.PolicyOperation) {
	_ = "STUB: not implemented"
	return
}

func (tb *TransactionBuffer) GetOperations() []persist.PolicyOperation {
	_ = "STUB: not implemented"
	return nil
}

func (tb *TransactionBuffer) Clear() { _ = "STUB: not implemented"; return }

func (tb *TransactionBuffer) GetModelSnapshot() model.Model {
	_ = "STUB: not implemented"
	return *new(model.Model)
}

func (tb *TransactionBuffer) ApplyOperationsToModel(baseModel model.Model) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func (tb *TransactionBuffer) HasOperations() bool { _ = "STUB: not implemented"; return false }

func (tb *TransactionBuffer) OperationCount() int { _ = "STUB: not implemented"; return 0 }
