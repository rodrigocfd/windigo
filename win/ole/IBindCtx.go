//go:build windows

package ole

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// [IBindCtx] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	bindCtx, _ := ole.CreateBindCtx(rel)
//
// [IBindCtx]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ibindctx
type IBindCtx struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IBindCtx) IID() co.IID {
	return co.IID_IBindCtx
}

type _IBindCtxVt struct {
	IUnknownVt
	RegisterObjectBound   uintptr
	RevokeObjectBound     uintptr
	ReleaseBoundObjects   uintptr
	SetBindOptions        uintptr
	GetBindOptions        uintptr
	GetRunningObjectTable uintptr
	RegisterObjectParam   uintptr
	GetObjectParam        uintptr
	EnumObjectParam       uintptr
	RevokeObjectParam     uintptr
}
