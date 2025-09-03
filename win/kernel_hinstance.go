//go:build windows

package win

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to an [instance]. This is the base address of the module in memory.
//
// [instance]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hinstance
type HINSTANCE HANDLE

// [GetModuleHandle] function.
//
// Example:
//
// Retrieving own .exe handle:
//
//	hInst, _ := win.GetModuleHandle("")
//
// [GetModuleHandle]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName string) (HINSTANCE, error) {
	var wModuleName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetModuleHandleW, "GetModuleHandleW"),
		uintptr(wModuleName.EmptyIsNil(moduleName)))
	if ret == 0 {
		return HINSTANCE(0), co.ERROR(err)
	}
	return HINSTANCE(ret), nil
}

var _GetModuleHandleW *syscall.Proc

// [LoadLibrary] function.
//
// ⚠️ You must defer [HINSTANCE.FreeLibrary].
//
// [LoadLibrary]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
func LoadLibrary(libFileName string) (HINSTANCE, error) {
	var wLibFileName wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LoadLibraryW, "LoadLibraryW"),
		uintptr(wLibFileName.EmptyIsNil(libFileName)))
	if ret == 0 {
		return HINSTANCE(0), co.ERROR(err)
	}
	return HINSTANCE(ret), nil
}

var _LoadLibraryW *syscall.Proc

// [FreeLibrary] function.
//
// Paired with [LoadLibrary].
//
// [FreeLibrary]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (hInst HINSTANCE) FreeLibrary() error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_FreeLibrary, "FreeLibrary"),
		uintptr(hInst))
	return utl.ZeroAsGetLastError(ret, err)
}

var _FreeLibrary *syscall.Proc

// [GetModuleFileName] function.
//
// Example:
//
// Retrieving own .exe path:
//
//	exePath, _ := win.HINSTANCE(0).GetModuleFileName()
//	println(exePath)
//
// [GetModuleFileName]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (hInst HINSTANCE) GetModuleFileName() (string, error) {
	sz := wstr.BUF_MAX
	var wBuf wstr.BufDecoder
	wBuf.Alloc(sz)

	for {
		ret, _, err := syscall.SyscallN(
			dll.Load(dll.KERNEL32, &_GetModuleFileNameW, "GetModuleFileNameW"),
			uintptr(hInst),
			uintptr(wBuf.Ptr()),
			uintptr(uint32(sz)))
		if ret == 0 {
			return "", co.ERROR(err)
		}
		chCopied := int(ret) + 1 // plus terminating null count

		if chCopied < sz { // to break, must have at least 1 char gap
			return wBuf.String(), nil
		}

		sz += 64
		wBuf.AllocAndZero(sz) // increase buffer size to try again
	}
}

var _GetModuleFileNameW *syscall.Proc
