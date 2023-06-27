//go:build windows

package co

// [DWMWA_GET_CLOAKED] return values.
//
// [DWMWA_GET_CLOAKED]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DWM_CLOAKED uint32

const (
	DWM_CLOAKED_APP       DWM_CLOAKED = 0x0000_0001
	DWM_CLOAKED_SHELL     DWM_CLOAKED = 0x0000_0002
	DWM_CLOAKED_INHERITED DWM_CLOAKED = 0x0000_0004
)

// [DwmSetIconicLivePreviewBitmap] dwSITFlags.
//
// [DwmSetIconicLivePreviewBitmap]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmseticoniclivepreviewbitmap
type DWM_SIT uint32

const (
	DWM_SIT_NONE         DWM_SIT = 0
	DWM_SIT_DISPLAYFRAME DWM_SIT = 0x0000_0001
)

// [DWMNCRENDERINGPOLICY] enumeration.
//
// [DWMNCRENDERINGPOLICY]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmncrenderingpolicy
type DWMNCRP uint32

const (
	DWMNCRP_USEWINDOWSTYLE = iota
	DWMNCRP_DISABLED
	DWMNCRP_ENABLED
)

// [DWMFLIP3DWINDOWPOLICY] enumeration.
//
// [DWMFLIP3DWINDOWPOLICY]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmflip3dwindowpolicy
type DWMFLIP3D uint32

const (
	DWMFLIP3D_DEFAULT = iota
	DWMFLIP3D_EXCLUDEBELOW
	DWMFLIP3D_EXCLUDEABOVE
)

// [DWMWINDOWATTRIBUTE] get enumeration.
//
// [DWMWINDOWATTRIBUTE]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DWMWA_GET uint32

const (
	DWMWA_GET_NCRENDERING_ENABLED   DWMWA_GET = 1  // Return type: bool.
	DWMWA_GET_CAPTION_BUTTON_BOUNDS DWMWA_GET = 5  // Return type: win.RECT.
	DWMWA_GET_EXTENDED_FRAME_BOUNDS DWMWA_GET = 9  // Return type: win.RECT.
	DWMWA_GET_CLOAKED               DWMWA_GET = 14 // Return type: co.DWM_CLOAKED.
)

// [DWMWINDOWATTRIBUTE] set enumeration.
//
// [DWMWINDOWATTRIBUTE]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DWMWA_SET uint32

const (
	DWMWA_SET_NCRENDERING_POLICY          DWMWA_SET = 2  // Type of val: co.DWMNCRP.
	DWMWA_SET_TRANSITIONS_FORCEDISABLED   DWMWA_SET = 3  // Type of val: bool.
	DWMWA_SET_ALLOW_NCPAINT               DWMWA_SET = 4  // Type of val: bool.
	DWMWA_SET_NONCLIENT_RTL_LAYOUT        DWMWA_SET = 6  // Type of val: bool.
	DWMWA_SET_FORCE_ICONIC_REPRESENTATION DWMWA_SET = 7  // Type of val: bool.
	DWMWA_SET_FLIP3D_POLICY               DWMWA_SET = 8  // Type of val: co.DWMFLIP3D.
	DWMWA_SET_HAS_ICONIC_BITMAP           DWMWA_SET = 10 // Type of val: bool.
	DWMWA_SET_DISALLOW_PEEK               DWMWA_SET = 11 // Type of val: bool.
	DWMWA_SET_EXCLUDED_FROM_PEEK          DWMWA_SET = 12 // Type of val: bool.
	DWMWA_SET_CLOAK                       DWMWA_SET = 13 // Type of val: (not used).
	DWMWA_SET_FREEZE_REPRESENTATION       DWMWA_SET = 15 // Type of val: (not used).
)
