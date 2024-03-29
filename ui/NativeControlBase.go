package ui

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
)

var (
	_globalBaseSubclassId  = uint32(0)  // incremented at each subclass installed
	_globalSubclassProcPtr = uintptr(0) // necessary for RemoveWindowSubclass()
)

//------------------------------------------------------------------------------

type _NativeControlBase struct {
	hWnd        win.HWND
	ctrlId      int
	parent      AnyParent
	eventsSubcl _EventsWm // subclass events
	subclassId  uint32
}

func (me *_NativeControlBase) new(parent AnyParent, ctrlId int) {
	me.hWnd = win.HWND(0)
	me.ctrlId = ctrlId
	me.parent = parent
	me.eventsSubcl.new()
	me.subclassId = 0
}

// Implements AnyWindow.
func (me *_NativeControlBase) Hwnd() win.HWND {
	return me.hWnd
}

// Implements AnyControl.
func (me *_NativeControlBase) CtrlId() int {
	return me.ctrlId
}

// Implements AnyControl.
func (me *_NativeControlBase) Parent() AnyParent {
	return me.parent
}

// Implements AnyNativeControl.
func (me *_NativeControlBase) OnSubclass() *_EventsWm {
	if me.Hwnd() != 0 {
		panic("Cannot add event handling after the control is created.")
	}
	return &me.eventsSubcl
}

// Calls CreateWindowEx() and installs subclass.
func (me *_NativeControlBase) createWindow(
	exStyle co.WS_EX, className, title string, style co.WS,
	pos win.POINT, size win.SIZE, hMenu win.HMENU) {

	if me.hWnd != 0 {
		panic(fmt.Sprintf("Control already created: \"%s\".", className))
	}

	me.hWnd = win.CreateWindowEx(exStyle, className, title, style,
		pos.X, pos.Y, size.Cx, size.Cy, me.parent.Hwnd(), hMenu,
		me.parent.Hwnd().Hinstance(), 0)

	me.installSubclassIfNeeded()
}

// Calls GetDlgItem() and installs subclass.
func (me *_NativeControlBase) assignDlgItem() {
	if me.hWnd != 0 {
		panic(fmt.Sprintf("Dialog control already assigned: \"%d\".", me.ctrlId))
	}
	if !me.parent.isDialog() {
		panic(fmt.Sprintf("Parent is not dialog, cannot assign control: \"%d\".", me.ctrlId))
	}

	me.hWnd = me.parent.Hwnd().GetDlgItem(int32(me.ctrlId))
	me.installSubclassIfNeeded()
}

func (me *_NativeControlBase) installSubclassIfNeeded() {
	if me.eventsSubcl.hasMessages() {
		if _globalSubclassProcPtr == 0 {
			_globalSubclassProcPtr = syscall.NewCallback(_SubclassProc)
		}
		_globalBaseSubclassId++
		me.subclassId = _globalBaseSubclassId

		// Subclass is installed after window creation, thus WM_CREATE can never
		// be handled for a subclassed control.
		me.hWnd.SetWindowSubclass(_globalSubclassProcPtr,
			me.subclassId, unsafe.Pointer(me)) // pass pointer to object itself
	}
}

// Keeps all *_NativeControlBase that were retrieved in _SubclassProc.
var _globalNativeControlBasePtrs = make(map[win.HWND]*_NativeControlBase, 10)

// Default window procedure for subclassed child controls.
func _SubclassProc(
	hWnd win.HWND, uMsg co.WM, wParam win.WPARAM, lParam win.LPARAM,
	uIdSubclass, dwRefData uintptr) uintptr {

	// Retrieve passed pointer.
	pMe := (*_NativeControlBase)(unsafe.Pointer(dwRefData))
	_globalNativeControlBasePtrs[hWnd] = pMe

	if pMe, hasPtr := _globalNativeControlBasePtrs[hWnd]; hasPtr {
		// Try to process the message with an user handler.
		retVal, meaningfulRet, wasHandled :=
			pMe.eventsSubcl.processMessage(uMsg, wParam, lParam)

		if uMsg == co.WM_NCDESTROY { // even if the user handles WM_NCDESTROY, we must ensure cleanup
			hWnd.RemoveWindowSubclass(_globalSubclassProcPtr, pMe.subclassId)
			delete(_globalNativeControlBasePtrs, hWnd) // clear our pointer
		}

		if wasHandled {
			if meaningfulRet {
				return retVal
			}
			return 0 // message processed, default return value
		}
	} else if uMsg == co.WM_NCDESTROY {
		// https://devblogs.microsoft.com/oldnewthing/20031111-00/?p=41883
		hWnd.RemoveWindowSubclass(_globalSubclassProcPtr, pMe.subclassId)
	}

	return hWnd.DefSubclassProc(uMsg, wParam, lParam) // message was not processed
}
