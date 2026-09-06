package persist

type BatchAdapter interface {
	Adapter

	AddPolicies(sec string, ptype string, rules [][]string) error

	RemovePolicies(sec string, ptype string, rules [][]string) error
}
