package casbin

import (
	"github.com/casbin/casbin/v3/persist"
)

type DistributedEnforcer struct {
	*SyncedEnforcer
}

func NewDistributedEnforcer(params ...interface{}) (*DistributedEnforcer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DistributedEnforcer) SetDispatcher(dispatcher persist.Dispatcher) {
	_ = "STUB: not implemented"
	return
}

func (d *DistributedEnforcer) AddPoliciesSelf(shouldPersist func() bool, sec string, ptype string, rules [][]string) (affected [][]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DistributedEnforcer) RemovePoliciesSelf(shouldPersist func() bool, sec string, ptype string, rules [][]string) (affected [][]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DistributedEnforcer) RemoveFilteredPolicySelf(shouldPersist func() bool, sec string, ptype string, fieldIndex int, fieldValues ...string) (affected [][]string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *DistributedEnforcer) ClearPolicySelf(shouldPersist func() bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *DistributedEnforcer) UpdatePolicySelf(shouldPersist func() bool, sec string, ptype string, oldRule, newRule []string) (affected bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *DistributedEnforcer) UpdatePoliciesSelf(shouldPersist func() bool, sec string, ptype string, oldRules, newRules [][]string) (affected bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *DistributedEnforcer) UpdateFilteredPoliciesSelf(shouldPersist func() bool, sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
