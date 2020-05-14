/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package api

import (
	c "wingows/consts"
)

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

type NMLVGETINFOTIP struct {
	Hdr        NMHDR
	DwFlags    uint32
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
