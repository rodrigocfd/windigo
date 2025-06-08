//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded version procedures.
var versionCache [3]*syscall.Proc

// Loads version procedures.
func Version(procId PROC_VERSION) uintptr {
	return LoadProc(SYSDLL_version, versionCache[:], versionProcStr, uint64(procId)).Addr()
}

type PROC_VERSION uint64 // Procedure identifiers for version.

// Auto-generated version procedure identifier: cache index | str start | str past-end.
const (
	PROC_GetFileVersionInfoW     PROC_VERSION = 0 | (9 << 16) | (28 << 32)
	PROC_GetFileVersionInfoSizeW PROC_VERSION = 1 | (29 << 16) | (52 << 32)
	PROC_VerQueryValueW          PROC_VERSION = 2 | (53 << 16) | (67 << 32)
)

// Declaration of version procedure names.
const versionProcStr = `
--funcs
GetFileVersionInfoW
GetFileVersionInfoSizeW
VerQueryValueW
`
