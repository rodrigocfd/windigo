//go:build windows

package dll

import (
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"
)

type SystemDll struct {
	dll  *syscall.DLL
	name string
}

var (
	dllMutex sync.Mutex

	Advapi  = SystemDll{nil, "advapi32"}
	Comctl  = SystemDll{nil, "comctl32"}
	Dwmapi  = SystemDll{nil, "dwmapi"}
	Gdi     = SystemDll{nil, "gdi32"}
	Kernel  = SystemDll{nil, "kernel32"}
	Ole     = SystemDll{nil, "ole32"}
	Oleaut  = SystemDll{nil, "oleaut32"}
	Psapi   = SystemDll{nil, "psapi"}
	Shell   = SystemDll{nil, "shell32"}
	Shlwapi = SystemDll{nil, "shlwapi"}
	User    = SystemDll{nil, "user32"}
	Userenv = SystemDll{nil, "userenv"}
	Uxtheme = SystemDll{nil, "uxtheme"}
	Version = SystemDll{nil, "version"}
)

func (me *SystemDll) Load(pDestProc **syscall.Proc, procName string) uintptr {
	if pProc := atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(pDestProc))); pProc != nil {
		return (*syscall.Proc)(pProc).Addr() // already cached
	}

	dllMutex.Lock()
	defer dllMutex.Unlock()

	if me.dll == nil {
		me.dll = syscall.MustLoadDLL(me.name)
	}

	*pDestProc = me.dll.MustFindProc(procName)
	return (*pDestProc).Addr()
}
