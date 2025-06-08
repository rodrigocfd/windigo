//go:build windows

package ole

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// Stores the loazy-loaded shlwapi procedures.
var shlwapiCache [1]*syscall.Proc

// Loads shlwapi procedures.
func dllShlwapi(procId _PROC_SHLWAPI) uintptr {
	return dll.LoadProc(dll.SYSDLL_shlwapi, shlwapiCache[:], shlwapiProcStr, uint64(procId)).Addr()
}

type _PROC_SHLWAPI uint64 // Procedure identifiers for shlwapi.

// Auto-generated shlwapi procedure identifier: cache index | str start | str past-end.
const (
	_PROC_SHCreateMemStream _PROC_SHLWAPI = 0 | (9 << 16) | (26 << 32)
)

// Declaration of shlwapi procedure names.
const shlwapiProcStr = `
--funcs
SHCreateMemStream
`
