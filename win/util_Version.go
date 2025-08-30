//go:build windows

package win

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
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
		return VersionInfo{}, err
	}

	data := NewVecSized(szData, byte(0))
	defer data.Free()

	if err := GetFileVersionInfo(moduleName, data.HotSlice()); err != nil {
		return VersionInfo{}, err
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
