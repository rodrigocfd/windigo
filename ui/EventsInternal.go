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
	msgsZero map[co.WM][]func(p wm.Any)            // ordinary WM messages
	nfysZero map[_HashNfy][]func(p unsafe.Pointer) // WM_NOTIFY messages
}

func (me *_EventsInternal) new() {
	me.msgsZero = make(map[co.WM][]func(p wm.Any), 5) // arbitrary
	me.nfysZero = make(map[_HashNfy][]func(p unsafe.Pointer), 10)
}

// Adds a WM event.
func (me *_EventsInternal) addMsgZero(uMsg co.WM, userFunc func(p wm.Any)) {
	var slice []func(p wm.Any)

	if existingSlice, hasSlice := me.msgsZero[uMsg]; hasSlice { // at least 1 handle exists?
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
	me.msgsZero[uMsg] = append(slice, userFunc)
}

// Adds a WM_NOTIFY event.
func (me *_EventsInternal) addNfyZero(
	idFrom int, code co.NM, userFunc func(p unsafe.Pointer)) {

	hash := _HashNfy{idFrom, code}
	var slice []func(p unsafe.Pointer)

	if existingSlice, hasSlice := me.nfysZero[hash]; hasSlice { // at least 1 handle exists?
		slice = existingSlice
	} else { // no handlers for this message yet
		slice = make([]func(p unsafe.Pointer), 0, 1)
	}
	me.nfysZero[hash] = append(slice, userFunc)
}

// Executes all handlers for the given message.
func (me *_EventsInternal) processMessages(
	uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) {

	if uMsg == co.WM_NOTIFY {
		nmhdrPtr := unsafe.Pointer(lParam)
		nmhdr := (*win.NMHDR)(nmhdrPtr)
		hash := _HashNfy{
			idFrom: int(nmhdr.IdFrom),
			code:   co.NM(nmhdr.Code),
		}

		if slice, hasSlice := me.nfysZero[hash]; hasSlice {
			for _, userFunc := range slice {
				userFunc(nmhdrPtr)
			}
		}

	} else { // ordinary WM message
		if slice, hasSlice := me.msgsZero[uMsg]; hasSlice {
			for _, userFunc := range slice {
				userFunc(wm.Any{WParam: wParam, LParam: lParam})
			}
		}
	}
}
