//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Events added only internally by the library, cannot be added by the user.
// Supports multiple events for the same message, all will be executed.
type _EventsInternal struct {
	msgsNoRet map[co.WM][]func(p wm.Any)            // ordinary WM messages
	nfysNoRet map[_HashNfy][]func(p unsafe.Pointer) // WM_NOTIFY messages
}

func (me *_EventsInternal) clear() {
	for key := range me.msgsNoRet {
		delete(me.msgsNoRet, key)
	}
	for key := range me.nfysNoRet {
		delete(me.nfysNoRet, key)
	}
}

func (me *_EventsInternal) new() {
	me.msgsNoRet = make(map[co.WM][]func(p wm.Any), 5) // arbitrary
	me.nfysNoRet = make(map[_HashNfy][]func(p unsafe.Pointer), 10)
}

// Adds a WM event.
func (me *_EventsInternal) addMsgNoRet(uMsg co.WM, userFunc func(p wm.Any)) {
	var slice []func(p wm.Any)

	if existingSlice, hasSlice := me.msgsNoRet[uMsg]; hasSlice { // at least 1 handle exists?
		slice = existingSlice
	} else { // no handlers for this message yet
		capacity := 1
		if uMsg == co.WM_CREATE || uMsg == co.WM_INITDIALOG { // special optimization cases
			capacity = 10
		} else if uMsg == co.WM_SIZE {
			capacity = 3
		}

		slice = make([]func(p wm.Any), 0, capacity)
	}
	me.msgsNoRet[uMsg] = append(slice, userFunc)
}

// Adds a WM_NOTIFY event.
func (me *_EventsInternal) addNfyNoRet(
	idFrom int, code co.NM, userFunc func(p unsafe.Pointer)) {

	hash := _HashNfy{idFrom, code}
	var slice []func(p unsafe.Pointer)

	if existingSlice, hasSlice := me.nfysNoRet[hash]; hasSlice { // at least 1 handle exists?
		slice = existingSlice
	} else { // no handlers for this message yet
		slice = make([]func(p unsafe.Pointer), 0, 1)
	}
	me.nfysNoRet[hash] = append(slice, userFunc)
}

// Executes all handlers for the given message.
func (me *_EventsInternal) processMessages(
	uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) (atLeast1 bool) {

	if uMsg == co.WM_NOTIFY {
		nmhdrPtr := unsafe.Pointer(lParam)
		nmhdr := (*win.NMHDR)(nmhdrPtr)
		hash := _HashNfy{
			idFrom: int(nmhdr.IdFrom),
			code:   co.NM(nmhdr.Code),
		}

		if userFuncs, hasFuncs := me.nfysNoRet[hash]; hasFuncs {
			atLeast1 = true
			for _, userFunc := range userFuncs {
				userFunc(nmhdrPtr)
			}
		}

	} else { // ordinary WM message
		if userFuncs, hasFuncs := me.msgsNoRet[uMsg]; hasFuncs {
			atLeast1 = true
			for _, userFunc := range userFuncs {
				userFunc(wm.Any{WParam: wParam, LParam: lParam})
			}
		}
	}
	return
}
