//go:build windows

package dll

import (
	"syscall"
)

// ---
// Loading functions have repeated code because, with a single function, there
// would be 2 pointer indirections causing a ~9% increase in syscall times.
// ---

func Advapi(proc **syscall.Proc, name string) uintptr {
	if advapi32 == nil {
		advapi32 = syscall.MustLoadDLL("advapi32")
	}
	if *proc == nil {
		*proc = advapi32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var advapi32 *syscall.DLL

func Comctl(proc **syscall.Proc, name string) uintptr {
	if comctl32 == nil {
		comctl32 = syscall.MustLoadDLL("comctl32")
	}
	if *proc == nil {
		*proc = comctl32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var comctl32 *syscall.DLL

func Dwmapi(proc **syscall.Proc, name string) uintptr {
	if dwmapi == nil {
		dwmapi = syscall.MustLoadDLL("dwmapi")
	}
	if *proc == nil {
		*proc = dwmapi.MustFindProc(name)
	}
	return (*proc).Addr()
}

var dwmapi *syscall.DLL

func Gdi(proc **syscall.Proc, name string) uintptr {
	if gdi32 == nil {
		gdi32 = syscall.MustLoadDLL("gdi32")
	}
	if *proc == nil {
		*proc = gdi32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var gdi32 *syscall.DLL

func Kernel(proc **syscall.Proc, name string) uintptr {
	if kernel32 == nil {
		kernel32 = syscall.MustLoadDLL("kernel32")
	}
	if *proc == nil {
		*proc = kernel32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var kernel32 *syscall.DLL

func Ole(proc **syscall.Proc, name string) uintptr {
	if ole32 == nil {
		ole32 = syscall.MustLoadDLL("ole32")
	}
	if *proc == nil {
		*proc = ole32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var ole32 *syscall.DLL

func Oleaut(proc **syscall.Proc, name string) uintptr {
	if oleaut32 == nil {
		oleaut32 = syscall.MustLoadDLL("oleaut32")
	}
	if *proc == nil {
		*proc = oleaut32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var oleaut32 *syscall.DLL

func Psapi(proc **syscall.Proc, name string) uintptr {
	if psapi == nil {
		psapi = syscall.MustLoadDLL("psapi")
	}
	if *proc == nil {
		*proc = psapi.MustFindProc(name)
	}
	return (*proc).Addr()
}

var psapi *syscall.DLL

func Shell(proc **syscall.Proc, name string) uintptr {
	if shell32 == nil {
		shell32 = syscall.MustLoadDLL("shell32")
	}
	if *proc == nil {
		*proc = shell32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var shell32 *syscall.DLL

func Shlwapi(proc **syscall.Proc, name string) uintptr {
	if shlwapi == nil {
		shlwapi = syscall.MustLoadDLL("shlwapi")
	}
	if *proc == nil {
		*proc = shlwapi.MustFindProc(name)
	}
	return (*proc).Addr()
}

var shlwapi *syscall.DLL

func User(proc **syscall.Proc, name string) uintptr {
	if user32 == nil {
		user32 = syscall.MustLoadDLL("user32")
	}
	if *proc == nil {
		*proc = user32.MustFindProc(name)
	}
	return (*proc).Addr()
}

var user32 *syscall.DLL

func Uxtheme(proc **syscall.Proc, name string) uintptr {
	if uxtheme == nil {
		uxtheme = syscall.MustLoadDLL("uxtheme")
	}
	if *proc == nil {
		*proc = uxtheme.MustFindProc(name)
	}
	return (*proc).Addr()
}

var uxtheme *syscall.DLL

func Version(proc **syscall.Proc, name string) uintptr {
	if version == nil {
		version = syscall.MustLoadDLL("version")
	}
	if *proc == nil {
		*proc = version.MustFindProc(name)
	}
	return (*proc).Addr()
}

var version *syscall.DLL
