//go:build windows

package dll

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

const dllStr = `
advapi32
comctl32
dwmapi
gdi32
kernel32
ole32
oleaut32
psapi
shell32
shlwapi
user32
uxtheme
version
`

type SYSDLL uint64 // Identifies a system DLL to be loaded into the global cache.

// System DLL identifier: cache index | str start | str past-end.
const (
	SYSDLL_advapi32 SYSDLL = 0 | (1 << 16) | (9 << 32)
	SYSDLL_comctl32 SYSDLL = 1 | (10 << 16) | (18 << 32)
	SYSDLL_dwmapi   SYSDLL = 2 | (19 << 16) | (25 << 32)
	SYSDLL_gdi32    SYSDLL = 3 | (26 << 16) | (31 << 32)
	SYSDLL_kernel32 SYSDLL = 4 | (32 << 16) | (40 << 32)
	SYSDLL_ole32    SYSDLL = 5 | (41 << 16) | (46 << 32)
	SYSDLL_oleaut32 SYSDLL = 6 | (47 << 16) | (55 << 32)
	SYSDLL_psapi    SYSDLL = 7 | (56 << 16) | (61 << 32)
	SYSDLL_shell32  SYSDLL = 8 | (62 << 16) | (69 << 32)
	SYSDLL_shlwapi  SYSDLL = 9 | (70 << 16) | (77 << 32)
	SYSDLL_user32   SYSDLL = 10 | (78 << 16) | (84 << 32)
	SYSDLL_uxtheme  SYSDLL = 11 | (85 << 16) | (92 << 32)
	SYSDLL_version  SYSDLL = 12 | (93 << 16) | (100 << 32)
)

var (
	dllCache [13]*syscall.DLL // Stores the lazy-loaded system DLLs.
	dllMutex sync.Mutex       // For loading all DLLs and procedures.
)

// Loads a system DLL into the global cache.
func loadDll(dllId SYSDLL) *syscall.DLL {
	idx := uint64(dllId) & 0xffff
	if pDll := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&dllCache[idx]))); pDll != nil {
		return (*syscall.DLL)(pDll)
	}

	rangeStart := uint64(dllId) >> 16 & 0xffff
	rangePastEnd := uint64(dllId) >> 32 & 0xffff
	name := dllStr[rangeStart:rangePastEnd]

	dllMutex.Lock()
	defer dllMutex.Unlock()

	dllObj := syscall.MustLoadDLL(name)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&dllCache[idx])), unsafe.Pointer(dllObj))
	return dllObj
}

// Loads a procedure from a system DLL into the global cache.
func LoadProc(dllId SYSDLL, procCache []*syscall.Proc, procStr string, procId uint64) *syscall.Proc {
	idx := procId & 0xffff
	if pProc := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&procCache[idx]))); pProc != nil {
		return (*syscall.Proc)(pProc)
	}

	rangeStart := procId >> 16 & 0xffff
	rangePastEnd := procId >> 32 & 0xffff
	name := procStr[rangeStart:rangePastEnd]

	dllObj := loadDll(dllId)

	dllMutex.Lock()
	defer dllMutex.Unlock()

	proc := dllObj.MustFindProc(name)
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&procCache[idx])), unsafe.Pointer(proc))
	return proc
}
