package api

import (
	"syscall"
	"unsafe"
	"wingows/api/proc"
	c "wingows/consts"
)

type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         c.FW
	LfItalic         uint8
	LfUnderline      uint8
	LfStrikeOut      uint8
	LfCharSet        uint8
	LfOutPrecision   uint8
	LfClipPrecision  uint8
	LfQuality        uint8
	LfPitchAndFamily uint8
	LfFaceName       [32]uint16 // LF_FACESIZE
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
	return syscall.UTF16ToString(lf.LfFaceName[:])
}
