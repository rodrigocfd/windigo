//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [EmptyWorkingSet] function.
//
// [EmptyWorkingSet]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-emptyworkingset
func (hProcess HPROCESS) EmptyWorkingSet() error {
	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_EmptyWorkingSet, "EmptyWorkingSet"),
		uintptr(hProcess))
	return utl.ZeroAsGetLastError(ret, err)
}

var _psapi_EmptyWorkingSet *syscall.Proc

// [EnumProcessModules] function.
//
// Example:
//
//	hProcess := win.GetCurrentProcess()
//	hModules, _ := hProcess.EnumProcessModules()
//
//	for _, hModule := range hModules {
//		fileName, _ := hModule.GetModuleFileName()
//		println(fileName)
//	}
//
// [EnumProcessModules]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-enumprocessmodules
func (hProcess HPROCESS) EnumProcessModules() ([]HINSTANCE, error) {
	for {
		var bytesNeeded uint32
		ret, _, err := syscall.SyscallN(
			dll.Psapi.Load(&_psapi_EnumProcessModules, "EnumProcessModules"),
			uintptr(hProcess),
			0, 0,
			uintptr(unsafe.Pointer(&bytesNeeded)))
		if ret == 0 {
			return nil, co.ERROR(err)
		}

		elemsNeeded := uintptr(bytesNeeded) / unsafe.Sizeof(HINSTANCE(0))
		elems := make([]HINSTANCE, elemsNeeded)

		var bytesGot uint32
		ret, _, err = syscall.SyscallN(
			dll.Psapi.Load(&_psapi_EnumProcessModules, "EnumProcessModules"),
			uintptr(hProcess),
			uintptr(unsafe.Pointer(unsafe.SliceData(elems))),
			uintptr(bytesNeeded),
			uintptr(unsafe.Pointer(&bytesGot)))
		if ret == 0 {
			return nil, co.ERROR(err)
		}

		if bytesNeeded == bytesGot {
			return elems, nil
		}
	}
}

var _psapi_EnumProcessModules *syscall.Proc

// [GetMappedFileName] function.
//
// [GetMappedFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmappedfilenamew
func (hProcess HPROCESS) GetMappedFileName(address uintptr) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_GetMappedFileNameW, "GetMappedFileNameW"),
		uintptr(hProcess),
		address,
		uintptr(wBuf.Ptr()),
		uintptr(uint32(wBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _psapi_GetMappedFileNameW *syscall.Proc

// [GetModuleBaseName] function.
//
// [GetModuleBaseName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulebasenamew
func (hProcess HPROCESS) GetModuleBaseName(hModule HINSTANCE) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_GetModuleBaseNameW, "GetModuleBaseNameW"),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(wBuf.Ptr()),
		uintptr(uint32(wBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _psapi_GetModuleBaseNameW *syscall.Proc

// [GetModuleFileNameEx] function.
//
// [GetModuleFileNameEx]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulefilenameexw
func (hProcess HPROCESS) GetModuleFileNameEx(hModule HINSTANCE) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_GetModuleFileNameExW, "GetModuleFileNameExW"),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(wBuf.Ptr()),
		uintptr(uint32(wBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _psapi_GetModuleFileNameExW *syscall.Proc

// [GetModuleInformation] function.
//
// [GetModuleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmoduleinformation
func (hProcess HPROCESS) GetModuleInformation(hModule HINSTANCE) (MODULEINFO, error) {
	var mi MODULEINFO
	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_GetModuleInformation, "GetModuleInformation"),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&mi)),
		uintptr(uint32(unsafe.Sizeof(mi))))
	if ret == 0 {
		return MODULEINFO{}, co.ERROR(err)
	}
	return mi, nil
}

var _psapi_GetModuleInformation *syscall.Proc

// [GetProcessImageFileName] function
//
// [GetProcessImageFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessimagefilenamew
func (hProcess HPROCESS) GetProcessImageFileName() (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_K32GetProcessImageFileNameW, "K32GetProcessImageFileNameW"),
		uintptr(hProcess),
		uintptr(wBuf.Ptr()),
		uintptr(uint32(wBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _psapi_K32GetProcessImageFileNameW *syscall.Proc

// [GetProcessMemoryInfo] function
//
// [GetProcessMemoryInfo]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessmemoryinfo
func (hProcess HPROCESS) GetProcessMemoryInfo() (PROCESS_MEMORY_COUNTERS_EX, error) {
	var pmc PROCESS_MEMORY_COUNTERS_EX
	pmc.SetCb()

	ret, _, err := syscall.SyscallN(
		dll.Psapi.Load(&_psapi_GetProcessMemoryInfo, "GetProcessMemoryInfo"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&pmc)),
		uintptr(uint32(unsafe.Sizeof(pmc))))
	if ret == 0 {
		return PROCESS_MEMORY_COUNTERS_EX{}, co.ERROR(err)
	}
	return pmc, nil
}

var _psapi_GetProcessMemoryInfo *syscall.Proc
