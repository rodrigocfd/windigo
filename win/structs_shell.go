//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// [NOTIFYICONDATA] struct.
//
// ⚠️ You must call SetCbSize() to initialize the struct.
//
// # Example
//
//	nic := &NOTIFYICONDATA{}
//	nic.SetCbSize()
//
// [NOTIFYICONDATA]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-notifyicondataw
type NOTIFYICONDATA struct {
	cbSize           uint32
	Hwnd             HWND
	UID              uint32
	UFlags           co.NIF
	UCallbackMessage co.WM
	HIcon            HICON
	szTip            [128]uint16
	DwState          co.NIS
	DwStateMask      co.NIS
	szInfo           [256]uint16
	UTimeoutVersion  uint32 // union
	szInfoTitle      [64]uint16
	DwInfoFlags      co.NIIF
	GuidItem         GUID
	HBalloonIcon     HICON
}

func (nid *NOTIFYICONDATA) SetCbSize() { nid.cbSize = uint32(unsafe.Sizeof(*nid)) }

func (nid *NOTIFYICONDATA) SzTip() string { return Str.FromNativeSlice(nid.szTip[:]) }
func (nid *NOTIFYICONDATA) SetSzTip(val string) {
	copy(nid.szTip[:], Str.ToNativeSlice(Str.Substr(val, 0, len(nid.szTip)-1)))
}

func (nid *NOTIFYICONDATA) SzInfo() string { return Str.FromNativeSlice(nid.szInfo[:]) }
func (nid *NOTIFYICONDATA) SetSzInfo(val string) {
	copy(nid.szInfo[:], Str.ToNativeSlice(Str.Substr(val, 0, len(nid.szInfo)-1)))
}

func (nid *NOTIFYICONDATA) SzInfoTitle() string { return Str.FromNativeSlice(nid.szInfoTitle[:]) }
func (nid *NOTIFYICONDATA) SetSzInfoTitle(val string) {
	copy(nid.szInfoTitle[:], Str.ToNativeSlice(Str.Substr(val, 0, len(nid.szInfoTitle)-1)))
}

// [SHFILEINFO] struct.
//
// [SHFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  co.SFGAO
	szDisplayName [_MAX_PATH]uint16
	szTypeName    [80]uint16
}

func (shf *SHFILEINFO) SzDisplayName() string { return Str.FromNativeSlice(shf.szDisplayName[:]) }
func (shf *SHFILEINFO) SetSzDisplayName(val string) {
	copy(shf.szDisplayName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(shf.szDisplayName)-1)))
}

func (shf *SHFILEINFO) SzTypeName() string { return Str.FromNativeSlice(shf.szTypeName[:]) }
func (shf *SHFILEINFO) SetSzTypeName(val string) {
	copy(shf.szTypeName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(shf.szTypeName)-1)))
}
