package rbac

import "context"

type ContextRoleManager interface {
	RoleManager

	ClearCtx(ctx context.Context) error

	AddLinkCtx(ctx context.Context, name1 string, name2 string, domain ...string) error

	DeleteLinkCtx(ctx context.Context, name1 string, name2 string, domain ...string) error

	HasLinkCtx(ctx context.Context, name1 string, name2 string, domain ...string) (bool, error)

	GetRolesCtx(ctx context.Context, name string, domain ...string) ([]string, error)

	GetUsersCtx(ctx context.Context, name string, domain ...string) ([]string, error)

	GetDomainsCtx(ctx context.Context, name string) ([]string, error)

	GetAllDomainsCtx(ctx context.Context) ([]string, error)
}
