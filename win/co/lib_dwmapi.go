package co

// DwmSetIconicLivePreviewBitmap() dwSITFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
type DWM_SIT uint32

const (
	DWM_SIT_NONE         DWM_SIT = 0
	DWM_SIT_DISPLAYFRAME DWM_SIT = 0x0000_0001
)

// DWMNCRENDERINGPOLICY enumeration.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmncrenderingpolicy
type DWMNCRP uint32

const (
	DWMNCRP_USEWINDOWSTYLE = iota
	DWMNCRP_DISABLED
	DWMNCRP_ENABLED
)

// DWMFLIP3DWINDOWPOLICY enumeration.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmflip3dwindowpolicy
type DWMFLIP3D uint32

const (
	DWMFLIP3D_DEFAULT = iota
	DWMFLIP3D_EXCLUDEBELOW
	DWMFLIP3D_EXCLUDEABOVE
)

// DWMWINDOWATTRIBUTE get enumeration.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DWMWA_GET uint32

const (
	DWMWA_GET_NCRENDERING_ENABLED   DWMWA_GET = 1
	DWMWA_GET_CAPTION_BUTTON_BOUNDS DWMWA_GET = 5
	DWMWA_GET_EXTENDED_FRAME_BOUNDS DWMWA_GET = 9
	DWMWA_GET_CLOAKED               DWMWA_GET = 14
)

// DWMWINDOWATTRIBUTE set enumeration.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DWMWA_SET uint32

const (
	DWMWA_SET_NCRENDERING_POLICY          DWMWA_SET = 2
	DWMWA_SET_TRANSITIONS_FORCEDISABLED   DWMWA_SET = 3
	DWMWA_SET_ALLOW_NCPAINT               DWMWA_SET = 4
	DWMWA_SET_NONCLIENT_RTL_LAYOUT        DWMWA_SET = 6
	DWMWA_SET_FORCE_ICONIC_REPRESENTATION DWMWA_SET = 7
	DWMWA_SET_FLIP3D_POLICY               DWMWA_SET = 8
	DWMWA_SET_HAS_ICONIC_BITMAP           DWMWA_SET = 10
	DWMWA_SET_DISALLOW_PEEK               DWMWA_SET = 11
	DWMWA_SET_EXCLUDED_FROM_PEEK          DWMWA_SET = 12
	DWMWA_SET_CLOAK                       DWMWA_SET = 13
	DWMWA_SET_FREEZE_REPRESENTATION       DWMWA_SET = 15
)
