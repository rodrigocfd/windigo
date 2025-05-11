//go:build windows

package shell

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
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

// [ITEMIDLIST] struct.
//
// [ITEMIDLIST]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-itemidlist
type ITEMIDLIST uintptr

// Calls [CoTaskMemFree] to release the memory.
//
// [CoTaskMemFree]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cotaskmemfree
func (il ITEMIDLIST) Free() {
	ole.HTASKMEM(il).CoTaskMemFree()
}

// [PROPERTYKEY] struct.
//
// [PROPERTYKEY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ns-wtypes-propertykey
type PROPERTYKEY struct {
	data [20]byte // packed
}

// Returns the fmtId GUID.
func (pk *PROPERTYKEY) FmtId() win.GUID {
	return *(*win.GUID)(unsafe.Pointer(&pk.data[0]))
}

// Sets the fmtId GUID. The pid field will be set to PID_FIRST_USABLE (0x2).
func (pk *PROPERTYKEY) SetFmtId(guid win.GUID) {
	*(*win.GUID)(unsafe.Pointer(&pk.data[0])) = guid
	*(*uint32)(unsafe.Pointer(&pk.data[16])) = wutil.PID_FIRST_USABLE
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
	return wstr.WstrSliceToStr(tb.szTip[:])
}
func (tb *THUMBBUTTON) SetSzTip(val string) {
	wstr.StrToWstrBuf(wstr.SubstrRunes(val, 0, uint(len(tb.szTip)-1)), tb.szTip[:])
}
