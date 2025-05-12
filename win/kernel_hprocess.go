//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Handle to a process.
type HPROCESS HANDLE

// [GetCurrentProcess] function.
//
// [GetCurrentProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocess
func GetCurrentProcess() HPROCESS {
	ret, _, _ := syscall.SyscallN(_GetCurrentProcess.Addr())
	return HPROCESS(ret)
}

var _GetCurrentProcess = dll.Kernel32.NewProc("GetCurrentProcess")

// [OpenProcess] function.
//
// ⚠️ You must defer [HPROCESS.CloseHandle].
//
// [OpenProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess
func OpenProcess(access co.PROCESS, inheritHandle bool, processId uint32) (HPROCESS, error) {
	ret, _, err := syscall.SyscallN(_OpenProcess.Addr(),
		uintptr(access), utl.BoolToUintptr(inheritHandle), uintptr(processId))
	if ret == 0 {
		return HPROCESS(0), co.ERROR(err)
	}
	return HPROCESS(ret), nil
}

var _OpenProcess = dll.Kernel32.NewProc("OpenProcess")

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcess HPROCESS) CloseHandle() error {
	return HANDLE(hProcess).CloseHandle()
}

// [GetExitCodeProcess] function.
//
// [GetExitCodeProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
func (hProcess HPROCESS) GetExitCodeProcess() (uint32, error) {
	var exitCode uint32
	ret, _, err := syscall.SyscallN(_GetExitCodeProcess.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return exitCode, nil
}

var _GetExitCodeProcess = dll.Kernel32.NewProc("GetExitCodeProcess")

// [GetModuleBaseName] function.
//
// [GetModuleBaseName]: https://learn.microsoft.com/en-us/windows/win32/api/psapi/nf-psapi-getmodulebasenamew
func (hProcess HPROCESS) GetModuleBaseName(hModule HINSTANCE) (string, error) {
	var processName [utl.MAX_PATH]uint16
	ret, _, err := syscall.SyscallN(_GetModuleBaseNameW.Addr(),
		uintptr(hProcess), uintptr(hModule),
		uintptr(unsafe.Pointer(&processName[0])),
		uintptr(len(processName)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(processName[:]), nil
}

var _GetModuleBaseNameW = dll.Kernel32.NewProc("GetModuleBaseNameW")

// [GetPriorityClass] function.
//
// [GetPriorityClass]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getpriorityclass
func (hProcess HPROCESS) GetPriorityClass() (co.PRIORITY, error) {
	ret, _, err := syscall.SyscallN(_GetPriorityClass.Addr(),
		uintptr(hProcess))
	if ret == 0 {
		return co.PRIORITY(0), co.ERROR(err)
	}
	return co.PRIORITY(ret), nil
}

var _GetPriorityClass = dll.Kernel32.NewProc("GetPriorityClass")

// [GetProcessHandleCount] function.
//
// [GetProcessHandleCount]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesshandlecount
func (hProcess HPROCESS) GetProcessHandleCount() (uint, error) {
	var count uint32
	ret, _, err := syscall.SyscallN(_GetProcessHandleCount.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&count)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(count), nil
}

var _GetProcessHandleCount = dll.Kernel32.NewProc("GetProcessHandleCount")

// [GetProcessId] function.
//
// [GetProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessid
func (hProcess HPROCESS) GetProcessId() (uint32, error) {
	ret, _, err := syscall.SyscallN(_GetProcessId.Addr(),
		uintptr(hProcess))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _GetProcessId = dll.Kernel32.NewProc("GetProcessId")

// [GetProcessTimes] function.
//
// [GetProcessTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes
func (hProcess HPROCESS) GetProcessTimes() (HprocessTimes, error) {
	var ftCreation, ftExit, ftKernel, ftUser FILETIME
	ret, _, err := syscall.SyscallN(_GetProcessTimes.Addr(),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&ftCreation)), uintptr(unsafe.Pointer(&ftExit)),
		uintptr(unsafe.Pointer(&ftKernel)), uintptr(unsafe.Pointer(&ftUser)))
	if ret == 0 {
		return HprocessTimes{}, co.ERROR(err)
	}

	return HprocessTimes{
		Creation: ftCreation.ToTime(),
		Exit:     ftExit.ToTime(),
		Kernel:   ftKernel.ToTime(),
		User:     ftUser.ToTime(),
	}, nil
}

// Returned by [HPROCESS.GetProcessTimes].
type HprocessTimes struct {
	Creation time.Time
	Exit     time.Time
	Kernel   time.Time
	User     time.Time
}

var _GetProcessTimes = dll.Kernel32.NewProc("GetProcessTimes")

// [IsProcessCritical] function.
//
// [IsProcessCritical]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-isprocesscritical
func (hProcess HPROCESS) IsProcessCritical() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(_IsProcessCritical.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _IsProcessCritical = dll.Kernel32.NewProc("IsProcessCritical")

// [IsWow64Process] function.
//
// [IsWow64Process]: https://learn.microsoft.com/en-us/windows/win32/api/wow64apiset/nf-wow64apiset-iswow64process
func (hProcess HPROCESS) IsWow64Process() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(_IsWow64Process.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _IsWow64Process = dll.Kernel32.NewProc("IsWow64Process")

// [QueryFullProcessImageName] function.
//
// [QueryFullProcessImageName]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-queryfullprocessimagenamew
func (hProcess HPROCESS) QueryFullProcessImageName(flags co.PROCESS_NAME) (string, error) {
	var buf [utl.MAX_PATH]uint16
	szBuf := uint32(len(buf))

	ret, _, err := syscall.SyscallN(_QueryFullProcessImageNameW.Addr(),
		uintptr(hProcess), uintptr(flags),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&szBuf)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wstr.WstrSliceToStr(buf[:]), nil
}

var _QueryFullProcessImageNameW = dll.Kernel32.NewProc("QueryFullProcessImageNameW")

// [QueryProcessAffinityUpdateMode] function.
//
// [QueryProcessAffinityUpdateMode]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-queryprocessaffinityupdatemode
func (hProcess HPROCESS) QueryProcessAffinityUpdateMode() (co.AFFINITY, error) {
	var flags co.AFFINITY
	ret, _, err := syscall.SyscallN(_QueryProcessAffinityUpdateMode.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&flags)))
	if ret == 0 {
		return co.AFFINITY(0), co.ERROR(err)
	}
	return flags, nil
}

var _QueryProcessAffinityUpdateMode = dll.Kernel32.NewProc("QueryProcessAffinityUpdateMode")

// [QueryProcessCycleTime] function.
//
// [QueryProcessCycleTime]: https://learn.microsoft.com/en-us/windows/win32/api/realtimeapiset/nf-realtimeapiset-queryprocesscycletime
func (hProcess HPROCESS) QueryProcessCycleTime() (uint64, error) {
	var cycle uint64
	ret, _, err := syscall.SyscallN(_QueryProcessCycleTime.Addr(),
		uintptr(hProcess), uintptr(unsafe.Pointer(&cycle)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return cycle, nil
}

var _QueryProcessCycleTime = dll.Kernel32.NewProc("QueryProcessCycleTime")

// [ReadProcessMemory] function.
//
// [ReadProcessMemory]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory
func (hProcess HPROCESS) ReadProcessMemory(
	baseAddress uintptr,
	dest []byte,
) (numBytesRead uint, wErr error) {
	ret, _, err := syscall.SyscallN(_ReadProcessMemory.Addr(),
		uintptr(hProcess), baseAddress, uintptr(unsafe.Pointer(&dest[0])),
		uintptr(len(dest)), uintptr(unsafe.Pointer(&numBytesRead)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return numBytesRead, nil
}

var _ReadProcessMemory = dll.Kernel32.NewProc("ReadProcessMemory")

// [SetPriorityClass] function.
//
// [SetPriorityClass]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass
func (hProcess HPROCESS) SetPriorityClass(pc co.PRIORITY) error {
	ret, _, err := syscall.SyscallN(_SetPriorityClass.Addr(),
		uintptr(hProcess), uintptr(pc))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetPriorityClass = dll.Kernel32.NewProc("SetPriorityClass")

// [SetProcessAffinityUpdateMode] function.
//
// [SetProcessAffinityUpdateMode]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setprocessaffinityupdatemode
func (hProcess HPROCESS) SetProcessAffinityUpdateMode(affinity co.AFFINITY) error {
	ret, _, err := syscall.SyscallN(_SetProcessAffinityUpdateMode.Addr(),
		uintptr(hProcess), uintptr(affinity))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetProcessAffinityUpdateMode = dll.Kernel32.NewProc("SetProcessAffinityUpdateMode")

// [TerminateProcess] function.
//
// [TerminateProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminateprocess
func (hProcess HPROCESS) TerminateProcess(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(_TerminateProcess.Addr(),
		uintptr(hProcess), uintptr(exitCode))
	return utl.ZeroAsGetLastError(ret, err)
}

var _TerminateProcess = dll.Kernel32.NewProc("TerminateProcess")

// [WriteProcessMemory] function.
//
// [WriteProcessMemory]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory
func (hProcess HPROCESS) WriteProcessMemory(
	baseAddress uintptr,
	src []byte,
) (numBytesWritten uint, wErr error) {
	ret, _, err := syscall.SyscallN(_WriteProcessMemory.Addr(),
		uintptr(hProcess), baseAddress, uintptr(unsafe.Pointer(&src[0])),
		uintptr(len(src)), uintptr(unsafe.Pointer(&numBytesWritten)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return numBytesWritten, nil
}

var _WriteProcessMemory = dll.Kernel32.NewProc("WriteProcessMemory")
