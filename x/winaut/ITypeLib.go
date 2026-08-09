//go:build windows

package winaut

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/coaut"
)

// [ITypeLib] COM interface.
//
// [ITypeLib]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypelib
type ITypeLib struct{ win.IUnknown }

type _ITypeLibVt struct {
	utl.IUnknownVt
	GetTypeInfoCount  uintptr
	GetTypeInfo       uintptr
	GetTypeInfoType   uintptr
	GetTypeInfoOfGuid uintptr
	GetLibAttr        uintptr
	GetTypeComp       uintptr
	GetDocumentation  uintptr
	IsName            uintptr
	FindName          uintptr
	ReleaseTLibAttr   uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITypeLib) IID() *co.IID {
	return &coaut.IID_ITypeLib
}

// [GetTypeInfo] method.
//
// [GetTypeInfo]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypelib-gettypeinfo
func (me *ITypeLib) GetTypeInfo(releaser *win.OleReleaser, index int) (*ITypeInfo, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeLibVt](me.Ppvt()).GetTypeInfo,
		me.Ppvt(),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITypeInfo](ret, ppvtQueried, releaser)
}

// [GetTypeInfoCount] method.
//
// [GetTypeInfoCount]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypelib-gettypeinfocount
func (me *ITypeLib) GetTypeInfoCount() int {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeLibVt](me.Ppvt()).GetTypeInfoCount,
		me.Ppvt())
	return int(ret)
}

// [GetTypeInfoOfGuid] method.
//
// [GetTypeInfoOfGuid]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-itypelib-gettypeinfoofguid
func (me *ITypeLib) GetTypeInfoOfGuid(releaser *win.OleReleaser, pGuid *co.GUID) (*ITypeInfo, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITypeLibVt](me.Ppvt()).GetTypeInfo,
		me.Ppvt(),
		uintptr(unsafe.Pointer(pGuid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITypeInfo](ret, ppvtQueried, releaser)
}
