package api

import (
	"syscall"
	"unsafe"
	"winffi/api/proc"
	c "winffi/consts"
)

type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         c.FW
	Italic         uint8
	Underline      uint8
	StrikeOut      uint8
	CharSet        uint8
	OutPrecision   uint8
	ClipPrecision  uint8
	Quality        uint8
	PitchAndFamily uint8
	FaceName       [32]uint16 // LF_FACESIZE
}

func (lf *LOGFONT) CreateFontIndirect() HFONT {
	ret, _, _ := syscall.Syscall(proc.CreateFontIndirect.Addr(), 1,
		uintptr(unsafe.Pointer(lf)), 0, 0)
	if ret == 0 {
		panic("CreateFontIndirect failed.")
	}
	return HFONT(ret)
}

func (lf *LOGFONT) GetFaceName() string {
	return syscall.UTF16ToString(lf.FaceName[:])
}
