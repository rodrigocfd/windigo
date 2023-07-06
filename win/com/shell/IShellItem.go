//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [IShellItem] COM interface.
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem interface {
	com.IUnknown

	// [Compare] COM method.
	//
	// [Compare]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-compare
	Compare(si IShellItem, hint shellco.SICHINT) bool

	// [GetAttributes] COM method.
	//
	// [GetAttributes]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getattributes
	GetAttributes(mask co.SFGAO) co.SFGAO

	// [GetParent] COM method.
	//
	// ⚠️ You must defer IShellItem.Release() on the returned object.
	//
	// [GetParent]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
	GetParent() IShellItem

	// [GetDisplayName] COM method.
	//
	// # Example
	//
	//	var shi shell.IShellItem // initialized somewhere
	//
	//	fullPath := shi.GetDisplayName(shellco.SIGDN_FILESYSPATH)
	//
	// [GetDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
	GetDisplayName(sigdnName shellco.SIGDN) string
}

type _IShellItem struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IShellItem.Release().
func NewIShellItem(base com.IUnknown) IShellItem {
	return &_IShellItem{IUnknown: base}
}

// [SHCreateItemFromParsingName]: function.
//
// ⚠️ You must defer IShellItem.Release().
//
// # Example
//
//	ish := shell.NewShellItemFromPath("C:\\Temp\\file.txt")
//	defer ish.Release()
//
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func SHCreateItemFromParsingName(folderOrFilePath string) (IShellItem, error) {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(proc.SHCreateItemFromParsingName.Addr(),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(folderOrFilePath))),
		0, uintptr(unsafe.Pointer(win.GuidFromIid(shellco.IID_IShellItem))),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(com.NewIUnknown(ppvQueried)), nil
	} else {
		return nil, hr
	}
}

func (me *_IShellItem) Compare(si IShellItem, hint shellco.SICHINT) bool {
	var piOrder uint32
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).Compare,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())),
		uintptr(hint),
		uintptr(unsafe.Pointer(&piOrder)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

func (me *_IShellItem) GetAttributes(mask co.SFGAO) co.SFGAO {
	var attribs co.SFGAO
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).GetAttributes,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&mask)),
		uintptr(unsafe.Pointer(&attribs)))

	if hr := errco.ERROR(ret); hr == errco.S_OK || hr == errco.S_FALSE {
		return attribs
	} else {
		panic(hr)
	}
}

func (me *_IShellItem) GetParent() IShellItem {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).GetParent,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IShellItem) GetDisplayName(sigdnName shellco.SIGDN) string {
	var pv uintptr
	ret, _, _ := syscall.SyscallN(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).GetDisplayName,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(sigdnName), uintptr(unsafe.Pointer(&pv)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		defer win.HTASKMEM(pv).CoTaskMemFree()
		name := win.Str.FromNativePtr((*uint16)(unsafe.Pointer(pv)))
		return name
	} else {
		panic(hr)
	}
}
