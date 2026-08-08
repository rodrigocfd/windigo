//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [IExecAction] COM interface.
//
// Implements [OleResource].
//
// [IExecAction]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iexecaction
type IExecAction struct{ IAction }

type _IExecActionVt struct {
	_IActionVt
	Get_Path             uintptr
	Put_Path             uintptr
	Get_Arguments        uintptr
	Put_Arguments        uintptr
	Get_WorkingDirectory uintptr
	Put_WorkingDirectory uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IExecAction) IID() *co.IID {
	return &cotasks.IID_IExecAction
}

// [get_Arguments] method.
//
// [get_Arguments]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-get_arguments
func (me *IExecAction) GetArguments() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IExecActionVt](me.Ppvt()).Get_Arguments)
}

// [get_Path] method.
//
// [get_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-get_path
func (me *IExecAction) GetPath() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IExecActionVt](me.Ppvt()).Get_Path)
}

// [get_WorkingDirectory] method.
//
// [get_WorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-get_workingdirectory
func (me *IExecAction) GetWorkingDirectory() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IExecActionVt](me.Ppvt()).Get_WorkingDirectory)
}

// [put_Arguments] method.
//
// [put_Arguments]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-put_arguments
func (me *IExecAction) PutArguments(arguments string) error {
	return oleCallSetBstr(me, arguments, utl.Vt[_IExecActionVt](me.Ppvt()).Put_Arguments)
}

// [put_Path] method.
//
// [put_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-put_path
func (me *IExecAction) PutPath(path string) error {
	return oleCallSetBstr(me, path, utl.Vt[_IExecActionVt](me.Ppvt()).Put_Path)
}

// [put_WorkingDirectory] method.
//
// [put_WorkingDirectory]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iexecaction-put_workingdirectory
func (me *IExecAction) PutWorkingDirectory(workingDirectory string) error {
	return oleCallSetBstr(me, workingDirectory, utl.Vt[_IExecActionVt](me.Ppvt()).Put_WorkingDirectory)
}
