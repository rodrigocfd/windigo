//go:build windows

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
// Used in parent window user events.
type _EventsWmNfy struct {
	_EventsWm
	cmdsZero map[_HashCmd]func(p wm.Command)
	nfysRet  map[_HashNfy]func(p unsafe.Pointer) uintptr // meaningful return value
	nfysZero map[_HashNfy]func(p unsafe.Pointer)         // just returns zero (or TRUE if dialog)
}

func (me *_EventsWmNfy) new() {
	me._EventsWm.new()
	me.cmdsZero = make(map[_HashCmd]func(p wm.Command), 10) // arbitrary
	me.nfysRet = make(map[_HashNfy]func(p unsafe.Pointer) uintptr, 10)
	me.nfysZero = make(map[_HashNfy]func(p unsafe.Pointer), 10)
}

func (me *_EventsWmNfy) clear() {
	me._EventsWm.clear()
	for key := range me.cmdsZero {
		delete(me.cmdsZero, key)
	}
	for key := range me.nfysRet {
		delete(me.nfysRet, key)
	}
	for key := range me.nfysZero {
		delete(me.nfysZero, key)
	}
}

// Adds a WM_COMMAND event with no meaningful return value, always returning zero.
func (me *_EventsWmNfy) addCmdZero(cmdId int, notifCode co.CMD, userFunc func(p wm.Command)) {
	me.cmdsZero[_HashCmd{
		cmdId:     cmdId,
		notifCode: notifCode,
	}] = userFunc
}

// Adds a WM_NOTIFY event with a meaningful return value.
func (me *_EventsWmNfy) addNfyRet(idFrom int, code co.NM, userFunc func(p unsafe.Pointer) uintptr) {
	me.nfysRet[_HashNfy{
		idFrom: idFrom,
		code:   code,
	}] = userFunc
}

// Adds a WM_NOTIFY event with no meaningful return value, always returning zero.
func (me *_EventsWmNfy) addNfyZero(idFrom int, code co.NM, userFunc func(p unsafe.Pointer)) {
	me.nfysZero[_HashNfy{
		idFrom: idFrom,
		code:   code,
	}] = userFunc
}

func (me *_EventsWmNfy) processMessage(
	uMsg co.WM,
	wParam win.WPARAM,
	lParam win.LPARAM) (retVal uintptr, meaningfulRet, wasHandled bool) {

	if uMsg == co.WM_COMMAND {
		hash := _HashCmd{
			cmdId:     int(wParam.LoWord()),
			notifCode: co.CMD(wParam.HiWord()),
		}
		if userFunc, hasFunc := me.cmdsZero[hash]; hasFunc {
			msgObj := wm.Any{WParam: wParam, LParam: lParam}
			userFunc(wm.Command{Msg: msgObj})
			return 0, false, true
		}

	} else if uMsg == co.WM_NOTIFY {
		pHdr := unsafe.Pointer(lParam)
		hdr := (*win.NMHDR)(pHdr)
		hash := _HashNfy{
			idFrom: int(hdr.IdFrom),
			code:   co.NM(hdr.Code),
		}

		if userFunc, hasFunc := me.nfysZero[hash]; hasFunc {
			userFunc(pHdr)
			return 0, false, true

		} else if userFunc, hasFunc := me.nfysRet[hash]; hasFunc {
			return userFunc(pHdr), true, true
		}
	}

	return me._EventsWm.processMessage(uMsg, wParam, lParam) // delegate WM processing
}

// Generic [WM_COMMAND] handler.
//
// Avoid this method, prefer the specific command notification handlers.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmNfy) WmCommand(cmdId int, notifCode co.CMD, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, notifCode, userFunc)
}

// Handles a [WM_COMMAND] notification when a menu item is clicked.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmNfy) WmCommandMenu(cmdId int, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, co.CMD_MENU, userFunc)
}

// Handles a [WM_COMMAND] notification when a keyboard shortcut key is hit.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmNfy) WmCommandAccel(cmdId int, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, co.CMD_ACCELERATOR, userFunc)
}

// Handles a [WM_COMMAND] notification when a keyboard shortcut key is hit, or a
// menu item is clicked.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmNfy) WmCommandAccelMenu(cmdId int, userFunc func(p wm.Command)) {
	me.addCmdZero(cmdId, co.CMD_MENU, userFunc)
	me.addCmdZero(cmdId, co.CMD_ACCELERATOR, userFunc)
}

// Generic [WM_NOTIFY] handler.
//
// Avoid this method, prefer the specific notification handlers.
//
// [WM_NOTIFY]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-notify
func (me *_EventsWmNfy) WmNotify(idFrom int, code co.NM, userFunc func(p unsafe.Pointer) uintptr) {
	me.addNfyRet(idFrom, code, userFunc)
}
