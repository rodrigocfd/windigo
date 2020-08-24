/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package gui

import (
	"sort"
	"strings"
	"syscall"
	"unsafe"
	"wingows/co"
	"wingows/win"
)

// Keeps all user message handlers.
type _DepotMsg struct {
	mapMsgs map[co.WM]func(p Wm) uintptr
	mapCmds map[int32]func(p WmCommand)
}

func (me *_DepotMsg) addMsg(msg co.WM, userFunc func(p Wm) uintptr) {
	if me.mapMsgs == nil { // guard
		me.mapMsgs = make(map[co.WM]func(p Wm) uintptr, 16) // arbitrary capacity
	}
	me.mapMsgs[msg] = userFunc
}

func (me *_DepotMsg) addCmd(cmd int32, userFunc func(p WmCommand)) {
	if me.mapCmds == nil { // guard
		me.mapCmds = make(map[int32]func(p WmCommand), 4) // arbitrary capacity
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
func (p WmActivateApp) ThreadId() uint32       { return uint32(p.LParam) }

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

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-char
func (me *_DepotMsg) WmChar(userFunc func(p WmChar)) {
	me.addMsg(co.WM_CHAR, func(p Wm) uintptr {
		userFunc(WmChar{_WmChar(Wm(p))})
		return 0
	})
}

type WmChar struct{ _WmChar }

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
func (me *_DepotMsg) WmCommand(cmd int32, userFunc func(p WmCommand)) {
	me.addCmd(cmd, userFunc)
}

type WmCommand struct{ _Wm }

func (p WmCommand) IsFromMenu() bool         { return p.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return p.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() int32            { return p.ControlId() }
func (p WmCommand) AcceleratorId() int32     { return p.ControlId() }
func (p WmCommand) ControlId() int32         { return int32(p.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() uint16 { return p.WParam.HiWord() }
func (p WmCommand) ControlHwnd() win.HWND    { return win.HWND(p.LParam) }

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
func (me *_DepotMsg) WmCreate(userFunc func(p WmCreate) int32) {
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
	me.addMsg(co.WM_CONTEXTMENU, func(p Wm) uintptr {
		return uintptr(userFunc(WmCtlColorBtn{_WmCtlColor(Wm(p))}))
	})
}

type WmCtlColorBtn struct{ _WmCtlColor }

// https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-deadchar
func (me *_DepotMsg) WmDeadChar(userFunc func(p WmDeadChar)) {
	me.addMsg(co.WM_DEADCHAR, func(p Wm) uintptr {
		userFunc(WmDeadChar{_WmChar(Wm(p))})
		return 0
	})
}

type WmDeadChar struct{ _WmChar }

// https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-destroy
func (me *_DepotMsg) WmDestroy(userFunc func()) {
	me.addMsg(co.WM_DESTROY, func(p Wm) uintptr {
		userFunc()
		return 0
	})
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

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-initmenupopup
func (me *_DepotMsg) WmInitMenuPopup(userFunc func(p WmInitMenuPopup)) {
	me.addMsg(co.WM_INITMENUPOPUP, func(p Wm) uintptr {
		userFunc(WmInitMenuPopup(p))
		return 0
	})
}

type WmInitMenuPopup struct{ _Wm }

func (p WmInitMenuPopup) Hmenu() win.HMENU        { return win.HMENU(p.WParam) }
func (p WmInitMenuPopup) MenuRelativePos() uint16 { return p.LParam.LoWord() }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return p.LParam.HiWord() != 0 }

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

func (p WmMenuChar) CharCode() uint16      { return p.WParam.LoWord() }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menucommand
func (me *_DepotMsg) WmMenuCommand(userFunc func(p WmMenuCommand)) {
	me.addMsg(co.WM_MENUCOMMAND, func(p Wm) uintptr {
		userFunc(WmMenuCommand(p))
		return 0
	})
}

type WmMenuCommand struct{ _Wm }

func (p WmMenuCommand) ItemIndex() uint16 { return uint16(p.WParam) }
func (p WmMenuCommand) Hmenu() win.HMENU  { return win.HMENU(p.LParam) }

// https://docs.microsoft.com/en-us/windows/win32/menurc/wm-menuselect
func (me *_DepotMsg) WmMenuSelect(userFunc func(p WmMenuSelect)) {
	me.addMsg(co.WM_MENUSELECT, func(p Wm) uintptr {
		userFunc(WmMenuSelect(p))
		return 0
	})
}

type WmMenuSelect struct{ _Wm }

func (p WmMenuSelect) Item() uint16     { return p.WParam.LoWord() }
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
