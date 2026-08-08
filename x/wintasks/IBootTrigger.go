//go:build windows

package wintasks

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [IBootTrigger] COM interface.
//
// Implements [OleResource].
//
// [IBootTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-iboottrigger
type IBootTrigger struct{ ITrigger }

type _IBootTriggerVt struct {
	_ITriggerVt
	Get_Delay uintptr
	Put_Delay uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IBootTrigger) IID() *co.IID {
	return &cotasks.IID_IBootTrigger
}

// [get_Delay] method.
//
// [get_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iboottrigger-get_delay
func (me *IBootTrigger) GetDelay() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IBootTriggerVt](me.Ppvt()).Get_Delay)
}

// [put_Delay] method.
//
// [put_Delay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-iboottrigger-put_delay
func (me *IBootTrigger) PutDelay(delay string) error {
	return oleCallSetBstr(me, delay, utl.Vt[_IBootTriggerVt](me.Ppvt()).Put_Delay)
}
