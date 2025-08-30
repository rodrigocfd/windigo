//go:build windows

package win

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// Handle to a process.
type HPROCESS HANDLE

// [GetCurrentProcess] function.
//
// [GetCurrentProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocess
func GetCurrentProcess() HPROCESS {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetCurrentProcess, "GetCurrentProcess"))
	return HPROCESS(ret)
}

var _GetCurrentProcess *syscall.Proc

// [OpenProcess] function.
//
// ⚠️ You must defer [HPROCESS.CloseHandle].
//
// [OpenProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess
func OpenProcess(access co.PROCESS, inheritHandle bool, processId uint32) (HPROCESS, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_OpenProcess, "OpenProcess"),
		uintptr(access),
		utl.BoolToUintptr(inheritHandle),
		uintptr(processId))
	if ret == 0 {
		return HPROCESS(0), co.ERROR(err)
	}
	return HPROCESS(ret), nil
}

var _OpenProcess *syscall.Proc

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcess HPROCESS) CloseHandle() error {
	return HANDLE(hProcess).CloseHandle()
}

// [FlushInstructionCache] function.
//
// [FlushInstructionCache]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-flushinstructioncache
func (hProcess HPROCESS) FlushInstructionCache(baseAddress uintptr, size uint) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_FlushInstructionCache, "FlushInstructionCache"),
		uintptr(hProcess),
		baseAddress,
		uintptr(size))
	return utl.ZeroAsGetLastError(ret, err)
}

var _FlushInstructionCache *syscall.Proc

// [GetExitCodeProcess] function.
//
// [GetExitCodeProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
func (hProcess HPROCESS) GetExitCodeProcess() (uint32, error) {
	var exitCode uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetExitCodeProcess, "GetExitCodeProcess"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return exitCode, nil
}

var _GetExitCodeProcess *syscall.Proc

// [GetPriorityClass] function.
//
// [GetPriorityClass]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getpriorityclass
func (hProcess HPROCESS) GetPriorityClass() (co.PRIORITY, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetPriorityClass, "GetPriorityClass"),
		uintptr(hProcess))
	if ret == 0 {
		return co.PRIORITY(0), co.ERROR(err)
	}
	return co.PRIORITY(ret), nil
}

var _GetPriorityClass *syscall.Proc

// [GetProcessHandleCount] function.
//
// [GetProcessHandleCount]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesshandlecount
func (hProcess HPROCESS) GetProcessHandleCount() (uint, error) {
	var count uint32
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessHandleCount, "GetProcessHandleCount"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&count)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(count), nil
}

var _GetProcessHandleCount *syscall.Proc

// [GetProcessId] function.
//
// [GetProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessid
func (hProcess HPROCESS) GetProcessId() (uint32, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessId, "GetProcessId"),
		uintptr(hProcess))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

var _GetProcessId *syscall.Proc

// [GetProcessPriorityBoost] function.
//
// [GetProcessPriorityBoost]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesspriorityboost
func (hProcess HPROCESS) GetProcessPriorityBoost() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessPriorityBoost, "GetProcessPriorityBoost"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _GetProcessPriorityBoost *syscall.Proc

// [GetProcessShutdownParameters] function.
//
// [GetProcessShutdownParameters]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessshutdownparameters
func (hProcess HPROCESS) GetProcessShutdownParameters() (priorityLevel uint32, flag co.SHUTDOWN, wErr error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessShutdownParameters, "GetProcessShutdownParameters"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&priorityLevel)),
		uintptr(unsafe.Pointer(&flag)))
	if ret == 0 {
		return 0, co.SHUTDOWN(0), co.ERROR(err)
	}
	return
}

var _GetProcessShutdownParameters *syscall.Proc

// [GetProcessTimes] function.
//
// [GetProcessTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes
func (hProcess HPROCESS) GetProcessTimes() (HprocessTimes, error) {
	var ftCreation, ftExit, ftKernel, ftUser FILETIME
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessTimes, "GetProcessTimes"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&ftCreation)),
		uintptr(unsafe.Pointer(&ftExit)),
		uintptr(unsafe.Pointer(&ftKernel)),
		uintptr(unsafe.Pointer(&ftUser)))
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

var _GetProcessTimes *syscall.Proc

// Returned by [HPROCESS.GetProcessTimes].
type HprocessTimes struct {
	Creation time.Time
	Exit     time.Time
	Kernel   time.Time
	User     time.Time
}

// [GetProcessVersion] function.
//
// [GetProcessVersion]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessversion
func (hProcess HPROCESS) GetProcessVersion() (maj, min uint16, wErr error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_GetProcessVersion, "GetProcessVersion"),
		uintptr(hProcess))
	if ret == 0 {
		return 0, 0, co.ERROR(err)
	}
	return HIWORD(uint32(ret)), LOWORD(uint32(ret)), nil
}

var _GetProcessVersion *syscall.Proc

// [IsProcessCritical] function.
//
// [IsProcessCritical]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-isprocesscritical
func (hProcess HPROCESS) IsProcessCritical() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_IsProcessCritical, "IsProcessCritical"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _IsProcessCritical *syscall.Proc

// [IsWow64Process] function.
//
// [IsWow64Process]: https://learn.microsoft.com/en-us/windows/win32/api/wow64apiset/nf-wow64apiset-iswow64process
func (hProcess HPROCESS) IsWow64Process() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_IsWow64Process, "IsWow64Process"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

var _IsWow64Process *syscall.Proc

// [QueryFullProcessImageName] function.
//
// [QueryFullProcessImageName]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-queryfullprocessimagenamew
func (hProcess HPROCESS) QueryFullProcessImageName(flags co.PROCESS_NAME) (string, error) {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)

	szBuf := uint32(wBuf.Len())

	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_QueryFullProcessImageNameW, "QueryFullProcessImageNameW"),
		uintptr(hProcess),
		uintptr(flags),
		uintptr(wBuf.Ptr()),
		uintptr(unsafe.Pointer(&szBuf)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return wBuf.String(), nil
}

var _QueryFullProcessImageNameW *syscall.Proc

// [QueryProcessAffinityUpdateMode] function.
//
// [QueryProcessAffinityUpdateMode]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-queryprocessaffinityupdatemode
func (hProcess HPROCESS) QueryProcessAffinityUpdateMode() (co.AFFINITY, error) {
	var flags co.AFFINITY
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_QueryProcessAffinityUpdateMode, "QueryProcessAffinityUpdateMode"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&flags)))
	if ret == 0 {
		return co.AFFINITY(0), co.ERROR(err)
	}
	return flags, nil
}

var _QueryProcessAffinityUpdateMode *syscall.Proc

// [QueryProcessCycleTime] function.
//
// [QueryProcessCycleTime]: https://learn.microsoft.com/en-us/windows/win32/api/realtimeapiset/nf-realtimeapiset-queryprocesscycletime
func (hProcess HPROCESS) QueryProcessCycleTime() (uint64, error) {
	var cycle uint64
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_QueryProcessCycleTime, "QueryProcessCycleTime"),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&cycle)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return cycle, nil
}

var _QueryProcessCycleTime *syscall.Proc

// [ReadProcessMemory] function.
//
// [ReadProcessMemory]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory
func (hProcess HPROCESS) ReadProcessMemory(
	baseAddress uintptr,
	dest []byte,
) (numBytesRead uint, wErr error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_ReadProcessMemory, "ReadProcessMemory"),
		uintptr(hProcess),
		baseAddress,
		uintptr(unsafe.Pointer(&dest[0])),
		uintptr(len(dest)),
		uintptr(unsafe.Pointer(&numBytesRead)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return numBytesRead, nil
}

var _ReadProcessMemory *syscall.Proc

// [SetPriorityClass] function.
//
// [SetPriorityClass]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass
func (hProcess HPROCESS) SetPriorityClass(pc co.PRIORITY) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SetPriorityClass, "SetPriorityClass"),
		uintptr(hProcess),
		uintptr(pc))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetPriorityClass *syscall.Proc

// [SetProcessAffinityUpdateMode] function.
//
// [SetProcessAffinityUpdateMode]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setprocessaffinityupdatemode
func (hProcess HPROCESS) SetProcessAffinityUpdateMode(affinity co.AFFINITY) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_SetProcessAffinityUpdateMode, "SetProcessAffinityUpdateMode"),
		uintptr(hProcess),
		uintptr(affinity))
	return utl.ZeroAsGetLastError(ret, err)
}

var _SetProcessAffinityUpdateMode *syscall.Proc

// [TerminateProcess] function.
//
// [TerminateProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminateprocess
func (hProcess HPROCESS) TerminateProcess(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_TerminateProcess, "TerminateProcess"),
		uintptr(hProcess),
		uintptr(exitCode))
	return utl.ZeroAsGetLastError(ret, err)
}

var _TerminateProcess *syscall.Proc

// [VirtualQueryEx] function.
//
// [VirtualQueryEx]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-virtualqueryex
func (hProcess HPROCESS) VirtualQueryEx(baseAddress uintptr) (MEMORY_BASIC_INFORMATION, error) {
	var mbi MEMORY_BASIC_INFORMATION
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_VirtualQueryEx, "VirtualQueryEx"),
		uintptr(hProcess),
		baseAddress,
		uintptr(unsafe.Pointer(&mbi)),
		uintptr(uint32(unsafe.Sizeof(mbi))))
	if ret == 0 {
		return MEMORY_BASIC_INFORMATION{}, co.ERROR(err)
	}
	return mbi, nil
}

var _VirtualQueryEx *syscall.Proc

// [WriteProcessMemory] function.
//
// [WriteProcessMemory]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory
func (hProcess HPROCESS) WriteProcessMemory(
	baseAddress uintptr,
	src []byte,
) (numBytesWritten uint, wErr error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_WriteProcessMemory, "WriteProcessMemory"),
		uintptr(hProcess),
		baseAddress,
		uintptr(unsafe.Pointer(&src[0])),
		uintptr(len(src)),
		uintptr(unsafe.Pointer(&numBytesWritten)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return numBytesWritten, nil
}

var _WriteProcessMemory *syscall.Proc
