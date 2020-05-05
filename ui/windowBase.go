package ui

import (
	"fmt"
	"syscall"
	"unsafe"
	"winffi/api"
	c "winffi/consts"
	"winffi/parm"
)

// Base to all window types.
type windowBase struct {
	hwnd api.HWND
	On   windowOn
}

// Constructor: must use.
func newWindowBase() windowBase {
	return windowBase{
		hwnd: api.HWND(0),
		On:   newWindowOn(),
	}
}

func (base *windowBase) Hwnd() api.HWND {
	return base.hwnd
}

func (base *windowBase) registerClass(wcx *api.WNDCLASSEX) api.ATOM {
	if wcx.LpszClassName == nil {
		panic("Class name not given.")
	}

	wcx.WndProc = syscall.NewCallback(wndProc)

	atom, errno := wcx.RegisterClassEx()
	if errno != 0 {
		if c.ERROR(errno) == c.ERROR_CLASS_ALREADY_EXISTS { // https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
			atom = api.ATOM(wcx.HInstance.GetClassInfo(wcx.LpszClassName, wcx))
		} else {
			panic(fmt.Sprintf("RegisterClassEx failed with atom %d: %d %s\n",
				atom, errno, errno.Error()))
		}
	}

	return atom
}

func wndProc(hwnd api.HWND, msg c.WM, wParam api.WPARAM, lParam api.LPARAM) uintptr {
	if msg == c.WM_NCCREATE {
		cs := (*api.CREATESTRUCT)(unsafe.Pointer(lParam))
		base := (*windowBase)(unsafe.Pointer(cs.CreateParams))
		hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))
		base.hwnd = hwnd // assign actual HWND
	}

	// Retrieve passed pointer.
	base := (*windowBase)(unsafe.Pointer(hwnd.GetWindowLongPtr(c.GWLP_USERDATA)))

	// Save *windowBase from being collected by GC.
	hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if base == nil {
		return hwnd.DefWindowProc(msg, wParam, lParam)
	}

	// Mount object to be passed to user handler.
	paramRaw := parm.Raw{
		Msg:    msg,
		WParam: wParam,
		LParam: lParam,
	}

	// Try to process the message with an user handler.
	userResult, wasProcessed := base.On.processMessage(paramRaw)

	// No further messages processed after this one.
	if msg == c.WM_NCDESTROY {
		base.hwnd.SetWindowLongPtr(c.GWLP_USERDATA, 0) // clear passed pointer
		base.hwnd = api.HWND(0)
	}

	if wasProcessed {
		return userResult
	}
	return hwnd.DefWindowProc(msg, wParam, lParam)
}
