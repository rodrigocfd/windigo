/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package api

import (
	c "wingows/consts"
)

type LITEM struct {
	Mask      c.LIF
	ILink     int32
	State     c.LIS
	StateMask c.LIS
	SzID      [48]uint16            // MAX_LINKID_TEXT
	SzUrl     [2048 + 32 + 3]uint16 // L_MAX_URL_LENGTH
}

type LVCOLUMN struct {
	Mask       c.LVCF
	Fmt        int32
	Cx         int32
	PszText    *uint16
	CchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
	CxMin      int32
	CxDefault  int32
	CxIdeal    int32
}

type LVFINDINFO struct {
	Flags       c.LVFI
	Psz         *uint16
	LParam      LPARAM
	Pt          POINT
	VkDirection uint32
}

type LVITEM struct {
	Mask       c.LVIF
	IItem      int32
	ISubItem   int32
	State      c.LVIS
	StateMask  c.LVIS
	PszText    *uint16
	CchTextMax int32
	IImage     int32
	LParam     uintptr
	IIndent    int32
	IGroupId   int32
	CColumns   uint32
	PuColumns  *uint32
	PiColFmt   *int32
	IGroup     int32
}

type NMCUSTOMDRAW struct {
	Hdr         NMHDR
	DwDrawStage c.CDDS
	Hdc         HDC
	Rc          RECT
	DwItemSpec  uintptr
	UItemState  c.CDIS
	LItemlParam LPARAM
}

type NMITEMACTIVATE struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    LPARAM
	UKeyFlags c.LVKF
}

type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    LPARAM
}

type NMLVCACHEHINT struct {
	Hdr   NMHDR
	IFrom int32
	ITo   int32
}

type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

type NMLVEMPTYMARKUP struct {
	Hdr      NMHDR
	DwFlags  c.EMF
	SzMarkup [2048 + 32 + 3]uint16 // L_MAX_URL_LENGTH
}

type NMLVFINDITEM struct {
	Hdr    NMHDR
	IStart int32
	Lvfi   LVFINDINFO
}

type NMLVGETINFOTIP struct {
	Hdr        NMHDR
	DwFlags    c.LVGIT
	PszText    *uint16
	CchTextMax int32
	IItem      int32
	ISubItem   int32
	LParam     LPARAM
}

type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey c.VK
	Flags uint32
}

type NMLVLINK struct {
	Hdr      NMHDR
	Link     LITEM
	IItem    int32
	ISubItem int32
}

type NMLVODSTATECHANGE struct {
	Hdr       NMHDR
	IFrom     int32
	ITo       int32
	UNewState c.LVIS
	UOldState c.LVIS
}

type NMLVSCROLL struct {
	Hdr NMHDR
	Dx  int32
	Dy  int32
}
