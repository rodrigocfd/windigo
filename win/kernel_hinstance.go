package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to an instance. This is the base address of the module in memory.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#hinstance
type HINSTANCE HANDLE

// If moduleName is nil, returns a handle to the file used to create the calling
// process (.exe file).
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName StrOrNil) HINSTANCE {
	ret, _, err := syscall.Syscall(proc.GetModuleHandle.Addr(), 1,
		uintptr(variantStrOrNil(moduleName)), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HINSTANCE(ret)
}

// ⚠️ You must defer HINSTANCE.FreeLibrary().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-loadlibraryw
func LoadLibrary(libFileName string) HINSTANCE {
	ret, _, err := syscall.Syscall(proc.LoadLibrary.Addr(), 1,
		uintptr(unsafe.Pointer(Str.ToNativePtr(libFileName))),
		0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return HINSTANCE(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-freelibrary
func (hInst HINSTANCE) FreeLibrary() {
	ret, _, err := syscall.Syscall(proc.FreeLibrary.Addr(), 1,
		uintptr(hInst), 0, 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
}

// Example retrieving own .exe path:
//
//  exePath := win.HINSTANCE(0).GetModuleFileName()
//  fmt.Printf("Current .exe path: %s\n", exePath)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulefilenamew
func (hInst HINSTANCE) GetModuleFileName() string {
	var buf [_MAX_PATH + 1]uint16
	ret, _, err := syscall.Syscall(proc.GetModuleFileName.Addr(), 3,
		uintptr(hInst), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return Str.FromNativeSlice(buf[:])
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/libloaderapi/nf-libloaderapi-getprocaddress
func (hInst HINSTANCE) GetProcAddress(procName string) uintptr {
	ascii := []byte(procName)
	ascii = append(ascii, 0x00) // terminating null

	ret, _, err := syscall.Syscall(proc.GetProcAddress.Addr(), 2,
		uintptr(hInst), uintptr(unsafe.Pointer(&ascii[0])), 0)
	if ret == 0 {
		panic(errco.ERROR(err))
	}
	return ret
}
