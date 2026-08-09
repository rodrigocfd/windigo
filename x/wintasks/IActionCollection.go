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

// [IActionCollection] COM interface.
//
// [IActionCollection]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iactioncollection
type IActionCollection struct{ winaut.IDispatch }

type _IActionCollectionVt struct {
	utl.IDispatchVt
	Get_Count    uintptr
	Get_Item     uintptr
	Get__NewEnum uintptr
	Get_XmlText  uintptr
	Put_XmlText  uintptr
	Create       uintptr
	Remove       uintptr
	Clear        uintptr
	Get_Context  uintptr
	Put_Context  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IActionCollection) IID() *co.IID {
	return &cotasks.IID_IAction
}

// [Clear] method.
//
// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-clear
func (me *IActionCollection) Clear() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_IActionCollectionVt](me.Ppvt()).Clear)
}

// [Create] method.
//
// [Create]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-create
func (me *IActionCollection) Create(
	releaser *win.OleReleaser,
	actionType cotasks.TASK_ACTION,
) (*IAction, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IActionCollectionVt](me.Ppvt()).Create,
		me.Ppvt(),
		uintptr(actionType),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IAction](ret, ppvtQueried, releaser)
}

// Returns all [IAction] objects by calling [IActionCollection.GetCount] and
// [IActionCollection.GetItem].
func (me *IActionCollection) Enum(releaser *win.OleReleaser) ([]*IAction, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	actions := make([]*IAction, count)
	for i := 0; i < count; i++ {
		action, err := me.GetItem(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		actions = append(actions, action)
	}
	return actions, nil
}

// [get_Context] method.
//
// [get_Context]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_context
func (me *IActionCollection) GetContext() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IActionCollectionVt](me.Ppvt()).Get_Context)
}

// [get_Count] method.
//
// [get_Count]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_count
func (me *IActionCollection) GetCount() (int, error) {
	var count int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IActionCollectionVt](me.Ppvt()).Get_Count,
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
// [get_Item]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_item
func (me *IActionCollection) GetItem(releaser *win.OleReleaser, index int) (*IAction, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IActionCollectionVt](me.Ppvt()).Get_Item,
		me.Ppvt(),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IAction](ret, ppvtQueried, releaser)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-get_xmltext
func (me *IActionCollection) GetXmlText() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IActionCollectionVt](me.Ppvt()).Get_XmlText)
}

// [put_Context] method.
//
// [put_Context]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-put_context
func (me *IActionCollection) PutContext(context string) error {
	return oleCallSetBstr(me, context, utl.Vt[_IActionCollectionVt](me.Ppvt()).Put_Context)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iactioncollection-put_xmltext
func (me *IActionCollection) PutXmlText(text string) error {
	return oleCallSetBstr(me, text, utl.Vt[_IActionCollectionVt](me.Ppvt()).Put_XmlText)
}
