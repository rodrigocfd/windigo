//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// [CHOOSECOLOR] struct.
//
// ⚠️ You must call SetLStructSize() to initialize the struct.
//
// # Example
//
//	cc := &CHOOSECOLOR{}
//	cc.SetLStructSize()
//
// [CHOOSECOLOR]: https://learn.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-choosecolorw-r1
type CHOOSECOLOR struct {
	lStructSize    uint32
	HwndOwner      HWND
	HInstance      HWND
	RgbResult      COLORREF
	LpCustColors   *COLORREF // Slice must have 16 values.
	Flags          co.CC
	LCustData      uintptr // LPARAM
	LpfnHook       uintptr // LPCCHOOKPROC
	LpTemplateName *uint16
}

func (cc *CHOOSECOLOR) SetLStructSize() { cc.lStructSize = uint32(unsafe.Sizeof(*cc)) }
