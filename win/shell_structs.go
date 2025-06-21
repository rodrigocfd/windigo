//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
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

// [ITEMIDLIST] struct.
//
// Implements [OleResource].
//
// [ITEMIDLIST]: https://learn.microsoft.com/en-us/windows/win32/api/shtypes/ns-shtypes-itemidlist
type ITEMIDLIST uintptr

// Implements [OleResource].
func (il *ITEMIDLIST) release() {
	if *il != 0 {
		HTASKMEM(*il).CoTaskMemFree()
		*il = 0
	}
}

// [NOTIFYICONDATA] struct.
//
// ⚠️ You must call [NOTIFYICONDATA.SetCbSize] to initialize the struct.
//
// # Example
//
//	var nid win.NOTIFYICONDATA
//	nid.SetCbSize()
//
// [NOTIFYICONDATA]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NOTIFYICONDATA struct {
	cbSize           uint32
	HWnd             HWND
	UID              uint32
	UFlags           co.NIF
	UCallbackMessage co.WM
	HIcon            HICON
	szTip            [128]uint16
	DwState          co.NIS
	DwStateMask      co.NIS
	szInfo           [256]uint16
	uVersion         uint32
	szInfoTitle      [64]uint16
	DwInfoFlags      co.NIIF
	GuidItem         GUID
	HBalloonIcon     HICON
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (nid *NOTIFYICONDATA) SetCbSize() {
	nid.cbSize = uint32(unsafe.Sizeof(*nid))
}

func (nid *NOTIFYICONDATA) SzTip() string {
	return wstr.WinSliceToGo(nid.szTip[:])
}
func (nid *NOTIFYICONDATA) SetSzTip(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(nid.szTip)-1)), nid.szTip[:])
}

func (nid *NOTIFYICONDATA) SzInfo() string {
	return wstr.WinSliceToGo(nid.szInfo[:])
}
func (nid *NOTIFYICONDATA) SetSzInfo(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(nid.szInfo)-1)), nid.szInfo[:])
}

func (nid *NOTIFYICONDATA) SzInfoTitle() string {
	return wstr.WinSliceToGo(nid.szInfoTitle[:])
}
func (nid *NOTIFYICONDATA) SetSzInfoTitle(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(nid.szInfoTitle)-1)), nid.szInfoTitle[:])
}

// [NOTIFYICONIDENTIFIER] struct.
//
// ⚠️ You must call [NOTIFYICONIDENTIFIER.SetCbSize] to initialize the struct.
//
// # Example
//
//	var nii win.NOTIFYICONIDENTIFIER
//	nii.SetCbSize()
//
// [NOTIFYICONIDENTIFIER]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyiconidentifier
type NOTIFYICONIDENTIFIER struct {
	cbSize   uint32
	HWnd     HWND
	UID      uint32
	GuidItem GUID
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (nii *NOTIFYICONIDENTIFIER) SetCbSize() {
	nii.cbSize = uint32(unsafe.Sizeof(*nii))
}

// [PROPERTYKEY] struct.
//
// [PROPERTYKEY]: https://learn.microsoft.com/en-us/windows/win32/api/wtypes/ns-wtypes-propertykey
type PROPERTYKEY struct {
	data [20]byte // packed
}

// Creates a [PROPERTYKEY] from a string representation.
func PropertykeyFrom(pkey co.PKEY) PROPERTYKEY {
	fmtId := GuidFrom(string(pkey)[0:36])
	pId := wstr.ParseUint(string(pkey)[37:])

	var out PROPERTYKEY
	out.SetFmdId(fmtId)
	out.SetPId(uint32(pId))
	return out
}

func (pk *PROPERTYKEY) FmtId() GUID {
	return *(*GUID)(unsafe.Pointer(&pk.data[0]))
}
func (pk *PROPERTYKEY) SetFmdId(fmtId GUID) {
	*(*GUID)(unsafe.Pointer(&pk.data[0])) = fmtId
}

func (pk *PROPERTYKEY) PId() uint32 {
	return *(*uint32)(unsafe.Pointer(&pk.data[16]))
}
func (pk *PROPERTYKEY) SetPId(pId uint32) {
	*(*uint32)(unsafe.Pointer(&pk.data[16])) = pId
}

// [SHFILEINFO] struct.
//
// [SHFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  co.SFGAO
	szDisplayName [utl.MAX_PATH]uint16
	szTypeName    [80]uint16
}

func (shf *SHFILEINFO) SzDisplayName() string {
	return wstr.WinSliceToGo(shf.szDisplayName[:])
}
func (shf *SHFILEINFO) SetSzDisplayName(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(shf.szDisplayName)-1)), shf.szDisplayName[:])
}

func (shf *SHFILEINFO) SzTypeName() string {
	return wstr.WinSliceToGo(shf.szTypeName[:])
}
func (shf *SHFILEINFO) SetSzTypeName(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(shf.szTypeName)-1)), shf.szTypeName[:])
}

// [THUMBBUTTON] struct.
//
// [THUMBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/ns-shobjidl_core-thumbbutton
type THUMBBUTTON struct {
	DwMask  co.THB
	IId     uint32
	IBitmap uint32
	HIcon   HICON
	szTip   [260]uint16
	DwFlags co.THBF
}

func (tb *THUMBBUTTON) SzTip() string {
	return wstr.WinSliceToGo(tb.szTip[:])
}
func (tb *THUMBBUTTON) SetSzTip(val string) {
	wstr.GoToWinBuf(wstr.SubstrRunes(val, 0, uint(len(tb.szTip)-1)), tb.szTip[:])
}
