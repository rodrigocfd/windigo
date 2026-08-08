//go:build windows

package wintasks

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// Calls the COM method without parameters, returns BSTR.
func oleCallRetBstr(me interface{ Ppvt() uintptr }, pMethod uintptr) (string, error) {
	var name winaut.BSTR
	defer name.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		pMethod,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&name)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return name.String(), nil
	} else {
		return "", hr
	}
}

// Calls the COM method to set a BSTR.
func oleCallSetBstr(me interface{ Ppvt() uintptr }, s string, pMethod uintptr) error {
	bstrS, err := winaut.SysAllocString(s)
	if err != nil {
		return err
	}
	defer bstrS.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		pMethod,
		me.Ppvt(),
		uintptr(bstrS))
	return utl.HresultToError(ret)
}
