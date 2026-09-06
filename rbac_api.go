package casbin

func (e *Enforcer) GetRolesForUser(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetUsersForRole(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) HasRoleForUser(name string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddRolesForUser(user string, roles []string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteRoleForUser(user string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteRolesForUser(user string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteUser(user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeleteRole(role string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeletePermission(permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddPermissionForUser(user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) AddPermissionsForUser(user string, permissions ...[]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeletePermissionForUser(user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) DeletePermissionsForUser(user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) GetPermissionsForUser(user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetNamedPermissionsForUser(ptype string, user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) HasPermissionForUser(user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) GetImplicitRolesForUser(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetNamedImplicitRolesForUser(ptype string, name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetImplicitUsersForRole(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetNamedImplicitPermissionsForUser(ptype string, gtype string, user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetImplicitUsersForPermission(permission ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetDomainsForUser(user string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetImplicitResourcesForUser(user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deepCopyPolicy(src []string) []string { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) GetAllowedObjectConditions(user string, action string, prefix string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeDuplicatePermissions(permissions [][]string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) GetImplicitUsersForResource(resource string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetNamedImplicitUsersForResource(ptype string, resource string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetImplicitUsersForResourceByDomain(resource string, domain string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) GetImplicitObjectPatternsForUser(user string, domain string, action string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) matchDomain(domainIndex int, domain string, rule []string) bool {
	_ = "STUB: not implemented"
	return false
}
