/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"windigo/co"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
	ACCEL struct {
		FVirt co.ACCELF
		Key   co.VK
		Cmd   uint16 // LOWORD(wParam) value
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-compareitemstruct
	COMPAREITEMSTRUCT struct {
		CtlType    co.ODT
		CtlID      uint32
		HwndItem   HWND
		ItemID1    uint32
		ItemData1  uintptr // ULONG_PTR
		ItemID2    uint32
		ItemData2  uintptr // ULONG_PTR
		DwLocaleId uint32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-copydatastruct
	COPYDATASTRUCT struct {
		DwData uintptr // ULONG_PTR
		CbData uint32
		LpData uintptr // PVOID
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-createstructw
	CREATESTRUCT struct {
		LpCreateParams uintptr // LPVOID
		HInstance      HINSTANCE
		HMenu          HMENU
		HwndParent     HWND
		Cy, Cx         int32
		Y, X           int32
		Style          co.WS
		LpszName       *uint16
		LpszClass      *uint16
		ExStyle        co.WS_EX
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-deleteitemstruct
	DELETEITEMSTRUCT struct {
		CtlType  co.ODT
		CtlID    uint32
		ItemID   uint32
		HwndItem HWND
		ItemData uintptr // ULONG_PTR
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-drawitemstruct
	DRAWITEMSTRUCT struct {
		CtlType    co.ODT
		CtlID      uint32
		ItemID     uint32
		ItemAction co.ODA
		ItemState  co.ODS
		HwndItem   HWND
		Hdc        HDC
		RcItem     RECT
		ItemData   uintptr // ULONG_PTR
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-filetime
	FILETIME struct {
		DwLowDateTime  uint32
		DwHighDateTime uint32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/guiddef/ns-guiddef-guid
	GUID struct {
		Data1 uint32
		Data2 uint16
		Data3 uint16
		Data4 uint64
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-helpinfo
	HELPINFO struct {
		CbSize       uint32
		IContextType co.HELPINFO
		ICtrlId      int32
		HItemHandle  HANDLE
		DwContextId  uintptr // DWORD_PTR
		MousePos     POINT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
	LOGFONT struct {
		LfHeight         int32
		LfWidth          int32
		LfEscapement     int32
		LfOrientation    int32
		LfWeight         co.FW
		LfItalic         uint8
		LfUnderline      uint8
		LfStrikeOut      uint8
		LfCharSet        uint8
		LfOutPrecision   uint8
		LfClipPrecision  uint8
		LfQuality        uint8
		LfPitchAndFamily uint8
		LfFaceName       [_LF_FACESIZE]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuinfo
	MENUINFO struct {
		CbSize          uint32
		FMask           co.MIM
		DwStyle         co.MNS
		CyMax           uint32
		HbrBack         HBRUSH
		DwContextHelpID uint32
		DwMenuData      uintptr // ULONG_PTR
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-menuiteminfow
	MENUITEMINFO struct {
		CbSize        uint32
		FMask         co.MIIM
		FType         co.MFT
		FState        co.MFS
		WId           uint32
		HSubMenu      HMENU
		HBmpChecked   HBITMAP
		HBmpUnchecked HBITMAP
		DwItemData    uintptr // ULONG_PTR
		DwTypeData    uintptr // LPWSTR, content changes according to fType
		Cch           uint32
		HBmpItem      HBITMAP
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-monitorinfoexw
	MONITORINFOEX struct {
		CbSize    uint32
		RcMonitor RECT
		RcWork    RECT
		Flags     uint32
		SzDevice  [_CCHDEVICENAME]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-msg
	MSG struct {
		HWnd   HWND
		Msg    uint32
		WParam WPARAM
		LParam LPARAM
		Time   uint32
		Pt     POINT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nmhdr
	NMHDR struct {
		HWndFrom HWND
		IdFrom   uintptr // UINT_PTR, actually it's a simple control ID
		Code     uint32  // in fact it should be int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nonclientmetricsw
	NONCLIENTMETRICS struct {
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

	// https://docs.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-openfilenamew
	OPENFILENAME struct {
		LStructSize       uint32
		HwndOwner         HWND
		HInstance         HINSTANCE
		LpstrFilter       *uint16
		LpstrCustomFilter *uint16
		NMaxCustFilter    uint32
		NFilterIndex      uint32
		LpstrFile         *uint16
		NMaxFile          uint32
		LpstrFileTitle    *uint16
		NMaxFileTitle     uint32
		LpstrInitialDir   *uint16
		LpstrTitle        *uint16
		Flags             co.OFN
		NFileOffset       uint16
		NFileExtension    uint16
		LpstrDefExt       *uint16
		LCustData         LPARAM
		LpfnHook          uintptr // LPOFNHOOKPROC
		LpTemplateName    *uint16
		PvReserved        uintptr // void*
		DwReserved        uint32
		FlagsEx           co.OFN_EX
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
	OSVERSIONINFOEX struct {
		DwOsVersionInfoSize uint32
		DwMajorVersion      uint32
		DwMinorVersion      uint32
		DwBuildNumber       uint32
		DWPlatformId        uint32
		SzCSDVersion        [128]uint16
		WServicePackMajor   uint16
		WServicePackMinor   uint16
		WSuiteMask          uint16
		WProductType        uint8
		WReserve            uint8
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-paintstruct
	PAINTSTRUCT struct {
		Hdc         HDC
		FErase      int32 // BOOL
		RcPaint     RECT
		FRestore    int32 // BOOL
		FIncUpdate  int32 // BOOL
		RgbReserved [32]byte
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
	POINT struct {
		X, Y int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-powerbroadcast_setting
	POWERBROADCAST_SETTING struct {
		PowerSetting GUID
		DataLength   uint32
		Data         [1]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
	RECT struct {
		Left, Top, Right, Bottom int32
	}

	// https://docs.microsoft.com/en-us/previous-versions/windows/desktop/legacy/aa379560(v=vs.85)
	SECURITY_ATTRIBUTES struct {
		NLength              uint32
		LpSecurityDescriptor uintptr // LPVOID
		BInheritHandle       int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/shellapi/ns-shellapi-shfileinfow
	SHFILEINFO struct {
		HIcon         HICON
		IIcon         int32
		DwAttributes  co.SFGAO
		SzDisplayName [_MAX_PATH]uint16
		SzTypeName    [80]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/windef/ns-windef-size
	SIZE struct {
		Cx, Cy int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-systemtime
	SYSTEMTIME struct {
		WYear         uint16
		WMonth        uint16
		WDayOfWeek    uint16
		WDay          uint16
		WHour         uint16
		WMinute       uint16
		WSecond       uint16
		WMilliseconds uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/timezoneapi/ns-timezoneapi-time_zone_information
	TIME_ZONE_INFORMATION struct {
		Bias         int32
		StandardName [32]uint16
		StandardDate SYSTEMTIME
		StandardBias int32
		DaylightName [32]uint16
		DaylightDate SYSTEMTIME
		DaylightBias int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-titlebarinfoex
	TITLEBARINFOEX struct {
		CbSize     uint32
		RcTitleBar RECT
		Rgstate    [_CCHILDREN_TITLEBAR + 1]uint32
		Rgrect     [_CCHILDREN_TITLEBAR + 1]RECT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/minwinbase/ns-minwinbase-win32_find_dataw
	WIN32_FIND_DATA struct {
		DwFileAttributes   co.FILE_ATTRIBUTE
		FtCreationTime     FILETIME
		FtLastAccessTime   FILETIME
		FtLastWriteTime    FILETIME
		NFileSizeHigh      uint32
		NFileSizeLow       uint32
		DwReserved0        uint32
		DwReserved1        uint32
		CFileName          [_MAX_PATH]uint16
		CAlternateFileName [14]uint16
		DwFileType         uint32
		DwCreatorType      uint32
		WFinderFlags       uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-wndclassexw
	WNDCLASSEX struct {
		CbSize        uint32
		Style         co.CS
		LpfnWndProc   uintptr // WNDPROC
		CbClsExtra    int32
		CbWndExtra    int32
		HInstance     HINSTANCE
		HIcon         HICON
		HCursor       HCURSOR
		HbrBackground HBRUSH
		LpszMenuName  *uint16
		LpszClassName *uint16
		HIconSm       HICON
	}
)
