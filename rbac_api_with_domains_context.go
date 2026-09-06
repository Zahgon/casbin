package casbin

import (
	"context"
)

func (e *ContextEnforcer) AddRoleForUserInDomainCtx(ctx context.Context, user string, role string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteRoleForUserInDomainCtx(ctx context.Context, user string, role string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteRolesForUserInDomainCtx(ctx context.Context, user string, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteAllUsersByDomainCtx(ctx context.Context, domain string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteDomainsCtx(ctx context.Context, domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
