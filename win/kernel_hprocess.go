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
	ret, _, _ := syscall.SyscallN(dll.Kernel(dll.PROC_GetCurrentProcess))
	return HPROCESS(ret)
}

// [OpenProcess] function.
//
// ⚠️ You must defer [HPROCESS.CloseHandle].
//
// [OpenProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-openprocess
func OpenProcess(access co.PROCESS, inheritHandle bool, processId uint32) (HPROCESS, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_OpenProcess),
		uintptr(access),
		utl.BoolToUintptr(inheritHandle),
		uintptr(processId))
	if ret == 0 {
		return HPROCESS(0), co.ERROR(err)
	}
	return HPROCESS(ret), nil
}

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
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetExitCodeProcess),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&exitCode)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return exitCode, nil
}

// [GetPriorityClass] function.
//
// [GetPriorityClass]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getpriorityclass
func (hProcess HPROCESS) GetPriorityClass() (co.PRIORITY, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetPriorityClass),
		uintptr(hProcess))
	if ret == 0 {
		return co.PRIORITY(0), co.ERROR(err)
	}
	return co.PRIORITY(ret), nil
}

// [GetProcessHandleCount] function.
//
// [GetProcessHandleCount]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesshandlecount
func (hProcess HPROCESS) GetProcessHandleCount() (uint, error) {
	var count uint32
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessHandleCount),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&count)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint(count), nil
}

// [GetProcessId] function.
//
// [GetProcessId]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessid
func (hProcess HPROCESS) GetProcessId() (uint32, error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessId),
		uintptr(hProcess))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return uint32(ret), nil
}

// [GetProcessPriorityBoost] function.
//
// [GetProcessPriorityBoost]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesspriorityboost
func (hProcess HPROCESS) GetProcessPriorityBoost() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessPriorityBoost),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

// [GetProcessShutdownParameters] function.
//
// [GetProcessShutdownParameters]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocessshutdownparameters
func (hProcess HPROCESS) GetProcessShutdownParameters() (priorityLevel uint32, flag co.SHUTDOWN, wErr error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessShutdownParameters),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&priorityLevel)),
		uintptr(unsafe.Pointer(&flag)))
	if ret == 0 {
		return 0, co.SHUTDOWN(0), co.ERROR(err)
	}
	return
}

// [GetProcessTimes] function.
//
// [GetProcessTimes]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getprocesstimes
func (hProcess HPROCESS) GetProcessTimes() (HprocessTimes, error) {
	var ftCreation, ftExit, ftKernel, ftUser FILETIME
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessTimes),
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
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_GetProcessVersion),
		uintptr(hProcess))
	if ret == 0 {
		return 0, 0, co.ERROR(err)
	}
	return HIWORD(uint32(ret)), LOWORD(uint32(ret)), nil
}

// [IsProcessCritical] function.
//
// [IsProcessCritical]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-isprocesscritical
func (hProcess HPROCESS) IsProcessCritical() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_IsProcessCritical),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

// [IsWow64Process] function.
//
// [IsWow64Process]: https://learn.microsoft.com/en-us/windows/win32/api/wow64apiset/nf-wow64apiset-iswow64process
func (hProcess HPROCESS) IsWow64Process() (bool, error) {
	var bVal int32 // BOOL
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_IsWow64Process),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&bVal)))
	if ret == 0 {
		return false, co.ERROR(err)
	}
	return bVal != 0, nil
}

// [QueryFullProcessImageName] function.
//
// [QueryFullProcessImageName]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-queryfullprocessimagenamew
func (hProcess HPROCESS) QueryFullProcessImageName(flags co.PROCESS_NAME) (string, error) {
	recvBuf := wstr.NewBufReceiver(wstr.BUF_MAX)
	defer recvBuf.Free()

	szBuf := uint32(recvBuf.Len())

	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_QueryFullProcessImageNameW),
		uintptr(hProcess),
		uintptr(flags),
		uintptr(recvBuf.UnsafePtr()),
		uintptr(unsafe.Pointer(&szBuf)))
	if ret == 0 {
		return "", co.ERROR(err)
	}
	return recvBuf.String(), nil
}

// [QueryProcessAffinityUpdateMode] function.
//
// [QueryProcessAffinityUpdateMode]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-queryprocessaffinityupdatemode
func (hProcess HPROCESS) QueryProcessAffinityUpdateMode() (co.AFFINITY, error) {
	var flags co.AFFINITY
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_QueryProcessAffinityUpdateMode),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&flags)))
	if ret == 0 {
		return co.AFFINITY(0), co.ERROR(err)
	}
	return flags, nil
}

// [QueryProcessCycleTime] function.
//
// [QueryProcessCycleTime]: https://learn.microsoft.com/en-us/windows/win32/api/realtimeapiset/nf-realtimeapiset-queryprocesscycletime
func (hProcess HPROCESS) QueryProcessCycleTime() (uint64, error) {
	var cycle uint64
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_QueryProcessCycleTime),
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&cycle)))
	if ret == 0 {
		return 0, co.ERROR(err)
	}
	return cycle, nil
}

// [ReadProcessMemory] function.
//
// [ReadProcessMemory]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-readprocessmemory
func (hProcess HPROCESS) ReadProcessMemory(
	baseAddress uintptr,
	dest []byte,
) (numBytesRead uint, wErr error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_ReadProcessMemory),
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

// [SetPriorityClass] function.
//
// [SetPriorityClass]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setpriorityclass
func (hProcess HPROCESS) SetPriorityClass(pc co.PRIORITY) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetPriorityClass),
		uintptr(hProcess),
		uintptr(pc))
	return utl.ZeroAsGetLastError(ret, err)
}

// [SetProcessAffinityUpdateMode] function.
//
// [SetProcessAffinityUpdateMode]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-setprocessaffinityupdatemode
func (hProcess HPROCESS) SetProcessAffinityUpdateMode(affinity co.AFFINITY) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_SetProcessAffinityUpdateMode),
		uintptr(hProcess),
		uintptr(affinity))
	return utl.ZeroAsGetLastError(ret, err)
}

// [TerminateProcess] function.
//
// [TerminateProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminateprocess
func (hProcess HPROCESS) TerminateProcess(exitCode uint32) error {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_TerminateProcess),
		uintptr(hProcess),
		uintptr(exitCode))
	return utl.ZeroAsGetLastError(ret, err)
}

// [VirtualQueryEx] function.
//
// [VirtualQueryEx]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-virtualqueryex
func (hProcess HPROCESS) VirtualQueryEx(baseAddress uintptr) (MEMORY_BASIC_INFORMATION, error) {
	var mbi MEMORY_BASIC_INFORMATION
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_VirtualQueryEx),
		uintptr(hProcess),
		baseAddress,
		uintptr(unsafe.Pointer(&mbi)),
		uintptr(uint32(unsafe.Sizeof(mbi))))
	if ret == 0 {
		return MEMORY_BASIC_INFORMATION{}, co.ERROR(err)
	}
	return mbi, nil
}

// [WriteProcessMemory] function.
//
// [WriteProcessMemory]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-writeprocessmemory
func (hProcess HPROCESS) WriteProcessMemory(
	baseAddress uintptr,
	src []byte,
) (numBytesWritten uint, wErr error) {
	ret, _, err := syscall.SyscallN(dll.Kernel(dll.PROC_WriteProcessMemory),
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
