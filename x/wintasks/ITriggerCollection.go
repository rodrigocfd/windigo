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

// [ITriggerCollection] COM interface.
//
// [ITriggerCollection]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itriggercollection
type ITriggerCollection struct{ winaut.IDispatch }

type _ITriggerCollectionVt struct {
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
func (*ITriggerCollection) IID() *co.IID {
	return &cotasks.IID_ITriggerCollection
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *ITriggerCollection) AddRef(releaser *win.OleReleaser) *ITriggerCollection {
	return utl.OleNewFromAddRef[*ITriggerCollection](me, releaser)
}

// [Clear] method.
//
// [Clear]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-clear
func (me *ITriggerCollection) Clear() error {
	return utl.OleCallWithoutParms(me, utl.Vt[_ITriggerCollectionVt](me.Ppvt()).Clear)
}

// [Create] method.
//
// [Create]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-create
func (me *ITriggerCollection) Create(
	releaser *win.OleReleaser,
	triggerType cotasks.TASK_TRIGGER2,
) (*ITrigger, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITriggerCollectionVt](me.Ppvt()).Create,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&triggerType)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITrigger](ret, ppvtQueried, releaser)
}

// Returns all [ITrigger] objects by calling [ITriggerCollection.GetCount] and
// [ITriggerCollection.GetItem].
func (me *ITriggerCollection) Enum(releaser *win.OleReleaser) ([]*ITrigger, error) {
	count, err := me.GetCount()
	if err != nil {
		return nil, err
	}

	triggers := make([]*ITrigger, 0, count)
	for i := 0; i < count; i++ {
		trigger, err := me.GetItem(releaser, i)
		if err != nil {
			return nil, err // stop immediately
		}
		triggers = append(triggers, trigger)
	}
	return triggers, nil
}

// [get_Count] method.
//
// [get_Count]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-get_count
func (me *ITriggerCollection) GetCount() (int, error) {
	var count int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITriggerCollectionVt](me.Ppvt()).Get_Count,
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
// [get_Item]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itriggercollection-get_item
func (me *ITriggerCollection) GetItem(releaser *win.OleReleaser, index int) (*ITrigger, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITriggerCollectionVt](me.Ppvt()).Get_Item,
		me.Ppvt(),
		uintptr(int32(index)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITrigger](ret, ppvtQueried, releaser)
}
