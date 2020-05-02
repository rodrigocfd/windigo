package winffi

import (
	"winffi/consts"
)

type MONITORINFOEX struct {
	Size      uint32
	RcMonitor RECT
	RcWork    RECT
	Flags     uint32
	Device    [consts.CCHDEVICENAME]uint16
}

type MSG struct {
	HWnd   HWND
	Msg    uint32
	WParam WPARAM
	LParam LPARAM
	Time   uint32
	Pt     POINT
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

type OSVERSIONINFOEX struct {
	OsVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CSDVersion        [128]uint16
	ServicePackMajor  uint16
	ServicePackMinor  uint16
	SuiteMask         uint16
	ProductType       uint8
	Reserve           uint8
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
