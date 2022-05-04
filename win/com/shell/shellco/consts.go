package shellco

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-advanceslideshow
type DSD uint32

const (
	DSD_FORWARD  DSD = 0
	DSD_BACKWARD DSD = 1
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getslideshowoptions
type DSO uint32

const (
	DSO_DISABLED      DSO = 0
	DSO_SHUFFLEIMAGES DSO = 0x01
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-idesktopwallpaper-getstatus
type DSS uint32

const (
	DSS_ENABLED                    DSS = 0x01
	DSS_SLIDESHOW                  DSS = 0x02
	DSS_DISABLED_BY_REMOTE_SESSION DSS = 0x04
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-desktop_wallpaper_position
type DWPOS uint32

const (
	DWPOS_CENTER  DWPOS = 0
	DWPOS_TILE    DWPOS = 1
	DWPOS_STRETCH DWPOS = 2
	DWPOS_FIT     DWPOS = 3
	DWPOS_FILL    DWPOS = 4
	DWPOS_SPAN    DWPOS = 5
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/com/dropeffect-constants
type DROPEFFECT uint32

const (
	DROPEFFECT_NONE   DROPEFFECT = 0
	DROPEFFECT_COPY   DROPEFFECT = 1
	DROPEFFECT_MOVE   DROPEFFECT = 2
	DROPEFFECT_LINK   DROPEFFECT = 4
	DROPEFFECT_SCROLL DROPEFFECT = 0x8000_0000
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_fileopendialogoptions
type FOS uint32

const (
	FOS_OVERWRITEPROMPT          FOS = 0x2
	FOS_STRICTFILETYPES          FOS = 0x4
	FOS_NOCHANGEDIR              FOS = 0x8
	FOS_PICKFOLDERS              FOS = 0x20
	FOS_FORCEFILESYSTEM          FOS = 0x40
	FOS_ALLNONSTORAGEITEMS       FOS = 0x80
	FOS_NOVALIDATE               FOS = 0x100
	FOS_ALLOWMULTISELECT         FOS = 0x200
	FOS_PATHMUSTEXIST            FOS = 0x800
	FOS_FILEMUSTEXIST            FOS = 0x1000
	FOS_CREATEPROMPT             FOS = 0x2000
	FOS_SHAREAWARE               FOS = 0x4000
	FOS_NOREADONLYRETURN         FOS = 0x8000
	FOS_NOTESTFILECREATE         FOS = 0x1_0000
	FOS_HIDEMRUPLACES            FOS = 0x2_0000
	FOS_HIDEPINNEDPLACES         FOS = 0x4_0000
	FOS_NODEREFERENCELINKS       FOS = 0x10_0000
	FOS_OKBUTTONNEEDSINTERACTION FOS = 0x20_0000
	FOS_DONTADDTORECENT          FOS = 0x200_0000
	FOS_FORCESHOWHIDDEN          FOS = 0x1000_0000
	FOS_DEFAULTNOMINIMODE        FOS = 0x2000_0000
	FOS_FORCEPREVIEWPANEON       FOS = 0x4000_0000
	FOS_SUPPORTSTREAMABLEITEMS   FOS = 0x8000_0000
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_sichintf
type SICHINT uint32

const (
	SICHINT_DISPLAY                       SICHINT = 0
	SICHINT_ALLFIELDS                     SICHINT = 0x8000_0000
	SICHINT_CANONICAL                     SICHINT = 0x1000_0000
	SICHINT_TEST_FILESYSPATH_IF_NOT_EQUAL SICHINT = 0x2000_0000
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-sigdn
type SIGDN uint32

const (
	SIGDN_NORMALDISPLAY               SIGDN = 0
	SIGDN_PARENTRELATIVEPARSING       SIGDN = 0x8001_8001
	SIGDN_DESKTOPABSOLUTEPARSING      SIGDN = 0x8002_8000
	SIGDN_PARENTRELATIVEEDITING       SIGDN = 0x8003_1001
	SIGDN_DESKTOPABSOLUTEEDITING      SIGDN = 0x8004_c000
	SIGDN_FILESYSPATH                 SIGDN = 0x8005_8000
	SIGDN_URL                         SIGDN = 0x8006_8000
	SIGDN_PARENTRELATIVEFORADDRESSBAR SIGDN = 0x8007_c001
	SIGDN_PARENTRELATIVE              SIGDN = 0x8008_0001
	SIGDN_PARENTRELATIVEFORUI         SIGDN = 0x8009_4001
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-getpath
type SLGP uint32

const (
	SLGP_SHORTPATH        SLGP = 0x1
	SLGP_UNCPRIORITY      SLGP = 0x2
	SLGP_RAWPATH          SLGP = 0x4
	SLGP_RELATIVEPRIORITY SLGP = 0x8
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-ishelllinkw-resolve
type SLR uint32

const (
	SLR_NONE                      SLR = 0
	SLR_NO_UI                     SLR = 0x1
	SLR_ANY_MATCH                 SLR = 0x2
	SLR_UPDATE                    SLR = 0x4
	SLR_NOUPDATE                  SLR = 0x8
	SLR_NOSEARCH                  SLR = 0x10
	SLR_NOTRACK                   SLR = 0x20
	SLR_NOLINKINFO                SLR = 0x40
	SLR_INVOKE_MSI                SLR = 0x80
	SLR_NO_UI_WITH_MSG_PUMP       SLR = 0x101
	SLR_OFFER_DELETE_WITHOUT_FILE SLR = 0x200
	SLR_KNOWNFOLDER               SLR = 0x400
	SLR_MACHINE_IN_LOCAL_TARGET   SLR = 0x800
	SLR_UPDATE_MACHINE_AND_SID    SLR = 0x1000
	SLR_NO_OBJECT_ID              SLR = 0x2000
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-stpflag
type STPFLAG uint32

const (
	STPFLAG_NONE                      STPFLAG = 0
	STPFLAG_USEAPPTHUMBNAILALWAYS     STPFLAG = 0x1
	STPFLAG_USEAPPTHUMBNAILWHENACTIVE STPFLAG = 0x2
	STPFLAG_USEAPPPEEKALWAYS          STPFLAG = 0x4
	STPFLAG_USEAPPPEEKWHENACTIVE      STPFLAG = 0x8
)

// ITaskbarList3::SetProgressState() tbpFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
type TBPF uint32

const (
	// Stops displaying progress and returns the button to its normal state.
	// Call the method with this flag to dismiss the progress bar when the
	// operation is complete or canceled.
	TBPF_NOPROGRESS TBPF = 0
	// The progress indicator does not grow in size, but cycles repeatedly along
	// the length of the taskbar button. This indicates activity without
	// specifying what proportion of the progress is complete. Progress is
	// taking place, but there is no prediction as to how long the operation
	// will take.
	TBPF_INDETERMINATE TBPF = 0x1
	// The progress indicator grows in size from left to right in proportion to
	// the estimated amount of the operation completed. This is a determinate
	// progress indicator; a prediction is being made as to the duration of the
	// operation.
	TBPF_NORMAL TBPF = 0x2
	// The progress indicator turns red to show that an error has occurred in
	// one of the windows that is broadcasting progress. This is a determinate
	// state. If the progress indicator is in the indeterminate state, it
	// switches to a red determinate display of a generic percentage not
	// indicative of actual progress.
	TBPF_ERROR TBPF = 0x4
	// The progress indicator turns yellow to show that progress is currently
	// stopped in one of the windows but can be resumed by the user. No error
	// condition exists and nothing is preventing the progress from continuing.
	// This is a determinate state. If the progress indicator is in the
	// indeterminate state, it switches to a yellow determinate display of a
	// generic percentage not indicative of actual progress.
	TBPF_PAUSED TBPF = 0x8
)

// THUMBBUTTON dwMask.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-thumbbuttonmask
type THB uint32

const (
	THB_BITMAP  THB = 0x1
	THB_ICON    THB = 0x2
	THB_TOOLTIP THB = 0x4
	THB_FLAGS   THB = 0x8
)

// THUMBBUTTON dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-thumbbuttonflags
type THBF uint32

const (
	THBF_ENABLED        THBF = 0
	THBF_DISABLED       THBF = 0x1
	THBF_DISMISSONCLICK THBF = 0x2
	THBF_NOBACKGROUND   THBF = 0x4
	THBF_HIDDEN         THBF = 0x8
	THBF_NONINTERACTIVE THBF = 0x10
)
