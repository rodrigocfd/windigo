//go:build windows

package wintasks

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [ITaskNamedValueCollection] COM interface.
//
// Implements [OleResource].
//
// [ITaskNamedValueCollection]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itasknamedvaluecollection
type ITaskNamedValueCollection struct{ winaut.IDispatch }

type _ITaskNamedValueCollectionVt struct {
	utl.IDispatchVt
	Get_Count    uintptr
	Get_Item     uintptr
	Get__NewEnum uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskNamedValueCollection) IID() *co.IID {
	return &cotasks.IID_ITaskNamedValueCollection
}

// [Clear] method;
//
// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-clear
func (me *ITaskNamedValueCollection) Clear() error {
	return utl.OleCallWithoutParms(me,
		utl.Vt[_ITaskNamedValueCollectionVt](me.Ppvt()).Clear)
}

// [Create] method.
//
// [Create]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-create
func (me *ITaskNamedValueCollection) Create(
	releaser *win.OleReleaser,
	name, value string,
) (*ITaskNamedValuePair, error) {
	var ppvtQueried uintptr

	bstrName, err := winaut.SysAllocString(name)
	if err != nil {
		return nil, err
	}
	defer bstrName.SysFreeString()

	bstrValue, err := winaut.SysAllocString(value)
	if err != nil {
		return nil, err
	}
	defer bstrValue.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskNamedValueCollectionVt](me.Ppvt()).Create,
		me.Ppvt(),
		uintptr(bstrName),
		uintptr(bstrValue),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITaskNamedValuePair](ret, ppvtQueried, releaser)
}

// Returns all [ITaskNamedValuePair] objects by calling
// [ITaskNamedValueCollection.GetCount] and [ITaskNamedValueCollection.GetItem].
func (me *ITaskNamedValueCollection) Enum(releaser *win.OleReleaser) ([]*ITaskNamedValuePair, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	pairs := make([]*ITaskNamedValuePair, count)
	for i := 0; i < count; i++ {
		pair, err := me.GetItem(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		pairs = append(pairs, pair)
	}
	return pairs, nil
}

// [get_Count] method.
//
// [get_Count]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-get_count
func (me *ITaskNamedValueCollection) GetCount() (int, error) {
	var count int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskNamedValueCollectionVt](me.Ppvt()).Get_Count,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&count)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(count), nil
	} else {
		return 0, hr
	}
}

// [get_Item] method.
//
// [get_Item]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-get_item
func (me *ITaskNamedValueCollection) GetItem(releaser *win.OleReleaser, index int) (*ITaskNamedValuePair, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskNamedValueCollectionVt](me.Ppvt()).Get_Item,
		me.Ppvt(),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITaskNamedValuePair](ret, ppvtQueried, releaser)
}

// [Remove] method.
//
// [Remove]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasknamedvaluecollection-remove
func (me *ITaskNamedValueCollection) Remove(index int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskNamedValueCollectionVt](me.Ppvt()).Remove,
		me.Ppvt(),
		uintptr(int32(index)))
	return utl.HresultToError(ret)
}
