//go:build windows

package win

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Version information from an EXE or DLL, loaded with [GetFileVersionInfo] and
// [VerQueryValue].
//
// Example:
//
//	hInst, _ := win.GetModuleHandle("")
//	exeName, _ := hInst.GetModuleFileName()
//	info, _ := win.VersionLoad(exeName)
type VersionInfo struct {
	Version          [4]uint16
	LangId           LANGID
	CodePage         co.CP
	Comments         string
	CompanyName      string
	FileDescription  string
	FileVersion      string
	InternalName     string
	LegalCopyright   string
	LegalTrademarks  string
	OriginalFilename string
	ProductName      string
	ProductVersion   string
	PrivateBuild     string
	SpecialBuild     string
}

// Loads the embedded version information from an EXE or DLL, with
// [GetFileVersionInfo] and [VerQueryValue].
//
// Example:
//
//	hInst, _ := win.GetModuleHandle("")
//	exeName, _ := hInst.GetModuleFileName()
//	info, _ := win.VersionLoad(exeName)
func VersionLoad(moduleName string) (VersionInfo, error) {
	szData, err := GetFileVersionInfoSize(moduleName)
	if err != nil {
		return VersionInfo{}, fmt.Errorf("VersionLoad GetFileVersionInfoSize: %w", err)
	}

	data := NewVecSized(szData, byte(0))
	defer data.Free()

	if err := GetFileVersionInfo(moduleName, data.HotSlice()); err != nil {
		return VersionInfo{}, fmt.Errorf("VersionLoad GetFileVersionInfo: %w", err)
	}

	var v VersionInfo // to be returned

	if pNfoRaw, _, ok := VerQueryValue(data.HotSlice(), "\\"); ok {
		pNfo := (*VS_FIXEDFILEINFO)(pNfoRaw)
		v.Version[0], v.Version[1], v.Version[2], v.Version[3] = pNfo.FileVersion()
	}

	type Block struct {
		LangId   LANGID
		CodePage co.CP
	}

	if pBlocks, count, ok := VerQueryValue(
		data.HotSlice(), "\\VarFileInfo\\Translation"); ok && count > 0 {

		blocks := unsafe.Slice((*Block)(pBlocks), count)
		bl := blocks[0] // we'll load the 1st block only, which is the most common case

		v.LangId = bl.LangId
		v.CodePage = bl.CodePage

		getStr := func(id string) string {
			if pStr, nChars, ok := VerQueryValue(data.HotSlice(),
				fmt.Sprintf("\\StringFileInfo\\%04x%04x\\%s",
					bl.LangId, bl.CodePage, id)); ok {

				wideStr := unsafe.Slice((*uint16)(pStr), nChars-1) // don't include terminating null
				return wstr.DecodeSlice(wideStr)
			}
			return ""
		}

		v.Comments = getStr("Comments")
		v.CompanyName = getStr("CompanyName")
		v.FileDescription = getStr("FileDescription")
		v.FileVersion = getStr("FileVersion")
		v.InternalName = getStr("InternalName")
		v.LegalCopyright = getStr("LegalCopyright")
		v.LegalTrademarks = getStr("LegalTrademarks")
		v.OriginalFilename = getStr("OriginalFilename")
		v.ProductName = getStr("ProductName")
		v.ProductVersion = getStr("ProductVersion")
		v.PrivateBuild = getStr("PrivateBuild")
		v.SpecialBuild = getStr("SpecialBuild")
	}

	return v, nil
}

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
	return utl.Make64(ffi.dwFileDateLS, ffi.dwFileDateMS)
}
func (ffi *VS_FIXEDFILEINFO) SetFileDate(val uint64) {
	ffi.dwFileDateLS, ffi.dwFileDateMS = utl.Break64(val)
}
