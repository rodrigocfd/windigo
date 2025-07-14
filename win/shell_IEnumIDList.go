//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// [IEnumIDList] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// # Example
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	_ = item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(rel, win.HWND(0),
//		co.SHCONTF_FOLDERS|co.SHCONTF_NONFOLDERS|co.SHCONTF_INCLUDEHIDDEN)
//
// [IEnumIDList]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ienumidlist
type IEnumIDList struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEnumIDList) IID() co.IID {
	return co.IID_IEnumIDList
}

// [Clone] method.
//
// [Clone]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-clone
func (me *IEnumIDList) Clone(releaser *OleReleaser) (*IEnumIDList, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Clone,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IEnumIDList{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

// Returns all [ITEMIDLIST] values by calling [IEnumIDList.Next].
//
// # Example
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var item *win.IShellItem
//	_ = win.SHCreateItemFromParsingName(rel, "C:\\Temp", &item)
//
//	var folder *win.IShellFolder
//	_ = item.BindToHandler(rel, nil, co.BHID_SFObject, &folder)
//
//	pidlList, _ := folder.EnumObjects(
//		rel,
//		win.HWND(0),
//		co.SHCONTF_FOLDERS|co.SHCONTF_NONFOLDERS|co.SHCONTF_INCLUDEHIDDEN,
//	)
//
//	pidls, _ := pidlList.Enum(rel)
//	for _, pidl := range pidls {
//		var child *win.IShellItem
//		_ = win.SHCreateItemFromIDList(rel, pidl, &child)
//		name, _ := child.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(name)
//	}
func (me *IEnumIDList) Enum(releaser *OleReleaser) ([]*ITEMIDLIST, error) {
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
func (me *IEnumIDList) Next(releaser *OleReleaser) (*ITEMIDLIST, error) {
	var idlFetched ITEMIDLIST
	var numFetched uint32

	ret, _, _ := syscall.SyscallN(
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Next,
		uintptr(unsafe.Pointer(me.Ppvt())),
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
	ret, _, _ := syscall.SyscallN(
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Reset,
		uintptr(unsafe.Pointer(me.Ppvt())))
	return utl.ErrorAsHResult(ret)
}

// [Skip] method.
//
// [Skip]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ienumidlist-skip
func (me *IEnumIDList) Skip(count uint) error {
	ret, _, _ := syscall.SyscallN(
		(*_IEnumIDListVt)(unsafe.Pointer(*me.Ppvt())).Skip,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(count)))
	return utl.ErrorAsHResult(ret)
}

type _IEnumIDListVt struct {
	_IUnknownVt
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}
