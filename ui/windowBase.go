package ui

import (
	"log"
	"syscall"
	"unsafe"
	a "winffi/api"
	c "winffi/consts"
)

type windowBase struct {
	hwnd a.HWND
	Wcx  a.WNDCLASSEX
}

func (base *windowBase) Hwnd() a.HWND {
	return base.hwnd
}

func newWindowBase() windowBase {
	return windowBase{
		Wcx: a.WNDCLASSEX{
			Style: c.CS_DBLCLKS,
		},
	}
}

func (base *windowBase) registerClass(hInst a.HINSTANCE) a.ATOM {
	base.Wcx.Size = uint32(unsafe.Sizeof(base.Wcx))
	base.Wcx.WndProc = syscall.NewCallback(wndProc)
	base.Wcx.HInstance = hInst

	atom, errno := base.Wcx.RegisterClassEx()
	if errno != 0 {
		if c.ERROR(errno) == c.ERROR_CLASS_ALREADY_EXISTS {
			atom = a.ATOM(hInst.GetClassInfo(base.Wcx.LpszClassName,
				&base.Wcx)) // https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
		} else {
			log.Panicf("RegisterClassEx failed with atom %d: %d\n%s",
				atom, errno, errno.Error())
		}
	}

	return atom
}

func wndProc(hwnd a.HWND, msg c.WM, wParam a.WPARAM, lParam a.LPARAM) uintptr {
	if msg == c.WM_NCCREATE {
		cs := (*a.CREATESTRUCT)(unsafe.Pointer(lParam))
		base := (*windowBase)(unsafe.Pointer(cs.CreateParams))
		hwnd.SetWindowLongPtr(c.GWLP_USERDATA, uintptr(unsafe.Pointer(base)))
		base.hwnd = hwnd // assign actual HWND
	}

	base := (*windowBase)(unsafe.Pointer(hwnd.GetWindowLongPtr(c.GWLP_USERDATA)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if base == nil {
		return hwnd.DefWindowProc(msg, wParam, lParam)
	}

	//...

	return hwnd.DefWindowProc(msg, wParam, lParam)
}
