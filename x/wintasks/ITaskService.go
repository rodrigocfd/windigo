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

// [ITaskService] COM interface.
//
// [ITaskService]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itaskservice
type ITaskService struct{ winaut.IDispatch }

type _ITaskServiceVt struct {
	utl.IDispatchVt
	GetFolder           uintptr
	GetRunningTasks     uintptr
	NewTask             uintptr
	Connect             uintptr
	Get_Connected       uintptr
	Get_TargetServer    uintptr
	Get_ConnectedUser   uintptr
	Get_ConnectedDomain uintptr
	Get_HighestVersion  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskService) IID() *co.IID {
	return &cotasks.IID_ITaskService
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *ITaskService) AddRef(releaser *win.OleReleaser) *ITaskService {
	return utl.OleNewFromAddRef[*ITaskService](me, releaser)
}

// [Connect] method.
//
// [Connect]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-connect
func (me *ITaskService) Connect(serverName, user, domain, password string) error {
	localRel := win.NewOleReleaser()
	defer localRel.Release()

	vServerName := winaut.NewVariant(localRel, serverName)
	vUser := winaut.NewVariant(localRel, user)
	vDomain := winaut.NewVariant(localRel, domain)
	vPassword := winaut.NewVariant(localRel, password)

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskServiceVt](me.Ppvt()).Connect,
		me.Ppvt(),
		uintptr(unsafe.Pointer(vServerName)),
		uintptr(unsafe.Pointer(vUser)),
		uintptr(unsafe.Pointer(vDomain)),
		uintptr(unsafe.Pointer(vPassword)))
	return utl.HresultToError(ret)
}

// [get_Connected] method.
//
// [get_Connected]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_connected
func (me *ITaskService) GetConnected() (bool, error) {
	var connected int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskServiceVt](me.Ppvt()).Get_Connected,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&connected)))
	return utl.HresultToBoolError(int32(connected), ret)
}

// [get_ConnectedDomain] method.
//
// [get_ConnectedDomain]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_connecteddomain
func (me *ITaskService) GetConnectedDomain() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskServiceVt](me.Ppvt()).Get_ConnectedDomain)
}

// [get_ConnectedUser] method.
//
// [get_ConnectedUser]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_connecteduser
func (me *ITaskService) GetConnectedUser() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskServiceVt](me.Ppvt()).Get_ConnectedUser)
}

// [get_HighestVersion] method.
//
// [get_HighestVersion]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_highestversion
func (me *ITaskService) GetHighestVersion() (int, error) {
	var version uint32
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskServiceVt](me.Ppvt()).Get_HighestVersion,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&version)))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		return 0, hr
	}
	return int(version), nil
}

// [get_TargetServer] method.
//
// [get_TargetServer]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-get_targetserver
func (me *ITaskService) GetTargetServer() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskServiceVt](me.Ppvt()).Get_TargetServer)
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-getfolder
func (me *ITaskService) GetFolder(releaser *win.OleReleaser, path string) (*ITaskFolder, error) {
	var ppvtQueried uintptr

	bstrPath, err := winaut.SysAllocString(path)
	if err != nil {
		return nil, err
	}
	defer bstrPath.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskServiceVt](me.Ppvt()).GetFolder,
		me.Ppvt(),
		uintptr(bstrPath),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITaskFolder](ret, ppvtQueried, releaser)
}

// [NewTask] method.
//
// [NewTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskservice-newtask
func (me *ITaskService) NewTask(releaser *win.OleReleaser) (*ITaskDefinition, error) {
	var ppvtQueried uintptr
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskServiceVt](me.Ppvt()).NewTask,
		me.Ppvt(),
		0,
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITaskDefinition](ret, ppvtQueried, releaser)
}
