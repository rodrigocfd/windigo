//go:build windows

package ole

import (
	"github.com/rodrigocfd/windigo/internal/utl"
)

// A [COM] object whose lifetime can be managed by an ole.Releaser, automating the
// cleanup.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type ComResource interface {
	// Frees the resources of the object immediately.
	//
	// You usually don't need to call this method directly, since every function
	// which returns a [COM] object will require a Releaser to manage the
	// object's lifetime.
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	Release()
}

// Stores multiple [COM] resources, releasing all them at once.
//
// Every function which returns a COM resource will require a Releaser to manage
// the object's lifetime.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type Releaser struct {
	objs []ComResource
}

// Creates a new [Releaser] to store multiple [COM] resources, releasing them
// all at once.
//
// Every function which returns a COM resource will require a Releaser to manage
// the object's lifetime.
//
// ⚠️ You must defer Releaser.Release().
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func NewReleaser() *Releaser {
	return new(Releaser)
}

// Adds a new [COM] resource to have its lifetime managed by the Releaser.
func (me *Releaser) Add(objs ...ComResource) {
	me.objs = append(me.objs, objs...)
}

// Releases all added [COM] resource, in the reverse order they were added.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
func (me *Releaser) Release() {
	for i := len(me.objs) - 1; i >= 0; i-- {
		me.objs[i].Release()
	}
	me.objs = nil
}

// Releases the specific [COM] resources, if present, immediately.
//
// These objects will be removed from the internal list, thus not being released
// when [Releaser.Release] is further called.
func (me *Releaser) ReleaseNow(objs ...ComResource) {
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

	newSlice := make([]ComResource, 0, len(me.objs)-numToRelease)
	for _, ourObj := range me.objs {
		for _, passedObj := range objs {
			if passedObj == ourObj {
				ourObj.Release()
			} else {
				newSlice = append(newSlice, ourObj)
			}
		}
	}
	me.objs = newSlice
}
