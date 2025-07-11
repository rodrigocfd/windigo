//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Stores multiple [COM] resources, releasing all them at once.
//
// Every function which returns a COM resource will require an [OleReleaser]
// to manage the object's lifetime.
//
// # Example
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type OleReleaser struct {
	objs []OleResource
}

// Creates a new [OleReleaser] to store multiple [COM] resources, releasing them
// all at once.
//
// Every function which returns a COM resource will require an [OleReleaser] to
// manage the object's lifetime.
//
// ⚠️ You must defer [OleReleaser.Release].
//
// # Example
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
// # Example
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
	numToRelease := 0
	for _, passedObj := range objs {
		if !utl.IsNil(passedObj) { // obj passed by the user is not nil
			for _, ourObj := range me.objs {
				if ourObj == passedObj { // we found this object in our list
					numToRelease++
				}
			}
		}
	}

	if numToRelease == 0 {
		return // no objects to be released
	}

	newSlice := make([]OleResource, 0, len(me.objs)-numToRelease)
	for _, ourObj := range me.objs {
		for _, passedObj := range objs {
			if passedObj == ourObj {
				ourObj.release()
			} else {
				newSlice = append(newSlice, ourObj)
			}
		}
	}
	me.objs = newSlice
}
