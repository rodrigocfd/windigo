package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/err"
)

// Base to all windows created by specifying all options, which will be passed
// to the underlying CreateWindowEx().
type _WindowBaseOpts struct {
	_WindowBase
}

func (me *_WindowBaseOpts) new() {
	me._WindowBase.new()
}

// Fills the WNDCLASSEX structure with the given parameters, and the class name,
// if not specified, will be auto-generated.
func (me *_WindowBaseOpts) generateWcx(
	wcx *win.WNDCLASSEX,
	hInst win.HINSTANCE, className string, classStyles co.CS,
	hCursor win.HCURSOR, hBrushBg win.HBRUSH, iconId int) string {

	wcx.CbSize = uint32(unsafe.Sizeof(wcx))
	wcx.LpfnWndProc = syscall.NewCallback(_WndProc)
	wcx.HInstance = hInst
	wcx.Style = classStyles
	wcx.HCursor = hCursor
	wcx.HbrBackground = hBrushBg

	if iconId != 0 {
		wcx.HIconSm, wcx.HIcon = me._WindowBase.loadIcons(hInst, iconId)
	}

	// After all the fields are set, if no class name, we generate one by hashing
	// all WNDCLASSEX fields. That's why it must be the last thing to be done.
	actualClassName := className
	if actualClassName == "" {
		actualClassName = fmt.Sprintf("%x.%x.%x.%x.%x.%x.%x.%x.%x.%x",
			wcx.Style, wcx.LpfnWndProc, wcx.CbClsExtra, wcx.CbWndExtra,
			wcx.HInstance, wcx.HIcon, wcx.HCursor, wcx.HbrBackground,
			wcx.LpszMenuName, wcx.HIconSm)
	}
	actualClassNameSlice := win.Str.ToUint16Slice(actualClassName)
	wcx.LpszClassName = &actualClassNameSlice[0]

	return actualClassName
}

// Calls RegisterClassEx().
func (me *_WindowBaseOpts) registerClass(wcx *win.WNDCLASSEX) win.ATOM {
	atom, fail := win.RegisterClassEx(wcx)
	if fail != nil {
		lerr, _ := fail.(err.ERROR)
		if lerr == err.CLASS_ALREADY_EXISTS {
			// https://devblogs.microsoft.com/oldnewthing/20150429-00/?p=44984
			// https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
			atom, fail = wcx.HInstance.GetClassInfoEx( // retrieve atom from existing window class
				(*uint16)(unsafe.Pointer(wcx.LpszClassName)), wcx)
			if fail != nil {
				panic(fail) // GetClassInfoEx failed
			}
		} else if lerr != err.SUCCESS {
			panic(lerr) // RegisterClassEx failed
		}
	}
	return atom
}

// Calls CreateWindowEx().
func (me *_WindowBaseOpts) createWindow(
	exStyle co.WS_EX, className, title string, style co.WS,
	pos win.POINT, size win.SIZE, hParent win.HWND, hMenu win.HMENU,
	hInst win.HINSTANCE) {

	if me._WindowBase.hWnd != 0 {
		panic(fmt.Sprintf("Window already created: \"%s\".", className))
	}

	// The hwnd member is saved in WM_NCCREATE processing in wndProc.
	win.CreateWindowEx(exStyle, className, title, style,
		pos.X, pos.Y, size.Cx, size.Cy, hParent, hMenu, hInst,
		win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Returns window coords at screen center, and window size from its client area.
func (me *_WindowBaseOpts) calcWndCoords(
	pClientArea *win.SIZE, hMenu win.HMENU,
	styles co.WS, exStyles co.WS_EX) (win.POINT, win.SIZE) {

	screenSize := win.SIZE{
		Cx: win.GetSystemMetrics(co.SM_CXSCREEN),
		Cy: win.GetSystemMetrics(co.SM_CYSCREEN),
	}

	_MultiplyDpi(nil, pClientArea) // in-place correct for DPI

	pos := win.POINT{
		X: screenSize.Cx/2 - pClientArea.Cx/2, // center on screen
		Y: screenSize.Cy/2 - pClientArea.Cy/2,
	}

	rc := win.RECT{
		Left:   pos.X,
		Top:    pos.Y,
		Right:  pClientArea.Cx + pos.X,
		Bottom: pClientArea.Cy + pos.Y,
	}
	win.AdjustWindowRectEx(&rc, styles,
		hMenu != 0 && hMenu.GetMenuItemCount() > 0, exStyles)

	return win.POINT{X: rc.Left, Y: rc.Top},
		win.SIZE{Cx: rc.Right - rc.Left, Cy: rc.Bottom - rc.Top}
}

// Default window procedure.
func _WndProc(
	hWnd win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if uMsg == co.WM_NCCREATE {
		cs := (*win.CREATESTRUCT)(unsafe.Pointer(lParam))
		pMe := (*_WindowBaseOpts)(unsafe.Pointer(cs.LpCreateParams))
		hWnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))
		pMe._WindowBase.hWnd = hWnd // assign actual HWND
	}

	// Retrieve passed pointer.
	pMe := (*_WindowBaseOpts)(unsafe.Pointer(hWnd.GetWindowLongPtr(co.GWLP_USERDATA)))

	// If the retrieved *_WindowBaseRaw stays here, the GC will collect it.
	// Sending it away will prevent the GC collection.
	// https://stackoverflow.com/a/51188315
	hWnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))

	// If no pointer stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if pMe != nil {
		// Process all internal events.
		pMe.internalEvents.processMessages(uMsg, wParam, lParam)

		// Try to process the message with an user handler.
		retVal, meaningfulRet, wasHandled :=
			pMe._WindowBase.events.processMessage(uMsg, wParam, lParam)

		// No further messages processed after this one.
		if uMsg == co.WM_NCDESTROY {
			pMe._WindowBase.hWnd.SetWindowLongPtr(co.GWLP_USERDATA, 0) // clear passed pointer
			pMe._WindowBase.hWnd = win.HWND(0)
		}

		if wasHandled {
			if meaningfulRet {
				return retVal
			}
			return 0 // message processed, default return value
		}
	}

	return hWnd.DefWindowProc(uMsg, wParam, lParam) // message not processed
}
