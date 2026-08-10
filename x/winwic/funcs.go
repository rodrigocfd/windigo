//go:build windows

package winwic

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/wstr"
)

// Calls the COM method which allocates a buffer and returns a string.
func oleCallAllocBufRetStr(me interface{ Ppvt() uintptr }, pMethod uintptr) (string, error) {
	var szBuf uint32
	ret, _, _ := syscall.SyscallN(
		pMethod,
		me.Ppvt(),
		0, 0,
		uintptr(unsafe.Pointer(&szBuf)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}

	buf := make([]uint16, szBuf)
	ret, _, _ = syscall.SyscallN(
		pMethod,
		me.Ppvt(),
		uintptr(szBuf),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&szBuf)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return "", hr
	}
	return wstr.DecodeSlice(buf), nil
}
