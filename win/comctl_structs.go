//go:build windows

package win

import (
	"encoding/binary"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// A handle to a tree view control [item].
//
// [item]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-controls#parent-and-child-items
type HTREEITEM HANDLE

// Predefined tree view control [item handle].
//
// [item handle]: https://learn.microsoft.com/en-us/windows/win32/controls/tree-view-controls#parent-and-child-items
const (
	HTREEITEM_ROOT  HTREEITEM = 0x1_0000
	HTREEITEM_FIRST HTREEITEM = 0x0_ffff
	HTREEITEM_LAST  HTREEITEM = 0x0_fffe
	HTREEITEM_SORT  HTREEITEM = 0x0_fffd
)

// [EDITBALLOONTIP] struct.
//
// ⚠️ You must call [EDITBALLOONTIP.SetCbStruct] to initialize the struct.
//
// # Example
//
//	var eb win.EDITBALLOONTIP
//	eb.SetCbStruct()
//
// [EDITBALLOONTIP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-editballoontip
type EDITBALLOONTIP struct {
	cbStruct uint32
	PszTitle *uint16
	PszText  *uint16
	TtiIcon  co.TTI
}

// Sets the cbStruct field to the size of the struct, correctly initializing it.
func (eb *EDITBALLOONTIP) SetCbStruct() {
	eb.cbStruct = uint32(unsafe.Sizeof(*eb))
}

// [HDITEM] struct.
//
// [HDITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-hditemw
type HDITEM struct {
	Mask       co.HDI
	Cxy        int32
	pszText    *uint16
	Hbm        HBITMAP
	cchTextMax int32
	Fmt        co.HDF
	LParam     LPARAM
	IImage     int32
	IOrder     int32
	Type       co.HDFT
	PvFilter   uintptr
	State      co.HDIS
}

func (hdi *HDITEM) PszText() []uint16 {
	return unsafe.Slice(hdi.pszText, hdi.cchTextMax)
}
func (hdi *HDITEM) SetPszText(val []uint16) {
	hdi.cchTextMax = int32(len(val))
	hdi.pszText = &val[0]
}

// [IMAGEINFO] struct.
//
// [IMAGEINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imageinfo
type IMAGEINFO struct {
	HbmImage HBITMAP
	HbmMask  HBITMAP
	unused1  int32
	unused2  int32
	RcImage  RECT
}

// [IMAGELISTDRAWPARAMS] struct.
//
// ⚠️ You must call [IMAGELISTDRAWPARAMS.SetCbSize] to initialize the struct.
//
// # Example
//
//	var idp win.IMAGELISTDRAWPARAMS
//	idp.SetCbSize()
//
// [IMAGELISTDRAWPARAMS]: https://learn.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type IMAGELISTDRAWPARAMS struct {
	cbSize       uint32
	Himl         HIMAGELIST
	I            int32
	HdcDst       HDC
	X, Y, Cx, Cy int32
	XBitmap      int32
	YBitmap      int32
	RgbBk        COLORREF
	RgbFg        COLORREF
	FStyle       co.ILD
	DwRop        co.ROP
	FState       co.ILS
	Frame        uint32
	CrEffect     COLORREF
}

// Sets the cbSize field to the size of the struct, correctly initializing it.
func (idp *IMAGELISTDRAWPARAMS) SetCbSize() {
	idp.cbSize = uint32(unsafe.Sizeof(*idp))
}

// [_INITCOMMONCONTROLSEX] struct.
//
// ⚠️ You must call [_INITCOMMONCONTROLSEX.SetDwSize] to initialize the struct.
//
// # Example
//
//	var icc win._INITCOMMONCONTROLSEX
//	icc.SetDwSize()
//
// [_INITCOMMONCONTROLSEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-initcommoncontrolsex
type _INITCOMMONCONTROLSEX struct {
	dwSize uint32
	DwICC  co.ICC
}

// Sets the dwSize field to the size of the struct, correctly initializing it.
func (icc *_INITCOMMONCONTROLSEX) SetDwSize() {
	icc.dwSize = uint32(unsafe.Sizeof(*icc))
}

// [LITEM] struct.
//
// [LITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LITEM struct {
	Mask      co.LIF
	ILink     int32
	State     co.LIS
	StateMask co.LIS
	szID      [utl.MAX_LINKID_TEXT]uint16
	szUrl     [utl.L_MAX_URL_LENGTH]uint16
}

func (li *LITEM) SzID() string {
	return wstr.DecodeSlice(li.szID[:])
}
func (li *LITEM) SetSzID(val string) {
	wstr.EncodeToBuf(val, li.szID[:])
}

func (li *LITEM) SzUrl() string {
	return wstr.DecodeSlice(li.szUrl[:])
}
func (li *LITEM) SetSzUrl(val string) {
	wstr.EncodeToBuf(val, li.szUrl[:])
}

// [LVCOLUMN] struct.
//
// [LVCOLUMN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
type LVCOLUMN struct {
	Mask       co.LVCF
	Fmt        co.LVCFMT_C
	Cx         int32
	pszText    *uint16
	cchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
	CxMin      int32
	CxDefault  int32
	CxIdeal    int32
}

func (lvc *LVCOLUMN) PszText() []uint16 {
	return unsafe.Slice(lvc.pszText, lvc.cchTextMax)
}
func (lvc *LVCOLUMN) SetPszText(val []uint16) {
	lvc.cchTextMax = int32(len(val))
	lvc.pszText = &val[0]
}

// [LVFINDINFO] struct.
//
// [LVFINDINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvfindinfow
type LVFINDINFO struct {
	Flags       co.LVFI
	Psz         *uint16
	LParam      LPARAM
	Pt          POINT
	vkDirection uint32 // should bt uint16
}

func (fi *LVFINDINFO) VkDirection() co.VK {
	return co.VK(fi.vkDirection)
}
func (fi *LVFINDINFO) SetVkDirection(val co.VK) {
	fi.vkDirection = uint32(val)
}

// [LVHITTESTINFO] struct.
//
// [LVHITTESTINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvhittestinfo
type LVHITTESTINFO struct {
	Pt       POINT // Coordinates relative to list view.
	Flags    co.LVHT
	IItem    int32 // -1 if no item.
	ISubItem int32
	IGroup   int32
}

// [LVITEM] struct.
//
// [LVITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVITEM struct {
	Mask       co.LVIF
	IItem      int32
	ISubItem   int32
	State      co.LVIS
	StateMask  co.LVIS
	pszText    *uint16
	cchTextMax int32
	IImage     int32
	LParam     LPARAM
	IIndent    int32
	IGroupId   co.LVI_GROUPID
	CColumns   uint32
	PuColumns  *uint32
	PiColFmt   *co.LVCFMT_I
	IGroup     int32
}

func (lvi *LVITEM) PszText() []uint16 {
	return unsafe.Slice(lvi.pszText, lvi.cchTextMax)
}
func (lvi *LVITEM) SetPszText(val []uint16) {
	lvi.cchTextMax = int32(len(val))
	lvi.pszText = &val[0]
}

// [LVITEMINDEX] struct.
//
// [LVITEMINDEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemindex
type LVITEMINDEX struct {
	IItem  int32
	IGroup int32
}

// [NMBCDROPDOWN] struct.
//
// [NMBCDROPDOWN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmbcdropdown
type NMBCDROPDOWN struct {
	Hdr      NMHDR
	RcButton RECT
}

// [NMBCHOTITEM] struct.
//
// [NMBCHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmbchotitem
type NMBCHOTITEM struct {
	Hdr     NMHDR
	DwFlags co.HICF
}

// [NMCHAR] struct.
//
// [NMCHAR]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmchar
type NMCHAR struct {
	Hdr        NMHDR
	Ch         uint32
	DwItemPrev uint32
	DwItemNext uint32
}

// [NMCUSTOMDRAW] struct;
//
// [NMCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type NMCUSTOMDRAW struct {
	Hdr         NMHDR
	DwDrawStage co.CDDS
	Hdc         HDC
	Rc          RECT
	DwItemSpec  uintptr // DWORD_PTR
	UItemState  co.CDIS
	LItemlParam LPARAM
}

// [NMDATETIMECHANGE] struct.
//
// [NMDATETIMECHANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimechange
type NMDATETIMECHANGE struct {
	Nmhdr   NMHDR
	DwFlags co.GDT
	St      SYSTEMTIME
}

// [NMDATETIMEFORMAT] struct.
//
// [NMDATETIMEFORMAT]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatw
type NMDATETIMEFORMAT struct {
	Nmhdr      NMHDR
	PszFormat  *uint16
	St         SYSTEMTIME
	PszDisplay *uint16
	szDisplay  [64]uint16
}

func (dtf *NMDATETIMEFORMAT) SzDisplay() string {
	return wstr.DecodeSlice(dtf.szDisplay[:])
}
func (dtf *NMDATETIMEFORMAT) SetSzDisplay(val string) {
	wstr.EncodeToBuf(val, dtf.szDisplay[:])
}

// [NMDATETIMEFORMATQUERY] struct.
//
// [NMDATETIMEFORMATQUERY]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatqueryw
type NMDATETIMEFORMATQUERY struct {
	Nmhdr     NMHDR
	PszFormat *uint16
	SzMax     SIZE
}

// [NMDATETIMESTRING] struct.
//
// [NMDATETIMESTRING]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimestringw
type NMDATETIMESTRING struct {
	Nmhdr         NMHDR
	PszUserString *uint16
	St            SYSTEMTIME
	DwFlags       co.GDT
}

// [NMDATETIMEWMKEYDOWN] struct.
//
// [NMDATETIMEWMKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimewmkeydownw
type NMDATETIMEWMKEYDOWN struct {
	Nmhdr     NMHDR
	nVirtKey  int32 // should be uint16
	PszFormat *uint16
	St        SYSTEMTIME
}

func (dtk *NMDATETIMEWMKEYDOWN) NVirtKey() co.VK {
	return co.VK(dtk.nVirtKey)
}
func (dtk *NMDATETIMEWMKEYDOWN) SetNVirtKey(val co.VK) {
	dtk.nVirtKey = int32(val)
}

// [NMDAYSTATE] struct.
//
// [NMDAYSTATE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdaystate
type NMDAYSTATE struct {
	Nmhdr       NMHDR
	StStart     SYSTEMTIME
	CDayState   int32
	PrgDayState *uint32 // *MONTHDAYSTATE
}

// [NMHDDISPINFO] struct.
//
// [NMHDDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmhddispinfow
type NMHDDISPINFO struct {
	Hdr        NMHDR
	IItem      int32
	Mask       co.HDI
	pszText    *uint16
	cchTextMax int32
	IImage     int32
	LParam     LPARAM
}

func (hdi *NMHDDISPINFO) PszText() []uint16 {
	return unsafe.Slice(hdi.pszText, hdi.cchTextMax)
}
func (hdi *NMHDDISPINFO) SetPszText(val []uint16) {
	hdi.cchTextMax = int32(len(val))
	hdi.pszText = &val[0]
}

// [NMHDFILTERBTNCLICK] struct.
//
// [NMHDFILTERBTNCLICK]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmhdfilterbtnclick
type NMHDFILTERBTNCLICK struct {
	Hdr   NMHDR
	IItem int32
	Rc    RECT
}

// [NMHDR] struct.
//
// [NMHDR]: https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-nmhdr
type NMHDR struct {
	HWndFrom HWND
	IdFrom   uintptr // UINT_PTR, actually it's a simple control ID
	Code     uint32  // in fact it should be int32
}

// [NMHEADER] struct.
//
// [NMHEADER]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmheaderw
type NMHEADER struct {
	Hdr     NMHDR
	IItem   int32
	IButton co.HEADER_BTN
	PItem   *HDITEM
}

// [NMITEMACTIVATE] struct.
//
// [NMITEMACTIVATE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmitemactivate
type NMITEMACTIVATE struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState co.LVIS
	UOldState co.LVIS
	UChanged  co.LVIF
	PtAction  POINT
	LParam    LPARAM
	UKeyFlags co.LVKF
}

// [NMKEY] struct.
//
// [NMKEY]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmkey
type NMKEY struct {
	Hdr    NMHDR
	nVKey  uint32 // should be uint16
	uFlags uint32
}

func (nmk *NMKEY) NVKey() co.VK       { return co.VK(nmk.nVKey) }
func (nmk *NMKEY) SetNVKey(val co.VK) { nmk.nVKey = uint32(val) }

func (nmk *NMKEY) ScanCode() uint8 { return LOBYTE(LOWORD(nmk.uFlags)) }
func (nmk *NMKEY) SetScanCode(val uint8) {
	nmk.uFlags = MAKELONG(
		MAKEWORD(val, HIBYTE(LOWORD(nmk.uFlags))),
		HIWORD(nmk.uFlags),
	)
}

func (nmk *NMKEY) IsExtendedKey() bool { return utl.BitIsSet(HIBYTE(LOWORD(nmk.uFlags)), 0) }
func (nmk *NMKEY) SetIsExtendedKey(val bool) {
	nmk.uFlags = MAKELONG(
		MAKEWORD(
			LOBYTE(LOWORD(nmk.uFlags)),
			utl.BitSet(HIBYTE(LOWORD(nmk.uFlags)), 0, val),
		),
		HIWORD(nmk.uFlags),
	)
}

func (nmk *NMKEY) ContextCode() bool { return utl.BitIsSet(HIBYTE(LOWORD(nmk.uFlags)), 5) }
func (nmk *NMKEY) SetContextCode(val bool) {
	nmk.uFlags = MAKELONG(
		MAKEWORD(
			LOBYTE(LOWORD(nmk.uFlags)),
			utl.BitSet(HIBYTE(LOWORD(nmk.uFlags)), 5, val),
		),
		HIWORD(nmk.uFlags),
	)
}

func (nmk *NMKEY) IsKeyDownBeforeSend() bool { return utl.BitIsSet(HIBYTE(LOWORD(nmk.uFlags)), 6) }
func (nmk *NMKEY) SetIsKeyDownBeforeSend(val bool) {
	nmk.uFlags = MAKELONG(
		MAKEWORD(
			LOBYTE(LOWORD(nmk.uFlags)),
			utl.BitSet(HIBYTE(LOWORD(nmk.uFlags)), 6, val),
		),
		HIWORD(nmk.uFlags),
	)
}

func (nmk *NMKEY) TransitionState() bool { return utl.BitIsSet(HIBYTE(LOWORD(nmk.uFlags)), 7) }
func (nmk *NMKEY) SetTransitionState(val bool) {
	nmk.uFlags = MAKELONG(
		MAKEWORD(
			LOBYTE(LOWORD(nmk.uFlags)),
			utl.BitSet(HIBYTE(LOWORD(nmk.uFlags)), 7, val),
		),
		HIWORD(nmk.uFlags),
	)
}

// [NMLINK] struct.
//
// [NMLINK]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlink
type NMLINK struct {
	Hdr  NMHDR
	Item LITEM
}

// [NMLISTVIEW] struct.
//
// [NMLISTVIEW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlistview
type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState co.LVIS
	UOldState co.LVIS
	UChanged  co.LVIF
	PtAction  POINT
	LParam    LPARAM
}

// [NMLVCACHEHINT] struct.
//
// [NMLVCACHEHINT]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcachehint
type NMLVCACHEHINT struct {
	Hdr   NMHDR
	IFrom int32
	ITo   int32
}

// [NMLVCUSTOMDRAW] struct.
//
// [NMLVCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
type NMLVCUSTOMDRAW struct {
	Nmcd        NMCUSTOMDRAW
	ClrText     COLORREF
	ClrTextBk   COLORREF
	ISubItem    int32
	DwItemType  co.LVCDI
	ClrFace     COLORREF
	IIconEffect int32
	IIconPhase  int32
	iPartId     int32
	iStateId    int32
	RcText      RECT
	UAlign      co.LVGA_HEADER
}

func (lcd *NMLVCUSTOMDRAW) PartStateId() co.VS {
	return co.VS(utl.Make32(uint16(lcd.iStateId), uint16(lcd.iPartId)))
}
func (lcd *NMLVCUSTOMDRAW) SetPartStateId(val co.VS) {
	lcd.iPartId = val.Part()
	lcd.iStateId = val.State()
}

// [NMLVDISPINFO] struct.
//
// [NMLVDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvdispinfow
type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

// [NMLVEMPTYMARKUP] struct.
//
// [NMLVEMPTYMARKUP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvemptymarkup
type NMLVEMPTYMARKUP struct {
	Hdr      NMHDR
	DwFlags  co.EMF
	szMarkup [utl.L_MAX_URL_LENGTH]uint16
}

func (lve *NMLVEMPTYMARKUP) SzMarkup() string {
	return wstr.DecodeSlice(lve.szMarkup[:])
}
func (lve *NMLVEMPTYMARKUP) SetSzMarkup(val string) {
	wstr.EncodeToBuf(val, lve.szMarkup[:])
}

// [NMLVFINDITEM] struct.
//
// [NMLVFINDITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvfinditemw
type NMLVFINDITEM struct {
	Hdr    NMHDR
	IStart int32
	Lvfi   LVFINDINFO
}

// [NMLVGETINFOTIP] struct.
//
// [NMLVGETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvgetinfotipw
type NMLVGETINFOTIP struct {
	Hdr        NMHDR
	DwFlags    co.LVGIT
	pszText    *uint16
	cchTextMax int32
	IItem      int32
	ISubItem   int32
	LParam     LPARAM
}

func (git *NMLVGETINFOTIP) PszText() []uint16 {
	return unsafe.Slice(git.pszText, git.cchTextMax)
}
func (git *NMLVGETINFOTIP) SetPszText(val []uint16) {
	git.cchTextMax = int32(len(val))
	git.pszText = &val[0]
}

// [NMLVKEYDOWN] struct.
//
// [NMLVKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvkeydown
type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// [NMLVLINK] struct.
//
// [NMLVLINK]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvlink
type NMLVLINK struct {
	Hdr      NMHDR
	Link     LITEM
	IItem    int32
	ISubItem int32
}

// [NMLVODSTATECHANGE] struct.
//
// [NMLVODSTATECHANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvodstatechange
type NMLVODSTATECHANGE struct {
	Hdr       NMHDR
	IFrom     int32
	ITo       int32
	UNewState co.LVIS
	UOldState co.LVIS
}

// [NMLVSCROLL] struct.
//
// [NMLVSCROLL]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvscroll
type NMLVSCROLL struct {
	Hdr    NMHDR
	Dx, Dy int32
}

// [NMMOUSE] struct.
//
// [NMMOUSE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmmouse
type NMMOUSE struct {
	Hdr        NMHDR
	DwItemSpec uintptr // DWORD_PTR
	DwItemData uintptr // DWORD_PTR
	Pt         POINT
	DwHitInfo  LPARAM
}

// [NMOBJECTNOTIFY] struct.
//
// [NMOBJECTNOTIFY]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmobjectnotify
type NMOBJECTNOTIFY struct {
	Hdr      NMHDR
	IItem    int32
	Piid     *co.IID
	Object   uintptr
	HrResult co.HRESULT
	DwFlags  uint32
}

// [NMSELCHANGE] struct.
//
// [NMSELCHANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmselchange
type NMSELCHANGE struct {
	Nmhdr      NMHDR
	StSelStart SYSTEMTIME
	StSelEnd   SYSTEMTIME
}

// [NMTBCUSTOMDRAW] struct.
//
// [NMTBCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbcustomdraw
type NMTBCUSTOMDRAW struct {
	Nmcd                 NMCUSTOMDRAW
	HbrMonoDither        HBRUSH
	HbrLines             HBRUSH
	HpenLines            HPEN
	ClrText              COLORREF
	ClrMark              COLORREF
	ClrTextHighlight     COLORREF
	ClrBtnFace           COLORREF
	ClrBtnHighlight      COLORREF
	ClrHighlightHotTrack COLORREF
	RcText               RECT
	NStringBkMode        co.BKMODE
	NHLStringBkMode      co.BKMODE
	IListGap             int32
}

// [NMTBDISPINFO] struct.
//
// [NMTBDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbdispinfow
type NMTBDISPINFO struct {
	Hdr       NMHDR
	DwMask    co.TBNF
	IdCommand int32
	LParam    LPARAM
	IImage    int32
	pszText   *uint16
	cchText   int32
}

func (di *NMTBDISPINFO) PszText() []uint16 {
	return unsafe.Slice(di.pszText, di.cchText)
}
func (di *NMTBDISPINFO) SetPszText(val []uint16) {
	di.cchText = int32(len(val))
	di.pszText = &val[0]
}

// [NMTBDUPACCELERATOR] struct.
//
// [NMTBDUPACCELERATOR]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-dupaccelerator
type NMTBDUPACCELERATOR struct {
	Hdr  NMHDR
	Ch   uint32
	FDup int32 // This is a BOOL value.
}

// [NMTBGETINFOTIP] struct.
//
// [NMTBGETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbgetinfotipw
type NMTBGETINFOTIP struct {
	Hdr        NMHDR
	pszText    *uint16
	cchTextMax int32
	IItem      int32
	LParam     LPARAM
}

func (git *NMTBGETINFOTIP) PszText() []uint16 {
	return unsafe.Slice(git.pszText, git.cchTextMax)
}
func (git *NMTBGETINFOTIP) SetPszText(val []uint16) {
	git.cchTextMax = int32(len(val))
	git.pszText = &val[0]
}

// [NMTBHOTITEM] struct.
//
// [NMTBHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbhotitem
type NMTBHOTITEM struct {
	Hdr     NMHDR
	IdOld   int32
	IdNew   int32
	DwFlags co.HICF
}

// [NMTBRESTORE] struct.
//
// [NMTBRESTORE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbrestore
type NMTBRESTORE struct {
	Hdr              NMHDR
	PData            *uint32
	PCurrent         *uint32
	CbData           uint32
	IItem            int32
	CButtons         int32
	CbBytesPerRecord int32
	TbButton         TBBUTTON
}

// [NMTBSAVE] struct.
//
// [NMTBSAVE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbsave
type NMTBSAVE struct {
	Hdr      NMHDR
	PData    *uint32
	PCurrent *uint32
	CbData   uint32
	IItem    int32
	CButtons int32
	TbButton TBBUTTON
}

// [NMTBWRAPACCELERATOR] struct.
//
// [NMTBWRAPACCELERATOR]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-wrapaccelerator
type NMTBWRAPACCELERATOR struct {
	Hdr     NMHDR
	Ch      uint32
	IButton int32
}

// [NMTBWRAPHOTITEM] struct.
//
// [NMTBWRAPHOTITEM]: https://learn.microsoft.com/en-us/windows/win32/controls/tbn-wraphotitem
type NMTBWRAPHOTITEM struct {
	Hdr     NMHDR
	IStart  int32
	IDir    int32
	NReason co.HICF
}

// [NMTCKEYDOWN] struct.
//
// [NMTCKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtckeydown
type NMTCKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// [NMTOOLBAR] struct.
//
// [NMTOOLBAR]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtoolbarw
type NMTOOLBAR struct {
	Hdr      NMHDR
	IItem    int32
	TbButton TBBUTTON
	cchText  int32
	pszText  *uint16
	RcButton RECT
}

func (tb *NMTOOLBAR) PszText() []uint16 {
	return unsafe.Slice(tb.pszText, tb.cchText)
}
func (tb *NMTOOLBAR) SetPszText(val []uint16) {
	tb.cchText = int32(len(val))
	tb.pszText = &val[0]
}

// [NMTOOLTIPSCREATED] struct.
//
// [NMTOOLTIPSCREATED]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtooltipscreated
type NMTOOLTIPSCREATED struct {
	Hdr          NMHDR
	HwndToolTips HWND
}

// [NMTRBTHUMBPOSCHANGING] struct.
//
// [NMTRBTHUMBPOSCHANGING]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtrbthumbposchanging
type NMTRBTHUMBPOSCHANGING struct {
	Hdr     NMHDR
	DwPos   uint32
	NReason co.TB_REQ
}

// [NMTREEVIEW] struct.
//
// [NMTREEVIEW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtreevieww
type NMTREEVIEW struct {
	Hdr     NMHDR
	Action  uint32 // co.TVE | co.TVC
	ItemOld TVITEM
	ItemNew TVITEM
	PtDrag  POINT
}

// [NMTVASYNCDRAW] struct.
//
// [NMTVASYNCDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type NMTVASYNCDRAW struct {
	Hdr            NMHDR
	Pimldp         *IMAGELISTDRAWPARAMS
	Hr             co.HRESULT
	Hitem          HTREEITEM
	LParam         LPARAM
	DwRetFlags     co.ADRF
	IRetImageIndex int32
}

// [NMTVCUSTOMDRAW] struct.
//
// [NMTVCUSTOMDRAW]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvcustomdraw
type NMTVCUSTOMDRAW struct {
	Nmcd      NMCUSTOMDRAW
	ClrText   COLORREF
	ClrTextBk COLORREF
	ILevel    int32
}

// [NMTVDISPINFO] struct.
//
// [NMTVDISPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfow
type NMTVDISPINFO struct {
	Hdr  NMHDR
	Item TVITEM
}

// [NMTVGETINFOTIP] struct.
//
// [NMTVGETINFOTIP]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvgetinfotipw
type NMTVGETINFOTIP struct {
	Hdr        NMHDR
	pszText    *uint16
	cchTextMax int32
	HItem      HTREEITEM
	LParam     LPARAM
}

func (git *NMTVGETINFOTIP) PszText() []uint16 {
	return unsafe.Slice(git.pszText, git.cchTextMax)
}
func (git *NMTVGETINFOTIP) SetPszText(val []uint16) {
	git.cchTextMax = int32(len(val))
	git.pszText = &val[0]
}

// [NMTVKEYDOWN] struct.
//
// [NMTVKEYDOWN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvkeydown
type NMTVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// [NMTVITEMCHANGE] struct.
//
// [NMTVITEMCHANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvitemchange
type NMTVITEMCHANGE struct {
	Hdr       NMHDR
	UChanged  co.TVIF
	HItem     HTREEITEM
	UStateNew co.TVIS
	UStateOld co.TVIS
	LParam    LPARAM
}

// [NMUPDOWN] struct.
//
// [NMUPDOWN]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmupdown
type NMUPDOWN struct {
	Hdr    NMHDR
	IPos   int32
	IDelta int32
}

// [NMVIEWCHANGE] struct.
//
// [NMVIEWCHANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmviewchange
type NMVIEWCHANGE struct {
	Nmhdr     NMHDR
	DwOldView co.MCMV
	DwNewView co.MCMV
}

// [PBRANGE] struct.
//
// [PBRANGE]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-pbrange
type PBRANGE struct {
	ILow  int32
	IHigh int32
}

// [TASKDIALOG_BUTTON] struct syntactic sugar.
//
// This struct originally has a packed alignment, so we serialized it before the
// syscall.
//
// [TASKDIALOG_BUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-taskdialog_button
type TASKDIALOG_BUTTON struct {
	Id   co.ID
	Text string
}

// [TASKDIALOGCONFIG] struct syntactic sugar.
//
// This struct originally has a packed alignment, so we serialized it before the
// syscall.
//
// [TASKDIALOGCONFIG]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-taskdialogconfig
type TASKDIALOGCONFIG struct {
	HwndParent           HWND
	HInstance            HINSTANCE
	Flags                co.TDF
	CommonButtons        co.TDCBF
	WindowTitle          string
	HMainIcon            TdcIcon // Union PCWSTR + HICON, but string resource won't be considered.
	MainInstruction      string
	Content              string
	Buttons              []TASKDIALOG_BUTTON
	DefaultButtonId      uint16
	RadioButtons         []TASKDIALOG_BUTTON
	DefaultRadioButton   uint16
	VerificationText     string
	ExpandedInformation  string
	ExpandedControlText  string
	CollapsedControlText string
	HFooterIcon          TdcIcon // Union PCWSTR + HICON, but string resource won't be considered.
	Footer               string
	PfCallback           uintptr
	LpCallbackData       uintptr
	Width                int
}

// Converts the syntactic sugar struct into the packed raw one.
func (tdc *TASKDIALOGCONFIG) serialize(
	pStrsBuf *wstr.BufEncoder, // buffer for all struct and button strings
	pTdcBuf *Vec[byte],
	pBtnsBuf *Vec[[12]byte], // buffer for Buttons and RadioButtons
) {
	dest := pTdcBuf.HotSlice()

	numButtons := uint(len(tdc.Buttons))
	numRadios := uint(len(tdc.RadioButtons))
	pBtnsBuf.Resize(numButtons+numRadios, [12]byte{}) // alloc buffer for buttons + radios

	tdc.put32(dest[0:], uint32(pTdcBuf.Len())) // cbSize
	tdc.put64(dest[4:], uint64(tdc.HwndParent))
	tdc.put64(dest[12:], uint64(tdc.HInstance))
	tdc.put32(dest[20:], uint32(tdc.Flags))
	tdc.put32(dest[24:], uint32(tdc.CommonButtons))

	tdc.putPtr(dest[28:], pStrsBuf.PtrEmptyIsNil(tdc.WindowTitle))

	if !tdc.HMainIcon.IsNone() {
		tdc.put64(dest[36:], tdc.HMainIcon.raw())
	}

	tdc.putPtr(dest[44:], pStrsBuf.PtrEmptyIsNil(tdc.MainInstruction))
	tdc.putPtr(dest[52:], pStrsBuf.PtrEmptyIsNil(tdc.Content))

	if len(tdc.Buttons) > 0 {
		for i, btn := range tdc.Buttons {
			pBtnBuf := pBtnsBuf.Get(uint(i)) // already allocated
			tdc.put32((*pBtnBuf)[:], uint32(btn.Id))
			tdc.putPtr((*pBtnBuf)[4:], pStrsBuf.PtrEmptyIsNil(btn.Text))
		}
		tdc.put32(dest[60:], uint32(len(tdc.Buttons))) // number of buttons
		tdc.putPtr(dest[64:], pBtnsBuf.UnsafePtr())    // ptr to buttons block
	}

	tdc.put32(dest[72:], uint32(tdc.DefaultButtonId))

	if len(tdc.RadioButtons) > 0 {
		baseRadioIdx := uint(len(tdc.Buttons)) // radios are appended after the buttons
		for i, btn := range tdc.RadioButtons {
			pBtnBuf := pBtnsBuf.Get(baseRadioIdx + uint(i)) // already allocated
			tdc.put32((*pBtnBuf)[:], uint32(btn.Id))
			tdc.putPtr((*pBtnBuf)[4:], pStrsBuf.PtrEmptyIsNil(btn.Text))
		}
		tdc.put32(dest[76:], uint32(len(tdc.RadioButtons)))               // number of radios
		tdc.putPtr(dest[80:], unsafe.Pointer(pBtnsBuf.Get(baseRadioIdx))) // ptr to radios block
	}

	tdc.put32(dest[88:], uint32(tdc.DefaultRadioButton))

	tdc.putPtr(dest[92:], pStrsBuf.PtrEmptyIsNil(tdc.VerificationText))
	tdc.putPtr(dest[100:], pStrsBuf.PtrEmptyIsNil(tdc.ExpandedInformation))
	tdc.putPtr(dest[108:], pStrsBuf.PtrEmptyIsNil(tdc.ExpandedControlText))
	tdc.putPtr(dest[116:], pStrsBuf.PtrEmptyIsNil(tdc.CollapsedControlText))

	if !tdc.HFooterIcon.IsNone() {
		tdc.put64(dest[124:], tdc.HFooterIcon.raw())
	}

	tdc.putPtr(dest[132:], pStrsBuf.PtrEmptyIsNil(tdc.Footer))

	tdc.put64(dest[140:], uint64(tdc.PfCallback))
	tdc.put64(dest[148:], uint64(tdc.LpCallbackData))
	tdc.put32(dest[156:], uint32(tdc.Width))
}

func (*TASKDIALOGCONFIG) put32(b []byte, v uint32) {
	binary.LittleEndian.PutUint32(b, v)
}
func (*TASKDIALOGCONFIG) put64(b []byte, v uint64) {
	binary.LittleEndian.PutUint64(b, v)
}
func (tdc *TASKDIALOGCONFIG) putPtr(b []byte, p unsafe.Pointer) {
	tdc.put64(b, uint64(uintptr(p)))
}

// [TCITEM] struct.
//
// [TCITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tcitemw
type TCITEM struct {
	Mask        co.TCIF
	DwState     co.TCIS
	DwStateMask co.TCIS
	pszText     *uint16
	cchTextMax  int32
	IImage      int32
	LParam      LPARAM
}

func (tci *TCITEM) PszText() []uint16 {
	return unsafe.Slice(tci.pszText, tci.cchTextMax)
}
func (tci *TCITEM) SetPszText(val []uint16) {
	tci.cchTextMax = int32(len(val))
	tci.pszText = &val[0]
}

// [TBBUTTON] struct.
//
// [TBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tbbutton
type TBBUTTON struct {
	IBitmap   int32 // With multiple image lists, HIWORD is the image list index.
	IdCommand int32
	FsState   co.TBSTATE
	FsStyle   co.BTNS
	bReserved [6]uint8 // This padding is 2 in 32-bit environments.
	DwData    uintptr
	IString   *uint16 // Can also be the index in the string list.
}

// [TVINSERTSTRUCT] struct.
//
// [TVINSERTSTRUCT]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvinsertstructw
type TVINSERTSTRUCT struct {
	HParent      HTREEITEM
	HInsertAfter HTREEITEM
	Itemex       TVITEMEX
}

// [TVITEM] struct.
//
// [TVITEM]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemw
type TVITEM struct {
	Mask           co.TVIF
	HItem          HTREEITEM
	State          co.TVIS
	StateMask      co.TVIS
	pszText        *uint16
	cchTextMax     int32
	IImage         int32
	ISelectedImage int32
	CChildren      co.TVI_CHILDREN
	LParam         LPARAM
}

func (tvi *TVITEM) PszText() []uint16 {
	return unsafe.Slice(tvi.pszText, tvi.cchTextMax)
}
func (tvi *TVITEM) SetPszText(val []uint16) {
	tvi.cchTextMax = int32(len(val))
	tvi.pszText = &val[0]
}

// [TVITEMEX] struct.
//
// [TVITEMEX]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVITEMEX struct {
	Mask           co.TVIF
	HItem          HTREEITEM
	State          co.TVIS
	StateMask      co.TVIS
	pszText        *uint16
	cchTextMax     int32
	IImage         int32
	ISelectedImage int32
	CChildren      co.TVI_CHILDREN
	LParam         LPARAM
	IIntegral      int32
	UStateEx       co.TVIS_EX
	Hwnd           HWND
	IExpandedImage int32
	iReserved      int32
}

func (tvx *TVITEMEX) PszText() []uint16 {
	return unsafe.Slice(tvx.pszText, tvx.cchTextMax)
}
func (tvx *TVITEMEX) SetPszText(val []uint16) {
	tvx.cchTextMax = int32(len(val))
	tvx.pszText = &val[0]
}
