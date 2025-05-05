//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

type _MainDlg struct {
	_BaseDlg
	iconId       uint16
	accelTableId uint16
}

// Constructor.
func newMainDlg(opts *VarOptsMainDlg) *_MainDlg {
	me := &_MainDlg{
		_BaseDlg:     newBaseDlg(opts.dlgId),
		iconId:       opts.iconId,
		accelTableId: opts.accelTableId,
	}
	me.defaultMessageHandlers()
	return me
}

func (me *_MainDlg) runAsMain(hInst win.HINSTANCE) int {
	me.createDialogParam(hInst, win.HWND(0))

	if me.iconId != 0 {
		me.setIcon(hInst, me.iconId)
	}

	var hAccel win.HACCEL
	if me.accelTableId != 0 {
		var err error
		hAccel, err = hInst.LoadAccelerators(win.ResIdInt(me.accelTableId))
		if err != nil {
			panic(err)
		}
	}

	me.hWnd.ShowWindow(co.SW_SHOW)
	return me.runMainLoop(hAccel, true)
}

func (me *_MainDlg) defaultMessageHandlers() {
	me._BaseDlg._BaseContainer.defaultMessageHandlers()

	me.userEvents.WmClose(func() {
		me.hWnd.DestroyWindow()
	})

	me.userEvents.WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})
}

// Options for ui.NewMainDlg(); returned by ui.OptsMainDlg().
type VarOptsMainDlg struct {
	dlgId        uint16
	iconId       uint16
	accelTableId uint16
}

// Options for ui.NewMainDlg().
func OptsMainDlg() *VarOptsMainDlg {
	return &VarOptsMainDlg{}
}

// Dialog resource ID passed to [CreateDialogParam].
//
// Panics if not informed.
//
// [CreateDialogParam]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createdialogparamw
func (o *VarOptsMainDlg) DlgId(id uint16) *VarOptsMainDlg { o.dlgId = id; return o }

// Dialog icon ID passed to [WM_SETICON].
//
// Defaults to none.
//
// [WM_SETICON]: https://learn.microsoft.com/en-us/windows/win32/winmsg/wm-seticon
func (o *VarOptsMainDlg) IconId(id uint16) *VarOptsMainDlg { o.iconId = id; return o }

// Accelerator table ID passed to [LoadAccelerators].
//
// Defaults to none.
//
// [LoadAccelerators]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadacceleratorsw
func (o *VarOptsMainDlg) AccelTableId(id uint16) *VarOptsMainDlg { o.accelTableId = id; return o }
