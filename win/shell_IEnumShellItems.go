//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IEnumShellItems] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var desktop *win.IShellItem
//	_ = win.SHGetKnownFolderItem(
//		rel,
//		co.FOLDERID_Desktop,
//		co.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	var enumItems *win.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, co.BHID_EnumItems, &enumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(path)
//	}
//
// [IEnumShellItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumshellitems
type IEnumShellItems struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumShellItems) IID() co.IID {
	return co.IID_IEnumShellItems
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-clone
func (me *IEnumShellItems) Clone(releaser *OleReleaser) (*IEnumShellItems, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IEnumShellItems{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// Returns all [IShellItem] values by calling [IEnumShellItems.Next].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var desktop *win.IShellItem
//	_ = win.SHGetKnownFolderItem(
//		rel,
//		co.FOLDERID_Desktop,
//		co.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	var enumItems *win.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, co.BHID_EnumItems, &enumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(path)
//	}
func (me *IEnumShellItems) Enum(releaser *OleReleaser) ([]*IShellItem, error) {
	items := make([]*IShellItem, 0)
	var item *IShellItem
	var hr error

	for {
		item, hr = me.Next(releaser)
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
func (me *IEnumShellItems) Next(releaser *OleReleaser) (*IShellItem, error) {
	var ppvtQueried **_IUnknownVt
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
		1,
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItem{IUnknown{ppvtQueried}}
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
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-skip
func (me *IEnumShellItems) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		(*_IEnumShellItemsVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.ErrorAsHResult(ret)
}

type _IEnumShellItemsVt struct {
	_IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}
