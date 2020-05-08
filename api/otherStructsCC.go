package api

import (
	c "gowinui/consts"
)

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

type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey c.VK
	Flags uint32
}
