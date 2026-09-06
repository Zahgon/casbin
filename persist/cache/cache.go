package cache

import "errors"

var ErrNoSuchKey = errors.New("there's no such key existing in cache")

type Cache interface {
	Set(key string, value bool, extra ...interface{}) error

	Get(key string) (bool, error)

	Delete(key string) error

	Clear() error
}
