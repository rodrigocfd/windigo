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

type wplp struct {
	WParam win.WPARAM
	LParam win.LPARAM
}

// Raw window message parameters.
type Wm struct {
	wplp
}

type WmActivate struct{ wplp }

func (p WmActivate) Event() co.WA                           { return co.WA(p.WParam.LoWord()) }
func (p WmActivate) IsMinimized() bool                      { return p.WParam.HiWord() != 0 }
func (p WmActivate) ActivatedOrDeactivatedWindow() win.HWND { return win.HWND(p.LParam) }

type WmActivateApp struct{ wplp }

func (p WmActivateApp) IsBeingActivated() bool { return p.WParam != 0 }
func (p WmActivateApp) ThreadId() uint32       { return uint32(p.LParam) }

type WmAppCommand struct{ wplp }

func (p WmAppCommand) OwnerWindow() win.HWND     { return win.HWND(p.WParam) }
func (p WmAppCommand) AppCommand() co.APPCOMMAND { return co.APPCOMMAND(p.LParam.HiWord() &^ 0xF000) }
func (p WmAppCommand) UDevice() co.FAPPCOMMAND   { return co.FAPPCOMMAND(p.LParam.HiWord() & 0xF000) }
func (p WmAppCommand) Keys() co.MK               { return co.MK(p.LParam.LoWord()) }

type bChar struct{ wplp } // base for other messages

func (p bChar) CharCode() uint16          { return uint16(p.WParam) }
func (p bChar) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p bChar) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p bChar) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p bChar) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p bChar) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }
func (p bChar) KeyBeingReleased() bool    { return (p.LParam.HiByteHiWord() & 0b1000_0000) != 0 }

type WmChar struct{ bChar }

type WmCommand struct{ wplp }

func (p WmCommand) IsFromMenu() bool         { return p.WParam.HiWord() == 0 }
func (p WmCommand) IsFromAccelerator() bool  { return p.WParam.HiWord() == 1 }
func (p WmCommand) IsFromControl() bool      { return !p.IsFromMenu() && !p.IsFromAccelerator() }
func (p WmCommand) MenuId() int32            { return p.ControlId() }
func (p WmCommand) AcceleratorId() int32     { return p.ControlId() }
func (p WmCommand) ControlId() int32         { return int32(p.WParam.LoWord()) }
func (p WmCommand) ControlNotifCode() uint16 { return p.WParam.HiWord() }
func (p WmCommand) ControlHwnd() win.HWND    { return win.HWND(p.LParam) }

type WmContextMenu struct{ wplp }

func (p WmContextMenu) RightClickedWindow() win.HWND { return win.HWND(p.WParam) }
func (p WmContextMenu) CursorPos() win.POINT         { return p.LParam.MakePoint() }

type WmCreate struct{ wplp }

func (p WmCreate) CreateStruct() *win.CREATESTRUCT {
	return (*win.CREATESTRUCT)(unsafe.Pointer(p.LParam))
}

type WmDeadChar struct{ bChar }

type WmDropFiles struct{ wplp }

func (p WmDropFiles) Hdrop() win.HDROP { return win.HDROP(p.WParam) }

// Calls DragQueryFile successively to retrieve all file names, and releases the
// HDROP handle.
func (p WmDropFiles) RetrieveAll() []string {
	count := p.Hdrop().DragQueryFile(0xFFFFFFFF, nil, 0)
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

type WmGetDlgCode struct{ wplp }

func (p WmGetDlgCode) VirtualKeyCode() co.VK { return co.VK(p.WParam) }
func (p WmGetDlgCode) IsQuery() bool         { return p.LParam == 0 }
func (p WmGetDlgCode) Msg() *win.MSG         { return (*win.MSG)(unsafe.Pointer(p.LParam)) }
func (p WmGetDlgCode) HasAlt() bool          { return (win.GetAsyncKeyState(co.VK_MENU) & 0x8000) != 0 }
func (p WmGetDlgCode) HasCtrl() bool         { return (win.GetAsyncKeyState(co.VK_CONTROL) & 0x8000) != 0 }
func (p WmGetDlgCode) HasShift() bool        { return (win.GetAsyncKeyState(co.VK_SHIFT) & 0x8000) != 0 }

type WmHelp struct{ wplp }

func (p WmHelp) HelpInfo() *win.HELPINFO { return (*win.HELPINFO)(unsafe.Pointer(p.LParam)) }

type WmHotKey struct{ wplp }

func (p WmHotKey) HotKey() co.IDHOT      { return co.IDHOT(p.WParam) }
func (p WmHotKey) OtherKeys() co.MOD     { return co.MOD(p.LParam.LoWord()) }
func (p WmHotKey) VirtualKeyCode() co.VK { return co.VK(p.LParam.HiWord()) }

type WmInitMenuPopup struct{ wplp }

func (p WmInitMenuPopup) Hmenu() win.HMENU        { return win.HMENU(p.WParam) }
func (p WmInitMenuPopup) MenuRelativePos() uint16 { return p.LParam.LoWord() }
func (p WmInitMenuPopup) IsWindowMenu() bool      { return p.LParam.HiWord() != 0 }

type bKeyUpDn struct{ wplp } // base for other messages

func (p bKeyUpDn) VirtualKeyCode() co.VK     { return co.VK(p.WParam) }
func (p bKeyUpDn) RepeatCount() uint16       { return p.LParam.LoWord() }
func (p bKeyUpDn) ScanCode() uint8           { return p.LParam.LoByteHiWord() }
func (p bKeyUpDn) IsExtendedKey() bool       { return (p.LParam.HiByteHiWord() & 0b0000_0001) != 0 }
func (p bKeyUpDn) HasAltKey() bool           { return (p.LParam.HiByteHiWord() & 0b0010_0000) != 0 }
func (p bKeyUpDn) IsKeyDownBeforeSend() bool { return (p.LParam.HiByteHiWord() & 0b0100_0000) != 0 }

type WmKeyDown struct{ bKeyUpDn }

type WmKeyUp struct{ bKeyUpDn }

type WmKillFocus struct{ wplp }

func (p WmKillFocus) WindowReceivingFocus() win.HWND { return win.HWND(p.WParam) }

type bButtonClk struct{ wplp } // base for other messages

func (p bButtonClk) HasCtrl() bool      { return (co.MK(p.WParam) & co.MK_CONTROL) != 0 }
func (p bButtonClk) HasLeftBtn() bool   { return (co.MK(p.WParam) & co.MK_LBUTTON) != 0 }
func (p bButtonClk) HasMiddleBtn() bool { return (co.MK(p.WParam) & co.MK_MBUTTON) != 0 }
func (p bButtonClk) HasRightBtn() bool  { return (co.MK(p.WParam) & co.MK_RBUTTON) != 0 }
func (p bButtonClk) HasShift() bool     { return (co.MK(p.WParam) & co.MK_SHIFT) != 0 }
func (p bButtonClk) HasXBtn1() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON1) != 0 }
func (p bButtonClk) HasXBtn2() bool     { return (co.MK(p.WParam) & co.MK_XBUTTON2) != 0 }
func (p bButtonClk) Pos() win.POINT     { return p.LParam.MakePoint() }

type WmLButtonDblClk struct{ bButtonClk }

type WmLButtonDown struct{ bButtonClk }

type WmLButtonUp struct{ bButtonClk }

type WmMButtonDblClk struct{ bButtonClk }

type WmMButtonDown struct{ bButtonClk }

type WmMButtonUp struct{ bButtonClk }

type WmMenuChar struct{ wplp }

func (p WmMenuChar) CharCode() uint16      { return p.WParam.LoWord() }
func (p WmMenuChar) ActiveMenuType() co.MF { return co.MF(p.WParam.HiWord()) }
func (p WmMenuChar) ActiveMenu() win.HMENU { return win.HMENU(p.LParam) }

type WmMenuCommand struct{ wplp }

func (p WmMenuCommand) ItemIndex() uint16 { return uint16(p.WParam) }
func (p WmMenuCommand) Hmenu() win.HMENU  { return win.HMENU(p.LParam) }

type WmMenuSelect struct{ wplp }

func (p WmMenuSelect) Item() uint16     { return p.WParam.LoWord() }
func (p WmMenuSelect) Flags() co.MF     { return co.MF(p.WParam.HiWord()) }
func (p WmMenuSelect) Hmenu() win.HMENU { return win.HMENU(p.LParam) }

type WmMouseHover struct{ bButtonClk }

type WmMouseMove struct{ bButtonClk }

type WmMove struct{ wplp }

func (p WmMove) Pos() win.POINT { return p.LParam.MakePoint() }

type WmNcPaint struct{ wplp }

func (p WmNcPaint) Hrgn() win.HRGN { return win.HRGN(p.WParam) }

type WmNotify struct{ wplp }

func (p WmNotify) NmHdr() *win.NMHDR { return (*win.NMHDR)(unsafe.Pointer(p.LParam)) }

type WmPrint struct{ wplp }

func (p WmPrint) Hdc() win.HDC           { return win.HDC(p.WParam) }
func (p WmPrint) DrawingOptions() co.PRF { return co.PRF(p.LParam) }

type WmRButtonDblClk struct{ bButtonClk }

type WmRButtonDown struct{ bButtonClk }

type WmRButtonUp struct{ bButtonClk }

type WmSetFocus struct{ wplp }

func (p WmSetFocus) UnfocusedWindow() win.HWND { return win.HWND(p.WParam) }

type WmSetFont struct{ wplp }

func (p WmSetFont) Hfont() win.HFONT   { return win.HFONT(p.WParam) }
func (p WmSetFont) ShouldRedraw() bool { return p.LParam == 1 }

type WmSize struct{ wplp }

func (p WmSize) Request() co.SIZE         { return co.SIZE(p.WParam) }
func (p WmSize) ClientAreaSize() win.SIZE { return p.LParam.MakeSize() }

type WmSysChar struct{ bChar }

type WmSysCommand struct{ wplp }

func (p WmSysCommand) RequestCommand() co.SC { return co.SC(p.WParam) }
func (p WmSysCommand) CursorPos() win.POINT  { return p.LParam.MakePoint() }

type WmSysDeadChar struct{ bChar }

type WmSysKeyDown struct{ bKeyUpDn }

type WmSysKeyUp struct{ bKeyUpDn }
