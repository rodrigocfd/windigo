//go:build windows

package co

// [VS_FIXEDFILEINFO] DwFileType.
//
// [VS_FIXEDFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VFT uint32

const (
	VFT_UNKNOWN    VFT = 0x0000_0000
	VFT_APP        VFT = 0x0000_0001
	VFT_DLL        VFT = 0x0000_0002
	VFT_DRV        VFT = 0x0000_0003
	VFT_FONT       VFT = 0x0000_0004
	VFT_VXD        VFT = 0x0000_0005
	VFT_STATIC_LIB VFT = 0x0000_0007
)

// [VS_FIXEDFILEINFO] DwFileSubType.
//
// [VS_FIXEDFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VFT2 uint32

const (
	VFT2_UNKNOWN               VFT2 = 0x0000_0000
	VFT2_DRV_PRINTER           VFT2 = 0x0000_0001
	VFT2_DRV_KEYBOARD          VFT2 = 0x0000_0002
	VFT2_DRV_LANGUAGE          VFT2 = 0x0000_0003
	VFT2_DRV_DISPLAY           VFT2 = 0x0000_0004
	VFT2_DRV_MOUSE             VFT2 = 0x0000_0005
	VFT2_DRV_NETWORK           VFT2 = 0x0000_0006
	VFT2_DRV_SYSTEM            VFT2 = 0x0000_0007
	VFT2_DRV_INSTALLABLE       VFT2 = 0x0000_0008
	VFT2_DRV_SOUND             VFT2 = 0x0000_0009
	VFT2_DRV_COMM              VFT2 = 0x0000_000a
	VFT2_DRV_INPUTMETHOD       VFT2 = 0x0000_000b
	VFT2_DRV_VERSIONED_PRINTER VFT2 = 0x0000_000c

	VFT2_FONT_RASTER   VFT2 = 0x0000_0001
	VFT2_FONT_VECTOR   VFT2 = 0x0000_0002
	VFT2_FONT_TRUETYPE VFT2 = 0x0000_0003
)

// [VS_FIXEDFILEINFO] DwFileOS.
//
// [VS_FIXEDFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VOS uint32

const (
	VOS_UNKNOWN VOS = 0x0000_0000
	VOS_DOS     VOS = 0x0001_0000
	VOS_OS216   VOS = 0x0002_0000
	VOS_OS232   VOS = 0x0003_0000
	VOS_NT      VOS = 0x0004_0000
	VOS_WINCE   VOS = 0x0005_0000

	VOS_BASE      VOS = 0x0000_0000
	VOS_WINDOWS16 VOS = 0x0000_0001
	VOS_PM16      VOS = 0x0000_0002
	VOS_PM32      VOS = 0x0000_0003
	VOS_WINDOWS32 VOS = 0x0000_0004

	VOS_DOS_WINDOWS16 VOS = 0x0001_0001
	VOS_DOS_WINDOWS32 VOS = 0x0001_0004
	VOS_OS216_PM16    VOS = 0x0002_0002
	VOS_OS232_PM32    VOS = 0x0003_0003
	VOS_NT_WINDOWS32  VOS = 0x0004_0004
)

// [VS_FIXEDFILEINFO] DwFileFlagsMask and DwFileFlags.
//
// [VS_FIXEDFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VS_FF uint32

const (
	VS_FF_DEBUG        VS_FF = 0x0000_0001
	VS_FF_PRERELEASE   VS_FF = 0x0000_0002
	VS_FF_PATCHED      VS_FF = 0x0000_0004
	VS_FF_PRIVATEBUILD VS_FF = 0x0000_0008
	VS_FF_INFOINFERRED VS_FF = 0x0000_0010
	VS_FF_SPECIALBUILD VS_FF = 0x0000_0020
)
