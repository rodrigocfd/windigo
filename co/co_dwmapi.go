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

// [DWMFLIP3DWINDOWPOLICY] enumeration.
//
// [DWMFLIP3DWINDOWPOLICY]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmflip3dwindowpolicy
type DWMFLIP3D uint32

const (
	DWMFLIP3D_DEFAULT = iota
	DWMFLIP3D_EXCLUDEBELOW
	DWMFLIP3D_EXCLUDEABOVE
)

// [DWMNCRENDERINGPOLICY] enumeration.
//
// [DWMNCRENDERINGPOLICY]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmncrenderingpolicy
type DWMNCRP uint32

const (
	DWMNCRP_USEWINDOWSTYLE DWMNCRP = iota
	DWMNCRP_DISABLED
	DWMNCRP_ENABLED
)

// [DWM_SYSTEMBACKDROP_TYPE] enumeration.
//
// [DWM_SYSTEMBACKDROP_TYPE]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwm_systembackdrop_type
type DWMSBT uint32

const (
	DWMSBT_AUTO DWMSBT = iota
	DWMSBT_NONE
	DWMSBT_MAINWINDOW
	DWMSBT_TRANSIENTWINDOW
	DWMSBT_TABBEDWINDOW
)

// [DwmShowContact] showContact.
//
// [DwmShowContact]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/nf-dwmapi-dwmshowcontact
type DWMSC uint32

const (
	DWMSC_NONE      DWMSC = 0x0000_0000
	DWMSC_DOWN      DWMSC = 0x0000_0001
	DWMSC_UP        DWMSC = 0x0000_0002
	DWMSC_DRAG      DWMSC = 0x0000_0004
	DWMSC_HOLD      DWMSC = 0x0000_0008
	DWMSC_PENBARREL DWMSC = 0x0000_0010
	DWMSC_ALL       DWMSC = 0xffff_ffff
)

// [DWMWINDOWATTRIBUTE] enumeration.
//
// [DWMWINDOWATTRIBUTE]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwmwindowattribute
type DWMWA uint32

const (
	DWMWA_NCRENDERING_ENABLED            DWMWA = 1
	DWMWA_NCRENDERING_POLICY             DWMWA = 2
	DWMWA_TRANSITIONS_FORCEDISABLED      DWMWA = 3
	DWMWA_ALLOW_NCPAINT                  DWMWA = 4
	DWMWA_CAPTION_BUTTON_BOUNDS          DWMWA = 5
	DWMWA_NONCLIENT_RTL_LAYOUT           DWMWA = 6
	DWMWA_FORCE_ICONIC_REPRESENTATION    DWMWA = 7
	DWMWA_FLIP3D_POLICY                  DWMWA = 8
	DWMWA_EXTENDED_FRAME_BOUNDS          DWMWA = 9
	DWMWA_HAS_ICONIC_BITMAP              DWMWA = 10
	DWMWA_DISALLOW_PEEK                  DWMWA = 11
	DWMWA_EXCLUDED_FROM_PEEK             DWMWA = 12
	DWMWA_CLOAK                          DWMWA = 13
	DWMWA_CLOAKED                        DWMWA = 14
	DWMWA_FREEZE_REPRESENTATION          DWMWA = 15
	DWMWA_PASSIVE_UPDATE_MODE            DWMWA = 16
	DWMWA_USE_HOSTBACKDROPBRUSH          DWMWA = 17 // Windows 11 Build 22000.
	DWMWA_USE_IMMERSIVE_DARK_MODE        DWMWA = 20 // Windows 11 Build 22000.
	DWMWA_WINDOW_CORNER_PREFERENCE       DWMWA = 33 // Windows 11 Build 22000.
	DWMWA_BORDER_COLOR                   DWMWA = 34 // Windows 11 Build 22000.
	DWMWA_CAPTION_COLOR                  DWMWA = 35 // Windows 11 Build 22000.
	DWMWA_TEXT_COLOR                     DWMWA = 36 // Windows 11 Build 22000.
	DWMWA_VISIBLE_FRAME_BORDER_THICKNESS DWMWA = 37 // Windows 11 Build 22000.
	DWMWA_SYSTEMBACKDROP_TYPE            DWMWA = 38 // Windows 11 Build 22621.
)

// [DWM_WINDOW_CORNER_PREFERENCE] enumeration.
//
// [DWM_WINDOW_CORNER_PREFERENCE]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ne-dwmapi-dwm_window_corner_preference
type DWMWCP uint32

const (
	DWMWCP_DEFAULT DWMWCP = iota
	DWMWCP_DONOTROUND
	DWMWCP_ROUND
	DWMWCP_ROUNDSMALL
)
