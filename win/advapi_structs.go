//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
)

// [VALENT] struct.
//
// [VALENT]: https://learn.microsoft.com/en-us/windows/win32/api/winreg/ns-winreg-valentw
type _VALENT struct {
	ValueName *uint16
	ValueLen  uint32
	ValuePtr  uintptr
	Type      co.REG
}

// Returns a projection over src, delimited by ValuePtr and ValueLen fields.
func (v *_VALENT) bufProjection(src []byte) []byte {
	srcPtrVal := uintptr(unsafe.Pointer(unsafe.SliceData(src)))
	offsetIdx := v.ValuePtr - srcPtrVal
	pastIdx := offsetIdx + uintptr(v.ValueLen)
	return src[offsetIdx:pastIdx]
}
