package effector

type DefaultEffector struct {
}

func NewDefaultEffector() *DefaultEffector { _ = "STUB: not implemented"; return nil }

func (e *DefaultEffector) MergeEffects(expr string, effects []Effect, matches []float64, policyIndex int, policyLength int) (Effect, int, error) {
	_ = "STUB: not implemented"
	return *new(Effect), 0, nil
}
