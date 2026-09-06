package casbin

import (
	"sync"
	"time"

	"github.com/casbin/casbin/v3/persist/cache"
)

type CachedEnforcer struct {
	*Enforcer
	expireTime  time.Duration
	cache       cache.Cache
	enableCache int32
	locker      *sync.RWMutex
}

type CacheableParam interface {
	GetCacheKey() string
}

func NewCachedEnforcer(params ...interface{}) (*CachedEnforcer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *CachedEnforcer) EnableCache(enableCache bool) { _ = "STUB: not implemented"; return }

func (e *CachedEnforcer) Enforce(rvals ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *CachedEnforcer) LoadPolicy() error { _ = "STUB: not implemented"; return nil }

func (e *CachedEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *CachedEnforcer) RemovePolicies(rules [][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *CachedEnforcer) getCachedResult(key string) (res bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (e *CachedEnforcer) SetExpireTime(expireTime time.Duration) { _ = "STUB: not implemented"; return }

func (e *CachedEnforcer) SetCache(c cache.Cache) { _ = "STUB: not implemented"; return }

func (e *CachedEnforcer) setCachedResult(key string, res bool, extra ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *CachedEnforcer) getKey(params ...interface{}) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (e *CachedEnforcer) InvalidateCache() error { _ = "STUB: not implemented"; return nil }

func GetCacheKey(params ...interface{}) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (e *CachedEnforcer) ClearPolicy() { _ = "STUB: not implemented"; return }
