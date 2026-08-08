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

// [ITaskSettings] COM interface.
//
// Implements [OleResource].
//
// [ITaskSettings]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itasksettings
type ITaskSettings struct{ winaut.IDispatch }

type _ITaskSettingsVt struct {
	utl.IDispatchVt
	Get_AllowDemandStart           uintptr
	Put_AllowDemandStart           uintptr
	Get_RestartInterval            uintptr
	Put_RestartInterval            uintptr
	Get_RestartCount               uintptr
	Put_RestartCount               uintptr
	Get_MultipleInstances          uintptr
	Put_MultipleInstances          uintptr
	Get_StopIfGoingOnBatteries     uintptr
	Put_StopIfGoingOnBatteries     uintptr
	Get_DisallowStartIfOnBatteries uintptr
	Put_DisallowStartIfOnBatteries uintptr
	Get_AllowHardTerminate         uintptr
	Put_AllowHardTerminate         uintptr
	Get_StartWhenAvailable         uintptr
	Put_StartWhenAvailable         uintptr
	Get_XmlText                    uintptr
	Put_XmlText                    uintptr
	Get_RunOnlyIfNetworkAvailable  uintptr
	Put_RunOnlyIfNetworkAvailable  uintptr
	Get_ExecutionTimeLimit         uintptr
	Put_ExecutionTimeLimit         uintptr
	Get_Enabled                    uintptr
	Put_Enabled                    uintptr
	Get_DeleteExpiredTaskAfter     uintptr
	Put_DeleteExpiredTaskAfter     uintptr
	Get_Priority                   uintptr
	Put_Priority                   uintptr
	Get_Compatibility              uintptr
	Put_Compatibility              uintptr
	Get_Hidden                     uintptr
	Put_Hidden                     uintptr
	Get_IdleSettings               uintptr
	Put_IdleSettings               uintptr
	Get_RunOnlyIfIdle              uintptr
	Put_RunOnlyIfIdle              uintptr
	Get_WakeToRun                  uintptr
	Put_WakeToRun                  uintptr
	Get_NetworkSettings            uintptr
	Put_NetworkSettings            uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskSettings) IID() *co.IID {
	return &cotasks.IID_ITaskSettings
}

// [get_AllowDemandStart] method.
//
// [get_AllowDemandStart]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_allowdemandstart
func (me *ITaskSettings) GetAllowDemandStart() (bool, error) {
	var allow int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_AllowDemandStart,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&allow)))
	return utl.HresultToBoolError(int32(allow), ret)
}

// [get_AllowHardTerminate] method.
//
// [get_AllowHardTerminate]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_allowhardterminate
func (me *ITaskSettings) GetAllowHardTerminate() (bool, error) {
	var allow int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_AllowHardTerminate,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&allow)))
	return utl.HresultToBoolError(int32(allow), ret)
}

// [get_Compatibility] method.
//
// [get_Compatibility]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_compatibility
func (me *ITaskSettings) GetCompatibility() (cotasks.TASK_COMPATIBILITY, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_COMPATIBILITY](me,
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_Compatibility)
}

// [get_DeleteExpiredTaskAfter] method.
//
// [get_DeleteExpiredTaskAfter]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_deleteexpiredtaskafter
func (me *ITaskSettings) GetDeleteExpiredTaskAfter() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_DeleteExpiredTaskAfter)
}

// [get_DisallowStartIfOnBatteries] method.
//
// [get_DisallowStartIfOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_disallowstartifonbatteries
func (me *ITaskSettings) GetDisallowStartIfOnBatteries() (bool, error) {
	var disallow int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_DisallowStartIfOnBatteries,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&disallow)))
	return utl.HresultToBoolError(int32(disallow), ret)
}

// [get_Enabled] method.
//
// [get_Enabled]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_enabled
func (me *ITaskSettings) GetEnabled() (bool, error) {
	var enabled int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_Enabled,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&enabled)))
	return utl.HresultToBoolError(int32(enabled), ret)
}

// [get_ExecutionTimeLimit] method.
//
// [get_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_executiontimelimit
func (me *ITaskSettings) GetExecutionTimeLimit() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_ExecutionTimeLimit)
}

// [get_Hidden] method.
//
// [get_Hidden]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_hidden
func (me *ITaskSettings) GetHidden() (bool, error) {
	var hidden int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_Hidden,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&hidden)))
	return utl.HresultToBoolError(int32(hidden), ret)
}

// [get_MultipleInstances] method.
//
// [get_MultipleInstances]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_multipleinstances
func (me *ITaskSettings) GetMultipleInstances() (cotasks.TASK_INSTANCES, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_INSTANCES](me,
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_MultipleInstances)
}

// [get_Priority] method.
//
// [get_Priority]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_priority
func (me *ITaskSettings) GetPriority() (int, error) {
	var priority int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_Priority,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&priority)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(priority), nil
	} else {
		return 0, hr
	}
}

// [get_RestartCount] method.
//
// [get_RestartCount]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_restartcount
func (me *ITaskSettings) GetRestartCount() (int, error) {
	var priority int32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_RestartCount,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&priority)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(priority), nil
	} else {
		return 0, hr
	}
}

// [get_RestartInterval] method.
//
// [get_RestartInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_restartinterval
func (me *ITaskSettings) GetRestartInterval() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_RestartInterval)
}

// [get_RunOnlyIfIdle] method.
//
// [get_RunOnlyIfIdle]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_runonlyifidle
func (me *ITaskSettings) GetRunOnlyIfIdle() (bool, error) {
	var runOnlyIfIdle int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_RunOnlyIfIdle,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&runOnlyIfIdle)))
	return utl.HresultToBoolError(int32(runOnlyIfIdle), ret)
}

// [get_RunOnlyIfNetworkAvailable] method.
//
// [get_RunOnlyIfNetworkAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_runonlyifnetworkavailable
func (me *ITaskSettings) GetRunOnlyIfNetworkAvailable() (bool, error) {
	var runOnlyIfNetworkAvailable int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_RunOnlyIfNetworkAvailable,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&runOnlyIfNetworkAvailable)))
	return utl.HresultToBoolError(int32(runOnlyIfNetworkAvailable), ret)
}

// [get_StartWhenAvailable] method.
//
// [get_StartWhenAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_startwhenavailable
func (me *ITaskSettings) GetStartWhenAvailable() (bool, error) {
	var startWhenAvailable int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_StartWhenAvailable,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&startWhenAvailable)))
	return utl.HresultToBoolError(int32(startWhenAvailable), ret)
}

// [get_StopIfGoingOnBatteries] method.
//
// [get_StopIfGoingOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_stopifgoingonbatteries
func (me *ITaskSettings) GetStopIfGoingOnBatteries() (bool, error) {
	var stopIfGoingOnBatt int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_StopIfGoingOnBatteries,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&stopIfGoingOnBatt)))
	return utl.HresultToBoolError(int32(stopIfGoingOnBatt), ret)
}

// [get_WakeToRun] method.
//
// [get_WakeToRun]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_waketorun
func (me *ITaskSettings) GetWakeToRun() (bool, error) {
	var wakeToRun int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_WakeToRun,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&wakeToRun)))
	return utl.HresultToBoolError(int32(wakeToRun), ret)
}

// [get_XmlText] method.
//
// [get_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-get_xmltext
func (me *ITaskSettings) GetXmlText() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Get_XmlText)
}

// [put_AllowDemandStart] method.
//
// [put_AllowDemandStart]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_allowdemandstart
func (me *ITaskSettings) PutAllowDemandStart(allowDemandStart bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_AllowDemandStart,
		me.Ppvt(),
		utl.BoolToUintptr(allowDemandStart))
	return utl.HresultToError(ret)
}

// [put_AllowHardTerminate] method.
//
// [put_AllowHardTerminate]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_allowhardterminate
func (me *ITaskSettings) PutAllowHardTerminate(allowHardTerminate bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_AllowHardTerminate,
		me.Ppvt(),
		utl.BoolToUintptr(allowHardTerminate))
	return utl.HresultToError(ret)
}

// [put_Compatibility] method.
//
// [put_Compatibility]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_compatibility
func (me *ITaskSettings) PutCompatibility(compatLevel cotasks.TASK_COMPATIBILITY) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_Compatibility,
		me.Ppvt(),
		uintptr(compatLevel))
	return utl.HresultToError(ret)
}

// [put_DeleteExpiredTaskAfter] method.
//
// [put_DeleteExpiredTaskAfter]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_deleteexpiredtaskafter
func (me *ITaskSettings) PutDeleteExpiredTaskAfter(expirationDelay string) error {
	return oleCallSetBstr(me, expirationDelay, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_DeleteExpiredTaskAfter)
}

// [put_DisallowStartIfOnBatteries] method.
//
// [put_DisallowStartIfOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_disallowstartifonbatteries
func (me *ITaskSettings) PutDisallowStartIfOnBatteries(disallowStart bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_DisallowStartIfOnBatteries,
		me.Ppvt(),
		utl.BoolToUintptr(disallowStart))
	return utl.HresultToError(ret)
}

// [put_ExecutionTimeLimit] method.
//
// [put_ExecutionTimeLimit]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_executiontimelimit
func (me *ITaskSettings) PutExecutionTimeLimit(executionTimeLimit string) error {
	return oleCallSetBstr(me, executionTimeLimit, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_ExecutionTimeLimit)
}

// [put_Hidden] method.
//
// [put_Hidden]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_hidden
func (me *ITaskSettings) PutHidden(hidden bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_Hidden,
		me.Ppvt(),
		utl.BoolToUintptr(hidden))
	return utl.HresultToError(ret)
}

// [put_MultipleInstances] method.
//
// [put_MultipleInstances]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_multipleinstances
func (me *ITaskSettings) PutMultipleInstances(policy cotasks.TASK_INSTANCES) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_MultipleInstances,
		me.Ppvt(),
		uintptr(policy))
	return utl.HresultToError(ret)
}

// [put_Priority] method.
//
// [put_Priority]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_priority
func (me *ITaskSettings) PutPriority(priority int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_Priority,
		me.Ppvt(),
		uintptr(int32(priority)))
	return utl.HresultToError(ret)
}

// [put_RestartCount] method.
//
// [put_RestartCount]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_restartcount
func (me *ITaskSettings) PutRestartCount(restartCount int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_RestartCount,
		me.Ppvt(),
		uintptr(int32(restartCount)))
	return utl.HresultToError(ret)
}

// [put_RestartInterval] method.
//
// [put_RestartInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_restartinterval
func (me *ITaskSettings) PutRestartInterval(restartInterval string) error {
	return oleCallSetBstr(me, restartInterval, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_RestartInterval)
}

// [put_RunOnlyIfIdle] method.
//
// [put_RunOnlyIfIdle]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_runonlyifidle
func (me *ITaskSettings) PutRunOnlyIfIdle(runOnlyIfIdle bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_RunOnlyIfIdle,
		me.Ppvt(),
		utl.BoolToUintptr(runOnlyIfIdle))
	return utl.HresultToError(ret)
}

// [put_RunOnlyIfNetworkAvailable] method.
//
// [put_RunOnlyIfNetworkAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_runonlyifnetworkavailable
func (me *ITaskSettings) PutRunOnlyIfNetworkAvailable(runOnlyIfNetworkAvailable bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_RunOnlyIfNetworkAvailable,
		me.Ppvt(),
		utl.BoolToUintptr(runOnlyIfNetworkAvailable))
	return utl.HresultToError(ret)
}

// [put_StartWhenAvailable] method.
//
// [put_StartWhenAvailable]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_startwhenavailable
func (me *ITaskSettings) PutStartWhenAvailable(startWhenAvailable bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_StartWhenAvailable,
		me.Ppvt(),
		utl.BoolToUintptr(startWhenAvailable))
	return utl.HresultToError(ret)
}

// [put_StopIfGoingOnBatteries] method.
//
// [put_StopIfGoingOnBatteries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_stopifgoingonbatteries
func (me *ITaskSettings) PutStopIfGoingOnBatteries(stopIfOnBatteries bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_StartWhenAvailable,
		me.Ppvt(),
		utl.BoolToUintptr(stopIfOnBatteries))
	return utl.HresultToError(ret)
}

// [put_WakeToRun] method.
//
// [put_WakeToRun]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_waketorun
func (me *ITaskSettings) PutWakeToRun(wake bool) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_WakeToRun,
		me.Ppvt(),
		utl.BoolToUintptr(wake))
	return utl.HresultToError(ret)
}

// [put_XmlText] method.
//
// [put_XmlText]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itasksettings-put_xmltext
func (me *ITaskSettings) PutXmlText(text string) error {
	return oleCallSetBstr(me, text, utl.Vt[_ITaskSettingsVt](me.Ppvt()).Put_XmlText)
}
