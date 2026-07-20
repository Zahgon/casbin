package persist

import "github.com/casbin/casbin/v3/model"

type WatcherEx interface {
	Watcher

	UpdateForAddPolicy(sec, ptype string, params ...string) error

	UpdateForRemovePolicy(sec, ptype string, params ...string) error

	UpdateForRemoveFilteredPolicy(sec, ptype string, fieldIndex int, fieldValues ...string) error

	UpdateForSavePolicy(model model.Model) error

	UpdateForAddPolicies(sec string, ptype string, rules ...[]string) error

	UpdateForRemovePolicies(sec string, ptype string, rules ...[]string) error
}
