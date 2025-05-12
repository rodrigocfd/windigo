//go:build windows

package utl

import (
	"sync"
	"unsafe"
)

type _PtrCache struct {
	mutex sync.Mutex
	cache map[unsafe.Pointer]struct{}
}

// A global synchronized cache for Go pointers which must not be collected by
// the GC.
var PtrCache _PtrCache

// Synchronously adds a new Go pointer to the cache, preventing GC collection.
func (me *_PtrCache) Add(ptr unsafe.Pointer) {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	if me.cache == nil {
		me.cache = make(map[unsafe.Pointer]struct{})
	}
	me.cache[ptr] = struct{}{}
}

// Synchronously deletes the Go pointer from the cache, allowing GC collection.
func (me *_PtrCache) Delete(ptr unsafe.Pointer) {
	me.mutex.Lock()
	defer me.mutex.Unlock()

	delete(me.cache, ptr)
}
