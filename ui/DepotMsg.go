/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Keeps user message handlers.
type _DepotMsg struct {
	mapMsgs map[co.WM]func(p Wm) uintptr
}

func (me *_DepotMsg) processMessage(msg co.WM, p Wm) (uintptr, bool) {
	if userFunc, hasMsg := me.mapMsgs[msg]; hasMsg {
		return userFunc(p), true // user handler found
	}
	return 0, false // no user handler found
}

func (me *_DepotMsg) hasMessages() bool {
	return len(me.mapMsgs) > 0
}

// Handles a raw, unspecific window message. There will be no treatment of
// WPARAM/LPARAM data, you'll have to unpack all the information manually. This
// is very dangerous.
//
// Unless you have a very good reason, always prefer the specific message
// handlers.
func (me *_DepotMsg) Wm(message co.WM, userFunc func(p Wm) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p Wm) uintptr, 4) // arbitrary capacity, just to speed-up the first allocations
	}
	me.mapMsgs[message] = userFunc
}

//------------------------------------------------------------------------------

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-activate
//
// Warning: default handled in WindowMain.
func (me *_DepotMsg) WmActivate(userFunc func(p WmActivate)) {
	me.Wm(co.WM_ACTIVATE, func(p Wm) uintptr {
		userFunc(WmActivate{m: p})
		return 0
	})
}

type WmActivate struct{ m Wm }

func (p WmActivate) Event() co.WA                           { return co.WA(p.m.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                      { return p.m.WParam.HiWord() != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-activateapp
func (me *_DepotMsg) WmActivateApp(userFunc func(p WmActivateApp)) {
	me.Wm(co.WM_ACTIVATEAPP, func(p Wm) uintptr {
		userFunc(WmActivateApp{m: p})
		return 0
	})
}

type WmActivateApp struct{ m Wm }

func (p WmActivateApp) IsBeingActivated() bool { return p.m.WParam != 0 }
func (p WmActivateApp) ThreadId() uint         { return uint(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
func (me *_DepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.Wm(co.WM_APPCOMMAND, func(p Wm) uintptr {
		userFunc(WmAppCommand{m: p})
		return 1
	})
}

type WmAppCommand struct{ m Wm }

func (p WmAppCommand) OwnerWindow() win.HWND     { return win.HWND(p.m.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.m.LParam.HiWord() &^ 0xF000) }
func (p WmAppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.m.LParam.HiWord() & 0xF000) }
func (p WmAppCommand) Keys() co.MK               { return co.MK(p.m.LParam.LoWord()) }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-askcbformatname
func (me *_DepotMsg) WmAskCbFormatName(userFunc func(p WmAskCbFormatName)) {
	me.Wm(co.WM_ASKCBFORMATNAME, func(p Wm) uintptr {
		userFunc(WmAskCbFormatName{m: p})
		return 0
	})
}

type WmAskCbFormatName struct{ m Wm }

func (p WmAskCbFormatName) BufferSize() uint { return uint(p.m.WParam) }
func (p WmAskCbFormatName) Buffer() *uint16  { return (*uint16)(unsafe.Pointer(p.m.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-cancelmode
func (me *_DepotMsg) WmCancelMode(userFunc func()) {
	me.Wm(co.WM_CANCELMODE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-capturechanged
func (me *_DepotMsg) WmCaptureChanged(userFunc func(hwndGainingMouse win.HWND)) {
	me.Wm(co.WM_CAPTURECHANGED, func(p Wm) uintptr {
		userFunc(win.HWND(p.LParam))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-changecbchain
func (me *_DepotMsg) WmChangeCbChain(userFunc func(p WmChangeCbChain)) {
	me.Wm(co.WM_CHANGECBCHAIN, func(p Wm) uintptr {
		userFunc(WmChangeCbChain{m: p})
		return 0
	})
}

type WmChangeCbChain struct{ m Wm }

func (p WmChangeCbChain) WindowBeingRemoved() win.HWND { return win.HWND(p.m.WParam) }
func (p WmChangeCbChain) NextWindow() win.HWND         { return win.HWND(p.m.LParam) }
func (p WmChangeCbChain) IsLastWindow() bool           { return p.m.LParam == 0 }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (me *_DepotMsg) WmChar(userFunc func(p WmChar)) {
	me.Wm(co.WM_CHAR, func(p Wm) uintptr {
		userFunc(WmChar{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-chartoitem
func (me *_DepotMsg) WmCharToItem(userFunc func(p WmCharToItem)) {
	me.Wm(co.WM_CHARTOITEM, func(p Wm) uintptr {
		userFunc(WmCharToItem{m: p})
		return 0
	})
}

type WmCharToItem struct{ m Wm }

func (p WmCharToItem) CharCode() rune        { return rune(p.m.WParam.LoWord()) }
func (p WmCharToItem) CurrentCaretPos() uint { return uint(p.m.WParam.HiWord()) }
func (p WmCharToItem) HwndListBox() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-childactivate
func (me *_DepotMsg) WmChildActivate(userFunc func()) {
	me.Wm(co.WM_CHILDACTIVATE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-clipboardupdate
func (me *_DepotMsg) WmClipboardUpdate(userFunc func()) {
	me.Wm(co.WM_CLIPBOARDUPDATE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-close
//
// Warning: default handled in WindowModal.
func (me *_DepotMsg) WmClose(userFunc func()) {
	me.Wm(co.WM_CLOSE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-compareitem
func (me *_DepotMsg) WmCompareItem(userFunc func(p WmCompareItem) int) {
	me.Wm(co.WM_COMPAREITEM, func(p Wm) uintptr {
		return uintptr(userFunc(WmCompareItem{m: p}))
	})
}

type WmCompareItem struct{ m Wm }

func (p WmCompareItem) ControlId() int { return int(p.m.WParam) }
func (p WmCompareItem) CompareItemStruct() *win.COMPAREITEMSTRUCT {
	return (*win.COMPAREITEMSTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-contextmenu
func (me *_DepotMsg) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.Wm(co.WM_CONTEXTMENU, func(p Wm) uintptr {
		userFunc(WmContextMenu{m: p})
		return 0
	})
}

type WmContextMenu struct{ m Wm }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.m.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.m.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-copydata
func (me *_DepotMsg) WmCopyData(userFunc func(p WmCopyData) bool) {
	me.Wm(co.WM_COPYDATA, func(p Wm) uintptr {
		return _Ui.BoolToUintptr(userFunc(WmCopyData{m: p}))
	})
}

type WmCopyData struct{ m Wm }

func (p WmCopyData) WindowPassingData() win.HWND { return win.HWND(p.m.WParam) }
func (p WmCopyData) CopyDataStruct() *win.COPYDATASTRUCT {
	return (*win.COPYDATASTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-create
func (me *_DepotMsg) WmCreate(userFunc func(p *win.CREATESTRUCT) int) {
	me.Wm(co.WM_CREATE, func(p Wm) uintptr {
		return uintptr(userFunc((*win.CREATESTRUCT)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorbtn
func (me *_DepotMsg) WmCtlColorBtn(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORBTN, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-ctlcolordlg
func (me *_DepotMsg) WmCtlColorDlg(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORDLG, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcoloredit
func (me *_DepotMsg) WmCtlColorEdit(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLOREDIT, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorlistbox
func (me *_DepotMsg) WmCtlColorListBox(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORLISTBOX, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorscrollbar
func (me *_DepotMsg) WmCtlColorScrollBar(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORSCROLLBAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
func (me *_DepotMsg) WmCtlColorStatic(userFunc func(p WmCtlColor) win.HBRUSH) {
	me.Wm(co.WM_CTLCOLORSTATIC, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColor{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
func (me *_DepotMsg) WmDeadChar(userFunc func(p WmChar)) {
	me.Wm(co.WM_DEADCHAR, func(p Wm) uintptr {
		userFunc(WmChar{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-deleteitem
func (me *_DepotMsg) WmDeleteItem(userFunc func(p WmDeleteItem)) {
	me.Wm(co.WM_DELETEITEM, func(p Wm) uintptr {
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
func (me *_DepotMsg) WmDestroy(userFunc func()) {
	me.Wm(co.WM_DESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-destroyclipboard
func (me *_DepotMsg) WmDestroyClipboard(userFunc func()) {
	me.Wm(co.WM_DESTROYCLIPBOARD, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-devmodechange
func (me *_DepotMsg) WmDevModeChange(userFunc func(deviceName string)) {
	me.Wm(co.WM_DEVMODECHANGE, func(p Wm) uintptr {
		userFunc(win.Str.FromUint16Ptr((*uint16)(unsafe.Pointer(p.LParam))))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-displaychange
func (me *_DepotMsg) WmDisplayChange(userFunc func(p WmDisplayChange)) {
	me.Wm(co.WM_DISPLAYCHANGE, func(p Wm) uintptr {
		userFunc(WmDisplayChange{m: p})
		return 0
	})
}

type WmDisplayChange struct{ m Wm }

func (p WmDisplayChange) BitsPerPixel() uint { return uint(p.m.WParam) }
func (p WmDisplayChange) Size() win.SIZE     { return p.m.LParam.MakeSize() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-drawclipboard
func (me *_DepotMsg) WmDrawClipboard(userFunc func()) {
	me.Wm(co.WM_DRAWCLIPBOARD, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-drawitem
func (me *_DepotMsg) WmDrawItem(userFunc func(p WmDrawItem)) {
	me.Wm(co.WM_DRAWITEM, func(p Wm) uintptr {
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
func (me *_DepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.Wm(co.WM_DROPFILES, func(p Wm) uintptr {
		userFunc(WmDropFiles{m: p})
		return 0
	})
}

type WmDropFiles struct{ m Wm }

func (p WmDropFiles) Hdrop() win.HDROP { return win.HDROP(p.m.WParam) }
func (p WmDropFiles) Count() uint      { return uint(p.Hdrop().DragQueryFile(0xFFFF_FFFF, nil, 0)) }

// Calls DragQueryFile() successively to retrieve all file names, and releases
// the HDROP handle with DragFinish().
func (p WmDropFiles) RetrieveAll() []string {
	count := uint32(p.Count())
	files := make([]string, 0, count)
	for i := uint32(0); i < count; i++ {
		pathLen := p.Hdrop().DragQueryFile(i, nil, 0) + 1 // room for terminating null
		pathBuf := make([]uint16, pathLen)
		p.Hdrop().DragQueryFile(i, &pathBuf[0], pathLen)
		files = append(files, syscall.UTF16ToString(pathBuf))
	}
	p.Hdrop().DragFinish()
	sort.Slice(files, func(i, j int) bool { // case insensitive
		return strings.ToUpper(files[i]) < strings.ToUpper(files[j])
	})
	return files
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-enable
func (me *_DepotMsg) WmEnable(userFunc func(hasBeenEnabled bool)) {
	me.Wm(co.WM_ENABLE, func(p Wm) uintptr {
		userFunc(p.WParam != 0)
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
func (me *_DepotMsg) WmEndSession(userFunc func(p WmEndSession)) {
	me.Wm(co.WM_ENDSESSION, func(p Wm) uintptr {
		userFunc(WmEndSession{m: p})
		return 0
	})
}

type WmEndSession struct{ m Wm }

func (p WmEndSession) IsSessionBeingEnded() bool { return p.m.WParam != 0 }
func (p WmEndSession) Event() co.ENDSESSION      { return co.ENDSESSION(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-enteridle
func (me *_DepotMsg) WmEnterIdle(userFunc func(p WmEnterIdle)) {
	me.Wm(co.WM_ENTERIDLE, func(p Wm) uintptr {
		userFunc(WmEnterIdle{m: p})
		return 0
	})
}

type WmEnterIdle struct{ m Wm }

func (p WmEnterIdle) Displayed() co.MSGF       { return co.MSGF(p.m.WParam) }
func (p WmEnterIdle) DialogOrWindow() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-entermenuloop
func (me *_DepotMsg) WmEnterMenuLoop(userFunc func(isTrackPopupMenu bool)) {
	me.Wm(co.WM_ENTERMENULOOP, func(p Wm) uintptr {
		userFunc(p.WParam != 0)
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-entersizemove
func (me *_DepotMsg) WmEnterSizeMove(userFunc func()) {
	me.Wm(co.WM_ENTERSIZEMOVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-erasebkgnd
func (me *_DepotMsg) WmEraseBkgnd(userFunc func(hdc win.HDC) int) {
	me.Wm(co.WM_ERASEBKGND, func(p Wm) uintptr {
		return uintptr(userFunc(win.HDC(p.WParam)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-exitmenuloop
func (me *_DepotMsg) WmExitMenuLoop(userFunc func(isShortcutMenu bool)) {
	me.Wm(co.WM_EXITMENULOOP, func(p Wm) uintptr {
		userFunc(p.WParam != 0)
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-exitsizemove
func (me *_DepotMsg) WmExitSizeMove(userFunc func()) {
	me.Wm(co.WM_EXITSIZEMOVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-fontchange
func (me *_DepotMsg) WmFontChange(userFunc func()) {
	me.Wm(co.WM_FONTCHANGE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
func (me *_DepotMsg) WmGetDlgCode(userFunc func(p WmGetDlgCode) co.DLGC) {
	me.Wm(co.WM_GETDLGCODE, func(p Wm) uintptr {
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
func (me *_DepotMsg) WmGetFont(userFunc func() win.HFONT) {
	me.Wm(co.WM_FONTCHANGE, func(p Wm) uintptr {
		return uintptr(userFunc())
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-gettitlebarinfoex
func (me *_DepotMsg) WmGetTitleBarInfoEx(userFunc func(p *win.TITLEBARINFOEX)) {
	me.Wm(co.WM_GETTITLEBARINFOEX, func(p Wm) uintptr {
		userFunc((*win.TITLEBARINFOEX)(unsafe.Pointer(p.LParam)))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/shell/wm-help
func (me *_DepotMsg) WmHelp(userFunc func(p *win.HELPINFO)) {
	me.Wm(co.WM_HELP, func(p Wm) uintptr {
		userFunc((*win.HELPINFO)(unsafe.Pointer(p.LParam)))
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
func (me *_DepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.Wm(co.WM_HOTKEY, func(p Wm) uintptr {
		userFunc(WmHotKey{m: p})
		return 0
	})
}

type WmHotKey struct{ m Wm }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.m.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.m.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.m.LParam.HiWord()) }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll
func (me *_DepotMsg) WmHScroll(userFunc func(p WmScroll)) {
	me.Wm(co.WM_HSCROLL, func(p Wm) uintptr {
		userFunc(WmScroll{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-hscrollclipboard
func (me *_DepotMsg) WmHScrollClipboard(userFunc func(p WmScroll)) {
	me.Wm(co.WM_HSCROLLCLIPBOARD, func(p Wm) uintptr {
		userFunc(WmScroll{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
func (me *_DepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.Wm(co.WM_INITMENUPOPUP, func(p Wm) uintptr {
		userFunc(WmInitMenuPopup{m: p})
		return 0
	})
}

type WmInitMenuPopup struct{ m Wm }

func (p WmInitMenuPopup) Hmenu() win.HMENU   { return win.HMENU(p.m.WParam) }
func (p WmInitMenuPopup) Pos() uint          { return uint(p.m.LParam.LoWord()) }
func (p WmInitMenuPopup) IsWindowMenu() bool { return p.m.LParam.HiWord() != 0 }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keydown
func (me *_DepotMsg) WmKeyDown(userFunc func(p WmKey)) {
	me.Wm(co.WM_KEYDOWN, func(p Wm) uintptr {
		userFunc(WmKey{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keyup
func (me *_DepotMsg) WmKeyUp(userFunc func(p WmKey)) {
	me.Wm(co.WM_KEYUP, func(p Wm) uintptr {
		userFunc(WmKey{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-killfocus
func (me *_DepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.Wm(co.WM_KILLFOCUS, func(p Wm) uintptr {
		userFunc(WmKillFocus{m: p})
		return 0
	})
}

type WmKillFocus struct{ m Wm }

func (p WmKillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondblclk
func (me *_DepotMsg) WmLButtonDblClk(userFunc func(p WmMouse)) {
	me.Wm(co.WM_LBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
func (me *_DepotMsg) WmLButtonDown(userFunc func(p WmMouse)) {
	me.Wm(co.WM_LBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttonup
func (me *_DepotMsg) WmLButtonUp(userFunc func(p WmMouse)) {
	me.Wm(co.WM_LBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondblclk
func (me *_DepotMsg) WmMButtonDblClk(userFunc func(p WmMouse)) {
	me.Wm(co.WM_MBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondown
func (me *_DepotMsg) WmMButtonDown(userFunc func(p WmMouse)) {
	me.Wm(co.WM_MBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttonup
func (me *_DepotMsg) WmMButtonUp(userFunc func(p WmMouse)) {
	me.Wm(co.WM_MBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
func (me *_DepotMsg) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.Wm(co.WM_MENUCHAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenuChar{m: p}))
	})
}

type WmMenuChar struct{ m Wm }

func (p WmMenuChar) CharCode() rune        { return rune(p.m.WParam.LoWord()) }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.m.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
func (me *_DepotMsg) WmMenuCommand(userFunc func(p WmMenu)) {
	me.Wm(co.WM_MENUCOMMAND, func(p Wm) uintptr {
		userFunc(WmMenu{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
func (me *_DepotMsg) WmMenuDrag(userFunc func(p WmMenu) co.MND) {
	me.Wm(co.WM_MENUDRAG, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenu{m: p}))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menugetobject
func (me *_DepotMsg) WmMenuGetObject(userFunc func(p *win.MENUGETOBJECTINFO) co.MNGO) {
	me.Wm(co.WM_MENUGETOBJECT, func(p Wm) uintptr {
		return uintptr(userFunc((*win.MENUGETOBJECTINFO)(unsafe.Pointer(p.LParam))))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menurbuttonup
func (me *_DepotMsg) WmMenuRButtonUp(userFunc func(p WmMenu)) {
	me.Wm(co.WM_MENURBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMenu{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
func (me *_DepotMsg) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.Wm(co.WM_MENUSELECT, func(p Wm) uintptr {
		userFunc(WmMenuSelect{m: p})
		return 0
	})
}

type WmMenuSelect struct{ m Wm }

func (p WmMenuSelect) Item() uint       { return uint(p.m.WParam.LoWord()) }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.m.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousehover
func (me *_DepotMsg) WmMouseHover(userFunc func(p WmMouse)) {
	me.Wm(co.WM_MOUSEHOVER, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mouseleave
func (me *_DepotMsg) WmMouseLeave(userFunc func()) {
	me.Wm(co.WM_MOUSELEAVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
func (me *_DepotMsg) WmMouseMove(userFunc func(p WmMouse)) {
	me.Wm(co.WM_MOUSEMOVE, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-move
func (me *_DepotMsg) WmMove(userFunc func(clientAreaPos win.POINT)) {
	me.Wm(co.WM_MOVE, func(p Wm) uintptr {
		userFunc(p.LParam.MakePoint())
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-moving
func (me *_DepotMsg) WmMoving(userFunc func(windowPos *win.RECT)) {
	me.Wm(co.WM_MOVING, func(p Wm) uintptr {
		userFunc((*win.RECT)(unsafe.Pointer(p.LParam)))
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncactivate
func (me *_DepotMsg) WmNcActivate(userFunc func(p WmNcActivate) bool) {
	me.Wm(co.WM_NCACTIVATE, func(p Wm) uintptr {
		return _Ui.BoolToUintptr(userFunc(WmNcActivate{m: p}))
	})
}

type WmNcActivate struct{ m Wm }

func (p WmNcActivate) IsActive() bool            { return p.m.WParam != 0 }
func (p WmNcActivate) IsVisualStyleActive() bool { return p.m.LParam == 0 }
func (p WmNcActivate) UpdatedRegion() win.HRGN   { return win.HRGN(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-nccalcsize
func (me *_DepotMsg) WmNcCalcSize(userFunc func(p WmNcCalcSize) co.WVR) {
	me.Wm(co.WM_NCCALCSIZE, func(p Wm) uintptr {
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
func (me *_DepotMsg) WmNcDestroy(userFunc func()) {
	me.Wm(co.WM_NCDESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nchittest
func (me *_DepotMsg) WmNcHitTest(userFunc func(cursorCoord win.POINT) co.HT) {
	me.Wm(co.WM_NCHITTEST, func(p Wm) uintptr {
		return uintptr(userFunc(p.LParam.MakePoint()))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondblclk
func (me *_DepotMsg) WmNcLButtonDblClk(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCLBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttondown
func (me *_DepotMsg) WmNcLButtonDown(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCLBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-nclbuttonup
func (me *_DepotMsg) WmNcLButtonUp(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCLBUTTONUP, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondblclk
func (me *_DepotMsg) WmNcMButtonDblClk(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCMBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttondown
func (me *_DepotMsg) WmNcMButtonDown(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCMBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmbuttonup
func (me *_DepotMsg) WmNcMButtonUp(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCMBUTTONUP, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousehover
func (me *_DepotMsg) WmNcMouseHover(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCMOUSEHOVER, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmouseleave
func (me *_DepotMsg) WmNcMouseLeave(userFunc func()) {
	me.Wm(co.WM_NCMOUSELEAVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncmousemove
func (me *_DepotMsg) WmNcMouseMove(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCMOUSEMOVE, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-ncpaint
//
// Warning: default handled in WindowControl.
func (me *_DepotMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.Wm(co.WM_NCPAINT, func(p Wm) uintptr {
		userFunc(WmNcPaint{m: p})
		return 0
	})
}

type WmNcPaint struct{ m Wm }

func (p WmNcPaint) Raw() Wm                 { return p.m }
func (p WmNcPaint) UpdatedRegion() win.HRGN { return win.HRGN(p.m.WParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondblclk
func (me *_DepotMsg) WmNcRButtonDblClk(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCRBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttondown
func (me *_DepotMsg) WmNcRButtonDown(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCRBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncrbuttonup
func (me *_DepotMsg) WmNcRButtonUp(userFunc func(p WmNcMouse)) {
	me.Wm(co.WM_NCRBUTTONUP, func(p Wm) uintptr {
		userFunc(WmNcMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondblclk
func (me *_DepotMsg) WmNcXButtonDblClk(userFunc func(p WmNcMouseX)) {
	me.Wm(co.WM_NCXBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmNcMouseX{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttondown
func (me *_DepotMsg) WmNcXButtonDown(userFunc func(p WmNcMouseX)) {
	me.Wm(co.WM_NCXBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmNcMouseX{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-ncxbuttonup
func (me *_DepotMsg) WmNcXButtonUp(userFunc func(p WmNcMouseX)) {
	me.Wm(co.WM_NCXBUTTONUP, func(p Wm) uintptr {
		userFunc(WmNcMouseX{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-nextmenu
func (me *_DepotMsg) WmNextMenu(userFunc func(p WmNextMenu)) {
	me.Wm(co.WM_NEXTMENU, func(p Wm) uintptr {
		userFunc(WmNextMenu{m: p})
		return 0
	})
}

type WmNextMenu struct{ m Wm }

func (p WmNextMenu) VirtualKeyCode() co.VK { return co.VK(p.m.WParam) }
func (p WmNextMenu) MdiNextMenu() *win.MDINEXTMENU {
	return (*win.MDINEXTMENU)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-paint
func (me *_DepotMsg) WmPaint(userFunc func()) {
	me.Wm(co.WM_PAINT, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-paintclipboard
func (me *_DepotMsg) WmPaintClipboard(userFunc func(WmPaintClipboard)) {
	me.Wm(co.WM_PAINTCLIPBOARD, func(p Wm) uintptr {
		userFunc(WmPaintClipboard{m: p})
		return 0
	})
}

type WmPaintClipboard struct{ m Wm }

func (p WmPaintClipboard) CbViewerWindow() win.HWND { return win.HWND(p.m.WParam) }
func (p WmPaintClipboard) PaintStruct() *win.PAINTSTRUCT {
	return (*win.PAINTSTRUCT)(unsafe.Pointer(p.m.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
func (me *_DepotMsg) WmPowerBroadcast(userFunc func(p WmPowerBroadcast)) {
	me.Wm(co.WM_POWERBROADCAST, func(p Wm) uintptr {
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
func (me *_DepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.Wm(co.WM_PRINT, func(p Wm) uintptr {
		userFunc(WmPrint{m: p})
		return 0
	})
}

type WmPrint struct{ m Wm }

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.m.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.m.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondblclk
func (me *_DepotMsg) WmRButtonDblClk(userFunc func(p WmMouse)) {
	me.Wm(co.WM_RBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondown
func (me *_DepotMsg) WmRButtonDown(userFunc func(p WmMouse)) {
	me.Wm(co.WM_RBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttonup
func (me *_DepotMsg) WmRButtonUp(userFunc func(p WmMouse)) {
	me.Wm(co.WM_RBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-renderallformats
func (me *_DepotMsg) WmRenderAllFormats(userFunc func()) {
	me.Wm(co.WM_RENDERALLFORMATS, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-renderformat
func (me *_DepotMsg) WmRenderFormat(userFunc func(clipboardFormat co.CF)) {
	me.Wm(co.WM_RENDERFORMAT, func(p Wm) uintptr {
		userFunc(co.CF(p.WParam))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-setfocus
//
// Warning: default handled in WindowMain and WindowModal.
func (me *_DepotMsg) WmSetFocus(userFunc func(hwndLosingFocus win.HWND)) {
	me.Wm(co.WM_SETFOCUS, func(p Wm) uintptr {
		userFunc(win.HWND(p.LParam))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-setfont
func (me *_DepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.Wm(co.WM_SETFONT, func(p Wm) uintptr {
		userFunc(WmSetFont{m: p})
		return 0
	})
}

type WmSetFont struct{ m Wm }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.m.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.m.LParam == 1 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-size
func (me *_DepotMsg) WmSize(userFunc func(p WmSize)) {
	me.Wm(co.WM_SIZE, func(p Wm) uintptr {
		userFunc(WmSize{m: p})
		return 0
	})
}

type WmSize struct{ m Wm }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.m.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.m.LParam.MakeSize() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-sizeclipboard
func (me *_DepotMsg) WmSizeClipboard(userFunc func(p WmSizeClipboard)) {
	me.Wm(co.WM_SIZECLIPBOARD, func(p Wm) uintptr {
		userFunc(WmSizeClipboard{m: p})
		return 0
	})
}

type WmSizeClipboard struct{ m Wm }

func (p WmSizeClipboard) CbViewerWindow() win.HWND { return win.HWND(p.m.WParam) }
func (p WmSizeClipboard) NewDimensions() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.m.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syschar
func (me *_DepotMsg) WmSysChar(userFunc func(p WmChar)) {
	me.Wm(co.WM_SYSCHAR, func(p Wm) uintptr {
		userFunc(WmChar{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
func (me *_DepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.Wm(co.WM_SYSCOMMAND, func(p Wm) uintptr {
		userFunc(WmSysCommand{m: p})
		return 0
	})
}

type WmSysCommand struct{ m Wm }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.m.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.m.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-sysdeadchar
func (me *_DepotMsg) WmSysDeadChar(userFunc func(p WmChar)) {
	me.Wm(co.WM_SYSDEADCHAR, func(p Wm) uintptr {
		userFunc(WmChar{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeydown
func (me *_DepotMsg) WmSysKeyDown(userFunc func(p WmKey)) {
	me.Wm(co.WM_SYSKEYDOWN, func(p Wm) uintptr {
		userFunc(WmKey{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeyup
func (me *_DepotMsg) WmSysKeyUp(userFunc func(p WmKey)) {
	me.Wm(co.WM_SYSKEYUP, func(p Wm) uintptr {
		userFunc(WmKey{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/sysinfo/wm-timechange
func (me *_DepotMsg) WmTimeChange(userFunc func()) {
	me.Wm(co.WM_TIMECHANGE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-uninitmenupopup
func (me *_DepotMsg) WmUnInitMenuPopup(userFunc func(menu win.HMENU)) {
	me.Wm(co.WM_UNINITMENUPOPUP, func(p Wm) uintptr {
		userFunc(win.HMENU(p.WParam))
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-vscroll
func (me *_DepotMsg) WmVScroll(userFunc func(p WmScroll)) {
	me.Wm(co.WM_VSCROLL, func(p Wm) uintptr {
		userFunc(WmScroll{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-vscrollclipboard
func (me *_DepotMsg) WmVScrollClipboard(userFunc func(p WmScroll)) {
	me.Wm(co.WM_VSCROLLCLIPBOARD, func(p Wm) uintptr {
		userFunc(WmScroll{m: p})
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondblclk
func (me *_DepotMsg) WmXButtonDblClk(userFunc func(p WmMouse)) {
	me.Wm(co.WM_XBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttondown
func (me *_DepotMsg) WmXButtonDown(userFunc func(p WmMouse)) {
	me.Wm(co.WM_XBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 1
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-xbuttonup
func (me *_DepotMsg) WmXButtonUp(userFunc func(p WmMouse)) {
	me.Wm(co.WM_XBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMouse{m: p})
		return 1
	})
}
