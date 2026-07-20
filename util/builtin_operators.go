package util

import (
	"regexp"
	"sync"

	"github.com/casbin/casbin/v3/rbac"

	"github.com/casbin/govaluate"
)

var (
	keyMatch2Re = regexp.MustCompile(`:[^/]+`)
	keyMatch3Re = regexp.MustCompile(`\{[^/]+\}`)
	keyMatch4Re = regexp.MustCompile(`{([^/]+)}`)
	keyMatch5Re = regexp.MustCompile(`\{[^/]+\}`)
	keyGet2Re1  = regexp.MustCompile(`:[^/]+`)
	keyGet3Re1  = regexp.MustCompile(`\{[^/]+?\}`)
	reCache     = map[string]*regexp.Regexp{}
	reCacheMu   = sync.RWMutex{}
)

func mustCompileOrGet(key string) *regexp.Regexp { _ = "STUB: not implemented"; return nil }

func validateVariadicArgs(expectedLen int, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func validateVariadicStringArgs(expectedLen int, args ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func KeyMatch(key1 string, key2 string) bool { _ = "STUB: not implemented"; return false }

func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyGet(key1, key2 string) string { _ = "STUB: not implemented"; return "" }

func KeyGetFunc(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyMatch2(key1 string, key2 string) bool { _ = "STUB: not implemented"; return false }

func KeyMatch2Func(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyGet2(key1, key2 string, pathVar string) string { _ = "STUB: not implemented"; return "" }

func KeyGet2Func(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyMatch3(key1 string, key2 string) bool { _ = "STUB: not implemented"; return false }

func KeyMatch3Func(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyGet3(key1, key2 string, pathVar string) string { _ = "STUB: not implemented"; return "" }

func KeyGet3Func(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyMatch4(key1 string, key2 string) bool { _ = "STUB: not implemented"; return false }

func KeyMatch4Func(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyMatch5(key1 string, key2 string) bool { _ = "STUB: not implemented"; return false }

func KeyMatch5Func(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RegexMatch(key1 string, key2 string) bool { _ = "STUB: not implemented"; return false }

func RegexMatchFunc(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IPMatch(ip1 string, ip2 string) bool { _ = "STUB: not implemented"; return false }

func IPMatchFunc(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GlobMatch(key1 string, key2 string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func GlobMatchFunc(args ...interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateGFunction(rm rbac.RoleManager, useCache bool) govaluate.ExpressionFunction {
	_ = "STUB: not implemented"
	return *new(govaluate.ExpressionFunction)
}

func GenerateConditionalGFunction(crm rbac.ConditionalRoleManager) govaluate.ExpressionFunction {
	_ = "STUB: not implemented"
	return *new(govaluate.ExpressionFunction)
}

func TimeMatchFunc(args ...string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func TimeMatch(startTime, endTime string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
