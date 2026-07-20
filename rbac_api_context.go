package casbin

import (
	"context"
)

func (e *ContextEnforcer) AddRoleForUserCtx(ctx context.Context, user string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteRoleForUserCtx(ctx context.Context, user string, role string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteRolesForUserCtx(ctx context.Context, user string, domain ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteUserCtx(ctx context.Context, user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeleteRoleCtx(ctx context.Context, role string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeletePermissionCtx(ctx context.Context, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddPermissionForUserCtx(ctx context.Context, user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) AddPermissionsForUserCtx(ctx context.Context, user string, permissions ...[]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeletePermissionForUserCtx(ctx context.Context, user string, permission ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *ContextEnforcer) DeletePermissionsForUserCtx(ctx context.Context, user string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
