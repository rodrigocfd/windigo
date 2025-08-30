//go:build windows

package ui

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// An item from a [header].
//
// [header]: https://learn.microsoft.com/en-us/windows/win32/controls/header-controls
type HeaderItem struct {
	owner *Header
	index int32
}

// Returns the zero-based index of the item.
func (me HeaderItem) Index() int {
	return int(me.index)
}

// Retrieves the text justification with [HDM_GETITEM].
//
// Possible values:
//   - [co.HDF_LEFT]
//   - [co.HDF_CENTER]
//   - [co.HDF_RIGHT]
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me HeaderItem) Justification() co.HDF {
	hdi := win.HDITEM{
		Mask: co.HDI_FORMAT,
	}
	me.owner.hWnd.SendMessage(co.HDM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))

	return hdi.Fmt & (co.HDF_LEFT | co.HDF_CENTER | co.HDF_RIGHT) // restrict bits
}

// Retrieves the order of the item with [HDM_GETITEM].
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me HeaderItem) Order() int {
	hdi := win.HDITEM{
		Mask: co.HDI_ORDER,
	}
	me.owner.hWnd.SendMessage(co.HDM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))

	return int(hdi.IOrder)
}

// Sets the text justification with [HDM_SETITEM].
//
// Possible values:
//   - [co.HDF_LEFT]
//   - [co.HDF_CENTER]
//   - [co.HDF_RIGHT]
//
// Returns the same item, so further operations can be chained.
//
// [HDM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setitem
func (me HeaderItem) SetJustification(hdf co.HDF) HeaderItem {
	hdi := win.HDITEM{
		Mask: co.HDI_FORMAT,
	}
	me.owner.hWnd.SendMessage(co.HDM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))

	hdi.Fmt &^= (co.HDF_LEFT | co.HDF_CENTER | co.HDF_RIGHT)        // remove bits
	hdi.Fmt |= (hdf & (co.HDF_LEFT | co.HDF_CENTER | co.HDF_RIGHT)) // restrict bits
	me.owner.hWnd.SendMessage(co.HDM_SETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))

	return me
}

// Sets the displayed sort arrow with [HDM_SETITEM].
//
// Possible values:
//   - [co.HDF_NONE]
//   - [co.HDF_SORTUP]
//   - [co.HDF_SORTDOWN]
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [HDM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setitem
func (me HeaderItem) SetSortArrow(hdf co.HDF) HeaderItem {
	count := me.owner.Items.Count()
	for i := uint(0); i < count; i++ {
		hdi := win.HDITEM{
			Mask: co.HDI_FORMAT,
		}
		me.owner.hWnd.SendMessage(co.HDM_GETITEM,
			win.WPARAM(i), win.LPARAM(unsafe.Pointer(&hdi))) // retrieve current style

		hdi.Fmt &^= (co.HDF_SORTDOWN | co.HDF_SORTUP) // remove bits

		if i == uint(me.index) { // only our item will be set
			hdi.Fmt |= (hdf & (co.HDF_SORTDOWN | co.HDF_SORTUP)) // restrict bits
		}
		me.owner.hWnd.SendMessage(co.HDM_SETITEM,
			win.WPARAM(i), win.LPARAM(unsafe.Pointer(&hdi)))
	}
	return me
}

// Sets the text with [HDM_SETITEM].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [HDM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setitem
func (me HeaderItem) SetText(text string) HeaderItem {
	hdi := win.HDITEM{
		Mask: co.HDI_TEXT,
	}

	var wText wstr.BufEncoder
	hdi.SetPszText(wText.Slice(text))

	ret, err := me.owner.hWnd.SendMessage(co.HDM_SETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("HDM_SETITEM %d to \"%s\" failed.", me.index, text))
	}

	return me
}

// Sets the width of the item with [HDM_SETITEM].
//
// Returns the same item, so further operations can be chained.
//
// Panics on error.
//
// [HDM_SETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-setitem
func (me HeaderItem) SetWidth(width int) HeaderItem {
	hdi := win.HDITEM{
		Mask: co.HDI_WIDTH,
		Cxy:  int32(width),
	}

	me.owner.hWnd.SendMessage(co.HDM_SETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))
	return me
}

// Retrieves the displayed sort arrow with [HDM_GETITEM].
//
// Possible values:
//   - [co.HDF_NONE]
//   - [co.HDF_SORTUP]
//   - [co.HDF_SORTDOWN]
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me HeaderItem) SortArrow() co.HDF {
	hdi := win.HDITEM{
		Mask: co.HDI_FORMAT,
	}
	me.owner.hWnd.SendMessage(co.HDM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))

	return hdi.Fmt & (co.HDF_SORTDOWN | co.HDF_SORTUP) // restrict bits
}

// Retrieves the text with [HDM_GETITEM].
//
// Panics on error.
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me HeaderItem) Text() string {
	hdi := win.HDITEM{
		Mask: co.HDI_TEXT,
	}

	var wBuf wstr.BufDecoder
	wBuf.Alloc(wstr.BUF_MAX)
	hdi.SetPszText(wBuf.HotSlice())

	ret, err := me.owner.hWnd.SendMessage(co.HDM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("HDM_GETITEM %d failed.", me.index))
	}

	return wBuf.String()
}

// Retrieves the width of the item with [HDM_GETITEM].
//
// Panics on error.
//
// [HDM_GETITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/hdm-getitem
func (me HeaderItem) Width() int {
	hdi := win.HDITEM{
		Mask: co.HDI_WIDTH,
	}

	ret, err := me.owner.hWnd.SendMessage(co.HDM_GETITEM,
		win.WPARAM(me.index), win.LPARAM(unsafe.Pointer(&hdi)))
	if err != nil || ret == 0 {
		panic(fmt.Sprintf("HDM_GETITEM %d failed.", me.index))
	}
	return int(hdi.Cxy)
}
