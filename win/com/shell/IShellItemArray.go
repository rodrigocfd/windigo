package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

type _IShellItemArrayVtbl struct {
	win.IUnknownVtbl
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr
	GetItemAt                  uintptr
	EnumItems                  uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray struct {
	win.IUnknown // Base IUnknown.
}

// ⚠️ You must defer Release().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) GetItemAt(index int) IShellItem {
	var ppvQueried **win.IUnknownVtbl
	ret, _, _ := syscall.Syscall(
		(*_IShellItemArrayVtbl)(unsafe.Pointer(*me.Ppv)).GetItemAt, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(index), uintptr(unsafe.Pointer(&ppvQueried)))

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return IShellItem{
		win.IUnknown{Ppv: ppvQueried},
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
func (me *IShellItemArray) GetCount() int {
	count := uint32(0)
	ret, _, _ := syscall.Syscall(
		(*_IShellItemArrayVtbl)(unsafe.Pointer(*me.Ppv)).GetCount, 3,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&count)), 0)

	if lerr := err.ERROR(ret); lerr != err.S_OK {
		panic(lerr)
	}
	return int(count)
}

// Syntactic sugar, calls GetDisplayName() on each IShellItem.
func (me *IShellItemArray) GetDisplayNames() []string {
	count := me.GetCount()
	files := make([]string, 0, count)

	for i := 0; i < count; i++ {
		shellItem := me.GetItemAt(i)
		files = append(files, shellItem.GetDisplayName(co.SIGDN_FILESYSPATH))
		shellItem.Release()
	}

	return files
}
