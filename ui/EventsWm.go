package ui

import (
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

// Ordinary events for WM messages and WM_TIMER.
// If an event for the given message already exists, it will be overwritten.
type _EventsWm struct {
	msgsRet  map[co.WM]func(p wm.Any) uintptr // meaningful return value
	msgsZero map[co.WM]func(p wm.Any)         // just returns zero (or TRUE if dialog)
	timers   map[int]func()                   // WM_TIMER
}

func (me *_EventsWm) new() {
	me.msgsRet = make(map[co.WM]func(p wm.Any) uintptr, 10) // arbitrary
	me.msgsZero = make(map[co.WM]func(p wm.Any), 10)
	me.timers = make(map[int]func(), 5)
}

// Adds a new WM message with a meaningful return value.
func (me *_EventsWm) addMsgRet(uMsg co.WM, userFunc func(p wm.Any) uintptr) {
	me.msgsRet[uMsg] = userFunc
}

// Adds a new WM message with no meaningful value, always returning zero.
func (me *_EventsWm) addMsgZero(uMsg co.WM, userFunc func(p wm.Any)) {
	me.msgsZero[uMsg] = userFunc
}

func (me *_EventsWm) hasMessages() bool {
	return len(me.msgsRet) > 0 ||
		len(me.msgsZero) > 0 ||
		len(me.timers) > 0
}

func (me *_EventsWm) processMessage(
	uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM,
) (retVal uintptr, meaningfulRet bool, wasHandled bool) {

	msgObj := wm.Any{WParam: wParam, LParam: lParam}

	if uMsg == co.WM_TIMER {
		if userFunc, hasFunc := me.timers[int(wParam)]; hasFunc {
			userFunc()
			retVal, meaningfulRet, wasHandled = 0, false, true
			return
		}

	} else if userFunc, hasFunc := me.msgsZero[uMsg]; hasFunc {
		userFunc(msgObj)
		retVal, meaningfulRet, wasHandled = 0, false, true
		return

	} else if userFunc, hasFunc := me.msgsRet[uMsg]; hasFunc {
		retVal, meaningfulRet, wasHandled = userFunc(msgObj), true, true
		return
	}

	retVal, meaningfulRet, wasHandled = 0, false, false
	return
}

// Generic message handler.
//
// Avoid this method, prefer the specific message handlers.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/learnwin32/window-messages
func (me *_EventsWm) Wm(uMsg co.WM, userFunc func(p wm.Any) uintptr) {
	me.addMsgRet(uMsg, userFunc)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-timer
func (me *_EventsWm) WmTimer(nIDEvent int, userFunc func()) {
	me.timers[nIDEvent] = userFunc
}

// ⚠️ By handling this message, you'll overwrite the default behavior in WindowMain.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-activate
func (me *_EventsWm) WmActivate(userFunc func(p wm.Activate)) {
	me.addMsgZero(co.WM_ACTIVATE, func(p wm.Any) {
		userFunc(wm.Activate{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-activateapp
func (me *_EventsWm) WmActivateApp(userFunc func(p wm.ActivateApp)) {
	me.addMsgZero(co.WM_ACTIVATEAPP, func(p wm.Any) {
		userFunc(wm.ActivateApp{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
func (me *_EventsWm) WmAppCommand(userFunc func(p wm.AppCommand)) {
	me.addMsgRet(co.WM_APPCOMMAND, func(p wm.Any) uintptr {
		userFunc(wm.AppCommand{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-askcbformatname
func (me *_EventsWm) WmAskCbFormatName(userFunc func(p wm.AskCbFormatName)) {
	me.addMsgZero(co.WM_ASKCBFORMATNAME, func(p wm.Any) {
		userFunc(wm.AskCbFormatName{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-cancelmode
func (me *_EventsWm) WmCancelMode(userFunc func()) {
	me.addMsgZero(co.WM_CANCELMODE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-capturechanged
func (me *_EventsWm) WmCaptureChanged(userFunc func(p wm.CaptureChanged)) {
	me.addMsgZero(co.WM_CAPTURECHANGED, func(p wm.Any) {
		userFunc(wm.CaptureChanged{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-changecbchain
func (me *_EventsWm) WmChangeCbChain(userFunc func(p wm.ChangeCbChain)) {
	me.addMsgZero(co.WM_CHANGECBCHAIN, func(p wm.Any) {
		userFunc(wm.ChangeCbChain{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (me *_EventsWm) WmChar(userFunc func(p wm.Char)) {
	me.addMsgZero(co.WM_CHAR, func(p wm.Any) {
		userFunc(wm.Char{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-chartoitem
func (me *_EventsWm) WmCharToItem(userFunc func(p wm.CharToItem) int) {
	me.addMsgRet(co.WM_CHARTOITEM, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CharToItem{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-childactivate
func (me *_EventsWm) WmChildActivate(userFunc func()) {
	me.addMsgZero(co.WM_CHILDACTIVATE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-clipboardupdate
func (me *_EventsWm) WmClipboardUpdate(userFunc func()) {
	me.addMsgZero(co.WM_CLIPBOARDUPDATE, func(_ wm.Any) {
		userFunc()
	})
}

// ⚠️ By handling this message, you'll overwrite the default behavior in WindowMain and WindowModal
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-close
func (me *_EventsWm) WmClose(userFunc func()) {
	me.addMsgZero(co.WM_CLOSE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-compareitem
func (me *_EventsWm) WmCompareItem(userFunc func(p wm.CompareItem) int) {
	me.addMsgRet(co.WM_COMPAREITEM, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CompareItem{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-contextmenu
func (me *_EventsWm) WmContextMenu(userFunc func(p wm.ContextMenu)) {
	me.addMsgZero(co.WM_CONTEXTMENU, func(p wm.Any) {
		userFunc(wm.ContextMenu{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-copydata
func (me *_EventsWm) WmCopyData(userFunc func(p wm.CopyData) bool) {
	me.addMsgRet(co.WM_COPYDATA, func(p wm.Any) uintptr {
		return util.BoolToUintptr(userFunc(wm.CopyData{Msg: p}))
	})
}

// Sent only to raw windows (not for dialog windows).
// Return 0 to continue window creation or -1 to abort it.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-create
func (me *_EventsWm) WmCreate(userFunc func(p wm.Create) int) {
	me.addMsgRet(co.WM_CREATE, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.Create{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorbtn
func (me *_EventsWm) WmCtlColorBtn(userFunc func(p wm.CtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORBTN, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CtlColor{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-ctlcolordlg
func (me *_EventsWm) WmCtlColorDlg(userFunc func(p wm.CtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORDLG, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CtlColor{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcoloredit
func (me *_EventsWm) WmCtlColorEdit(userFunc func(p wm.CtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLOREDIT, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CtlColor{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorlistbox
func (me *_EventsWm) WmCtlColorListBox(userFunc func(p wm.CtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORLISTBOX, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CtlColor{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorscrollbar
func (me *_EventsWm) WmCtlColorScrollBar(userFunc func(p wm.CtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORSCROLLBAR, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CtlColor{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
func (me *_EventsWm) WmCtlColorStatic(userFunc func(p wm.CtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORSTATIC, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.CtlColor{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
func (me *_EventsWm) WmDeadChar(userFunc func(p wm.Char)) {
	me.addMsgZero(co.WM_DEADCHAR, func(p wm.Any) {
		userFunc(wm.Char{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-deleteitem
func (me *_EventsWm) WmDeleteItem(userFunc func(p wm.DeleteItem)) {
	me.addMsgRet(co.WM_DELETEITEM, func(p wm.Any) uintptr {
		userFunc(wm.DeleteItem{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
func (me *_EventsWm) WmDestroy(userFunc func()) {
	me.addMsgZero(co.WM_DESTROY, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-destroyclipboard
func (me *_EventsWm) WmDestroyClipboard(userFunc func()) {
	me.addMsgZero(co.WM_DESTROYCLIPBOARD, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-devmodechange
func (me *_EventsWm) WmDevModeChange(userFunc func(wm.DevModeChange)) {
	me.addMsgZero(co.WM_DEVMODECHANGE, func(p wm.Any) {
		userFunc(wm.DevModeChange{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-displaychange
func (me *_EventsWm) WmDisplayChange(userFunc func(p wm.DisplayChange)) {
	me.addMsgZero(co.WM_DISPLAYCHANGE, func(p wm.Any) {
		userFunc(wm.DisplayChange{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-drawclipboard
func (me *_EventsWm) WmDrawClipboard(userFunc func()) {
	me.addMsgZero(co.WM_DRAWCLIPBOARD, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-drawitem
func (me *_EventsWm) WmDrawItem(userFunc func(p wm.DrawItem)) {
	me.addMsgRet(co.WM_DRAWITEM, func(p wm.Any) uintptr {
		userFunc(wm.DrawItem{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/shell/wm-dropfiles
func (me *_EventsWm) WmDropFiles(userFunc func(p wm.DropFiles)) {
	me.addMsgZero(co.WM_DROPFILES, func(p wm.Any) {
		userFunc(wm.DropFiles{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-enable
func (me *_EventsWm) WmEnable(userFunc func(p wm.Enable)) {
	me.addMsgZero(co.WM_ENABLE, func(p wm.Any) {
		userFunc(wm.Enable{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
func (me *_EventsWm) WmEndSession(userFunc func(p wm.EndSession)) {
	me.addMsgZero(co.WM_ENDSESSION, func(p wm.Any) {
		userFunc(wm.EndSession{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
func (me *_EventsWm) WmEnterIdle(userFunc func(p wm.EnterIdle)) {
	me.addMsgZero(co.WM_ENTERIDLE, func(p wm.Any) {
		userFunc(wm.EnterIdle{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-entermenuloop
func (me *_EventsWm) WmEnterMenuLoop(userFunc func(wm.EnterMenuLoop)) {
	me.addMsgZero(co.WM_ENTERMENULOOP, func(p wm.Any) {
		userFunc(wm.EnterMenuLoop{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-entersizemove
func (me *_EventsWm) WmEnterSizeMove(userFunc func()) {
	me.addMsgZero(co.WM_ENTERSIZEMOVE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-erasebkgnd
func (me *_EventsWm) WmEraseBkgnd(userFunc func(wm.EraseBkgnd) int) {
	me.addMsgRet(co.WM_ERASEBKGND, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.EraseBkgnd{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-exitmenuloop
func (me *_EventsWm) WmExitMenuLoop(userFunc func(wm.ExitMenuLoop)) {
	me.addMsgZero(co.WM_EXITMENULOOP, func(p wm.Any) {
		userFunc(wm.ExitMenuLoop{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-exitsizemove
func (me *_EventsWm) WmExitSizeMove(userFunc func()) {
	me.addMsgZero(co.WM_EXITSIZEMOVE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-fontchange
func (me *_EventsWm) WmFontChange(userFunc func()) {
	me.addMsgZero(co.WM_FONTCHANGE, func(p wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
func (me *_EventsWm) WmGetDlgCode(userFunc func(p wm.GetDlgCode) co.DLGC) {
	me.addMsgRet(co.WM_GETDLGCODE, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.GetDlgCode{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-getfont
func (me *_EventsWm) WmGetFont(userFunc func() win.HFONT) {
	me.addMsgRet(co.WM_FONTCHANGE, func(_ wm.Any) uintptr {
		return uintptr(userFunc())
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-geticon
func (me *_EventsWm) WmGetIcon(userFunc func(p wm.GetIcon) win.HICON) {
	me.addMsgRet(co.WM_GETICON, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.GetIcon{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-getminmaxinfo
func (me *_EventsWm) WmGetMinMaxInfo(userFunc func(p wm.GetMinMaxInfo)) {
	me.addMsgZero(co.WM_GETMINMAXINFO, func(p wm.Any) {
		userFunc(wm.GetMinMaxInfo{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-gettitlebarinfoex
func (me *_EventsWm) WmGetTitleBarInfoEx(userFunc func(p wm.GetTitleBarInfoEx)) {
	me.addMsgZero(co.WM_GETTITLEBARINFOEX, func(p wm.Any) {
		userFunc(wm.GetTitleBarInfoEx{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/shell/wm-help
func (me *_EventsWm) WmHelp(userFunc func(p wm.Help)) {
	me.addMsgRet(co.WM_HELP, func(p wm.Any) uintptr {
		userFunc(wm.Help{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
func (me *_EventsWm) WmHotKey(userFunc func(p wm.HotKey)) {
	me.addMsgZero(co.WM_HOTKEY, func(p wm.Any) {
		userFunc(wm.HotKey{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll
func (me *_EventsWm) WmHScroll(userFunc func(p wm.HScroll)) {
	me.addMsgZero(co.WM_HSCROLL, func(p wm.Any) {
		userFunc(wm.HScroll{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-hscrollclipboard
func (me *_EventsWm) WmHScrollClipboard(userFunc func(p wm.HScrollClipboard)) {
	me.addMsgZero(co.WM_HSCROLLCLIPBOARD, func(p wm.Any) {
		userFunc(wm.HScrollClipboard{Msg: p})
	})
}

// Sent only to dialog windows (not for raw windows).
// Return true to direct the system to set the keyboard focus to the first
// control.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-initdialog
func (me *_EventsWm) WmInitDialog(userFunc func(p wm.InitDialog) bool) {
	me.addMsgRet(co.WM_INITDIALOG, func(p wm.Any) uintptr {
		return util.BoolToUintptr(userFunc(wm.InitDialog{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
func (me *_EventsWm) WmInitMenuPopup(userFunc func(p wm.InitMenuPopup)) {
	me.addMsgZero(co.WM_INITMENUPOPUP, func(p wm.Any) {
		userFunc(wm.InitMenuPopup{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keydown
func (me *_EventsWm) WmKeyDown(userFunc func(p wm.Key)) {
	me.addMsgZero(co.WM_KEYDOWN, func(p wm.Any) {
		userFunc(wm.Key{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keyup
func (me *_EventsWm) WmKeyUp(userFunc func(p wm.Key)) {
	me.addMsgZero(co.WM_KEYUP, func(p wm.Any) {
		userFunc(wm.Key{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-killfocus
func (me *_EventsWm) WmKillFocus(userFunc func(p wm.KillFocus)) {
	me.addMsgZero(co.WM_KILLFOCUS, func(p wm.Any) {
		userFunc(wm.KillFocus{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondblclk
func (me *_EventsWm) WmLButtonDblClk(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_LBUTTONDBLCLK, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
func (me *_EventsWm) WmLButtonDown(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_LBUTTONDOWN, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttonup
func (me *_EventsWm) WmLButtonUp(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_LBUTTONUP, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondblclk
func (me *_EventsWm) WmMButtonDblClk(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_MBUTTONDBLCLK, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondown
func (me *_EventsWm) WmMButtonDown(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_MBUTTONDOWN, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttonup
func (me *_EventsWm) WmMButtonUp(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_MBUTTONUP, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
func (me *_EventsWm) WmMenuChar(userFunc func(p wm.MenuChar) co.MNC) {
	me.addMsgRet(co.WM_MENUCHAR, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.MenuChar{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
func (me *_EventsWm) WmMenuCommand(userFunc func(p wm.Menu)) {
	me.addMsgZero(co.WM_MENUCOMMAND, func(p wm.Any) {
		userFunc(wm.Menu{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
func (me *_EventsWm) WmMenuDrag(userFunc func(p wm.Menu) co.MND) {
	me.addMsgRet(co.WM_MENUDRAG, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.Menu{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
func (me *_EventsWm) WmMenuGetObject(userFunc func(p wm.MenuGetObject) co.MNGO) {
	me.addMsgRet(co.WM_MENUGETOBJECT, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.MenuGetObject{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menurbuttonup
func (me *_EventsWm) WmMenuRButtonUp(userFunc func(p wm.Menu)) {
	me.addMsgZero(co.WM_MENURBUTTONUP, func(p wm.Any) {
		userFunc(wm.Menu{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
func (me *_EventsWm) WmMenuSelect(userFunc func(p wm.MenuSelect)) {
	me.addMsgZero(co.WM_MENUSELECT, func(p wm.Any) {
		userFunc(wm.MenuSelect{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousehover
func (me *_EventsWm) WmMouseHover(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_MOUSEHOVER, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mouseleave
func (me *_EventsWm) WmMouseLeave(userFunc func()) {
	me.addMsgZero(co.WM_MOUSELEAVE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
func (me *_EventsWm) WmMouseMove(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_MOUSEMOVE, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-move
func (me *_EventsWm) WmMove(userFunc func(p wm.Move)) {
	me.addMsgZero(co.WM_MOVE, func(p wm.Any) {
		userFunc(wm.Move{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-moving
func (me *_EventsWm) WmMoving(userFunc func(p wm.Moving)) {
	me.addMsgRet(co.WM_MOVING, func(p wm.Any) uintptr {
		userFunc(wm.Moving{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncactivate
func (me *_EventsWm) WmNcActivate(userFunc func(p wm.NcActivate) bool) {
	me.addMsgRet(co.WM_NCACTIVATE, func(p wm.Any) uintptr {
		return util.BoolToUintptr(userFunc(wm.NcActivate{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
func (me *_EventsWm) WmNcCalcSize(userFunc func(p wm.NcCalcSize) co.WVR) {
	me.addMsgRet(co.WM_NCCALCSIZE, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.NcCalcSize{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccreate
func (me *_EventsWm) WmNcCreate(userFunc func(p wm.Create) bool) {
	me.addMsgRet(co.WM_NCCREATE, func(p wm.Any) uintptr {
		return util.BoolToUintptr(userFunc(wm.Create{Msg: p}))
	})
}

// ⚠️ By handling this message, you'll overwrite the default behavior in WindowMain.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncdestroy
func (me *_EventsWm) WmNcDestroy(userFunc func()) {
	me.addMsgZero(co.WM_NCDESTROY, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
func (me *_EventsWm) WmNcHitTest(userFunc func(wm.NcHitTest) co.HT) {
	me.addMsgRet(co.WM_NCHITTEST, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.NcHitTest{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondblclk
func (me *_EventsWm) WmNcLButtonDblClk(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCLBUTTONDBLCLK, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondown
func (me *_EventsWm) WmNcLButtonDown(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCLBUTTONDOWN, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttonup
func (me *_EventsWm) WmNcLButtonUp(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCLBUTTONUP, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondblclk
func (me *_EventsWm) WmNcMButtonDblClk(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCMBUTTONDBLCLK, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondown
func (me *_EventsWm) WmNcMButtonDown(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCMBUTTONDOWN, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttonup
func (me *_EventsWm) WmNcMButtonUp(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCMBUTTONUP, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousehover
func (me *_EventsWm) WmNcMouseHover(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCMOUSEHOVER, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmouseleave
func (me *_EventsWm) WmNcMouseLeave(userFunc func()) {
	me.addMsgZero(co.WM_NCMOUSELEAVE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousemove
func (me *_EventsWm) WmNcMouseMove(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCMOUSEMOVE, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// ⚠️ By handling this message, you'll overwrite the default behavior in WindowControl.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-ncpaint
func (me *_EventsWm) WmNcPaint(userFunc func(p wm.NcPaint)) {
	me.addMsgZero(co.WM_NCPAINT, func(p wm.Any) {
		userFunc(wm.NcPaint{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondblclk
func (me *_EventsWm) WmNcRButtonDblClk(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCRBUTTONDBLCLK, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondown
func (me *_EventsWm) WmNcRButtonDown(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCRBUTTONDOWN, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttonup
func (me *_EventsWm) WmNcRButtonUp(userFunc func(p wm.NcMouse)) {
	me.addMsgZero(co.WM_NCRBUTTONUP, func(p wm.Any) {
		userFunc(wm.NcMouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondblclk
func (me *_EventsWm) WmNcXButtonDblClk(userFunc func(p wm.NcMouseX)) {
	me.addMsgRet(co.WM_NCXBUTTONDBLCLK, func(p wm.Any) uintptr {
		userFunc(wm.NcMouseX{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondown
func (me *_EventsWm) WmNcXButtonDown(userFunc func(p wm.NcMouseX)) {
	me.addMsgRet(co.WM_NCXBUTTONDOWN, func(p wm.Any) uintptr {
		userFunc(wm.NcMouseX{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttonup
func (me *_EventsWm) WmNcXButtonUp(userFunc func(p wm.NcMouseX)) {
	me.addMsgRet(co.WM_NCXBUTTONUP, func(p wm.Any) uintptr {
		userFunc(wm.NcMouseX{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-nextmenu
func (me *_EventsWm) WmNextMenu(userFunc func(p wm.NextMenu)) {
	me.addMsgZero(co.WM_NEXTMENU, func(p wm.Any) {
		userFunc(wm.NextMenu{Msg: p})
	})
}

// Note that you must call BeginPaint/EndPaint, or the window may get stuck.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-paint
func (me *_EventsWm) WmPaint(userFunc func()) {
	me.addMsgZero(co.WM_PAINT, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-paintclipboard
func (me *_EventsWm) WmPaintClipboard(userFunc func(wm.PaintClipboard)) {
	me.addMsgZero(co.WM_PAINTCLIPBOARD, func(p wm.Any) {
		userFunc(wm.PaintClipboard{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
func (me *_EventsWm) WmPowerBroadcast(userFunc func(p wm.PowerBroadcast)) {
	me.addMsgRet(co.WM_POWERBROADCAST, func(p wm.Any) uintptr {
		userFunc(wm.PowerBroadcast{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/gdi/wm-print
func (me *_EventsWm) WmPrint(userFunc func(p wm.Print)) {
	me.addMsgZero(co.WM_PRINT, func(p wm.Any) {
		userFunc(wm.Print{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondblclk
func (me *_EventsWm) WmRButtonDblClk(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_RBUTTONDBLCLK, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondown
func (me *_EventsWm) WmRButtonDown(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_RBUTTONDOWN, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttonup
func (me *_EventsWm) WmRButtonUp(userFunc func(p wm.Mouse)) {
	me.addMsgZero(co.WM_RBUTTONUP, func(p wm.Any) {
		userFunc(wm.Mouse{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-renderallformats
func (me *_EventsWm) WmRenderAllFormats(userFunc func()) {
	me.addMsgZero(co.WM_RENDERALLFORMATS, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-renderformat
func (me *_EventsWm) WmRenderFormat(userFunc func(p wm.RenderFormat)) {
	me.addMsgZero(co.WM_RENDERFORMAT, func(p wm.Any) {
		userFunc(wm.RenderFormat{Msg: p})
	})
}

// ⚠️ By handling this message, you'll overwrite the default behavior in WindowMain.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-setfocus
func (me *_EventsWm) WmSetFocus(userFunc func(p wm.SetFocus)) {
	me.addMsgZero(co.WM_SETFOCUS, func(p wm.Any) {
		userFunc(wm.SetFocus{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-setfont
func (me *_EventsWm) WmSetFont(userFunc func(p wm.SetFont)) {
	me.addMsgZero(co.WM_SETFONT, func(p wm.Any) {
		userFunc(wm.SetFont{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
func (me *_EventsWm) WmSetIcon(userFunc func(p wm.SetIcon) win.HICON) {
	me.addMsgRet(co.WM_SETICON, func(p wm.Any) uintptr {
		return uintptr(userFunc(wm.SetIcon{Msg: p}))
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-size
func (me *_EventsWm) WmSize(userFunc func(p wm.Size)) {
	me.addMsgZero(co.WM_SIZE, func(p wm.Any) {
		userFunc(wm.Size{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-sizeclipboard
func (me *_EventsWm) WmSizeClipboard(userFunc func(p wm.SizeClipboard)) {
	me.addMsgZero(co.WM_SIZECLIPBOARD, func(p wm.Any) {
		userFunc(wm.SizeClipboard{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syschar
func (me *_EventsWm) WmSysChar(userFunc func(p wm.Char)) {
	me.addMsgZero(co.WM_SYSCHAR, func(p wm.Any) {
		userFunc(wm.Char{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
func (me *_EventsWm) WmSysCommand(userFunc func(p wm.SysCommand)) {
	me.addMsgZero(co.WM_SYSCOMMAND, func(p wm.Any) {
		userFunc(wm.SysCommand{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-sysdeadchar
func (me *_EventsWm) WmSysDeadChar(userFunc func(p wm.Char)) {
	me.addMsgZero(co.WM_SYSDEADCHAR, func(p wm.Any) {
		userFunc(wm.Char{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeydown
func (me *_EventsWm) WmSysKeyDown(userFunc func(p wm.Key)) {
	me.addMsgZero(co.WM_SYSKEYDOWN, func(p wm.Any) {
		userFunc(wm.Key{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeyup
func (me *_EventsWm) WmSysKeyUp(userFunc func(p wm.Key)) {
	me.addMsgZero(co.WM_SYSKEYUP, func(p wm.Any) {
		userFunc(wm.Key{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/sysinfo/wm-timechange
func (me *_EventsWm) WmTimeChange(userFunc func()) {
	me.addMsgZero(co.WM_TIMECHANGE, func(_ wm.Any) {
		userFunc()
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/menurc/wm-uninitmenupopup
func (me *_EventsWm) WmUnInitMenuPopup(userFunc func(p wm.UnInitMenuPopup)) {
	me.addMsgZero(co.WM_UNINITMENUPOPUP, func(p wm.Any) {
		userFunc(wm.UnInitMenuPopup{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/wm-vscroll
func (me *_EventsWm) WmVScroll(userFunc func(p wm.VScroll)) {
	me.addMsgZero(co.WM_VSCROLL, func(p wm.Any) {
		userFunc(wm.VScroll{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-vscrollclipboard
func (me *_EventsWm) WmVScrollClipboard(userFunc func(p wm.VScrollClipboard)) {
	me.addMsgZero(co.WM_VSCROLLCLIPBOARD, func(p wm.Any) {
		userFunc(wm.VScrollClipboard{Msg: p})
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondblclk
func (me *_EventsWm) WmXButtonDblClk(userFunc func(p wm.Mouse)) {
	me.addMsgRet(co.WM_XBUTTONDBLCLK, func(p wm.Any) uintptr {
		userFunc(wm.Mouse{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondown
func (me *_EventsWm) WmXButtonDown(userFunc func(p wm.Mouse)) {
	me.addMsgRet(co.WM_XBUTTONDOWN, func(p wm.Any) uintptr {
		userFunc(wm.Mouse{Msg: p})
		return 1
	})
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttonup
func (me *_EventsWm) WmXButtonUp(userFunc func(p wm.Mouse)) {
	me.addMsgRet(co.WM_XBUTTONUP, func(p wm.Any) uintptr {
		userFunc(wm.Mouse{Msg: p})
		return 1
	})
}
