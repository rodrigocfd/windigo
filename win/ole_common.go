//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// A [COM] object whose lifetime can be managed by an [OleReleaser], automating
// the cleanup.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type OleResource interface {
	release()
}

// A [COM] object, derived from [IUnknown].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type OleObj interface {
	OleResource

	// Returns the unique [COM] [interface ID].
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
	IID() co.IID

	// Returns the [COM] virtual table pointer.
	//
	// This is a low-level method, used internally by the library. Incorrect usage
	// may lead to segmentation faults.
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	Ppvt() **_IUnknownVt
}

// Returns the virtual table pointer, performing a nil check.
func ppvtOrNil(obj OleObj) unsafe.Pointer {
	if !utl.IsNil(obj) {
		return unsafe.Pointer(obj.Ppvt())
	}
	return nil
}

// Stores multiple [COM] resources, releasing all them at once.
//
// Every function which returns a COM resource will require an [OleReleaser]
// to manage the object's lifetime.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type OleReleaser struct {
	objs []OleResource
}

// Constructs a new [OleReleaser] to store multiple [COM] resources, releasing
// them all at once.
//
// Every function which returns a COM resource will require an [OleReleaser] to
// manage the object's lifetime.
//
// ⚠️ You must defer [OleReleaser.Release].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func NewOleReleaser() *OleReleaser {
	return new(OleReleaser)
}

// Adds a new [COM] resource to have its lifetime managed by the [OleReleaser].
func (me *OleReleaser) Add(objs ...OleResource) {
	me.objs = append(me.objs, objs...)
}

// Releases all added [COM] resource, in the reverse order they were added.
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *OleReleaser) Release() {
	for i := len(me.objs) - 1; i >= 0; i-- {
		me.objs[i].release()
	}
	me.objs = nil
}

// Releases the specific [COM] resources, if present, immediately.
//
// These objects will be removed from the internal list, thus not being released
// when [OleReleaser.Release] is further called.
func (me *OleReleaser) ReleaseNow(objs ...OleResource) {
NextHisObj:
	for _, hisObj := range objs {
		if utl.IsNil(hisObj) {
			continue // skip nil objects
		}

		for ourIdx, ourObj := range me.objs {
			if ourObj == hisObj { // we found the passed object in our array
				hisObj.release()
				copy(me.objs[ourIdx:len(me.objs)-1], me.objs[ourIdx+1:len(me.objs)]) // move subsequent elements into the gap
				me.objs[len(me.objs)-1] = nil
				me.objs = me.objs[:len(me.objs)-1] // shrink our slice over the same memory
				continue NextHisObj
			}
		}
	}
}
