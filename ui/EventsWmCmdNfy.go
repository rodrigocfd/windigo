/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

type (
	_CmdHash struct {
		CmdId     uint16
		NotifCode uint16
	}
	_NfyHash struct {
		IdFrom int
		Code   co.NM
	}
)

//------------------------------------------------------------------------------

// Besides WM and WM_TIMER, also keeps WM_COMMAND and WM_NOTIFY message callbacks.
type _EventsWmCmdNfy struct {
	*_EventsWm
	mapCmds    map[_CmdHash]func(p WmCommand)
	mapNfysRet map[_NfyHash]func(p unsafe.Pointer) uintptr // meaningful return value
	mapNfys    map[_NfyHash]func(p unsafe.Pointer)         // just returns zero (or TRUE if dialog)
}

// Constructor.
func _NewEventsWmCmdNfy() *_EventsWmCmdNfy {
	return &_EventsWmCmdNfy{
		_EventsWm:  _NewEventsWm(),
		mapCmds:    make(map[_CmdHash]func(p WmCommand)),
		mapNfysRet: make(map[_NfyHash]func(p unsafe.Pointer) uintptr),
		mapNfys:    make(map[_NfyHash]func(p unsafe.Pointer)),
	}
}

func (me *_EventsWmCmdNfy) processMessage(
	msg co.WM, p Wm) (retVal uintptr, useRetVal bool, wasHandled bool) {

	if msg == co.WM_COMMAND {
		hash := _CmdHash{
			CmdId:     p.WParam.LoWord(),
			NotifCode: p.WParam.HiWord(),
		}
		if userFunc, hasFunc := me.mapCmds[hash]; hasFunc {
			userFunc(WmCommand{m: p})
			retVal, useRetVal, wasHandled = 0, false, true
		}

	} else if msg == co.WM_NOTIFY {
		nmhdrPtr := unsafe.Pointer(p.LParam)
		nmhdr := (*win.NMHDR)(nmhdrPtr)
		hash := _NfyHash{
			IdFrom: int(nmhdr.IdFrom),
			Code:   co.NM(nmhdr.Code),
		}

		if userFunc, hasFunc := me.mapNfys[hash]; hasFunc {
			userFunc(nmhdrPtr)
			retVal, useRetVal, wasHandled = 0, false, true
			return

		} else if userFunc, hasFunc := me.mapNfysRet[hash]; hasFunc {
			retVal, useRetVal, wasHandled = userFunc(nmhdrPtr), true, true
			return
		}
	}

	return me._EventsWm.processMessage(msg, p) // delegate WM processing
}

func (me *_EventsWmCmdNfy) hasMessages() bool {
	return len(me.mapCmds) > 0 ||
		len(me.mapNfysRet) > 0 ||
		len(me.mapNfys) > 0 ||
		me._EventsWm.hasMessages()
}

func (me *_EventsWmCmdNfy) addNfyRet(
	idFrom int, code co.NM, userFunc func(p unsafe.Pointer) uintptr) {

	hash := _NfyHash{
		IdFrom: idFrom,
		Code:   code,
	}
	me.mapNfysRet[hash] = userFunc
}

func (me *_EventsWmCmdNfy) addNfy(
	idFrom int, code co.NM, userFunc func(p unsafe.Pointer)) {

	hash := _NfyHash{
		IdFrom: idFrom,
		Code:   code,
	}
	me.mapNfys[hash] = userFunc
}

//------------------------------------------------------------------------------

// Generic WM_COMMAND handler.
//
// Avoid this method, prefer the specific command notification handlers.
//
// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmCmdNfy) WmCommand(cmdId int, notifCode int, userFunc func(p WmCommand)) {
	hash := _CmdHash{
		CmdId:     uint16(cmdId),
		NotifCode: uint16(notifCode),
	}
	me.mapCmds[hash] = userFunc
}

// Handles a WM_COMMAND notification when a menu item is clicked.
//
// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmCmdNfy) WmCommandMenu(cmdId int, userFunc func(p WmCommand)) {
	me.WmCommand(cmdId, 0, userFunc)
}

// Handles a WM_COMMAND notification when a keyboard shortcut key is hit.
//
// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmCmdNfy) WmCommandAccel(cmdId int, userFunc func(p WmCommand)) {
	me.WmCommand(cmdId, 1, userFunc)
}

// Handles a WM_COMMAND notification when a keyboard shortcut key is hit, or a
// menu item is clicked.
//
// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_EventsWmCmdNfy) WmCommandAccelMenu(cmdId int, userFunc func(p WmCommand)) {
	me.WmCommand(cmdId, 0, userFunc)
	me.WmCommand(cmdId, 1, userFunc)
}

type WmCommand struct{ m Wm }

func (p WmCommand) IsFromMenu() bool        { return p.m.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool { return p.m.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool     { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() int             { return p.ControlId() }
func (p WmCommand) AcceleratorId() int      { return p.ControlId() }
func (p WmCommand) ControlId() int          { return int(p.m.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() int   { return int(p.m.WParam.HiWord()) }
func (p WmCommand) ControlHwnd() win.HWND   { return win.HWND(p.m.LParam) }

//------------------------------------------------------------------------------

// Generic WM_NOTIFY handler.
//
// Avoid this method, prefer the specific notification handlers.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/wm-notify
func (me *_EventsWmCmdNfy) WmNotify(idFrom int, code co.NM, userFunc func(p unsafe.Pointer) uintptr) {
	me.addNfyRet(idFrom, code, userFunc)
}
