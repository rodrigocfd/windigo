//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/co"
)

// [ITypeLib] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [ITypeLib]: https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-itypelib
type ITypeLib struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITypeLib) IID() co.IID {
	return co.IID_ITypeLib
}

type _ITypeLibVt struct {
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
