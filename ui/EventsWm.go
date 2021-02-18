/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"sort"
	"strings"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Keeps WM and WM_TIMER message callbacks.
type _EventsWm struct {
	mapMsgsRet map[co.WM]func(p Wm) uintptr // meaningful return value
	mapMsgs    map[co.WM]func(p Wm)         // just returns zero (or TRUE if dialog)
	mapTims    map[int]func()
}

// Constructor.
func _NewEventsWm() *_EventsWm {
	return &_EventsWm{
		mapMsgsRet: make(map[co.WM]func(p Wm) uintptr),
		mapMsgs:    make(map[co.WM]func(p Wm)),
		mapTims:    make(map[int]func()),
	}
}

func (me *_EventsWm) processMessage(
	msg co.WM, p Wm) (retVal uintptr, useRetVal bool, wasHandled bool) {

	if msg == co.WM_TIMER {
		if userFunc, hasFunc := me.mapTims[int(p.WParam)]; hasFunc {
			userFunc()
			retVal, useRetVal, wasHandled = 0, false, true
			return
		}

	} else if userFunc, hasFunc := me.mapMsgs[msg]; hasFunc {
		userFunc(p)
		retVal, useRetVal, wasHandled = 0, false, true
		return

	} else if userFunc, hasFunc := me.mapMsgsRet[msg]; hasFunc {
		retVal, useRetVal, wasHandled = userFunc(p), true, true
		return
	}

	retVal, useRetVal, wasHandled = 0, false, false
	return
}

func (me *_EventsWm) hasMessages() bool {
	return len(me.mapMsgsRet) > 0 ||
		len(me.mapMsgs) > 0 ||
		len(me.mapTims) > 0
}

func (me *_EventsWm) addMsgRet(msg co.WM, userFunc func(p Wm) uintptr) {
	me.mapMsgsRet[msg] = userFunc
}

func (me *_EventsWm) addMsg(msg co.WM, userFunc func(p Wm)) {
	me.mapMsgs[msg] = userFunc
}

//------------------------------------------------------------------------------

// Generic message handler.
//
// Avoid this method, prefer the specific message handlers.
//
// https://docs.microsoft.com/en-us/windows/win32/learnwin32/window-messages
func (me *_EventsWm) Wm(msg co.WM, userFunc func(p Wm) uintptr) {
	me.addMsgRet(msg, userFunc)
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-timer
func (me *_EventsWm) WmTimer(nIDEvent int, userFunc func()) {
	if me.mapTims == nil {
		me.mapTims = make(map[int]func())
	}
	me.mapTims[nIDEvent] = userFunc
}

//------------------------------------------------------------------------------

// Warning: default handled in WindowMain.
//
// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-activate
func (me *_EventsWm) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(co.WM_ACTIVATE, func(p Wm) {
		userFunc(WmActivate{m: p})
	})
}

type WmActivate struct{ m Wm }

func (p WmActivate) Event() co.WA                           { return co.WA(p.m.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                      { return p.m.WParam.HiWord() != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.m.LParam) }

type WmActivateApp struct{ m Wm }

func (p WmActivateApp) IsBeingActivated() bool { return p.m.WParam != 0 }
func (p WmActivateApp) ThreadId() uint         { return uint(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
func (me *_EventsWm) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsgRet(co.WM_APPCOMMAND, func(p Wm) uintptr {
		userFunc(WmAppCommand{m: p})
		return 1
	})
}

type WmAppCommand struct{ m Wm }

func (p WmAppCommand) OwnerWindow() win.HWND     { return win.HWND(p.m.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.m.LParam.HiWord() &^ 0xf000) }
func (p WmAppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.m.LParam.HiWord() & 0xf000) }
func (p WmAppCommand) Keys() co.MK               { return co.MK(p.m.LParam.LoWord()) }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-askcbformatname
func (me *_EventsWm) WmAskCbFormatName(userFunc func(p WmAskCbFormatName)) {
	me.addMsg(co.WM_ASKCBFORMATNAME, func(p Wm) {
		userFunc(WmAskCbFormatName{m: p})
	})
}

type WmAskCbFormatName struct{ m Wm }

func (p WmAskCbFormatName) BufferSize() int { return int(p.m.WParam) }
func (p WmAskCbFormatName) Buffer() *uint16 { return (*uint16)(unsafe.Pointer(p.m.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-cancelmode
func (me *_EventsWm) WmCancelMode(userFunc func()) {
	me.addMsg(co.WM_CANCELMODE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-capturechanged
func (me *_EventsWm) WmCaptureChanged(userFunc func(hwndGainingMouse win.HWND)) {
	me.addMsg(co.WM_CAPTURECHANGED, func(p Wm) {
		userFunc(win.HWND(p.LParam))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-changecbchain
func (me *_EventsWm) WmChangeCbChain(userFunc func(p WmChangeCbChain)) {
	me.addMsg(co.WM_CHANGECBCHAIN, func(p Wm) {
		userFunc(WmChangeCbChain{m: p})
	})
}

type WmChangeCbChain struct{ m Wm }

func (p WmChangeCbChain) WindowBeingRemoved() win.HWND { return win.HWND(p.m.WParam) }
func (p WmChangeCbChain) NextWindow() win.HWND         { return win.HWND(p.m.LParam) }
func (p WmChangeCbChain) IsLastWindow() bool           { return p.m.LParam == 0 }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (me *_EventsWm) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p Wm) {
		userFunc(WmChar{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-chartoitem
func (me *_EventsWm) WmCharToItem(userFunc func(p WmCharToItem)) {
	me.addMsg(co.WM_CHARTOITEM, func(p Wm) {
		userFunc(WmCharToItem{m: p})
	})
}

type WmCharToItem struct{ m Wm }

func (p WmCharToItem) CharCode() rune        { return rune(p.m.WParam.LoWord()) }
func (p WmCharToItem) CurrentCaretPos() int  { return int(p.m.WParam.HiWord()) }
func (p WmCharToItem) HwndListBox() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-childactivate
func (me *_EventsWm) WmChildActivate(userFunc func()) {
	me.addMsg(co.WM_CHILDACTIVATE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-clipboardupdate
func (me *_EventsWm) WmClipboardUpdate(userFunc func()) {
	me.addMsg(co.WM_CLIPBOARDUPDATE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-close
//
// Warning: default handled in WindowModal.
func (me *_EventsWm) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-compareitem
func (me *_EventsWm) WmCompareItem(userFunc func(p WmCompareItem) int) {
	me.addMsgRet(co.WM_COMPAREITEM, func(p Wm) uintptr {
		return uintptr(userFunc(WmCompareItem{m: p}))
	})
}

type WmCompareItem struct{ m Wm }

func (p WmCompareItem) ControlId() int { return int(p.m.WParam) }
func (p WmCompareItem) CompareItemStruct() *win.COMPAREITEMSTRUCT {
	return (*win.COMPAREITEMSTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-contextmenu
func (me *_EventsWm) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p Wm) {
		userFunc(WmContextMenu{m: p})
	})
}

type WmContextMenu struct{ m Wm }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.m.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.m.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-copydata
func (me *_EventsWm) WmCopyData(userFunc func(p WmCopyData) bool) {
	me.addMsgRet(co.WM_COPYDATA, func(p Wm) uintptr {
		return _global.BoolToUintptr(userFunc(WmCopyData{m: p}))
	})
}

type WmCopyData struct{ m Wm }

func (p WmCopyData) WindowPassingData() win.HWND { return win.HWND(p.m.WParam) }
func (p WmCopyData) CopyDataStruct() *win.COPYDATASTRUCT {
	return (*win.COPYDATASTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-create
func (me *_EventsWm) WmCreate(userFunc func(p *win.CREATESTRUCT) int) {
	me.addMsgRet(co.WM_CREATE, func(p Wm) uintptr {
		return uintptr(userFunc((*win.CREATESTRUCT)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorbtn
func (me *_EventsWm) WmCtlColorBtn(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORBTN, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-ctlcolordlg
func (me *_EventsWm) WmCtlColorDlg(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORDLG, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcoloredit
func (me *_EventsWm) WmCtlColorEdit(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLOREDIT, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorlistbox
func (me *_EventsWm) WmCtlColorListBox(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORLISTBOX, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorscrollbar
func (me *_EventsWm) WmCtlColorScrollBar(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORSCROLLBAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
func (me *_EventsWm) WmCtlColorStatic(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.addMsgRet(co.WM_CTLCOLORSTATIC, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
func (me *_EventsWm) WmDeadChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p Wm) {
		userFunc(WmChar{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-deleteitem
func (me *_EventsWm) WmDeleteItem(userFunc func(p WmDeleteItem)) {
	me.addMsgRet(co.WM_DELETEITEM, func(p Wm) uintptr {
		userFunc(WmDeleteItem{m: p})
		return 1
	})
}

type WmDeleteItem struct{ m Wm }

func (p WmDeleteItem) ControlId() int { return int(p.m.WParam) }
func (p WmDeleteItem) DeleteItemStruct() *win.DELETEITEMSTRUCT {
	return (*win.DELETEITEMSTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
func (me *_EventsWm) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-destroyclipboard
func (me *_EventsWm) WmDestroyClipboard(userFunc func()) {
	me.addMsg(co.WM_DESTROYCLIPBOARD, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-devmodechange
func (me *_EventsWm) WmDevModeChange(userFunc func(deviceName string)) {
	me.addMsg(co.WM_DEVMODECHANGE, func(p Wm) {
		userFunc(win.Str.FromUint16Ptr((*uint16)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-displaychange
func (me *_EventsWm) WmDisplayChange(userFunc func(p WmDisplayChange)) {
	me.addMsg(co.WM_DISPLAYCHANGE, func(p Wm) {
		userFunc(WmDisplayChange{m: p})
	})
}

type WmDisplayChange struct{ m Wm }

func (p WmDisplayChange) BitsPerPixel() int { return int(p.m.WParam) }
func (p WmDisplayChange) Size() win.SIZE    { return p.m.LParam.MakeSize() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-drawclipboard
func (me *_EventsWm) WmDrawClipboard(userFunc func()) {
	me.addMsg(co.WM_DRAWCLIPBOARD, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-drawitem
func (me *_EventsWm) WmDrawItem(userFunc func(p WmDrawItem)) {
	me.addMsgRet(co.WM_DRAWITEM, func(p Wm) uintptr {
		userFunc(WmDrawItem{m: p})
		return 1
	})
}

type WmDrawItem struct{ m Wm }

func (p WmDrawItem) ControlId() int   { return int(p.m.WParam) }
func (p WmDrawItem) IsFromMenu() bool { return p.m.WParam == 0 }
func (p WmDrawItem) DrawItemStruct() *win.DRAWITEMSTRUCT {
	return (*win.DRAWITEMSTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/shell/wm-dropfiles
func (me *_EventsWm) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p Wm) {
		userFunc(WmDropFiles{m: p})
	})
}

type WmDropFiles struct{ m Wm }

func (p WmDropFiles) Hdrop() win.HDROP { return win.HDROP(p.m.WParam) }
func (p WmDropFiles) Count() int       { return int(p.Hdrop().DragQueryFile(0xffff_ffff, nil, 0)) }

// Calls DragQueryFile() successively to retrieve all file names, and releases
// the HDROP handle with DragFinish().
func (p WmDropFiles) RetrieveAll() []string {
	count := uint32(p.Count())
	files := make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		pathLen := p.Hdrop().DragQueryFile(i, nil, 0) + 1 // room for terminating null
		pathBuf := make([]uint16, pathLen)
		p.Hdrop().DragQueryFile(i, &pathBuf[0], pathLen)
		files = append(files, win.Str.FromUint16Slice(pathBuf))
	}
	p.Hdrop().DragFinish()
	sort.Slice(files, func(i, j int) bool { // case insensitive
		return strings.ToUpper(files[i]) < strings.ToUpper(files[j])
	})
	return files
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-enable
func (me *_EventsWm) WmEnable(userFunc func(hasBeenEnabled bool)) {
	me.addMsg(co.WM_ENABLE, func(p Wm) {
		userFunc(p.WParam != 0)
	})
}

// https://docs.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
func (me *_EventsWm) WmEndSession(userFunc func(p WmEndSession)) {
	me.addMsg(co.WM_ENDSESSION, func(p Wm) {
		userFunc(WmEndSession{m: p})
	})
}

type WmEndSession struct{ m Wm }

func (p WmEndSession) IsSessionBeingEnded() bool { return p.m.WParam != 0 }
func (p WmEndSession) Event() co.ENDSESSION      { return co.ENDSESSION(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
func (me *_EventsWm) WmEnterIdle(userFunc func(p WmEnterIdle)) {
	me.addMsg(co.WM_ENTERIDLE, func(p Wm) {
		userFunc(WmEnterIdle{m: p})
	})
}

type WmEnterIdle struct{ m Wm }

func (p WmEnterIdle) Displayed() co.MSGF       { return co.MSGF(p.m.WParam) }
func (p WmEnterIdle) DialogOrWindow() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-entermenuloop
func (me *_EventsWm) WmEnterMenuLoop(userFunc func(isTrackPopupMenu bool)) {
	me.addMsg(co.WM_ENTERMENULOOP, func(p Wm) {
		userFunc(p.WParam != 0)
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-entersizemove
func (me *_EventsWm) WmEnterSizeMove(userFunc func()) {
	me.addMsg(co.WM_ENTERSIZEMOVE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-erasebkgnd
func (me *_EventsWm) WmEraseBkgnd(userFunc func(hdc win.HDC) int) {
	me.addMsgRet(co.WM_ERASEBKGND, func(p Wm) uintptr {
		return uintptr(userFunc(win.HDC(p.WParam)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-exitmenuloop
func (me *_EventsWm) WmExitMenuLoop(userFunc func(isShortcutMenu bool)) {
	me.addMsg(co.WM_EXITMENULOOP, func(p Wm) {
		userFunc(p.WParam != 0)
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-exitsizemove
func (me *_EventsWm) WmExitSizeMove(userFunc func()) {
	me.addMsg(co.WM_EXITSIZEMOVE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-fontchange
func (me *_EventsWm) WmFontChange(userFunc func()) {
	me.addMsg(co.WM_FONTCHANGE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
func (me *_EventsWm) WmGetDlgCode(userFunc func(p WmGetDlgCode) co.DLGC) {
	me.addMsgRet(co.WM_GETDLGCODE, func(p Wm) uintptr {
		return uintptr(userFunc(WmGetDlgCode{m: p}))
	})
}

type WmGetDlgCode struct{ m Wm }

func (p WmGetDlgCode) Raw() Wm               { return p.m }
func (p WmGetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.m.WParam) }
func (p WmGetDlgCode) IsQuery() bool         { return p.m.LParam == 0 }
func (p WmGetDlgCode) Msg() *win.MSG         { return (*win.MSG)(unsafe.Pointer(p.m.LParam)) }
func (p WmGetDlgCode) HasAlt() bool          { return (win.GetAsyncKeyState(co.VK_MENU) & 0x8000) != 0 }
func (p WmGetDlgCode) HasCtrl() bool         { return (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0 }
func (p WmGetDlgCode) HasShift() bool        { return (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-getfont
func (me *_EventsWm) WmGetFont(userFunc func() win.HFONT) {
	me.addMsgRet(co.WM_FONTCHANGE, func(p Wm) uintptr {
		return uintptr(userFunc())
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-gettitlebarinfoex
func (me *_EventsWm) WmGetTitleBarInfoEx(userFunc func(p *win.TITLEBARINFOEX)) {
	me.addMsg(co.WM_GETTITLEBARINFOEX, func(p Wm) {
		userFunc((*win.TITLEBARINFOEX)(unsafe.Pointer(p.LParam)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/shell/wm-help
func (me *_EventsWm) WmHelp(userFunc func(p *win.HELPINFO)) {
	me.addMsgRet(co.WM_HELP, func(p Wm) uintptr {
		userFunc((*win.HELPINFO)(unsafe.Pointer(p.LParam)))
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
func (me *_EventsWm) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p Wm) {
		userFunc(WmHotKey{m: p})
	})
}

type WmHotKey struct{ m Wm }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.m.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.m.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.m.LParam.HiWord()) }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll
func (me *_EventsWm) WmHScroll(userFunc func(p WmScroll)) {
	me.addMsg(co.WM_HSCROLL, func(p Wm) {
		userFunc(WmScroll{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-hscrollclipboard
func (me *_EventsWm) WmHScrollClipboard(userFunc func(p WmScroll)) {
	me.addMsg(co.WM_HSCROLLCLIPBOARD, func(p Wm) {
		userFunc(WmScroll{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-initdialog
func (me *_EventsWm) WmInitDialog(userFunc func(hFocusedCtrl win.HWND) bool) {
	me.addMsgRet(co.WM_INITDIALOG, func(p Wm) uintptr {
		return _global.BoolToUintptr(userFunc(win.HWND(p.WParam)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
func (me *_EventsWm) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p Wm) {
		userFunc(WmInitMenuPopup{m: p})
	})
}

type WmInitMenuPopup struct{ m Wm }

func (p WmInitMenuPopup) Hmenu() win.HMENU   { return win.HMENU(p.m.WParam) }
func (p WmInitMenuPopup) Pos() int           { return int(p.m.LParam.LoWord()) }
func (p WmInitMenuPopup) IsWindowMenu() bool { return p.m.LParam.HiWord() != 0 }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keydown
func (me *_EventsWm) WmKeyDown(userFunc func(p WmKey)) {
	me.addMsg(co.WM_KEYDOWN, func(p Wm) {
		userFunc(WmKey{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keyup
func (me *_EventsWm) WmKeyUp(userFunc func(p WmKey)) {
	me.addMsg(co.WM_KEYUP, func(p Wm) {
		userFunc(WmKey{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-killfocus
func (me *_EventsWm) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p Wm) {
		userFunc(WmKillFocus{m: p})
	})
}

type WmKillFocus struct{ m Wm }

func (p WmKillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondblclk
func (me *_EventsWm) WmLButtonDblClk(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
func (me *_EventsWm) WmLButtonDown(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttonup
func (me *_EventsWm) WmLButtonUp(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_LBUTTONUP, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondblclk
func (me *_EventsWm) WmMButtonDblClk(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondown
func (me *_EventsWm) WmMButtonDown(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttonup
func (me *_EventsWm) WmMButtonUp(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_MBUTTONUP, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
func (me *_EventsWm) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.addMsgRet(co.WM_MENUCHAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenuChar{m: p}))
	})
}

type WmMenuChar struct{ m Wm }

func (p WmMenuChar) CharCode() rune        { return rune(p.m.WParam.LoWord()) }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.m.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
func (me *_EventsWm) WmMenuCommand(userFunc func(p WmMenu)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p Wm) {
		userFunc(WmMenu{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
func (me *_EventsWm) WmMenuDrag(userFunc func(p WmMenu) co.MND) {
	me.addMsgRet(co.WM_MENUDRAG, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenu{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
func (me *_EventsWm) WmMenuGetObject(userFunc func(p *win.MENUGETOBJECTINFO) co.MNGO) {
	me.addMsgRet(co.WM_MENUGETOBJECT, func(p Wm) uintptr {
		return uintptr(userFunc((*win.MENUGETOBJECTINFO)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menurbuttonup
func (me *_EventsWm) WmMenuRButtonUp(userFunc func(p WmMenu)) {
	me.addMsg(co.WM_MENURBUTTONUP, func(p Wm) {
		userFunc(WmMenu{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
func (me *_EventsWm) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p Wm) {
		userFunc(WmMenuSelect{m: p})
	})
}

type WmMenuSelect struct{ m Wm }

func (p WmMenuSelect) Item() int        { return int(p.m.WParam.LoWord()) }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.m.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousehover
func (me *_EventsWm) WmMouseHover(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mouseleave
func (me *_EventsWm) WmMouseLeave(userFunc func()) {
	me.addMsg(co.WM_MOUSELEAVE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
func (me *_EventsWm) WmMouseMove(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-move
func (me *_EventsWm) WmMove(userFunc func(clientAreaPos win.POINT)) {
	me.addMsg(co.WM_MOVE, func(p Wm) {
		userFunc(p.LParam.MakePoint())
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-moving
func (me *_EventsWm) WmMoving(userFunc func(windowPos *win.RECT)) {
	me.addMsgRet(co.WM_MOVING, func(p Wm) uintptr {
		userFunc((*win.RECT)(unsafe.Pointer(p.LParam)))
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncactivate
func (me *_EventsWm) WmNcActivate(userFunc func(p WmNcActivate) bool) {
	me.addMsgRet(co.WM_NCACTIVATE, func(p Wm) uintptr {
		return _global.BoolToUintptr(userFunc(WmNcActivate{m: p}))
	})
}

type WmNcActivate struct{ m Wm }

func (p WmNcActivate) IsActive() bool            { return p.m.WParam != 0 }
func (p WmNcActivate) IsVisualStyleActive() bool { return p.m.LParam == 0 }
func (p WmNcActivate) UpdatedRegion() win.HRGN   { return win.HRGN(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
func (me *_EventsWm) WmNcCalcSize(userFunc func(p WmNcCalcSize) co.WVR) {
	me.addMsgRet(co.WM_NCCALCSIZE, func(p Wm) uintptr {
		return uintptr(userFunc(WmNcCalcSize{m: p}))
	})
}

type WmNcCalcSize struct{ m Wm }

func (p WmNcCalcSize) ShouldIndicateValidPart() bool { return p.m.WParam != 0 }
func (p WmNcCalcSize) NcCalcSizeParams() *win.NCCALCSIZE_PARAMS {
	return (*win.NCCALCSIZE_PARAMS)(unsafe.Pointer(p.m.LParam))
}
func (p WmNcCalcSize) Rect() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.m.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncdestroy
//
// Warning: default handled in WindowMain.
func (me *_EventsWm) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
func (me *_EventsWm) WmNcHitTest(userFunc func(cursorCoord win.POINT) co.HT) {
	me.addMsgRet(co.WM_NCHITTEST, func(p Wm) uintptr {
		return uintptr(userFunc(p.LParam.MakePoint()))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondblclk
func (me *_EventsWm) WmNcLButtonDblClk(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCLBUTTONDBLCLK, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondown
func (me *_EventsWm) WmNcLButtonDown(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCLBUTTONDOWN, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttonup
func (me *_EventsWm) WmNcLButtonUp(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCLBUTTONUP, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondblclk
func (me *_EventsWm) WmNcMButtonDblClk(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCMBUTTONDBLCLK, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondown
func (me *_EventsWm) WmNcMButtonDown(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCMBUTTONDOWN, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttonup
func (me *_EventsWm) WmNcMButtonUp(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCMBUTTONUP, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousehover
func (me *_EventsWm) WmNcMouseHover(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCMOUSEHOVER, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmouseleave
func (me *_EventsWm) WmNcMouseLeave(userFunc func()) {
	me.addMsg(co.WM_NCMOUSELEAVE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousemove
func (me *_EventsWm) WmNcMouseMove(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCMOUSEMOVE, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-ncpaint
//
// Warning: default handled in WindowControl.
func (me *_EventsWm) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(co.WM_NCPAINT, func(p Wm) {
		userFunc(WmNcPaint{m: p})
	})
}

type WmNcPaint struct{ m Wm }

func (p WmNcPaint) Raw() Wm               { return p.m }
func (p WmNcPaint) UpdatedHrgn() win.HRGN { return win.HRGN(p.m.WParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondblclk
func (me *_EventsWm) WmNcRButtonDblClk(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCRBUTTONDBLCLK, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondown
func (me *_EventsWm) WmNcRButtonDown(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCRBUTTONDOWN, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttonup
func (me *_EventsWm) WmNcRButtonUp(userFunc func(p WmNcMouse)) {
	me.addMsg(co.WM_NCRBUTTONUP, func(p Wm) {
		userFunc(WmNcMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondblclk
func (me *_EventsWm) WmNcXButtonDblClk(userFunc func(p WmNcMouseX)) {
	me.addMsgRet(co.WM_NCXBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmNcMouseX{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondown
func (me *_EventsWm) WmNcXButtonDown(userFunc func(p WmNcMouseX)) {
	me.addMsgRet(co.WM_NCXBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmNcMouseX{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttonup
func (me *_EventsWm) WmNcXButtonUp(userFunc func(p WmNcMouseX)) {
	me.addMsgRet(co.WM_NCXBUTTONUP, func(p Wm) uintptr {
		userFunc(WmNcMouseX{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-nextmenu
func (me *_EventsWm) WmNextMenu(userFunc func(p WmNextMenu)) {
	me.addMsg(co.WM_NEXTMENU, func(p Wm) {
		userFunc(WmNextMenu{m: p})
	})
}

type WmNextMenu struct{ m Wm }

func (p WmNextMenu) VirtualKeyCode() co.VK { return co.VK(p.m.WParam) }
func (p WmNextMenu) MdiNextMenu() *win.MDINEXTMENU {
	return (*win.MDINEXTMENU)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-paint
func (me *_EventsWm) WmPaint(userFunc func()) {
	me.addMsg(co.WM_PAINT, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-paintclipboard
func (me *_EventsWm) WmPaintClipboard(userFunc func(WmPaintClipboard)) {
	me.addMsg(co.WM_PAINTCLIPBOARD, func(p Wm) {
		userFunc(WmPaintClipboard{m: p})
	})
}

type WmPaintClipboard struct{ m Wm }

func (p WmPaintClipboard) CbViewerWindow() win.HWND { return win.HWND(p.m.WParam) }
func (p WmPaintClipboard) PaintStruct() *win.PAINTSTRUCT {
	return (*win.PAINTSTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
func (me *_EventsWm) WmPowerBroadcast(userFunc func(p WmPowerBroadcast)) {
	me.addMsgRet(co.WM_POWERBROADCAST, func(p Wm) uintptr {
		userFunc(WmPowerBroadcast{m: p})
		return 1
	})
}

type WmPowerBroadcast struct{ m Wm }

func (p WmPowerBroadcast) Event() co.PBT { return co.PBT(p.m.WParam) }
func (p WmPowerBroadcast) PowerBroadcastSetting() *win.POWERBROADCAST_SETTING {
	if p.Event() == co.PBT_POWERSETTINGCHANGE {
		return (*win.POWERBROADCAST_SETTING)(unsafe.Pointer(p.m.LParam))
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-print
func (me *_EventsWm) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(co.WM_PRINT, func(p Wm) {
		userFunc(WmPrint{m: p})
	})
}

type WmPrint struct{ m Wm }

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.m.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondblclk
func (me *_EventsWm) WmRButtonDblClk(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondown
func (me *_EventsWm) WmRButtonDown(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttonup
func (me *_EventsWm) WmRButtonUp(userFunc func(p WmMouse)) {
	me.addMsg(co.WM_RBUTTONUP, func(p Wm) {
		userFunc(WmMouse{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-renderallformats
func (me *_EventsWm) WmRenderAllFormats(userFunc func()) {
	me.addMsg(co.WM_RENDERALLFORMATS, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-renderformat
func (me *_EventsWm) WmRenderFormat(userFunc func(clipboardFormat co.CF)) {
	me.addMsg(co.WM_RENDERFORMAT, func(p Wm) {
		userFunc(co.CF(p.WParam))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-setfocus
//
// Warning: default handled in WindowMain and WindowModal.
func (me *_EventsWm) WmSetFocus(userFunc func(hwndLosingFocus win.HWND)) {
	me.addMsg(co.WM_SETFOCUS, func(p Wm) {
		userFunc(win.HWND(p.LParam))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-setfont
func (me *_EventsWm) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(co.WM_SETFONT, func(p Wm) {
		userFunc(WmSetFont{m: p})
	})
}

type WmSetFont struct{ m Wm }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.m.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.m.LParam == 1 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
func (me *_EventsWm) WmSetIcon(userFunc func(p WmSetIcon) win.HICON) {
	me.addMsgRet(co.WM_SETICON, func(p Wm) uintptr {
		return uintptr(userFunc(WmSetIcon{m: p}))
	})
}

type WmSetIcon struct{ m Wm }

func (p WmSetIcon) Size() co.ICON_SZ { return co.ICON_SZ(p.m.WParam) }
func (p WmSetIcon) Hicon() win.HICON { return win.HICON(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-size
func (me *_EventsWm) WmSize(userFunc func(p WmSize)) {
	me.addMsg(co.WM_SIZE, func(p Wm) {
		userFunc(WmSize{m: p})
	})
}

type WmSize struct{ m Wm }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.m.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.m.LParam.MakeSize() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-sizeclipboard
func (me *_EventsWm) WmSizeClipboard(userFunc func(p WmSizeClipboard)) {
	me.addMsg(co.WM_SIZECLIPBOARD, func(p Wm) {
		userFunc(WmSizeClipboard{m: p})
	})
}

type WmSizeClipboard struct{ m Wm }

func (p WmSizeClipboard) CbViewerWindow() win.HWND { return win.HWND(p.m.WParam) }
func (p WmSizeClipboard) NewDimensions() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.m.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syschar
func (me *_EventsWm) WmSysChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p Wm) {
		userFunc(WmChar{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
func (me *_EventsWm) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p Wm) {
		userFunc(WmSysCommand{m: p})
	})
}

type WmSysCommand struct{ m Wm }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.m.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.m.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-sysdeadchar
func (me *_EventsWm) WmSysDeadChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p Wm) {
		userFunc(WmChar{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeydown
func (me *_EventsWm) WmSysKeyDown(userFunc func(p WmKey)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p Wm) {
		userFunc(WmKey{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeyup
func (me *_EventsWm) WmSysKeyUp(userFunc func(p WmKey)) {
	me.addMsg(co.WM_SYSKEYUP, func(p Wm) {
		userFunc(WmKey{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/sysinfo/wm-timechange
func (me *_EventsWm) WmTimeChange(userFunc func()) {
	me.addMsg(co.WM_TIMECHANGE, func(p Wm) {
		userFunc()
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-uninitmenupopup
func (me *_EventsWm) WmUnInitMenuPopup(userFunc func(menu win.HMENU)) {
	me.addMsg(co.WM_UNINITMENUPOPUP, func(p Wm) {
		userFunc(win.HMENU(p.WParam))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-vscroll
func (me *_EventsWm) WmVScroll(userFunc func(p WmScroll)) {
	me.addMsg(co.WM_VSCROLL, func(p Wm) {
		userFunc(WmScroll{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-vscrollclipboard
func (me *_EventsWm) WmVScrollClipboard(userFunc func(p WmScroll)) {
	me.addMsg(co.WM_VSCROLLCLIPBOARD, func(p Wm) {
		userFunc(WmScroll{m: p})
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondblclk
func (me *_EventsWm) WmXButtonDblClk(userFunc func(p WmMouse)) {
	me.addMsgRet(co.WM_XBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondown
func (me *_EventsWm) WmXButtonDown(userFunc func(p WmMouse)) {
	me.addMsgRet(co.WM_XBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttonup
func (me *_EventsWm) WmXButtonUp(userFunc func(p WmMouse)) {
	me.addMsgRet(co.WM_XBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 1
	})
}
