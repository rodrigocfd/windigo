//go:build windows

package wintasks

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [IEventTrigger] COM interface.
//
// [IEventTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-ieventtrigger
type IEventTrigger struct{ ITrigger }

type _IEventTriggerVt struct {
	_ITriggerVt
	Get_Subscription uintptr
	Put_Subscription uintptr
	Get_Delay        uintptr
	Put_Delay        uintptr
	Get_ValueQueries uintptr
	Put_ValueQueries uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IEventTrigger) IID() *co.IID {
	return &cotasks.IID_IEventTrigger
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IEventTrigger) AddRef(releaser *win.OleReleaser) *IEventTrigger {
	return utl.OleNewFromAddRef[*IEventTrigger](me, releaser)
}

// [get_Delay] method.
//
// [get_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-get_delay
func (me *IEventTrigger) GetDelay() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEventTriggerVt](me.Ppvt()).Get_Delay)
}

// [get_Subscription] method.
//
// [get_Subscription]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-get_subscription
func (me *IEventTrigger) GetSubscription() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IEventTriggerVt](me.Ppvt()).Get_Subscription)
}

// [get_ValueQueries] method.
//
// [get_ValueQueries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-get_valuequeries
func (me *IEventTrigger) GetValueQueries(releaser *win.OleReleaser) (*ITaskNamedValueCollection, error) {
	return utl.OleNewFromCallWithoutParms[*ITaskNamedValueCollection](me, releaser,
		utl.Vt[_IEventTriggerVt](me.Ppvt()).Get_ValueQueries)
}

// [put_Delay] method.
//
// [put_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-put_delay
func (me *IEventTrigger) PutDelay(delay string) error {
	return oleCallSetBstr(me, delay, utl.Vt[_IEventTriggerVt](me.Ppvt()).Put_Delay)
}

// [put_Subscription] method.
//
// [put_Subscription]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-put_subscription
func (me *IEventTrigger) PutSubscription(subscription string) error {
	return oleCallSetBstr(me, subscription, utl.Vt[_IEventTriggerVt](me.Ppvt()).Put_Subscription)
}

// [put_ValueQueries] method.
//
// [put_ValueQueries]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-ieventtrigger-put_valuequeries
func (me *IEventTrigger) PutValueQueries(namedXPaths *ITaskNamedValueCollection) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IEventTriggerVt](me.Ppvt()).Put_ValueQueries,
		me.Ppvt(),
		namedXPaths.Ppvt())
	return utl.HresultToError(ret)
}
