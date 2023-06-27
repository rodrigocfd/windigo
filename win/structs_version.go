//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/win/co"
)

// [VS_FIXEDFILEINFO] struct.
//
// [VS_FIXEDFILEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/verrsrc/ns-verrsrc-vs_fixedfileinfo
type VS_FIXEDFILEINFO struct {
	DwSignature        uint32
	DwStrucVersion     uint32
	dwFileVersionMS    uint32
	dwFileVersionLS    uint32
	dwProductVersionMS uint32
	dwProductVersionLS uint32
	DwFileFlagsMask    co.VS_FF
	DwFileFlags        co.VS_FF
	DwFileOS           co.VOS
	DwFileType         co.VFT
	DwFileSubtype      co.VFT2
	dwFileDateMS       uint32
	dwFileDateLS       uint32
}

func (ffi *VS_FIXEDFILEINFO) FileVersion() (major, minor, patch, build uint16) {
	return HIWORD(ffi.dwFileVersionMS), LOWORD(ffi.dwFileVersionMS),
		HIWORD(ffi.dwFileVersionLS), LOWORD(ffi.dwFileVersionLS)
}
func (ffi *VS_FIXEDFILEINFO) SetFileVersion(major, minor, patch, build uint16) {
	ffi.dwFileVersionMS = MAKELONG(minor, major)
	ffi.dwFileVersionLS = MAKELONG(build, patch)
}

func (ffi *VS_FIXEDFILEINFO) ProductVersion() (major, minor, patch, build uint16) {
	return HIWORD(ffi.dwProductVersionMS), LOWORD(ffi.dwProductVersionMS),
		HIWORD(ffi.dwProductVersionLS), LOWORD(ffi.dwProductVersionLS)
}
func (ffi *VS_FIXEDFILEINFO) SetProductVersion(major, minor, patch, build uint16) {
	ffi.dwProductVersionMS = MAKELONG(minor, major)
	ffi.dwProductVersionLS = MAKELONG(build, patch)
}

func (ffi *VS_FIXEDFILEINFO) FileDate() uint64 {
	return util.Make64(ffi.dwFileDateLS, ffi.dwFileDateMS)
}
func (ffi *VS_FIXEDFILEINFO) SetFileDate(val uint64) {
	ffi.dwFileDateLS, ffi.dwFileDateMS = util.Break64(val)
}
