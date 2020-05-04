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
	Wcx  api.WNDCLASSEX
	On   windowOn
}

// Constructor: must use.
func newWindowBase() windowBase {
	return windowBase{
		Wcx: api.WNDCLASSEX{
			Style: c.CS_DBLCLKS,
		},
		On: newWindowOn(),
	}
}

func (base *windowBase) Hwnd() api.HWND {
	return base.hwnd
}

func (base *windowBase) registerClass(hInst api.HINSTANCE) api.ATOM {
	base.Wcx.Size = uint32(unsafe.Sizeof(base.Wcx))
	base.Wcx.WndProc = syscall.NewCallback(wndProc)
	base.Wcx.HInstance = hInst

	atom, errno := base.Wcx.RegisterClassEx()
	if errno != 0 {
		if c.ERROR(errno) == c.ERROR_CLASS_ALREADY_EXISTS {
			atom = api.ATOM(hInst.GetClassInfo(base.Wcx.LpszClassName,
				&base.Wcx)) // https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
		} else {
			panic(fmt.Sprintf("RegisterClassEx failed with atom %d: %d %s",
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

	// Save *windowBase from being collected by GC.
	hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))

	if wasProcessed {
		return userResult
	}
	return hwnd.DefWindowProc(msg, wParam, lParam)
}
