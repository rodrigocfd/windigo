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

// [IEnumShellItems] COM interface.
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
//	var desktop *winsh.IShellItem
//	_ = winsh.SHGetKnownFolderItem(
//		rel,
//		&cosh.FOLDERID_Desktop,
//		cosh.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	var enumItems *winsh.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, &cosh.BHID_EnumItems, &enumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//		println(path)
//	}
//
// [IEnumShellItems]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumshellitems
type IEnumShellItems struct{ win.IUnknown }

type _IEnumShellItemsVt struct {
	utl.IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumShellItems) IID() *co.IID {
	return &cosh.IID_IEnumShellItems
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-clone
func (me *IEnumShellItems) Clone(releaser *win.OleReleaser) (*IEnumShellItems, error) {
	return utl.OleNewFromCallWithoutParms[*IEnumShellItems](me, releaser,
		utl.Vt[_IEnumShellItemsVt](me.Ppvt()).Clone)
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
//	var desktop *winsh.IShellItem
//	_ = winsh.SHGetKnownFolderItem(
//		rel,
//		&cosh.FOLDERID_Desktop,
//		cosh.KF_DEFAULT,
//		win.HANDLE(0),
//		&desktop,
//	)
//
//	var enumItems *winsh.IEnumShellItems
//	_ = desktop.BindToHandler(rel, nil, &cosh.BHID_EnumItems, &enumItems)
//
//	items, _ := enumItems.Enum(rel)
//	for _, item := range items {
//		path, _ := item.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//		println(path)
//	}
func (me *IEnumShellItems) Enum(releaser *win.OleReleaser) ([]*IShellItem, error) {
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
func (me *IEnumShellItems) Next(releaser *win.OleReleaser) (*IShellItem, error) {
	var ppvtQueried uintptr
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumShellItemsVt](me.Ppvt()).Next,
		me.Ppvt(),
		1,
		uintptr(unsafe.Pointer(&ppvtQueried)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := utl.OleNew[*IShellItem](ppvtQueried, releaser)
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
	return utl.OleCallWithoutParms(me, utl.Vt[_IEnumShellItemsVt](me.Ppvt()).Reset)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumshellitems-skip
func (me *IEnumShellItems) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumShellItemsVt](me.Ppvt()).Skip,
		me.Ppvt(),
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}
