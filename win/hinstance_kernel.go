//go:build windows

package win

import (
	"runtime"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an [instance]. This is the base address of the module in memory.
//
// [instance]: https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hinstance
type HINSTANCE HANDLE

// [GetModuleHandle] function.
//
// If moduleName is nil, returns a handle to the file used to create the calling
// process (.exe file).
//
// [GetModuleHandle]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName StrOpt) HINSTANCE {
	ret, _, err := syscall.SyscallN(proc.GetModuleHandle.Addr(),
		uintptr(moduleName.Raw()))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HINSTANCE(ret)
}

// [LoadLibrary] function.
//
// ⚠️ You must defer HINSTANCE.FreeLibrary().
//
// [LoadLibrary]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
func LoadLibrary(libFileName string) HINSTANCE {
	ret, _, err := syscall.SyscallN(proc.LoadLibrary.Addr(),
		uintptr(unsafe.Pointer(Str.ToNativePtr(libFileName))))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HINSTANCE(ret)
}

// [FindResource] function.
//
// [FindResource]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-findresourcew
func (hInst HINSTANCE) FindResource(
	name ResId, rsrcType RsrcType) (HRSRC, error) {

	nameVal, nameBuf := name.raw()
	rsrcTypeVal, rsrcTypeBuf := rsrcType.raw()

	ret, _, err := syscall.SyscallN(proc.FindResource.Addr(),
		uintptr(hInst), nameVal, rsrcTypeVal)
	runtime.KeepAlive(nameBuf)
	runtime.KeepAlive(rsrcTypeBuf)

	if ret == 0 {
		return HRSRC(0), errco.ERROR(err)
	}
	return HRSRC(ret), nil
}

// [FindResourceEx] function.
//
// [FindResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-findresourceexw
func (hInst HINSTANCE) FindResourceEx(
	name ResId,
	rsrcType RsrcType,
	language LANGID) (HRSRC, error) {

	nameVal, nameBuf := name.raw()
	rsrcTypeVal, rsrcTypeBuf := rsrcType.raw()

	ret, _, err := syscall.SyscallN(proc.FindResourceEx.Addr(),
		uintptr(hInst), nameVal, rsrcTypeVal, uintptr(language))
	runtime.KeepAlive(nameBuf)
	runtime.KeepAlive(rsrcTypeBuf)

	if ret == 0 {
		return HRSRC(0), errco.ERROR(err)
	}
	return HRSRC(ret), nil
}

// [FreeLibrary] function.
//
// [FreeLibrary]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (hInst HINSTANCE) FreeLibrary() error {
	ret, _, err := syscall.SyscallN(proc.FreeLibrary.Addr(),
		uintptr(hInst))
	if ret == 0 {
		return errco.ERROR(err)
	}
	return nil
}

// [GetModuleFileName] function.
//
// # Example
//
// Retrieving own .exe path:
//
//	exePath := win.HINSTANCE(0).GetModuleFileName()
//	fmt.Printf("Current .exe path: %s\n", exePath)
//
// [GetModuleFileName]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (hInst HINSTANCE) GetModuleFileName() string {
	var buf [_MAX_PATH + 1]uint16
	ret, _, err := syscall.SyscallN(proc.GetModuleFileName.Addr(),
		uintptr(hInst), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

// [GetProcAddress] function.
//
// [GetProcAddress]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getprocaddress
func (hInst HINSTANCE) GetProcAddress(procName string) (uintptr, error) {
	ascii := []byte(procName)
	ascii = append(ascii, 0x00) // terminating null

	ret, _, err := syscall.SyscallN(proc.GetProcAddress.Addr(),
		uintptr(hInst), uintptr(unsafe.Pointer(&ascii[0])))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return ret, nil
}

// [LoadResource] function.
//
// [LoadResource]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadresource
func (hInst HINSTANCE) LoadResource(hResInfo HRSRC) (HRSRCMEM, error) {
	ret, _, err := syscall.SyscallN(proc.LoadResource.Addr(),
		uintptr(hInst), uintptr(hResInfo))
	if ret == 0 {
		return HRSRCMEM(0), errco.ERROR(err)
	}
	return HRSRCMEM(ret), nil
}

// [LockResource] function.
//
// This method should belong to HRSRCMEM, but in order to make it safe, we
// automatically call HINSTANCE.SizeofResource(), so it's implemented here.
//
// [LockResource]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-lockresource
func (hInst HINSTANCE) LockResource(
	hResInfo HRSRC, hResLoaded HRSRCMEM) ([]byte, error) {

	sz, szErr := hInst.SizeofResource(hResInfo)
	if szErr != nil {
		return nil, szErr
	}

	ret, _, err := syscall.SyscallN(proc.LockResource.Addr(),
		uintptr(hInst), uintptr(hResInfo), uintptr(hResLoaded))
	if ret == 0 {
		return nil, errco.ERROR(err)
	}

	return unsafe.Slice((*byte)(unsafe.Pointer(ret)), sz), nil
}

// [SizeofResource] function.
//
// [SizeofResource]: https://learn.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-sizeofresource
func (hInst HINSTANCE) SizeofResource(hResInfo HRSRC) (int, error) {
	ret, _, err := syscall.SyscallN(proc.SizeofResource.Addr(),
		uintptr(hInst), uintptr(hResInfo))
	if ret == 0 {
		return 0, errco.ERROR(err)
	}
	return int(ret), nil
}
