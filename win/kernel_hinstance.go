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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pModuleName := wbuf.PtrEmptyIsNil(moduleName)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetModuleHandleW, "GetModuleHandleW"),
		uintptr(pModuleName))
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
	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pLibFileName := wbuf.PtrEmptyIsNil(libFileName)

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_LoadLibraryW, "LoadLibraryW"),
		uintptr(pLibFileName))
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
// # Example
//
// Retrieving own .exe path:
//
//	exePath, _ := win.HINSTANCE(0).GetModuleFileName()
//	println(exePath)
//
// [GetModuleFileName]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (hInst HINSTANCE) GetModuleFileName() (string, error) {
	sz := wstr.BUF_MAX
	buf := wstr.NewBufDecoder(sz)
	defer buf.Free()

	for {
		ret, _, err := syscall.SyscallN(
			dll.Load(dll.KERNEL32, &_GetModuleFileNameW, "GetModuleFileNameW"),
			uintptr(hInst),
			uintptr(buf.UnsafePtr()),
			uintptr(uint32(sz)))
		if ret == 0 {
			return "", co.ERROR(err)
		}
		chCopied := uint(ret) + 1 // plus terminating null count

		if chCopied < sz { // to break, must have at least 1 char gap
			return buf.String(), nil
		}

		sz += 64
		buf.Resize(sz) // increase buffer size to try again
	}
}

var _GetModuleFileNameW *syscall.Proc
