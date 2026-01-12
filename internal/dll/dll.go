//go:build windows

package dll

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

type DLL_INDEX int // Indexes globals dllCache and dllNames.

const (
	ADVAPI32 DLL_INDEX = iota
	COMCTL32
	DWMAPI
	GDI32
	KERNEL32
	OLE32
	OLEAUT32
	PSAPI
	SHELL32
	SHLWAPI
	USER32
	USERENV
	UXTHEME
	VERSION
)

var (
	dllCache [14]*syscall.DLL // Indexed by DLL_INDEX.
	dllMutex sync.Mutex
	dllNames = [...]string{ // Indexed by DLL_INDEX.
		"advapi32",
		"comctl32",
		"dwmapi",
		"gdi32",
		"kernel32",
		"ole32",
		"oleaut32",
		"psapi",
		"shell32",
		"shlwapi",
		"user32",
		"userenv",
		"uxtheme",
		"version",
	}
)

func loadDll(idx DLL_INDEX, name string) *syscall.DLL {
	if pDll := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&dllCache[idx]))); pDll != nil {
		return (*syscall.DLL)(pDll)
	}

	dllMutex.Lock()
	defer dllMutex.Unlock()

	dllObj := syscall.MustLoadDLL(name)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&dllCache[idx])), unsafe.Pointer(dllObj))
	return dllObj
}

// Dynamically loads a procedure from a system DLL.
func Load(dllIdx DLL_INDEX, pDestProc **syscall.Proc, procName string) uintptr {
	if pProc := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pDestProc))); pProc != nil {
		return (*syscall.Proc)(pProc).Addr()
	}

	dllObj := loadDll(dllIdx, dllNames[dllIdx])

	dllMutex.Lock()
	defer dllMutex.Unlock()

	proc := dllObj.MustFindProc(procName)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(pDestProc)), unsafe.Pointer(proc))
	return proc.Addr()
}
