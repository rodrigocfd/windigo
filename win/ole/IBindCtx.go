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

// Returns the unique COM interface identifier.
func (*IBindCtx) IID() co.IID {
	return co.IID_IBindCtx
}
