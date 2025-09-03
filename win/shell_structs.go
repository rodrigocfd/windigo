//go:build windows

package win

import (
	"strconv"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
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
// You can retrieve the ITEMIDLIST of an [IShellItem] with
// [SHGetIDListFromObject].
//
// You can retrieve the [IShellItem] if an ITEMIDLIST with
// [SHCreateItemFromIDList].
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
// Example:
//
//	var nid win.NOTIFYICONDATA
//	nid.SetCbSize()
//
// [NOTIFYICONDATA]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NOTIFYICONDATA struct {
	cbSize           uint32
	HWnd             HWND   // Handle to the window that receives notifications from the icon in the notification area.
	UID              uint32 // Application-defined identifier of the taskbar icon.
	UFlags           co.NIF // Indicate which other members of the struct are valid, also provides other information.
	UCallbackMessage co.WM  // Application-defined message identifier. Sent to the hWnd as an ordinary message.
	HIcon            HICON  // Handle to the icon to be added, modified or deleted.
	szTip            [128]uint16
	DwState          co.NIS // State of the icon.
	DwStateMask      co.NIS // Specifies which bits of DwState are retrieved of modified.
	szInfo           [256]uint16
	uVersion         uint32
	szInfoTitle      [64]uint16
	DwInfoFlags      co.NIIF // Modifies the behavior of the balloon notification.
	GuidItem         GUID    // A registered GUID that identifies the icon, overriding UID. Requires co.NIF_GUID in UFlags.
	HBalloonIcon     HICON   // Handle to the icon of the balloon notification.
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (nid *NOTIFYICONDATA) SetCbSize() {
	nid.cbSize = uint32(unsafe.Sizeof(*nid))
}

// Retrieves the text of the icon tooltip.
func (nid *NOTIFYICONDATA) SzTip() string {
	return wstr.DecodeSlice(nid.szTip[:])
}

// Sets the text of the icon tooltip.
func (nid *NOTIFYICONDATA) SetSzTip(val string) {
	wstr.EncodeToBuf(nid.szTip[:], val)
}

// Retrieves the text displayed in a balloon notification.
func (nid *NOTIFYICONDATA) SzInfo() string {
	return wstr.DecodeSlice(nid.szInfo[:])
}

// Sets the text displayed in a balloon notification.
func (nid *NOTIFYICONDATA) SetSzInfo(val string) {
	wstr.EncodeToBuf(nid.szInfo[:], val)
}

// Retrieves the title of the ballon notification.
func (nid *NOTIFYICONDATA) SzInfoTitle() string {
	return wstr.DecodeSlice(nid.szInfoTitle[:])
}

// Sets the title of the ballon notification.
func (nid *NOTIFYICONDATA) SetSzInfoTitle(val string) {
	wstr.EncodeToBuf(nid.szInfoTitle[:], val)
}

// [NOTIFYICONIDENTIFIER] struct.
//
// ⚠️ You must call [NOTIFYICONIDENTIFIER.SetCbSize] to initialize the struct.
//
// Example:
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
	pId, _ := strconv.ParseUint(string(pkey)[37:], 10, 32)

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
	return wstr.DecodeSlice(shf.szDisplayName[:])
}
func (shf *SHFILEINFO) SetSzDisplayName(val string) {
	wstr.EncodeToBuf(shf.szDisplayName[:], val)
}

func (shf *SHFILEINFO) SzTypeName() string {
	return wstr.DecodeSlice(shf.szTypeName[:])
}
func (shf *SHFILEINFO) SetSzTypeName(val string) {
	wstr.EncodeToBuf(shf.szTypeName[:], val)
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
	return wstr.DecodeSlice(tb.szTip[:])
}
func (tb *THUMBBUTTON) SetSzTip(val string) {
	wstr.EncodeToBuf(tb.szTip[:], val)
}
