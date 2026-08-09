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

// [IPrincipal] COM interface.
//
// [IPrincipal]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iprincipal
type IPrincipal struct{ winaut.IDispatch }

type _IPrincipalVt struct {
	utl.IDispatchVt
	Get_Id          uintptr
	Put_Id          uintptr
	Get_DisplayName uintptr
	Put_DisplayName uintptr
	Get_UserId      uintptr
	Put_UserId      uintptr
	Get_LogonType   uintptr
	Put_LogonType   uintptr
	Get_GroupId     uintptr
	Put_GroupId     uintptr
	Get_RunLevel    uintptr
	Put_RunLevel    uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IPrincipal) IID() *co.IID {
	return &cotasks.IID_IPrincipal
}

// [get_DisplayName] method.
//
// [get_DisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_displayname
func (me *IPrincipal) GetDisplayName() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IPrincipalVt](me.Ppvt()).Get_DisplayName)
}

// [get_GroupId] method.
//
// [get_GroupId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_groupid
func (me *IPrincipal) GetGroupId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IPrincipalVt](me.Ppvt()).Get_GroupId)
}

// [get_Id] method.
//
// [get_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_id
func (me *IPrincipal) GetId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IPrincipalVt](me.Ppvt()).Get_Id)
}

// [get_LogonType] method.
//
// [get_LogonType]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_logontype
func (me *IPrincipal) GetLogonType() (cotasks.TASK_LOGON, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_LOGON](me,
		utl.Vt[_IPrincipalVt](me.Ppvt()).Get_LogonType)
}

// [get_RunLevel] method.
//
// [get_RunLevel]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_runlevel
func (me *IPrincipal) GetRunLevel() (cotasks.TASK_RUNLEVEL, error) {
	return utl.OleCallReturnStruct[cotasks.TASK_RUNLEVEL](me,
		utl.Vt[_IPrincipalVt](me.Ppvt()).Get_RunLevel)
}

// [get_UserId] method.
//
// [get_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-get_userid
func (me *IPrincipal) GetUserId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IPrincipalVt](me.Ppvt()).Get_UserId)
}

// [put_DisplayName] method.
//
// [put_DisplayName]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_displayname
func (me *IPrincipal) PutDisplayName(name string) error {
	return oleCallSetBstr(me, name, utl.Vt[_IPrincipalVt](me.Ppvt()).Put_DisplayName)
}

// [put_GroupId] method.
//
// [put_GroupId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_groupid
func (me *IPrincipal) PutGroupId(groupId string) error {
	return oleCallSetBstr(me, groupId, utl.Vt[_IPrincipalVt](me.Ppvt()).Put_GroupId)
}

// [put_Id] method.
//
// [put_Id]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_id
func (me *IPrincipal) PutId(id string) error {
	return oleCallSetBstr(me, id, utl.Vt[_IPrincipalVt](me.Ppvt()).Put_Id)
}

// [put_LogonType] method.
//
// [put_LogonType]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_logontype
func (me *IPrincipal) PutLogonType(logon cotasks.TASK_LOGON) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPrincipalVt](me.Ppvt()).Put_LogonType,
		me.Ppvt(),
		uintptr(logon))
	return utl.HresultToError(ret)
}

// [put_RunLevel] method.
//
// [put_RunLevel]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_runlevel
func (me *IPrincipal) PutRunLevel(runLevel cotasks.TASK_RUNLEVEL) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IPrincipalVt](me.Ppvt()).Put_RunLevel,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&runLevel)))
	return utl.HresultToError(ret)
}

// [put_UserId] method.
//
// [put_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iprincipal-put_userid
func (me *IPrincipal) PutUserId(userId string) error {
	return oleCallSetBstr(me, userId, utl.Vt[_IPrincipalVt](me.Ppvt()).Put_UserId)
}
