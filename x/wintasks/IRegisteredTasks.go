//go:build windows

package wintasks

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
	"github.com/rodrigocfd/windigo/x/winaut"
)

// [IRegisteredTask] COM interface.
//
// Implements [OleResource].
//
// [IRegisteredTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iregisteredtask
type IRegisteredTask struct{ winaut.IDispatch }

type _IRegisteredTaskVt struct {
	utl.IDispatchVt
	Get_Name               uintptr
	Get_Path               uintptr
	Get_State              uintptr
	Get_Enabled            uintptr
	Put_Enabled            uintptr
	Run                    uintptr
	RunEx                  uintptr
	GetInstances           uintptr
	Get_LastRunTime        uintptr
	Get_LastTaskResult     uintptr
	Get_NumberOfMissedRuns uintptr
	Get_NextRunTime        uintptr
	Get_Definition         uintptr
	Get_Xml                uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
	Stop                   uintptr
	GetRunTimes            uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IRegisteredTask) IID() *co.IID {
	return &cotasks.IID_IRegisteredTask
}

// [get_Definition] method.
//
// [get_Definition]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_definition
func (me *IRegisteredTask) GetDefinition(releaser *win.OleReleaser) (*ITaskDefinition, error) {
	return utl.OleNewFromCallWithoutParms[*ITaskDefinition](me, releaser,
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_Definition)
}

// [get_Enabled] method.
//
// [get_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itrigger-get_enabled
func (me *IRegisteredTask) GetEnabled() (bool, error) {
	var enabled int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_Enabled,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&enabled)))
	return utl.HresultToBoolError(int32(enabled), ret)
}

// [get_LastTaskResult] method.
//
// [get_LastTaskResult]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_lasttaskresult
func (me *IRegisteredTask) GetLastTaskResult() (int, error) {
	var last int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_LastTaskResult,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&last)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(last), nil
	} else {
		return 0, hr
	}
}

// [get_Name] method.
//
// [get_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_name
func (me *IRegisteredTask) GetName() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_Name)
}

// [get_NumberOfMissedRuns] method.
//
// [get_NumberOfMissedRuns]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_numberofmissedruns
func (me *IRegisteredTask) GetNumberOfMissedRuns() (int, error) {
	var nmr int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_NumberOfMissedRuns,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&nmr)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(nmr), nil
	} else {
		return 0, hr
	}
}

// [get_Path] method.
//
// [get_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_path
func (me *IRegisteredTask) GetPath() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_Path)
}

// [GetRunTimes] method.
//
// [GetRunTimes]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-getruntimes
func (me *IRegisteredTask) GetRunTimes(start, end time.Time, count int) ([]time.Time, error) {
	var stStart, stEnd win.SYSTEMTIME
	stStart.SetTime(start)
	stEnd.SetTime(end)

	count32 := uint32(count)
	var pRunTimes win.HTASKMEM // will be allocated by the OS
	defer pRunTimes.CoTaskMemFree()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).GetRunTimes,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&stStart)),
		uintptr(unsafe.Pointer(&stEnd)),
		uintptr(unsafe.Pointer(&count32)),
		uintptr(unsafe.Pointer(&pRunTimes)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		memStRunTimes := unsafe.Slice((*win.SYSTEMTIME)(unsafe.Pointer(pRunTimes)), count32)
		runTimes := make([]time.Time, 0, count32)
		for _, st := range memStRunTimes {
			runTimes = append(runTimes, st.ToTime())
		}
		return runTimes, nil
	} else {
		return nil, hr
	}
}

// [get_State] method.
//
// [get_State]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_state
func (me *IRegisteredTask) GetState() (cotasks.TASK_STATE, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_STATE](me,
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_State)
}

// [get_Xml] method.
//
// [get_Xml]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-get_xml
func (me *IRegisteredTask) GetXml() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Get_Xml)
}

// [put_Enabled] method.
//
// [put_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-put_enabled
func (me *IRegisteredTask) PutEnabled(enabled bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Put_Enabled,
		me.Ppvt(),
		utl.BoolToUintptr(enabled))
	return utl.HresultToError(ret)
}

// [Stop] method.
//
// [Stop]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iregisteredtask-stop
func (me *IRegisteredTask) Stop() error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IRegisteredTaskVt](me.Ppvt()).Stop,
		me.Ppvt(),
		0)
	return utl.HresultToError(ret)
}
