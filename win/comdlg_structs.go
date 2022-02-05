package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// ⚠️ You must call SetLStructSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-choosecolorw-r1
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
