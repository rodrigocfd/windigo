//go:build windows

package win

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to an embedded resource, which can be loaded from an
// executable of DLL file.
//
// Created with ResourceInfoLoad().
type ResourceInfo struct {
	resBuf []byte
}

// Reads and stores an embedded resource from an executable or DLL file.
//
// # Example
//
//	resNfo, _ := win.ResourceInfoLoad(win.HINSTANCE(0).GetModuleFileName())
//	verNfo, _ := resNfo.FixedFileInfo()
//	vMaj, vMin, vPat, _ := verNfo.ProductVersion() // like 1.0.0
//
//	blocks := resNfo.Blocks() // each block contains one language
//	productName, _ := blocks[0].ProductName()
//	companyName, _ := blocks[0].CompanyName()
func ResourceInfoLoad(exePath string) (*ResourceInfo, error) {
	resBuf, err := GetFileVersionInfo(exePath)
	if err != nil {
		return nil, err
	}

	return &ResourceInfo{resBuf: resBuf}, nil
}

// Returns the VS_FIXEDFILEINFO struct, which contains version information.
func (me *ResourceInfo) FixedFileInfo() (*VS_FIXEDFILEINFO, bool) {
	ptr, _, ok := VerQueryValue(me.resBuf, "\\")
	if !ok {
		return nil, false
	}

	return (*VS_FIXEDFILEINFO)(ptr), true
}

// Calls ResourceInfo.FixedFileInfo() and automatically retrieves the product
// version, or all zeros if not available.
func (me *ResourceInfo) ProductVersion() (major, minor, patch, build uint16) {
	if verNfo, ok := me.FixedFileInfo(); ok {
		return verNfo.ProductVersion()
	} else {
		return 0, 0, 0, 0
	}
}

// Returns the string information blocks, one per language and code page, which contain several strings.
func (me *ResourceInfo) Blocks() []ResourceInfoBlock {
	type _RawBlock struct {
		langId   LANGID
		codePage co.CP
	}

	if ptr, sz, ok := VerQueryValue(me.resBuf, "\\VarFileInfo\\Translation"); !ok {
		return []ResourceInfoBlock{}
	} else {
		rawBlocks := unsafe.Slice((*_RawBlock)(ptr), sz)
		blocks := make([]ResourceInfoBlock, 0, len(rawBlocks))
		for _, rawBlock := range rawBlocks {
			blocks = append(blocks, ResourceInfoBlock{
				resNfo:   me,
				langId:   rawBlock.langId,
				codePage: rawBlock.codePage,
			})
		}
		return blocks
	}
}

//------------------------------------------------------------------------------

// A block of information retrieved by ResourceInfo.
type ResourceInfoBlock struct {
	resNfo   *ResourceInfo
	langId   LANGID
	codePage co.CP
}

func (me *ResourceInfoBlock) LangId() LANGID  { return me.langId }
func (me *ResourceInfoBlock) CodePage() co.CP { return me.codePage }

func (me *ResourceInfoBlock) strVal(info string) (string, bool) {
	ptr, sz, ok := VerQueryValue(me.resNfo.resBuf,
		fmt.Sprintf("\\StringFileInfo\\%04x%04x\\%s", me.langId, me.codePage, info))
	if !ok {
		return "", false
	}
	return Str.FromNativeSlice(unsafe.Slice((*uint16)(ptr), sz)), true
}

func (me *ResourceInfoBlock) Comments() (string, bool)         { return me.strVal("Comments") }
func (me *ResourceInfoBlock) CompanyName() (string, bool)      { return me.strVal("CompanyName") }
func (me *ResourceInfoBlock) FileDescription() (string, bool)  { return me.strVal("FileDescription") }
func (me *ResourceInfoBlock) FileVersion() (string, bool)      { return me.strVal("FileVersion") }
func (me *ResourceInfoBlock) InternalName() (string, bool)     { return me.strVal("InternalName") }
func (me *ResourceInfoBlock) LegalCopyright() (string, bool)   { return me.strVal("LegalCopyright") }
func (me *ResourceInfoBlock) LegalTrademarks() (string, bool)  { return me.strVal("LegalTrademarks") }
func (me *ResourceInfoBlock) OriginalFilename() (string, bool) { return me.strVal("OriginalFilename") }
func (me *ResourceInfoBlock) ProductName() (string, bool)      { return me.strVal("ProductName") }
func (me *ResourceInfoBlock) ProductVersion() (string, bool)   { return me.strVal("ProductVersion") }
func (me *ResourceInfoBlock) PrivateBuild() (string, bool)     { return me.strVal("PrivateBuild") }
func (me *ResourceInfoBlock) SpecialBuild() (string, bool)     { return me.strVal("SpecialBuild") }
