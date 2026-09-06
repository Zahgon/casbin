package constant

const (
	ActionIndex   = "act"
	DomainIndex   = "dom"
	SubjectIndex  = "sub"
	ObjectIndex   = "obj"
	PriorityIndex = "priority"
)

const (
	AllowOverrideEffect   = "some(where (p_eft == allow))"
	DenyOverrideEffect    = "!some(where (p_eft == deny))"
	AllowAndDenyEffect    = "some(where (p_eft == allow)) && !some(where (p_eft == deny))"
	PriorityEffect        = "priority(p_eft) || deny"
	SubjectPriorityEffect = "subjectPriority(p_eft) || deny"
)
