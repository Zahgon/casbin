package cache

import "time"

type cacheItem struct {
	value     bool
	expiresAt time.Time
	ttl       time.Duration
}

type DefaultCache map[string]cacheItem

func (c *DefaultCache) Set(key string, value bool, extra ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *DefaultCache) Get(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *DefaultCache) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (c *DefaultCache) Clear() error { _ = "STUB: not implemented"; return nil }

func NewDefaultCache() (Cache, error) { _ = "STUB: not implemented"; return *new(Cache), nil }
