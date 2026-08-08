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

// [IEnumIDList] COM interface.
//
// Implements [OleResource].
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
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *winsh.IShellFolder
//	_ = item.BindToHandler(rel, nil, &cosh.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(rel, win.HWND(0),
//		cosh.SHCONTF_FOLDERS|cosh.SHCONTF_NONFOLDERS|cosh.SHCONTF_INCLUDEHIDDEN)
//
// [IEnumIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumidlist
type IEnumIDList struct{ win.IUnknown }

type _IEnumIDListVt struct {
	utl.IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumIDList) IID() *co.IID {
	return &cosh.IID_IEnumIDList
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-clone
func (me *IEnumIDList) Clone(releaser *win.OleReleaser) (*IEnumIDList, error) {
	return utl.OleNewFromCallWithoutParms[*IEnumIDList](me, releaser,
		utl.Vt[_IEnumIDListVt](me.Ppvt()).Clone)
}

// Returns all [ITEMIDLIST] values by calling [IEnumIDList.Next].
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
//	var item *winsh.IShellItem
//	_ = winsh.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *winsh.IShellFolder
//	_ = item.BindToHandler(rel, nil, &cosh.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(
//		rel,
//		win.HWND(0),
//		cosh.SHCONTF_FOLDERS|cosh.SHCONTF_NONFOLDERS|cosh.SHCONTF_INCLUDEHIDDEN,
//	)
//
//	pidls, _ := pidlList.Enum(rel)
//	for _, pidl := range pidls {
//		var child *winsh.IShellItem
//		_ = winsh.SHCreateItemFromIDList(rel, pidl, &child)
//		name, _ := child.GetDisplayName(cosh.SIGDN_FILESYSPATH)
//		println(name)
//	}
func (me *IEnumIDList) Enum(releaser *win.OleReleaser) ([]*ITEMIDLIST, error) {
	items := make([]*ITEMIDLIST, 0)
	var item *ITEMIDLIST
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
// This is a low-level method, prefer using [IEnumIDList.Enum].
//
// [Next]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-next
func (me *IEnumIDList) Next(releaser *win.OleReleaser) (*ITEMIDLIST, error) {
	var idlFetched ITEMIDLIST
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumIDListVt](me.Ppvt()).Next,
		me.Ppvt(),
		1,
		uintptr(unsafe.Pointer(&idlFetched)),
		uintptr(unsafe.Pointer(&numFetched)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pIdlFetched := &idlFetched
		releaser.Add(pIdlFetched)
		return pIdlFetched, nil
	} else if hr == co.HRESULT_S_FALSE {
		return nil, nil
	} else {
		return nil, hr
	}
}

// [Reset] method.
//
// [Reset]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-reset
func (me *IEnumIDList) Reset() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IEnumIDListVt](me.Ppvt()).Reset)
}

// [Skip] method.
//
// Panics if count is negative.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-skip
func (me *IEnumIDList) Skip(count int) error {
	utl.PanicNeg(count)
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEnumIDListVt](me.Ppvt()).Skip,
		me.Ppvt(),
		uintptr(uint32(count)))
	return utl.HresultToError(ret)
}
