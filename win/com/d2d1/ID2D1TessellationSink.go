//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
)

// [ID2D1TessellationSink] COM interface.
//
// [ID2D1TessellationSink]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1tessellationsink
type ID2D1TessellationSink interface {
	com.IUnknown

	// [Close] COM method.
	//
	// [Close]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1tessellationsink-close
	Close()

	// [AddTriangles] COM method.
	//
	// [AddTriangles]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1tessellationsink-addtriangles
	AddTriangles(triangles []TRIANGLE)
}

type _ID2D1TessellationSink struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1TessellationSink.Release().
func NewID2D1TessellationSink(base com.IUnknown) ID2D1TessellationSink {
	return &_ID2D1TessellationSink{IUnknown: base}
}

func (me *_ID2D1TessellationSink) Close() {
	syscall.SyscallN(
		(*d2d1vt.ID2D1TessellationSink)(unsafe.Pointer(*me.Ptr())).Close,
		uintptr(unsafe.Pointer(me.Ptr())))
}

func (me *_ID2D1TessellationSink) AddTriangles(triangles []TRIANGLE) {
	syscall.SyscallN(
		(*d2d1vt.ID2D1TessellationSink)(unsafe.Pointer(*me.Ptr())).AddTriangles,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&triangles[0])), uintptr(uint32(len(triangles))))
}
