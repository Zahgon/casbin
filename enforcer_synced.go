package casbin

import (
	"sync"
	"time"

	"github.com/casbin/govaluate"

	"github.com/casbin/casbin/v3/persist"
	"github.com/casbin/casbin/v3/rbac"
)

type SyncedEnforcer struct {
	*Enforcer
	m               sync.RWMutex
	stopAutoLoad    chan struct{}
	autoLoadRunning int32
}

func NewSyncedEnforcer(params ...interface{}) (*SyncedEnforcer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetLock() *sync.RWMutex { _ = "STUB: not implemented"; return nil }

func (e *SyncedEnforcer) GetRoleManager() rbac.RoleManager {
	_ = "STUB: not implemented"
	return *new(rbac.RoleManager)
}

func (e *SyncedEnforcer) GetNamedRoleManager(ptype string) rbac.RoleManager {
	_ = "STUB: not implemented"
	return *new(rbac.RoleManager)
}

func (e *SyncedEnforcer) SetRoleManager(rm rbac.RoleManager) { _ = "STUB: not implemented"; return }

func (e *SyncedEnforcer) SetNamedRoleManager(ptype string, rm rbac.RoleManager) {
	_ = "STUB: not implemented"
	return
}

func (e *SyncedEnforcer) IsAutoLoadingRunning() bool { _ = "STUB: not implemented"; return false }

func (e *SyncedEnforcer) StartAutoLoadPolicy(d time.Duration) { _ = "STUB: not implemented"; return }

func (e *SyncedEnforcer) StopAutoLoadPolicy() { _ = "STUB: not implemented"; return }

func (e *SyncedEnforcer) SetWatcher(watcher persist.Watcher) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedEnforcer) LoadModel() error { _ = "STUB: not implemented"; return nil }

func (e *SyncedEnforcer) ClearPolicy() { _ = "STUB: not implemented"; return }

func (e *SyncedEnforcer) LoadPolicy() error { _ = "STUB: not implemented"; return nil }

func (e *SyncedEnforcer) LoadFilteredPolicy(filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedEnforcer) LoadIncrementalFilteredPolicy(filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedEnforcer) SavePolicy() error { _ = "STUB: not implemented"; return nil }

func (e *SyncedEnforcer) BuildRoleLinks() error { _ = "STUB: not implemented"; return nil }

func (e *SyncedEnforcer) Enforce(rvals ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) EnforceWithMatcher(matcher string, rvals ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) EnforceEx(rvals ...interface{}) (bool, []string, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e *SyncedEnforcer) EnforceExWithMatcher(matcher string, rvals ...interface{}) (bool, []string, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e *SyncedEnforcer) BatchEnforce(requests [][]interface{}) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) BatchEnforceWithMatcher(matcher string, requests [][]interface{}) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllSubjects() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllNamedSubjects(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllObjects() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllNamedObjects(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllActions() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllNamedActions(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllRoles() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllNamedRoles(ptype string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetAllUsers() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetPolicy() ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetFilteredPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetNamedPolicy(ptype string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetGroupingPolicy() ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetNamedGroupingPolicy(ptype string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) HasPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) HasNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddPoliciesEx(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddNamedPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddNamedPoliciesEx(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdatePolicy(oldPolicy []string, newPolicy []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateNamedPolicy(ptype string, p1 []string, p2 []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdatePolicies(oldPolices [][]string, newPolicies [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateNamedPolicies(ptype string, p1 [][]string, p2 [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateFilteredPolicies(newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateFilteredNamedPolicies(ptype string, newPolicies [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemovePolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveNamedPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveNamedPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveFilteredNamedPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) HasGroupingPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) HasNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddGroupingPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddGroupingPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddGroupingPoliciesEx(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddNamedGroupingPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddNamedGroupingPoliciesEx(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveGroupingPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveGroupingPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveFilteredGroupingPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveNamedGroupingPolicy(ptype string, params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveNamedGroupingPolicies(ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateGroupingPolicy(oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateGroupingPolicies(oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateNamedGroupingPolicy(ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) UpdateNamedGroupingPolicies(ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) RemoveFilteredNamedGroupingPolicy(ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddFunction(name string, function govaluate.ExpressionFunction) {
	_ = "STUB: not implemented"
	return
}

func (e *SyncedEnforcer) SelfAddPolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfAddPolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfAddPoliciesEx(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfRemovePolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfRemovePolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfRemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfUpdatePolicy(sec string, ptype string, oldRule, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) SelfUpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
