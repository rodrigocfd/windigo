//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded shell procedures.
var shellCache [8]*syscall.Proc

// Loads shell procedures.
func Shell(procId PROC_SHELL) uintptr {
	return LoadProc(SYSDLL_shell32, shellCache[:], shellProcStr, uint64(procId)).Addr()
}

type PROC_SHELL uint64 // Procedure identifiers for shell.

// Auto-generated shell procedure identifier: cache index | str start | str past-end.
const (
	PROC_CommandLineToArgvW      PROC_SHELL = 0 | (9 << 16) | (27 << 32)
	PROC_Shell_NotifyIconW       PROC_SHELL = 1 | (28 << 16) | (45 << 32)
	PROC_Shell_NotifyIconGetRect PROC_SHELL = 2 | (46 << 16) | (69 << 32)
	PROC_SHGetFileInfoW          PROC_SHELL = 3 | (70 << 16) | (84 << 32)
	PROC_DragFinish              PROC_SHELL = 4 | (94 << 16) | (104 << 32)
	PROC_DragQueryFileW          PROC_SHELL = 5 | (105 << 16) | (119 << 32)
	PROC_DragQueryPoint          PROC_SHELL = 6 | (120 << 16) | (134 << 32)
	PROC_ShellAboutW             PROC_SHELL = 7 | (143 << 16) | (154 << 32)
)

// Declaration of shell procedure names.
const shellProcStr = `
--funcs
CommandLineToArgvW
Shell_NotifyIconW
Shell_NotifyIconGetRect
SHGetFileInfoW

--hdrop
DragFinish
DragQueryFileW
DragQueryPoint

--hwnd
ShellAboutW
`
