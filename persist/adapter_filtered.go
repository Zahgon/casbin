package persist

import (
	"github.com/casbin/casbin/v3/model"
)

type FilteredAdapter interface {
	Adapter

	LoadFilteredPolicy(model model.Model, filter interface{}) error

	IsFiltered() bool
}
