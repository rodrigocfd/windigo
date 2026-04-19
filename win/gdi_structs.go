//go:build windows

package win

import (
	"encoding/binary"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [BITMAP] struct.
//
// [BITMAP]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmap
type BITMAP struct {
	bmType     int32
	Width      int32
	Height     int32
	WidthBytes int32
	Planes     uint16
	BitsPixel  uint16
	Bits       *byte
}

// Useful when [capturing an image].
//
// [capturing an image]: https://learn.microsoft.com/en-gb/windows/win32/gdi/capturing-an-image
func (bm *BITMAP) CalcBitmapSize(bitCount co.BITCOUNT) int {
	return int(((bm.Width*int32(bitCount) + 31) / 32) * 4 * bm.Height)
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
// Example:
//
//	var bfh win.BITMAPFILEHEADER
//	bfh.SetBfType()
//
// [BITMAPFILEHEADER]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapfileheader
type BITMAPFILEHEADER struct {
	data [14]byte // sizeof(BITMAPFILEHEADER) packed
}

// Sets the internal struct type for bitmap, correctly initializing it.
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
// Example:
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
// Note that the Height field might be [twice] the actual height.
//
// ⚠️ You must call [BITMAPINFOHEADER.SetBiSize] to initialize the struct.
//
// Example:
//
//	var bih win.BITMAPINFOHEADER
//	bih.SetBiSize()
//
// [BITMAPINFOHEADER]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
// [twice]: https://stackoverflow.com/q/5812849
type BITMAPINFOHEADER struct {
	biSize        uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      co.BITCOUNT
	Compression   co.BI
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

// Sets the internal biSize field to the size of the struct, correctly
// initializing it.
func (bih *BITMAPINFOHEADER) SetBiSize() {
	bih.biSize = uint32(unsafe.Sizeof(*bih))
}

func (bih *BITMAPINFOHEADER) Serialize() []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(bih)), unsafe.Sizeof(*bih))
}

// [BITMAPV5HEADER] struct.
//
// ⚠️ You must call [BITMAPV5HEADER.SetBV5Size] to initialize the struct.
//
// Example:
//
//	var bvh win.BITMAPV5HEADER
//	bih.SetBV5Size()
//
// [BITMAPV5HEADER]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapv5header
type BITMAPV5HEADER struct {
	bV5Size       uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      co.BITCOUNT
	Compression   co.BI
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
	RedMask       uint32
	GreenMask     uint32
	BlueMask      uint32
	AlphaMask     uint32
	CSType        co.LCS
	Endpoints     CIEXYZTRIPLE
	GammaRed      uint32
	GammaGreen    uint32
	GammaBlue     uint32
	Intent        co.LCS_GM
	ProfileData   uint32
	ProfileSize   uint32
	bV5Reserved   uint32
}

// Sets the internal bV5Size field to the size of the struct, correctly
// initializing it.
func (bvh *BITMAPV5HEADER) SetBV5Size() {
	bvh.bV5Size = uint32(unsafe.Sizeof(*bvh))
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

// [CIEXYZ] struct.
//
// [CIEXYZ]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-ciexyz
type CIEXYZ struct {
	CiexyzX int32
	CiexyzY int32
	CiexyzZ int32
}

// [CIEXYZTRIPLE] struct.
//
// [CIEXYZTRIPLE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-ciexyztriple
type CIEXYZTRIPLE struct {
	CiexyzRed   CIEXYZ
	CiexyzGreen CIEXYZ
	CiexyzBlue  CIEXYZ
}

// [DOCINFO] struct.
//
// ⚠️ You must call [DOCINFO.SetCbSize] to initialize the struct.
//
// Example:
//
//	var di win.DOCINFO
//	di.SetCbSize()
//
// [DOCINFO]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-docinfow
type DOCINFO struct {
	cbSize       int32
	LpszDocName  *uint16
	LpszOutput   *uint16
	LpszDataType *uint16
	FwType       co.DIPJ
}

// Sets the internal cbSize field to the size of the struct, correctly
// initializing it.
func (di *DOCINFO) SetCbSize() {
	di.cbSize = int32(unsafe.Sizeof(*di))
}

// [DEVMODE] struct.
//
// ⚠️ You must call [DEVMODE.SetDmSize] to initialize the struct.
//
// Example:
//
//	var dm win.DEVMODE
//	dm.SetDmSize()
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DEVMODE struct {
	deviceName       [utl.CCHDEVICENAME]uint16
	dmSpecVersion    uint16
	DriverVersion    uint16
	dmSize           uint16
	DriverExtra      uint16
	Fields           co.DM
	union0           DEVMODE_Printer // or DEVMODE_Display
	Color            co.DMCOLOR
	Duplex           co.DMDUP
	YResolution      int16
	TTOption         co.DMTT
	Collate          co.DMCOLLATE
	formName         [utl.CCHFORMNAME]uint16
	LogPixels        uint16
	BitsPerPel       uint32
	PelsWidth        uint32
	PelsHeight       uint32
	union1           uint32 // co.DMDISPLAYFLAGS or co.DMNUP
	DisplayFrequency uint32
	ICMMethod        co.DMICMMETHOD
	ICMIntent        co.DMICM
	MediaType        co.DMMEDIA
	DitherType       co.DMDITHER
	reserved1        uint32
	reserved2        uint32
	PanningWidth     uint32
	PanningHeight    uint32
}

// 1st variation of 1st union of [DEVMODE] struct.
type DEVMODE_Printer struct {
	Orientation   co.DMORIENT
	PaperSize     co.DMPAPER
	PaperLength   int16
	PaperWidth    int16
	Scale         int16
	Copies        int16
	DefaultSource co.DMBIN
	PrintQuality  co.DMRES
}

// 2st variation of 1st union of [DEVMODE] struct.
type DEVMODE_Display struct {
	Position           POINT
	DisplayOrientation co.DMDO
	DisplayFixedOutput co.DMDFO
}

func (dm *DEVMODE) DeviceName() string {
	return wstr.DecodeSlice(dm.deviceName[:])
}
func (dm *DEVMODE) SetDeviceName(val string) {
	wstr.EncodeToBuf(dm.deviceName[:], val)
}

// Sets the internal dmSize field to the size of the struct, correctly
// initializing it. Also sets dmSpecVersion.
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

func (dm *DEVMODE) FormName() string {
	return wstr.DecodeSlice(dm.formName[:])
}
func (dm *DEVMODE) SetFormName(val string) {
	wstr.EncodeToBuf(dm.formName[:], val)
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
	Style co.BRS
	Color COLORREF
	Hatch uintptr // ULONG_PTR
}

// [LOGFONT] struct.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         co.FW
	Italic         uint8 // This is a BOOL value.
	Underline      uint8 // This is a BOOL value.
	StrikeOut      uint8 // This is a BOOL value.
	CharSet        co.CHARSET
	OutPrecision   co.OUT_PRECIS
	ClipPrecision  co.CLIP_PRECIS
	Quality        co.QUALITY
	pitchAndFamily uint8 // combination of co.PITCH and co.FF
	faceName       [utl.LF_FACESIZE]uint16
}

func (lf *LOGFONT) FaceName() string {
	return wstr.DecodeSlice(lf.faceName[:])
}
func (lf *LOGFONT) SetFaceName(val string) {
	wstr.EncodeToBuf(lf.faceName[:], val)
}

func (lf *LOGFONT) Pitch() co.PITCH {
	return co.PITCH(lf.pitchAndFamily & 0b1111)
}
func (lf *LOGFONT) SetPitch(val co.PITCH) {
	lf.pitchAndFamily &^= 0b1111 // clear bits
	lf.pitchAndFamily |= uint8(val & 0b1111)
}

func (lf *LOGFONT) Family() co.FF {
	return co.FF(lf.pitchAndFamily & 0b1111_0000)
}
func (lf *LOGFONT) SetFamily(val co.FF) {
	lf.pitchAndFamily &^= 0b1111_0000 // clear bits
	lf.pitchAndFamily |= uint8(val & 0b1111_0000)
}

// [LOGPEN] struct.
//
// [LOGPEN]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logpen
type LOGPEN struct {
	Style co.PS
	Width POINT
	Color COLORREF
}

// [PALETTEENTRY] struct.
//
// [PALETTEENTRY]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-paletteentry
type PALETTEENTRY struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Flags co.PC
}

// [PIXELFORMATDESCRIPTOR] struct.
//
// ⚠️ You must call [PIXELFORMATDESCRIPTOR.SetNSize] to initialize the struct.
//
// Example:
//
//	var pfd win.PIXELFORMATDESCRIPTOR
//	pfd.SetNSize()
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

// Sets the internal nSize field to the size of the struct, correctly
// initializing it. Also sets nVersion.
func (mix *PIXELFORMATDESCRIPTOR) SetNSize() {
	mix.nSize = uint16(unsafe.Sizeof(*mix))
	mix.nVersion = 1
}

// [RGBQUAD] struct.
//
// Stores red, green and blue values in the range 0-255.
//
// [RGBQUAD]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-rgbquad
type RGBQUAD struct {
	data [4]uint8
}

func (rq *RGBQUAD) Blue() uint8 {
	return rq.data[0]
}
func (rq *RGBQUAD) SetBlue(val uint8) {
	rq.data[0] = val
}

func (rq *RGBQUAD) Green() uint8 {
	return rq.data[1]
}
func (rq *RGBQUAD) SetGreen(val uint8) {
	rq.data[1] = val
}

func (rq *RGBQUAD) Red() uint8 {
	return rq.data[2]
}
func (rq *RGBQUAD) SetRed(val uint8) {
	rq.data[2] = val
}

func (rq *RGBQUAD) ToColorref() COLORREF {
	return RGB(rq.Red(), rq.Green(), rq.Blue())
}

// [TEXTMETRIC] struct.
//
// [TEXTMETRIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type TEXTMETRIC struct {
	Height           uint32
	Ascent           uint32
	Descent          uint32
	InternalLeading  uint32
	ExternalLeading  uint32
	AveCharWidth     uint32
	MaxCharWidth     uint32
	Weight           uint32
	Overhang         uint32
	DigitizedAspectX uint32
	DigitizedAspectY uint32
	FirstChar        uint16
	LastChar         uint16
	DefaultChar      uint16
	BreakChar        uint16
	Italic           uint8
	Underlined       uint8
	StruckOut        uint8
	pitchAndFamily   uint8 // combination of co.TMPF and co.FF
	CharSet          co.CHARSET
}

func (tm *TEXTMETRIC) Pitch() co.TMPF {
	return co.TMPF(tm.pitchAndFamily & 0b1111)
}
func (tm *TEXTMETRIC) SetPitch(val co.TMPF) {
	tm.pitchAndFamily &^= 0b1111 // clear bits
	tm.pitchAndFamily |= uint8(val & 0b1111)
}

func (tm *TEXTMETRIC) Family() co.FF {
	return co.FF(tm.pitchAndFamily & 0b1111_0000)
}
func (tm *TEXTMETRIC) SetFamily(val co.FF) {
	tm.pitchAndFamily &^= 0b1111_0000 // clear bits
	tm.pitchAndFamily |= uint8(val & 0b1111_0000)
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
