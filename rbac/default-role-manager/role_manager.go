package defaultrolemanager

import (
	"sync"

	"github.com/casbin/casbin/v3/rbac"
	"github.com/casbin/casbin/v3/util"
)

const defaultDomain string = ""

type Role struct {
	name                       string
	roles                      *sync.Map
	users                      *sync.Map
	matched                    *sync.Map
	matchedBy                  *sync.Map
	linkConditionFuncMap       *sync.Map
	linkConditionFuncParamsMap *sync.Map
}

func newRole(name string) *Role { _ = "STUB: not implemented"; return nil }

func (r *Role) addRole(role *Role) { _ = "STUB: not implemented"; return }

func (r *Role) removeRole(role *Role) { _ = "STUB: not implemented"; return }

func (r *Role) addUser(user *Role) { _ = "STUB: not implemented"; return }

func (r *Role) removeUser(user *Role) { _ = "STUB: not implemented"; return }

func (r *Role) addMatch(role *Role) { _ = "STUB: not implemented"; return }

func (r *Role) removeMatch(role *Role) { _ = "STUB: not implemented"; return }

func (r *Role) removeMatches() { _ = "STUB: not implemented"; return }

func (r *Role) rangeRoles(fn func(key, value interface{}) bool) { _ = "STUB: not implemented"; return }

func (r *Role) rangeUsers(fn func(key, value interface{}) bool) { _ = "STUB: not implemented"; return }

func (r *Role) getRoles() []string { _ = "STUB: not implemented"; return nil }

func (r *Role) getUsers() []string { _ = "STUB: not implemented"; return nil }

type linkConditionFuncKey struct {
	roleName   string
	domainName string
}

func (r *Role) addLinkConditionFunc(role *Role, domain string, fn rbac.LinkConditionFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Role) getLinkConditionFunc(role *Role, domain string) (rbac.LinkConditionFunc, bool) {
	_ = "STUB: not implemented"
	return *new(rbac.LinkConditionFunc), false
}

func (r *Role) setLinkConditionFuncParams(role *Role, domain string, params ...string) {
	_ = "STUB: not implemented"
	return
}

func (r *Role) getLinkConditionFuncParams(role *Role, domain string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type RoleManagerImpl struct {
	allRoles           *sync.Map
	maxHierarchyLevel  int
	matchingFunc       rbac.MatchingFunc
	domainMatchingFunc rbac.MatchingFunc
	matchingFuncCache  *util.SyncLRUCache
	mutex              sync.Mutex
}

func NewRoleManagerImpl(maxHierarchyLevel int) *RoleManagerImpl {
	_ = "STUB: not implemented"
	return nil
}

func newRoleManagerWithMatchingFunc(maxHierarchyLevel int, fn rbac.MatchingFunc) *RoleManagerImpl {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RoleManagerImpl) rebuild() { _ = "STUB: not implemented"; return }

func (rm *RoleManagerImpl) Match(str string, pattern string) bool {
	_ = "STUB: not implemented"
	return false
}

func (rm *RoleManagerImpl) rangeMatchingRoles(name string, isPattern bool, fn func(role *Role) bool) {
	_ = "STUB: not implemented"
	return
}

func (rm *RoleManagerImpl) load(name interface{}) (value *Role, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (rm *RoleManagerImpl) getRole(name string) (r *Role, created bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func loadAndDelete(m *sync.Map, name string) (value interface{}, loaded bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (rm *RoleManagerImpl) removeRole(name string) { _ = "STUB: not implemented"; return }

func (rm *RoleManagerImpl) AddMatchingFunc(name string, fn rbac.MatchingFunc) {
	_ = "STUB: not implemented"
	return
}

func (rm *RoleManagerImpl) AddDomainMatchingFunc(name string, fn rbac.MatchingFunc) {
	_ = "STUB: not implemented"
	return
}

func (rm *RoleManagerImpl) Clear() error { _ = "STUB: not implemented"; return nil }

func (rm *RoleManagerImpl) AddLink(name1 string, name2 string, domains ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RoleManagerImpl) DeleteLink(name1 string, name2 string, domains ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RoleManagerImpl) HasLink(name1 string, name2 string, domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (rm *RoleManagerImpl) hasLinkHelper(targetName string, roles map[string]*Role, level int) bool {
	_ = "STUB: not implemented"
	return false
}

func (rm *RoleManagerImpl) GetRoles(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RoleManagerImpl) GetUsers(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RoleManagerImpl) GetImplicitRoles(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RoleManagerImpl) GetImplicitUsers(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RoleManagerImpl) getImplicitRolesHelper(roles map[string]*Role, roleSet map[string]bool, res []string, level int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RoleManagerImpl) getImplicitUsersHelper(users map[string]*Role, userSet map[string]bool, res []string, level int) []string {
	_ = "STUB: not implemented"
	return nil
}

func (rm *RoleManagerImpl) PrintRoles() error { _ = "STUB: not implemented"; return nil }

func (rm *RoleManagerImpl) GetDomains(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RoleManagerImpl) GetAllDomains() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rm *RoleManagerImpl) copyFrom(other *RoleManagerImpl) { _ = "STUB: not implemented"; return }

func rangeLinks(users *sync.Map, fn func(name1, name2 string, domain ...string) bool) {
	_ = "STUB: not implemented"
	return
}

func (rm *RoleManagerImpl) Range(fn func(name1, name2 string, domain ...string) bool) {
	_ = "STUB: not implemented"
	return
}

func (rm *RoleManagerImpl) BuildRelationship(name1 string, name2 string, domain ...string) error {
	_ = "STUB: not implemented"
	return nil
}

type DomainManager struct {
	rmMap              *sync.Map
	maxHierarchyLevel  int
	matchingFunc       rbac.MatchingFunc
	domainMatchingFunc rbac.MatchingFunc
	matchingFuncCache  *util.SyncLRUCache
}

func NewDomainManager(maxHierarchyLevel int) *DomainManager { _ = "STUB: not implemented"; return nil }

func (dm *DomainManager) AddMatchingFunc(name string, fn rbac.MatchingFunc) {
	_ = "STUB: not implemented"
	return
}

func (dm *DomainManager) AddDomainMatchingFunc(name string, fn rbac.MatchingFunc) {
	_ = "STUB: not implemented"
	return
}

func (dm *DomainManager) rebuild() { _ = "STUB: not implemented"; return }

func (dm *DomainManager) Clear() error { _ = "STUB: not implemented"; return nil }

func (dm *DomainManager) getDomain(domains ...string) (domain string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (dm *DomainManager) Match(str string, pattern string) bool {
	_ = "STUB: not implemented"
	return false
}

func (dm *DomainManager) rangeAffectedRoleManagers(domain string, fn func(rm *RoleManagerImpl)) {
	_ = "STUB: not implemented"
	return
}

func (dm *DomainManager) load(name interface{}) (value *RoleManagerImpl, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (dm *DomainManager) getRoleManager(domain string, store bool) *RoleManagerImpl {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DomainManager) AddLink(name1 string, name2 string, domains ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DomainManager) DeleteLink(name1 string, name2 string, domains ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DomainManager) HasLink(name1 string, name2 string, domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (dm *DomainManager) GetRoles(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DomainManager) GetUsers(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DomainManager) GetImplicitRoles(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DomainManager) GetImplicitUsers(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DomainManager) PrintRoles() error { _ = "STUB: not implemented"; return nil }

func (dm *DomainManager) GetDomains(name string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DomainManager) GetAllDomains() ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dm *DomainManager) BuildRelationship(name1 string, name2 string, domain ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (dm *DomainManager) DeleteDomain(domain string) error { _ = "STUB: not implemented"; return nil }

type RoleManager struct {
	*DomainManager
}

func NewRoleManager(maxHierarchyLevel int) *RoleManager { _ = "STUB: not implemented"; return nil }

func (rm *RoleManagerImpl) DeleteDomain(domain string) error { _ = "STUB: not implemented"; return nil }

type ConditionalRoleManager struct {
	RoleManagerImpl
}

func (crm *ConditionalRoleManager) copyFrom(other *ConditionalRoleManager) {
	_ = "STUB: not implemented"
	return
}

func newConditionalRoleManagerWithMatchingFunc(maxHierarchyLevel int, fn rbac.MatchingFunc) *ConditionalRoleManager {
	_ = "STUB: not implemented"
	return nil
}

func NewConditionalRoleManager(maxHierarchyLevel int) *ConditionalRoleManager {
	_ = "STUB: not implemented"
	return nil
}

func (crm *ConditionalRoleManager) HasLink(name1 string, name2 string, domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (crm *ConditionalRoleManager) hasLinkHelper(targetName string, roles map[string]*Role, level int, domains ...string) bool {
	_ = "STUB: not implemented"
	return false
}

func (crm *ConditionalRoleManager) getNextRoles(currentRole, nextRole *Role, domains []string, nextRoles map[string]*Role) bool {
	_ = "STUB: not implemented"
	return false
}

func (crm *ConditionalRoleManager) checkLinkCondition(name1, name2 string, domain []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (crm *ConditionalRoleManager) GetRoles(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (crm *ConditionalRoleManager) GetUsers(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (crm *ConditionalRoleManager) GetImplicitRoles(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (crm *ConditionalRoleManager) GetImplicitUsers(name string, domain ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (crm *ConditionalRoleManager) getImplicitRolesHelper(roles map[string]*Role, roleSet map[string]bool, res []string, level int, domains []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (crm *ConditionalRoleManager) getImplicitUsersHelper(users map[string]*Role, userSet map[string]bool, res []string, level int, domains []string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (crm *ConditionalRoleManager) GetLinkConditionFunc(userName, roleName string) (rbac.LinkConditionFunc, bool) {
	_ = "STUB: not implemented"
	return *new(rbac.LinkConditionFunc), false
}

func (crm *ConditionalRoleManager) GetDomainLinkConditionFunc(userName, roleName, domain string) (rbac.LinkConditionFunc, bool) {
	_ = "STUB: not implemented"
	return *new(rbac.LinkConditionFunc), false
}

func (crm *ConditionalRoleManager) GetLinkConditionFuncParams(userName, roleName string, domain ...string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (crm *ConditionalRoleManager) AddLinkConditionFunc(userName, roleName string, fn rbac.LinkConditionFunc) {
	_ = "STUB: not implemented"
	return
}

func (crm *ConditionalRoleManager) AddDomainLinkConditionFunc(userName, roleName, domain string, fn rbac.LinkConditionFunc) {
	_ = "STUB: not implemented"
	return
}

func (crm *ConditionalRoleManager) SetLinkConditionFuncParams(userName, roleName string, params ...string) {
	_ = "STUB: not implemented"
	return
}

func (crm *ConditionalRoleManager) SetDomainLinkConditionFuncParams(userName, roleName, domain string, params ...string) {
	_ = "STUB: not implemented"
	return
}

type ConditionalDomainManager struct {
	ConditionalRoleManager
	DomainManager
}

func NewConditionalDomainManager(maxHierarchyLevel int) *ConditionalDomainManager {
	_ = "STUB: not implemented"
	return nil
}

func (cdm *ConditionalDomainManager) load(name interface{}) (value *ConditionalRoleManager, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cdm *ConditionalDomainManager) getConditionalRoleManager(domain string, store bool) *ConditionalRoleManager {
	_ = "STUB: not implemented"
	return nil
}

func (cdm *ConditionalDomainManager) HasLink(name1 string, name2 string, domains ...string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cdm *ConditionalDomainManager) GetRoles(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdm *ConditionalDomainManager) GetUsers(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdm *ConditionalDomainManager) GetImplicitRoles(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdm *ConditionalDomainManager) GetImplicitUsers(name string, domains ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cdm *ConditionalDomainManager) AddLink(name1 string, name2 string, domains ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cdm *ConditionalDomainManager) DeleteLink(name1 string, name2 string, domains ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cdm *ConditionalDomainManager) AddLinkConditionFunc(userName, roleName string, fn rbac.LinkConditionFunc) {
	_ = "STUB: not implemented"
	return
}

func (cdm *ConditionalDomainManager) AddDomainLinkConditionFunc(userName, roleName, domain string, fn rbac.LinkConditionFunc) {
	_ = "STUB: not implemented"
	return
}

func (cdm *ConditionalDomainManager) SetLinkConditionFuncParams(userName, roleName string, params ...string) {
	_ = "STUB: not implemented"
	return
}

func (cdm *ConditionalDomainManager) SetDomainLinkConditionFuncParams(userName, roleName, domain string, params ...string) {
	_ = "STUB: not implemented"
	return
}

func (cdm *ConditionalDomainManager) AddDomainMatchingFunc(name string, fn rbac.MatchingFunc) {
	_ = "STUB: not implemented"
	return
}

func (cdm *ConditionalDomainManager) rebuild() { _ = "STUB: not implemented"; return }
