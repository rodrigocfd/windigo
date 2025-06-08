//go:build windows

package dll

import (
	"syscall"
)

// Stores the loazy-loaded advapi procedures.
var advapiCache [32]*syscall.Proc

// Loads advapi procedures.
func Advapi(procId PROC_ADVAPI) uintptr {
	return LoadProc(SYSDLL_advapi32, advapiCache[:], advapiProcStr, uint64(procId)).Addr()
}

type PROC_ADVAPI uint64 // Procedure identifiers for advapi.

// Auto-generated advapi procedure identifier: cache index | str start | str past-end.
const (
	PROC_RegDisablePredefinedCache   PROC_ADVAPI = 0 | (9 << 16) | (34 << 32)
	PROC_RegDisablePredefinedCacheEx PROC_ADVAPI = 1 | (35 << 16) | (62 << 32)
	PROC_RegConnectRegistryW         PROC_ADVAPI = 2 | (71 << 16) | (90 << 32)
	PROC_RegOpenCurrentUser          PROC_ADVAPI = 3 | (91 << 16) | (109 << 32)
	PROC_RegCloseKey                 PROC_ADVAPI = 4 | (110 << 16) | (121 << 32)
	PROC_RegCopyTreeW                PROC_ADVAPI = 5 | (122 << 16) | (134 << 32)
	PROC_RegCreateKeyExW             PROC_ADVAPI = 6 | (135 << 16) | (150 << 32)
	PROC_RegDeleteKeyW               PROC_ADVAPI = 7 | (151 << 16) | (164 << 32)
	PROC_RegDeleteKeyExW             PROC_ADVAPI = 8 | (165 << 16) | (180 << 32)
	PROC_RegDeleteKeyValueW          PROC_ADVAPI = 9 | (181 << 16) | (199 << 32)
	PROC_RegDeleteTreeW              PROC_ADVAPI = 10 | (200 << 16) | (214 << 32)
	PROC_RegDeleteValueW             PROC_ADVAPI = 11 | (215 << 16) | (230 << 32)
	PROC_RegDisableReflectionKey     PROC_ADVAPI = 12 | (231 << 16) | (254 << 32)
	PROC_RegEnableReflectionKey      PROC_ADVAPI = 13 | (255 << 16) | (277 << 32)
	PROC_RegEnumKeyExW               PROC_ADVAPI = 14 | (278 << 16) | (291 << 32)
	PROC_RegEnumValueW               PROC_ADVAPI = 15 | (292 << 16) | (305 << 32)
	PROC_RegFlushKey                 PROC_ADVAPI = 16 | (306 << 16) | (317 << 32)
	PROC_RegGetValueW                PROC_ADVAPI = 17 | (318 << 16) | (330 << 32)
	PROC_RegLoadKeyW                 PROC_ADVAPI = 18 | (331 << 16) | (342 << 32)
	PROC_RegOpenKeyExW               PROC_ADVAPI = 19 | (343 << 16) | (356 << 32)
	PROC_RegQueryInfoKeyW            PROC_ADVAPI = 20 | (357 << 16) | (373 << 32)
	PROC_RegQueryMultipleValuesW     PROC_ADVAPI = 21 | (374 << 16) | (397 << 32)
	PROC_RegQueryReflectionKey       PROC_ADVAPI = 22 | (398 << 16) | (419 << 32)
	PROC_RegQueryValueExW            PROC_ADVAPI = 23 | (420 << 16) | (436 << 32)
	PROC_RegRenameKey                PROC_ADVAPI = 24 | (437 << 16) | (449 << 32)
	PROC_RegReplaceKeyW              PROC_ADVAPI = 25 | (450 << 16) | (464 << 32)
	PROC_RegRestoreKeyW              PROC_ADVAPI = 26 | (465 << 16) | (479 << 32)
	PROC_RegSaveKeyW                 PROC_ADVAPI = 27 | (480 << 16) | (491 << 32)
	PROC_RegSaveKeyExW               PROC_ADVAPI = 28 | (492 << 16) | (505 << 32)
	PROC_RegSetKeyValueW             PROC_ADVAPI = 29 | (506 << 16) | (521 << 32)
	PROC_RegSetValueExW              PROC_ADVAPI = 30 | (522 << 16) | (536 << 32)
	PROC_RegUnLoadKeyW               PROC_ADVAPI = 31 | (537 << 16) | (550 << 32)
)

// Declaration of advapi procedure names.
const advapiProcStr = `
--funcs
RegDisablePredefinedCache
RegDisablePredefinedCacheEx

--hkey
RegConnectRegistryW
RegOpenCurrentUser
RegCloseKey
RegCopyTreeW
RegCreateKeyExW
RegDeleteKeyW
RegDeleteKeyExW
RegDeleteKeyValueW
RegDeleteTreeW
RegDeleteValueW
RegDisableReflectionKey
RegEnableReflectionKey
RegEnumKeyExW
RegEnumValueW
RegFlushKey
RegGetValueW
RegLoadKeyW
RegOpenKeyExW
RegQueryInfoKeyW
RegQueryMultipleValuesW
RegQueryReflectionKey
RegQueryValueExW
RegRenameKey
RegReplaceKeyW
RegRestoreKeyW
RegSaveKeyW
RegSaveKeyExW
RegSetKeyValueW
RegSetValueExW
RegUnLoadKeyW
`
