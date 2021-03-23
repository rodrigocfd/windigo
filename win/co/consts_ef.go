package co

// Edit control messages.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-messages
const (
	EM_GETSEL              WM = 0x00b0
	EM_SETSEL              WM = 0x00b1
	EM_GETRECT             WM = 0x00b2
	EM_SETRECT             WM = 0x00b3
	EM_SETRECTNP           WM = 0x00b4
	EM_SCROLL              WM = 0x00b5
	EM_LINESCROLL          WM = 0x00b6
	EM_SCROLLCARET         WM = 0x00b7
	EM_GETMODIFY           WM = 0x00b8
	EM_SETMODIFY           WM = 0x00b9
	EM_GETLINECOUNT        WM = 0x00ba
	EM_LINEINDEX           WM = 0x00bb
	EM_SETHANDLE           WM = 0x00bc
	EM_GETHANDLE           WM = 0x00bd
	EM_GETTHUMB            WM = 0x00be
	EM_LINELENGTH          WM = 0x00c1
	EM_REPLACESEL          WM = 0x00c2
	EM_GETLINE             WM = 0x00c4
	EM_LIMITTEXT           WM = 0x00c5
	EM_CANUNDO             WM = 0x00c6
	EM_UNDO                WM = 0x00c7
	EM_FMTLINES            WM = 0x00c8
	EM_LINEFROMCHAR        WM = 0x00c9
	EM_SETTABSTOPS         WM = 0x00cb
	EM_SETPASSWORDCHAR     WM = 0x00cc
	EM_EMPTYUNDOBUFFER     WM = 0x00cd
	EM_GETFIRSTVISIBLELINE WM = 0x00ce
	EM_SETREADONLY         WM = 0x00cf
	EM_SETWORDBREAKPROC    WM = 0x00d0
	EM_GETWORDBREAKPROC    WM = 0x00d1
	EM_GETPASSWORDCHAR     WM = 0x00d2
	EM_SETMARGINS          WM = 0x00d3
	EM_GETMARGINS          WM = 0x00d4
	EM_SETLIMITTEXT        WM = EM_LIMITTEXT
	EM_GETLIMITTEXT        WM = 0x00d5
	EM_POSFROMCHAR         WM = 0x00d6
	EM_CHARFROMPOS         WM = 0x00d7
	EM_SETIMESTATUS        WM = 0x00d8
	EM_GETIMESTATUS        WM = 0x00d9
)

// NMLVEMPTYMARKUP dwFlags.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmlvemptymarkup
type EMF uint32

const (
	EMF_NULL     EMF = 0x00000000
	EMF_CENTERED EMF = 0x00000001
)

// Edit control notifications, sent via WM_COMMAND.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-edit-control-reference-notifications
const (
	EN_SETFOCUS     CMD = 0x0100
	EN_KILLFOCUS    CMD = 0x0200
	EN_CHANGE       CMD = 0x0300
	EN_UPDATE       CMD = 0x0400
	EN_ERRSPACE     CMD = 0x0500
	EN_MAXTEXT      CMD = 0x0501
	EN_HSCROLL      CMD = 0x0601
	EN_VSCROLL      CMD = 0x0602
	EN_ALIGN_LTR_EC CMD = 0x0700
	EN_ALIGN_RTL_EC CMD = 0x0701
	EN_BEFORE_PASTE CMD = 0x0800
	EN_AFTER_PASTE  CMD = 0x0801
)

// WM_ENDSESSION event.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
type ENDSESSION uint32

const (
	ENDSESSION_RESTARTORSHUTDOWN ENDSESSION = 0
	ENDSESSION_CLOSEAPP          ENDSESSION = 0x00000001
	ENDSESSION_CRITICAL          ENDSESSION = 0x40000000
	ENDSESSION_LOGOFF            ENDSESSION = 0x80000000
)

// Edit control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/edit-control-styles
type ES WS

const (
	ES_LEFT        ES = 0x0000
	ES_CENTER      ES = 0x0001
	ES_RIGHT       ES = 0x0002
	ES_MULTILINE   ES = 0x0004
	ES_UPPERCASE   ES = 0x0008
	ES_LOWERCASE   ES = 0x0010
	ES_PASSWORD    ES = 0x0020
	ES_AUTOVSCROLL ES = 0x0040
	ES_AUTOHSCROLL ES = 0x0080
	ES_NOHIDESEL   ES = 0x0100
	ES_OEMCONVERT  ES = 0x0400
	ES_READONLY    ES = 0x0800
	ES_WANTRETURN  ES = 0x1000
	ES_NUMBER      ES = 0x2000
)

// WM_APPCOMMAND input event.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
type FAPPCOMMAND uint32

const (
	FAPPCOMMAND_MOUSE FAPPCOMMAND = 0x8000
	FAPPCOMMAND_KEY   FAPPCOMMAND = 0
	FAPPCOMMAND_OEM   FAPPCOMMAND = 0x1000
)

// CreateFile() dwFlagsAndAttributes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type FILE_ATTRIBUTE uint32

const (
	FILE_ATTRIBUTE_INVALID               FILE_ATTRIBUTE = 0xffffffff // -1
	FILE_ATTRIBUTE_READONLY              FILE_ATTRIBUTE = 0x00000001
	FILE_ATTRIBUTE_HIDDEN                FILE_ATTRIBUTE = 0x00000002
	FILE_ATTRIBUTE_SYSTEM                FILE_ATTRIBUTE = 0x00000004
	FILE_ATTRIBUTE_DIRECTORY             FILE_ATTRIBUTE = 0x00000010
	FILE_ATTRIBUTE_ARCHIVE               FILE_ATTRIBUTE = 0x00000020
	FILE_ATTRIBUTE_DEVICE                FILE_ATTRIBUTE = 0x00000040
	FILE_ATTRIBUTE_NORMAL                FILE_ATTRIBUTE = 0x00000080
	FILE_ATTRIBUTE_TEMPORARY             FILE_ATTRIBUTE = 0x00000100
	FILE_ATTRIBUTE_SPARSE_FILE           FILE_ATTRIBUTE = 0x00000200
	FILE_ATTRIBUTE_REPARSE_POINT         FILE_ATTRIBUTE = 0x00000400
	FILE_ATTRIBUTE_COMPRESSED            FILE_ATTRIBUTE = 0x00000800
	FILE_ATTRIBUTE_OFFLINE               FILE_ATTRIBUTE = 0x00001000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   FILE_ATTRIBUTE = 0x00002000
	FILE_ATTRIBUTE_ENCRYPTED             FILE_ATTRIBUTE = 0x00004000
	FILE_ATTRIBUTE_INTEGRITY_STREAM      FILE_ATTRIBUTE = 0x00008000
	FILE_ATTRIBUTE_VIRTUAL               FILE_ATTRIBUTE = 0x00010000
	FILE_ATTRIBUTE_NO_SCRUB_DATA         FILE_ATTRIBUTE = 0x00020000
	FILE_ATTRIBUTE_EA                    FILE_ATTRIBUTE = 0x00040000
	FILE_ATTRIBUTE_PINNED                FILE_ATTRIBUTE = 0x00080000
	FILE_ATTRIBUTE_UNPINNED              FILE_ATTRIBUTE = 0x00100000
	FILE_ATTRIBUTE_RECALL_ON_OPEN        FILE_ATTRIBUTE = 0x00040000
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS FILE_ATTRIBUTE = 0x00400000
)

// CreateFile() dwFlagsAndAttributes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type FILE_FLAG uint32

const (
	FILE_FLAG_NONE                  FILE_FLAG = 0
	FILE_FLAG_WRITE_THROUGH         FILE_FLAG = 0x80000000
	FILE_FLAG_OVERLAPPED            FILE_FLAG = 0x40000000
	FILE_FLAG_NO_BUFFERING          FILE_FLAG = 0x20000000
	FILE_FLAG_RANDOM_ACCESS         FILE_FLAG = 0x10000000
	FILE_FLAG_SEQUENTIAL_SCAN       FILE_FLAG = 0x08000000
	FILE_FLAG_DELETE_ON_CLOSE       FILE_FLAG = 0x04000000
	FILE_FLAG_BACKUP_SEMANTICS      FILE_FLAG = 0x02000000
	FILE_FLAG_POSIX_SEMANTICS       FILE_FLAG = 0x01000000
	FILE_FLAG_SESSION_AWARE         FILE_FLAG = 0x00800000
	FILE_FLAG_OPEN_REPARSE_POINT    FILE_FLAG = 0x00200000
	FILE_FLAG_OPEN_NO_RECALL        FILE_FLAG = 0x00100000
	FILE_FLAG_FIRST_PIPE_INSTANCE   FILE_FLAG = 0x00080000
	FILE_FLAG_OPEN_REQUIRING_OPLOCK FILE_FLAG = 0x00040000
)

// SetFilePointerEx() dwMoveMethod. Originally with FILE prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
type FILE_FROM uint32

const (
	FILE_FROM_BEGIN   FILE_FROM = 0
	FILE_FROM_CURRENT FILE_FROM = 1
	FILE_FROM_END     FILE_FROM = 2
)

// MapViewOfFile() dwDesiredAccess.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
type FILE_MAP uint32

const (
	_SECTION_QUERY                FILE_MAP = 0x0001
	_SECTION_MAP_WRITE            FILE_MAP = 0x0002
	_SECTION_MAP_READ             FILE_MAP = 0x0004
	_SECTION_MAP_EXECUTE          FILE_MAP = 0x0008
	_SECTION_EXTEND_SIZE          FILE_MAP = 0x0010
	_SECTION_MAP_EXECUTE_EXPLICIT FILE_MAP = 0x0020
	_SECTION_ALL_ACCESS           FILE_MAP = FILE_MAP(STANDARD_RIGHTS_REQUIRED) | _SECTION_QUERY | _SECTION_MAP_WRITE | _SECTION_MAP_READ | _SECTION_MAP_EXECUTE | _SECTION_EXTEND_SIZE

	FILE_MAP_WRITE           FILE_MAP = _SECTION_MAP_WRITE
	FILE_MAP_READ            FILE_MAP = _SECTION_MAP_READ
	FILE_MAP_ALL_ACCESS      FILE_MAP = _SECTION_ALL_ACCESS
	FILE_MAP_EXECUTE         FILE_MAP = _SECTION_MAP_EXECUTE_EXPLICIT
	FILE_MAP_COPY            FILE_MAP = 0x00000001
	FILE_MAP_RESERVE         FILE_MAP = 0x80000000
	FILE_MAP_TARGETS_INVALID FILE_MAP = 0x40000000
	FILE_MAP_LARGE_PAGES     FILE_MAP = 0x20000000
)

// CreateFile() dwShareMode.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type FILE_SHARE uint32

const (
	FILE_SHARE_NONE   FILE_SHARE = 0
	FILE_SHARE_READ   FILE_SHARE = 0x00000001
	FILE_SHARE_WRITE  FILE_SHARE = 0x00000002
	FILE_SHARE_DELETE FILE_SHARE = 0x00000004
)

// LOGFONT lfWeight.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-logfontw
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
