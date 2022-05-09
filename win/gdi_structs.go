//go:build windows

package win

import (
	"encoding/binary"
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

func (bm *BITMAP) CalcBitmapSize(bitCount uint16) int {
	// https://docs.microsoft.com/en-gb/windows/win32/gdi/capturing-an-image
	return int(((bm.BmWidth*int32(bitCount) + 31) / 32) * 4 * bm.BmHeight)
}

// ⚠️ You must call SetBfType() to initialize the struct.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapfileheader
type BITMAPFILEHEADER struct {
	data [14]byte // sizeof(BITMAPFILEHEADER) packed
}

func (bfh *BITMAPFILEHEADER) SetBfType() { binary.LittleEndian.PutUint32(bfh.data[0:], 0x4d42) } // https://docs.microsoft.com/en-gb/windows/win32/gdi/capturing-an-image

func (bfh *BITMAPFILEHEADER) BfSize() uint32       { return binary.LittleEndian.Uint32(bfh.data[2:]) }
func (bfh *BITMAPFILEHEADER) SetBfSize(val uint32) { binary.LittleEndian.PutUint32(bfh.data[2:], val) }

func (bfh *BITMAPFILEHEADER) BfOffBits() uint32 { return binary.LittleEndian.Uint32(bfh.data[10:]) }
func (bfh *BITMAPFILEHEADER) SetBfOffBits(val uint32) {
	binary.LittleEndian.PutUint32(bfh.data[10:], val)
}

func (bfh *BITMAPFILEHEADER) Serialize() []byte { return bfh.data[:] }

// ⚠️ You must call BmiHeader.SetBiSize() to initialize the struct.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfo
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}

// ⚠️ You must call SetBiSize() to initialize the struct.
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

func (bih *BITMAPINFOHEADER) Serialize() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(bih)), unsafe.Sizeof(*bih))
}

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
