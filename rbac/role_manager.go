package rbac

type MatchingFunc func(arg1 string, arg2 string) bool

type LinkConditionFunc = func(args ...string) (bool, error)

type RoleManager interface {
	Clear() error

	AddLink(name1 string, name2 string, domain ...string) error

	BuildRelationship(name1 string, name2 string, domain ...string) error

	DeleteLink(name1 string, name2 string, domain ...string) error

	HasLink(name1 string, name2 string, domain ...string) (bool, error)

	GetRoles(name string, domain ...string) ([]string, error)

	GetUsers(name string, domain ...string) ([]string, error)

	GetImplicitRoles(name string, domain ...string) ([]string, error)

	GetImplicitUsers(name string, domain ...string) ([]string, error)

	GetDomains(name string) ([]string, error)

	GetAllDomains() ([]string, error)

	PrintRoles() error

	Match(str string, pattern string) bool

	AddMatchingFunc(name string, fn MatchingFunc)

	AddDomainMatchingFunc(name string, fn MatchingFunc)

	DeleteDomain(domain string) error
}

type ConditionalRoleManager interface {
	RoleManager

	AddLinkConditionFunc(userName, roleName string, fn LinkConditionFunc)

	SetLinkConditionFuncParams(userName, roleName string, params ...string)

	AddDomainLinkConditionFunc(user string, role string, domain string, fn LinkConditionFunc)

	SetDomainLinkConditionFuncParams(user string, role string, domain string, params ...string)
}
