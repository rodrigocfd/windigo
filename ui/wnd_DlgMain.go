//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

// Dialog-based main application window.
type _DlgMain struct {
	_DlgBase
	iconId       uint16
	accelTableId uint16
	dropFiles    bool
}

// Constructor.
func newMainDlg(opts *VarOptsMainDlg) *_DlgMain {
	me := &_DlgMain{
		_DlgBase:     newBaseDlg(opts.dlgId),
		iconId:       opts.iconId,
		accelTableId: opts.accelTableId,
		dropFiles:    opts.dropFiles,
	}
	me.defaultMessageHandlers()
	return me
}

func (me *_DlgMain) runAsMain(hInst win.HINSTANCE) int {
	if me.dropFiles {
		oleRel := me._DlgBase._BaseContainer.initDropTarget()
		defer oleRel.Release()
	}

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

func (me *_DlgMain) defaultMessageHandlers() {
	me._DlgBase._BaseContainer.defaultMessageHandlers()

	me.userEvents.WmClose(func() {
		me.hWnd.DestroyWindow()
	})

	me.userEvents.WmNcDestroy(func() {
		win.PostQuitMessage(0)
	})
}

// Options for [NewMainDlg]; returned by [OptsMainDlg].
type VarOptsMainDlg struct {
	dlgId        uint16
	iconId       uint16
	accelTableId uint16
	dropFiles    bool
}

// Options for [NewMainDlg].
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

// Declares an internal [IDropTarget] object, so files can be [dragged] onto the
// window.
//
// When a drag and drop happens, the window will automatically receive a
// [WM_DROPFILES] message:
//
//	wnd.On().WmDropFiles(func(p ui.WmDropFiles) {
//		filePaths, _ := p.HDrop().DragQueryFile()
//		// ...
//	})
//
// In order to make it work, don't forget to call [win.OleInitialize] and defer
// [win.OleUninitialize] at the beginning of your program.
//
// [IDropTarget]: https://learn.microsoft.com/en-us/windows/win32/api/oleidl/nn-oleidl-idroptarget
// [dragged]: https://learn.microsoft.com/en-us/windows/win32/com/drag-and-drop
// [WM_DROPFILES]: https://learn.microsoft.com/en-us/windows/win32/shell/wm-dropfiles
func (o *VarOptsMainDlg) DropFiles(d bool) *VarOptsMainDlg { o.dropFiles = d; return o }
