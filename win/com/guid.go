/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

import (
	"encoding/binary"
	"fmt"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win/proc"
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]uint8
}

func makeGuid(d1 uint32, d2, d3 uint16, d4 uint64) GUID {
	guid := GUID{Data1: d1, Data2: d2, Data3: d3}
	binary.BigEndian.PutUint64(guid.Data4[:], d4)
	return guid
}

func (clsid *GUID) CoCreateInstance(iid *GUID) *IUnknown {
	if iid == nil {
		iid = &Guid_IUnknown
	}
	unk := &IUnknown{}
	ret, _, _ := syscall.Syscall6(proc.CoCreateInstance.Addr(), 5,
		uintptr(unsafe.Pointer(clsid)), 0, uintptr(co.CLSCTX_INPROC_SERVER),
		uintptr(unsafe.Pointer(iid)), uintptr(unsafe.Pointer(&unk)), 0)

	if co.ERROR(ret) != co.ERROR_S_OK {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("CoCreateInstance failed: %d %s",
			lerr, lerr.Error()))
	}
	return unk
}

func CoInitializeEx(dwCoInit co.COINIT) {
	ret, _, _ := syscall.Syscall(proc.CoInitializeEx.Addr(), 2,
		0, uintptr(dwCoInit), 0)
	if co.ERROR(ret) != co.ERROR_S_OK && co.ERROR(ret) != co.ERROR_S_FALSE {
		lerr := syscall.Errno(ret)
		panic(fmt.Sprintf("CoInitializeEx failed: %d %s",
			lerr, lerr.Error()))
	}
}

func CoUninitialize() {
	syscall.Syscall(proc.CoUninitialize.Addr(), 0, 0, 0, 0)
}
