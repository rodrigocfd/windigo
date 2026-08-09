//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [ILogonTrigger] COM interface.
//
// [ILogonTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iLogontrigger
type ILogonTrigger struct{ ITrigger }

type _ILogonTriggerVt struct {
	_ITriggerVt
	Get_Delay  uintptr
	Put_Delay  uintptr
	Get_UserId uintptr
	Put_UserId uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*ILogonTrigger) IID() *co.IID {
	return &cotasks.IID_ILogonTrigger
}

// [get_Delay] method.
//
// [get_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-get_delay
func (me *ILogonTrigger) GetDelay() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ILogonTriggerVt](me.Ppvt()).Get_Delay)
}

// [get_UserId] method.
//
// [get_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-get_userid
func (me *ILogonTrigger) GetUserId() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_ILogonTriggerVt](me.Ppvt()).Get_UserId)
}

// [put_Delay] method.
//
// [put_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-put_delay
func (me *ILogonTrigger) PutDelay(delay string) error {
	return oleCallSetBstr(me, delay, utl.Vt[_ILogonTriggerVt](me.Ppvt()).Put_Delay)
}

// [put_UserId] method.
//
// [put_UserId]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ilogontrigger-put_userid
func (me *ILogonTrigger) PutUserId(userId string) error {
	return oleCallSetBstr(me, userId, utl.Vt[_ILogonTriggerVt](me.Ppvt()).Put_UserId)
}
