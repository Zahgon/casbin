package casbin

func (e *SyncedEnforcer) GetRolesForUser(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetUsersForRole(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) HasRoleForUser(name string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddRoleForUser(user string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddRolesForUser(user string, roles []string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteRoleForUser(user string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteRolesForUser(user string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteUser(user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteRole(role string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeletePermission(permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddPermissionForUser(user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) AddPermissionsForUser(user string, permissions ...[]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeletePermissionForUser(user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeletePermissionsForUser(user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) GetPermissionsForUser(user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetNamedPermissionsForUser(ptype string, user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) HasPermissionForUser(user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) GetImplicitRolesForUser(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetImplicitPermissionsForUser(user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetNamedImplicitPermissionsForUser(ptype string, gtype string, user string, domain ...string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetImplicitUsersForPermission(permission ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedEnforcer) GetImplicitObjectPatternsForUser(user string, domain string, action string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
