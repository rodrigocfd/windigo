//go:build windows

package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// Base to all windows created with CreateWindowEx.
type _RawBase struct {
	_BaseContainer
}

// Constructor.
func newBaseRaw() _RawBase {
	return _RawBase{
		_BaseContainer: newBaseContainer(_WNDTY_RAW),
	}
}

func (me *_RawBase) registerClass(
	hInst win.HINSTANCE,
	className string,
	classStyle co.CS,
	classIconId uint16,
	classBgBrush win.HBRUSH,
	classCursor win.HCURSOR,
) win.ATOM {
	wcx := win.WNDCLASSEX{
		LpfnWndProc:   wndProcCallback(),
		HInstance:     hInst,
		Style:         classStyle,
		HbrBackground: classBgBrush,
		HCursor:       classCursor,
	}
	wcx.SetCbSize()

	if classIconId != 0 {
		hIcon, err := hInst.LoadIcon(win.IconResId(classIconId))
		if err != nil {
			panic(err)
		}
		wcx.HIcon = hIcon
		wcx.HIconSm = hIcon
	}

	if className == "" {
		className = fmt.Sprintf("WNDCLASS %x.%x.%x.%x %x.%x.%x %x.%x.%x",
			wcx.Style, wcx.LpfnWndProc, wcx.CbClsExtra, wcx.CbWndExtra,
			wcx.HInstance, wcx.HIcon, wcx.HIconSm,
			wcx.HbrBackground, wcx.HIconSm, wcx.LpszMenuName)
	}

	wbuf := wstr.NewBufEncoder()
	defer wbuf.Free()
	wcx.LpszClassName = (*uint16)(wbuf.PtrEmptyIsNil(className))

	atom, err := win.RegisterClassEx(&wcx)
	if err != nil {
		if wErr, _ := err.(co.ERROR); wErr == co.ERROR_CLASS_ALREADY_EXISTS {
			// https://devblogs.microsoft.com/oldnewthing/20150429-00/?p=44984
			// https://devblogs.microsoft.com/oldnewthing/20041011-00/?p=37603
			// Retrieve atom from existing window class.
			atom, err = wcx.HInstance.GetClassInfoEx(wcx.LpszClassName, &wcx)
			if err != nil {
				panic(err) // GetClassInfoEx failed
			}
		} else if wErr != co.ERROR_SUCCESS {
			panic(wErr) // RegisterClassEx failed
		}
	}

	return atom
}

func (me *_RawBase) createWindow(
	exStyle co.WS_EX,
	className win.ATOM,
	title string,
	style co.WS,
	pos win.POINT,
	size win.SIZE,
	hParent win.HWND,
	hMenu win.HMENU,
	hInst win.HINSTANCE,
) {
	if me.hWnd != 0 {
		panic("Cannot create window twice.")
	}

	// The hWnd member is saved in WM_NCCREATE processing in wndProc.
	_, err := win.CreateWindowEx(exStyle, win.ClassNameAtom(className),
		title, style, int(pos.X), int(pos.Y), uint(size.Cx), uint(size.Cy),
		hParent, hMenu, hInst, win.LPARAM(unsafe.Pointer(me))) // pass pointer to object itself
	if err != nil {
		panic(err)
	}
}

func (me *_RawBase) delegateFocusToFirstChild() error {
	if hFocus := win.GetFocus(); hFocus == me.hWnd {
		// https://stackoverflow.com/a/2835220/6923555
		hFirstChild, err := me.hWnd.GetWindow(co.GW_CHILD)
		if err != nil {
			return err
		}
		hFirstChild.SetFocus()
	}
	return nil
}

var _wndProcCallback uintptr

func wndProcCallback() uintptr {
	if _wndProcCallback != 0 {
		return _wndProcCallback
	}

	_wndProcCallback = syscall.NewCallback(
		func(hWnd win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM) uintptr {
			var pMe *_RawBase

			if uMsg == co.WM_NCCREATE {
				cs := (*win.CREATESTRUCT)(unsafe.Pointer(lParam))
				pMe = (*_RawBase)(unsafe.Pointer(cs.LpCreateParams))
				pMe.hWnd = hWnd
				hWnd.SetWindowLongPtr(co.GWLP_USERDATA, uintptr(unsafe.Pointer(pMe))) // store
			} else {
				ptr, _ := hWnd.GetWindowLongPtr(co.GWLP_USERDATA) // retrieve
				pMe = (*_RawBase)(unsafe.Pointer(ptr))
			}

			// If no pointer stored, then no processing is done.
			// Prevents processing before WM_NCCREATE and after WM_NCDESTROY.
			if pMe == nil {
				return hWnd.DefWindowProc(uMsg, wParam, lParam)
			}

			// Execute before-user closures, keep track if at least one was executed.
			msg := Wm{uMsg, wParam, lParam}
			atLeastOneBeforeUser := pMe.beforeUserEvents.processAll(msg)

			// Execute user closure, if any.
			userRet, hasUserRet := pMe.userEvents.processLast(msg)

			// Execute post-user closures, keep track if at least one was executed.
			atLeastOneAfterUser := pMe.afterUserEvents.processAll(msg)

			switch uMsg {
			case co.WM_CREATE:
				pMe.removeWmCreateInitdialog() // will release all memory in these closures
			case co.WM_NCDESTROY: // always check
				hWnd.SetWindowLongPtr(co.GWLP_USERDATA, 0)
				pMe.hWnd = win.HWND(0)
				pMe.clearMessages()
			}

			if hasUserRet {
				return userRet
			} else if atLeastOneBeforeUser || atLeastOneAfterUser {
				return 0
			} else {
				return hWnd.DefWindowProc(uMsg, wParam, lParam)
			}
		},
	)
	return _wndProcCallback
}
