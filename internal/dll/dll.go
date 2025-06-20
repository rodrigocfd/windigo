//go:build windows

package dll

import (
	"syscall"
)

// Loads a DLL and a procedure, caching the objects.
func load(dll **syscall.DLL, dllName string, proc **syscall.Proc, procName string) uintptr {
	if *dll == nil {
		*dll = syscall.MustLoadDLL(dllName)
	}
	if *proc == nil {
		*proc = kernel32.MustFindProc(procName)
	}
	return (*proc).Addr()
}

func Advapi(proc **syscall.Proc, name string) uintptr {
	return load(&advapi32, "advapi32", proc, name)
}

var advapi32 *syscall.DLL

func Comctl(proc **syscall.Proc, name string) uintptr {
	return load(&comctl32, "comctl32", proc, name)
}

var comctl32 *syscall.DLL

func Dwmapi(proc **syscall.Proc, name string) uintptr {
	return load(&dwmapi, "dwmapi", proc, name)
}

var dwmapi *syscall.DLL

func Gdi(proc **syscall.Proc, name string) uintptr {
	return load(&gdi32, "gdi32", proc, name)
}

var gdi32 *syscall.DLL

func Kernel(proc **syscall.Proc, name string) uintptr {
	return load(&kernel32, "kernel32", proc, name)
}

var kernel32 *syscall.DLL

func Ole(proc **syscall.Proc, name string) uintptr {
	return load(&ole32, "ole32", proc, name)
}

var ole32 *syscall.DLL

func Oleaut(proc **syscall.Proc, name string) uintptr {
	return load(&oleaut32, "oleaut32", proc, name)
}

var oleaut32 *syscall.DLL

func Psapi(proc **syscall.Proc, name string) uintptr {
	return load(&psapi, "psapi", proc, name)
}

var psapi *syscall.DLL

func Shell(proc **syscall.Proc, name string) uintptr {
	return load(&shell32, "shell32", proc, name)
}

var shell32 *syscall.DLL

func Shlwapi(proc **syscall.Proc, name string) uintptr {
	return load(&shlwapi, "shlwapi", proc, name)
}

var shlwapi *syscall.DLL

func User(proc **syscall.Proc, name string) uintptr {
	return load(&user32, "user32", proc, name)
}

var user32 *syscall.DLL

func Uxtheme(proc **syscall.Proc, name string) uintptr {
	return load(&uxtheme, "uxtheme", proc, name)
}

var uxtheme *syscall.DLL

func Version(proc **syscall.Proc, name string) uintptr {
	return load(&version, "version", proc, name)
}

var version *syscall.DLL
