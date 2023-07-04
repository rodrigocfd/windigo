//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [ID2D1Mesh] COM interface.
//
// [ID2D1Mesh]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1mesh
type ID2D1Mesh interface {
	ID2D1Resource

	// [Open] COM method.
	//
	// ⚠️ You must defer ID2D1TessellationSink.Release().
	//
	// [Open]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1mesh-open
	Open() ID2D1TessellationSink
}

type _ID2D1Mesh struct{ ID2D1Resource }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1Mesh.Release().
func NewID2D1Mesh(base com.IUnknown) ID2D1Mesh {
	return &_ID2D1Mesh{ID2D1Resource: NewID2D1Resource(base)}
}

func (me *_ID2D1Mesh) Open() ID2D1TessellationSink {
	var ppQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1Mesh)(unsafe.Pointer(*me.Ptr())).Open,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewID2D1TessellationSink(com.NewIUnknown(ppQueried))
	} else {
		panic(hr)
	}
}
