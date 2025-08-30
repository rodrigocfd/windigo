//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// Handle to a process [snapshot], which iterates over modules, processes and
// threads.
//
// [snapshot]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
type HPROCSNAP HANDLE

// [CreateToolhelp32Snapshot] function.
//
// ⚠️ You must defer [HPROCSNAP.CloseHandle].
//
// Example:
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPMODULE, 0)
//	defer hSnap.CloseHandle()
//
// [CreateToolhelp32Snapshot]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
func CreateToolhelp32Snapshot(flags co.TH32CS, processId uint32) (HPROCSNAP, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_CreateToolhelp32Snapshot, "CreateToolhelp32Snapshot"),
		uintptr(flags),
		uintptr(processId))
	if int(ret) == utl.INVALID_HANDLE_VALUE {
		return HPROCSNAP(0), co.ERROR(err)
	}
	return HPROCSNAP(ret), nil
}

var _CreateToolhelp32Snapshot *syscall.Proc

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcSnap HPROCSNAP) CloseHandle() error {
	return HANDLE(hProcSnap).CloseHandle()
}

// Returns the modules by calling [HPROCSNAP.Module32First] and
// [HPROCSNAP.Module32Next].
//
// Example:
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPMODULE, 0)
//	defer hSnap.CloseHandle()
//
//	modules, _ := hSnap.EnumModules()
//	for _, nfo := range modules {
//		println(nfo.SzExePath())
//	}
func (hProcSnap HPROCSNAP) EnumModules() ([]MODULEENTRY32, error) {
	modules := make([]MODULEENTRY32, 0)
	var me32 MODULEENTRY32
	me32.SetDwSize()

	found, err := hProcSnap.Module32First(&me32)

	for {
		if err != nil {
			return nil, err
		} else if !found {
			return modules, nil
		}
		modules = append(modules, me32)
		found, err = hProcSnap.Module32Next(&me32)
	}
}

// Returns the processes by calling [HPROCSNAP.Process32First] and
// [HPROCSNAP.Process32Next].
//
// Example:
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPPROCESS, 0)
//	defer hSnap.CloseHandle()
//
//	processes, _ := hSnap.EnumProcesses()
//	for _, nfo := range processes {
//		println(nfo.SzExeFile())
//	}
func (hProcSnap HPROCSNAP) EnumProcesses() ([]PROCESSENTRY32, error) {
	processes := make([]PROCESSENTRY32, 0)
	var pe32 PROCESSENTRY32
	pe32.SetDwSize()

	found, err := hProcSnap.Process32First(&pe32)

	for {
		if err != nil {
			return nil, err
		} else if !found {
			return processes, nil
		}
		processes = append(processes, pe32)
		found, err = hProcSnap.Process32Next(&pe32)
	}
}

// Returns the threads by calling [HPROCSNAP.Thread32First] and
// [HPROCSNAP.Thread32Next].
//
// Example:
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPTHREAD, 0)
//	defer hSnap.CloseHandle()
//
//	threads, _ := hSnap.EnumThreads()
//	for _, nfo := range threads {
//		println(nfo.Th32ThreadID)
//	}
func (hProcSnap HPROCSNAP) EnumThreads() ([]THREADENTRY32, error) {
	threads := make([]THREADENTRY32, 0)
	var te32 THREADENTRY32
	te32.SetDwSize()

	found, err := hProcSnap.Thread32First(&te32)

	for {
		if err != nil {
			return nil, err
		} else if !found {
			return threads, nil
		}
		threads = append(threads, te32)
		found, err = hProcSnap.Thread32Next(&te32)
	}
}

// [Module32First] function.
//
// This is a low-level function, prefer using [HPROCSNAP.EnumModules].
//
// [Module32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32firstw
func (hProcSnap HPROCSNAP) Module32First(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_Module32FirstW, "Module32FirstW"),
		uintptr(hProcSnap),
		uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

var _Module32FirstW *syscall.Proc

// [Module32Next] function.
//
// This is a low-level function, prefer using [HPROCSNAP.EnumModules].
//
// [Module32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32nextw
func (hProcSnap HPROCSNAP) Module32Next(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_Module32NextW, "Module32NextW"),
		uintptr(hProcSnap),
		uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

var _Module32NextW *syscall.Proc

// [Process32First] function.
//
// This is a low-level function, prefer using [HPROCSNAP.EnumProcesses].
//
// [Process32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAP) Process32First(buf *PROCESSENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_Process32FirstW, "Process32FirstW"),
		uintptr(hProcSnap),
		uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

var _Process32FirstW *syscall.Proc

// [Process32Next] function.
//
// This is a low-level function, prefer using [HPROCSNAP.EnumProcesses].
//
// [Process32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAP) Process32Next(buf *PROCESSENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_Process32NextW, "Process32NextW"),
		uintptr(hProcSnap),
		uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

var _Process32NextW *syscall.Proc

// [Thread32First] function.
//
// This is a low-level function, prefer using [HPROCSNAP.EnumThreads].
//
// [Thread32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32first
func (hProcSnap HPROCSNAP) Thread32First(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_Thread32First, "Thread32First"),
		uintptr(hProcSnap),
		uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}

var _Thread32First *syscall.Proc

// [Thread32Next] function.
//
// This is a low-level function, prefer using [HPROCSNAP.EnumThreads].
//
// [Thread32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32next
func (hProcSnap HPROCSNAP) Thread32Next(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(
		dll.Load(dll.KERNEL32, &_Thread32Next, "Thread32Next"),
		uintptr(hProcSnap),
		uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}

var _Thread32Next *syscall.Proc
