package persist

type UpdatableAdapter interface {
	Adapter

	UpdatePolicy(sec string, ptype string, oldRule, newRule []string) error

	UpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error

	UpdateFilteredPolicies(sec string, ptype string, newRules [][]string, fieldIndex int, fieldValues ...string) ([][]string, error)
}
