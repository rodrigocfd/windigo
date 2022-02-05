package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmap
type BITMAP struct {
	bmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       *byte
}

// ⚠️ You must call BmiHeader.SetBiSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfo
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}

// ⚠️ You must call SetBiSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BITMAPINFOHEADER struct {
	biSize          uint32
	BiWidth         int32
	BiHeight        int32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   co.BI
	BiSizeImage     uint32
	BiXPelsPerMeter int32
	BiYPelsPerMeter int32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

func (bih *BITMAPINFOHEADER) SetBiSize() { bih.biSize = uint32(unsafe.Sizeof(*bih)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logbrush
type LOGBRUSH struct {
	LbStyle co.BRS
	LbColor COLORREF
	LbHatch uintptr // ULONG_PTR
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         co.FW
	LfItalic         uint8
	LfUnderline      uint8
	LfStrikeOut      uint8
	LfCharSet        uint8
	LfOutPrecision   uint8
	LfClipPrecision  uint8
	LfQuality        uint8
	LfPitchAndFamily uint8
	lfFaceName       [_LF_FACESIZE]uint16
}

func (lf *LOGFONT) LfFaceName() string { return Str.FromNativeSlice(lf.lfFaceName[:]) }
func (lf *LOGFONT) SetLfFaceName(val string) {
	copy(lf.lfFaceName[:], Str.ToNativeSlice(Str.Substr(val, 0, len(lf.lfFaceName)-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logpen
type LOGPEN struct {
	LopnStyle co.PS
	LopnWidth POINT
	LopnColor COLORREF
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-rgbquad
type RGBQUAD struct {
	data [4]byte
}

func (rq *RGBQUAD) Blue() uint8       { return *(*uint8)(unsafe.Pointer(&rq.data[0])) }
func (rq *RGBQUAD) SetBlue(val uint8) { *(*uint8)(unsafe.Pointer(&rq.data[0])) = val }

func (rq *RGBQUAD) Green() uint8       { return *(*uint8)(unsafe.Pointer(&rq.data[1])) }
func (rq *RGBQUAD) SetGreen(val uint8) { *(*uint8)(unsafe.Pointer(&rq.data[1])) = val }

func (rq *RGBQUAD) Red() uint8       { return *(*uint8)(unsafe.Pointer(&rq.data[2])) }
func (rq *RGBQUAD) SetRed(val uint8) { *(*uint8)(unsafe.Pointer(&rq.data[2])) = val }

func (rq *RGBQUAD) ToColorref() COLORREF { return RGB(rq.Red(), rq.Green(), rq.Blue()) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type TEXTMETRIC struct {
	TmHeight           uint32
	TmAscent           uint32
	TmDescent          uint32
	TmInternalLeading  uint32
	TmExternalLeading  uint32
	TmAveCharWidth     uint32
	TmMaxCharWidth     uint32
	TmWeight           uint32
	TmOverhang         uint32
	TmDigitizedAspectX uint32
	TmDigitizedAspectY uint32
	TmFirstChar        uint16
	TmLastChar         uint16
	TmDefaultChar      uint16
	TmBreakChar        uint16
	TmItalic           uint8
	TmUnderlined       uint8
	TmStruckOut        uint8
	TmPitchAndFamily   uint8
	TmCharSet          co.CHARSET
}
