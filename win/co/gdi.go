//go:build windows

package co

// [SetArcDirection] dir.
//
// [SetArcDirection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setarcdirection
type AD int32

const (
	AD_COUNTERCLOCKWISE AD = 1
	AD_CLOCKWISE        AD = 2
)

// [BITMAPINFOHEADER] biCompression.
//
// [BITMAPINFOHEADER]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BI uint32

const (
	BI_RGB       BI = 0
	BI_RLE8      BI = 1
	BI_RLE4      BI = 2
	BI_BITFIELDS BI = 3
	BI_JPEG      BI = 4
	BI_PNG       BI = 5
)

// [SetBkMode] mode. Originally has no prefix.
//
// [SetBkMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
type BKMODE int32

const (
	BKMODE_TRANSPARENT BKMODE = 1
	BKMODE_OPAQUE      BKMODE = 2
)

// [LOGBRUSH] lbStyle. Originally with BS prefix.
//
// [LOGBRUSH]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logbrush
type BRS uint32

const (
	BRS_SOLID         BRS = 0
	BRS_NULL          BRS = 1
	BRS_HOLLOW        BRS = BRS_NULL
	BRS_HATCHED       BRS = 2
	BRS_PATTERN       BRS = 3
	BRS_INDEXED       BRS = 4
	BRS_DIBPATTERN    BRS = 5
	BRS_DIBPATTERNPT  BRS = 6
	BRS_PATTERN8X8    BRS = 7
	BRS_DIBPATTERN8X8 BRS = 8
	BRS_MONOPATTERN   BRS = 9
)

// [TEXTMETRIC] tmCharSet. Originally with _CHARSET suffix.
//
// [TEXTMETRIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type CHARSET uint8

const (
	CHARSET_ANSI        CHARSET = 0
	CHARSET_DEFAULT     CHARSET = 1
	CHARSET_SYMBOL      CHARSET = 2
	CHARSET_SHIFTJIS    CHARSET = 128
	CHARSET_HANGUL      CHARSET = 129
	CHARSET_GB2312      CHARSET = 134
	CHARSET_CHINESEBIG5 CHARSET = 136
	CHARSET_OEM         CHARSET = 255
	CHARSET_JOHAB       CHARSET = 130
	CHARSET_HEBREW      CHARSET = 177
	CHARSET_ARABIC      CHARSET = 178
	CHARSET_GREEK       CHARSET = 161
	CHARSET_TURKISH     CHARSET = 162
	CHARSET_VIETNAMESE  CHARSET = 163
	CHARSET_THAI        CHARSET = 222
	CHARSET_EASTEUROPE  CHARSET = 238
	CHARSET_RUSSIAN     CHARSET = 204
	CHARSET_MAC         CHARSET = 77
	CHARSET_BALTIC      CHARSET = 186
)

// [CreateDIBSection] usage.
//
// [CreateDIBSection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
type DIB uint32

const (
	DIB_RGB_COLORS DIB = 0 // Color table in RGBs.
	DIB_PAL_COLORS DIB = 1 // Color table in palette indices.
)

// [LOGFONT] lfWeight.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type FW uint32

const (
	FW_DONTCARE   FW = 0
	FW_THIN       FW = 100
	FW_EXTRALIGHT FW = 200
	FW_ULTRALIGHT FW = FW_EXTRALIGHT
	FW_LIGHT      FW = 300
	FW_NORMAL     FW = 400
	FW_REGULAR    FW = 400
	FW_MEDIUM     FW = 500
	FW_SEMIBOLD   FW = 600
	FW_DEMIBOLD   FW = FW_SEMIBOLD
	FW_BOLD       FW = 700
	FW_EXTRABOLD  FW = 800
	FW_ULTRABOLD  FW = FW_EXTRABOLD
	FW_HEAVY      FW = 900
	FW_BLACK      FW = FW_HEAVY
)

// [GetDeviceCaps] index. Originally has no prefix.
//
// [GetDeviceCaps]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdevicecaps
type GDC int32

const (
	GDC_DRIVERVERSION   GDC = 0
	GDC_TECHNOLOGY      GDC = 2
	GDC_HORZSIZE        GDC = 4
	GDC_VERTSIZE        GDC = 6
	GDC_HORZRES         GDC = 8
	GDC_VERTRES         GDC = 10
	GDC_BITSPIXEL       GDC = 12
	GDC_PLANES          GDC = 14
	GDC_NUMBRUSHES      GDC = 16
	GDC_NUMPENS         GDC = 18
	GDC_NUMMARKERS      GDC = 20
	GDC_NUMFONTS        GDC = 22
	GDC_NUMCOLORS       GDC = 24
	GDC_PDEVICESIZE     GDC = 26
	GDC_CURVECAPS       GDC = 28
	GDC_LINECAPS        GDC = 30
	GDC_POLYGONALCAPS   GDC = 32
	GDC_TEXTCAPS        GDC = 34
	GDC_CLIPCAPS        GDC = 36
	GDC_RASTERCAPS      GDC = 38
	GDC_ASPECTX         GDC = 40
	GDC_ASPECTY         GDC = 42
	GDC_ASPECTXY        GDC = 44
	GDC_LOGPIXELSX      GDC = 88
	GDC_LOGPIXELSY      GDC = 90
	GDC_SIZEPALETTE     GDC = 104
	GDC_NUMRESERVED     GDC = 106
	GDC_COLORRES        GDC = 108
	GDC_PHYSICALWIDTH   GDC = 110
	GDC_PHYSICALHEIGHT  GDC = 111
	GDC_PHYSICALOFFSETX GDC = 112
	GDC_PHYSICALOFFSETY GDC = 113
	GDC_SCALINGFACTORX  GDC = 114
	GDC_SCALINGFACTORY  GDC = 115
	GDC_VREFRESH        GDC = 116
	GDC_DESKTOPVERTRES  GDC = 117
	GDC_DESKTOPHORZRES  GDC = 118
	GDC_BLTALIGNMENT    GDC = 119
	GDC_SHADEBLENDCAPS  GDC = 120
	GDC_COLORMGMTCAPS   GDC = 121
)

// [CreateHatchBrush] iHatch.
//
// [CreateHatchBrush]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createhatchbrush
type HS int32

const (
	HS_HORIZONTAL HS = 0 // Pattern: -----
	HS_VERTICAL   HS = 1 // Pattern: |||||
	HS_FDIAGONAL  HS = 2 // Pattern: \\\\\
	HS_BDIAGONAL  HS = 3 // Pattern: /////
	HS_CROSS      HS = 4 // Pattern: +++++
	HS_DIAGCROSS  HS = 5 // Pattern: xxxxx
)

// [SetPolyFillMode] mode. Originally has no prefix.
//
// [SetPolyFillMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setpolyfillmode
type POLYF int32

const (
	POLYF_ALTERNATE POLYF = 1
	POLYF_WINDING   POLYF = 2
)

// [WM_PRINT] drawing options.
//
// [WM_PRINT]: https://learn.microsoft.com/en-us/windows/win32/gdi/wm-print
type PRF uint32

const (
	PRF_CHECKVISIBLE PRF = 0x0000_0001
	PRF_NONCLIENT    PRF = 0x0000_0002
	PRF_CLIENT       PRF = 0x0000_0004
	PRF_ERASEBKGND   PRF = 0x0000_0008
	PRF_CHILDREN     PRF = 0x0000_0010
	PRF_OWNED        PRF = 0x0000_0020
)

// [CreatePen] iStyle.
//
// [CreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createpen
type PS int32

const (
	PS_SOLID       PS = 0
	PS_DASH        PS = 1
	PS_DOT         PS = 2
	PS_DASHDOT     PS = 3
	PS_DASHDOTDOT  PS = 4
	PS_NULL        PS = 5
	PS_INSIDEFRAME PS = 6
)

// [PolyDraw] aj.
//
// [PolyDraw]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-polydraw
type PT uint8

const (
	PT_CLOSEFIGURE PT = 0x01
	PT_LINETO      PT = 0x02
	PT_BEZIERTO    PT = 0x04
	PT_MOVETO      PT = 0x06
)

// [SelectObject] return value. Originally with REGION suffix.
//
// [SelectObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
type REGION uint32

const (
	REGION_NULL    REGION = 1
	REGION_SIMPLE  REGION = 2
	REGION_COMPLEX REGION = 3
)

// [CombineRgn] and [SelectClipPath] mode.
//
// [CombineRgn]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-combinergn
// [SelectClipPath]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectclippath
type RGN int32

const (
	RGN_AND  RGN = 1
	RGN_OR   RGN = 2
	RGN_XOR  RGN = 3
	RGN_DIFF RGN = 4
	RGN_COPY RGN = 5
)

// [BitBlt] rop, [IMAGELISTDRAWPARAMS] dwRop.
//
// [BitBlt]: https://learn.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
// [IMAGELISTDRAWPARAMS]: https://learn.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type ROP uint32

const (
	ROP_SRCCOPY        ROP = 0x00cc_0020
	ROP_SRCPAINT       ROP = 0x00ee_0086
	ROP_SRCAND         ROP = 0x0088_00c6
	ROP_SRCINVERT      ROP = 0x0066_0046
	ROP_SRCERASE       ROP = 0x0044_0328
	ROP_NOTSRCCOPY     ROP = 0x0033_0008
	ROP_NOTSRCERASE    ROP = 0x0011_00a6
	ROP_MERGECOPY      ROP = 0x00c0_00ca
	ROP_MERGEPAINT     ROP = 0x00bb_0226
	ROP_PATCOPY        ROP = 0x00f0_0021
	ROP_PATPAINT       ROP = 0x00fb_0a09
	ROP_PATINVERT      ROP = 0x005a_0049
	ROP_DSTINVERT      ROP = 0x0055_0009
	ROP_BLACKNESS      ROP = 0x0000_0042
	ROP_WHITENESS      ROP = 0x00ff_0062
	ROP_NOMIRRORBITMAP ROP = 0x8000_0000
	ROP_CAPTUREBLT     ROP = 0x4000_0000
)

// [SetStretchBltMode] mode.
//
// [SetStretchBltMode]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setstretchbltmode
type STRETCH int32

const (
	STRETCH_BLACKONWHITE STRETCH = 1
	STRETCH_WHITEONBLACK STRETCH = 2
	STRETCH_COLORONCOLOR STRETCH = 3
	STRETCH_HALFTONE     STRETCH = 4
	STRETCH_ANDSCANS     STRETCH = STRETCH_BLACKONWHITE
	STRETCH_ORSCANS      STRETCH = STRETCH_WHITEONBLACK
	STRETCH_DELETESCANS  STRETCH = STRETCH_COLORONCOLOR
)

// [SetTextAlign] align. Includes values with VTA prefix.
//
// [SetTextAlign]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-settextalign
type TA uint32

const (
	TA_NOUPDATECP TA = 0
	TA_UPDATECP   TA = 1
	TA_LEFT       TA = 0
	TA_RIGHT      TA = 2
	TA_CENTER     TA = 6
	TA_TOP        TA = 0
	TA_BOTTOM     TA = 8
	TA_BASELINE   TA = 24
	TA_RTLREADING TA = 256
)

// [TrackPopupMenu] uFlags.
//
// [TrackPopupMenu]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-trackpopupmenu
type TPM uint32

const (
	TPM_LEFTBUTTON      TPM = 0x0000
	TPM_RIGHTBUTTON     TPM = 0x0002
	TPM_LEFTALIGN       TPM = 0x0000
	TPM_CENTERALIGN     TPM = 0x0004
	TPM_RIGHTALIGN      TPM = 0x0008
	TPM_TOPALIGN        TPM = 0x0000
	TPM_VCENTERALIGN    TPM = 0x0010
	TPM_BOTTOMALIGN     TPM = 0x0020
	TPM_HORIZONTAL      TPM = 0x0000
	TPM_VERTICAL        TPM = 0x0040
	TPM_NONOTIFY        TPM = 0x0080
	TPM_RETURNCMD       TPM = 0x0100
	TPM_RECURSE         TPM = 0x0001
	TPM_HORPOSANIMATION TPM = 0x0400
	TPM_HORNEGANIMATION TPM = 0x0800
	TPM_VERPOSANIMATION TPM = 0x1000
	TPM_VERNEGANIMATION TPM = 0x2000
	TPM_NOANIMATION     TPM = 0x4000
	TPM_LAYOUTRTL       TPM = 0x8000
	TPM_WORKAREA        TPM = 0x1_0000
)
