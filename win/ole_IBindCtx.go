//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [IBindCtx] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	bindCtx, _ := win.CreateBindCtx(rel)
//
// [IBindCtx]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nn-objidl-ibindctx
type IBindCtx struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IBindCtx) IID() co.IID {
	return co.IID_IBindCtx
}

// [EnumObjectParam] method.
//
// [EnumObjectParam]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-enumobjectparam
func (me *IBindCtx) EnumObjectParam(releaser *OleReleaser) (*IEnumString, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).EnumObjectParam,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IEnumString{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// [GetBindOptions] method.
//
// [GetBindOptions]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-getbindoptions
func (me *IBindCtx) GetBindOptions() (BIND_OPTS3, error) {
	var bo BIND_OPTS3
	bo.SetCbStruct()

	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).GetBindOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&bo)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return bo, nil
	} else {
		return BIND_OPTS3{}, hr
	}
}

// [GetObjectParam] method.
//
// [GetObjectParam]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-getobjectparam
func (me *IBindCtx) GetObjectParam(releaser *OleReleaser, key string, ppOut interface{}) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	pKey := wbuf.PtrAllowEmpty(key)

	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).GetObjectParam,
		uintptr(pKey),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.OleCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(OleObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

// [RegisterObjectBound] method.
//
// [RegisterObjectBound]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-registerobjectbound
func (me *IBindCtx) RegisterObjectBound(obj *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).RegisterObjectBound,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(obj.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [ReleaseBoundObjects] method.
//
// [ReleaseBoundObjects]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-releaseboundobjects
func (me *IBindCtx) ReleaseBoundObjects() error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).ReleaseBoundObjects,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [RevokeObjectBound] method.
//
// [RevokeObjectBound]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-revokeobjectbound
func (me *IBindCtx) RevokeObjectBound(obj *IUnknown) error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).RevokeObjectBound,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(obj.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [SetBindOptions] method.
//
// [SetBindOptions]: https://learn.microsoft.com/en-us/windows/win32/api/objidl/nf-objidl-ibindctx-setbindoptions
func (me *IBindCtx) SetBindOptions(bindOpts *BIND_OPTS3) error {
	ret, _, _ := syscall.SyscallN(
		(*_IBindCtxVt)(unsafe.Pointer(*me.Ppvt())).SetBindOptions,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(bindOpts)))
	return utl.ErrorAsHResult(ret)
}

type _IBindCtxVt struct {
	_IUnknownVt
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
