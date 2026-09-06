package casbin

const (
	notImplemented = "not implemented"
)

func (e *Enforcer) shouldPersist() bool { _ = "STUB: not implemented"; return false }

func (e *Enforcer) shouldNotify() bool { _ = "STUB: not implemented"; return false }

func (e *Enforcer) validateConstraintsForGroupingPolicy() error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) addPolicyWithoutNotify(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) addPoliciesWithoutNotify(sec string, ptype string, rules [][]string, autoRemoveRepeat bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) removePolicyWithoutNotify(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) updatePolicyWithoutNotify(sec string, ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) updatePoliciesWithoutNotify(sec string, ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) removePoliciesWithoutNotify(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) removeFilteredPolicyWithoutNotify(sec string, ptype string, fieldIndex int, fieldValues []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) updateFilteredPoliciesWithoutNotify(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) addPolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) addPolicies(sec string, ptype string, rules [][]string, autoRemoveRepeat bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) removePolicy(sec string, ptype string, rule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) updatePolicy(sec string, ptype string, oldRule []string, newRule []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) updatePolicies(sec string, ptype string, oldRules [][]string, newRules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) removePolicies(sec string, ptype string, rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) removeFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) updateFilteredPolicies(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) GetFieldIndex(ptype string, field string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *Enforcer) SetFieldIndex(ptype string, field string, index int) {
	_ = "STUB: not implemented"
	return
}
