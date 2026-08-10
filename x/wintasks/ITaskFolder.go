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

// [ITaskFolder] COM interface.
//
// [ITaskFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-itaskfolder
type ITaskFolder struct{ winaut.IDispatch }

type _ITaskFolderVt struct {
	utl.IDispatchVt
	Get_Name               uintptr
	Get_Path               uintptr
	GetFolder              uintptr
	GetFolders             uintptr
	CreateFolder           uintptr
	DeleteFolder           uintptr
	GetTask                uintptr
	GetTasks               uintptr
	DeleteTask             uintptr
	RegisterTask           uintptr
	RegisterTaskDefinition uintptr
	GetSecurityDescriptor  uintptr
	SetSecurityDescriptor  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ITaskFolder) IID() *co.IID {
	return &cotasks.IID_ITaskFolder
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *ITaskFolder) AddRef(releaser *win.OleReleaser) *ITaskFolder {
	return utl.OleNewFromAddRef[*ITaskFolder](me, releaser)
}

// [DeleteFolder] method.
//
// [DeleteFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-deletefolder
func (me *ITaskFolder) DeleteFolder(subFolderName string) error {
	bstrSubFolderName, err := winaut.SysAllocString(subFolderName)
	if err != nil {
		return err
	}
	defer bstrSubFolderName.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskFolderVt](me.Ppvt()).DeleteFolder,
		me.Ppvt(),
		uintptr(bstrSubFolderName),
		0)
	return utl.HresultToError(ret)
}

// [DeleteTask] method.
//
// [DeleteTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-deletetask
func (me *ITaskFolder) DeleteTask(name string) error {
	bstrName, err := winaut.SysAllocString(name)
	if err != nil {
		return err
	}
	defer bstrName.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskFolderVt](me.Ppvt()).DeleteTask,
		me.Ppvt(),
		uintptr(bstrName),
		0)
	return utl.HresultToError(ret)
}

// [get_Name] method.
//
// [get_Name]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-get_name
func (me *ITaskFolder) GetName() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskFolderVt](me.Ppvt()).Get_Name)
}

// [get_Path] method.
//
// [get_Path]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-get_path
func (me *ITaskFolder) GetPath() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ITaskFolderVt](me.Ppvt()).Get_Path)
}

// [GetFolder] method.
//
// [GetFolder]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-getfolder
func (me *ITaskFolder) GetFolder(releaser *win.OleReleaser, path string) (*ITaskFolder, error) {
	var ppvtQueried uintptr

	bstrPath, err := winaut.SysAllocString(path)
	if err != nil {
		return nil, err
	}
	defer bstrPath.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskFolderVt](me.Ppvt()).GetFolder,
		me.Ppvt(),
		uintptr(bstrPath),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*ITaskFolder](ret, ppvtQueried, releaser)
}

// [GetTask] method.
//
// [GetTask]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-itaskfolder-gettask
func (me *ITaskFolder) GetTask(releaser *win.OleReleaser, path string) (*IRegisteredTask, error) {
	var ppvtQueried uintptr

	bstrPath, err := winaut.SysAllocString(path)
	if err != nil {
		return nil, err
	}
	defer bstrPath.SysFreeString()

	ret, _, _ := syscall.SyscallN(
		utl.Vt[_ITaskFolderVt](me.Ppvt()).GetTask,
		me.Ppvt(),
		uintptr(bstrPath),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return utl.OleNewIfOk[*IRegisteredTask](ret, ppvtQueried, releaser)
}
