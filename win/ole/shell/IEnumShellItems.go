//go:build windows

package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IEnumShellItems] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	desktop, _ := shell.SHGetKnownFolderItem[shell.IShellItem](
//		rel, co.FOLDERID_Desktop, co.KF_FLAG_DEFAULT, win.HANDLE(0))
//
//	enumItems, _ := shell.BindToHandler[shell.IEnumShellItems](
//		desktop, rel, nil, co.BHID_EnumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(path)
//	}
//
// [IEnumShellItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumshellitems
type IEnumShellItems struct{ ole.IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumShellItems) IID() co.IID {
	return co.IID_IEnumShellItems
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-clone
func (me *IEnumShellItems) Clone(releaser *ole.Releaser) (*IEnumShellItems, error) {
	var ppvtQueried **ole.IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := ole.ComObj[IEnumShellItems](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// Returns all objects by calling [IEnumShellItems.Next].
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	desktop, _ := shell.SHGetKnownFolderItem[shell.IShellItem](
//		rel, co.FOLDERID_Desktop, co.KF_FLAG_DEFAULT, win.HANDLE(0))
//
//	enumItems, _ := shell.BindToHandler[shell.IEnumShellItems](
//		desktop, rel, nil, co.BHID_EnumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(path)
//	}
func (me *IEnumShellItems) Enum(rel *ole.Releaser) ([]*IShellItem, error) {
	items := make([]*IShellItem, 0)
	var item *IShellItem
	var hr error

	for {
		item, hr = me.Next(rel)
		if hr != nil { // actual error
			return nil, hr
		} else if item == nil { // no more items to fetch
			return items, nil
		} else { // item fetched
			items = append(items, item)
		}
	}
}

// [Next] method.
//
// If there are no more items, nil is returned.
//
// This is a low-level method, prefer using [IEnumShellItems.Enum].
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-next
func (me *IEnumShellItems) Next(releaser *ole.Releaser) (*IShellItem, error) {
	var ppvtQueried **ole.IUnknownVt
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
		1, uintptr(unsafe.Pointer(&ppvtQueried)), uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := ole.ComObj[IShellItem](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else if hr == co.HRESULT_S_FALSE {
		return nil, nil
	} else {
		return nil, hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-reset
func (me *IEnumShellItems) Reset() error {
	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Reset,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [Skip] method.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-skip
func (me *IEnumShellItems) Skip(count uint) error {
	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.ErrorAsHResult(ret)
}

type _IEnumShellItemsVt struct {
	ole.IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}
