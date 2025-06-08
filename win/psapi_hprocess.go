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

// [EmptyWorkingSet] function.
//
// [EmptyWorkingSet]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-emptyworkingset
func (hProcess HPROCESS) EmptyWorkingSet() error {
	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_EmptyWorkingSet),
		uintptr(hProcess))
	return utl.ZeroAsGetLastError(ret, err)
}

// [GetMappedFileName] function.
//
// [GetMappedFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmappedfilenamew
func (hProcess HPROCESS) GetMappedFileName(address uintptr) (string, error) {
	var fileName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_GetMappedFileNameW),
		uintptr(hProcess),
		address,
		uintptr(unsafe.Pointer(&fileName[0])),
		uintptr(len(fileName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(fileName[:]), nil
}

// [GetModuleBaseName] function.
//
// [GetModuleBaseName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulebasenamew
func (hProcess HPROCESS) GetModuleBaseName(hModule HINSTANCE) (string, error) {
	var baseName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_GetModuleBaseNameW),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&baseName[0])),
		uintptr(len(baseName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(baseName[:]), nil
}

// [GetModuleFileNameEx] function.
//
// [GetModuleFileNameEx]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulefilenameexw
func (hProcess HPROCESS) GetModuleFileNameEx(hModule HINSTANCE) (string, error) {
	var fileName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_GetModuleFileNameExW),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&fileName[0])),
		uintptr(len(fileName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(fileName[:]), nil
}

// [GetModuleInformation] function.
//
// [GetModuleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmoduleinformation
func (hProcess HPROCESS) GetModuleInformation(hModule HINSTANCE) (MODULEINFO, error) {
	var mi MODULEINFO
	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_GetModuleInformation),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&mi)),
		uintptr(uint32(unsafe.Sizeof(mi))))
	if ret == 0 {
		return MODULEINFO{}, co.ERROR(err)
	}
	return mi, nil
}

// [GetProcessImageFileName] function
//
// [GetProcessImageFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessimagefilenamew
func (hProcess HPROCESS) GetProcessImageFileName() (string, error) {
	var buf [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_K32GetProcessImageFileNameW),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(uint32(len(buf))))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(buf[:]), nil
}

// [GetProcessMemoryInfo] function
//
// [GetProcessMemoryInfo]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessmemoryinfo
func (hProcess HPROCESS) GetProcessMemoryInfo() (PROCESS_MEMORY_COUNTERS_EX, error) {
	var pmc PROCESS_MEMORY_COUNTERS_EX
	pmc.SetCb()

	ret, _, err := syscall.SyscallN(dll.Psapi(dll.PROC_GetProcessMemoryInfo),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&pmc)),
		uintptr(uint32(unsafe.Sizeof(pmc))))
	if ret == 0 {
		return PROCESS_MEMORY_COUNTERS_EX{}, co.ERROR(err)
	}
	return pmc, nil
}
