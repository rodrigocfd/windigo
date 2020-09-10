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

// Keeps all user message handlers.
type _DepotMsg struct {
	mapMsgs map[co.WM]func(p Wm) uintptr
	mapCmds map[int]func(p WmCommand)
}

func (me *_DepotMsg) addMsg(msg co.WM, userFunc func(p Wm) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p Wm) uintptr, 16) // arbitrary capacity
	}
	me.mapMsgs[msg] = userFunc
}

func (me *_DepotMsg) addCmd(cmd int, userFunc func(p WmCommand)) {
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[int]func(p WmCommand), 4) // arbitrary capacity
	}
	me.mapCmds[cmd] = userFunc
}

func (me *_DepotMsg) processMessage(msg co.WM, p Wm) (uintptr, bool) {
	if msg == co.WM_COMMAND {
		pCmd := WmCommand(p)
		if userFunc, hasCmd := me.mapCmds[pCmd.ControlId()]; hasCmd {
			userFunc(pCmd)
			return 0, true // always return zero
		}
	} else {
		if userFunc, hasMsg := me.mapMsgs[msg]; hasMsg {
			return userFunc(p), true
		}
	}

	return 0, false // no user handler found
}

//------------------------------------------------------------------------------

// Handles a raw, unspecific window message. There will be no treatment of
// WPARAM/LPARAM data, you'll have to unpack all the information manually. This
// is very dangerous.
//
// Unless you have a very good reason, always prefer the specific message
// handlers.
func (me *_DepotMsg) Wm(message co.WM, userFunc func(p Wm) uintptr) {
	me.addMsg(message, userFunc)
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-activate
//
// Warning: default handled in WindowMain.
func (me *_DepotMsg) WmActivate(userFunc func(p WmActivate)) {
	me.addMsg(co.WM_ACTIVATE, func(p Wm) uintptr {
		userFunc(WmActivate(p))
		return 0
	})
}

type WmActivate struct{ _Wm }

func (p WmActivate) Event() co.WA                           { return co.WA(p.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                      { return p.WParam.HiWord() != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-activateapp
func (me *_DepotMsg) WmActivateApp(userFunc func(p WmActivateApp)) {
	me.addMsg(co.WM_ACTIVATEAPP, func(p Wm) uintptr {
		userFunc(WmActivateApp(p))
		return 0
	})
}

type WmActivateApp struct{ _Wm }

func (p WmActivateApp) IsBeingActivated() bool { return p.WParam != 0 }
func (p WmActivateApp) ThreadId() uint         { return uint(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
func (me *_DepotMsg) WmAppCommand(userFunc func(p WmAppCommand)) {
	me.addMsg(co.WM_APPCOMMAND, func(p Wm) uintptr {
		userFunc(WmAppCommand(p))
		return 1
	})
}

type WmAppCommand struct{ _Wm }

func (p WmAppCommand) OwnerWindow() win.HWND     { return win.HWND(p.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.LParam.HiWord() &^ 0xF000) }
func (p WmAppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.LParam.HiWord() & 0xF000) }
func (p WmAppCommand) Keys() co.MK               { return co.MK(p.LParam.LoWord()) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-capturechanged
func (me *_DepotMsg) WmCaptureChanged(userFunc func(p WmCaptureChanged)) {
	me.addMsg(co.WM_CAPTURECHANGED, func(p Wm) uintptr {
		userFunc(WmCaptureChanged(p))
		return 0
	})
}

type WmCaptureChanged struct{ _Wm }

func (p WmCaptureChanged) WindowGainingMouse() win.HWND { return win.HWND(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-changecbchain
func (me *_DepotMsg) WmChangeCbChain(userFunc func(p WmChangeCbChain)) {
	me.addMsg(co.WM_CHANGECBCHAIN, func(p Wm) uintptr {
		userFunc(WmChangeCbChain(p))
		return 0
	})
}

type WmChangeCbChain struct{ _Wm }

func (p WmChangeCbChain) WindowBeingRemoved() win.HWND { return win.HWND(p.WParam) }
func (p WmChangeCbChain) NextWindow() win.HWND         { return win.HWND(p.LParam) }
func (p WmChangeCbChain) IsLastWindow() bool           { return p.LParam == 0 }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (me *_DepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p Wm) uintptr {
		userFunc(WmChar{_WmChar(Wm(p))})
		return 0
	})
}

type WmChar struct{ _WmChar }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-chartoitem
func (me *_DepotMsg) WmCharToItem(userFunc func(p WmCharToItem)) {
	me.addMsg(co.WM_CHARTOITEM, func(p Wm) uintptr {
		userFunc(WmCharToItem(p))
		return 0
	})
}

type WmCharToItem struct{ _Wm }

func (p WmCharToItem) CharCode() rune        { return rune(p.WParam.LoWord()) }
func (p WmCharToItem) CurrentCaretPos() uint { return uint(p.WParam.HiWord()) }
func (p WmCharToItem) HwndListBox() win.HWND { return win.HWND(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-childactivate
func (me *_DepotMsg) WmChildActivate(userFunc func()) {
	me.addMsg(co.WM_CHILDACTIVATE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-close
//
// Warning: default handled in WindowModal.
func (me *_DepotMsg) WmClose(userFunc func()) {
	me.addMsg(co.WM_CLOSE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
func (me *_DepotMsg) WmCommand(cmd int, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

type WmCommand struct{ _Wm }

func (p WmCommand) IsFromMenu() bool        { return p.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool { return p.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool     { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() int             { return p.ControlId() }
func (p WmCommand) AcceleratorId() int      { return p.ControlId() }
func (p WmCommand) ControlId() int          { return int(p.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() int   { return int(p.WParam.HiWord()) }
func (p WmCommand) ControlHwnd() win.HWND   { return win.HWND(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-compareitem
func (me *_DepotMsg) WmCompareItem(userFunc func(p WmCompareItem) int) {
	me.addMsg(co.WM_COMPAREITEM, func(p Wm) uintptr {
		return uintptr(userFunc(WmCompareItem(p)))
	})
}

type WmCompareItem struct{ _Wm }

func (p WmCompareItem) ControlId() int { return int(p.WParam) }
func (p WmCompareItem) CompareItemStruct() *win.COMPAREITEMSTRUCT {
	return (*win.COMPAREITEMSTRUCT)(unsafe.Pointer(p.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-contextmenu
func (me *_DepotMsg) WmContextMenu(userFunc func(p WmContextMenu)) {
	me.addMsg(co.WM_CONTEXTMENU, func(p Wm) uintptr {
		userFunc(WmContextMenu(p))
		return 0
	})
}

type WmContextMenu struct{ _Wm }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-create
func (me *_DepotMsg) WmCreate(userFunc func(p WmCreate) int) {
	me.addMsg(co.WM_CREATE, func(p Wm) uintptr {
		return uintptr(userFunc(WmCreate(p)))
	})
}

type WmCreate struct{ _Wm }

func (p WmCreate) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorbtn
func (me *_DepotMsg) WmCtlColorBtn(userFunc func(p WmCtlColorBtn) win.HBRUSH) {
	me.addMsg(co.WM_CTLCOLORBTN, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorBtn{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorBtn struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-ctlcolordlg
func (me *_DepotMsg) WmCtlColorDlg(userFunc func(p WmCtlColorDlg) win.HBRUSH) {
	me.addMsg(co.WM_CTLCOLORDLG, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorDlg{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorDlg struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcoloredit
func (me *_DepotMsg) WmCtlColorEdit(userFunc func(p WmCtlColorEdit) win.HBRUSH) {
	me.addMsg(co.WM_CTLCOLOREDIT, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorEdit{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorEdit struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorlistbox
func (me *_DepotMsg) WmCtlColorListBox(userFunc func(p WmCtlColorListBox) win.HBRUSH) {
	me.addMsg(co.WM_CTLCOLORLISTBOX, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorListBox{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorListBox struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorscrollbar
func (me *_DepotMsg) WmCtlColorScrollBar(userFunc func(p WmCtlColorScrollBar) win.HBRUSH) {
	me.addMsg(co.WM_CTLCOLORSCROLLBAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorScrollBar{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorScrollBar struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-ctlcolorstatic
func (me *_DepotMsg) WmCtlColorStatic(userFunc func(p WmCtlColorStatic) win.HBRUSH) {
	me.addMsg(co.WM_CTLCOLORSTATIC, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorStatic{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorStatic struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
func (me *_DepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p Wm) uintptr {
		userFunc(WmDeadChar{_WmChar(Wm(p))})
		return 0
	})
}

type WmDeadChar struct{ _WmChar }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-deleteitem
func (me *_DepotMsg) WmDeleteItem(userFunc func(p WmDeleteItem)) {
	me.addMsg(co.WM_DELETEITEM, func(p Wm) uintptr {
		userFunc(WmDeleteItem(p))
		return 1
	})
}

type WmDeleteItem struct{ _Wm }

func (p WmDeleteItem) ControlId() int { return int(p.WParam) }
func (p WmDeleteItem) DeleteItemStruct() *win.DELETEITEMSTRUCT {
	return (*win.DELETEITEMSTRUCT)(unsafe.Pointer(p.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
func (me *_DepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-destroyclipboard
func (me *_DepotMsg) WmDestroyClipboard(userFunc func()) {
	me.addMsg(co.WM_DESTROYCLIPBOARD, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-displaychange
func (me *_DepotMsg) WmDisplayChange(userFunc func(p WmDisplayChange)) {
	me.addMsg(co.WM_DISPLAYCHANGE, func(p Wm) uintptr {
		userFunc(WmDisplayChange(p))
		return 0
	})
}

type WmDisplayChange struct{ _Wm }

func (p WmDisplayChange) BitsPerPixel() uint { return uint(p.WParam) }
func (p WmDisplayChange) Size() win.SIZE     { return p.LParam.MakeSize() }

// https://docs.microsoft.com/en-us/windows/win32/dataxchg/wm-drawclipboard
func (me *_DepotMsg) WmDrawClipboard(userFunc func()) {
	me.addMsg(co.WM_DRAWCLIPBOARD, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-drawitem
func (me *_DepotMsg) WmDrawItem(userFunc func(p WmDrawItem)) {
	me.addMsg(co.WM_DRAWITEM, func(p Wm) uintptr {
		userFunc(WmDrawItem(p))
		return 1
	})
}

type WmDrawItem struct{ _Wm }

func (p WmDrawItem) ControlId() int   { return int(p.WParam) }
func (p WmDrawItem) IsFromMenu() bool { return p.WParam == 0 }
func (p WmDrawItem) DrawItemStruct() *win.DRAWITEMSTRUCT {
	return (*win.DRAWITEMSTRUCT)(unsafe.Pointer(p.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/shell/wm-dropfiles
func (me *_DepotMsg) WmDropFiles(userFunc func(p WmDropFiles)) {
	me.addMsg(co.WM_DROPFILES, func(p Wm) uintptr {
		userFunc(WmDropFiles(p))
		return 0
	})
}

type WmDropFiles struct{ _Wm }

func (p WmDropFiles) Hdrop() win.HDROP { return win.HDROP(p.WParam) }

// Calls DragQueryFile successively to retrieve all file names, and releases the
// HDROP handle.
func (p WmDropFiles) RetrieveAll() []string {
	count := p.Hdrop().DragQueryFile(0xFFFF_FFFF, nil, 0)
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
func (me *_DepotMsg) WmEnable(userFunc func(p WmEnable)) {
	me.addMsg(co.WM_ENABLE, func(p Wm) uintptr {
		userFunc(WmEnable(p))
		return 0
	})
}

type WmEnable struct{ _Wm }

func (p WmEnable) Enabled() bool { return p.WParam != 0 }

// https://docs.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
func (me *_DepotMsg) WmEndSession(userFunc func(p WmEndSession)) {
	me.addMsg(co.WM_ENDSESSION, func(p Wm) uintptr {
		userFunc(WmEndSession(p))
		return 0
	})
}

type WmEndSession struct{ _Wm }

func (p WmEndSession) IsSessionBeingEnded() bool { return p.WParam != 0 }
func (p WmEndSession) Event() co.ENDSESSION      { return co.ENDSESSION(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-entermenuloop
func (me *_DepotMsg) WmEnterMenuLoop(userFunc func(p WmEnterMenuLoop)) {
	me.addMsg(co.WM_ENTERMENULOOP, func(p Wm) uintptr {
		userFunc(WmEnterMenuLoop(p))
		return 0
	})
}

type WmEnterMenuLoop struct{ _Wm }

func (p WmEnterMenuLoop) IsTrackPopupMenu() bool { return p.WParam != 0 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-entersizemove
func (me *_DepotMsg) WmEnterSizeMove(userFunc func()) {
	me.addMsg(co.WM_ENTERSIZEMOVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-erasebkgnd
func (me *_DepotMsg) WmEraseBkGnd(userFunc func(p WmEraseBkGnd) int) {
	me.addMsg(co.WM_ERASEBKGND, func(p Wm) uintptr {
		return uintptr(userFunc(WmEraseBkGnd(p)))
	})
}

type WmEraseBkGnd struct{ _Wm }

func (p WmEraseBkGnd) Hdc() win.HDC { return win.HDC(p.WParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-exitmenuloop
func (me *_DepotMsg) WmExitMenuLoop(userFunc func(p WmExitMenuLoop)) {
	me.addMsg(co.WM_EXITMENULOOP, func(p Wm) uintptr {
		userFunc(WmExitMenuLoop(p))
		return 0
	})
}

type WmExitMenuLoop struct{ _Wm }

func (p WmExitMenuLoop) IsShortcutMenu() bool { return p.WParam != 0 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-exitsizemove
func (me *_DepotMsg) WmExitSizeMove(userFunc func()) {
	me.addMsg(co.WM_EXITSIZEMOVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-fontchange
func (me *_DepotMsg) WmFontChange(userFunc func()) {
	me.addMsg(co.WM_FONTCHANGE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/dlgbox/wm-getdlgcode
func (me *_DepotMsg) WmGetDlgCode(userFunc func(p WmGetDlgCode) co.DLGC) {
	me.addMsg(co.WM_GETDLGCODE, func(p Wm) uintptr {
		return uintptr(userFunc(WmGetDlgCode(p)))
	})
}

type WmGetDlgCode struct{ _Wm }

func (p WmGetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.WParam) }
func (p WmGetDlgCode) IsQuery() bool         { return p.LParam == 0 }
func (p WmGetDlgCode) Msg() *win.MSG         { return (*win.MSG)(unsafe.Pointer(p.LParam)) }
func (p WmGetDlgCode) HasAlt() bool          { return (win.GetAsyncKeyState(co.VK_MENU) & 0x8000) != 0 }
func (p WmGetDlgCode) HasCtrl() bool         { return (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0 }
func (p WmGetDlgCode) HasShift() bool        { return (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-getfont
func (me *_DepotMsg) WmGetFont(userFunc func() win.HFONT) {
	me.addMsg(co.WM_FONTCHANGE, func(p Wm) uintptr {
		return uintptr(userFunc())
	})
}

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-gettitlebarinfoex
func (me *_DepotMsg) WmGetTitleBarInfoEx(userFunc func(WmGetTitleBarInfoEx)) {
	me.addMsg(co.WM_GETTITLEBARINFOEX, func(p Wm) uintptr {
		userFunc(WmGetTitleBarInfoEx(p))
		return 0
	})
}

type WmGetTitleBarInfoEx struct{ _Wm }

func (p WmHelp) TitleBarInfoEx() *win.TITLEBARINFOEX {
	return (*win.TITLEBARINFOEX)(unsafe.Pointer(p.LParam))
}

// https://docs.microsoft.com/en-us/windows/win32/shell/wm-help
func (me *_DepotMsg) WmHelp(userFunc func(p WmHelp)) {
	me.addMsg(co.WM_HELP, func(p Wm) uintptr {
		userFunc(WmHelp(p))
		return 1
	})
}

type WmHelp struct{ _Wm }

func (p WmHelp) HelpInfo() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-hotkey
func (me *_DepotMsg) WmHotKey(userFunc func(p WmHotKey)) {
	me.addMsg(co.WM_HOTKEY, func(p Wm) uintptr {
		userFunc(WmHotKey(p))
		return 0
	})
}

type WmHotKey struct{ _Wm }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.LParam.HiWord()) }

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-hscroll
func (me *_DepotMsg) WmHScroll(userFunc func(p WmHScroll)) {
	me.addMsg(co.WM_HSCROLL, func(p Wm) uintptr {
		userFunc(WmHScroll{_WmScroll(Wm(p))})
		return 0
	})
}

type WmHScroll struct{ _WmScroll }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
func (me *_DepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p Wm) uintptr {
		userFunc(WmInitMenuPopup(p))
		return 0
	})
}

type WmInitMenuPopup struct{ _Wm }

func (p WmInitMenuPopup) Hmenu() win.HMENU   { return win.HMENU(p.WParam) }
func (p WmInitMenuPopup) Pos() uint          { return uint(p.LParam.LoWord()) }
func (p WmInitMenuPopup) IsWindowMenu() bool { return p.LParam.HiWord() != 0 }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keydown
func (me *_DepotMsg) WmKeyDown(userFunc func(p WmKeyDown)) {
	me.addMsg(co.WM_KEYDOWN, func(p Wm) uintptr {
		userFunc(WmKeyDown{_WmKey(Wm(p))})
		return 0
	})
}

type WmKeyDown struct{ _WmKey }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-keyup
func (me *_DepotMsg) WmKeyUp(userFunc func(p WmKeyUp)) {
	me.addMsg(co.WM_KEYUP, func(p Wm) uintptr {
		userFunc(WmKeyUp{_WmKey(Wm(p))})
		return 0
	})
}

type WmKeyUp struct{ _WmKey }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-killfocus
func (me *_DepotMsg) WmKillFocus(userFunc func(p WmKillFocus)) {
	me.addMsg(co.WM_KILLFOCUS, func(p Wm) uintptr {
		userFunc(WmKillFocus(p))
		return 0
	})
}

type WmKillFocus struct{ _Wm }

func (p WmKillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.WParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondblclk
func (me *_DepotMsg) WmLButtonDblClk(userFunc func(p WmLButtonDblClk)) {
	me.addMsg(co.WM_LBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmLButtonDblClk{_WmButton(Wm(p))})
		return 0
	})
}

type WmLButtonDblClk struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttondown
func (me *_DepotMsg) WmLButtonDown(userFunc func(p WmLButtonDown)) {
	me.addMsg(co.WM_LBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmLButtonDown{_WmButton(Wm(p))})
		return 0
	})
}

type WmLButtonDown struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-lbuttonup
func (me *_DepotMsg) WmLButtonUp(userFunc func(p WmLButtonUp)) {
	me.addMsg(co.WM_LBUTTONUP, func(p Wm) uintptr {
		userFunc(WmLButtonUp{_WmButton(Wm(p))})
		return 0
	})
}

type WmLButtonUp struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondblclk
func (me *_DepotMsg) WmMButtonDblClk(userFunc func(p WmMButtonDblClk)) {
	me.addMsg(co.WM_MBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmMButtonDblClk{_WmButton(Wm(p))})
		return 0
	})
}

type WmMButtonDblClk struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttondown
func (me *_DepotMsg) WmMButtonDown(userFunc func(p WmMButtonDown)) {
	me.addMsg(co.WM_MBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmMButtonDown{_WmButton(Wm(p))})
		return 0
	})
}

type WmMButtonDown struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mbuttonup
func (me *_DepotMsg) WmMButtonUp(userFunc func(p WmMButtonUp)) {
	me.addMsg(co.WM_MBUTTONUP, func(p Wm) uintptr {
		userFunc(WmMButtonUp{_WmButton(Wm(p))})
		return 0
	})
}

type WmMButtonUp struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuchar
func (me *_DepotMsg) WmMenuChar(userFunc func(p WmMenuChar) co.MNC) {
	me.addMsg(co.WM_MENUCHAR, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenuChar(p)))
	})
}

type WmMenuChar struct{ _Wm }

func (p WmMenuChar) CharCode() rune        { return rune(p.WParam.LoWord()) }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
func (me *_DepotMsg) WmMenuCommand(userFunc func(p WmMenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p Wm) uintptr {
		userFunc(WmMenuCommand{_WmMenu(Wm(p))})
		return 0
	})
}

type WmMenuCommand struct{ _WmMenu }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menudrag
func (me *_DepotMsg) WmMenuDrag(userFunc func(p WmMenuDrag) co.MND) {
	me.addMsg(co.WM_MENUDRAG, func(p Wm) uintptr {
		return uintptr(userFunc(WmMenuDrag{_WmMenu(Wm(p))}))
	})
}

type WmMenuDrag struct{ _WmMenu }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
func (me *_DepotMsg) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p Wm) uintptr {
		userFunc(WmMenuSelect(p))
		return 0
	})
}

type WmMenuSelect struct{ _Wm }

func (p WmMenuSelect) Item() uint       { return uint(p.WParam.LoWord()) }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousehover
func (me *_DepotMsg) WmMouseHover(userFunc func(p WmMouseHover)) {
	me.addMsg(co.WM_MOUSEHOVER, func(p Wm) uintptr {
		userFunc(WmMouseHover{_WmButton(Wm(p))})
		return 0
	})
}

type WmMouseHover struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mouseleave
func (me *_DepotMsg) WmMouseLeave(userFunc func()) {
	me.addMsg(co.WM_MOUSELEAVE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-mousemove
func (me *_DepotMsg) WmMouseMove(userFunc func(p WmMouseMove)) {
	me.addMsg(co.WM_MOUSEMOVE, func(p Wm) uintptr {
		userFunc(WmMouseMove{_WmButton(Wm(p))})
		return 0
	})
}

type WmMouseMove struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-move
func (me *_DepotMsg) WmMove(userFunc func(p WmMove)) {
	me.addMsg(co.WM_MOVE, func(p Wm) uintptr {
		userFunc(WmMove(p))
		return 0
	})
}

type WmMove struct{ _Wm }

func (p WmMove) Pos() win.POINT { return p.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-moving
func (me *_DepotMsg) WmMoving(userFunc func(p WmMoving)) {
	me.addMsg(co.WM_MOVING, func(p Wm) uintptr {
		userFunc(WmMoving(p))
		return 1
	})
}

type WmMoving struct{ _Wm }

func (p WmMoving) ScreenCoords() *win.RECT { return (*win.RECT)(unsafe.Pointer(p.LParam)) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncactivate
func (me *_DepotMsg) WmNcActivate(userFunc func(p WmNcActivate) bool) {
	me.addMsg(co.WM_NCACTIVATE, func(p Wm) uintptr {
		return _Util.BoolToUintptr(userFunc(WmNcActivate(p)))
	})
}

type WmNcActivate struct{ _Wm }

func (p WmNcActivate) IsActive() bool { return p.WParam != 0 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-ncdestroy
//
// Warning: default handled in WindowMain.
func (me *_DepotMsg) WmNcDestroy(userFunc func()) {
	me.addMsg(co.WM_NCDESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-ncpaint
//
// Warning: default handled in WindowControl.
func (me *_DepotMsg) WmNcPaint(userFunc func(p WmNcPaint)) {
	me.addMsg(co.WM_NCPAINT, func(p Wm) uintptr {
		userFunc(WmNcPaint(p))
		return 0
	})
}

type WmNcPaint struct{ _Wm }

func (p WmNcPaint) Hrgn() win.HRGN { return win.HRGN(p.WParam) }

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-paint
func (me *_DepotMsg) WmPaint(userFunc func()) {
	me.addMsg(co.WM_PAINT, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
func (me *_DepotMsg) WmPowerBroadcast(userFunc func(p WmPowerBroadcast)) {
	me.addMsg(co.WM_POWERBROADCAST, func(p Wm) uintptr {
		userFunc(WmPowerBroadcast(p))
		return 1
	})
}

type WmPowerBroadcast struct{ _Wm }

func (p WmPowerBroadcast) Event() co.PBT { return co.PBT(p.WParam) }
func (p WmPowerBroadcast) PowerBroadcastSetting() *win.POWERBROADCAST_SETTING {
	if p.Event() == co.PBT_POWERSETTINGCHANGE {
		return (*win.POWERBROADCAST_SETTING)(unsafe.Pointer(p.LParam))
	}
	return nil
}

// https://docs.microsoft.com/en-us/windows/win32/gdi/wm-print
func (me *_DepotMsg) WmPrint(userFunc func(p WmPrint)) {
	me.addMsg(co.WM_PRINT, func(p Wm) uintptr {
		userFunc(WmPrint(p))
		return 0
	})
}

type WmPrint struct{ _Wm }

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondblclk
func (me *_DepotMsg) WmRButtonDblClk(userFunc func(p WmRButtonDblClk)) {
	me.addMsg(co.WM_RBUTTONDBLCLK, func(p Wm) uintptr {
		userFunc(WmRButtonDblClk{_WmButton(Wm(p))})
		return 0
	})
}

type WmRButtonDblClk struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttondown
func (me *_DepotMsg) WmRButtonDown(userFunc func(p WmRButtonDown)) {
	me.addMsg(co.WM_RBUTTONDOWN, func(p Wm) uintptr {
		userFunc(WmRButtonDown{_WmButton(Wm(p))})
		return 0
	})
}

type WmRButtonDown struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-rbuttonup
func (me *_DepotMsg) WmRButtonUp(userFunc func(p WmRButtonUp)) {
	me.addMsg(co.WM_RBUTTONUP, func(p Wm) uintptr {
		userFunc(WmRButtonUp{_WmButton(Wm(p))})
		return 0
	})
}

type WmRButtonUp struct{ _WmButton }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-setfocus
//
// Warning: default handled in WindowMain and WindowModal.
func (me *_DepotMsg) WmSetFocus(userFunc func(p WmSetFocus)) {
	me.addMsg(co.WM_SETFOCUS, func(p Wm) uintptr {
		userFunc(WmSetFocus(p))
		return 0
	})
}

type WmSetFocus struct{ _Wm }

func (p WmSetFocus) UnfocusedWindow() win.HWND { return win.HWND(p.WParam) }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-setfont
func (me *_DepotMsg) WmSetFont(userFunc func(p WmSetFont)) {
	me.addMsg(co.WM_SETFONT, func(p Wm) uintptr {
		userFunc(WmSetFont(p))
		return 0
	})
}

type WmSetFont struct{ _Wm }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.LParam == 1 }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-size
func (me *_DepotMsg) WmSize(userFunc func(p WmSize)) {
	me.addMsg(co.WM_SIZE, func(p Wm) uintptr {
		userFunc(WmSize(p))
		return 0
	})
}

type WmSize struct{ _Wm }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.LParam.MakeSize() }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syschar
func (me *_DepotMsg) WmSysChar(userFunc func(p WmSysChar)) {
	me.addMsg(co.WM_SYSCHAR, func(p Wm) uintptr {
		userFunc(WmSysChar{_WmChar(Wm(p))})
		return 0
	})
}

type WmSysChar struct{ _WmChar }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-syscommand
func (me *_DepotMsg) WmSysCommand(userFunc func(p WmSysCommand)) {
	me.addMsg(co.WM_SYSCOMMAND, func(p Wm) uintptr {
		userFunc(WmSysCommand(p))
		return 0
	})
}

type WmSysCommand struct{ _Wm }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.LParam.MakePoint() }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-sysdeadchar
func (me *_DepotMsg) WmSysDeadChar(userFunc func(p WmSysDeadChar)) {
	me.addMsg(co.WM_SYSDEADCHAR, func(p Wm) uintptr {
		userFunc(WmSysDeadChar{_WmChar(Wm(p))})
		return 0
	})
}

type WmSysDeadChar struct{ _WmChar }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeydown
func (me *_DepotMsg) WmSysKeyDown(userFunc func(p WmSysKeyDown)) {
	me.addMsg(co.WM_SYSKEYDOWN, func(p Wm) uintptr {
		userFunc(WmSysKeyDown{_WmKey(Wm(p))})
		return 0
	})
}

type WmSysKeyDown struct{ _WmKey }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-syskeyup
func (me *_DepotMsg) WmSysKeyUp(userFunc func(p WmSysKeyUp)) {
	me.addMsg(co.WM_SYSKEYUP, func(p Wm) uintptr {
		userFunc(WmSysKeyUp{_WmKey(Wm(p))})
		return 0
	})
}

type WmSysKeyUp struct{ _WmKey }

// https://docs.microsoft.com/en-us/windows/win32/sysinfo/wm-timechange
func (me *_DepotMsg) WmTimeChange(userFunc func()) {
	me.addMsg(co.WM_TIMECHANGE, func(p Wm) uintptr {
		userFunc()
		return 0
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/wm-vscroll
func (me *_DepotMsg) WmVScroll(userFunc func(p WmVScroll)) {
	me.addMsg(co.WM_VSCROLL, func(p Wm) uintptr {
		userFunc(WmVScroll{_WmScroll(Wm(p))})
		return 0
	})
}

type WmVScroll struct{ _WmScroll }
