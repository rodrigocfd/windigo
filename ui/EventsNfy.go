package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type (
	_HashCmd struct {
		cmdId     int
		notifCode co.CMD
	}
	_HashNfy struct {
		idFrom int
		code   co.NM
	}
)

//------------------------------------------------------------------------------

// Ordinary events for WM messages and WM_TIMER, plus WM_COMMAND and WM_NOTIFY.
// If an event for the given message already exists, it will be overwritten.
type _EventsNfy struct {
	_EventsWm
	cmdsZero map[_HashCmd]func(p wm.Command)
	nfysRet  map[_HashNfy]func(p unsafe.Pointer) uintptr // meaningful return value
	nfysZero map[_HashNfy]func(p unsafe.Pointer)         // just returns zero (or TRUE if dialog)
}

func (me *_EventsNfy) new() {
	me._EventsWm.new()
	me.cmdsZero = make(map[_HashCmd]func(p wm.Command))
	me.nfysRet = make(map[_HashNfy]func(p unsafe.Pointer) uintptr)
	me.nfysZero = make(map[_HashNfy]func(p unsafe.Pointer))
}

// Adds a WM_COMMAND event with no meaningful return value, always returning zero.
func (me *_EventsNfy) addCmdZero(cmdId int, notifCode co.CMD, userFunc func(p wm.Command)) {
	me.cmdsZero[_HashCmd{
		cmdId:     cmdId,
		notifCode: notifCode,
	}] = userFunc
}

// Adds a WM_NOTIFY event with a meaningful return value.
func (me *_EventsNfy) addNfyRet(idFrom int, code co.NM, userFunc func(p unsafe.Pointer) uintptr) {
	me.nfysRet[_HashNfy{
		idFrom: idFrom,
		code:   code,
	}] = userFunc
}

// Adds a WM_NOTIFY event with no meaningful return value, always returning zero.
func (me *_EventsNfy) addNfyZero(idFrom int, code co.NM, userFunc func(p unsafe.Pointer)) {
	me.nfysZero[_HashNfy{
		idFrom: idFrom,
		code:   code,
	}] = userFunc
}

func (me *_EventsNfy) hasMessages() bool {
	return len(me.cmdsZero) > 0 ||
		len(me.nfysRet) > 0 ||
		len(me.nfysZero) > 0 ||
		me._EventsWm.hasMessages()
}

func (me *_EventsNfy) processMessage(
	uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM,
) (retVal uintptr, meaningfulRet bool, wasHandled bool) {

	if uMsg == co.WM_COMMAND {
		hash := _HashCmd{
			cmdId:     int(wParam.LoWord()),
			notifCode: co.CMD(wParam.HiWord()),
		}
		if userFunc, hasFunc := me.cmdsZero[hash]; hasFunc {
			msgObj := wm.Any{WParam: wParam, LParam: lParam}
			userFunc(wm.Command{Msg: msgObj})
			retVal, meaningfulRet, wasHandled = 0, false, true
		}

	} else if uMsg == co.WM_NOTIFY {
		nmhdrPtr := unsafe.Pointer(lParam)
		nmhdr := (*win.NMHDR)(nmhdrPtr)
		hash := _HashNfy{
			idFrom: int(nmhdr.IdFrom),
			code:   co.NM(nmhdr.Code),
		}

		if userFunc, hasFunc := me.nfysZero[hash]; hasFunc {
			userFunc(nmhdrPtr)
			retVal, meaningfulRet, wasHandled = 0, false, true
			return

		} else if userFunc, hasFunc := me.nfysRet[hash]; hasFunc {
			retVal, meaningfulRet, wasHandled = userFunc(nmhdrPtr), true, true
			return
		}
	}

	return me._EventsWm.processMessage(uMsg, wParam, lParam) // delegate WM processing
}

// Generic WM_COMMAND handler.
//
// Avoid this method, prefer the specific command notification handlers.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsNfy) WmCommand(cmdId int, notifCode co.CMD, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, notifCode, userFunc)
}

// Handles a WM_COMMAND notification when a menu item is clicked.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsNfy) WmCommandMenu(cmdId int, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, co.CMD_MENU, userFunc)
}

// Handles a WM_COMMAND notification when a keyboard shortcut key is hit.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsNfy) WmCommandAccel(cmdId int, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, co.CMD_ACCELERATOR, userFunc)
}

// Handles a WM_COMMAND notification when a keyboard shortcut key is hit, or a
// menu item is clicked.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsNfy) WmCommandAccelMenu(cmdId int, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, co.CMD_MENU, userFunc)
	me.addCmdZero(cmdId, co.CMD_ACCELERATOR, userFunc)
}

// Generic WM_NOTIFY handler.
//
// Avoid this method, prefer the specific notification handlers.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-notify
func (me *_EventsNfy) WmNotify(idFrom int, code co.NM, userFunc func(p unsafe.Pointer) uintptr) {
	me.addNfyRet(idFrom, code, userFunc)
}
