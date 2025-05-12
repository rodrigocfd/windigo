//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Handle to an [instance]. This is the base address of the module in memory.
//
// [instance]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hinstance
type HINSTANCE HANDLE

// [GetModuleHandle] function.
//
// # Example
//
// Retrieving own .exe handle:
//
//	hInst, _ := win.GetModuleHandle("")
//
// [GetModuleHandle]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName string) (HINSTANCE, error) {
	moduleName16 := wstr.NewBufWith[wstr.Stack20](moduleName, wstr.EMPTY_IS_NIL)
	ret, _, err := syscall.SyscallN(_GetModuleHandleW.Addr(),
		uintptr(moduleName16.UnsafePtr()))
	if ret == 0 {
		return HINSTANCE(0), co.ERROR(err)
	}
	return HINSTANCE(ret), nil
}

var _GetModuleHandleW = dll.Kernel32.NewProc("GetModuleHandleW")

// [LoadLibrary] function.
//
// ⚠️ You must defer [HINSTANCE.FreeLibrary].
//
// [LoadLibrary]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
func LoadLibrary(libFileName string) (HINSTANCE, error) {
	libFileName16 := wstr.NewBufWith[wstr.Stack20](libFileName, wstr.EMPTY_IS_NIL)
	ret, _, err := syscall.SyscallN(_LoadLibraryW.Addr(),
		uintptr(libFileName16.UnsafePtr()))
	if ret == 0 {
		return HINSTANCE(0), co.ERROR(err)
	}
	return HINSTANCE(ret), nil
}

var _LoadLibraryW = dll.Kernel32.NewProc("LoadLibraryW")

// [FreeLibrary] function.
//
// Paired with [LoadLibrary].
//
// [FreeLibrary]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (hInst HINSTANCE) FreeLibrary() error {
	ret, _, err := syscall.SyscallN(_FreeLibrary.Addr(),
		uintptr(hInst))
	return utl.ZeroAsGetLastError(ret, err)
}

var _FreeLibrary = dll.Kernel32.NewProc("FreeLibrary")

// [GetModuleFileName] function.
//
// # Example
//
// Retrieving own .exe path:
//
//	exePath, _ := win.HINSTANCE(0).GetModuleFileName()
//	println(exePath)
//
// [GetModuleFileName]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (hInst HINSTANCE) GetModuleFileName() (string, error) {
	buf := wstr.NewBufSized[wstr.Stack260](260) // start allocating on the stack

	for {
		ret, _, err := syscall.SyscallN(_GetModuleFileNameW.Addr(),
			uintptr(hInst), uintptr(buf.UnsafePtr()), uintptr(buf.Len()))
		if ret == 0 {
			return "", co.ERROR(err)
		}
		chCopied := uint(ret) + 1 // plus terminating null count

		if chCopied < buf.Len() { // to break, must have at least 1 char gap
			return wstr.WstrSliceToStr(buf.HotSlice()), nil
		}

		buf.Resize(buf.Len() + 64) // increase buffer size to try again
	}
}

var _GetModuleFileNameW = dll.Kernel32.NewProc("GetModuleFileNameW")
