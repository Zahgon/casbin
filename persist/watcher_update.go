package persist

type UpdatableWatcher interface {
	Watcher

	UpdateForUpdatePolicy(sec string, ptype string, oldRule, newRule []string) error

	UpdateForUpdatePolicies(sec string, ptype string, oldRules, newRules [][]string) error
}
