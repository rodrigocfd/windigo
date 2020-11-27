/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package co

// Up-down control messages.
type UDM WM

const (
	UDM_SETRANGE         UDM = UDM(WM_USER + 101)
	UDM_GETRANGE         UDM = UDM(WM_USER + 102)
	UDM_SETPOS           UDM = UDM(WM_USER + 103)
	UDM_GETPOS           UDM = UDM(WM_USER + 104)
	UDM_SETBUDDY         UDM = UDM(WM_USER + 105)
	UDM_GETBUDDY         UDM = UDM(WM_USER + 106)
	UDM_SETACCEL         UDM = UDM(WM_USER + 107)
	UDM_GETACCEL         UDM = UDM(WM_USER + 108)
	UDM_SETBASE          UDM = UDM(WM_USER + 109)
	UDM_GETBASE          UDM = UDM(WM_USER + 110)
	UDM_SETRANGE32       UDM = UDM(WM_USER + 111)
	UDM_GETRANGE32       UDM = UDM(WM_USER + 112)
	UDM_SETUNICODEFORMAT UDM = UDM(CCM_SETUNICODEFORMAT)
	UDM_GETUNICODEFORMAT UDM = UDM(CCM_GETUNICODEFORMAT)
)

// Up-down control notifications, sent via WM_NOTIFY.
type UDN NM

const (
	_UDN_FIRST UDN = -721

	UDN_DELTAPOS UDN = _UDN_FIRST - 1
)

// VerifyVersionInfo() dwTypeMask.
type VER uint32

const (
	VER_BUILDNUMBER      VER = 0x0000004
	VER_MAJORVERSION     VER = 0x0000002
	VER_MINORVERSION     VER = 0x0000001
	VER_PLATFORMID       VER = 0x0000008
	VER_PRODUCT_TYPE     VER = 0x0000080
	VER_SERVICEPACKMAJOR VER = 0x0000020
	VER_SERVICEPACKMINOR VER = 0x0000010
	VER_SUITENAME        VER = 0x0000040
)

// VerifyVersionInfo() dwlConditionMask.
type VER_COND uint8

const (
	VER_COND_EQUAL         VER_COND = 1
	VER_COND_GREATER       VER_COND = 2
	VER_COND_GREATER_EQUAL VER_COND = 3
	VER_COND_LESS          VER_COND = 4
	VER_COND_LESS_EQUAL    VER_COND = 5

	VER_COND_AND VER_COND = 6
	VER_COND_OR  VER_COND = 7
)

// Virtual key codes.
type VK uint16

const (
	VK_LBUTTON             VK = 0x01
	VK_RBUTTON             VK = 0x02
	VK_CANCEL              VK = 0x03
	VK_MBUTTON             VK = 0x04
	VK_XBUTTON1            VK = 0x05
	VK_XBUTTON2            VK = 0x06
	VK_BACK                VK = 0x08
	VK_TAB                 VK = 0x09
	VK_CLEAR               VK = 0x0c
	VK_RETURN              VK = 0x0d
	VK_SHIFT               VK = 0x10
	VK_CONTROL             VK = 0x11
	VK_MENU                VK = 0x12
	VK_PAUSE               VK = 0x13
	VK_CAPITAL             VK = 0x14
	VK_KANA                VK = 0x15
	VK_HANGEUL             VK = 0x15
	VK_HANGUL              VK = 0x15
	VK_JUNJA               VK = 0x17
	VK_FINAL               VK = 0x18
	VK_HANJA               VK = 0x19
	VK_KANJI               VK = 0x19
	VK_ESCAPE              VK = 0x1b
	VK_CONVERT             VK = 0x1c
	VK_NONCONVERT          VK = 0x1d
	VK_ACCEPT              VK = 0x1e
	VK_MODECHANGE          VK = 0x1f
	VK_SPACE               VK = 0x20
	VK_PRIOR               VK = 0x21
	VK_NEXT                VK = 0x22
	VK_END                 VK = 0x23
	VK_HOME                VK = 0x24
	VK_LEFT                VK = 0x25
	VK_UP                  VK = 0x26
	VK_RIGHT               VK = 0x27
	VK_DOWN                VK = 0x28
	VK_SELECT              VK = 0x29
	VK_PRINT               VK = 0x2a
	VK_EXECUTE             VK = 0x2b
	VK_SNAPSHOT            VK = 0x2c
	VK_INSERT              VK = 0x2d
	VK_DELETE              VK = 0x2e
	VK_HELP                VK = 0x2f
	VK_LWIN                VK = 0x5b
	VK_RWIN                VK = 0x5c
	VK_APPS                VK = 0x5d
	VK_SLEEP               VK = 0x5f
	VK_NUMPAD0             VK = 0x60
	VK_NUMPAD1             VK = 0x61
	VK_NUMPAD2             VK = 0x62
	VK_NUMPAD3             VK = 0x63
	VK_NUMPAD4             VK = 0x64
	VK_NUMPAD5             VK = 0x65
	VK_NUMPAD6             VK = 0x66
	VK_NUMPAD7             VK = 0x67
	VK_NUMPAD8             VK = 0x68
	VK_NUMPAD9             VK = 0x69
	VK_MULTIPLY            VK = 0x6a
	VK_ADD                 VK = 0x6b
	VK_SEPARATOR           VK = 0x6c
	VK_SUBTRACT            VK = 0x6d
	VK_DECIMAL             VK = 0x6e
	VK_DIVIDE              VK = 0x6f
	VK_F1                  VK = 0x70
	VK_F2                  VK = 0x71
	VK_F3                  VK = 0x72
	VK_F4                  VK = 0x73
	VK_F5                  VK = 0x74
	VK_F6                  VK = 0x75
	VK_F7                  VK = 0x76
	VK_F8                  VK = 0x77
	VK_F9                  VK = 0x78
	VK_F10                 VK = 0x79
	VK_F11                 VK = 0x7a
	VK_F12                 VK = 0x7b
	VK_F13                 VK = 0x7c
	VK_F14                 VK = 0x7d
	VK_F15                 VK = 0x7e
	VK_F16                 VK = 0x7f
	VK_F17                 VK = 0x80
	VK_F18                 VK = 0x81
	VK_F19                 VK = 0x82
	VK_F20                 VK = 0x83
	VK_F21                 VK = 0x84
	VK_F22                 VK = 0x85
	VK_F23                 VK = 0x86
	VK_F24                 VK = 0x87
	VK_NUMLOCK             VK = 0x90
	VK_SCROLL              VK = 0x91
	VK_OEM_NEC_EQUAL       VK = 0x92
	VK_OEM_FJ_JISHO        VK = 0x92
	VK_OEM_FJ_MASSHOU      VK = 0x93
	VK_OEM_FJ_TOUROKU      VK = 0x94
	VK_OEM_FJ_LOYA         VK = 0x95
	VK_OEM_FJ_ROYA         VK = 0x96
	VK_LSHIFT              VK = 0xa0
	VK_RSHIFT              VK = 0xa1
	VK_LCONTROL            VK = 0xa2
	VK_RCONTROL            VK = 0xa3
	VK_LMENU               VK = 0xa4
	VK_RMENU               VK = 0xa5
	VK_BROWSER_BACK        VK = 0xa6
	VK_BROWSER_FORWARD     VK = 0xa7
	VK_BROWSER_REFRESH     VK = 0xa8
	VK_BROWSER_STOP        VK = 0xa9
	VK_BROWSER_SEARCH      VK = 0xaa
	VK_BROWSER_FAVORITES   VK = 0xab
	VK_BROWSER_HOME        VK = 0xac
	VK_VOLUME_MUTE         VK = 0xad
	VK_VOLUME_DOWN         VK = 0xae
	VK_VOLUME_UP           VK = 0xaf
	VK_MEDIA_NEXT_TRACK    VK = 0xb0
	VK_MEDIA_PREV_TRACK    VK = 0xb1
	VK_MEDIA_STOP          VK = 0xb2
	VK_MEDIA_PLAY_PAUSE    VK = 0xb3
	VK_LAUNCH_MAIL         VK = 0xb4
	VK_LAUNCH_MEDIA_SELECT VK = 0xb5
	VK_LAUNCH_APP1         VK = 0xb6
	VK_LAUNCH_APP2         VK = 0xb7
	VK_OEM_1               VK = 0xba
	VK_OEM_PLUS            VK = 0xbb
	VK_OEM_COMMA           VK = 0xbc
	VK_OEM_MINUS           VK = 0xbd
	VK_OEM_PERIOD          VK = 0xbe
	VK_OEM_2               VK = 0xbf
	VK_OEM_3               VK = 0xc0
	VK_OEM_4               VK = 0xdb
	VK_OEM_5               VK = 0xdc
	VK_OEM_6               VK = 0xdd
	VK_OEM_7               VK = 0xde
	VK_OEM_8               VK = 0xdf
	VK_OEM_AX              VK = 0xe1
	VK_OEM_102             VK = 0xe2
	VK_ICO_HELP            VK = 0xe3
	VK_ICO_00              VK = 0xe4
	VK_PROCESSKEY          VK = 0xe5
	VK_ICO_CLEAR           VK = 0xe6
	VK_PACKET              VK = 0xe7
	VK_OEM_RESET           VK = 0xe9
	VK_OEM_JUMP            VK = 0xea
	VK_OEM_PA1             VK = 0xeb
	VK_OEM_PA2             VK = 0xec
	VK_OEM_PA3             VK = 0xed
	VK_OEM_WSCTRL          VK = 0xee
	VK_OEM_CUSEL           VK = 0xef
	VK_OEM_ATTN            VK = 0xf0
	VK_OEM_FINISH          VK = 0xf1
	VK_OEM_COPY            VK = 0xf2
	VK_OEM_AUTO            VK = 0xf3
	VK_OEM_ENLW            VK = 0xf4
	VK_OEM_BACKTAB         VK = 0xf5
	VK_ATTN                VK = 0xf6
	VK_CRSEL               VK = 0xf7
	VK_EXSEL               VK = 0xf8
	VK_EREOF               VK = 0xf9
	VK_PLAY                VK = 0xfa
	VK_ZOOM                VK = 0xfb
	VK_NONAME              VK = 0xfc
	VK_PA1                 VK = 0xfd
	VK_OEM_CLEAR           VK = 0xfe
)

// https://docs.microsoft.com/en-us/windows/win32/controls/parts-and-states
type VS_PART int32

const ( // list view styles
	VS_PART_LVP_LISTITEM         VS_PART = 1
	VS_PART_LVP_LISTGROUP        VS_PART = 2
	VS_PART_LVP_LISTDETAIL       VS_PART = 3
	VS_PART_LVP_LISTSORTEDDETAIL VS_PART = 4
	VS_PART_LVP_EMPTYTEXT        VS_PART = 5
	VS_PART_LVP_GROUPHEADER      VS_PART = 6
	VS_PART_LVP_GROUPHEADERLINE  VS_PART = 7
	VS_PART_LVP_EXPANDBUTTON     VS_PART = 8
	VS_PART_LVP_COLLAPSEBUTTON   VS_PART = 9
	VS_PART_LVP_COLUMNDETAIL     VS_PART = 10
)

// https://docs.microsoft.com/en-us/windows/win32/controls/parts-and-states
type VS_STATE int32

const ( // list view states
	VS_STATE_NONE VS_STATE = 0

	VS_STATE_LISS_NORMAL           VS_STATE = 1
	VS_STATE_LISS_HOT              VS_STATE = 2
	VS_STATE_LISS_SELECTED         VS_STATE = 3
	VS_STATE_LISS_DISABLED         VS_STATE = 4
	VS_STATE_LISS_SELECTEDNOTFOCUS VS_STATE = 5
	VS_STATE_LISS_HOTSELECTED      VS_STATE = 6

	VS_STATE_LVGH_OPEN                       VS_STATE = 1
	VS_STATE_LVGH_OPENHOT                    VS_STATE = 2
	VS_STATE_LVGH_OPENSELECTED               VS_STATE = 3
	VS_STATE_LVGH_OPENSELECTEDHOT            VS_STATE = 4
	VS_STATE_LVGH_OPENSELECTEDNOTFOCUSED     VS_STATE = 5
	VS_STATE_LVGH_OPENSELECTEDNOTFOCUSEDHOT  VS_STATE = 6
	VS_STATE_LVGH_OPENMIXEDSELECTION         VS_STATE = 7
	VS_STATE_LVGH_OPENMIXEDSELECTIONHOT      VS_STATE = 8
	VS_STATE_LVGH_CLOSE                      VS_STATE = 9
	VS_STATE_LVGH_CLOSEHOT                   VS_STATE = 10
	VS_STATE_LVGH_CLOSESELECTED              VS_STATE = 11
	VS_STATE_LVGH_CLOSESELECTEDHOT           VS_STATE = 12
	VS_STATE_LVGH_CLOSESELECTEDNOTFOCUSED    VS_STATE = 13
	VS_STATE_LVGH_CLOSESELECTEDNOTFOCUSEDHOT VS_STATE = 14
	VS_STATE_LVGH_CLOSEMIXEDSELECTION        VS_STATE = 15
	VS_STATE_LVGH_CLOSEMIXEDSELECTIONHOT     VS_STATE = 16

	VS_STATE_LVGHL_OPEN                       VS_STATE = 1
	VS_STATE_LVGHL_OPENHOT                    VS_STATE = 2
	VS_STATE_LVGHL_OPENSELECTED               VS_STATE = 3
	VS_STATE_LVGHL_OPENSELECTEDHOT            VS_STATE = 4
	VS_STATE_LVGHL_OPENSELECTEDNOTFOCUSED     VS_STATE = 5
	VS_STATE_LVGHL_OPENSELECTEDNOTFOCUSEDHOT  VS_STATE = 6
	VS_STATE_LVGHL_OPENMIXEDSELECTION         VS_STATE = 7
	VS_STATE_LVGHL_OPENMIXEDSELECTIONHOT      VS_STATE = 8
	VS_STATE_LVGHL_CLOSE                      VS_STATE = 9
	VS_STATE_LVGHL_CLOSEHOT                   VS_STATE = 10
	VS_STATE_LVGHL_CLOSESELECTED              VS_STATE = 11
	VS_STATE_LVGHL_CLOSESELECTEDHOT           VS_STATE = 12
	VS_STATE_LVGHL_CLOSESELECTEDNOTFOCUSED    VS_STATE = 13
	VS_STATE_LVGHL_CLOSESELECTEDNOTFOCUSEDHOT VS_STATE = 14
	VS_STATE_LVGHL_CLOSEMIXEDSELECTION        VS_STATE = 15
	VS_STATE_LVGHL_CLOSEMIXEDSELECTIONHOT     VS_STATE = 16

	VS_STATE_LVEB_NORMAL VS_STATE = 1
	VS_STATE_LVEB_HOVER  VS_STATE = 2
	VS_STATE_LVEB_PUSHED VS_STATE = 3

	VS_STATE_LVCB_NORMAL VS_STATE = 1
	VS_STATE_LVCB_HOVER  VS_STATE = 2
	VS_STATE_LVCB_PUSHED VS_STATE = 3
)
