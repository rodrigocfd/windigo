//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [GetMappedFileName] function.
//
// [GetMappedFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmappedfilenamew
func (hProcess HPROCESS) GetMappedFileName(address uintptr) (string, error) {
	var fileName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(_GetMappedFileNameW.Addr(),
		uintptr(hProcess),
		address,
		uintptr(unsafe.Pointer(&fileName[0])),
		uintptr(len(fileName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(fileName[:]), nil
}

var _GetMappedFileNameW = dll.Psapi.NewProc("GetMappedFileNameW")

// [GetModuleBaseName] function.
//
// [GetModuleBaseName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulebasenamew
func (hProcess HPROCESS) GetModuleBaseName(hModule HINSTANCE) (string, error) {
	var baseName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(_GetModuleBaseNameW.Addr(),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&baseName[0])),
		uintptr(len(baseName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(baseName[:]), nil
}

var _GetModuleBaseNameW = dll.Psapi.NewProc("GetModuleBaseNameW")

// [GetModuleInformation] function.
//
// [GetModuleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmoduleinformation
func (hProcess HPROCESS) GetModuleInformation(hModule HINSTANCE) (MODULEINFO, error) {
	var mi MODULEINFO
	ret, _, err := syscall.SyscallN(_GetModuleInformation.Addr(),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&mi)),
		uintptr(uint32(unsafe.Sizeof(mi))))
	if ret == 0 {
		return MODULEINFO{}, co.ERROR(err)
	}
	return mi, nil
}

var _GetModuleInformation = dll.Psapi.NewProc("GetModuleInformation")

// [GetModuleFileNameEx] function.
//
// [GetModuleFileNameEx]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulefilenameexw
func (hProcess HPROCESS) GetModuleFileNameEx(hModule HINSTANCE) (string, error) {
	var fileName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(_GetModuleFileNameExW.Addr(),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&fileName[0])),
		uintptr(len(fileName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(fileName[:]), nil
}

var _GetModuleFileNameExW = dll.Psapi.NewProc("GetModuleFileNameExW")
