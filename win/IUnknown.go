/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package win

import (
	"encoding/binary"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	proc "github.com/rodrigocfd/windigo/win/internal"
)

// Returns a GUID struct from hex numbers, which can be copied straight from
// standard GUID definitions.
//
// Example for IUnknown:
// g := NewGuid(0x00000000, 0x0000, 0x0000, 0xc000, 0x000000000046)
func NewGuid(p1 uint32, p2, p3, p4 uint16, p5 uint64) *GUID {
	newGuid := GUID{
		Data1: p1,
		Data2: p2,
		Data3: p3,
		Data4: (uint64(p4) << 48) | p5,
	}

	buf64 := [8]byte{}
	binary.BigEndian.PutUint64(buf64[:], newGuid.Data4)
	newGuid.Data4 = binary.LittleEndian.Uint64(buf64[:]) // reverse bytes of Data4
	return &newGuid
}

//------------------------------------------------------------------------------

type (
	// Base to all COM interfaces.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	IUnknown struct {
		Ppv **IUnknownVtbl // Pointer to pointer to the COM virtual table.
	}

	// IUnknown virtual table.
	IUnknownVtbl struct {
		QueryInterface uintptr
		AddRef         uintptr
		Release        uintptr
	}
)

// Returns an IUnknown COM object. The inner Ppv can be cast to any COM interface.
// Typically uses CLSCTX_INPROC_SERVER.
//
// You must defer Release().
//
// https://docs.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(
	rclsid *GUID, pUnkOuter unsafe.Pointer,
	dwClsContext co.CLSCTX, riid *GUID) (*IUnknown, error) {

	var ppv **IUnknownVtbl
	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(rclsid)), uintptr(pUnkOuter),
		uintptr(dwClsContext), uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppv)), 0)

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		return nil, NewWinError(lerr, "CoCreateInstance")
	}
	return &IUnknown{Ppv: ppv}, nil
}

// Queries an IUnknown COM object. The inner Ppv can be cast to any COM interface.
//
// You must defer Release() on the returned COM object.
//
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (me *IUnknown) QueryInterface(riid *GUID) (*IUnknown, error) {
	var ppvQueried **IUnknownVtbl
	ret, _, _ := syscall.Syscall((*me.Ppv).QueryInterface, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if lerr := co.ERROR(ret); lerr != co.ERROR_S_OK {
		return nil, NewWinError(lerr, "IUnknown.QueryInterface")
	}
	return &IUnknown{Ppv: ppvQueried}, nil
}

// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IUnknown) AddRef() uint32 {
	ret, _, _ := syscall.Syscall((*me.Ppv).AddRef, 1,
		uintptr(unsafe.Pointer(me.Ppv)), 0, 0)
	return uint32(ret)
}

// Can be called any number of times, will actually release only while the
// internal ref count is greater than zero.
//
// https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (me *IUnknown) Release() uint32 {
	if me.Ppv != nil {
		ret, _, _ := syscall.Syscall((*me.Ppv).Release, 1,
			uintptr(unsafe.Pointer(me.Ppv)), 0, 0)
		if ret == 0 { // COM pointer was released
			me.Ppv = nil
		}
		return uint32(ret)
	}
	return 0
}
