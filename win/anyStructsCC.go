/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"wingows/co"
)

// https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imageinfo
type IMAGEINFO struct {
	HbmImage HBITMAP
	HbmMask  HBITMAP
	Unused1  int32
	Unused2  int32
	RcImage  RECT
}

// https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type IMAGELISTDRAWPARAMS struct {
	CbSize   uint32
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

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-litem
type LITEM struct {
	Mask      co.LIF
	ILink     int32
	State     co.LIS
	StateMask co.LIS
	SzID      [_MAX_LINKID_TEXT]uint16
	SzUrl     [_L_MAX_URL_LENGTH]uint16
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
type LVCOLUMN struct {
	Mask       co.LVCF
	Fmt        int32
	Cx         int32
	PszText    uintptr // LPWSTR
	CchTextMax int32
	ISubItem   int32
	IImage     int32
	IOrder     int32
	CxMin      int32
	CxDefault  int32
	CxIdeal    int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvfindinfow
type LVFINDINFO struct {
	Flags       co.LVFI
	Psz         uintptr // LPCWSTR
	LParam      LPARAM
	Pt          POINT
	VkDirection uint32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvhittestinfo
type LVHITTESTINFO struct {
	Pt       POINT // Coordinates relative to list view.
	Flags    co.LVHT
	IItem    int32 // -1 if no item.
	ISubItem int32
	IGroup   int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
type LVITEM struct {
	Mask       co.LVIF
	IItem      int32
	ISubItem   int32
	State      co.LVIS
	StateMask  co.LVIS
	PszText    uintptr // LPWSTR
	CchTextMax int32
	IImage     int32
	LParam     LPARAM
	IIndent    int32
	IGroupId   int32
	CColumns   uint32
	PuColumns  uintptr // *uint32
	PiColFmt   uintptr // *int32
	IGroup     int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type NMCUSTOMDRAW struct {
	Hdr         NMHDR
	DwDrawStage co.CDDS
	Hdc         HDC
	Rc          RECT
	DwItemSpec  uintptr
	UItemState  co.CDIS
	LItemlParam LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimechange
type NMDATETIMECHANGE struct {
	Nmhdr   NMHDR
	DwFlags co.GDT
	St      SYSTEMTIME
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatw
type NMDATETIMEFORMAT struct {
	Nmhdr      NMHDR
	PszFormat  uintptr // LPCWSTR
	St         SYSTEMTIME
	pszDisplay uintptr // LPCWSTR
	SzDisplay  [64]uint16
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatqueryw
type NMDATETIMEFORMATQUERY struct {
	Nmhdr     NMHDR
	PszFormat uintptr // LPCWSTR
	SzMax     SIZE
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimestringw
type NMDATETIMESTRING struct {
	Nmhdr         NMHDR
	PszUserString uintptr // LPCWSTR
	St            SYSTEMTIME
	DwFlags       co.GDT
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimewmkeydownw
type NMDATETIMEWMKEYDOWN struct {
	Nmhdr     NMHDR
	NVirtKey  int32
	PszFormat uintptr // LPCWSTR
	St        SYSTEMTIME
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmitemactivate
type NMITEMACTIVATE struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    LPARAM
	UKeyFlags co.LVKF
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlistview
type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int32
	ISubItem  int32
	UNewState uint32
	UOldState uint32
	UChanged  uint32
	PtAction  POINT
	LParam    LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcachehint
type NMLVCACHEHINT struct {
	Hdr   NMHDR
	IFrom int32
	ITo   int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
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

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvdispinfow
type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvemptymarkup
type NMLVEMPTYMARKUP struct {
	Hdr      NMHDR
	DwFlags  co.EMF
	SzMarkup [_L_MAX_URL_LENGTH]uint16
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvfinditemw
type NMLVFINDITEM struct {
	Hdr    NMHDR
	IStart int32
	Lvfi   LVFINDINFO
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvgetinfotipw
type NMLVGETINFOTIP struct {
	Hdr        NMHDR
	DwFlags    co.LVGIT
	PszText    uintptr // LPWSTR
	CchTextMax int32
	IItem      int32
	ISubItem   int32
	LParam     LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvkeydown
type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvlink
type NMLVLINK struct {
	Hdr      NMHDR
	Link     LITEM
	IItem    int32
	ISubItem int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvodstatechange
type NMLVODSTATECHANGE struct {
	Hdr       NMHDR
	IFrom     int32
	ITo       int32
	UNewState co.LVIS
	UOldState co.LVIS
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvscroll
type NMLVSCROLL struct {
	Hdr NMHDR
	Dx  int32
	Dy  int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmmouse
type NMMOUSE struct {
	Hdr        NMHDR
	DwItemSpec uintptr
	DwItemData uintptr
	Pt         POINT
	DwHitInfo  LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtreevieww
type NMTREEVIEW struct {
	Hdr     NMHDR
	Action  uint32 // co.TVE | co.TVC
	ItemOld TVITEM
	ItemNew TVITEM
	PtDrag  POINT
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type NMTVASYNCDRAW struct {
	Hdr            NMHDR
	Pimldp         uintptr  // LPIMAGELISTDRAWPARAMS
	Hr             co.ERROR // HRESULT
	Hitem          HTREEITEM
	LParam         LPARAM
	DwRetFlags     co.ADRF
	IRetImageIndex int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvcustomdraw
type NMTVCUSTOMDRAW struct {
	Nmcd      NMCUSTOMDRAW
	ClrText   COLORREF
	ClrTextBk COLORREF
	ILevel    int32
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfow
type NMTVDISPINFO struct {
	Hdr  NMHDR
	Item TVITEM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvgetinfotipw
type NMTVGETINFOTIP struct {
	Hdr        NMHDR
	PszText    uintptr // LPWSTR
	CchTextMax int32
	HItem      HTREEITEM
	LParam     LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvitemchange
type NMTVITEMCHANGE struct {
	Hdr       NMHDR
	UChanged  co.TVIF
	HItem     HTREEITEM
	UStateNew co.TVIS
	UStateOld co.TVIS
	LParam    LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvkeydown
type NMTVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

// https://www.google.com/search?client=firefox-b-d&q=TVINSERTSTRUCTW
type TVINSERTSTRUCT struct {
	HParent      HTREEITEM
	HInsertAfter HTREEITEM
	Itemex       TVITEMEX
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemw
type TVITEM struct {
	Mask           co.TVIF
	HItem          HTREEITEM
	State          co.TVIS
	StateMask      co.TVIS
	PszText        uintptr // LPWSTR
	CchTextMax     int32
	IImage         int32
	ISelectedImage int32
	CChildren      co.TVI_CHILDREN
	LParam         LPARAM
}

// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVITEMEX struct {
	Mask           co.TVIF
	HItem          HTREEITEM
	State          co.TVIS
	StateMask      co.TVIS
	PszText        uintptr // LPWSTR
	CchTextMax     int32
	IImage         int32
	ISelectedImage int32
	CChildren      co.TVI_CHILDREN
	LParam         LPARAM
	IIntegral      int32
	UStateEx       co.TVIS_EX
	Hwnd           HWND
	IExpandedImage int32
	IReserved      int32
}
