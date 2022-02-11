package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
type IShellItem struct{ win.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IShellItem.Release().
func NewIShellItem(base win.IUnknown) IShellItem {
	return IShellItem{IUnknown: base}
}

// Creates an IShellItem from a string path.
//
// ⚠️ You must defer IShellItem.Release().
//
// Example:
//
//  ish := shell.NewShellItemFromPath("C:\\Temp\\file.txt")
//  defer ish.Release()
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
func NewShellItemFromPath(folderOrFilePath string) (IShellItem, error) {
	var ppvQueried win.IUnknown
	ret, _, _ := syscall.Syscall6(proc.SHCreateItemFromParsingName.Addr(), 4,
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(folderOrFilePath))),
		0, uintptr(unsafe.Pointer(win.GuidFromIid(shellco.IID_IShellItem))),
		uintptr(unsafe.Pointer(&ppvQueried)),
		0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(ppvQueried), nil
	} else {
		return IShellItem{}, hr
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-compare
func (me *IShellItem) Compare(si IShellItem, hint shellco.SICHINT) bool {
	var piOrder uint32
	ret, _, _ := syscall.Syscall6(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).Compare, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(si.Ptr())),
		uintptr(hint),
		uintptr(unsafe.Pointer(&piOrder)), 0, 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return true
	} else if hr == errco.S_FALSE {
		return false
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getattributes
func (me *IShellItem) GetAttributes(mask co.SFGAO) co.SFGAO {
	var attribs co.SFGAO
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).GetAttributes, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&mask)),
		uintptr(unsafe.Pointer(&attribs)))

	if hr := errco.ERROR(ret); hr == errco.S_OK || hr == errco.S_FALSE {
		return attribs
	} else {
		panic(hr)
	}
}

// ⚠️ You must defer IShellItem.Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getparent
func (me *IShellItem) GetParent() IShellItem {
	var ppvQueried win.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).GetParent, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&ppvQueried)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(ppvQueried)
	} else {
		panic(hr)
	}
}

// Example:
//
//  var shi shell.IShellItem // initialized somewhere
//
//  fullPath := shi.GetDisplayName(shellco.SIGDN_FILESYSPATH)
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItem) GetDisplayName(sigdnName shellco.SIGDN) string {
	var pv uintptr
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellItem)(unsafe.Pointer(*me.Ptr())).GetDisplayName, 3,
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
