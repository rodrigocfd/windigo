//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
)

type (
	_StorageMsg struct { // ordinary WM messages
		id  co.WM
		fun func(p Wm) uintptr
	}
	_StorageCmd struct { // WM_COMMAND
		cmdId     uint16
		notifCode co.CMD
		fun       func()
	}
	_StorageNfy struct { // WM_NOTIFY
		idFrom uint16
		code   co.NM
		fun    func(p unsafe.Pointer) uintptr
	}
	_StorageTmr struct { // WM_TIMER
		timerId uintptr
		fun     func()
	}
	_WNDTY uint8 // Tells if the window is raw or dialog.
)

const (
	_WNDTY_DLG _WNDTY = 0x01 // A dialog window, loaded from resource.
	_WNDTY_RAW _WNDTY = 0x02 // An ordinary window created with CreateWindowEx.
)

// Exposes events for window [messages].
//
// You cannot create this object directly, it will be created automatically
// by the owning window.
//
// [messages]: https://learn.microsoft.com/en-us/windows/win32/winmsg/about-messages-and-message-queues
type EventsWindow struct {
	defProcVal uintptr // 0 for ordinary windows, TRUE for dialogs

	// We use simple arrays instead of maps, because the closures are never too
	// many, therefore a simple linear search is more efficient.

	wmCreate     func(p WmCreate) int      // WM_CREATE message; deleted after first processing
	wmInitDialog func(p WmInitDialog) bool // WM_INITDIALOG message; deleted after first processing
	msgs         []_StorageMsg             // ordinary WM messages
	cmds         []_StorageCmd             // WM_COMMAND
	nfys         []_StorageNfy             // WM_NOTIFY
	tmrs         []_StorageTmr             // WM_TIMER
}

// Constructor.
func newEventsWindow(wndTy _WNDTY) EventsWindow {
	defProcVal := 0 // for raw windows
	if wndTy == _WNDTY_DLG {
		defProcVal = 1 // TRUE; for dialogs
	}

	return EventsWindow{
		defProcVal:   uintptr(defProcVal),
		wmCreate:     nil,
		wmInitDialog: nil,
		msgs:         make([]_StorageMsg, 0),
		cmds:         make([]_StorageCmd, 0),
		nfys:         make([]_StorageNfy, 0),
		tmrs:         make([]_StorageTmr, 0),
	}
}

// To be called after the first WM_CREATE/INITDIALOG processing. Releases the
// memory in all these closures, which are never called again.
func (me *EventsWindow) removeWmCreateInitdialog() {
	me.wmCreate = nil
	me.wmInitDialog = nil
}

// Releases the memory of all closures.
func (me *EventsWindow) clear() {
	me.removeWmCreateInitdialog()
	me.msgs = nil
	me.cmds = nil
	me.nfys = nil
	me.tmrs = nil
}

func (me *EventsWindow) hasMessage() bool {
	return me.wmCreate != nil ||
		me.wmInitDialog != nil ||
		len(me.msgs) > 0 ||
		len(me.cmds) > 0 ||
		len(me.nfys) > 0 ||
		len(me.tmrs) > 0
}

// For user events. When the user adds a message handler, it will overwrite a
// previously added one. It would be too costly to search and remove the
// previous one, so we just keep them all, and run the last one.
func (me *EventsWindow) processLast(p Wm) (userRet uintptr, wasHandled bool) {
	switch p.Msg {
	case co.WM_CREATE:
		if me.wmCreate != nil {
			return uintptr(me.wmCreate(WmCreate{p})), true // handled, stop here
		}
	case co.WM_INITDIALOG:
		if me.wmInitDialog != nil {
			return utl.BoolToUintptr(me.wmInitDialog(WmInitDialog{p})), true // handled, stop here
		}
	case co.WM_COMMAND:
		cmdId := p.WParam.LoWord()
		notifCode := co.CMD(p.WParam.HiWord())
		for i := len(me.cmds) - 1; i >= 0; i-- {
			if me.cmds[i].cmdId == cmdId && me.cmds[i].notifCode == notifCode {
				me.cmds[i].fun()
				return me.defProcVal, true // handled, stop here
			}
		}
	case co.WM_NOTIFY:
		pHdr := unsafe.Pointer(p.LParam)
		hdr := (*win.NMHDR)(pHdr)
		for i := len(me.nfys) - 1; i >= 0; i-- {
			if me.nfys[i].idFrom == uint16(hdr.IdFrom) && me.nfys[i].code == co.NM(hdr.Code) {
				return me.nfys[i].fun(pHdr), true // handled, stop here
			}
		}
	case co.WM_TIMER:
		for i := len(me.tmrs) - 1; i >= 0; i-- {
			if me.tmrs[i].timerId == uintptr(p.WParam) {
				me.tmrs[i].fun()
				return me.defProcVal, true // handled, stop here
			}
		}
	default:
		for i := len(me.msgs) - 1; i >= 0; i-- {
			if me.msgs[i].id == p.Msg {
				return me.msgs[i].fun(p), true // handled, stop here
			}
		}
	}
	return 0, false // no message found
}

// [WM_CREATE] message handler.
//
// Return 0 to continue window creation, or -1 to abort it.
//
// Sent only to windows created with [CreateWindowEx]; dialog windows will
// receive [EventsWindow.WmInitDialog] instead.
//
// Example:
//
//	var wnd ui.Parent // initialized somewhere
//
//	wnd.On().WmCreate(func(_ ui.WmCreate) bool {
//		return 0
//	}
//
// [WM_CREATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-create
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (me *EventsWindow) WmCreate(fun func(p WmCreate) int) {
	me.wmCreate = fun
}

// [WM_INITDIALOG] message handler.
//
// Return true to direct the system to set the keyboard focus to the first
// control.
//
// Sent only to dialog windows; those created with [CreateWindowEx] will receive
// [EventsWindow.WmCreate] instead.
//
// Example:
//
//	var wnd ui.Parent // initialized somewhere
//
//	wnd.On().WmInitDialog(func(_ ui.WmInitDialog) bool {
//		return true
//	}
//
// [WM_INITDIALOG]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-initdialog
// [CreateWindowEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw
func (me *EventsWindow) WmInitDialog(fun func(p WmInitDialog) bool) {
	me.wmInitDialog = fun
}

// Generic [message handler].
//
// Avoid this method, prefer the specific message handlers.
//
// [message handler]: https://learn.microsoft.com/en-us/windows/win32/learnwin32/window-messages
func (me *EventsWindow) Wm(id co.WM, fun func(p Wm) uintptr) {
	me.msgs = append(me.msgs, _StorageMsg{id, fun})
}

// Generic [WM_COMMAND] handler.
//
// Avoid this method, prefer the specific command notification handlers.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *EventsWindow) WmCommand(cmdId uint16, notifCode co.CMD, fun func()) {
	me.cmds = append(me.cmds, _StorageCmd{cmdId, notifCode, fun})
}

// [WM_COMMAND] handler for both accelerator and menu events. Ideal for IDs
// shared between accelerator keys and menu items.
//
// [WM_COMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *EventsWindow) WmCommandAccelMenu(cmdId uint16, fun func()) {
	me.WmCommand(cmdId, co.CMD_MENU, fun)
	me.WmCommand(cmdId, co.CMD_ACCELERATOR, fun)
}

// Generic [WM_NOTIFY] handler.
//
// Avoid this method, prefer the specific notification handlers.
//
// [WM_NOTIFY]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-notify
func (me *EventsWindow) WmNotify(idFrom uint16, code co.NM, fun func(p unsafe.Pointer) uintptr) {
	me.nfys = append(me.nfys, _StorageNfy{idFrom, code, fun})
}

// [WM_TIMER] message handler.
//
// [WM_TIMER]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-timer
func (me *EventsWindow) WmTimer(timerId uintptr, fun func()) {
	me.tmrs = append(me.tmrs, _StorageTmr{timerId, fun})
}

// [WM_ACTIVATE] message handler.
//
// [WM_ACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-activate
func (me *EventsWindow) WmActivate(fun func(p WmActivate)) {
	me.Wm(co.WM_ACTIVATE, func(p Wm) uintptr {
		fun(WmActivate{p})
		return me.defProcVal
	})
}

// [WM_ACTIVATEAPP] message handler.
//
// [WM_ACTIVATEAPP]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-activateapp
func (me *EventsWindow) WmActivateApp(fun func(p WmActivateApp)) {
	me.Wm(co.WM_ACTIVATEAPP, func(p Wm) uintptr {
		fun(WmActivateApp{p})
		return me.defProcVal
	})
}

// [WM_APPCOMMAND] message handler.
//
// [WM_APPCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
func (me *EventsWindow) WmAppCommand(fun func(p WmAppCommand)) {
	me.Wm(co.WM_APPCOMMAND, func(p Wm) uintptr {
		fun(WmAppCommand{p})
		return 1
	})
}

// [WM_ASKCBFORMATNAME] message handler.
//
// [WM_ASKCBFORMATNAME]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-askcbformatname
func (me *EventsWindow) WmAskCbFormatName(fun func(p WmAskCbFormatName)) {
	me.Wm(co.WM_ASKCBFORMATNAME, func(p Wm) uintptr {
		fun(WmAskCbFormatName{p})
		return me.defProcVal
	})
}

// [WM_CANCELMODE] message handler.
//
// [WM_CANCELMODE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-cancelmode
func (me *EventsWindow) WmCancelMode(fun func()) {
	me.Wm(co.WM_CANCELMODE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_CAPTURECHANGED] message handler.
//
// [WM_CAPTURECHANGED]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-capturechanged
func (me *EventsWindow) WmCaptureChanged(fun func(p WmCaptureChanged)) {
	me.Wm(co.WM_CAPTURECHANGED, func(p Wm) uintptr {
		fun(WmCaptureChanged{p})
		return me.defProcVal
	})
}

// [WM_CHANGECBCHAIN] message handler.
//
// [WM_CHANGECBCHAIN]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-changecbchain
func (me *EventsWindow) WmChangeCbChain(fun func(p WmChangeCbChain)) {
	me.Wm(co.WM_CHANGECBCHAIN, func(p Wm) uintptr {
		fun(WmChangeCbChain{p})
		return me.defProcVal
	})
}

// [WM_CHAR] message handler.
//
// [WM_CHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (me *EventsWindow) WmChar(fun func(p WmChar)) {
	me.Wm(co.WM_CHAR, func(p Wm) uintptr {
		fun(WmChar{p})
		return me.defProcVal
	})
}

// [WM_CHARTOITEM] message handler.
//
// [WM_CHARTOITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-chartoitem
func (me *EventsWindow) WmCharToItem(fun func(p WmCharToItem) int) {
	me.Wm(co.WM_CHARTOITEM, func(p Wm) uintptr {
		return uintptr(fun(WmCharToItem{p}))
	})
}

// [WM_CHILDACTIVATE] message handler.
//
// [WM_CHILDACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-childactivate
func (me *EventsWindow) WmChildActivate(fun func()) {
	me.Wm(co.WM_CHILDACTIVATE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_CLIPBOARDUPDATE] message handler.
//
// [WM_CLIPBOARDUPDATE]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-clipboardupdate
func (me *EventsWindow) WmClipboardUpdate(fun func()) {
	me.Wm(co.WM_CLIPBOARDUPDATE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_CLOSE] message handler.
//
// ⚠️ By handling this message, you'll overwrite the default behavior in
// WindowMain and WindowModal:
//   - dialog WindowMain: calls [DestroyWindow];
//   - dialog WindowModal: calls [EndDialog];
//   - raw WindowModal: re-enables parent window and calls [DestroyWindow].
//
// [WM_CLOSE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-close
// [DestroyWindow]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-destroywindow
// [EndDialog]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-enddialog
func (me *EventsWindow) WmClose(fun func()) {
	me.Wm(co.WM_CLOSE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_COMPAREITEM] message handler.
//
// [WM_COMPAREITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-compareitem
func (me *EventsWindow) WmCompareItem(fun func(p WmCompareItem) int) {
	me.Wm(co.WM_COMPAREITEM, func(p Wm) uintptr {
		return uintptr(fun(WmCompareItem{p}))
	})
}

// [WM_CONTEXTMENU] message handler.
//
// [WM_CONTEXTMENU]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-contextmenu
func (me *EventsWindow) WmContextMenu(fun func(p WmContextMenu)) {
	me.Wm(co.WM_CONTEXTMENU, func(p Wm) uintptr {
		fun(WmContextMenu{p})
		return me.defProcVal
	})
}

// [WM_COPYDATA] message handler.
//
// [WM_COPYDATA]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-copydata
func (me *EventsWindow) WmCopyData(fun func(p WmCopyData) bool) {
	me.Wm(co.WM_COPYDATA, func(p Wm) uintptr {
		return utl.BoolToUintptr(fun(WmCopyData{p}))
	})
}

// [WM_CTLCOLORBTN] message handler.
//
// [WM_CTLCOLORBTN]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorbtn
func (me *EventsWindow) WmCtlColorBtn(fun func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORBTN, func(p Wm) uintptr {
		return uintptr(fun(WmCtlColor{p}))
	})
}

// [WM_CTLCOLORDLG] message handler.
//
// [WM_CTLCOLORDLG]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-ctlcolordlg
func (me *EventsWindow) WmCtlColorDlg(fun func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORDLG, func(p Wm) uintptr {
		return uintptr(fun(WmCtlColor{p}))
	})
}

// [WM_CTLCOLOREDIT] message handler.
//
// [WM_CTLCOLOREDIT]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcoloredit
func (me *EventsWindow) WmCtlColorEdit(fun func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLOREDIT, func(p Wm) uintptr {
		return uintptr(fun(WmCtlColor{p}))
	})
}

// [WM_CTLCOLORLISTBOX] message handler.
//
// [WM_CTLCOLORLISTBOX]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorlistbox
func (me *EventsWindow) WmCtlColorListBox(fun func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORLISTBOX, func(p Wm) uintptr {
		return uintptr(fun(WmCtlColor{p}))
	})
}

// [WM_CTLCOLORSCROLLBAR] message handler.
//
// [WM_CTLCOLORSCROLLBAR]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorscrollbar
func (me *EventsWindow) WmCtlColorScrollBar(fun func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORSCROLLBAR, func(p Wm) uintptr {
		return uintptr(fun(WmCtlColor{p}))
	})
}

// [WM_CTLCOLORSTATIC] message handler.
//
// [WM_CTLCOLORSTATIC]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
func (me *EventsWindow) WmCtlColorStatic(fun func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORSTATIC, func(p Wm) uintptr {
		return uintptr(fun(WmCtlColor{p}))
	})
}

// [WM_DEADCHAR] message handler.
//
// [WM_DEADCHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
func (me *EventsWindow) WmDeadChar(fun func(p WmChar)) {
	me.Wm(co.WM_DEADCHAR, func(p Wm) uintptr {
		fun(WmChar{p})
		return me.defProcVal
	})
}

// [WM_DELETEITEM] message handler.
//
// [WM_DELETEITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-deleteitem
func (me *EventsWindow) WmDeleteItem(fun func(p WmDeleteItem)) {
	me.Wm(co.WM_DELETEITEM, func(p Wm) uintptr {
		fun(WmDeleteItem{p})
		return 1
	})
}

// [WM_DESTROY] message handler.
//
// [WM_DESTROY]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
func (me *EventsWindow) WmDestroy(fun func()) {
	me.Wm(co.WM_DESTROY, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_DESTROYCLIPBOARD] message handler.
//
// [WM_DESTROYCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-destroyclipboard
func (me *EventsWindow) WmDestroyClipboard(fun func()) {
	me.Wm(co.WM_DESTROYCLIPBOARD, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_DEVICECHANGE] message handler.
//
// [WM_DEVICECHANGE]: https://learn.microsoft.com/en-us/windows/win32/devio/wm-devicechange
func (me *EventsWindow) WmDeviceChange(fun func(WmDeviceChange) co.BROADCAST_QUERY) {
	me.Wm(co.WM_DEVICECHANGE, func(p Wm) uintptr {
		return uintptr(fun(WmDeviceChange{p}))
	})
}

// [WM_DEVMODECHANGE] message handler.
//
// [WM_DEVMODECHANGE]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-devmodechange
func (me *EventsWindow) WmDevModeChange(fun func(WmDevModeChange)) {
	me.Wm(co.WM_DEVMODECHANGE, func(p Wm) uintptr {
		fun(WmDevModeChange{p})
		return me.defProcVal
	})
}

// [WM_DISPLAYCHANGE] message handler.
//
// [WM_DISPLAYCHANGE]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-displaychange
func (me *EventsWindow) WmDisplayChange(fun func(p WmDisplayChange)) {
	me.Wm(co.WM_DISPLAYCHANGE, func(p Wm) uintptr {
		fun(WmDisplayChange{p})
		return me.defProcVal
	})
}

// [WM_DRAWCLIPBOARD] message handler.
//
// [WM_DRAWCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-drawclipboard
func (me *EventsWindow) WmDrawClipboard(fun func()) {
	me.Wm(co.WM_DRAWCLIPBOARD, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_DRAWITEM] message handler.
//
// [WM_DRAWITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-drawitem
func (me *EventsWindow) WmDrawItem(fun func(p WmDrawItem)) {
	me.Wm(co.WM_DRAWITEM, func(p Wm) uintptr {
		fun(WmDrawItem{p})
		return 1
	})
}

// [WM_DROPFILES] message handler.
//
// [WM_DROPFILES]: https://learn.microsoft.com/en-us/windows/win32/shell/wm-dropfiles
func (me *EventsWindow) WmDropFiles(fun func(p WmDropFiles)) {
	me.Wm(co.WM_DROPFILES, func(p Wm) uintptr {
		fun(WmDropFiles{p})
		return me.defProcVal
	})
}

// [WM_DWMCOLORIZATIONCOLORCHANGED] message handler.
//
// [WM_DWMCOLORIZATIONCOLORCHANGED]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmcolorizationcolorchanged
func (me *EventsWindow) WmDwmColorizationColorChanged(fun func(p WmDwmColorizationColorChanged)) {
	me.Wm(co.WM_DWMCOLORIZATIONCOLORCHANGED, func(p Wm) uintptr {
		fun(WmDwmColorizationColorChanged{p})
		return me.defProcVal
	})
}

// [WM_DWMCOMPOSITIONCHANGED] message handler.
//
// [WM_DWMCOMPOSITIONCHANGED]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmcompositionchanged
func (me *EventsWindow) WmDwmCompositionChanged(fun func()) {
	me.Wm(co.WM_DWMCOMPOSITIONCHANGED, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_DWMNCRENDERINGCHANGED] message handler.
//
// [WM_DWMNCRENDERINGCHANGED]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmncrenderingchanged
func (me *EventsWindow) WmDwmNcRenderingChanged(fun func(p WmDwmNcRenderingChanged)) {
	me.Wm(co.WM_DWMNCRENDERINGCHANGED, func(p Wm) uintptr {
		fun(WmDwmNcRenderingChanged{p})
		return me.defProcVal
	})
}

// [WM_DWMSENDICONICLIVEPREVIEWBITMAP] message handler.
//
// [WM_DWMSENDICONICLIVEPREVIEWBITMAP]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmsendiconiclivepreviewbitmap
func (me *EventsWindow) WmDwmSendIconicLivePreviewBitmap(fun func()) {
	me.Wm(co.WM_DWMSENDICONICLIVEPREVIEWBITMAP, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_DWMSENDICONICTHUMBNAIL] message handler.
//
// [WM_DWMSENDICONICTHUMBNAIL]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmsendiconicthumbnail
func (me *EventsWindow) WmDwmSendIconicThumbnail(fun func(p WmDwmSendIconicThumbnail)) {
	me.Wm(co.WM_DWMSENDICONICTHUMBNAIL, func(p Wm) uintptr {
		fun(WmDwmSendIconicThumbnail{p})
		return me.defProcVal
	})
}

// [WM_DWMWINDOWMAXIMIZEDCHANGE] message handler.
//
// [WM_DWMWINDOWMAXIMIZEDCHANGE]: https://learn.microsoft.com/en-us/windows/win32/dwm/wm-dwmwindowmaximizedchange
func (me *EventsWindow) WmDwmWindowMaximizedChange(fun func(p WmDwmWindowMaximizedChange)) {
	me.Wm(co.WM_DWMWINDOWMAXIMIZEDCHANGE, func(p Wm) uintptr {
		fun(WmDwmWindowMaximizedChange{p})
		return me.defProcVal
	})
}

// [WM_ENABLE] message handler.
//
// [WM_ENABLE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-enable
func (me *EventsWindow) WmEnable(fun func(p WmEnable)) {
	me.Wm(co.WM_ENABLE, func(p Wm) uintptr {
		fun(WmEnable{p})
		return me.defProcVal
	})
}

// [WM_ENDSESSION] message handler.
//
// [WM_ENDSESSION]: https://learn.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
func (me *EventsWindow) WmEndSession(fun func(p WmEndSession)) {
	me.Wm(co.WM_ENDSESSION, func(p Wm) uintptr {
		fun(WmEndSession{p})
		return me.defProcVal
	})
}

// [WM_ENTERIDLE] message handler.
//
// [WM_ENTERIDLE]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
func (me *EventsWindow) WmEnterIdle(fun func(p WmEnterIdle)) {
	me.Wm(co.WM_ENTERIDLE, func(p Wm) uintptr {
		fun(WmEnterIdle{p})
		return me.defProcVal
	})
}

// [WM_ENTERMENULOOP] message handler.
//
// [WM_ENTERMENULOOP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-entermenuloop
func (me *EventsWindow) WmEnterMenuLoop(fun func(WmEnterMenuLoop)) {
	me.Wm(co.WM_ENTERMENULOOP, func(p Wm) uintptr {
		fun(WmEnterMenuLoop{p})
		return me.defProcVal
	})
}

// [WM_ENTERSIZEMOVE] message handler.
//
// [WM_ENTERSIZEMOVE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-entersizemove
func (me *EventsWindow) WmEnterSizeMove(fun func()) {
	me.Wm(co.WM_ENTERSIZEMOVE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_ERASEBKGND] message handler.
//
// [WM_ERASEBKGND]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-erasebkgnd
func (me *EventsWindow) WmEraseBkgnd(fun func(WmEraseBkgnd) int) {
	me.Wm(co.WM_ERASEBKGND, func(p Wm) uintptr {
		return uintptr(fun(WmEraseBkgnd{p}))
	})
}

// [WM_EXITMENULOOP] message handler.
//
// [WM_EXITMENULOOP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-exitmenuloop
func (me *EventsWindow) WmExitMenuLoop(fun func(WmExitMenuLoop)) {
	me.Wm(co.WM_EXITMENULOOP, func(p Wm) uintptr {
		fun(WmExitMenuLoop{p})
		return me.defProcVal
	})
}

// [WM_EXITSIZEMOVE] message handler.
//
// [WM_EXITSIZEMOVE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-exitsizemove
func (me *EventsWindow) WmExitSizeMove(fun func()) {
	me.Wm(co.WM_EXITSIZEMOVE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_FONTCHANGE] message handler.
//
// [WM_FONTCHANGE]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-fontchange
func (me *EventsWindow) WmFontChange(fun func()) {
	me.Wm(co.WM_FONTCHANGE, func(p Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_GETDLGCODE] message handler.
//
// [WM_GETDLGCODE]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
func (me *EventsWindow) WmGetDlgCode(fun func(p WmGetDlgCode) co.DLGC) {
	me.Wm(co.WM_GETDLGCODE, func(p Wm) uintptr {
		return uintptr(fun(WmGetDlgCode{p}))
	})
}

// [WM_GETFONT] message handler.
//
// [WM_GETFONT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-getfont
func (me *EventsWindow) WmGetFont(fun func() win.HFONT) {
	me.Wm(co.WM_GETFONT, func(_ Wm) uintptr {
		return uintptr(fun())
	})
}

// [MN_GETHMENU] message handler.
//
// [MN_GETHMENU]: https://learn.microsoft.com/en-us/windows/win32/winmsg/mn-gethmenu
func (me *EventsWindow) WmGetHMenu(fun func() win.HMENU) {
	me.Wm(co.WM_MN_GETHMENU, func(_ Wm) uintptr {
		return uintptr(fun())
	})
}

// [WM_GETICON] message handler.
//
// [WM_GETICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-geticon
func (me *EventsWindow) WmGetIcon(fun func(p WmGetIcon) win.HICON) {
	me.Wm(co.WM_GETICON, func(p Wm) uintptr {
		return uintptr(fun(WmGetIcon{p}))
	})
}

// [WM_GETMINMAXINFO] message handler.
//
// [WM_GETMINMAXINFO]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-getminmaxinfo
func (me *EventsWindow) WmGetMinMaxInfo(fun func(p WmGetMinMaxInfo)) {
	me.Wm(co.WM_GETMINMAXINFO, func(p Wm) uintptr {
		fun(WmGetMinMaxInfo{p})
		return me.defProcVal
	})
}

// [WM_GETTEXT] message handler.
//
// [WM_GETTEXT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-gettext
func (me *EventsWindow) WmGetText(fun func(p WmGetText) uint) {
	me.Wm(co.WM_GETTEXT, func(p Wm) uintptr {
		return uintptr(fun(WmGetText{p}))
	})
}

// [WM_GETTEXTLENGTH] message handler.
//
// [WM_GETTEXTLENGTH]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-gettextlength
func (me *EventsWindow) WmGetTextLength(fun func() uint) {
	me.Wm(co.WM_GETTEXTLENGTH, func(p Wm) uintptr {
		return uintptr(fun())
	})
}

// [WM_GETTITLEBARINFOEX] message handler.
//
// [WM_GETTITLEBARINFOEX]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-gettitlebarinfoex
func (me *EventsWindow) WmGetTitleBarInfoEx(fun func(p WmGetTitleBarInfoEx)) {
	me.Wm(co.WM_GETTITLEBARINFOEX, func(p Wm) uintptr {
		fun(WmGetTitleBarInfoEx{p})
		return me.defProcVal
	})
}

// [WM_HELP] message handler.
//
// [WM_HELP]: https://learn.microsoft.com/en-us/windows/win32/shell/wm-help
func (me *EventsWindow) WmHelp(fun func(p WmHelp)) {
	me.Wm(co.WM_HELP, func(p Wm) uintptr {
		fun(WmHelp{p})
		return 1
	})
}

// [WM_HOTKEY] message handler.
//
// [WM_HOTKEY]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
func (me *EventsWindow) WmHotKey(fun func(p WmHotKey)) {
	me.Wm(co.WM_HOTKEY, func(p Wm) uintptr {
		fun(WmHotKey{p})
		return me.defProcVal
	})
}

// [WM_HSCROLL] message handler.
//
// [WM_HSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-hscroll
func (me *EventsWindow) WmHScroll(fun func(p WmScroll)) {
	me.Wm(co.WM_HSCROLL, func(p Wm) uintptr {
		fun(WmScroll{p})
		return me.defProcVal
	})
}

// [WM_HSCROLLCLIPBOARD] message handler.
//
// [WM_HSCROLLCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-hscrollclipboard
func (me *EventsWindow) WmHScrollClipboard(fun func(p WmScrollClipboard)) {
	me.Wm(co.WM_HSCROLLCLIPBOARD, func(p Wm) uintptr {
		fun(WmScrollClipboard{p})
		return me.defProcVal
	})
}

// [WM_INITMENUPOPUP] message handler.
//
// [WM_INITMENUPOPUP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
func (me *EventsWindow) WmInitMenuPopup(fun func(p WmInitMenuPopup)) {
	me.Wm(co.WM_INITMENUPOPUP, func(p Wm) uintptr {
		fun(WmInitMenuPopup{p})
		return me.defProcVal
	})
}

// [WM_KEYDOWN] message handler.
//
// [WM_KEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-keydown
func (me *EventsWindow) WmKeyDown(fun func(p WmKey)) {
	me.Wm(co.WM_KEYDOWN, func(p Wm) uintptr {
		fun(WmKey{p})
		return me.defProcVal
	})
}

// [WM_KEYUP] message handler.
//
// [WM_KEYUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-keyup
func (me *EventsWindow) WmKeyUp(fun func(p WmKey)) {
	me.Wm(co.WM_KEYUP, func(p Wm) uintptr {
		fun(WmKey{p})
		return me.defProcVal
	})
}

// [WM_KILLFOCUS] message handler.
//
// [WM_KILLFOCUS]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-killfocus
func (me *EventsWindow) WmKillFocus(fun func(p WmKillFocus)) {
	me.Wm(co.WM_KILLFOCUS, func(p Wm) uintptr {
		fun(WmKillFocus{p})
		return me.defProcVal
	})
}

// [WM_LBUTTONDBLCLK] message handler.
//
// [WM_LBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondblclk
func (me *EventsWindow) WmLButtonDblClk(fun func(p WmMouse)) {
	me.Wm(co.WM_LBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_LBUTTONDOWN] message handler.
//
// [WM_LBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
func (me *EventsWindow) WmLButtonDown(fun func(p WmMouse)) {
	me.Wm(co.WM_LBUTTONDOWN, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_LBUTTONUP] message handler.
//
// [WM_LBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttonup
func (me *EventsWindow) WmLButtonUp(fun func(p WmMouse)) {
	me.Wm(co.WM_LBUTTONUP, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_MBUTTONDBLCLK] message handler.
//
// [WM_MBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondblclk
func (me *EventsWindow) WmMButtonDblClk(fun func(p WmMouse)) {
	me.Wm(co.WM_MBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_MBUTTONDOWN] message handler.
//
// [WM_MBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondown
func (me *EventsWindow) WmMButtonDown(fun func(p WmMouse)) {
	me.Wm(co.WM_MBUTTONDOWN, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_MBUTTONUP] message handler.
//
// [WM_MBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttonup
func (me *EventsWindow) WmMButtonUp(fun func(p WmMouse)) {
	me.Wm(co.WM_MBUTTONUP, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_MENUCHAR] message handler.
//
// [WM_MENUCHAR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
func (me *EventsWindow) WmMenuChar(fun func(p WmMenuChar) co.MNC) {
	me.Wm(co.WM_MENUCHAR, func(p Wm) uintptr {
		return uintptr(fun(WmMenuChar{p}))
	})
}

// [WM_MENUCOMMAND] message handler.
//
// [WM_MENUCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
func (me *EventsWindow) WmMenuCommand(fun func(p WmMenu)) {
	me.Wm(co.WM_MENUCOMMAND, func(p Wm) uintptr {
		fun(WmMenu{p})
		return me.defProcVal
	})
}

// [WM_MENUDRAG] message handler.
//
// [WM_MENUDRAG]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
func (me *EventsWindow) WmMenuDrag(fun func(p WmMenu) co.MND) {
	me.Wm(co.WM_MENUDRAG, func(p Wm) uintptr {
		return uintptr(fun(WmMenu{p}))
	})
}

// [WM_MENUGETOBJECT] message handler.
//
// [WM_MENUGETOBJECT]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
func (me *EventsWindow) WmMenuGetObject(fun func(p WmMenuGetObject) co.MNGO) {
	me.Wm(co.WM_MENUGETOBJECT, func(p Wm) uintptr {
		return uintptr(fun(WmMenuGetObject{p}))
	})
}

// [WM_MENURBUTTONUP] message handler.
//
// [WM_MENURBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menurbuttonup
func (me *EventsWindow) WmMenuRButtonUp(fun func(p WmMenu)) {
	me.Wm(co.WM_MENURBUTTONUP, func(p Wm) uintptr {
		fun(WmMenu{p})
		return me.defProcVal
	})
}

// [WM_MENUSELECT] message handler.
//
// [WM_MENUSELECT]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
func (me *EventsWindow) WmMenuSelect(fun func(p WmMenuSelect)) {
	me.Wm(co.WM_MENUSELECT, func(p Wm) uintptr {
		fun(WmMenuSelect{p})
		return me.defProcVal
	})
}

// [WM_MOUSEHOVER] message handler.
//
// [WM_MOUSEHOVER]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mousehover
func (me *EventsWindow) WmMouseHover(fun func(p WmMouse)) {
	me.Wm(co.WM_MOUSEHOVER, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_MOUSELEAVE] message handler.
//
// [WM_MOUSELEAVE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mouseleave
func (me *EventsWindow) WmMouseLeave(fun func()) {
	me.Wm(co.WM_MOUSELEAVE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_MOUSEMOVE] message handler.
//
// [WM_MOUSEMOVE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
func (me *EventsWindow) WmMouseMove(fun func(p WmMouse)) {
	me.Wm(co.WM_MOUSEMOVE, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_MOVE] message handler.
//
// [WM_MOVE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-move
func (me *EventsWindow) WmMove(fun func(p WmMove)) {
	me.Wm(co.WM_MOVE, func(p Wm) uintptr {
		fun(WmMove{p})
		return me.defProcVal
	})
}

// [WM_MOVING] message handler.
//
// [WM_MOVING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-moving
func (me *EventsWindow) WmMoving(fun func(p WmMoving)) {
	me.Wm(co.WM_MOVING, func(p Wm) uintptr {
		fun(WmMoving{p})
		return 1
	})
}

// [WM_NCACTIVATE] message handler.
//
// [WM_NCACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-ncactivate
func (me *EventsWindow) WmNcActivate(fun func(p WmNcActivate) bool) {
	me.Wm(co.WM_NCACTIVATE, func(p Wm) uintptr {
		return utl.BoolToUintptr(fun(WmNcActivate{p}))
	})
}

// [WM_NCCALCSIZE] message handler.
//
// [WM_NCCALCSIZE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
func (me *EventsWindow) WmNcCalcSize(fun func(p WmNcCalcSize) co.WVR) {
	me.Wm(co.WM_NCCALCSIZE, func(p Wm) uintptr {
		return uintptr(fun(WmNcCalcSize{p}))
	})
}

// [WM_NCCREATE] message handler.
//
// [WM_NCCREATE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-nccreate
func (me *EventsWindow) WmNcCreate(fun func(p WmCreate) bool) {
	me.Wm(co.WM_NCCREATE, func(p Wm) uintptr {
		return utl.BoolToUintptr(fun(WmCreate{p}))
	})
}

// [WM_NCDESTROY] message handler.
//
// ⚠️ By handling this message, you'll overwrite the default behavior in raw
// and dialog WindowMain, in both cases calling [PostQuitMessage].
//
// [WM_NCDESTROY]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-ncdestroy
// [PostQuitMessage]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postquitmessage
func (me *EventsWindow) WmNcDestroy(fun func()) {
	me.Wm(co.WM_NCDESTROY, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_NCHITTEST] message handler.
//
// [WM_NCHITTEST]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
func (me *EventsWindow) WmNcHitTest(fun func(WmNcHitTest) co.HT) {
	me.Wm(co.WM_NCHITTEST, func(p Wm) uintptr {
		return uintptr(fun(WmNcHitTest{p}))
	})
}

// [WM_NCLBUTTONDBLCLK] message handler.
//
// [WM_NCLBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondblclk
func (me *EventsWindow) WmNcLButtonDblClk(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCLBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCLBUTTONDOWN] message handler.
//
// [WM_NCLBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondown
func (me *EventsWindow) WmNcLButtonDown(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCLBUTTONDOWN, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCLBUTTONUP] message handler.
//
// [WM_NCLBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttonup
func (me *EventsWindow) WmNcLButtonUp(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCLBUTTONUP, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCMBUTTONDBLCLK] message handler.
//
// [WM_NCMBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondblclk
func (me *EventsWindow) WmNcMButtonDblClk(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCMBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCMBUTTONDOWN] message handler.
//
// [WM_NCMBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondown
func (me *EventsWindow) WmNcMButtonDown(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCMBUTTONDOWN, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCMBUTTONUP] message handler.
//
// [WM_NCMBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttonup
func (me *EventsWindow) WmNcMButtonUp(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCMBUTTONUP, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCMOUSEHOVER] message handler.
//
// [WM_NCMOUSEHOVER]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousehover
func (me *EventsWindow) WmNcMouseHover(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCMOUSEHOVER, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCMOUSELEAVE] message handler.
//
// [WM_NCMOUSELEAVE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmouseleave
func (me *EventsWindow) WmNcMouseLeave(fun func()) {
	me.Wm(co.WM_NCMOUSELEAVE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_NCMOUSEMOVE] message handler.
//
// [WM_NCMOUSEMOVE]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousemove
func (me *EventsWindow) WmNcMouseMove(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCMOUSEMOVE, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCPAINT] message handler.
//
// [WM_NCPAINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-ncpaint
func (me *EventsWindow) WmNcPaint(fun func(p WmNcPaint)) {
	me.Wm(co.WM_NCPAINT, func(p Wm) uintptr {
		fun(WmNcPaint{p})
		return me.defProcVal
	})
}

// [WM_NCRBUTTONDBLCLK] message handler.
//
// [WM_NCRBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondblclk
func (me *EventsWindow) WmNcRButtonDblClk(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCRBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCRBUTTONDOWN] message handler.
//
// [WM_NCRBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondown
func (me *EventsWindow) WmNcRButtonDown(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCRBUTTONDOWN, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCRBUTTONUP] message handler.
//
// [WM_NCRBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttonup
func (me *EventsWindow) WmNcRButtonUp(fun func(p WmNcMouse)) {
	me.Wm(co.WM_NCRBUTTONUP, func(p Wm) uintptr {
		fun(WmNcMouse{p})
		return me.defProcVal
	})
}

// [WM_NCXBUTTONDBLCLK] message handler.
//
// [WM_NCXBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondblclk
func (me *EventsWindow) WmNcXButtonDblClk(fun func(p WmNcMouseX)) {
	me.Wm(co.WM_NCXBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmNcMouseX{p})
		return 1
	})
}

// [WM_NCXBUTTONDOWN] message handler.
//
// [WM_NCXBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondown
func (me *EventsWindow) WmNcXButtonDown(fun func(p WmNcMouseX)) {
	me.Wm(co.WM_NCXBUTTONDOWN, func(p Wm) uintptr {
		fun(WmNcMouseX{p})
		return 1
	})
}

// [WM_NCXBUTTONUP] message handler.
//
// [WM_NCXBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttonup\
func (me *EventsWindow) WmNcXButtonUp(fun func(p WmNcMouseX)) {
	me.Wm(co.WM_NCXBUTTONUP, func(p Wm) uintptr {
		fun(WmNcMouseX{p})
		return 1
	})
}

// [WM_NEXTDLGCTL] message handler.
//
// [WM_NEXTDLGCTL]: https://learn.microsoft.com/en-us/windows/win32/dlgbox/wm-nextdlgctl
func (me *EventsWindow) WmNextDlgCtl(fun func(p WmNextDlgCtl)) {
	me.Wm(co.WM_NEXTDLGCTL, func(p Wm) uintptr {
		fun(WmNextDlgCtl{p})
		return me.defProcVal
	})
}

// [WM_NEXTMENU] message handler.
//
// [WM_NEXTMENU]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-nextmenu
func (me *EventsWindow) WmNextMenu(fun func(p WmNextMenu)) {
	me.Wm(co.WM_NEXTMENU, func(p Wm) uintptr {
		fun(WmNextMenu{p})
		return me.defProcVal
	})
}

// [WM_NULL] message handler.
//
// [WM_NULL]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-null
func (me *EventsWindow) WmNull(fun func()) {
	me.Wm(co.WM_NULL, func(p Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_PAINT] message handler.
//
// Note that, even if you don't actually paint anything, you still must call
// [win.HWND.BeginPaint] and [win.HWND.EndPaint], otherwise the window may get
// stuck.
//
// Example:
//
//	var wnd ui.Parent // initialized somewhere
//
//	wnd.On().WmPaint(func() {
//		var ps win.PAINTSTRUCT
//		hdc, _ := wnd.Hwnd().BeginPaint(&ps)
//		defer wnd.Hwnd().EndPaint(&ps)
//	}
//
// [WM_PAINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-paint
func (me *EventsWindow) WmPaint(fun func()) {
	me.Wm(co.WM_PAINT, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_PAINTCLIPBOARD] message handler.
//
// [WM_PAINTCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-paintclipboard
func (me *EventsWindow) WmPaintClipboard(fun func(WmPaintClipboard)) {
	me.Wm(co.WM_PAINTCLIPBOARD, func(p Wm) uintptr {
		fun(WmPaintClipboard{p})
		return me.defProcVal
	})
}

// [WM_PARENTNOTIFY] message handler.
//
// [WM_PARENTNOTIFY]: https://learn.microsoft.com/en-us/windows/win32/inputmsg/wm-parentnotify
func (me *EventsWindow) WmParentNotify(fun func(WmParentNotify)) {
	me.Wm(co.WM_PARENTNOTIFY, func(p Wm) uintptr {
		fun(WmParentNotify{p})
		return me.defProcVal
	})
}

// [WM_POWERBROADCAST] message handler.
//
// [WM_POWERBROADCAST]: https://learn.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
func (me *EventsWindow) WmPowerBroadcast(fun func(p WmPowerBroadcast)) {
	me.Wm(co.WM_POWERBROADCAST, func(p Wm) uintptr {
		fun(WmPowerBroadcast{p})
		return 1
	})
}

// [WM_PRINT] message handler.
//
// [WM_PRINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-print
func (me *EventsWindow) WmPrint(fun func(p WmPrint)) {
	me.Wm(co.WM_PRINT, func(p Wm) uintptr {
		fun(WmPrint{p})
		return me.defProcVal
	})
}

// [WM_QUERYDRAGICON] message handler.
//
// [WM_QUERYDRAGICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-querydragicon
func (me *EventsWindow) WmQueryDragIcon(fun func() win.HICON) {
	me.Wm(co.WM_QUERYDRAGICON, func(p Wm) uintptr {
		return uintptr(fun())
	})
}

// [WM_QUERYOPEN] message handler.
//
// [WM_QUERYOPEN]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-queryopen
func (me *EventsWindow) WmQueryOpen(fun func() bool) {
	me.Wm(co.WM_QUERYOPEN, func(p Wm) uintptr {
		return utl.BoolToUintptr(fun())
	})
}

// [WM_RBUTTONDBLCLK] message handler.
//
// [WM_RBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondblclk
func (me *EventsWindow) WmRButtonDblClk(fun func(p WmMouse)) {
	me.Wm(co.WM_RBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_RBUTTONDOWN] message handler.
//
// [WM_RBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondown
func (me *EventsWindow) WmRButtonDown(fun func(p WmMouse)) {
	me.Wm(co.WM_RBUTTONDOWN, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_RBUTTONUP] message handler.
//
// [WM_RBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttonup
func (me *EventsWindow) WmRButtonUp(fun func(p WmMouse)) {
	me.Wm(co.WM_RBUTTONUP, func(p Wm) uintptr {
		fun(WmMouse{p})
		return me.defProcVal
	})
}

// [WM_RENDERALLFORMATS] message handler.
//
// [WM_RENDERALLFORMATS]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-renderallformats
func (me *EventsWindow) WmRenderAllFormats(fun func()) {
	me.Wm(co.WM_RENDERALLFORMATS, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_RENDERFORMAT] message handler.
//
// [WM_RENDERFORMAT]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-renderformat
func (me *EventsWindow) WmRenderFormat(fun func(p WmRenderFormat)) {
	me.Wm(co.WM_RENDERFORMAT, func(p Wm) uintptr {
		fun(WmRenderFormat{p})
		return me.defProcVal
	})
}

// [WM_SETCURSOR] message handler.
//
// [WM_SETCURSOR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-setcursor
func (me *EventsWindow) WmSetCursor(fun func(p WmSetCursor) bool) {
	me.Wm(co.WM_SETCURSOR, func(p Wm) uintptr {
		return utl.BoolToUintptr(fun(WmSetCursor{p}))
	})
}

// [WM_SETFOCUS] message handler.
//
// [WM_SETFOCUS]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-setfocus
func (me *EventsWindow) WmSetFocus(fun func(p WmSetFocus)) {
	me.Wm(co.WM_SETFOCUS, func(p Wm) uintptr {
		fun(WmSetFocus{p})
		return me.defProcVal
	})
}

// [WM_SETFONT] message handler.
//
// [WM_SETFONT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-setfont
func (me *EventsWindow) WmSetFont(fun func(p WmSetFont)) {
	me.Wm(co.WM_SETFONT, func(p Wm) uintptr {
		fun(WmSetFont{p})
		return me.defProcVal
	})
}

// [WM_SETICON] message handler.
//
// [WM_SETICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
func (me *EventsWindow) WmSetIcon(fun func(p WmSetIcon) win.HICON) {
	me.Wm(co.WM_SETICON, func(p Wm) uintptr {
		return uintptr(fun(WmSetIcon{p}))
	})
}

// [WM_SETREDRAW] message handler.
//
// [WM_SETREDRAW]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-setredraw
func (me *EventsWindow) WmSetRedraw(fun func(p WmSetRedraw)) {
	me.Wm(co.WM_SETREDRAW, func(p Wm) uintptr {
		fun(WmSetRedraw{p})
		return me.defProcVal
	})
}

// [WM_SETTEXT] message handler.
//
// [WM_SETTEXT]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-settext
func (me *EventsWindow) WmSetText(fun func(p WmSetText) uintptr) {
	me.Wm(co.WM_SETTEXT, func(p Wm) uintptr {
		return fun(WmSetText{p})
	})
}

// [WM_SHOWWINDOW] message handler.
//
// [WM_SHOWWINDOW]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-showwindow
func (me *EventsWindow) WmShowWindow(fun func(p WmShowWindow)) {
	me.Wm(co.WM_SHOWWINDOW, func(p Wm) uintptr {
		fun(WmShowWindow{p})
		return me.defProcVal
	})
}

// [WM_SIZE] message handler.
//
// [WM_SIZE]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-size
func (me *EventsWindow) WmSize(fun func(p WmSize)) {
	me.Wm(co.WM_SIZE, func(p Wm) uintptr {
		fun(WmSize{p})
		return me.defProcVal
	})
}

// [WM_SIZECLIPBOARD] message handler.
//
// [WM_SIZECLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-sizeclipboard
func (me *EventsWindow) WmSizeClipboard(fun func(p WmSizeClipboard)) {
	me.Wm(co.WM_SIZECLIPBOARD, func(p Wm) uintptr {
		fun(WmSizeClipboard{p})
		return me.defProcVal
	})
}

// [WM_SIZING] message handler.
//
// [WM_SIZING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-sizing
func (me *EventsWindow) WmSizing(fun func(p WmSizing)) {
	me.Wm(co.WM_SIZING, func(p Wm) uintptr {
		fun(WmSizing{p})
		return 1
	})
}

// [WM_STYLECHANGED] message handler.
//
// [WM_STYLECHANGED]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-stylechanged
func (me *EventsWindow) WmStyleChanged(fun func(p WmStyles)) {
	me.Wm(co.WM_STYLECHANGED, func(p Wm) uintptr {
		fun(WmStyles{p})
		return me.defProcVal
	})
}

// [WM_STYLECHANGING] message handler.
//
// [WM_STYLECHANGING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-stylechanging
func (me *EventsWindow) WmStyleChanging(fun func(p WmStyles)) {
	me.Wm(co.WM_STYLECHANGING, func(p Wm) uintptr {
		fun(WmStyles{p})
		return me.defProcVal
	})
}

// [WM_SYNCPAINT] message handler.
//
// [WM_SYNCPAINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-syncpaint
func (me *EventsWindow) WmSyncPaint(fun func()) {
	me.Wm(co.WM_SYNCPAINT, func(p Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_SYSCHAR] message handler.
//
// [WM_SYSCHAR]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-syschar
func (me *EventsWindow) WmSysChar(fun func(p WmChar)) {
	me.Wm(co.WM_SYSCHAR, func(p Wm) uintptr {
		fun(WmChar{p})
		return me.defProcVal
	})
}

// [WM_SYSCOMMAND] message handler.
//
// [WM_SYSCOMMAND]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
func (me *EventsWindow) WmSysCommand(fun func(p WmSysCommand)) {
	me.Wm(co.WM_SYSCOMMAND, func(p Wm) uintptr {
		fun(WmSysCommand{p})
		return me.defProcVal
	})
}

// [WM_SYSDEADCHAR] message handler.
//
// [WM_SYSDEADCHAR]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-sysdeadchar
func (me *EventsWindow) WmSysDeadChar(fun func(p WmChar)) {
	me.Wm(co.WM_SYSDEADCHAR, func(p Wm) uintptr {
		fun(WmChar{p})
		return me.defProcVal
	})
}

// [WM_SYSKEYDOWN] message handler.
//
// [WM_SYSKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-syskeydown
func (me *EventsWindow) WmSysKeyDown(fun func(p WmKey)) {
	me.Wm(co.WM_SYSKEYDOWN, func(p Wm) uintptr {
		fun(WmKey{p})
		return me.defProcVal
	})
}

// [WM_SYSKEYUP] message handler.
//
// [WM_SYSKEYUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-syskeyup
func (me *EventsWindow) WmSysKeyUp(fun func(p WmKey)) {
	me.Wm(co.WM_SYSKEYUP, func(p Wm) uintptr {
		fun(WmKey{p})
		return me.defProcVal
	})
}

// [WM_THEMECHANGED] message handler.
//
// [WM_THEMECHANGED]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-themechanged
func (me *EventsWindow) WmThemeChanged(fun func()) {
	me.Wm(co.WM_THEMECHANGED, func(p Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_TIMECHANGE] message handler.
//
// [WM_TIMECHANGE]: https://learn.microsoft.com/en-us/windows/win32/sysinfo/wm-timechange
func (me *EventsWindow) WmTimeChange(fun func()) {
	me.Wm(co.WM_TIMECHANGE, func(_ Wm) uintptr {
		fun()
		return me.defProcVal
	})
}

// [WM_UNINITMENUPOPUP] message handler.
//
// [WM_UNINITMENUPOPUP]: https://learn.microsoft.com/en-us/windows/win32/menurc/wm-uninitmenupopup
func (me *EventsWindow) WmUnInitMenuPopup(fun func(p WmUnInitMenuPopup)) {
	me.Wm(co.WM_UNINITMENUPOPUP, func(p Wm) uintptr {
		fun(WmUnInitMenuPopup{p})
		return me.defProcVal
	})
}

// [WM_UNDO] message handler.
//
// [WM_UNDO]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-undo
func (me *EventsWindow) WmUndo(fun func() bool) {
	me.Wm(co.WM_UNDO, func(p Wm) uintptr {
		return utl.BoolToUintptr(fun())
	})
}

// [WM_VSCROLL] message handler.
//
// [WM_VSCROLL]: https://learn.microsoft.com/en-us/windows/win32/controls/wm-vscroll
func (me *EventsWindow) WmVScroll(fun func(p WmScroll)) {
	me.Wm(co.WM_VSCROLL, func(p Wm) uintptr {
		fun(WmScroll{p})
		return me.defProcVal
	})
}

// [WM_VSCROLLCLIPBOARD] message handler.
//
// [WM_VSCROLLCLIPBOARD]: https://learn.microsoft.com/en-us/windows/win32/dataxchg/wm-vscrollclipboard
func (me *EventsWindow) WmVScrollClipboard(fun func(p WmScrollClipboard)) {
	me.Wm(co.WM_VSCROLLCLIPBOARD, func(p Wm) uintptr {
		fun(WmScrollClipboard{p})
		return me.defProcVal
	})
}

// [WM_WINDOWPOSCHANGED] message handler.
//
// [WM_WINDOWPOSCHANGED]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-windowposchanged
func (me *EventsWindow) WmWindowPosChanged(fun func(p WmWindowPos)) {
	me.Wm(co.WM_WINDOWPOSCHANGED, func(p Wm) uintptr {
		fun(WmWindowPos{p})
		return me.defProcVal
	})
}

// [WM_WINDOWPOSCHANGING] message handler.
//
// [WM_WINDOWPOSCHANGING]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-windowposchanging
func (me *EventsWindow) WmWindowPosChanging(fun func(p WmWindowPos)) {
	me.Wm(co.WM_WINDOWPOSCHANGING, func(p Wm) uintptr {
		fun(WmWindowPos{p})
		return me.defProcVal
	})
}

// [WM_WTSSESSION_CHANGE] message handler.
//
// [WM_WTSSESSION_CHANGE]: https://learn.microsoft.com/en-us/windows/win32/termserv/wm-wtssession-change
func (me *EventsWindow) WmWtsSessionChange(fun func(p WmWtsSessionChange)) {
	me.Wm(co.WM_WTSSESSION_CHANGE, func(p Wm) uintptr {
		fun(WmWtsSessionChange{p})
		return me.defProcVal
	})
}

// [WM_XBUTTONDBLCLK] message handler.
//
// [WM_XBUTTONDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondblclk
func (me *EventsWindow) WmXButtonDblClk(fun func(p WmMouse)) {
	me.Wm(co.WM_XBUTTONDBLCLK, func(p Wm) uintptr {
		fun(WmMouse{p})
		return 1
	})
}

// [WM_XBUTTONDOWN] message handler.
//
// [WM_XBUTTONDOWN]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondown
func (me *EventsWindow) WmXButtonDown(fun func(p WmMouse)) {
	me.Wm(co.WM_XBUTTONDOWN, func(p Wm) uintptr {
		fun(WmMouse{p})
		return 1
	})
}

// [WM_XBUTTONUP] message handler.
//
// [WM_XBUTTONUP]: https://learn.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttonup
func (me *EventsWindow) WmXButtonUp(fun func(p WmMouse)) {
	me.Wm(co.WM_XBUTTONUP, func(p Wm) uintptr {
		fun(WmMouse{p})
		return 1
	})
}
