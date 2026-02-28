//go:build windows

package ui

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
)

// Native [status bar] control.
//
// [status bar]: https://learn.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBar struct {
	_BaseCtrl
	events      StatusBarEvents
	iconCache16 _IconCacheHicon
	partsData   []_StatusBarPartData
	rightEdges  []int32 // buffer to speed up ResizeToFitParent() calls
}

type _StatusBarPartData struct {
	sizePixels   int
	resizeWeight int
}

func (me *_StatusBarPartData) IsFixedWidth() bool {
	return me.resizeWeight == 0
}

// Creates a new [StatusBar] with [win.CreateWindowEx].
//
// Example:
//
//	runtime.LockOSThread()
//
//	wnd := ui.NewMain(
//		ui.OptsMain().
//			Title("Hello world"),
//	)
//	sbar := ui.NewStatusBar(
//		wnd
//		win.OptsStatusBar().
//			FixedPart(ui.DpiX(100), "First").
//			FlexPart(1, "Second"),
//	)
//	wnd.RunAsMain()
func NewStatusBar(parent Parent, opts *VarOptsStatusBar) *StatusBar {
	if len(opts.parts) == 0 {
		panic("Cannot create a StatusBar control without parts.")
	}

	setUniqueCtrlId(&opts.ctrlId)
	me := &StatusBar{
		_BaseCtrl:   newBaseCtrl(opts.ctrlId),
		events:      StatusBarEvents{opts.ctrlId, &parent.base().userEvents},
		iconCache16: newIconCacheHicon(),
		partsData:   make([]_StatusBarPartData, 0, len(opts.parts)),
		rightEdges:  make([]int32, len(opts.parts)), // initially filled with zeros
	}

	parent.base().beforeUserEvents.wmCreateOrInitdialog(func() {
		sbStyle := co.WS_CHILD | co.WS_VISIBLE | co.WS(co.SBARS_TOOLTIPS)
		parentStyle, _ := parent.Hwnd().Style()
		isParentResizable := (parentStyle&co.WS_MAXIMIZEBOX) != 0 ||
			(parentStyle&co.WS_SIZEBOX) != 0
		if isParentResizable {
			sbStyle |= co.WS(co.SBARS_SIZEGRIP)
		}

		me.createWindow(co.WS_EX_NONE, "msctls_statusbar32", "",
			sbStyle, win.POINT{}, win.SIZE{}, parent, false)

		me.addParts(opts.parts)
	})

	parent.base().beforeUserEvents.wm(co.WM_SIZE, func(p Wm) {
		me.resizeToFitParent(WmSize{p})
	})

	parent.base().afterUserEvents.wm(co.WM_DESTROY, func(_ Wm) {
		me.iconCache16.Release()
	})

	return me
}

func (me *StatusBar) addParts(parts []_StatusBarOptPart) {
	for _, part := range parts {
		if part.width < 0 {
			panic("StatusBar part width cannot be negative.")
		} else if part.flex < 0 {
			panic("StatusBar part flex cannot be negative.")
		}

		me.partsData = append(me.partsData, _StatusBarPartData{
			sizePixels:   part.width,
			resizeWeight: part.flex,
		})
	}

	hParent, _ := me.Hwnd().GetParent()
	rc, _ := hParent.GetClientRect()
	me.resizeToFitParent(WmSize{ // force the creation of the parts, so we can set text and icon
		Raw: Wm{
			WParam: win.WPARAM(co.SIZE_REQ_RESTORED),
			LParam: win.MAKELPARAM(uint16(rc.Right-rc.Left), 0),
		},
	})

	for i, part := range parts {
		me.Part(i).SetText(part.text)
		if part.icon.isValid() {
			me.Part(i).SetIcon(part.icon)
		}
	}
}

func (me *StatusBar) resizeToFitParent(parm WmSize) {
	if parm.Request() == co.SIZE_REQ_MINIMIZED || me.Hwnd() == 0 {
		return
	}
	me.Hwnd().SendMessage(co.WM_SIZE, 0, 0) // tell status bar to fit parent

	if len(me.partsData) == 0 {
		return // no parts added, nothing else to do
	}

	cx := int(parm.ClientAreaSize().Cx) // available width

	totalWeight := 0 // total weight of all variable-width parts
	cxVariable := cx // total width to be divided among variable-width parts
	for i := range me.partsData {
		if me.partsData[i].IsFixedWidth() {
			cxVariable -= me.partsData[i].sizePixels
		} else {
			totalWeight += me.partsData[i].resizeWeight
		}
	}

	cxTotal := cx
	for i := len(me.partsData) - 1; i >= 0; i-- { // fill right edges array with the right edge of each part
		me.rightEdges[i] = int32(cxTotal)
		if me.partsData[i].IsFixedWidth() {
			cxTotal -= me.partsData[i].sizePixels
		} else {
			cxTotal -= (cxVariable / totalWeight) * me.partsData[i].resizeWeight
		}
	}
	me.Hwnd().SendMessage(co.SB_SETPARTS,
		win.WPARAM(int32(len(me.rightEdges))),
		win.LPARAM(unsafe.Pointer(&me.rightEdges[0])))
}

// Exposes all the control notifications the can be handled.
//
// Panics if called after the control has been created.
func (me *StatusBar) On() *StatusBarEvents {
	me.panicIfAddingEventAfterCreated()
	return &me.events
}

// Returns all parts.
func (me *StatusBar) AllParts() []StatusBarPart {
	nParts := me.PartCount()
	parts := make([]StatusBarPart, 0, nParts)
	for i := 0; i < nParts; i++ {
		parts = append(parts, me.Part(i))
	}
	return parts
}

// Returns the last part.
func (me *StatusBar) LastPart() StatusBarPart {
	return me.Part(me.PartCount() - 1)
}

// Returns the part at the given zero-based index.
func (me *StatusBar) Part(index int) StatusBarPart {
	return StatusBarPart{me, int32(index)}
}

// Returns the number of parts.
func (me *StatusBar) PartCount() int {
	return len(me.partsData)
}

// Options for [NewStatusBar]; returned by [OptsStatusBar].
type VarOptsStatusBar struct {
	ctrlId uint16
	parts  []_StatusBarOptPart
}

type _StatusBarOptPart struct {
	width int
	flex  int
	text  string
	icon  Ico
}

// Options for [NewStatusBar].
func OptsStatusBar() *VarOptsStatusBar {
	return &VarOptsStatusBar{}
}

// Control ID. Must be unique within a same parent window.
//
// Defaults to an auto-generated ID.
func (o *VarOptsStatusBar) CtrlId(id uint16) *VarOptsStatusBar { o.ctrlId = id; return o }

// Adds a fixed-width part to the StatusBar, with the given width.
//
// Example:
//
//	win.OptsStatusBar().
//		FixedPart(ui.DpiX(100), "Foo")
func (o *VarOptsStatusBar) FixedPart(cx int, text string) *VarOptsStatusBar {
	o.parts = append(o.parts, _StatusBarOptPart{cx, 0, text, Ico{}})
	return o
}

// Adds a fixed-width part to the StatusBar, with the given width. Also adds an
// icon, either from the resource or from a shell file extension.
//
// Example:
//
//	win.OptsStatusBar().
//		FixedPartIcon(ui.DpiX(100), "Foo", ui.IcoId(101))
func (o *VarOptsStatusBar) FixedPartIcon(cx int, text string, icon Ico) *VarOptsStatusBar {
	o.parts = append(o.parts, _StatusBarOptPart{cx, 0, text, icon})
	return o
}

// Adds a variable-sized part to the StatusBar, which will resize according to
// the remaining space.
//
// Example:
//
//	win.OptsStatusBar().
//		FlexPart(1, "Foo")
func (o *VarOptsStatusBar) FlexPart(flex int, text string) *VarOptsStatusBar {
	o.parts = append(o.parts, _StatusBarOptPart{0, flex, text, Ico{}})
	return o
}

// Adds a variable-sized part to the StatusBar, which will resize according to
// the remaining space. Also adds an icon, either from the resource or from a
// shell file extension.
//
// Example:
//
//	win.OptsStatusBar().
//		FlexPartIcon(1, "Foo", ui.IcoId(101))
func (o *VarOptsStatusBar) FlexPartIcon(flex int, text string, icon Ico) *VarOptsStatusBar {
	o.parts = append(o.parts, _StatusBarOptPart{0, flex, text, icon})
	return o
}

// Native [status bar] control events.
//
// You cannot create this object directly, it will be created automatically
// by the owning control.
//
// [status bar]: https://learn.microsoft.com/en-us/windows/win32/controls/status-bars
type StatusBarEvents struct {
	ctrlId       uint16
	parentEvents *WindowEvents
}

// [NM_CLICK] message handler.
//
// [NM_CLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-click-status-bar
func (me *StatusBarEvents) NmClick(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_CLICK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_DBLCLK] message handler.
//
// [NM_DBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-dblclk-status-bar
func (me *StatusBarEvents) NmDblClk(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_DBLCLK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RCLICK] message handler.
//
// [NM_RCLICK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rclick-status-bar
func (me *StatusBarEvents) NmRClick(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RCLICK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [NM_RDBLCLK] message handler.
//
// [NM_RDBLCLK]: https://learn.microsoft.com/en-us/windows/win32/controls/nm-rdblclk-status-bar
func (me *StatusBarEvents) NmRDblClk(fun func(p *win.NMMOUSE) bool) {
	me.parentEvents.WmNotify(me.ctrlId, co.NM_RDBLCLK, func(p unsafe.Pointer) uintptr {
		return utl.BoolToUintptr(fun((*win.NMMOUSE)(p)))
	})
}

// [SBN_SIMPLEMODECHANGE] message handler.
//
// [SBN_SIMPLEMODECHANGE]: https://learn.microsoft.com/en-us/windows/win32/controls/sbn-simplemodechange
func (me *StatusBarEvents) SbnSimpleModeChange(fun func(p *win.NMMOUSE)) {
	me.parentEvents.WmNotify(me.ctrlId, co.SBN_SIMPLEMODECHANGE, func(p unsafe.Pointer) uintptr {
		fun((*win.NMMOUSE)(p))
		return me.parentEvents.defProcVal
	})
}
