package co

// VerifyVersionInfo() dwTypeMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
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
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
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

// OSVERSIONINFOEX WSuiteMask. Includes values with VER_NT prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
type VER_SUITE uint16

const (
	VER_SUITE_BACKOFFICE               VER_SUITE = 0x00000004
	VER_SUITE_BLADE                    VER_SUITE = 0x00000400
	VER_SUITE_COMPUTE_SERVER           VER_SUITE = 0x00004000
	VER_SUITE_DATACENTER               VER_SUITE = 0x00000080
	VER_SUITE_ENTERPRISE               VER_SUITE = 0x00000002
	VER_SUITE_EMBEDDEDNT               VER_SUITE = 0x00000040
	VER_SUITE_PERSONAL                 VER_SUITE = 0x00000200
	VER_SUITE_SINGLEUSERTS             VER_SUITE = 0x00000100
	VER_SUITE_SMALLBUSINESS            VER_SUITE = 0x00000001
	VER_SUITE_SMALLBUSINESS_RESTRICTED VER_SUITE = 0x00000020
	VER_SUITE_STORAGE_SERVER           VER_SUITE = 0x00002000
	VER_SUITE_TERMINAL                 VER_SUITE = 0x00000010
	VER_SUITE_WH_SERVER                VER_SUITE = 0x00008000

	VER_SUITE_NT_DOMAIN_CONTROLLER VER_SUITE = 0x0000002
	VER_SUITE_NT_SERVER            VER_SUITE = 0x0000003
	VER_SUITE_NT_WORKSTATION       VER_SUITE = 0x0000001
)

// VS_FIXEDFILEINFO DwFileType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VFT uint32

const (
	VFT_UNKNOWN    VFT = 0x00000000
	VFT_APP        VFT = 0x00000001
	VFT_DLL        VFT = 0x00000002
	VFT_DRV        VFT = 0x00000003
	VFT_FONT       VFT = 0x00000004
	VFT_VXD        VFT = 0x00000005
	VFT_STATIC_LIB VFT = 0x00000007
)

// VS_FIXEDFILEINFO DwFileSubType.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VFT2 uint32

const (
	VFT2_UNKNOWN               VFT2 = 0x00000000
	VFT2_DRV_PRINTER           VFT2 = 0x00000001
	VFT2_DRV_KEYBOARD          VFT2 = 0x00000002
	VFT2_DRV_LANGUAGE          VFT2 = 0x00000003
	VFT2_DRV_DISPLAY           VFT2 = 0x00000004
	VFT2_DRV_MOUSE             VFT2 = 0x00000005
	VFT2_DRV_NETWORK           VFT2 = 0x00000006
	VFT2_DRV_SYSTEM            VFT2 = 0x00000007
	VFT2_DRV_INSTALLABLE       VFT2 = 0x00000008
	VFT2_DRV_SOUND             VFT2 = 0x00000009
	VFT2_DRV_COMM              VFT2 = 0x0000000a
	VFT2_DRV_INPUTMETHOD       VFT2 = 0x0000000b
	VFT2_DRV_VERSIONED_PRINTER VFT2 = 0x0000000c

	VFT2_FONT_RASTER   VFT2 = 0x00000001
	VFT2_FONT_VECTOR   VFT2 = 0x00000002
	VFT2_FONT_TRUETYPE VFT2 = 0x00000003
)

// Virtual key codes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
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

// VS_FIXEDFILEINFO DwFileOS.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VOS uint32

const (
	VOS_UNKNOWN VOS = 0x00000000
	VOS_DOS     VOS = 0x00010000
	VOS_OS216   VOS = 0x00020000
	VOS_OS232   VOS = 0x00030000
	VOS_NT      VOS = 0x00040000
	VOS_WINCE   VOS = 0x00050000

	VOS_BASE      VOS = 0x00000000
	VOS_WINDOWS16 VOS = 0x00000001
	VOS_PM16      VOS = 0x00000002
	VOS_PM32      VOS = 0x00000003
	VOS_WINDOWS32 VOS = 0x00000004

	VOS_DOS_WINDOWS16 VOS = 0x00010001
	VOS_DOS_WINDOWS32 VOS = 0x00010004
	VOS_OS216_PM16    VOS = 0x00020002
	VOS_OS232_PM32    VOS = 0x00030003
	VOS_NT_WINDOWS32  VOS = 0x00040004
)

// VS_FIXEDFILEINFO DwFileFlagsMask and DwFileFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VS_FF uint32

const (
	VS_FF_DEBUG        VS_FF = 0x00000001
	VS_FF_PRERELEASE   VS_FF = 0x00000002
	VS_FF_PATCHED      VS_FF = 0x00000004
	VS_FF_PRIVATEBUILD VS_FF = 0x00000008
	VS_FF_INFOINFERRED VS_FF = 0x00000010
	VS_FF_SPECIALBUILD VS_FF = 0x00000020
)
