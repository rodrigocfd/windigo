package api

import c "winffi/consts"

type ACCEL struct {
	virt uint8
	key  uint16
	cmd  uint16
}

type CREATESTRUCT struct {
	CreateParams    uintptr
	Instance        HINSTANCE
	Menu            HMENU
	Parent          HWND
	Cy, Cx, Y, X    int32
	Style           int32
	Name, ClassName uintptr
	ExStyle         uint32
}

type MONITORINFOEX struct {
	Size      uint32
	RcMonitor RECT
	RcWork    RECT
	Flags     uint32
	Device    [c.CCHDEVICENAME]uint16
}

type NMHDR struct {
	HWndFrom HWND
	IdFrom   uintptr
	Code     uint32
}

type NONCLIENTMETRICS struct {
	Size            uint32
	BorderWidth     int32
	ScrollWidth     int32
	ScrollHeight    int32
	CaptionWidth    int32
	CaptionHeight   int32
	CaptionFont     LOGFONT
	SmCaptionWidth  int32
	SmCaptionHeight int32
	SmCaptionFont   LOGFONT
	MenuWidth       int32
	MenuHeight      int32
	MenuFont        LOGFONT
	StatusFont      LOGFONT
	MessageFont     LOGFONT
}

type POINT struct {
	X, Y int32
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type SIZE struct {
	Cx, Cy int32
}
