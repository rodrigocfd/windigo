//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IShellItemArray] COM interface.
//
// [IShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItemArray) IID() co.IID {
	return co.IID_IShellItemArray
}

// Returns the path names of each [IShellItem] object by calling [GetCount],
// [GetItemAt] and [GetDisplayName].
//
// # Example
//
//	var arr shell.IShellItemArray // initialized somewhere
//
//	names, _ := arr.EnumDisplayNames(co.SIGDN_FILESYSPATH)
//	for _, fullPath := range names {
//		println(fullPath)
//	}
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
// [GetItemAt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
// [GetDisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitem-getdisplayname
func (me *IShellItemArray) EnumDisplayNames(sigdnName co.SIGDN) ([]string, error) {
	localRel := ole.NewReleaser()
	defer localRel.Release()

	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, count)

	for i := uint(0); i < count; i++ {
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

// Returns all [IShellItem] objects by calling [GetCount] and [GetItemAt].
//
// If you just want to retrieve the paths, prefer using EnumDisplayNames().
//
// # Example
//
//	var arr shell.IShellItemArray // initialized somewhere
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	items, _ := arr.EnumItems(rel)
//	for _, item := range items {
//		fullPath, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(fullPath)
//	}
//
// [IShellItem]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem
// [GetCount]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getcount
// [GetItemAt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) EnumItems(releaser *ole.Releaser) ([]*IShellItem, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	items := make([]*IShellItem, 0, count)

	for i := uint(0); i < count; i++ {
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
func (me *IShellItemArray) GetCount() (uint, error) {
	var count uint32
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItemArray)(unsafe.Pointer(*me.Ppvt())).GetCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&count)), 0)

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return uint(count), nil
	} else {
		return 0, hr
	}
}

// [GetItemAt] method.
//
// [GetItemAt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) GetItemAt(releaser *ole.Releaser, index uint) (*IShellItem, error) {
	var ppvtQueried **vt.IUnknown
	ret, _, _ := syscall.SyscallN(
		(*vt.IShellItemArray)(unsafe.Pointer(*me.Ppvt())).GetItemAt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(index), uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := vt.NewObj[IShellItem](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}
