//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [GetFileVersionInfo] function.
//
// This is a low-level function, prefer using [VersionLoad].
//
// # Example
//
//	hInst, _ := win.GetModuleHandle("")
//	exeName, _ := hInst.GetModuleFileName()
//	szData, _ := win.GetFileVersionInfoSize(exeName)
//
//	data := heap.NewVecSized(szData, byte(0))
//	defer data.Free()
//
//	win.GetFileVersionInfo(exeName, data.HotSlice())
//
// [GetFileVersionInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfow
func GetFileVersionInfo(fileName string, dest []byte) error {
	fileName16 := wstr.NewBufWith[wstr.Stack20](fileName, wstr.EMPTY_IS_NIL)
	ret, _, err := syscall.SyscallN(_GetFileVersionInfoW.Addr(),
		uintptr(fileName16.UnsafePtr()), 0,
		uintptr(len(dest)), uintptr(unsafe.Pointer(&dest[0])))
	return wutil.ZeroAsGetLastError(ret, err)
}

var _GetFileVersionInfoW = dll.Version.NewProc("GetFileVersionInfoW")

// [GetFileVersionInfoSize] function.
//
// [GetFileVersionInfoSize]: https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-getfileversioninfosizew
func GetFileVersionInfoSize(fileName string) (uint, error) {
	fileName16 := wstr.NewBufWith[wstr.Stack20](fileName, wstr.EMPTY_IS_NIL)
	var dummy uint32

	ret, _, err := syscall.SyscallN(_GetFileVersionInfoSizeW.Addr(),
		uintptr(fileName16.UnsafePtr()), uintptr(unsafe.Pointer(&dummy)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(ret), nil
}

var _GetFileVersionInfoSizeW = dll.Version.NewProc("GetFileVersionInfoSizeW")

// [VerQueryValue] function.
//
// This is a low-level function, prefer using [VersionLoad].
//
// # Example
//
//	hInst, _ := win.GetModuleHandle("")
//	exeName, _ := hInst.GetModuleFileName()
//	szData, _ := win.GetFileVersionInfoSize(exeName)
//
//	data := heap.NewVecSized(szData, byte(0))
//	defer data.Free()
//
//	win.GetFileVersionInfo(exeName, data.HotSlice())
//
//	if pNfoRaw, _, ok := win.VerQueryValue(data.HotSlice(), "\\"); ok {
//		pNfo := (*win.VS_FIXEDFILEINFO)(pNfoRaw)
//		println(pNfo.FileVersion())
//	}
//
//	type Block struct {
//		LangId   win.LANGID
//		CodePage co.CP
//	}
//
//	if pBlocks, count, ok := win.VerQueryValue(
//		data.HotSlice(), "\\VarFileInfo\\Translation"); ok {
//
//		blocks := unsafe.Slice((*Block)(pBlocks), count)
//		for _, block := range blocks {
//			if pStr, nChars, ok := win.VerQueryValue(data.HotSlice(),
//				fmt.Sprintf("\\StringFileInfo\\%04x%04x\\%s",
//					block.LangId, block.CodePage, "ProductName")); ok {
//
//				wideStr := unsafe.Slice((*uint16)(pStr), nChars)
//				str := wstr.WstrSliceToStr(wideStr)
//				println(str)
//			}
//		}
//	}
//
// [VerQueryValue]: https://learn.microsoft.com/en-us/windows/win32/api/winver/nf-winver-verqueryvaluew
func VerQueryValue(block []byte, subBlock string) (unsafe.Pointer, uint, bool) {
	subBlock16 := wstr.NewBufWith[wstr.Stack20](subBlock, wstr.ALLOW_EMPTY)
	var lplpBuffer uintptr
	var puLen uint32

	ret, _, _ := syscall.SyscallN(_VerQueryValueW.Addr(),
		uintptr(unsafe.Pointer(&block[0])), uintptr(subBlock16.UnsafePtr()),
		uintptr(unsafe.Pointer(&lplpBuffer)), uintptr(unsafe.Pointer(&puLen)))
	if ret == 0 {
		return nil, 0, false
	}
	return unsafe.Pointer(lplpBuffer), uint(puLen), true
}

var _VerQueryValueW = dll.Version.NewProc("VerQueryValueW")
