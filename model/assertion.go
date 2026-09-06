package model

import (
	"sync"

	"github.com/casbin/casbin/v3/rbac"
)

type Assertion struct {
	Key             string
	Value           string
	Tokens          []string
	ParamsTokens    []string
	Policy          [][]string
	PolicyMap       map[string]int
	RM              rbac.RoleManager
	CondRM          rbac.ConditionalRoleManager
	FieldIndexMap   map[string]int
	FieldIndexMutex sync.RWMutex
}

func (ast *Assertion) buildIncrementalRoleLinks(rm rbac.RoleManager, op PolicyOp, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ast *Assertion) buildRoleLinks(rm rbac.RoleManager) error {
	_ = "STUB: not implemented"
	return nil
}

func (ast *Assertion) buildIncrementalConditionalRoleLinks(condRM rbac.ConditionalRoleManager, op PolicyOp, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ast *Assertion) buildConditionalRoleLinks(condRM rbac.ConditionalRoleManager) error {
	_ = "STUB: not implemented"
	return nil
}

func (ast *Assertion) addConditionalRoleLink(rule []string, domainRule []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ast *Assertion) copy() *Assertion { _ = "STUB: not implemented"; return nil }
