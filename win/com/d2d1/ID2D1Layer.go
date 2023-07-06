//go:build windows

package d2d1

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/d2d1/d2d1vt"
)

// [ID2D1Layer] COM interface.
//
// [ID2D1Layer]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nn-d2d1-id2d1layer
type ID2D1Layer interface {
	ID2D1Resource

	// [GetSize] COM method.
	//
	// [GetSize]: https://learn.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1layer-getsize
	GetSize() SIZE_F
}

type _ID2D1Layer struct{ ID2D1Resource }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer ID2D1Layer.Release().
func NewID2D1Layer(base com.IUnknown) ID2D1Layer {
	return &_ID2D1Layer{ID2D1Resource: NewID2D1Resource(base)}
}

func (me *_ID2D1Layer) GetSize() SIZE_F {
	ret, _, _ := syscall.SyscallN(
		(*d2d1vt.ID2D1Layer)(unsafe.Pointer(*me.Ptr())).GetSize,
		uintptr(unsafe.Pointer(me.Ptr())))

	lo, hi := util.Break64(uint64(ret))
	return SIZE_F{Width: float32(lo), Height: float32(hi)}
}
