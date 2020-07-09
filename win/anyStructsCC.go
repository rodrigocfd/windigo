/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"wingows/co"
)

type IMAGEINFO struct {
	HbmImage HBITMAP
	HbmMask  HBITMAP
	Unused1  int32
	Unused2  int32
	RcImage  RECT
}

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

type LITEM struct {
	Mask      co.LIF
	ILink     int32
	State     co.LIS
	StateMask co.LIS
	SzID      [48]uint16            // MAX_LINKID_TEXT
	SzUrl     [2048 + 32 + 3]uint16 // L_MAX_URL_LENGTH
}

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

type LVFINDINFO struct {
	Flags       co.LVFI
	Psz         uintptr // LPCWSTR
	LParam      LPARAM
	Pt          POINT
	VkDirection uint32
}

type LVHITTESTINFO struct {
	Pt       POINT // Coordinates relative to list view.
	Flags    co.LVHT
	IItem    int32 // -1 if no item.
	ISubItem int32
	IGroup   int32
}

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

type NMCUSTOMDRAW struct {
	Hdr         NMHDR
	DwDrawStage co.CDDS
	Hdc         HDC
	Rc          RECT
	DwItemSpec  uintptr
	UItemState  co.CDIS
	LItemlParam LPARAM
}

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

type NMLVCACHEHINT struct {
	Hdr   NMHDR
	IFrom int32
	ITo   int32
}

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

type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

type NMLVEMPTYMARKUP struct {
	Hdr      NMHDR
	DwFlags  co.EMF
	SzMarkup [2048 + 32 + 3]uint16 // L_MAX_URL_LENGTH
}

type NMLVFINDITEM struct {
	Hdr    NMHDR
	IStart int32
	Lvfi   LVFINDINFO
}

type NMLVGETINFOTIP struct {
	Hdr        NMHDR
	DwFlags    co.LVGIT
	PszText    uintptr // LPWSTR
	CchTextMax int32
	IItem      int32
	ISubItem   int32
	LParam     LPARAM
}

type NMLVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

type NMLVLINK struct {
	Hdr      NMHDR
	Link     LITEM
	IItem    int32
	ISubItem int32
}

type NMLVODSTATECHANGE struct {
	Hdr       NMHDR
	IFrom     int32
	ITo       int32
	UNewState co.LVIS
	UOldState co.LVIS
}

type NMLVSCROLL struct {
	Hdr NMHDR
	Dx  int32
	Dy  int32
}

type NMMOUSE struct {
	Hdr        NMHDR
	DwItemSpec uintptr
	DwItemData uintptr
	Pt         POINT
	DwHitInfo  LPARAM
}

type NMTREEVIEW struct {
	Hdr     NMHDR
	Action  uint32 // co.TVE | co.TVC
	ItemOld TVITEM
	ItemNew TVITEM
	PtDrag  POINT
}

type NMTVASYNCDRAW struct {
	Hdr            NMHDR
	Pimldp         uintptr  // LPIMAGELISTDRAWPARAMS
	Hr             co.ERROR // HRESULT
	Hitem          HTREEITEM
	LParam         LPARAM
	DwRetFlags     co.ADRF
	IRetImageIndex int32
}

type NMTVCUSTOMDRAW struct {
	Nmcd      NMCUSTOMDRAW
	ClrText   COLORREF
	ClrTextBk COLORREF
	ILevel    int32
}

type NMTVDISPINFO struct {
	Hdr  NMHDR
	Item TVITEM
}

type NMTVGETINFOTIP struct {
	Hdr        NMHDR
	PszText    uintptr // LPWSTR
	CchTextMax int32
	HItem      HTREEITEM
	LParam     LPARAM
}

type NMTVITEMCHANGE struct {
	Hdr       NMHDR
	UChanged  co.TVIF
	HItem     HTREEITEM
	UStateNew co.TVIS
	UStateOld co.TVIS
	LParam    LPARAM
}

type NMTVKEYDOWN struct {
	Hdr   NMHDR
	WVKey co.VK
	Flags uint32
}

type TVINSERTSTRUCT struct {
	HParent      HTREEITEM
	HInsertAfter HTREEITEM
	Itemex       TVITEMEX
}

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
