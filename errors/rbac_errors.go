package errors

import "errors"

var (
	ErrNameNotFound                = errors.New("error: name does not exist")
	ErrDomainParameter             = errors.New("error: domain should be 1 parameter")
	ErrLinkNotFound                = errors.New("error: link between name1 and name2 does not exist")
	ErrUseDomainParameter          = errors.New("error: useDomain should be 1 parameter")
	ErrInvalidFieldValuesParameter = errors.New("fieldValues requires at least one parameter")

	ErrObjCondition   = errors.New("need to meet the prefix required by the object condition")
	ErrEmptyCondition = errors.New("GetAllowedObjectConditions have an empty condition")
)
