package casbin

func (e *SyncedEnforcer) GetUsersForRoleInDomain(name string, domain string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedEnforcer) GetRolesForUserInDomain(name string, domain string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedEnforcer) GetPermissionsForUserInDomain(user string, domain string) [][]string {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedEnforcer) AddRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteRoleForUserInDomain(user string, role string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteRolesForUserInDomain(user string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedEnforcer) DeleteDomains(domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
