//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [IAction] COM interface.
//
// [IAction]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iaction
type IAction struct{ winaut.IDispatch }

type _IActionVt struct {
	utl.IDispatchVt
	Get_Id   uintptr
	Put_Id   uintptr
	Get_Type uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IAction) IID() *co.IID {
	return &cotasks.IID_IAction
}

// [get_Id] method.
//
// [get_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iaction-get_id
func (me *IAction) GetId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IActionVt](me.Ppvt()).Get_Id)
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iaction-get_type
func (me *IAction) GetType() (cotasks.TASK_ACTION, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_ACTION](me,
		utl.Vt[_IActionVt](me.Ppvt()).Get_Type)
}

// [put_Id] method.
//
// [put_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iaction-put_id
func (me *IAction) PutId(id string) error {
	return oleCallSetBstr(me, id, utl.Vt[_IActionVt](me.Ppvt()).Put_Id)
}
