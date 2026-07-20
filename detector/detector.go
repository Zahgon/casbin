package detector

import "github.com/casbin/casbin/v3/rbac"

type Detector interface {
	Check(rm rbac.RoleManager) error
}
