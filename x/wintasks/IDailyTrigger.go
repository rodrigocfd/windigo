//go:build windows

package wintasks

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/x/cotasks"
)

// [IDailyTrigger] COM interface.
//
// [IDailyTrigger]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nn-taskschd-idailytrigger
type IDailyTrigger struct{ ITrigger }

type _IDailyTriggerVt struct {
	_ITriggerVt
	Get_DaysInterval uintptr
	Put_DaysInterval uintptr
	Get_RandomDelay  uintptr
	Put_RandomDelay  uintptr
}

// Returns the unique COM [interface ID].
//
// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
func (*IDailyTrigger) IID() *co.IID {
	return &cotasks.IID_IDailyTrigger
}

// [AddRef] method.
//
// [AddRef]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (me *IDailyTrigger) AddRef(releaser *win.OleReleaser) *IDailyTrigger {
	return utl.OleNewFromAddRef[*IDailyTrigger](me, releaser)
}

// [get_DaysInterval] method.
//
// [get_DaysInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-get_daysinterval
func (me *IDailyTrigger) GetDaysInterval() (int, error) {
	var days int16
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDailyTriggerVt](me.Ppvt()).Get_DaysInterval,
		me.Ppvt(),
		uintptr(unsafe.Pointer(&days)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return int(days), nil
	} else {
		return 0, hr
	}
}

// [get_RandomDelay] method.
//
// [get_RandomDelay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-get_randomdelay
func (me *IDailyTrigger) GetRandomDelay() (string, error) {
	return oleCallRetBstr(me, utl.Vt[_IDailyTriggerVt](me.Ppvt()).Get_RandomDelay)
}

// [put_DaysInterval] method.
//
// [put_DaysInterval]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-put_daysinterval
func (me *IDailyTrigger) PutDaysInterval(days int) error {
	ret, _, _ := syscall.SyscallN(
		utl.Vt[_IDailyTriggerVt](me.Ppvt()).Put_DaysInterval,
		me.Ppvt(),
		uintptr(int16(days)))
	return utl.HresultToError(ret)
}

// [put_RandomDelay] method.
//
// [put_RandomDelay]: https://learn.microsoft.com/en-us/windows/win32/api/taskschd/nf-taskschd-idailytrigger-put_randomdelay
func (me *IDailyTrigger) PutRandomDelay(randomDelay string) error {
	return oleCallSetBstr(me, randomDelay, utl.Vt[_IDailyTriggerVt](me.Ppvt()).Put_RandomDelay)
}
