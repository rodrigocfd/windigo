package win

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// High-level abstraction to an embedded resource, which can be loaded from an
// executable of DLL file.
//
// Created with LoadResourceInfo().
type ResourceInfo interface {
	FixedFileInfo() (*VS_FIXEDFILEINFO, bool)
	Blocks() []_ResourceInfoBlock

	Comments(langId LANGID, codePage co.CP) (string, bool)
	CompanyName(langId LANGID, codePage co.CP) (string, bool)
	FileDescription(langId LANGID, codePage co.CP) (string, bool)
	FileVersion(langId LANGID, codePage co.CP) (string, bool)
	InternalName(langId LANGID, codePage co.CP) (string, bool)
	LegalCopyright(langId LANGID, codePage co.CP) (string, bool)
	LegalTrademarks(langId LANGID, codePage co.CP) (string, bool)
	OriginalFilename(langId LANGID, codePage co.CP) (string, bool)
	ProductName(langId LANGID, codePage co.CP) (string, bool)
	ProductVersion(langId LANGID, codePage co.CP) (string, bool)
	PrivateBuild(langId LANGID, codePage co.CP) (string, bool)
	SpecialBuild(langId LANGID, codePage co.CP) (string, bool)
}

//------------------------------------------------------------------------------

type _ResourceInfo struct {
	resBuf []byte
}

// Reads and stores an embedded resource from an executable or DLL file.
func LoadResourceInfo(exePath string) (ResourceInfo, error) {
	resBuf, err := GetFileVersionInfo(exePath)
	if err != nil {
		return nil, err
	}

	return &_ResourceInfo{resBuf: resBuf}, nil
}

func (me *_ResourceInfo) FixedFileInfo() (*VS_FIXEDFILEINFO, bool) {
	ptr, _, ok := VerQueryValue(me.resBuf, "\\")
	if !ok {
		return nil, false
	}

	return (*VS_FIXEDFILEINFO)(ptr), true
}

type _ResourceInfoBlock struct {
	LangId   LANGID
	CodePage co.CP
}

func (me *_ResourceInfo) Blocks() []_ResourceInfoBlock {
	ptr, sz, ok := VerQueryValue(me.resBuf, "\\VarFileInfo\\Translation")
	if !ok {
		return []_ResourceInfoBlock{}
	}

	return unsafe.Slice((*_ResourceInfoBlock)(ptr), sz)
}

func (me *_ResourceInfo) Comments(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "Comments")
}
func (me *_ResourceInfo) CompanyName(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "CompanyName")
}
func (me *_ResourceInfo) FileDescription(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "FileDescription")
}
func (me *_ResourceInfo) FileVersion(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "FileVersion")
}
func (me *_ResourceInfo) InternalName(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "InternalName")
}
func (me *_ResourceInfo) LegalCopyright(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "LegalCopyright")
}
func (me *_ResourceInfo) LegalTrademarks(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "LegalTrademarks")
}
func (me *_ResourceInfo) OriginalFilename(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "OriginalFilename")
}
func (me *_ResourceInfo) ProductName(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "ProductName")
}
func (me *_ResourceInfo) ProductVersion(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "ProductVersion")
}
func (me *_ResourceInfo) PrivateBuild(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "PrivateBuild")
}
func (me *_ResourceInfo) SpecialBuild(langId LANGID, codePage co.CP) (string, bool) {
	return me.genericStringInfo(langId, codePage, "SpecialBuild")
}

func (me *_ResourceInfo) genericStringInfo(
	langId LANGID, codePage co.CP, info string) (string, bool) {

	ptr, sz, ok := VerQueryValue(me.resBuf,
		fmt.Sprintf("\\StringFileInfo\\%04x%04x\\%s", langId, codePage, info))
	if !ok {
		return "", false
	}

	return Str.FromNativeSlice(unsafe.Slice((*uint16)(ptr), sz)), true
}
