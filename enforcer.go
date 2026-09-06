package casbin

import (
	"sync"

	"github.com/casbin/casbin/v3/detector"
	"github.com/casbin/casbin/v3/effector"
	"github.com/casbin/casbin/v3/log"
	"github.com/casbin/casbin/v3/model"
	"github.com/casbin/casbin/v3/persist"
	"github.com/casbin/casbin/v3/rbac"

	"github.com/casbin/govaluate"
)

type Enforcer struct {
	modelPath string
	model     model.Model
	fm        model.FunctionMap
	eft       effector.Effector

	adapter    persist.Adapter
	watcher    persist.Watcher
	dispatcher persist.Dispatcher
	rmMap      map[string]rbac.RoleManager
	condRmMap  map[string]rbac.ConditionalRoleManager
	matcherMap sync.Map
	logger     log.Logger
	detectors  []detector.Detector

	enabled              bool
	autoSave             bool
	autoBuildRoleLinks   bool
	autoNotifyWatcher    bool
	autoNotifyDispatcher bool
	acceptJsonRequest    bool
	gFunctionCache       bool

	aiConfig AIConfig
}

type EnforceContext struct {
	RType string
	PType string
	EType string
	MType string
}

func (e EnforceContext) GetCacheKey() string { _ = "STUB: not implemented"; return "" }

func NewEnforcer(params ...interface{}) (*Enforcer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) InitWithFile(modelPath string, policyPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) InitWithAdapter(modelPath string, adapter persist.Adapter) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) InitWithModelAndAdapter(m model.Model, adapter persist.Adapter) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) initialize() { _ = "STUB: not implemented"; return }

func (e *Enforcer) LoadModel() error { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) GetModel() model.Model { _ = "STUB: not implemented"; return *new(model.Model) }

func (e *Enforcer) SetModel(m model.Model) { _ = "STUB: not implemented"; return }

func (e *Enforcer) GetAdapter() persist.Adapter {
	_ = "STUB: not implemented"
	return *new(persist.Adapter)
}

func (e *Enforcer) SetAdapter(adapter persist.Adapter) { _ = "STUB: not implemented"; return }

func (e *Enforcer) SetWatcher(watcher persist.Watcher) error { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) GetRoleManager() rbac.RoleManager {
	_ = "STUB: not implemented"
	return *new(rbac.RoleManager)
}

func (e *Enforcer) GetNamedRoleManager(ptype string) rbac.RoleManager {
	_ = "STUB: not implemented"
	return *new(rbac.RoleManager)
}

func (e *Enforcer) SetRoleManager(rm rbac.RoleManager) { _ = "STUB: not implemented"; return }

func (e *Enforcer) SetNamedRoleManager(ptype string, rm rbac.RoleManager) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) SetEffector(eft effector.Effector) { _ = "STUB: not implemented"; return }

func (e *Enforcer) SetLogger(logger log.Logger) { _ = "STUB: not implemented"; return }

func (e *Enforcer) SetDetector(d detector.Detector) { _ = "STUB: not implemented"; return }

func (e *Enforcer) SetDetectors(detectors []detector.Detector) { _ = "STUB: not implemented"; return }

func (e *Enforcer) RunDetections() error { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) ClearPolicy() { _ = "STUB: not implemented"; return }

func (e *Enforcer) LoadPolicy() error { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) loadPolicyFromAdapter(baseModel model.Model) (model.Model, error) {
	_ = "STUB: not implemented"
	return *new(model.Model), nil
}

func (e *Enforcer) applyModifiedModel(newModel model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) rebuildRoleLinks(newModel model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) rebuildConditionalRoleLinks(newModel model.Model) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) loadFilteredPolicy(filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) LoadFilteredPolicy(filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) LoadIncrementalFilteredPolicy(filter interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) IsFiltered() bool { _ = "STUB: not implemented"; return false }

func (e *Enforcer) SavePolicy() error { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) getDomainTokens() (rDomainToken, pDomainToken string) {
	_ = "STUB: not implemented"
	return "", ""
}

func (e *Enforcer) registerDomainMatchingFunc(ptype string) { _ = "STUB: not implemented"; return }

func (e *Enforcer) initRmMap() { _ = "STUB: not implemented"; return }

func (e *Enforcer) EnableEnforce(enable bool) { _ = "STUB: not implemented"; return }

func (e *Enforcer) EnableAutoNotifyWatcher(enable bool) { _ = "STUB: not implemented"; return }

func (e *Enforcer) EnableAutoNotifyDispatcher(enable bool) { _ = "STUB: not implemented"; return }

func (e *Enforcer) EnableAutoSave(autoSave bool) { _ = "STUB: not implemented"; return }

func (e *Enforcer) EnableAutoBuildRoleLinks(autoBuildRoleLinks bool) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) EnableAcceptJsonRequest(acceptJsonRequest bool) {
	_ = "STUB: not implemented"
	return
}

func (e *Enforcer) EnableGFunctionCache(enabled bool) { _ = "STUB: not implemented"; return }

func (e *Enforcer) BuildRoleLinks() error { _ = "STUB: not implemented"; return nil }

func (e *Enforcer) BuildIncrementalRoleLinks(op model.PolicyOp, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Enforcer) BuildIncrementalConditionalRoleLinks(op model.PolicyOp, ptype string, rules [][]string) error {
	_ = "STUB: not implemented"
	return nil
}

func NewEnforceContext(suffix string) EnforceContext {
	_ = "STUB: not implemented"
	return *new(EnforceContext)
}

func (e *Enforcer) invalidateMatcherMap() { _ = "STUB: not implemented"; return }

func (e *Enforcer) enforce(matcher string, explains *[]string, rvals ...interface{}) (ok bool, err error) {
	_ = "STUB: not implemented" //nolint:funlen,cyclop,gocyclo // TODO: reduce function complexity
	return false, nil
}

//nolint:nestif // TODO: reduce function complexity

func (e *Enforcer) getAndStoreMatcherExpression(hasEval bool, expString string, functions map[string]govaluate.ExpressionFunction) (*govaluate.EvaluableExpression, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) Enforce(rvals ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) EnforceWithMatcher(matcher string, rvals ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *Enforcer) EnforceEx(rvals ...interface{}) (bool, []string, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e *Enforcer) EnforceExWithMatcher(matcher string, rvals ...interface{}) (bool, []string, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (e *Enforcer) BatchEnforce(requests [][]interface{}) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) BatchEnforceWithMatcher(matcher string, requests [][]interface{}) ([]bool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Enforcer) AddNamedMatchingFunc(ptype, name string, fn rbac.MatchingFunc) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Enforcer) AddNamedDomainMatchingFunc(ptype, name string, fn rbac.MatchingFunc) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Enforcer) AddNamedLinkConditionFunc(ptype, user, role string, fn rbac.LinkConditionFunc) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Enforcer) AddNamedDomainLinkConditionFunc(ptype, user, role string, domain string, fn rbac.LinkConditionFunc) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Enforcer) SetNamedLinkConditionFuncParams(ptype, user, role string, params ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *Enforcer) SetNamedDomainLinkConditionFuncParams(ptype, user, role, domain string, params ...string) bool {
	_ = "STUB: not implemented"
	return false
}

type enforceParameters struct {
	rTokens map[string]int
	rVals   []interface{}

	pTokens map[string]int
	pVals   []string
}

func (p enforceParameters) Get(name string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func generateEvalFunction(functions map[string]govaluate.ExpressionFunction, parameters *enforceParameters) govaluate.ExpressionFunction {
	_ = "STUB: not implemented"
	return *new(govaluate.ExpressionFunction)
}
