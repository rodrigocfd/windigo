package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// ⚠️ You must call SetCbSize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type IMAGELISTDRAWPARAMS struct {
	cbSize   uint32
	Himl     HIMAGELIST
	I        int32
	HdcDst   HDC
	X        int32
	Y        int32
	Cx       int32
	Cy       int32
	XBitmap  int32
	YBitmap  int32
	RgbBk    COLORREF
	RgbFg    COLORREF
	FStyle   co.ILD
	DwRop    co.ROP
	FState   co.ILS
	Frame    uint32
	CrEffect COLORREF
}

func (ildp *IMAGELISTDRAWPARAMS) SetCbSize() { ildp.cbSize = uint32(unsafe.Sizeof(*ildp)) }

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LITEM struct {
	Mask      co.LIF
	ILink     int32
	State     co.LIS
	StateMask co.LIS
	szID      [_MAX_LINKID_TEXT]uint16
	szUrl     [_L_MAX_URL_LENGTH]uint16
}

func (li *LITEM) SzID() string { return Str.FromUint16Slice(li.szID[:]) }
func (li *LITEM) SetSzID(val string) {
	copy(li.szID[:], Str.ToUint16Slice(Str.Substr(val, 0, _MAX_LINKID_TEXT-1)))
}

func (li *LITEM) SzUrl() string { return Str.FromUint16Slice(li.szUrl[:]) }
func (li *LITEM) SetSzUrl(val string) {
	copy(li.szUrl[:], Str.ToUint16Slice(Str.Substr(val, 0, _L_MAX_URL_LENGTH-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
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

func (lvc *LVCOLUMN) PszText() []uint16 { return unsafe.Slice(lvc.pszText, lvc.cchTextMax) }
func (lvc *LVCOLUMN) SetPszText(val []uint16) {
	lvc.cchTextMax = int32(len(val))
	lvc.pszText = &val[0]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvfindinfow
type LVFINDINFO struct {
	Flags       co.LVFI
	Psz         *uint16
	LParam      LPARAM
	Pt          POINT
	VkDirection uint32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvhittestinfo
type LVHITTESTINFO struct {
	Pt       POINT // Coordinates relative to list view.
	Flags    co.LVHT
	IItem    int32 // -1 if no item.
	ISubItem int32
	IGroup   int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
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

func (lvi *LVITEM) PszText() []uint16 { return unsafe.Slice(lvi.pszText, lvi.cchTextMax) }
func (lvi *LVITEM) SetPszText(val []uint16) {
	lvi.cchTextMax = int32(len(val))
	lvi.pszText = &val[0]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmbcdropdown
type NMBCDROPDOWN struct {
	Hdr      NMHDR
	RcButton RECT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmbchotitem
type NMBCHOTITEM struct {
	Hdr     NMHDR
	DwFlags co.HICF
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmchar
type NMCHAR struct {
	Hdr        NMHDR
	Ch         uint32
	DwItemPrev uint32
	DwItemNext uint32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type NMCUSTOMDRAW struct {
	Hdr         NMHDR
	DwDrawStage co.CDDS
	Hdc         HDC
	Rc          RECT
	DwItemSpec  uintptr // DWORD_PTR
	UItemState  co.CDIS
	LItemlParam LPARAM
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimechange
type NMDATETIMECHANGE struct {
	Nmhdr   NMHDR
	DwFlags co.GDT
	St      SYSTEMTIME
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatw
type NMDATETIMEFORMAT struct {
	Nmhdr      NMHDR
	PszFormat  *uint16
	St         SYSTEMTIME
	PszDisplay *uint16
	SzDisplay  [64]uint16
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatqueryw
type NMDATETIMEFORMATQUERY struct {
	Nmhdr     NMHDR
	PszFormat *uint16
	SzMax     SIZE
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimestringw
type NMDATETIMESTRING struct {
	Nmhdr         NMHDR
	PszUserString *uint16
	St            SYSTEMTIME
	DwFlags       co.GDT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimewmkeydownw
type NMDATETIMEWMKEYDOWN struct {
	Nmhdr     NMHDR
	NVirtKey  int32
	PszFormat *uint16
	St        SYSTEMTIME
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdaystate
type NMDAYSTATE struct {
	Nmhdr       NMHDR
	StStart     SYSTEMTIME
	CDayState   int32
	PrgDayState *uint32 // *MONTHDAYSTATE
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmitemactivate
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmkey
type NMKEY struct {
	Hdr    NMHDR
	nVKey  uint32
	uFlags uint32
}

func (nmk *NMKEY) NVKey() co.VK       { return co.VK(nmk.nVKey) }
func (nmk *NMKEY) SetNVKey(val co.VK) { nmk.nVKey = uint32(val) }

func (nmk *NMKEY) ScanCode() uint { return uint(Bytes.Lo8(Bytes.Hi16(nmk.uFlags))) }

// func (nmk *NMKEY) SetScanCode(val uint) {
// 	nmk.uFlags |= uint32(uint8(val)) << 24

// 	nmk.uFlags = Bytes.Make32(
// 		Bytes.Lo16(nmk.uFlags),
// 		Bytes.Make16(uint8(val), Bytes.Hi8(Bytes.Hi16(nmk.uFlags))),
// 	)
// }

func (nmk *NMKEY) IsExtendedKey() bool { return (Bytes.Hi8(Bytes.Hi16(nmk.uFlags)) & 0b0000_0001) != 0 }

// func (nmk *NMKEY) SetIsExtendedKey(val bool) {
// 	nmk.uFlags = Bytes.Make32(
// 		Bytes.Lo16(nmk.uFlags),
// 		Bytes.Make16(Bytes.Lo8(Bytes.Hi16(nmk.uFlags)),   ),
// 	)
// }

func (nmk *NMKEY) HasAltKey() bool { return (Bytes.Hi8(Bytes.Hi16(nmk.uFlags)) & 0b0010_0000) != 0 }

func (nmk *NMKEY) IsKeyDownBeforeSend() bool {
	return (Bytes.Hi8(Bytes.Hi16(nmk.uFlags)) & 0b0100_0000) != 0
}

func (nmk *NMKEY) IsReleasingKey() bool {
	return (Bytes.Hi8(Bytes.Hi16(nmk.uFlags)) & 0b1000_0000) != 0
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlink
type NMLINK struct {
	Hdr  NMHDR
	Item LITEM
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlistview
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcachehint
type NMLVCACHEHINT struct {
	Hdr   NMHDR
	IFrom int32
	ITo   int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
type NMLVCUSTOMDRAW struct {
	Nmcd        NMCUSTOMDRAW
	ClrText     COLORREF
	ClrTextBk   COLORREF
	ISubItem    int32
	DwItemType  co.LVCDI
	ClrFace     COLORREF
	IIconEffect int32
	IIconPhase  int32
	IPartId     int32
	IStateId    int32
	RcText      RECT
	UAlign      co.LVGA_HEADER
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvdispinfow
type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvemptymarkup
type NMLVEMPTYMARKUP struct {
	Hdr      NMHDR
	DwFlags  co.EMF
	szMarkup [_L_MAX_URL_LENGTH]uint16
}

func (lve *NMLVEMPTYMARKUP) SzMarkup() string { return Str.FromUint16Slice(lve.szMarkup[:]) }
func (lve *NMLVEMPTYMARKUP) SetSzMarkup(val string) {
	copy(lve.szMarkup[:], Str.ToUint16Slice(Str.Substr(val, 0, _L_MAX_URL_LENGTH-1)))
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvfinditemw
type NMLVFINDITEM struct {
	Hdr    NMHDR
	IStart int32
	Lvfi   LVFINDINFO
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvgetinfotipw
type NMLVGETINFOTIP struct {
	Hdr        NMHDR
	DwFlags    co.LVGIT
	pszText    *uint16
	cchTextMax int32
	IItem      int32
	ISubItem   int32
	LParam     LPARAM
}

func (git *NMLVGETINFOTIP) PszText() []uint16 { return unsafe.Slice(git.pszText, git.cchTextMax) }
func (git *NMLVGETINFOTIP) SetPszText(val []uint16) {
	git.cchTextMax = int32(len(val))
	git.pszText = &val[0]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvkeydown
type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvlink
type NMLVLINK struct {
	Hdr      NMHDR
	Link     LITEM
	IItem    int32
	ISubItem int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvodstatechange
type NMLVODSTATECHANGE struct {
	Hdr       NMHDR
	IFrom     int32
	ITo       int32
	UNewState co.LVIS
	UOldState co.LVIS
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvscroll
type NMLVSCROLL struct {
	Hdr NMHDR
	Dx  int32
	Dy  int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmmouse
type NMMOUSE struct {
	Hdr        NMHDR
	DwItemSpec uintptr // DWORD_PTR
	DwItemData uintptr // DWORD_PTR
	Pt         POINT
	DwHitInfo  LPARAM
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmselchange
type NMSELCHANGE struct {
	Nmhdr      NMHDR
	StSelStart SYSTEMTIME
	StSelEnd   SYSTEMTIME
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtbcustomdraw
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

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtooltipscreated
type NMTOOLTIPSCREATED struct {
	Hdr          NMHDR
	HwndToolTips HWND
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtrbthumbposchanging
type NMTRBTHUMBPOSCHANGING struct {
	Hdr     NMHDR
	DwPos   uint32
	NReason co.TB_REQ
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtreevieww
type NMTREEVIEW struct {
	Hdr     NMHDR
	Action  uint32 // co.TVE | co.TVC
	ItemOld TVITEM
	ItemNew TVITEM
	PtDrag  POINT
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type NMTVASYNCDRAW struct {
	Hdr            NMHDR
	Pimldp         *IMAGELISTDRAWPARAMS
	Hr             errco.ERROR // HRESULT
	Hitem          HTREEITEM
	LParam         LPARAM
	DwRetFlags     co.ADRF
	IRetImageIndex int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvcustomdraw
type NMTVCUSTOMDRAW struct {
	Nmcd      NMCUSTOMDRAW
	ClrText   COLORREF
	ClrTextBk COLORREF
	ILevel    int32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfow
type NMTVDISPINFO struct {
	Hdr  NMHDR
	Item TVITEM
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvgetinfotipw
type NMTVGETINFOTIP struct {
	Hdr        NMHDR
	pszText    *uint16
	cchTextMax int32
	HItem      HTREEITEM
	LParam     LPARAM
}

func (git *NMTVGETINFOTIP) PszText() []uint16 { return unsafe.Slice(git.pszText, git.cchTextMax) }
func (git *NMTVGETINFOTIP) SetPszText(val []uint16) {
	git.cchTextMax = int32(len(val))
	git.pszText = &val[0]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvitemchange
type NMTVITEMCHANGE struct {
	Hdr       NMHDR
	UChanged  co.TVIF
	HItem     HTREEITEM
	UStateNew co.TVIS
	UStateOld co.TVIS
	LParam    LPARAM
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvkeydown
type NMTVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmviewchange
type NMVIEWCHANGE struct {
	Nmhdr     NMHDR
	DwOldView co.MCMV
	DwNewView co.MCMV
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tbbutton
type TBBUTTON struct {
	iBitmap   int32
	IdCommand int32
	FsState   co.TBSTATE
	FsStyle   co.BTNS
	bReserved [6]uint8 // this padding is 2 in 32-bit environments
	DwData    uintptr
	IString   *uint16 // can also be the index in the string list
}

func (tbb *TBBUTTON) IBitmap() (icon, imgList int) {
	icon = int(Bytes.Lo16(uint32(tbb.iBitmap)))
	imgList = int(Bytes.Hi16(uint32(tbb.iBitmap)))
	return
}
func (tbb *TBBUTTON) SetIBitmap(icon, imgList int) {
	tbb.iBitmap = int32(Bytes.Make32(uint16(icon), uint16(imgList)))
}

// 📑 https://www.google.com/search?client=firefox-b-d&q=TVINSERTSTRUCTW
type TVINSERTSTRUCT struct {
	HParent      HTREEITEM
	HInsertAfter HTREEITEM
	Itemex       TVITEMEX
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemw
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

func (tvi *TVITEM) PszText() []uint16 { return unsafe.Slice(tvi.pszText, tvi.cchTextMax) }
func (tvi *TVITEM) SetPszText(val []uint16) {
	tvi.cchTextMax = int32(len(val))
	tvi.pszText = &val[0]
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
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

func (tvx *TVITEMEX) PszText() []uint16 { return unsafe.Slice(tvx.pszText, tvx.cchTextMax) }
func (tvx *TVITEMEX) SetPszText(val []uint16) {
	tvx.cchTextMax = int32(len(val))
	tvx.pszText = &val[0]
}
