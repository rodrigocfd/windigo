//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-idroptarget
type IDropTarget interface {
	com.IUnknown

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragenter
	DragEnter(dataObj IDataObject, keyState co.MK,
		pt win.POINT, effect *shellco.DROPEFFECT)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragleave
	DragLeave()

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-dragover
	DragOver(keyState co.MK, pt win.POINT, effect *shellco.DROPEFFECT)

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/oleidl/nf-oleidl-idroptarget-drop
	Drop(dataObj IDataObject, keyState co.MK,
		pt win.POINT, effect *shellco.DROPEFFECT)
}

type _IDropTarget struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IDropTarget.Release().
func NewIDropTarget(base com.IUnknown) IDropTarget {
	return &_IDropTarget{IUnknown: base}
}

func (me *_IDropTarget) DragEnter(
	dataObj IDataObject, keyState co.MK,
	pt win.POINT, effect *shellco.DROPEFFECT) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDropTarget)(unsafe.Pointer(*me.Ptr())).DragEnter,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(dataObj.Ptr())),
		uintptr(keyState), uintptr(pt.X), uintptr(pt.Y),
		uintptr(unsafe.Pointer(effect)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDropTarget) DragLeave() {
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDropTarget)(unsafe.Pointer(*me.Ptr())).DragLeave,
		uintptr(unsafe.Pointer(me.Ptr())))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDropTarget) DragOver(
	keyState co.MK, pt win.POINT, effect *shellco.DROPEFFECT) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDropTarget)(unsafe.Pointer(*me.Ptr())).DragOver,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(keyState), uintptr(pt.X), uintptr(pt.Y),
		uintptr(unsafe.Pointer(effect)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

func (me *_IDropTarget) Drop(
	dataObj IDataObject, keyState co.MK,
	pt win.POINT, effect *shellco.DROPEFFECT) {

	ret, _, _ := syscall.SyscallN(
		(*shellvt.IDropTarget)(unsafe.Pointer(*me.Ptr())).Drop,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(dataObj.Ptr())),
		uintptr(keyState), uintptr(pt.X), uintptr(pt.Y),
		uintptr(unsafe.Pointer(effect)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
