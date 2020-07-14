/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"wingows/co"
)

type ACCEL struct {
	FVirt co.ACCELF
	Key   co.VK
	Cmd   uint16 // LOWORD(wParam) value
}

type CREATESTRUCT struct {
	LpCreateParams uintptr
	HInstance      HINSTANCE
	HMenu          HMENU
	HwndParent     HWND
	Cy, Cx         uint32 // actually int32
	Y, X           int32
	Style          co.WS
	LpszName       uintptr // LPCWSTR
	LpszClass      uintptr // LPCWSTR
	ExStyle        co.WS_EX
}

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 uint64
}

type HELPINFO struct {
	CbSize       uint32
	IContextType co.HELPINFO
	ICtrlId      int32
	HItemHandle  HANDLE
	DwContextId  uintptr
	MousePos     POINT
}

type MENUINFO struct {
	CbSize          uint32
	FMask           co.MIM
	DwStyle         co.MNS
	CyMax           uint32
	HbrBack         HBRUSH
	DwContextHelpID uint32
	DwMenuData      uintptr
}

type MENUITEMINFO struct {
	CbSize        uint32
	FMask         co.MIIM
	FType         co.MFT
	FState        co.MFS
	WId           uint32
	HSubMenu      HMENU
	HBmpChecked   HBITMAP
	HBmpUnchecked HBITMAP
	DwItemData    uintptr
	DwTypeData    uintptr // LPWSTR
	Cch           uint32
	HBmpItem      HBITMAP
}

type MONITORINFOEX struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	Flags     uint32
	SzDevice  [32]uint16 // CCHDEVICENAME
}

type NMHDR struct {
	HWndFrom HWND
	IdFrom   uintptr
	Code     uint32 // in fact it should be int32
}

type NONCLIENTMETRICS struct {
	CbSize             uint32
	IBorderWidth       int32
	IScrollWidth       int32
	IScrollHeight      int32
	ICaptionWidth      int32
	ICaptionHeight     int32
	LfCaptionFont      LOGFONT
	ISmCaptionWidth    int32
	ISmCaptionHeight   int32
	LfSmCaptionFont    LOGFONT
	IMenuWidth         int32
	IMenuHeight        int32
	LfMenuFont         LOGFONT
	LfStatusFont       LOGFONT
	LfMessageFont      LOGFONT
	IPaddedBorderWidth int32
}

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type SECURITY_ATTRIBUTES struct {
	NLength              uint32
	LpSecurityDescriptor uintptr
	BInheritHandle       int32
}

type SHFILEINFO struct {
	HIcon         HICON
	IIcon         int32
	DwAttributes  co.SFGAO
	SzDisplayName [260]uint16 // MAX_PATH
	SzTypeName    [80]uint16
}

type SIZE struct {
	Cx, Cy int32
}

type WIN32_FIND_DATA struct {
	DwFileAttributes   co.FILE_ATTRIBUTE
	FtCreationTime     FILETIME
	FtLastAccessTime   FILETIME
	FtLastWriteTime    FILETIME
	NFileSizeHigh      uint32
	NFileSizeLow       uint32
	DwReserved0        uint32
	DwReserved1        uint32
	CFileName          [260]uint16 // MAX_PATH
	CAlternateFileName [14]uint16
	DwFileType         uint32
	DwCreatorType      uint32
	WFinderFlags       uint16
}
