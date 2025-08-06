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
	BRS_HOLLOW            = BRS_NULL
	BRS_HATCHED       BRS = 2
	BRS_PATTERN       BRS = 3
	BRS_INDEXED       BRS = 4
	BRS_DIBPATTERN    BRS = 5
	BRS_DIBPATTERNPT  BRS = 6
	BRS_PATTERN8X8    BRS = 7
	BRS_DIBPATTERN8X8 BRS = 8
	BRS_MONOPATTERN   BRS = 9
)

// [TEXTMETRIC] tmCharSet and [LOGFONT] lfCharSet. Originally with _CHARSET
// suffix.
//
// [TEXTMETRIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
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

// [GetDCEx] flags.
//
// [GetDCEx]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getdc
type DCX uint32

const (
	DCX_WINDOW           DCX = 0x0000_0001
	DCX_CACHE            DCX = 0x0000_0002
	DCX_NORESETATTRS     DCX = 0x0000_0004
	DCX_CLIPCHILDREN     DCX = 0x0000_0008
	DCX_CLIPSIBLINGS     DCX = 0x0000_0010
	DCX_PARENTCLIP       DCX = 0x0000_0020
	DCX_EXCLUDERGN       DCX = 0x0000_0040
	DCX_INTERSECTRGN     DCX = 0x0000_0080
	DCX_EXCLUDEUPDATE    DCX = 0x0000_0100
	DCX_LOCKWINDOWUPDATE DCX = 0x0000_0400
)

// [CreateDIBSection], [GetDIBits] and [SetDIBitsToDevice] usage. Originally has
// DIV prefix and COLORS suffix.
//
// [CreateDIBSection]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-createdibsection
// [GetDIBits]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getdibits
// [SetDIBitsToDevice]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setdibitstodevice
type DIB_COLORS uint32

const (
	DIB_COLORS_RGB DIB_COLORS = 0 // Color table in RGBs.
	DIB_COLORS_PAL DIB_COLORS = 1 // Color table in palette indices.
)

// [LOGFONT] lfClipPrecision. Originally with CLIP prefix and PRECIS suffix.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type CLIP_PRECIS uint8

const (
	CLIP_PRECIS_DEFAULT     CLIP_PRECIS = 0
	CLIP_PRECIS_CHARACTER   CLIP_PRECIS = 1
	CLIP_PRECIS_STROKE      CLIP_PRECIS = 2
	CLIP_PRECIS_MASK        CLIP_PRECIS = 0xf
	CLIP_PRECIS_LH_ANGLES   CLIP_PRECIS = 1 << 4
	CLIP_PRECIS_TT_ALWAYS   CLIP_PRECIS = 2 << 4
	CLIP_PRECIS_DFA_DISABLE CLIP_PRECIS = 4 << 4
	CLIP_PRECIS_EMBEDDED    CLIP_PRECIS = 8 << 4
)

// [DOCINFO] fwType. Originally has DI prefix.
//
// [DOCINFO]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-docinfow
type DIPJ uint32

const (
	DIPJ_NONE                  DIPJ = 0
	DIPJ_APPBANDING            DIPJ = 0x0000_0001
	DIPJ_ROPS_READ_DESTINATION DIPJ = 0x0000_0002
)

// [DEVMODE] dmFields.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DM uint32

const (
	DM_ORIENTATION        DM = 0x0000_0001
	DM_PAPERSIZE          DM = 0x0000_0002
	DM_PAPERLENGTH        DM = 0x0000_0004
	DM_PAPERWIDTH         DM = 0x0000_0008
	DM_SCALE              DM = 0x0000_0010
	DM_POSITION           DM = 0x0000_0020
	DM_NUP                DM = 0x0000_0040
	DM_DISPLAYORIENTATION DM = 0x0000_0080
	DM_COPIES             DM = 0x0000_0100
	DM_DEFAULTSOURCE      DM = 0x0000_0200
	DM_PRINTQUALITY       DM = 0x0000_0400
	DM_COLOR              DM = 0x0000_0800
	DM_DUPLEX             DM = 0x0000_1000
	DM_YRESOLUTION        DM = 0x0000_2000
	DM_TTOPTION           DM = 0x0000_4000
	DM_COLLATE            DM = 0x0000_8000
	DM_FORMNAME           DM = 0x0001_0000
	DM_LOGPIXELS          DM = 0x0002_0000
	DM_BITSPERPEL         DM = 0x0004_0000
	DM_PELSWIDTH          DM = 0x0008_0000
	DM_PELSHEIGHT         DM = 0x0010_0000
	DM_DISPLAYFLAGS       DM = 0x0020_0000
	DM_DISPLAYFREQUENCY   DM = 0x0040_0000
	DM_ICMMETHOD          DM = 0x0080_0000
	DM_ICMINTENT          DM = 0x0100_0000
	DM_MEDIATYPE          DM = 0x0200_0000
	DM_DITHERTYPE         DM = 0x0400_0000
	DM_PANNINGWIDTH       DM = 0x0800_0000
	DM_PANNINGHEIGHT      DM = 0x1000_0000
	DM_DISPLAYFIXEDOUTPUT DM = 0x2000_0000
)

// [DEVMODE] dmDefaultSource.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMBIN int16

const (
	DMBIN_UPPER         DMBIN = 1
	DMBIN_ONLYONE       DMBIN = 1
	DMBIN_LOWER         DMBIN = 2
	DMBIN_MIDDLE        DMBIN = 3
	DMBIN_MANUAL        DMBIN = 4
	DMBIN_ENVELOPE      DMBIN = 5
	DMBIN_ENVMANUAL     DMBIN = 6
	DMBIN_AUTO          DMBIN = 7
	DMBIN_TRACTOR       DMBIN = 8
	DMBIN_SMALLFMT      DMBIN = 9
	DMBIN_LARGEFMT      DMBIN = 10
	DMBIN_LARGECAPACITY DMBIN = 11
	DMBIN_CASSETTE      DMBIN = 14
	DMBIN_FORMSOURCE    DMBIN = 15
	DMBIN_LAST                = DMBIN_FORMSOURCE
	DMBIN_USER          DMBIN = 256
)

// [DEVMODE] dmCollate.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMCOLLATE int16

const (
	DMCOLLATE_FALSE DMCOLLATE = 0
	DMCOLLATE_TRUE  DMCOLLATE = 1
)

// [DEVMODE] dmColor.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMCOLOR int16

const (
	DMCOLOR_MONOCHROME DMCOLOR = 1
	DMCOLOR_COLOR      DMCOLOR = 2
)

// [DEVMODE] dmDisplayFixedOutput.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMDFO uint32

const (
	DMDFO_DEFAULT DMDFO = 0
	DMDFO_STRETCH DMDFO = 1
	DMDFO_CENTER  DMDFO = 2
)

// [DEVMODE] dmDisplayFlags.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMDISPLAYFLAGS uint32

const (
	DMDISPLAYFLAGS_INTERLACED DMDISPLAYFLAGS = 0x0000_0002
	DMDISPLAYFLAGS_TEXTMODE   DMDISPLAYFLAGS = 0x0000_0004
)

// [DEVMODE] dmDitherType.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMDITHER uint32

const (
	DMDITHER_NONE           DMDITHER = 1
	DMDITHER_COARSE         DMDITHER = 2
	DMDITHER_FINE           DMDITHER = 3
	DMDITHER_LINEART        DMDITHER = 4
	DMDITHER_ERRORDIFFUSION DMDITHER = 5
	DMDITHER_RESERVED6      DMDITHER = 6
	DMDITHER_RESERVED7      DMDITHER = 7
	DMDITHER_RESERVED8      DMDITHER = 8
	DMDITHER_RESERVED9      DMDITHER = 9
	DMDITHER_GRAYSCALE      DMDITHER = 10
	DMDITHER_USER           DMDITHER = 256
)

// [DEVMODE] dmDisplayOrientation.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMDO uint32

const (
	DMDO_D90  DMDO = 1
	DMDO_D180 DMDO = 2
	DMDO_D270 DMDO = 3
)

// [DEVMODE] dmDuplex.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMDUP int16

const (
	DMDUP_SIMPLEX    DMDUP = 1
	DMDUP_VERTICAL   DMDUP = 2
	DMDUP_HORIZONTAL DMDUP = 3
)

// [DEVMODE] dmICMIntent.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMICM uint32

const (
	DMICM_SATURATE         DMICM = 1
	DMICM_CONTRAST         DMICM = 2
	DMICM_COLORIMETRIC     DMICM = 3
	DMICM_ABS_COLORIMETRIC DMICM = 4
	DMICM_USER             DMICM = 256
)

// [DEVMODE] dmICMMethod.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMICMMETHOD uint32

const (
	DMICMMETHOD_NONE   DMICMMETHOD = 1
	DMICMMETHOD_SYSTEM DMICMMETHOD = 2
	DMICMMETHOD_DRIVER DMICMMETHOD = 3
	DMICMMETHOD_DEVICE DMICMMETHOD = 4
	DMICMMETHOD_USER   DMICMMETHOD = 256
)

// [DEVMODE] dmMediaType.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMMEDIA uint32

const (
	DMMEDIA_STANDARD     DMMEDIA = 1
	DMMEDIA_TRANSPARENCY DMMEDIA = 2
	DMMEDIA_GLOSSY       DMMEDIA = 3
	DMMEDIA_USER         DMMEDIA = 256
)

// [DEVMODE] dmNup.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMNUP uint32

const (
	DMNUP_SYSTEM DMNUP = 1
	DMNUP_ONEUP  DMNUP = 2
)

// [DEVMODE] dmOrientation.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMORIENT int16

const (
	DMORIENT_PORTRAIT  DMORIENT = 1
	DMORIENT_LANDSCAPE DMORIENT = 2
)

// [DEVMODE] dmPaperSize.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMPAPER int16

const (
	DMPAPER_LETTER                        DMPAPER = 1   // Letter 8 1/2 x 11 in.
	DMPAPER_LETTERSMALL                   DMPAPER = 2   // Letter Small 8 1/2 x 11 in.
	DMPAPER_TABLOID                       DMPAPER = 3   // Tabloid 11 x 17 in.
	DMPAPER_LEDGER                        DMPAPER = 4   // Ledger 17 x 11 in.
	DMPAPER_LEGAL                         DMPAPER = 5   // Legal 8 1/2 x 14 in.
	DMPAPER_STATEMENT                     DMPAPER = 6   // Statement 5 1/2 x 8 1/2 in.
	DMPAPER_EXECUTIVE                     DMPAPER = 7   // Executive 7 1/4 x 10 1/2 in.
	DMPAPER_A3                            DMPAPER = 8   // A3 297 x 420 mm.
	DMPAPER_A4                            DMPAPER = 9   // A4 210 x 297 mm.
	DMPAPER_A4SMALL                       DMPAPER = 10  // A4 Small 210 x 297 mm.
	DMPAPER_A5                            DMPAPER = 11  // A5 148 x 210 mm.
	DMPAPER_B4                            DMPAPER = 12  // B4 (JIS) 250 x 354.
	DMPAPER_B5                            DMPAPER = 13  // B5 (JIS) 182 x 257 mm.
	DMPAPER_FOLIO                         DMPAPER = 14  // Folio 8 1/2 x 13 in.
	DMPAPER_QUARTO                        DMPAPER = 15  // Quarto 215 x 275 mm.
	DMPAPER_P10X14                        DMPAPER = 16  // 10x14 in.
	DMPAPER_P11X17                        DMPAPER = 17  // 11x17 in.
	DMPAPER_NOTE                          DMPAPER = 18  // Note 8 1/2 x 11 in.
	DMPAPER_ENV_9                         DMPAPER = 19  // Envelope #9 3 7/8 x 8 7/8.
	DMPAPER_ENV_10                        DMPAPER = 20  // Envelope #10 4 1/8 x 9 1/2.
	DMPAPER_ENV_11                        DMPAPER = 21  // Envelope #11 4 1/2 x 10 3/8.
	DMPAPER_ENV_12                        DMPAPER = 22  // Envelope #12 4 \276 x 11.
	DMPAPER_ENV_14                        DMPAPER = 23  // Envelope #14 5 x 11 1/2.
	DMPAPER_CSHEET                        DMPAPER = 24  // C size sheet.
	DMPAPER_DSHEET                        DMPAPER = 25  // D size sheet.
	DMPAPER_ESHEET                        DMPAPER = 26  // E size sheet.
	DMPAPER_ENV_DL                        DMPAPER = 27  // Envelope DL 110 x 220mm.
	DMPAPER_ENV_C5                        DMPAPER = 28  // Envelope C5 162 x 229 mm.
	DMPAPER_ENV_C3                        DMPAPER = 29  // Envelope C3 324 x 458 mm.
	DMPAPER_ENV_C4                        DMPAPER = 30  // Envelope C4 229 x 324 mm.
	DMPAPER_ENV_C6                        DMPAPER = 31  // Envelope C6 114 x 162 mm.
	DMPAPER_ENV_C65                       DMPAPER = 32  // Envelope C65 114 x 229 mm.
	DMPAPER_ENV_B4                        DMPAPER = 33  // Envelope B4 250 x 353 mm.
	DMPAPER_ENV_B5                        DMPAPER = 34  // Envelope B5 176 x 250 mm.
	DMPAPER_ENV_B6                        DMPAPER = 35  // Envelope B6 176 x 125 mm.
	DMPAPER_ENV_ITALY                     DMPAPER = 36  // Envelope 110 x 230 mm.
	DMPAPER_ENV_MONARCH                   DMPAPER = 37  // Envelope Monarch 3.875 x 7.5 in.
	DMPAPER_ENV_PERSONAL                  DMPAPER = 38  // 6 3/4 Envelope 3 5/8 x 6 1/2 in.
	DMPAPER_FANFOLD_US                    DMPAPER = 39  // US Std Fanfold 14 7/8 x 11 in.
	DMPAPER_FANFOLD_STD_GERMAN            DMPAPER = 40  // German Std Fanfold 8 1/2 x 12 in.
	DMPAPER_FANFOLD_LGL_GERMAN            DMPAPER = 41  // German Legal Fanfold 8 1/2 x 13 in.
	DMPAPER_ISO_B4                        DMPAPER = 42  // B4 (ISO) 250 x 353 mm.
	DMPAPER_JAPANESE_POSTCARD             DMPAPER = 43  // Japanese Postcard 100 x 148 mm.
	DMPAPER_P9X11                         DMPAPER = 44  // 9 x 11 in.
	DMPAPER_P10X11                        DMPAPER = 45  // 10 x 11 in.
	DMPAPER_P15X11                        DMPAPER = 46  // 15 x 11 in.
	DMPAPER_ENV_INVITE                    DMPAPER = 47  // Envelope Invite 220 x 220 mm.
	DMPAPER_LETTER_EXTRA                  DMPAPER = 50  // Letter Extra 9 275 x 12 in.
	DMPAPER_LEGAL_EXTRA                   DMPAPER = 51  // Legal Extra 9 275 x 15 in.
	DMPAPER_TABLOID_EXTRA                 DMPAPER = 52  // Tabloid Extra 11.69 x 18 in.
	DMPAPER_A4_EXTRA                      DMPAPER = 53  // A4 Extra 9.27 x 12.69 in.
	DMPAPER_LETTER_TRANSVERSE             DMPAPER = 54  // Letter Transverse 8 275 x 11 in.
	DMPAPER_A4_TRANSVERSE                 DMPAPER = 55  // A4 Transverse 210 x 297 mm.
	DMPAPER_LETTER_EXTRA_TRANSVERSE       DMPAPER = 56  // Letter Extra Transverse 9\275 x 12 in.
	DMPAPER_A_PLUS                        DMPAPER = 57  // SuperA/SuperA/A4 227 x 356 mm.
	DMPAPER_B_PLUS                        DMPAPER = 58  // SuperB/SuperB/A3 305 x 487 mm.
	DMPAPER_LETTER_PLUS                   DMPAPER = 59  // Letter Plus 8.5 x 12.69 in.
	DMPAPER_A4_PLUS                       DMPAPER = 60  // A4 Plus 210 x 330 mm.
	DMPAPER_A5_TRANSVERSE                 DMPAPER = 61  // A5 Transverse 148 x 210 mm.
	DMPAPER_B5_TRANSVERSE                 DMPAPER = 62  // B5 (JIS) Transverse 182 x 257 mm.
	DMPAPER_A3_EXTRA                      DMPAPER = 63  // A3 Extra 322 x 445 mm.
	DMPAPER_A5_EXTRA                      DMPAPER = 64  // A5 Extra 174 x 235 mm.
	DMPAPER_B5_EXTRA                      DMPAPER = 65  // B5 (ISO) Extra 201 x 276 mm.
	DMPAPER_A2                            DMPAPER = 66  // A2 420 x 594 mm.
	DMPAPER_A3_TRANSVERSE                 DMPAPER = 67  // A3 Transverse 297 x 420 mm.
	DMPAPER_A3_EXTRA_TRANSVERSE           DMPAPER = 68  // A3 Extra Transverse 322 x 445 mm.
	DMPAPER_DBL_JAPANESE_POSTCARD         DMPAPER = 69  // Japanese Double Postcard 200 x 148 mm.
	DMPAPER_A6                            DMPAPER = 70  // A6 105 x 148 mm.
	DMPAPER_JENV_KAKU2                    DMPAPER = 71  // Japanese Envelope Kaku #2.
	DMPAPER_JENV_KAKU3                    DMPAPER = 72  // Japanese Envelope Kaku #3.
	DMPAPER_JENV_CHOU3                    DMPAPER = 73  // Japanese Envelope Chou #3.
	DMPAPER_JENV_CHOU4                    DMPAPER = 74  // Japanese Envelope Chou #4.
	DMPAPER_LETTER_ROTATED                DMPAPER = 75  // Letter Rotated 11 x 8 1/2 11 in.
	DMPAPER_A3_ROTATED                    DMPAPER = 76  // A3 Rotated 420 x 297 mm.
	DMPAPER_A4_ROTATED                    DMPAPER = 77  // A4 Rotated 297 x 210 mm.
	DMPAPER_A5_ROTATED                    DMPAPER = 78  // A5 Rotated 210 x 148 mm.
	DMPAPER_B4_JIS_ROTATED                DMPAPER = 79  // B4 (JIS) Rotated 364 x 257 mm.
	DMPAPER_B5_JIS_ROTATED                DMPAPER = 80  // B5 (JIS) Rotated 257 x 182 mm.
	DMPAPER_JAPANESE_POSTCARD_ROTATED     DMPAPER = 81  // Japanese Postcard Rotated 148 x 100 mm.
	DMPAPER_DBL_JAPANESE_POSTCARD_ROTATED DMPAPER = 82  // Double Japanese Postcard Rotated 148 x 200 mm.
	DMPAPER_A6_ROTATED                    DMPAPER = 83  // A6 Rotated 148 x 105 mm.
	DMPAPER_JENV_KAKU2_ROTATED            DMPAPER = 84  // Japanese Envelope Kaku #2 Rotated.
	DMPAPER_JENV_KAKU3_ROTATED            DMPAPER = 85  // Japanese Envelope Kaku #3 Rotated.
	DMPAPER_JENV_CHOU3_ROTATED            DMPAPER = 86  // Japanese Envelope Chou #3 Rotated.
	DMPAPER_JENV_CHOU4_ROTATED            DMPAPER = 87  // Japanese Envelope Chou #4 Rotated.
	DMPAPER_B6_JIS                        DMPAPER = 88  // B6 (JIS) 128 x 182 mm.
	DMPAPER_B6_JIS_ROTATED                DMPAPER = 89  // B6 (JIS) Rotated 182 x 128 mm.
	DMPAPER_P12X11                        DMPAPER = 90  // 12 x 11 in.
	DMPAPER_JENV_YOU4                     DMPAPER = 91  // Japanese Envelope You #4.
	DMPAPER_JENV_YOU4_ROTATED             DMPAPER = 92  // Japanese Envelope You #4 Rotated.
	DMPAPER_P16K                          DMPAPER = 93  // PRC 16K 146 x 215 mm.
	DMPAPER_P32K                          DMPAPER = 94  // PRC 32K 97 x 151 mm.
	DMPAPER_P32KBIG                       DMPAPER = 95  // PRC 32K (Big) 97 x 151 mm.
	DMPAPER_PENV_1                        DMPAPER = 96  // PRC Envelope #1 102 x 165 mm.
	DMPAPER_PENV_2                        DMPAPER = 97  // PRC Envelope #2 102 x 176 mm.
	DMPAPER_PENV_3                        DMPAPER = 98  // PRC Envelope #3 125 x 176 mm.
	DMPAPER_PENV_4                        DMPAPER = 99  // PRC Envelope #4 110 x 208 mm.
	DMPAPER_PENV_5                        DMPAPER = 100 // PRC Envelope #5 110 x 220 mm.
	DMPAPER_PENV_6                        DMPAPER = 101 // PRC Envelope #6 120 x 230 mm.
	DMPAPER_PENV_7                        DMPAPER = 102 // PRC Envelope #7 160 x 230 mm.
	DMPAPER_PENV_8                        DMPAPER = 103 // PRC Envelope #8 120 x 309 mm.
	DMPAPER_PENV_9                        DMPAPER = 104 // PRC Envelope #9 229 x 324 mm.
	DMPAPER_PENV_10                       DMPAPER = 105 // PRC Envelope #10 324 x 458 mm.
	DMPAPER_P16K_ROTATED                  DMPAPER = 106 // PRC 16K Rotated.
	DMPAPER_P32K_ROTATED                  DMPAPER = 107 // PRC 32K Rotated.
	DMPAPER_P32KBIG_ROTATED               DMPAPER = 108 // PRC 32K(Big) Rotated.
	DMPAPER_PENV_1_ROTATED                DMPAPER = 109 // PRC Envelope #1 Rotated 165 x 102 mm.
	DMPAPER_PENV_2_ROTATED                DMPAPER = 110 // PRC Envelope #2 Rotated 176 x 102 mm.
	DMPAPER_PENV_3_ROTATED                DMPAPER = 111 // PRC Envelope #3 Rotated 176 x 125 mm.
	DMPAPER_PENV_4_ROTATED                DMPAPER = 112 // PRC Envelope #4 Rotated 208 x 110 mm.
	DMPAPER_PENV_5_ROTATED                DMPAPER = 113 // PRC Envelope #5 Rotated 220 x 110 mm.
	DMPAPER_PENV_6_ROTATED                DMPAPER = 114 // PRC Envelope #6 Rotated 230 x 120 mm.
	DMPAPER_PENV_7_ROTATED                DMPAPER = 115 // PRC Envelope #7 Rotated 230 x 160 mm.
	DMPAPER_PENV_8_ROTATED                DMPAPER = 116 // PRC Envelope #8 Rotated 309 x 120 mm.
	DMPAPER_PENV_9_ROTATED                DMPAPER = 117 // PRC Envelope #9 Rotated 324 x 229 mm.
	DMPAPER_PENV_10_ROTATED               DMPAPER = 118 // PRC Envelope #10 Rotated 458 x 324 mm.
	DMPAPER_USER                          DMPAPER = 256 // Other papers start here.
)

// [DEVMODE] dmPrintQuality.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMRES int16

const (
	DMRES_DRAFT  DMRES = -1
	DMRES_LOW    DMRES = -2
	DMRES_MEDIUM DMRES = -3
	DMRES_HIGH   DMRES = -4
)

// [DEVMODE] dmTTOption.
//
// [DEVMODE]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-devmodew
type DMTT int16

const (
	DMTT_BITMAP           DMTT = 1
	DMTT_DOWNLOAD         DMTT = 2
	DMTT_SUBDEV           DMTT = 3
	DMTT_DOWNLOAD_OUTLINE DMTT = 4
)

// [LOGFONT] family.
//
// The values set the bits 4 to 7. Bits 0 to 3 are usually set by [PITCH].
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type FF uint8

const (
	FF_DONTCARE   FF = 0 << 4
	FF_ROMAN      FF = 1 << 4
	FF_SWISS      FF = 2 << 4
	FF_MODERN     FF = 3 << 4
	FF_SCRIPT     FF = 4 << 4
	FF_DECORATIVE FF = 5 << 4
)

// [AddFontResourceEx] fl.
//
// [AddFontResourceEx]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-addfontresourceexw
type FR uint32

const (
	FR_PRIVATE  FR = 0x10
	FR_NOT_ENUM FR = 0x20
)

// [LOGFONT] lfWeight.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type FW uint32

const (
	FW_DONTCARE   FW = 0
	FW_THIN       FW = 100
	FW_EXTRALIGHT FW = 200
	FW_ULTRALIGHT    = FW_EXTRALIGHT
	FW_LIGHT      FW = 300
	FW_NORMAL     FW = 400
	FW_REGULAR    FW = 400
	FW_MEDIUM     FW = 500
	FW_SEMIBOLD   FW = 600
	FW_DEMIBOLD      = FW_SEMIBOLD
	FW_BOLD       FW = 700
	FW_EXTRABOLD  FW = 800
	FW_ULTRABOLD     = FW_EXTRABOLD
	FW_HEAVY      FW = 900
	FW_BLACK         = FW_HEAVY
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

// [GradientFill] mode.
//
// [GradientFill]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-gradientfill
type GRADIENT_FILL uint32

const (
	GRADIENT_FILL_RECT_H   GRADIENT_FILL = 0x0000_0000
	GRADIENT_FILL_RECT_V   GRADIENT_FILL = 0x0000_0001
	GRADIENT_FILL_TRIANGLE GRADIENT_FILL = 0x0000_0002
)

// [LOGFONT] lfOutPrecision. Originally with OUT prefix and PRECIS suffix.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type OUT_PRECIS uint8

const (
	OUT_PRECIS_DEFAULT        OUT_PRECIS = 0
	OUT_PRECIS_STRING         OUT_PRECIS = 1
	OUT_PRECIS_CHARACTER      OUT_PRECIS = 2
	OUT_PRECIS_STROKE         OUT_PRECIS = 3
	OUT_PRECIS_TT             OUT_PRECIS = 4
	OUT_PRECIS_DEVICE         OUT_PRECIS = 5
	OUT_PRECIS_RASTER         OUT_PRECIS = 6
	OUT_PRECIS_TT_ONLY        OUT_PRECIS = 7
	OUT_PRECIS_OUTLINE        OUT_PRECIS = 8
	OUT_PRECIS_SCREEN_OUTLINE OUT_PRECIS = 9
	OUT_PRECIS_PS_ONLY        OUT_PRECIS = 10
)

// [PALETTEENTRY] PeFlags.
//
// [PALETTEENTRY]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-paletteentry
type PC uint8

const (
	PC_RESERVED   PC = 0x01 // Palette index used for animation.
	PC_EXPLICIT   PC = 0x02 // Palette index is explicit to device.
	PC_NOCOLLAPSE PC = 0x04 // Do not match color to system palette.
)

// [PIXELFORMATDESCRIPTOR] dwFlags.
//
// [PIXELFORMATDESCRIPTOR]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-pixelformatdescriptor
type PFD uint32

const (
	PFD_DRAW_TO_WINDOW      PFD = 0x0000_0004
	PFD_DRAW_TO_BITMAP      PFD = 0x0000_0008
	PFD_SUPPORT_GDI         PFD = 0x0000_0010
	PFD_SUPPORT_OPENGL      PFD = 0x0000_0020
	PFD_GENERIC_ACCELERATED PFD = 0x0000_1000
	PFD_GENERIC_FORMAT      PFD = 0x0000_0040
	PFD_NEED_PALETTE        PFD = 0x0000_0080
	PFD_NEED_SYSTEM_PALETTE PFD = 0x0000_0100
	PFD_DOUBLEBUFFER        PFD = 0x0000_0001
	PFD_STEREO              PFD = 0x0000_0002
	PFD_SWAP_LAYER_BUFFERS  PFD = 0x0000_0800

	PFD_DEPTH_DONTCARE        PFD = 0x2000_0000
	PFD_DOUBLEBUFFER_DONTCARE PFD = 0x4000_0000
	PFD_STEREO_DONTCARE       PFD = 0x8000_0000
	PFD_SWAP_COPY             PFD = 0x0000_0400
	PFD_SWAP_EXCHANGE         PFD = 0x0000_0200
)

// [PIXELFORMATDESCRIPTOR] dwPixelType.
//
// [PIXELFORMATDESCRIPTOR]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-pixelformatdescriptor
type PFD_TYPE uint8

const (
	PFD_TYPE_RGBA       PFD_TYPE = 0
	PFD_TYPE_COLORINDEX PFD_TYPE = 1
)

// [LOGFONT] pitch. Originally has PITCH suffix.
//
// The values set the bits 0 to 3. Bits 4 to 7 are usually set by [FF].
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type PITCH uint8

const (
	PITCH_DEFAULT  PITCH = 0
	PITCH_FIXED    PITCH = 1
	PITCH_VARIABLE PITCH = 2
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

// [ExtCreatePen] end cap. Originally has PS prefix.
//
// [ExtCreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-extcreatepen
type PS_ENDCAP uint32

const (
	PS_ENDCAP_ROUND  PS_ENDCAP = 0x0000_0000
	PS_ENDCAP_SQUARE PS_ENDCAP = 0x0000_0100
	PS_ENDCAP_FLAT   PS_ENDCAP = 0x0000_0200
)

// [ExtCreatePen] style. Originally has PS prefix.
//
// [ExtCreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-extcreatepen
type PS_STYLE uint32

const (
	PS_STYLE_ALTERNATE   PS_STYLE = 8
	PS_STYLE_SOLID                = PS_STYLE(PS_SOLID)
	PS_STYLE_DASH                 = PS_STYLE(PS_DASH)
	PS_STYLE_DOT                  = PS_STYLE(PS_DOT)
	PS_STYLE_DASHDOT              = PS_STYLE(PS_DASHDOT)
	PS_STYLE_DASHDOTDOT           = PS_STYLE(PS_DASHDOTDOT)
	PS_STYLE_NULL                 = PS_STYLE(PS_NULL)
	PS_STYLE_USERSTYLE   PS_STYLE = 7
	PS_STYLE_INSIDEFRAME          = PS_STYLE(PS_INSIDEFRAME)
)

// [ExtCreatePen] type. Originally has PS prefix.
//
// [ExtCreatePen]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-extcreatepen
type PS_TYPE uint32

const (
	PS_TYPE_COSMETIC  PS_TYPE = 0x0000_0000
	PS_TYPE_GEOMETRIC PS_TYPE = 0x0001_0000
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

// [LOGFONT] lfQuality. Originally with QUALITY suffix.
//
// [LOGFONT]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
type QUALITY uint8

const (
	QUALITY_DEFAULT           QUALITY = 0
	QUALITY_DRAFT             QUALITY = 1
	QUALITY_PROOF             QUALITY = 2
	QUALITY_NONANTIALIASED    QUALITY = 3
	QUALITY_ANTIALIASED       QUALITY = 4
	QUALITY_CLEARTYPE         QUALITY = 5
	QUALITY_CLEARTYPE_NATURAL QUALITY = 6
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

// [GetStockObject] type. Originally has no prefix.
//
// [GetStockObject]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-getstockobject
type STOCK int32

const (
	STOCK_WHITE_BRUSH         STOCK = 0
	STOCK_LTGRAY_BRUSH        STOCK = 1
	STOCK_GRAY_BRUSH          STOCK = 2
	STOCK_DKGRAY_BRUSH        STOCK = 3
	STOCK_BLACK_BRUSH         STOCK = 4
	STOCK_NULL_BRUSH          STOCK = 5
	STOCK_HOLLOW_BRUSH              = STOCK_NULL_BRUSH
	STOCK_WHITE_PEN           STOCK = 6
	STOCK_BLACK_PEN           STOCK = 7
	STOCK_NULL_PEN            STOCK = 8
	STOCK_OEM_FIXED_FONT      STOCK = 10
	STOCK_ANSI_FIXED_FONT     STOCK = 11
	STOCK_ANSI_VAR_FONT       STOCK = 12
	STOCK_SYSTEM_FONT         STOCK = 13
	STOCK_DEVICE_DEFAULT_FONT STOCK = 14
	STOCK_DEFAULT_PALETTE     STOCK = 15
	STOCK_SYSTEM_FIXED_FONT   STOCK = 16
	STOCK_DEFAULT_GUI_FONT    STOCK = 17
	STOCK_DC_BRUSH            STOCK = 18
	STOCK_DC_PEN              STOCK = 19
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
	STRETCH_ANDSCANS             = STRETCH_BLACKONWHITE
	STRETCH_ORSCANS              = STRETCH_WHITEONBLACK
	STRETCH_DELETESCANS          = STRETCH_COLORONCOLOR
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

// [TEXTMETRIC] tmPitchAndFamily.
//
// The values set the bits 0 to 3. Bits 4 to 7 are usually set by [FF].
//
// [TEXTMETRIC]: https://learn.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-textmetricw
type TMPF uint8

const (
	TMPF_FIXED_PITCH TMPF = 0x01
	TMPF_VECTOR      TMPF = 0x02
	TMPF_DEVICE      TMPF = 0x08
	TMPF_TRUETYPE    TMPF = 0x04
)
