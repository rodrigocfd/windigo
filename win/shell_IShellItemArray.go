//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
)

// [IShellItemArray] COM interface.
//
// Implements [OleObj] and [OleResource].
//
// [IShellItemArray]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitemarray
type IShellItemArray struct{ IUnknown }

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IShellItemArray) IID() co.IID {
	return co.IID_IShellItemArray
}

// Returns the path names of each [IShellItem] object by calling
// [IShellItemArray.GetCount], [IShellItemArray.GetItemAt] and
// [IShellItem.GetDisplayName].
//
// Example:
//
//	var arr win.IShellItemArray // initialized somewhere
//
//	names, _ := arr.EnumDisplayNames(co.SIGDN_FILESYSPATH)
//	for _, fullPath := range names {
//		println(fullPath)
//	}
func (me *IShellItemArray) EnumDisplayNames(sigdnName co.SIGDN) ([]string, error) {
	localRel := NewOleReleaser()
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
//	var arr win.IShellItemArray // initialized somewhere
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	items, _ := arr.EnumItems(rel)
//	for _, item := range items {
//		fullPath, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
//		println(fullPath)
//	}
func (me *IShellItemArray) EnumItems(releaser *OleReleaser) ([]*IShellItem, error) {
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
		(*_IShellItemArrayVt)(unsafe.Pointer(*me.Ppvt())).GetCount,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(unsafe.Pointer(&count)),
		0)

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [GetItemAt] method.
//
// Panics if index is negative.
//
// [GetItemAt]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishellitemarray-getitemat
func (me *IShellItemArray) GetItemAt(releaser *OleReleaser, index int) (*IShellItem, error) {
	utl.PanicNeg(index)
	var ppvtQueried **_IUnknownVt

	ret, _, _ := syscall.SyscallN(
		(*_IShellItemArrayVt)(unsafe.Pointer(*me.Ppvt())).GetItemAt,
		uintptr(unsafe.Pointer(me.Ppvt())),
		uintptr(uint32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IShellItem{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

type _IShellItemArrayVt struct {
	_IUnknownVt
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr
	GetItemAt                  uintptr
	EnumItems                  uintptr
}
