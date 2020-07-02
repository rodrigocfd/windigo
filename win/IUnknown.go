/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type IUnknown struct {
	lpVtbl uintptr
}

type iUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

func (v *IUnknown) Release() uint32 {
	lpVtbl := (*iUnknownVtbl)(unsafe.Pointer(v.lpVtbl))
	ret, _, _ := syscall.Syscall(lpVtbl.Release, 1,
		uintptr(unsafe.Pointer(v)), 0, 0)
	return uint32(ret)
}

// Creates any COM interface, returning the base IUnknown.
// To retrieve the other interface itself, cast the inner lpVtbl.
func coCreateInstance(clsid *co.GUID, iid *co.GUID) *IUnknown {
	if iid == nil {
		iid = &co.Guid_IUnknown
	}

	// Returns a new GUID with the last uint64 member bytes flipped.
	// This is better than having a makeGuid() function being called to
	// initialize all the GUIDs with correct by order, or even having to flip
	// them manually on each global declaration.
	flipLastUint64 := func(guid *co.GUID) co.GUID {
		buf64 := [8]byte{}
		binary.BigEndian.PutUint64(buf64[:], guid.Data4)
		guidCopy := *guid
		guidCopy.Data4 = binary.LittleEndian.Uint64(buf64[:])
		return guidCopy
	}

	clsidFlip := flipLastUint64(clsid)
	iidFlip := flipLastUint64(iid)
	retIUnk := &IUnknown{}

	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(&clsidFlip)), 0,
		uintptr(co.CLSCTX_INPROC_SERVER),
		uintptr(unsafe.Pointer(&iidFlip)), uintptr(unsafe.Pointer(&retIUnk)), 0)

	if co.ERROR(ret) != co.ERROR_S_OK {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("CoCreateInstance failed: %d %s",
			lerr, lerr.Error()))
	}
	return retIUnk
}
