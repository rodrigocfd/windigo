//go:build windows

package win

import (
	"encoding/binary"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [BITMAP] struct.
//
// [BITMAP]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmap
type BITMAP struct {
	bmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       *byte
}

// Useful when [capturing an image].
//
// [capturing an image]: https://learn.microsoft.com/en-gb/windows/win32/gdi/capturing-an-image
func (bm *BITMAP) CalcBitmapSize(bitCount uint16) uint {
	return uint(((bm.BmWidth*int32(bitCount) + 31) / 32) * 4 * bm.BmHeight)
}

// [BLENDFUNCTION] struct.
//
// [BLENDFUNCTION]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-blendfunction
type BLENDFUNCTION struct {
	BlendOp             byte
	BlendFlags          byte
	SourceConstantAlpha byte
	AlphaFormat         byte
}

// [BITMAPFILEHEADER] struct.
//
// ⚠️ You must call [BITMAPFILEHEADER.SetBfType] to initialize the struct.
//
// # Example
//
//	var bfh win.BITMAPFILEHEADER
//	bfh.SetBfType()
//
// [BITMAPFILEHEADER]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapfileheader
type BITMAPFILEHEADER struct {
	data [14]byte // sizeof(BITMAPFILEHEADER) packed
}

// Sets the struct type for bitmap, correctly initializing it.
func (bfh *BITMAPFILEHEADER) SetBfType() {
	binary.LittleEndian.PutUint32(bfh.data[0:], 0x4d42) // https://learn.microsoft.com/en-gb/windows/win32/gdi/capturing-an-image
}

func (bfh *BITMAPFILEHEADER) BfSize() uint32 {
	return binary.LittleEndian.Uint32(bfh.data[2:])
}
func (bfh *BITMAPFILEHEADER) SetBfSize(val uint32) {
	binary.LittleEndian.PutUint32(bfh.data[2:], val)
}

func (bfh *BITMAPFILEHEADER) BfOffBits() uint32 {
	return binary.LittleEndian.Uint32(bfh.data[10:])
}
func (bfh *BITMAPFILEHEADER) SetBfOffBits(val uint32) {
	binary.LittleEndian.PutUint32(bfh.data[10:], val)
}

func (bfh *BITMAPFILEHEADER) Serialize() []byte { return bfh.data[:] }

// [BITMAPINFO] struct.
//
// ⚠️ You must call [BITMAPINFOHEADER.SetBiSize] on BmiHeader field, to
// initialize the struct.
//
// # Example
//
//	var bi win.BITMAPINFO
//	bi.BmiHeader.SetBiSize()
//
// [BITMAPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfo
type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors [1]RGBQUAD
}

// [BITMAPINFOHEADER] struct.
//
// ⚠️ You must call [BITMAPINFOHEADER.SetBiSize] to initialize the struct.
//
// # Example
//
//	var bih win.BITMAPINFOHEADER
//	bih.SetBiSize()
//
// [BITMAPINFOHEADER]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
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

// Sets the biSize field to the size of the struct, correctly initializing it.
func (bih *BITMAPINFOHEADER) SetBiSize() {
	bih.biSize = uint32(unsafe.Sizeof(*bih))
}

func (bih *BITMAPINFOHEADER) Serialize() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(bih)), unsafe.Sizeof(*bih))
}

// [COLORREF] struct.
//
// Specifies an RGB color.
//
// [COLORREF]: https://learn.microsoft.com/en-us/windows/win32/gdi/colorref
type COLORREF uint32

const (
	COLORREF_NONE    COLORREF = 0xffff_ffff // No color.
	COLORREF_DEFAULT COLORREF = 0xff00_0000 // Default color.

	COLORREF_DWMA_NONE    COLORREF = 0xffff_fffe // No color in a DWMA context.
	COLORREF_DWMA_DEFAULT COLORREF = 0xffff_ffff // Default color in a DWMA context.
)

// [RGB] macro.
//
// [RGB]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-rgb
func RGB(red, green, blue uint8) COLORREF {
	return COLORREF(uint32(red) | (uint32(green) << 8) | (uint32(blue) << 16))
}

// [GetRValue] macro.
//
// [GetRValue]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getrvalue
func (c COLORREF) Red() uint8 {
	return LOBYTE(LOWORD(uint32(c)))
}

// [GetGValue] macro.
//
// [GetGValue]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getgvalue
func (c COLORREF) Green() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 8))
}

// [GetBValue] macro.
//
// [GetBValue]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getbvalue
func (c COLORREF) Blue() uint8 {
	return LOBYTE(LOWORD(uint32(c) >> 16))
}

// Converts the COLORREF to an RGBQUAD struct.
func (c COLORREF) ToRgbquad() RGBQUAD {
	var rq RGBQUAD
	rq.SetBlue(c.Blue())
	rq.SetGreen(c.Green())
	rq.SetRed(c.Red())
	return rq
}

// [DOCINFO] struct.
//
// ⚠️ You must call [DOCINFO.SetCbSize] to initialize the struct.
//
// [DOCINFO]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-docinfow
type DOCINFO struct {
	cbSize       int32
	LpszDocName  *uint16
	LpszOutput   *uint16
	LpszDataType *uint16
	FwType       co.DIPJ
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (di *DOCINFO) SetCbSize() {
	di.cbSize = int32(unsafe.Sizeof(*di))
}

// [DEVMODE] struct.
//
// ⚠️ You must call [DEVMODE.SetDmSize] to initialize the struct.
//
// # Example
//
//	var dm win.DEVMODE
//	dm.SetDmSize()
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DEVMODE struct {
	dmDeviceName       [utl.CCHDEVICENAME]uint16
	dmSpecVersion      uint16
	DmDriverVersion    uint16
	dmSize             uint16
	DmDriverExtra      uint16
	DmFields           co.DM
	union0             DEVMODE_Printer
	DmColor            co.DMCOLOR
	DmDuplex           co.DMDUP
	DmYResolution      int16
	DmTTOption         co.DMTT
	DmCollate          co.DMCOLLATE
	dmFormName         [utl.CCHFORMNAME]uint16
	DmLogPixels        uint16
	DmBitsPerPel       uint32
	DmPelsWidth        uint32
	DmPelsHeight       uint32
	union1             uint32 // co.DMDISPLAYFLAGS | co.DMNUP
	DmDisplayFrequency uint32
	DmICMMethod        co.DMICMMETHOD
	DmICMIntent        co.DMICM
	DmMediaType        co.DMMEDIA
	DmDitherType       co.DMDITHER
	dmReserved1        uint32
	dmReserved2        uint32
	DmPanningWidth     uint32
	DmPanningHeight    uint32
}

// 1st variation of 1st union of [DEVMODE] struct.
type DEVMODE_Printer struct {
	DmOrientation   co.DMORIENT
	DmPaperSize     co.DMPAPER
	DmPaperLength   int16
	DmPaperWidth    int16
	DmScale         int16
	DmCopies        int16
	DmDefaultSource co.DMBIN
	DmPrintQuality  co.DMRES
}

// 2st variation of 1st union of [DEVMODE] struct.
type DEVMODE_Display struct {
	DmPosition           POINT
	DmDisplayOrientation co.DMDO
	DmDisplayFixedOutput co.DMDFO
}

func (dm *DEVMODE) DmDeviceName() string {
	return wstr.DecodeSlice(dm.dmDeviceName[:])
}
func (dm *DEVMODE) SetDmDeviceName(val string) {
	wstr.EncodeToBuf(val, dm.dmDeviceName[:])
}

// Sets the dmSize field to the size of the struct, correctly initializing it.
// Also sets dmSpecVersion.
func (dm *DEVMODE) SetDmSize() {
	dm.dmSpecVersion = utl.DM_SPECVERSION
	dm.dmSize = uint16(unsafe.Sizeof(*dm))
}

// Returns the 1st variation of the 1st union.
func (dm *DEVMODE) Printer() *DEVMODE_Printer {
	return &dm.union0
}

// Returns the 2nd variation of the 1st union.
func (dm *DEVMODE) Display() *DEVMODE_Display {
	return (*DEVMODE_Display)(unsafe.Pointer(&dm.union0))
}

func (dm *DEVMODE) DmFormName() string {
	return wstr.DecodeSlice(dm.dmFormName[:])
}
func (dm *DEVMODE) SetDmFormName(val string) {
	wstr.EncodeToBuf(val, dm.dmFormName[:])
}

// Returns the 1st variation of the 2nd union.
func (dm *DEVMODE) DmDisplayFlags() *co.DMDISPLAYFLAGS {
	return (*co.DMDISPLAYFLAGS)(unsafe.Pointer(&dm.union1))
}

// Returns the 2st variation of the 2nd union.
func (dm *DEVMODE) DmNup() *co.DMNUP {
	return (*co.DMNUP)(unsafe.Pointer(&dm.union1))
}

// [GRADIENT_RECT] struct.
//
// [GRADIENT_RECT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-gradient_rect
type GRADIENT_RECT struct {
	UpperLeft  uint32
	LowerRight uint32
}

// [GRADIENT_TRIANGLE] struct.
//
// [GRADIENT_TRIANGLE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-gradient_triangle
type GRADIENT_TRIANGLE struct {
	Vertex1 uint32
	Vertex2 uint32
	Vertex3 uint32
}

// [LOGBRUSH] struct.
//
// [LOGBRUSH]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logbrush
type LOGBRUSH struct {
	LbStyle co.BRS
	LbColor COLORREF
	LbHatch uintptr // ULONG_PTR
}

// [LOGFONT] struct.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type LOGFONT struct {
	LfHeight         int32
	LfWidth          int32
	LfEscapement     int32
	LfOrientation    int32
	LfWeight         co.FW
	LfItalic         uint8 // This is a BOOL value.
	LfUnderline      uint8 // This is a BOOL value.
	LfStrikeOut      uint8 // This is a BOOL value.
	LfCharSet        co.CHARSET
	LfOutPrecision   co.OUT_PRECIS
	LfClipPrecision  co.CLIP_PRECIS
	LfQuality        co.QUALITY
	lfPitchAndFamily uint8 // combination of co.PITCH and co.FF
	lfFaceName       [utl.LF_FACESIZE]uint16
}

func (lf *LOGFONT) LfFaceName() string {
	return wstr.DecodeSlice(lf.lfFaceName[:])
}
func (lf *LOGFONT) SetLfFaceName(val string) {
	wstr.EncodeToBuf(val, lf.lfFaceName[:])
}

func (lf *LOGFONT) Pitch() co.PITCH {
	return co.PITCH(lf.lfPitchAndFamily & 0b1111)
}
func (lf *LOGFONT) SetPitch(val co.PITCH) {
	lf.lfPitchAndFamily &^= 0b1111 // clear bits
	lf.lfPitchAndFamily |= uint8(val & 0b1111)
}

func (lf *LOGFONT) Family() co.FF {
	return co.FF(lf.lfPitchAndFamily & 0b1111_0000)
}
func (lf *LOGFONT) SetFamily(val co.FF) {
	lf.lfPitchAndFamily &^= 0b1111_0000 // clear bits
	lf.lfPitchAndFamily |= uint8(val & 0b1111_0000)
}

// [LOGPEN] struct.
//
// [LOGPEN]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logpen
type LOGPEN struct {
	LopnStyle co.PS
	LopnWidth POINT
	LopnColor COLORREF
}

// [PALETTEENTRY] struct.
//
// [PALETTEENTRY]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-paletteentry
type PALETTEENTRY struct {
	PeRed   uint8
	PeGreen uint8
	PeBlue  uint8
	PeFlags co.PC
}

// [PIXELFORMATDESCRIPTOR] struct.
//
// ⚠️ You must call [PIXELFORMATDESCRIPTOR.SetNSize] to initialize the struct.
//
// [PIXELFORMATDESCRIPTOR]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-pixelformatdescriptor
type PIXELFORMATDESCRIPTOR struct {
	nSize           uint16
	nVersion        uint16
	DwFlags         co.PFD
	IPixelType      co.PFD_TYPE
	CColorBits      uint8
	CRedBits        uint8
	CRedShift       uint8
	CGreenBits      uint8
	CGreenShift     uint8
	CBlueBits       uint8
	CBlueShift      uint8
	CAlphaBits      uint8
	CAlphaShift     uint8
	CAccumBits      uint8
	CAccumRedBits   uint8
	CAccumGreenBits uint8
	CAccumBlueBits  uint8
	CAccumAlphaBits uint8
	CDepthBits      uint8
	CStencilBits    uint8
	CAuxBuffers     uint8
	iLayerType      uint8
	BReserved       uint8
	dwLayerMask     uint32
	DwVisibleMask   uint32
	dwDamageMask    uint32
}

// Sets the nSize field to the size of the struct, correctly initializing it.
// Also sets nVersion.
func (pfd *PIXELFORMATDESCRIPTOR) SetNSize() {
	pfd.nSize = uint16(unsafe.Sizeof(*pfd))
	pfd.nVersion = 1
}

// [RGBQUAD] struct.
//
// [RGBQUAD]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-rgbquad
type RGBQUAD struct {
	data [4]byte
}

func (rq *RGBQUAD) Blue() uint8 {
	return *(*uint8)(unsafe.Pointer(&rq.data[0]))
}
func (rq *RGBQUAD) SetBlue(val uint8) {
	*(*uint8)(unsafe.Pointer(&rq.data[0])) = val
}

func (rq *RGBQUAD) Green() uint8 {
	return *(*uint8)(unsafe.Pointer(&rq.data[1]))
}
func (rq *RGBQUAD) SetGreen(val uint8) {
	*(*uint8)(unsafe.Pointer(&rq.data[1])) = val
}

func (rq *RGBQUAD) Red() uint8 {
	return *(*uint8)(unsafe.Pointer(&rq.data[2]))
}
func (rq *RGBQUAD) SetRed(val uint8) {
	*(*uint8)(unsafe.Pointer(&rq.data[2])) = val
}

func (rq *RGBQUAD) ToColorref() COLORREF {
	return RGB(rq.Red(), rq.Green(), rq.Blue())
}

// [TEXTMETRIC] struct.
//
// [TEXTMETRIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
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
	tmPitchAndFamily   uint8 // combination of co.TMPF and co.FF
	TmCharSet          co.CHARSET
}

func (tm *TEXTMETRIC) Pitch() co.TMPF {
	return co.TMPF(tm.tmPitchAndFamily & 0b1111)
}
func (tm *TEXTMETRIC) SetPitch(val co.TMPF) {
	tm.tmPitchAndFamily &^= 0b1111 // clear bits
	tm.tmPitchAndFamily |= uint8(val & 0b1111)
}

func (tm *TEXTMETRIC) Family() co.FF {
	return co.FF(tm.tmPitchAndFamily & 0b1111_0000)
}
func (tm *TEXTMETRIC) SetFamily(val co.FF) {
	tm.tmPitchAndFamily &^= 0b1111_0000 // clear bits
	tm.tmPitchAndFamily |= uint8(val & 0b1111_0000)
}

// [TRIVERTEX] struct.
//
// [TRIVERTEX]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-trivertex
type TRIVERTEX struct {
	X     int32
	Y     int32
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}
