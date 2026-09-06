package casbin

import (
	"sync"
	"time"

	"github.com/casbin/casbin/v3/persist/cache"
)

type SyncedCachedEnforcer struct {
	*SyncedEnforcer
	expireTime  time.Duration
	cache       cache.Cache
	enableCache int32
	locker      *sync.RWMutex
}

func NewSyncedCachedEnforcer(params ...interface{}) (*SyncedCachedEnforcer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *SyncedCachedEnforcer) EnableCache(enableCache bool) { _ = "STUB: not implemented"; return }

func (e *SyncedCachedEnforcer) Enforce(rvals ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) LoadPolicy() error { _ = "STUB: not implemented"; return nil }

func (e *SyncedCachedEnforcer) AddPolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) AddPolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) RemovePolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) getCachedResult(key string) (res bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) SetExpireTime(expireTime time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (e *SyncedCachedEnforcer) SetCache(c cache.Cache) { _ = "STUB: not implemented"; return }

func (e *SyncedCachedEnforcer) setCachedResult(key string, res bool, extra ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *SyncedCachedEnforcer) getKey(params ...interface{}) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (e *SyncedCachedEnforcer) InvalidateCache() error { _ = "STUB: not implemented"; return nil }

func (e *SyncedCachedEnforcer) checkOneAndRemoveCache(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *SyncedCachedEnforcer) checkManyAndRemoveCache(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
