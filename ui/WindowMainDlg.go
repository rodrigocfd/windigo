package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// Implements WindowMain interface.
type _WindowMainDlg struct {
	_WindowBaseDlg
	iconId       int
	accelTableId int
}

// Creates a new WindowMain by loading a dialog resource.
//
// Parameters iconId and accelTableId are optional.
func NewWindowMainDlg(dialogId, iconId, accelTableId int) WindowMain {
	me := _WindowMainDlg{}
	me._WindowBaseDlg.new(dialogId)
	me.iconId = iconId
	me.accelTableId = accelTableId

	me.defaultMessages()
	return &me
}

func (me *_WindowMainDlg) RunAsMain() int {
	if win.IsWindowsVistaOrGreater() {
		win.SetProcessDPIAware()
	}
	win.InitCommonControls()
	_CreateGlobalUiFont()
	defer _globalUiFont.DeleteObject()

	hInst := win.GetModuleHandle("")
	me._WindowBaseDlg.createDialog(win.HWND(0), hInst)

	me.setIcon(me.iconId, hInst)
	me.Hwnd().ShowWindow(co.SW_SHOW)

	hAccel := win.HACCEL(0)
	if me.accelTableId != 0 {
		hAccel = hInst.LoadAccelerators(uintptr(me.accelTableId)) // automatically freed by system
	}

	return _RunMainLoop(me.Hwnd(), hAccel)
}

func (me *_WindowMainDlg) isDialog() bool {
	return true
}

func (me *_WindowMainDlg) defaultMessages() {
	me.On().WmClose(func() {
		me.Hwnd().DestroyWindow()
	})

	me.On().WmNcDestroy(func() {
		win.PostQuitMessage(int32(err.SUCCESS))
	})
}

func (me *_WindowMainDlg) setIcon(iconId int, hInst win.HINSTANCE) {
	if me.iconId != 0 {
		hIcon16, hIcon32 := me._WindowBaseDlg._WindowBase.loadIcons(hInst, iconId)
		me.Hwnd().SendMessage(co.WM_SETICON,
			win.WPARAM(co.ICON_SZ_SMALL), win.LPARAM(hIcon16))
		me.Hwnd().SendMessage(co.WM_SETICON,
			win.WPARAM(co.ICON_SZ_BIG), win.LPARAM(hIcon32))
	}
}
