package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IShellItemVtbl struct {
	win.IUnknownVtbl
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem struct {
	win.IUnknown // Base IUnknown.
}

// Creates an IShellItem from a string path.
//
// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func NewShellItem(thePath string) (IShellItem, error) {
	var ppv **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall6(proc.SHCreateItemFromParsingName.Addr(), 4,
		uintptr(unsafe.Pointer(win.Str.ToUint16Ptr(thePath))),
		0, uintptr(unsafe.Pointer(win.NewGuidFromIid(shellco.IID_IShellItem))),
		uintptr(unsafe.Pointer(&ppv)),
		0, 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		return IShellItem{}, lerr
	}
	return IShellItem{
		win.IUnknown{Ppv: ppv},
	}, nil
}

func (me *IShellItem) Compare(psi IShellItem, hint shellco.SICHINT) bool {
	piOrder := uint32(0)
	ret, _, _ := syscall.Syscall6(
		(*_IShellItemVtbl)(unsafe.Pointer(*me.Ppv)).Compare, 4,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(psi.Ppv)),
		uintptr(hint),
		uintptr(unsafe.Pointer(&piOrder)), 0, 0)

	if lerr := err.ERROR(ret); lerr == err.S_OK {
		return true
	} else if lerr == err.S_FALSE {
		return false
	} else {
		panic(lerr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getattributes
func (me *IShellItem) GetAttributes(sfgaoMask co.SFGAO) co.SFGAO {
	attribs := co.SFGAO(0)
	ret, _, _ := syscall.Syscall(
		(*_IShellItemVtbl)(unsafe.Pointer(*me.Ppv)).GetAttributes, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&sfgaoMask)),
		uintptr(unsafe.Pointer(&attribs)))

	if lerr := err.ERROR(ret); lerr != err.S_OK && lerr != err.S_FALSE {
		panic(lerr)
	}
	return attribs
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
func (me *IShellItem) GetParent() IShellItem {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IShellItemVtbl)(unsafe.Pointer(*me.Ppv)).GetParent, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return IShellItem{
		win.IUnknown{Ppv: ppvQueried},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName shellco.SIGDN) string {
	var pv *uint16
	ret, _, _ := syscall.Syscall(
		(*_IShellItemVtbl)(unsafe.Pointer(*me.Ppv)).GetDisplayName, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(sigdnName), uintptr(unsafe.Pointer(&pv)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	name := win.Str.FromUint16Ptr(pv)
	win.CoTaskMemFree(unsafe.Pointer(pv))
	return name
}
