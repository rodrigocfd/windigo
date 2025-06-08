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
	_PROC_SHCreateItemFromIDList            _PROC_SHELL = 0 | (1 << 16) | (23 << 32)
	_PROC_SHCreateItemFromParsingName       _PROC_SHELL = 1 | (24 << 16) | (51 << 32)
	_PROC_SHCreateItemFromRelativeName      _PROC_SHELL = 2 | (52 << 16) | (80 << 32)
	_PROC_SHCreateShellItemArray            _PROC_SHELL = 3 | (81 << 16) | (103 << 32)
	_PROC_SHCreateShellItemArrayFromIDLists _PROC_SHELL = 4 | (104 << 16) | (137 << 32)
	_PROC_SHGetDesktopFolder                _PROC_SHELL = 5 | (138 << 16) | (156 << 32)
	_PROC_SHGetKnownFolderItem              _PROC_SHELL = 6 | (157 << 16) | (177 << 32)
	_PROC_SHGetIDListFromObject             _PROC_SHELL = 7 | (178 << 16) | (199 << 32)
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
