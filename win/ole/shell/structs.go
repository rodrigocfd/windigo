//go:build windows

package shell

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [COMDLG_FILTERSPEC] struct syntactic sugar.
//
// When the native syscall is made, this struct is converted into the raw
// struct.
//
// [COMDLG_FILTERSPEC]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-comdlg_filterspec
type COMDLG_FILTERSPEC struct {
	Name string
	Spec string
}

// [COMDLG_FILTERSPEC] struct.
//
// [COMDLG_FILTERSPEC]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-comdlg_filterspec
type _COMDLG_FILTERSPEC struct {
	PszName *uint16
	PszSpec *uint16
}

// [THUMBBUTTON] struct.
//
// [THUMBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ns-shobjidl_core-thumbbutton
type THUMBBUTTON struct {
	DwMask  co.THB
	IId     uint32
	IBitmap uint32
	HIcon   win.HICON
	szTip   [260]uint16
	DwFlags co.THBF
}

func (tb *THUMBBUTTON) SzTip() string {
	return wstr.Utf16SliceToStr(tb.szTip[:])
}
func (tb *THUMBBUTTON) SetSzTip(val string) {
	wstr.StrToUtf16(wstr.SubstrRunes(val, 0, uint(len(tb.szTip)-1)), tb.szTip[:])
}
