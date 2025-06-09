//go:build windows

package shell

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// Stores the loazy-loaded shell procedures.
var shellCache [8]*syscall.Proc

// Loads shell procedures.
func dllShell(procId _PROC_SHELL) uintptr {
	return dll.LoadProc(dll.SYSDLL_shell32, shellCache[:], shellProcStr, uint64(procId)).Addr()
}

type _PROC_SHELL uint64 // Procedure identifiers for shell.

// Auto-generated shell procedure identifier: cache index | str start | str past-end.
const (
	_PROC_SHCreateItemFromIDList            _PROC_SHELL = 0 | (9 << 16) | (31 << 32)
	_PROC_SHCreateItemFromParsingName       _PROC_SHELL = 1 | (32 << 16) | (59 << 32)
	_PROC_SHCreateItemFromRelativeName      _PROC_SHELL = 2 | (60 << 16) | (88 << 32)
	_PROC_SHCreateShellItemArray            _PROC_SHELL = 3 | (89 << 16) | (111 << 32)
	_PROC_SHCreateShellItemArrayFromIDLists _PROC_SHELL = 4 | (112 << 16) | (145 << 32)
	_PROC_SHGetDesktopFolder                _PROC_SHELL = 5 | (146 << 16) | (164 << 32)
	_PROC_SHGetKnownFolderItem              _PROC_SHELL = 6 | (165 << 16) | (185 << 32)
	_PROC_SHGetIDListFromObject             _PROC_SHELL = 7 | (186 << 16) | (207 << 32)
)

// Declaration of shell procedure names.
const shellProcStr = `
--funcs
SHCreateItemFromIDList
SHCreateItemFromParsingName
SHCreateItemFromRelativeName
SHCreateShellItemArray
SHCreateShellItemArrayFromIDLists
SHGetDesktopFolder
SHGetKnownFolderItem
SHGetIDListFromObject
`
