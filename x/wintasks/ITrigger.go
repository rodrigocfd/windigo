//go:build windows

package wintasks

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [ITrigger] COM interface.
//
// [ITrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itrigger
type ITrigger struct{ winaut.IDispatch }

type _ITriggerVt struct {
	utl.IDispatchVt
	Get_Type               uintptr
	Get_Id                 uintptr
	Put_Id                 uintptr
	Get_Repetition         uintptr
	Put_Repetition         uintptr
	Get_ExecutionTimeLimit uintptr
	Put_ExecutionTimeLimit uintptr
	Get_StartBoundary      uintptr
	Put_StartBoundary      uintptr
	Get_EndBoundary        uintptr
	Put_EndBoundary        uintptr
	Get_Enabled            uintptr
	Put_Enabled            uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITrigger) IID() *co.IID {
	return &cotasks.IID_ITrigger
}

// [get_Enabled] method.
//
// [get_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_enabled
func (me *ITrigger) GetEnabled() (bool, error) {
	var enabled int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITriggerVt](me.Ppvt()).Get_Enabled,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&enabled)))
	return utl.HresultToBoolError(int32(enabled), ret)
}

// [get_EndBoundary] method.
//
// [get_EndBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_endboundary
func (me *ITrigger) GetEndBoundary() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITriggerVt](me.Ppvt()).Get_EndBoundary)
}

// [get_ExecutionTimeLimit] method.
//
// [get_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_executiontimelimit
func (me *ITrigger) GetExecutionTimeLimit() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITriggerVt](me.Ppvt()).Get_ExecutionTimeLimit)
}

// [get_Id] method.
//
// [get_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_id
func (me *ITrigger) GetId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITriggerVt](me.Ppvt()).Get_Id)
}

// [get_StartBoundary] method.
//
// [get_StartBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_startboundary
func (me *ITrigger) GetStartBoundary() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITriggerVt](me.Ppvt()).Get_StartBoundary)
}

// [get_Type] method.
//
// [get_Type]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_type
func (me *ITrigger) GetType() (cotasks.TASK_TRIGGER2, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_TRIGGER2](me,
		utl.Vt[_ITriggerVt](me.Ppvt()).Get_Type)
}

// [put_Enabled] method.
//
// [put_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_enabled
func (me *ITrigger) PutEnabled(enabled bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITriggerVt](me.Ppvt()).Put_Enabled,
		me.Ppvt(),
		utl.BoolToUintptr(enabled))
	return utl.HresultToError(ret)
}

// [put_EndBoundary] method.
//
// [put_EndBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_endboundary
func (me *ITrigger) PutEndBoundary(end string) error {
	return oleCallSetBstr(me, end, utl.Vt[_ITriggerVt](me.Ppvt()).Put_EndBoundary)
}

// [put_ExecutionTimeLimit] method.
//
// [put_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_executiontimelimit
func (me *ITrigger) PutExecutionTimeLimit(timeLimit string) error {
	return oleCallSetBstr(me, timeLimit, utl.Vt[_ITriggerVt](me.Ppvt()).Put_ExecutionTimeLimit)
}

// [put_Id] method.
//
// [put_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_id
func (me *ITrigger) PutId(id string) error {
	return oleCallSetBstr(me, id, utl.Vt[_ITriggerVt](me.Ppvt()).Put_Id)
}

// [put_StartBoundary] method.
//
// [put_StartBoundary]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-put_startboundary
func (me *ITrigger) PutStartBoundary(start string) error {
	return oleCallSetBstr(me, start, utl.Vt[_ITriggerVt](me.Ppvt()).Put_StartBoundary)
}
