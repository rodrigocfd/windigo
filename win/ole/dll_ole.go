//go:build windows

package ole

import (
	"syscall"

	"github.com/rodrigocfd/windigo/internal/dll"
)

// Stores the loazy-loaded ole procedures.
var oleCache [13]*syscall.Proc

// Loads ole procedures.
func dllOle(procId _PROC_OLE) uintptr {
	return dll.LoadProc(dll.SYSDLL_ole32, oleCache[:], oleProcStr, uint64(procId)).Addr()
}

type _PROC_OLE uint64 // Procedure identifiers for ole.

// Auto-generated ole procedure identifier: cache index | str start | str past-end.
const (
	_PROC_CLSIDFromProgID  _PROC_OLE = 0 | (9 << 16) | (24 << 32)
	_PROC_CoCreateInstance _PROC_OLE = 1 | (25 << 16) | (41 << 32)
	_PROC_CoInitializeEx   _PROC_OLE = 2 | (42 << 16) | (56 << 32)
	_PROC_CoUninitialize   _PROC_OLE = 3 | (57 << 16) | (71 << 32)
	_PROC_CreateBindCtx    _PROC_OLE = 4 | (72 << 16) | (85 << 32)
	_PROC_OleInitialize    _PROC_OLE = 5 | (86 << 16) | (99 << 32)
	_PROC_OleUninitialize  _PROC_OLE = 6 | (100 << 16) | (115 << 32)
	_PROC_RegisterDragDrop _PROC_OLE = 7 | (116 << 16) | (132 << 32)
	_PROC_ReleaseStgMedium _PROC_OLE = 8 | (133 << 16) | (149 << 32)
	_PROC_RevokeDragDrop   _PROC_OLE = 9 | (150 << 16) | (164 << 32)
	_PROC_CoTaskMemAlloc   _PROC_OLE = 10 | (177 << 16) | (191 << 32)
	_PROC_CoTaskMemFree    _PROC_OLE = 11 | (192 << 16) | (205 << 32)
	_PROC_CoTaskMemRealloc _PROC_OLE = 12 | (206 << 16) | (222 << 32)
)

// Declaration of ole procedure names.
const oleProcStr = `
--funcs
CLSIDFromProgID
CoCreateInstance
CoInitializeEx
CoUninitialize
CreateBindCtx
OleInitialize
OleUninitialize
RegisterDragDrop
ReleaseStgMedium
RevokeDragDrop

--htaskmem
CoTaskMemAlloc
CoTaskMemFree
CoTaskMemRealloc
`
