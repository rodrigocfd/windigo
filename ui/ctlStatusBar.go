/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package ui

import (
	"fmt"
	"unsafe"
	"windigo/co"
	"windigo/win"
)

// Native status bar control.
//
// https://docs.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBar struct {
	*_NativeControlBase
	events *_EventsStatusBar
	parts  *_StatusBarPartCollection
}

// Constructor. Optionally receives a control ID.
func NewStatusBar(parent Parent, ctrlId ...int) *StatusBar {
	base := _NewNativeControlBase(parent, ctrlId...)
	me := &StatusBar{
		_NativeControlBase: base,
		events:             _NewEventsStatusBar(base),
	}
	me.parts = _NewStatusBarPartCollection(me)
	return me
}

// Calls CreateWindowEx().
//
// Control will be docked at bottom of parent window.
//
// Should be called at On().WmCreate(), or at On().WmInitDialog() if dialog.
func (me *StatusBar) Create() *StatusBar {
	sbStyle := co.WS_CHILD | co.WS_VISIBLE

	parentStyle := me.parent.Hwnd().GetStyle()
	isParentResizable := (parentStyle&co.WS_MAXIMIZEBOX) != 0 ||
		(parentStyle&co.WS_SIZEBOX) != 0

	if isParentResizable {
		sbStyle |= co.WS(co.SBARS_SIZEGRIP)
	}

	me._NativeControlBase.create("msctls_statusbar32", "", Pos{}, Size{},
		sbStyle, co.WS_EX_NONE)
	return me
}

// Exposes all StatusBar notifications.
//
// Cannot be called after the parent window was created.
func (me *StatusBar) On() *_EventsStatusBar {
	if me.hwnd != 0 {
		panic("Cannot add notifications after the StatusBar was created.")
	}
	return me.events
}

// Access to the parts.
func (me *StatusBar) Parts() *_StatusBarPartCollection {
	return me.parts
}

// Resizes the StatusBar to fill the available width on parent window.
// Intended to be called with parent's WM_SIZE processing.
func (me *StatusBar) ResizeToFitParent(p WmSize) *StatusBar {
	if p.Request() != co.SIZE_MINIMIZED && me.Hwnd() != 0 {
		cx := int(p.ClientAreaSize().Cx)        // available width
		me.Hwnd().SendMessage(co.WM_SIZE, 0, 0) // tell status bar to fit parent

		// Find the space to be divided among variable-width parts,
		// and total weight of variable-width parts.
		totalWeight := 0
		cxVariable := cx
		for i := range me.parts.partsData {
			if me.parts.partsData[i].IsFixedWidth() {
				cxVariable -= me.parts.partsData[i].sizePixels
			} else {
				totalWeight += me.parts.partsData[i].resizeWeight
			}
		}

		// Fill right edges array with the right edge of each part.
		cxTotal := cx
		for i := len(me.parts.partsData) - 1; i >= 0; i-- {
			me.parts.rightEdges[i] = int32(cxTotal)
			if me.parts.partsData[i].IsFixedWidth() {
				cxTotal -= me.parts.partsData[i].sizePixels
			} else {
				cxTotal -= (cxVariable / totalWeight) * me.parts.partsData[i].resizeWeight
			}
		}
		me.Hwnd().SendMessage(co.WM(co.SB_SETPARTS),
			win.WPARAM(len(me.parts.rightEdges)),
			win.LPARAM(unsafe.Pointer(&me.parts.rightEdges[0])))
	}

	return me
}

//------------------------------------------------------------------------------

type _StatusBarPartCollection struct {
	ctrl            *StatusBar
	partsData       []_StatusBarPartData
	rightEdges      []int32 // buffer to speed up ResizeToFitParent() calls
	initialParentCx int     // cache used when adding parts
}

// Constructor.
func _NewStatusBarPartCollection(ctrl *StatusBar) *_StatusBarPartCollection {
	return &_StatusBarPartCollection{
		ctrl: ctrl,
	}
}

// Adds a new part with fixed width.
//
// Width will be adjusted to the current system DPI.
func (me *_StatusBarPartCollection) AddFixed(sizePixels int) *StatusBar {
	if sizePixels < 0 {
		panic("Width of StatusBar part can't be negative.")
	}

	me.cacheInitialParentCx()

	size := Size{Cx: sizePixels, Cy: 0}
	_global.MultiplyDpi(nil, &size)

	me.partsData = append(me.partsData, _StatusBarPartData{
		sizePixels: size.Cx,
	})
	me.rightEdges = append(me.rightEdges, 0)

	me.ctrl.ResizeToFitParent(WmSize{
		m: Wm{
			WParam: win.WPARAM(co.SIZE_RESTORED),
			LParam: win.MakeLParam(uint16(me.initialParentCx), 0),
		},
	})
	return me.ctrl
}

// Adds a new resizable part.
//
// How resizeWeight works:
//
// - Suppose you have 3 parts, respectively with weights of 1, 1 and 2.
//
// - If available client area is 400px, respective part widths will be 100, 100 and 200px.
func (me *_StatusBarPartCollection) AddResizable(resizeWeight int) *StatusBar {
	if resizeWeight <= 0 {
		panic("Resize weight must be equal or greater than 1.")
	}

	me.cacheInitialParentCx()

	// Zero weight means a fixed-width part, which internally should have sizePixels set.
	me.partsData = append(me.partsData, _StatusBarPartData{
		resizeWeight: resizeWeight,
	})
	me.rightEdges = append(me.rightEdges, 0)

	me.ctrl.ResizeToFitParent(WmSize{
		m: Wm{
			WParam: win.WPARAM(co.SIZE_RESTORED),
			LParam: win.MakeLParam(uint16(me.initialParentCx), 0),
		},
	})
	return me.ctrl
}

// Returns the number of parts.
func (me *_StatusBarPartCollection) Count() int {
	return len(me.partsData)
}

// Returns the part at the given index.
//
// Does not perform bound checking.
func (me *_StatusBarPartCollection) Get(index int) *StatusBarPart {
	return _NewStatusBarPart(me.ctrl, index)
}

// Sets the text of all parts at once.
func (me *_StatusBarPartCollection) SetTexts(texts ...string) *StatusBar {
	if len(texts) > len(me.partsData) {
		panic("Number of texts is bigger than the number of parts.")
	}
	for i, txt := range texts {
		me.Get(i).SetText(txt)
	}
	return me.ctrl
}

func (me *_StatusBarPartCollection) cacheInitialParentCx() {
	if me.initialParentCx == 0 {
		rc := me.ctrl.parent.Hwnd().GetClientRect()
		me.initialParentCx = int(rc.Right) // initial width of parent's client area
	}
}

//------------------------------------------------------------------------------

// A single StatusBar part.
type StatusBarPart struct {
	ctrl  *StatusBar
	index int
}

// Constructor.
func _NewStatusBarPart(ctrl *StatusBar, index int) *StatusBarPart {
	return &StatusBarPart{
		ctrl:  ctrl,
		index: index,
	}
}

// Retrieves the HICON of the part.
//
// The icon is shared, the StatusBar doesn't own it.
func (me *StatusBarPart) Icon() win.HICON {
	return win.HICON(
		me.ctrl.Hwnd().SendMessage(co.WM(co.SB_GETICON), win.WPARAM(me.index), 0),
	)
}

// Puts the HICON on the part.
//
// The icon is shared, the StatusBar doesn't own it.
func (me *StatusBarPart) SetIcon(hIcon win.HICON) *StatusBarPart {
	me.ctrl.Hwnd().SendMessage(co.WM(co.SB_SETICON),
		win.WPARAM(me.index), win.LPARAM(hIcon))
	return me
}

// Sets the text of the part.
func (me *StatusBarPart) SetText(text string) *StatusBarPart {
	ret := me.ctrl.Hwnd().SendMessage(co.WM(co.SB_SETTEXT),
		win.MakeWParam(win.MakeWord(uint8(me.index), 0), 0),
		win.LPARAM(unsafe.Pointer(win.Str.ToUint16Ptr(text))))
	if ret == 0 {
		panic(fmt.Sprintf("SB_SETTEXT failed: \"%s\".", text))
	}
	return me
}

// Retrieves the text of the part.
func (me *StatusBarPart) Text() string {
	len := uint16(me.ctrl.Hwnd().
		SendMessage(co.WM(co.SB_GETTEXTLENGTH), win.WPARAM(me.index), 0),
	)
	if len == 0 {
		return ""
	}

	buf := make([]uint16, len+1)
	me.ctrl.Hwnd().SendMessage(co.WM(co.SB_GETTEXT),
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&buf[0])))
	return win.Str.FromUint16Slice(buf)
}

//------------------------------------------------------------------------------

// Internal data kept to each added part.
type _StatusBarPartData struct {
	sizePixels   int
	resizeWeight int
}

func (me *_StatusBarPartData) IsFixedWidth() bool {
	return me.resizeWeight == 0
}

//------------------------------------------------------------------------------

// StatusBar control notifications.
type _EventsStatusBar struct {
	ctrl *_NativeControlBase
}

// Constructor.
func _NewEventsStatusBar(ctrl *_NativeControlBase) *_EventsStatusBar {
	return &_EventsStatusBar{
		ctrl: ctrl,
	}
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-click-status-bar
func (me *_EventsStatusBar) NmClick(userFunc func(p *win.NMMOUSE) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-dblclk-status-bar
func (me *_EventsStatusBar) NmDblClk(userFunc func(p *win.NMMOUSE) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rclick-status-bar
func (me *_EventsStatusBar) NmRClick(userFunc func(p *win.NMMOUSE) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-status-bar
func (me *_EventsStatusBar) NmRDblClk(userFunc func(p *win.NMMOUSE) bool) {
	me.ctrl.parent.On().addNfyRet(me.ctrl.CtrlId(), co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		return _global.BoolToUintptr(userFunc((*win.NMMOUSE)(p)))
	})
}

// https://docs.microsoft.com/en-us/windows/win32/controls/sbn-simplemodechange
func (me *_EventsStatusBar) SbnSimpleModeChange(userFunc func(p *win.NMMOUSE)) {
	me.ctrl.parent.On().addNfy(me.ctrl.CtrlId(), co.NM(co.SBN_SIMPLEMODECHANGE), func(p unsafe.Pointer) {
		userFunc((*win.NMMOUSE)(p))
	})
}
