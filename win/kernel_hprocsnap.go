//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/wutil"
	"github.com/rodrigocfd/windigo/win/co"
)

// Handle to a process [snapshot], which iterates over modules, processes and
// threads.
//
// [snapshot]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
type HPROCSNAP HANDLE

// [CreateToolhelp32Snapshot] function.
//
// ⚠️ You must defer HPROCSNAP.CloseHandle().
//
// [CreateToolhelp32Snapshot]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
func CreateToolhelp32Snapshot(flags co.TH32CS, processId uint32) (HPROCSNAP, error) {
	ret, _, err := syscall.SyscallN(_CreateToolhelp32Snapshot.Addr(),
		uintptr(flags), uintptr(processId))
	if int(ret) == wutil.INVALID_HANDLE_VALUE {
		return HPROCSNAP(0), co.ERROR(err)
	}
	return HPROCSNAP(ret), nil
}

var _CreateToolhelp32Snapshot = dll.Kernel32.NewProc("CreateToolhelp32Snapshot")

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hProcSnap HPROCSNAP) CloseHandle() error {
	return HANDLE(hProcSnap).CloseHandle()
}

// Returns the modules by calling [Module32First] and [Module32Next].
//
// # Example
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPMODULE, 0)
//	defer hSnap.CloseHandle()
//
//	modules, _ := hSnap.EnumModules()
//	for nfo, _ := range modules {
//		println(nfo.SzExePath())
//	}
//
// [Module32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32firstw
// [Module32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32nextw
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

// Returns the processes by calling [Process32First] and [Process32Next].
//
// # Example
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPPROCESS, 0)
//	defer hSnap.CloseHandle()
//
//	processes, _ := hSnap.EnumProcesses()
//	for nfo, _ := range processes {
//		println(nfo.SzExeFile())
//	}
//
// [Process32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
// [Process32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
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

// Returns the threads by calling [Thread32First] and [Thread32Next].
//
// # Example
//
//	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPTHREAD, 0)
//	defer hSnap.CloseHandle()
//
//	threads, _ := hSnap.EnumThreads()
//	for nfo, _ := range threads {
//		println(nfo.Th32ThreadID)
//	}
//
// [Thread32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32first
// [Thread32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32next
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
// This is a low level function, prefer using HPROCSNAP.IterModules().
//
// [Module32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32firstw
func (hProcSnap HPROCSNAP) Module32First(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(_Module32FirstW.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

var _Module32FirstW = dll.Kernel32.NewProc("Module32FirstW")

// [Module32Next] function.
//
// This is a low level function, prefer using HPROCSNAP.IterModules().
//
// [Module32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-module32nextw
func (hProcSnap HPROCSNAP) Module32Next(buf *MODULEENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(_Module32NextW.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a module was found
}

var _Module32NextW = dll.Kernel32.NewProc("Module32NextW")

// [Process32First] function.
//
// This is a low level function, prefer using HPROCSNAP.IterProcesses().
//
// [Process32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAP) Process32First(buf *PROCESSENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(_Process32FirstW.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

var _Process32FirstW = dll.Kernel32.NewProc("Process32FirstW")

// [Process32Next] function.
//
// This is a low level function, prefer using HPROCSNAP.IterProcesses().
//
// [Process32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-process32firstw
func (hProcSnap HPROCSNAP) Process32Next(buf *PROCESSENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(_Process32NextW.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a process was found
}

var _Process32NextW = dll.Kernel32.NewProc("Process32NextW")

// [Thread32First] function.
//
// This is a low level function, prefer using HPROCSNAP.IterThreads().
//
// [Thread32First]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32first
func (hProcSnap HPROCSNAP) Thread32First(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(_Thread32First.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}

var _Thread32First = dll.Kernel32.NewProc("Thread32First")

// [Thread32Next] function.
//
// This is a low level function, prefer using HPROCSNAP.IterThreads().
//
// [Thread32Next]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-thread32next
func (hProcSnap HPROCSNAP) Thread32Next(buf *THREADENTRY32) (bool, error) {
	ret, _, err := syscall.SyscallN(_Thread32Next.Addr(),
		uintptr(hProcSnap), uintptr(unsafe.Pointer(buf)))
	if ret == 0 {
		if wErr := co.ERROR(err); wErr == co.ERROR_NO_MORE_FILES {
			return false, nil // not an error, search ended
		} else {
			return false, wErr
		}
	}
	return true, nil // a thread was found
}

var _Thread32Next = dll.Kernel32.NewProc("Thread32Next")
