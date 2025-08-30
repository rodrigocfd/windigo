//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// The buttons collection
//
// You cannot create this object directly, it will be created automatically
// by the owning [Toolbar].
type CollectionToolbarButtons struct {
	owner *Toolbar
}

// Adds a new button with [TB_ADDBUTTONS].
//
// The iconIndex is the zero-based index of the icon previously inserted into
// the control's image list.
//
// [TB_ADDBUTTONS]: https://learn.microsoft.com/en-us/windows/win32/controls/tb-addbuttons
func (me *CollectionToolbarButtons) Add(cmdId uint16, text string, iconIndex int) {
	var wText wstr.BufEncoder

	tbb := win.TBBUTTON{
		IBitmap:   int32(iconIndex),
		IdCommand: int32(cmdId),
		FsStyle:   co.BTNS_AUTOSIZE,
		FsState:   co.TBSTATE_ENABLED,
		IString:   (*uint16)(wText.AllowEmpty(text)),
	}

	ret, _ := me.owner.hWnd.SendMessage(co.TB_ADDBUTTONS,
		1, win.LPARAM(unsafe.Pointer(&tbb)))
	if ret == 0 {
		panic(fmt.Sprintf("TB_ADDBUTTONS \"%s\" failed.", text))
	}

	me.owner.hWnd.SendMessage(co.TB_AUTOSIZE, 0, 0)
}

// Retrieves the number of buttons with [TB_BUTTONCOUNT].
//
// [TB_BUTTONCOUNT]: https://learn.microsoft.com/en-us/windows/win32/controls/tb-buttoncount
func (me *CollectionToolbarButtons) Count() uint {
	ret, _ := me.owner.hWnd.SendMessage(co.TB_BUTTONCOUNT, 0, 0)
	return uint(ret)
}
