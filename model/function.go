package model

import (
	"sync"

	"github.com/casbin/govaluate"
)

type FunctionMap struct {
	fns *sync.Map
}

func (fm *FunctionMap) AddFunction(name string, function govaluate.ExpressionFunction) {
	_ = "STUB: not implemented"
	return
}

func LoadFunctionMap() FunctionMap { _ = "STUB: not implemented"; return *new(FunctionMap) }

func (fm *FunctionMap) GetFunctions() map[string]govaluate.ExpressionFunction {
	_ = "STUB: not implemented"
	return nil
}
