//go:build windows

package winsh

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cosh"
)

// [IShellItemArray] COM interface.
//
// [IShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray struct{ win.IUnknown }

type _IShellItemArrayVt struct {
	utl.IUnknownVt
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr
	GetItemAt                  uintptr
	EnumItems                  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItemArray) IID() *co.IID {
	return &cosh.IID_IShellItemArray
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IShellItemArray) AddRef(releaser *win.OleReleaser) *IShellItemArray {
	return utl.OleNewFromAddRef[*IShellItemArray](me, releaser)
}

// Returns the path names of each [IShellItem] object by calling
// [IShellItemArray.GetCount], [IShellItemArray.GetItemAt] and
// [IShellItem.GetDisplayName].
//
// Example:
//
//	var arr *winsh.IShellItemArray // initialized somewhere
//
//	names, _ := arr.EnumDisplayNames(cosh.SIGDN_FILESYSPATH)
//	for _, fullPath := range names {
//		println(fullPath)
//	}
func (me *IShellItemArray) EnumDisplayNames(sigdnName cosh.SIGDN) ([]string, error) {
	localRel := win.NewOleReleaser()
	defer localRel.Release()

	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, count)
	for i := 0; i < count; i++ {
		shellItem, err := me.GetItemAt(localRel, i)
		if err != nil {
			return nil, err
		}

		name, err := shellItem.GetDisplayName(sigdnName)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

// Returns all [IShellItem] objects by calling [IShellItemArray.GetCount] and
// [IShellItemArray.GetItemAt].
//
// If you just want to retrieve the paths, prefer using
// [IShellItemArray.EnumDisplayNames].
//
// Example:
//
//	var arr *winsh.IShellItemArray // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	items, _ := arr.EnumItems(rel)
//	for _, item := range items {
//		fullPath, _ := item.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//		println(fullPath)
//	}
func (me *IShellItemArray) EnumItems(releaser *win.OleReleaser) ([]*IShellItem, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	items := make([]*IShellItem, 0, count)
	for i := 0; i < count; i++ {
		shellItem, err := me.GetItemAt(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		items = append(items, shellItem)
	}
	return items, nil
}

// [GetCount] method.
//
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
func (me *IShellItemArray) GetCount() (int, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItemArrayVt](me.Ppvt()).GetCount,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&count)),
		0)
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(count), nil
}

// [GetItemAt] method.
//
// Panics if index is negative.
//
// [GetItemAt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) GetItemAt(releaser *win.OleReleaser, index int) (*IShellItem, error) {
	utl.PanicNeg(index)
	var ppvtQueried uintptr

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IShellItemArrayVt](me.Ppvt()).GetItemAt,
		me.Ppvt(),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IShellItem](ret, ppvtQueried, releaser)
}
