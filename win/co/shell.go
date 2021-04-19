package co

const (
	CLSID_FileOpenDialog CLSID = "dc1c5a9c-e88a-4dde-a5a1-60f82a20aef7"
	CLSID_FileSaveDialog CLSID = "c0b4e2f3-ba21-4773-8dba-335ec946eb8b"
	CLSID_TaskbarList    CLSID = "56fdf344-fd6d-11d0-958a-006097c9a090"

	IID_IFileOpenDialog IID = "d57c7288-d4ad-4768-be02-9d969532d960"
	IID_IFileSaveDialog IID = "84bccd23-5fde-4cdb-aea4-af64b83d78ab"
	IID_IShellItem      IID = "43826d1e-e718-42ee-bc55-a1e261c37bfe"
	IID_ITaskbarList    IID = "56fdf342-fd6d-11d0-958a-006097c9a090"
	IID_ITaskbarList2   IID = "602d4995-b13a-429b-a66e-1935e44f4317"
	IID_ITaskbarList3   IID = "ea1afb91-9e28-4b86-90e9-9e9f8a5eefaf"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/com/dropeffect-constants
type DROPEFFECT uint32

const (
	DROPEFFECT_NONE   DROPEFFECT = 0
	DROPEFFECT_COPY   DROPEFFECT = 1
	DROPEFFECT_MOVE   DROPEFFECT = 2
	DROPEFFECT_LINK   DROPEFFECT = 4
	DROPEFFECT_SCROLL DROPEFFECT = 0x80000000
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
	FOS_NOTESTFILECREATE         FOS = 0x10000
	FOS_HIDEMRUPLACES            FOS = 0x20000
	FOS_HIDEPINNEDPLACES         FOS = 0x40000
	FOS_NODEREFERENCELINKS       FOS = 0x100000
	FOS_OKBUTTONNEEDSINTERACTION FOS = 0x200000
	FOS_DONTADDTORECENT          FOS = 0x2000000
	FOS_FORCESHOWHIDDEN          FOS = 0x10000000
	FOS_DEFAULTNOMINIMODE        FOS = 0x20000000
	FOS_FORCEPREVIEWPANEON       FOS = 0x40000000
	FOS_SUPPORTSTREAMABLEITEMS   FOS = 0x80000000
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-_sichintf
type SICHINT uint32

const (
	SICHINT_DISPLAY                       SICHINT = 0
	SICHINT_ALLFIELDS                     SICHINT = 0x80000000
	SICHINT_CANONICAL                     SICHINT = 0x10000000
	SICHINT_TEST_FILESYSPATH_IF_NOT_EQUAL SICHINT = 0x20000000
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/ne-shobjidl_core-sigdn
type SIGDN uint32

const (
	SIGDN_NORMALDISPLAY               SIGDN = 0
	SIGDN_PARENTRELATIVEPARSING       SIGDN = 0x80018001
	SIGDN_DESKTOPABSOLUTEPARSING      SIGDN = 0x80028000
	SIGDN_PARENTRELATIVEEDITING       SIGDN = 0x80031001
	SIGDN_DESKTOPABSOLUTEEDITING      SIGDN = 0x8004c000
	SIGDN_FILESYSPATH                 SIGDN = 0x80058000
	SIGDN_URL                         SIGDN = 0x80068000
	SIGDN_PARENTRELATIVEFORADDRESSBAR SIGDN = 0x8007c001
	SIGDN_PARENTRELATIVE              SIGDN = 0x80080001
	SIGDN_PARENTRELATIVEFORUI         SIGDN = 0x80094001
)

// ITaskbarList3::SetProgressState() tbpFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
type TBPF uint32

const (
	TBPF_NOPROGRESS    TBPF = 0
	TBPF_INDETERMINATE TBPF = 0x1
	TBPF_NORMAL        TBPF = 0x2
	TBPF_ERROR         TBPF = 0x4
	TBPF_PAUSED        TBPF = 0x8
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
