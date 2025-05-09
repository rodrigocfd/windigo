//go:build windows

package ole

// Stores multiple [COM] objects, releasing all them at once.
//
// Every function which returns a COM object will require a Releaser to manage the
// object's lifetime.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type Releaser struct {
	objs []Releasable
}

// Creates a new Releaser to stores multiple [COM] objects, releasing all them
// at once.
//
// Every function which returns a COM object will require a Releaser to manage the
// object's lifetime.
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

// Adds a new [COM] object to have its lifetime managed by the Releaser.
func (me *Releaser) Add(objs ...Releasable) {
	me.objs = append(me.objs, objs...)
}

// Releases the resources of all added [COM] objects, in the reverse order they
// were added.
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
