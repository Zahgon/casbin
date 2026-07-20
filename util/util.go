package util

import (
	"regexp"
	"sync"
)

var evalReg = regexp.MustCompile(`\beval\((?P<rule>[^)]*)\)`)

var escapeAssertionRegex = regexp.MustCompile(`([()\s|&,=!><+\-*/]|^)((r|p)[0-9]*)\.`)

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EscapeAssertion(s string) string { _ = "STUB: not implemented"; return "" }

func RemoveComments(s string) string { _ = "STUB: not implemented"; return "" }

func ArrayEquals(a []string, b []string) bool { _ = "STUB: not implemented"; return false }

func Array2DEquals(a [][]string, b [][]string) bool { _ = "STUB: not implemented"; return false }

func SortArray2D(arr [][]string) { _ = "STUB: not implemented"; return }

func SortedArray2DEquals(a [][]string, b [][]string) bool { _ = "STUB: not implemented"; return false }

func ArrayRemoveDuplicates(s *[]string) { _ = "STUB: not implemented"; return }

func ArrayToString(s []string) string { _ = "STUB: not implemented"; return "" }

func ParamsToString(s ...string) string { _ = "STUB: not implemented"; return "" }

func SetEquals(a []string, b []string) bool { _ = "STUB: not implemented"; return false }

func SetEqualsInt(a []int, b []int) bool { _ = "STUB: not implemented"; return false }

func Set2DEquals(a [][]string, b [][]string) bool { _ = "STUB: not implemented"; return false }

func JoinSlice(a string, b ...string) []string { _ = "STUB: not implemented"; return nil }

func JoinSliceAny(a string, b ...string) []interface{} { _ = "STUB: not implemented"; return nil }

func SetSubtract(a []string, b []string) []string { _ = "STUB: not implemented"; return nil }

func HasEval(s string) bool { _ = "STUB: not implemented"; return false }

func ReplaceEval(s string, rule string) string { _ = "STUB: not implemented"; return "" }

func ReplaceEvalWithMap(src string, sets map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func GetEvalValue(s string) []string { _ = "STUB: not implemented"; return nil }

func EscapeStringLiterals(expr string) string { _ = "STUB: not implemented"; return "" }

func RemoveDuplicateElement(s []string) []string { _ = "STUB: not implemented"; return nil }

type node struct {
	key   interface{}
	value interface{}
	prev  *node
	next  *node
}

type LRUCache struct {
	capacity int
	m        map[interface{}]*node
	head     *node
	tail     *node
}

func NewLRUCache(capacity int) *LRUCache { _ = "STUB: not implemented"; return nil }

func (cache *LRUCache) remove(n *node, listOnly bool) { _ = "STUB: not implemented"; return }

func (cache *LRUCache) add(n *node, listOnly bool) { _ = "STUB: not implemented"; return }

func (cache *LRUCache) moveToHead(n *node) { _ = "STUB: not implemented"; return }

func (cache *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cache *LRUCache) Put(key interface{}, value interface{}) { _ = "STUB: not implemented"; return }

type SyncLRUCache struct {
	rwm sync.RWMutex
	*LRUCache
}

func NewSyncLRUCache(capacity int) *SyncLRUCache { _ = "STUB: not implemented"; return nil }

func (cache *SyncLRUCache) Get(key interface{}) (value interface{}, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (cache *SyncLRUCache) Put(key interface{}, value interface{}) {
	_ = "STUB: not implemented"
	return
}
