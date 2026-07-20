package casbin

func (e *Enforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) GetRolesForUserInDomain(name string, domain string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) GetAllUsersByDomain(domain string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) DeleteAllUsersByDomain(domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteDomains(domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) GetAllDomains() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Enforcer) GetAllRolesByDomain(domain string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
