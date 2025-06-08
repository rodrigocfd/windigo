//go:build windows

package oleaut

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// Stores the loazy-loaded oleaut procedures.
var oleautCache [9]*syscall.Proc

// Loads oleaut procedures.
func dllOleaut(procId _PROC_OLEAUT) uintptr {
	return dll.LoadProc(dll.SYSDLL_oleaut32, oleautCache[:], oleautProcStr, uint64(procId)).Addr()
}

type _PROC_OLEAUT uint64 // Procedure identifiers for oleaut.

// Auto-generated oleaut procedure identifier: cache index | str start | str past-end.
const (
	_PROC_SysAllocString          _PROC_OLEAUT = 0 | (8 << 16) | (22 << 32)
	_PROC_SysFreeString           _PROC_OLEAUT = 1 | (23 << 16) | (36 << 32)
	_PROC_SysReAllocString        _PROC_OLEAUT = 2 | (37 << 16) | (53 << 32)
	_PROC_OleLoadPicture          _PROC_OLEAUT = 3 | (63 << 16) | (77 << 32)
	_PROC_OleLoadPicturePath      _PROC_OLEAUT = 4 | (78 << 16) | (96 << 32)
	_PROC_VariantClear            _PROC_OLEAUT = 5 | (108 << 16) | (120 << 32)
	_PROC_VariantInit             _PROC_OLEAUT = 6 | (121 << 16) | (132 << 32)
	_PROC_SystemTimeToVariantTime _PROC_OLEAUT = 7 | (133 << 16) | (156 << 32)
	_PROC_VariantTimeToSystemTime _PROC_OLEAUT = 8 | (157 << 16) | (180 << 32)
)

// Declaration of oleaut procedure names.
const oleautProcStr = `
--bstr
SysAllocString
SysFreeString
SysReAllocString

--funcs
OleLoadPicture
OleLoadPicturePath

--variant
VariantClear
VariantInit
SystemTimeToVariantTime
VariantTimeToSystemTime
`
