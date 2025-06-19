//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

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
