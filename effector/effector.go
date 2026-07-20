package effector //nolint:cyclop // TODO

type Effect int

const (
	Allow Effect = iota
	Indeterminate
	Deny
)

type Effector interface {
	MergeEffects(expr string, effects []Effect, matches []float64, policyIndex int, policyLength int) (Effect, int, error)
}
