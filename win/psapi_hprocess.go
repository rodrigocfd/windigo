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
	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_EmptyWorkingSet, "EmptyWorkingSet"),
		uintptr(hProcess))
	return utl.ZeroAsGetLastError(ret, err)
}

var _EmptyWorkingSet *syscall.Proc

// [GetMappedFileName] function.
//
// [GetMappedFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmappedfilenamew
func (hProcess HPROCESS) GetMappedFileName(address uintptr) (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_GetMappedFileNameW, "GetMappedFileNameW"),
		uintptr(hProcess),
		address,
		uintptr(recvBuf.UnsafePtr()),
		uintptr(uint32(recvBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return recvBuf.String(), nil
}

var _GetMappedFileNameW *syscall.Proc

// [GetModuleBaseName] function.
//
// [GetModuleBaseName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulebasenamew
func (hProcess HPROCESS) GetModuleBaseName(hModule HINSTANCE) (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_GetModuleBaseNameW, "GetModuleBaseNameW"),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(uint32(recvBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return recvBuf.String(), nil
}

var _GetModuleBaseNameW *syscall.Proc

// [GetModuleFileNameEx] function.
//
// [GetModuleFileNameEx]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulefilenameexw
func (hProcess HPROCESS) GetModuleFileNameEx(hModule HINSTANCE) (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_GetModuleFileNameExW, "GetModuleFileNameExW"),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(uint32(recvBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return recvBuf.String(), nil
}

var _GetModuleFileNameExW *syscall.Proc

// [GetModuleInformation] function.
//
// [GetModuleInformation]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmoduleinformation
func (hProcess HPROCESS) GetModuleInformation(hModule HINSTANCE) (MODULEINFO, error) {
	var mi MODULEINFO
	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_GetModuleInformation, "GetModuleInformation"),
		uintptr(hProcess),
		uintptr(hModule),
		uintptr(unsafe.Pointer(&mi)),
		uintptr(uint32(unsafe.Sizeof(mi))))
	if ret == 0 {
		return MODULEINFO{}, co.ERROR(err)
	}
	return mi, nil
}

var _GetModuleInformation *syscall.Proc

// [GetProcessImageFileName] function
//
// [GetProcessImageFileName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessimagefilenamew
func (hProcess HPROCESS) GetProcessImageFileName() (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_K32GetProcessImageFileNameW, "K32GetProcessImageFileNameW"),
		uintptr(hProcess),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(uint32(recvBuf.Len())))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return recvBuf.String(), nil
}

var _K32GetProcessImageFileNameW *syscall.Proc

// [GetProcessMemoryInfo] function
//
// [GetProcessMemoryInfo]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getprocessmemoryinfo
func (hProcess HPROCESS) GetProcessMemoryInfo() (PROCESS_MEMORY_COUNTERS_EX, error) {
	var pmc PROCESS_MEMORY_COUNTERS_EX
	pmc.SetCb()

	ret, _, err := syscall.SyscallN(
		dll.Psapi(&_GetProcessMemoryInfo, "GetProcessMemoryInfo"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&pmc)),
		uintptr(uint32(unsafe.Sizeof(pmc))))
	if ret == 0 {
		return PROCESS_MEMORY_COUNTERS_EX{}, co.ERROR(err)
	}
	return pmc, nil
}

var _GetProcessMemoryInfo *syscall.Proc
