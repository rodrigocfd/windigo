//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded psapi procedures.
var psapiCache [8]*syscall.Proc

// Loads psapi procedures.
func Psapi(procId PROC_PSAPI) uintptr {
	return LoadProc(SYSDLL_psapi, psapiCache[:], psapiProcStr, uint64(procId)).Addr()
}

type PROC_PSAPI uint64 // Procedure identifiers for psapi.

// Auto-generated psapi procedure identifier: cache index | str start | str past-end.
const (
	PROC_K32GetPerformanceInfo       PROC_PSAPI = 0 | (9 << 16) | (30 << 32)
	PROC_EmptyWorkingSet             PROC_PSAPI = 1 | (43 << 16) | (58 << 32)
	PROC_GetMappedFileNameW          PROC_PSAPI = 2 | (59 << 16) | (77 << 32)
	PROC_GetModuleBaseNameW          PROC_PSAPI = 3 | (78 << 16) | (96 << 32)
	PROC_GetModuleFileNameExW        PROC_PSAPI = 4 | (97 << 16) | (117 << 32)
	PROC_GetModuleInformation        PROC_PSAPI = 5 | (118 << 16) | (138 << 32)
	PROC_K32GetProcessImageFileNameW PROC_PSAPI = 6 | (139 << 16) | (166 << 32)
	PROC_GetProcessMemoryInfo        PROC_PSAPI = 7 | (167 << 16) | (187 << 32)
)

// Declaration of psapi procedure names.
const psapiProcStr = `
--funcs
K32GetPerformanceInfo

--hprocess
EmptyWorkingSet
GetMappedFileNameW
GetModuleBaseNameW
GetModuleFileNameExW
GetModuleInformation
K32GetProcessImageFileNameW
GetProcessMemoryInfo
`
