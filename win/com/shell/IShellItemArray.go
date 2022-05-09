//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray interface {
	com.IUnknown

	// Prefer using IShellItemArray.ListDisplayNames(), which directly retrieves
	// all paths at once.
	//
	// ⚠️ You must defer IShellItem.Release() on the returned object.
	//
	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
	GetItemAt(index int) IShellItem

	// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
	GetCount() int

	// This helper method calls IShellItemArray.GetItemAt() to retrieve all
	// IShellItem objects, then calls IShellItem.GetDisplayName() on each one,
	// retrieving the names as strings.
	//
	// Example:
	//
	//		var shia shell.IShellItemArray // initialized somewhere
	//
	//		fullPaths := shia.ListDisplayNames(shellco.SIGDN_FILESYSPATH)
	ListDisplayNames(sigdnName shellco.SIGDN) []string
}

type _IShellItemArray struct{ com.IUnknown }

// Constructs a COM object from the base IUnknown.
//
// ⚠️ You must defer IShellItemArray.Release().
func NewIShellItemArray(base com.IUnknown) IShellItemArray {
	return &_IShellItemArray{IUnknown: base}
}

func (me *_IShellItemArray) GetItemAt(index int) IShellItem {
	var ppvQueried **comvt.IUnknown
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellItemArray)(unsafe.Pointer(*me.Ptr())).GetItemAt, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(index), uintptr(unsafe.Pointer(&ppvQueried)))

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return NewIShellItem(com.NewIUnknown(ppvQueried))
	} else {
		panic(hr)
	}
}

func (me *_IShellItemArray) GetCount() int {
	var count uint32
	ret, _, _ := syscall.Syscall(
		(*shellvt.IShellItemArray)(unsafe.Pointer(*me.Ptr())).GetCount, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(unsafe.Pointer(&count)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(count)
	} else {
		panic(hr)
	}
}

func (me *_IShellItemArray) ListDisplayNames(sigdnName shellco.SIGDN) []string {
	count := me.GetCount()
	names := make([]string, 0, count)

	for i := 0; i < count; i++ {
		shellItem := me.GetItemAt(i)
		defer shellItem.Release() // will pile up at the end of the function, but it's fine
		names = append(names, shellItem.GetDisplayName(sigdnName))
	}

	return names
}
