//go:build windows

package ole

import (
	"github.com/rodrigocfd/windigo/internal/utl"
)

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

// Release the specific [COM] resources, if present, immediately. They won't be
// released again when [Releaser.Release] is called.
func (me *Releaser) ReleaseNow(objs ...ComResource) {
	nRelease := 0
	for _, objToRelease := range objs {
		if !utl.IsNil(objToRelease) {
			for _, obj := range me.objs {
				if obj == objToRelease {
					nRelease++
				}
			}
		}
	}

	if nRelease == 0 {
		return
	}

	newSlice := make([]ComResource, 0, nRelease)
	for _, objToRelease := range objs {
		if !utl.IsNil(objToRelease) {
			for _, obj := range me.objs {
				if obj == objToRelease {
					obj.Release()
				} else {
					newSlice = append(newSlice, obj)
				}
			}
		}
	}
	me.objs = newSlice
}
