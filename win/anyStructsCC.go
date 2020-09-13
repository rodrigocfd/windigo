/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

import (
	"windigo/co"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-comboboxexitemw
	COMBOBOXEXITEM struct {
		Mask           co.CBEIF
		IItem          uintptr // INT_PTR
		PszText        *uint16
		CchTextMax     int32
		IImage         int32
		ISelectedImage int32
		IOverlay       int32
		IIndent        int32
		LParam         LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-hditemw
	HDITEM struct {
		Mask       co.HDI
		Cxy        int32
		PszText    *uint16
		Hbm        HBITMAP
		CchTextMax int32
		Fmt        co.HDF
		LParam     LPARAM
		IImage     int32
		IOrder     int32
		Type       co.HDFT
		PvFilter   uintptr // void*
		State      co.HDIS
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imageinfo
	IMAGEINFO struct {
		HbmImage HBITMAP
		HbmMask  HBITMAP
		Unused1  int32
		Unused2  int32
		RcImage  RECT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
	IMAGELISTDRAWPARAMS struct {
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
	LITEM struct {
		Mask      co.LIF
		ILink     int32
		State     co.LIS
		StateMask co.LIS
		SzID      [_MAX_LINKID_TEXT]uint16
		SzUrl     [_L_MAX_URL_LENGTH]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvcolumnw
	LVCOLUMN struct {
		Mask       co.LVCF
		Fmt        int32
		Cx         int32
		PszText    *uint16
		CchTextMax int32
		ISubItem   int32
		IImage     int32
		IOrder     int32
		CxMin      int32
		CxDefault  int32
		CxIdeal    int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvfindinfow
	LVFINDINFO struct {
		Flags       co.LVFI
		Psz         *uint16
		LParam      LPARAM
		Pt          POINT
		VkDirection uint32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvhittestinfo
	LVHITTESTINFO struct {
		Pt       POINT // Coordinates relative to list view.
		Flags    co.LVHT
		IItem    int32 // -1 if no item.
		ISubItem int32
		IGroup   int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-lvitemw
	LVITEM struct {
		Mask       co.LVIF
		IItem      int32
		ISubItem   int32
		State      co.LVIS
		StateMask  co.LVIS
		PszText    *uint16
		CchTextMax int32
		IImage     int32
		LParam     LPARAM
		IIndent    int32
		IGroupId   int32
		CColumns   uint32
		PuColumns  *uint32
		PiColFmt   *int32
		IGroup     int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcbedragbeginw
	NMCBEDRAGBEGIN struct {
		Hdr     NMHDR
		IItemid int32
		SzText  [_CBEMAXSTRLEN]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcbeendeditw
	NMCBEENDEDIT struct {
		Hdr           NMHDR
		FChanged      int32 // BOOL
		INewSelection int32
		SzText        [_CBEMAXSTRLEN]uint16
		IWhy          co.CBENF
	}

	// https://docs.microsoft.com/pt-br/windows/win32/controls/cben-deleteitem
	NMCOMBOBOXEX struct {
		Hdr    NMHDR
		CeItem COMBOBOXEXITEM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
	NMCUSTOMDRAW struct {
		Hdr         NMHDR
		DwDrawStage co.CDDS
		Hdc         HDC
		Rc          RECT
		DwItemSpec  uintptr // DWORD_PTR
		UItemState  co.CDIS
		LItemlParam LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimechange
	NMDATETIMECHANGE struct {
		Nmhdr   NMHDR
		DwFlags co.GDT
		St      SYSTEMTIME
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatw
	NMDATETIMEFORMAT struct {
		Nmhdr      NMHDR
		PszFormat  *uint16
		St         SYSTEMTIME
		pszDisplay *uint16
		SzDisplay  [64]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimeformatqueryw
	NMDATETIMEFORMATQUERY struct {
		Nmhdr     NMHDR
		PszFormat *uint16
		SzMax     SIZE
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimestringw
	NMDATETIMESTRING struct {
		Nmhdr         NMHDR
		PszUserString *uint16
		St            SYSTEMTIME
		DwFlags       co.GDT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmdatetimewmkeydownw
	NMDATETIMEWMKEYDOWN struct {
		Nmhdr     NMHDR
		NVirtKey  int32
		PszFormat *uint16
		St        SYSTEMTIME
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmhdfilterbtnclick
	NMHDFILTERBTNCLICK struct {
		Hdr   NMHDR
		IItem int32
		Rc    RECT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmhddispinfow
	NMHDDISPINFO struct {
		Hdr        NMHDR
		IItem      int32
		Mask       co.HDI
		PszText    *uint16
		CchTextMax int32
		IImage     int32
		LParam     LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmheaderw
	NMHEADER struct {
		Hdr     NMHDR
		IItem   int32
		IButton int32
		Pitem   *HDITEM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/Commctrl/ns-commctrl-nmipaddress
	NMIPADDRESS struct {
		Hdr    NMHDR
		IField int32
		IValue int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmitemactivate
	NMITEMACTIVATE struct {
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

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlink
	NMLINK struct {
		Hdr  NMHDR
		Item LITEM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlistview
	NMLISTVIEW struct {
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
	NMLVCACHEHINT struct {
		Hdr   NMHDR
		IFrom int32
		ITo   int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvcustomdraw
	NMLVCUSTOMDRAW struct {
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
	NMLVDISPINFO struct {
		Hdr  NMHDR
		Item LVITEM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvemptymarkup
	NMLVEMPTYMARKUP struct {
		Hdr      NMHDR
		DwFlags  co.EMF
		SzMarkup [_L_MAX_URL_LENGTH]uint16
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvfinditemw
	NMLVFINDITEM struct {
		Hdr    NMHDR
		IStart int32
		Lvfi   LVFINDINFO
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvgetinfotipw
	NMLVGETINFOTIP struct {
		Hdr        NMHDR
		DwFlags    co.LVGIT
		PszText    *uint16
		CchTextMax int32
		IItem      int32
		ISubItem   int32
		LParam     LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvkeydown
	NMLVKEYDOWN struct {
		Hdr   NMHDR
		WVKey co.VK
		Flags uint32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvlink
	NMLVLINK struct {
		Hdr      NMHDR
		Link     LITEM
		IItem    int32
		ISubItem int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvodstatechange
	NMLVODSTATECHANGE struct {
		Hdr       NMHDR
		IFrom     int32
		ITo       int32
		UNewState co.LVIS
		UOldState co.LVIS
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvscroll
	NMLVSCROLL struct {
		Hdr NMHDR
		Dx  int32
		Dy  int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmmouse
	NMMOUSE struct {
		Hdr        NMHDR
		DwItemSpec uintptr // DWORD_PTR
		DwItemData uintptr // DWORD_PTR
		Pt         POINT
		DwHitInfo  LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtreevieww
	NMTREEVIEW struct {
		Hdr     NMHDR
		Action  uint32 // co.TVE | co.TVC
		ItemOld TVITEM
		ItemNew TVITEM
		PtDrag  POINT
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
	NMTVASYNCDRAW struct {
		Hdr            NMHDR
		Pimldp         *IMAGELISTDRAWPARAMS
		Hr             co.ERROR // HRESULT
		Hitem          HTREEITEM
		LParam         LPARAM
		DwRetFlags     co.ADRF
		IRetImageIndex int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvcustomdraw
	NMTVCUSTOMDRAW struct {
		Nmcd      NMCUSTOMDRAW
		ClrText   COLORREF
		ClrTextBk COLORREF
		ILevel    int32
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfow
	NMTVDISPINFO struct {
		Hdr  NMHDR
		Item TVITEM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvgetinfotipw
	NMTVGETINFOTIP struct {
		Hdr        NMHDR
		PszText    *uint16
		CchTextMax int32
		HItem      HTREEITEM
		LParam     LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvitemchange
	NMTVITEMCHANGE struct {
		Hdr       NMHDR
		UChanged  co.TVIF
		HItem     HTREEITEM
		UStateNew co.TVIS
		UStateOld co.TVIS
		LParam    LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvkeydown
	NMTVKEYDOWN struct {
		Hdr   NMHDR
		WVKey co.VK
		Flags uint32
	}

	// https://www.google.com/search?client=firefox-b-d&q=TVINSERTSTRUCTW
	TVINSERTSTRUCT struct {
		HParent      HTREEITEM
		HInsertAfter HTREEITEM
		Itemex       TVITEMEX
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemw
	TVITEM struct {
		Mask           co.TVIF
		HItem          HTREEITEM
		State          co.TVIS
		StateMask      co.TVIS
		PszText        *uint16
		CchTextMax     int32
		IImage         int32
		ISelectedImage int32
		CChildren      co.TVI_CHILDREN
		LParam         LPARAM
	}

	// https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tvitemexw
	TVITEMEX struct {
		Mask           co.TVIF
		HItem          HTREEITEM
		State          co.TVIS
		StateMask      co.TVIS
		PszText        *uint16
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
)
