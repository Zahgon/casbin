package persist

type Dispatcher interface {
	AddPolicies(sec string, ptype string, rules [][]string) error

	RemovePolicies(sec string, ptype string, rules [][]string) error

	RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error

	ClearPolicy() error

	UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error

	UpdatePolicies(sec string, ptype string, oldrules, newRules [][]string) error

	UpdateFilteredPolicies(sec string, ptype string, oldRules [][]string, newRules [][]string) error
}
