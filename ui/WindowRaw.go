//go:build windows

package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// Base to all windows created by specifying all options, which will be passed
// to the underlying CreateWindowEx().
type _WindowRaw struct {
	_WindowBase
}

func (me *_WindowRaw) new() {
	me._WindowBase.new()
}

// Fills the WNDCLASSEX structure with the given parameters, and the class name,
// if not specified, will be auto-generated.
func (me *_WindowRaw) generateWcx(
	wcx *win.WNDCLASSEX,
	hInst win.HINSTANCE, className string, classStyles co.CS,
	hCursor win.HCURSOR, hBrushBg win.HBRUSH, iconId int) string {

	wcx.SetCbSize()
	wcx.LpfnWndProc = _globalWndProc
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
	actualClassNameSlice := win.Str.ToNativeSlice(actualClassName)
	wcx.LpszClassName = &actualClassNameSlice[0]

	return actualClassName
}

// Calls RegisterClassEx().
func (me *_WindowRaw) registerClass(wcx *win.WNDCLASSEX) win.ATOM {
	atom, err := win.RegisterClassEx(wcx)
	if err != nil {
		if wErr, _ := err.(errco.ERROR); wErr == errco.CLASS_ALREADY_EXISTS {
			// https://devblogs.microsoft.com/oldnewthing/20150429-00/?p=44984
			// https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
			// Retrieve atom from existing window class.
			atom, err = wcx.HInstance.GetClassInfoEx(wcx.LpszClassName, wcx)
			if err != nil {
				panic(err) // GetClassInfoEx failed
			}
		} else if wErr != errco.SUCCESS {
			panic(wErr) // RegisterClassEx failed
		}
	}
	return atom
}

// Calls CreateWindowEx().
func (me *_WindowRaw) createWindow(
	exStyle co.WS_EX, className win.ClassName, title win.StrOpt, style co.WS,
	pos win.POINT, size win.SIZE, hParent win.HWND, hMenu win.HMENU,
	hInst win.HINSTANCE) {

	if me._WindowBase.hWnd != 0 {
		panic("Window already created.")
	}

	_globalWindowRawPtrs[me] = struct{}{} // store pointer in the set

	// The hwnd member is saved in WM_NCCREATE processing in wndProc.
	win.CreateWindowEx(exStyle, className, title, style,
		pos.X, pos.Y, size.Cx, size.Cy, hParent, hMenu, hInst,
		win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
}

// Returns window coords at screen center, and window size from its client area.
func (me *_WindowRaw) calcWndCoords(
	pClientArea *win.SIZE, hMenu win.HMENU,
	styles co.WS, exStyles co.WS_EX) (win.POINT, win.SIZE) {

	_MultiplyDpi(nil, pClientArea) // in-place correct for DPI
	rc := win.RECT{
		Right:  pClientArea.Cx,
		Bottom: pClientArea.Cy,
	}
	win.AdjustWindowRectEx(&rc, styles,
		hMenu != 0 && hMenu.GetMenuItemCount() > 0, exStyles)

	pClientArea.Cx = rc.Right - rc.Left // as if top corner is 0,0
	pClientArea.Cy = rc.Bottom - rc.Top

	screenSize := win.SIZE{
		Cx: win.GetSystemMetrics(co.SM_CXSCREEN),
		Cy: win.GetSystemMetrics(co.SM_CYSCREEN),
	}
	pos := win.POINT{
		X: screenSize.Cx/2 - pClientArea.Cx/2, // center on screen
		Y: screenSize.Cy/2 - pClientArea.Cy/2,
	}

	return pos, *pClientArea
}

var (
	// A set keeping all *_WindowRaw that were retrieved in _WndProc.
	_globalWindowRawPtrs = make(map[*_WindowRaw]struct{}, 10)

	// Default window procedure.
	_globalWndProc uintptr = syscall.NewCallback(_WndProc)
)

func _WndProc(
	hWnd win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {

	var pMe *_WindowRaw

	// https://devblogs.microsoft.com/oldnewthing/20050422-08/?p=35813
	if uMsg == co.WM_NCCREATE {
		cs := (*win.CREATESTRUCT)(unsafe.Pointer(lParam))
		pMe = (*_WindowRaw)(unsafe.Pointer(cs.LpCreateParams))
		pMe._WindowBase.hWnd = hWnd // assign actual HWND
		hWnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe)))
	} else {
		pMe = (*_WindowRaw)(unsafe.Pointer(hWnd.GetWindowLongPtr(co.GWLP_USERDATA)))
	}

	// If object pointer is not stored, then no processing is done.
	// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
	if _, isStored := _globalWindowRawPtrs[pMe]; isStored {
		// Process all internal events.
		pMe.internalEvents.processMessages(uMsg, wParam, lParam)

		// Try to process the message with an user handler.
		retVal, meaningfulRet, wasHandled :=
			pMe._WindowBase.events.processMessage(uMsg, wParam, lParam)

		// No further messages processed after this one.
		if uMsg == co.WM_NCDESTROY {
			delete(_globalWindowRawPtrs, pMe) // remove from set
			hWnd.SetWindowLongPtr(co.GWLP_USERDATA, 0)
			pMe._WindowBase.hWnd = win.HWND(0)
			pMe._WindowBase.clearMessages() // prevents circular references
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
