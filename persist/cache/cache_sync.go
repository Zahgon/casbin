package cache

import (
	"sync"
)

type SyncCache struct {
	cache DefaultCache
	sync.RWMutex
}

func (c *SyncCache) Set(key string, value bool, extra ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SyncCache) Get(key string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *SyncCache) Delete(key string) error { _ = "STUB: not implemented"; return nil }

func (c *SyncCache) Clear() error { _ = "STUB: not implemented"; return nil }

func NewSyncCache() (Cache, error) { _ = "STUB: not implemented"; return *new(Cache), nil }
