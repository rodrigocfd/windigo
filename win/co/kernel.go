//go:build windows

package co

// [CONSOLE_READCONSOLE_CONTROL] DwControlKeyState.
//
// [CONSOLE_READCONSOLE_CONTROL]: https://learn.microsoft.com/en-us/windows/console/console-readconsole-control
type CKS uint32

const (
	CAPSLOCK_ON        CKS = 0x0080
	ENHANCED_KEY       CKS = 0x0100
	LEFT_ALT_PRESSED   CKS = 0x0002
	LEFT_CTRL_PRESSED  CKS = 0x0008
	NUMLOCK_ON         CKS = 0x0020
	RIGHT_ALT_PRESSED  CKS = 0x0001
	RIGHT_CTRL_PRESSED CKS = 0x0004
	SCROLLLOCK_ON      CKS = 0x0040
	SHIFT_PRESSED      CKS = 0x0010
)

// [SetConsoleDisplayMode] mode.
//
// [SetConsoleDisplayMode]: https://learn.microsoft.com/en-us/windows/console/setconsoledisplaymode
type CONSOLE uint32

const (
	CONSOLE_FULLSCREEN_MODE CONSOLE = 1
	CONSOLE_WINDOWED_MODE   CONSOLE = 2
)

// [Code page] identifiers.
//
// [Code page]: https://learn.microsoft.com/en-us/windows/win32/intl/code-page-identifiers
type CP uint16

const (
	CP_ACP        CP = 0  // The system default Windows ANSI code page.
	CP_OEMCP      CP = 1  // The current system OEM code page.
	CP_MACCP      CP = 2  // The current system Macintosh code page.
	CP_THREAD_ACP CP = 3  // The Windows ANSI code page for the current thread.
	CP_SYMBOL     CP = 42 // Symbol code page (42).

	CP_IBM1026      CP = 1026  // IBM EBCDIC Turkish (Latin 5).
	CP_IBM01047     CP = 1047  // IBM EBCDIC Latin 1/Open System.
	CP_IBM01140     CP = 1140  // IBM EBCDIC US-Canada (037 + Euro symbol); IBM EBCDIC (US-Canada-Euro).
	CP_IBM01141     CP = 1141  // IBM EBCDIC Germany (20273 + Euro symbol); IBM EBCDIC (Germany-Euro).
	CP_IBM01142     CP = 1142  // IBM EBCDIC Denmark-Norway (20277 + Euro symbol); IBM EBCDIC (Denmark-Norway-Euro).
	CP_IBM01143     CP = 1143  // IBM EBCDIC Finland-Sweden (20278 + Euro symbol); IBM EBCDIC (Finland-Sweden-Euro).
	CP_IBM01144     CP = 1144  // IBM EBCDIC Italy (20280 + Euro symbol); IBM EBCDIC (Italy-Euro).
	CP_IBM01145     CP = 1145  // IBM EBCDIC Latin America-Spain (20284 + Euro symbol); IBM EBCDIC (Spain-Euro).
	CP_IBM01146     CP = 1146  // IBM EBCDIC United Kingdom (20285 + Euro symbol); IBM EBCDIC (UK-Euro).
	CP_IBM01147     CP = 1147  // IBM EBCDIC France (20297 + Euro symbol); IBM EBCDIC (France-Euro).
	CP_IBM01148     CP = 1148  // IBM EBCDIC International (500 + Euro symbol); IBM EBCDIC (International-Euro).
	CP_IBM01149     CP = 1149  // IBM EBCDIC Icelandic (20871 + Euro symbol); IBM EBCDIC (Icelandic-Euro).
	CP_UTF16        CP = 1200  // Unicode UTF-16, little endian byte order (BMP of ISO 10646).
	CP_UNICODE_FFFE CP = 1201  // Unicode UTF-16, big endian byte order.
	CP_WINDOWS_1250 CP = 1250  // ANSI Central European; Central European (Windows).
	CP_WINDOWS_1251 CP = 1251  // ANSI Cyrillic; Cyrillic (Windows).
	CP_WINDOWS_1252 CP = 1252  // ANSI Latin 1; Western European (Windows).
	CP_WINDOWS_1253 CP = 1253  // ANSI Greek; Greek (Windows).
	CP_WINDOWS_1254 CP = 1254  // ANSI Turkish; Turkish (Windows).
	CP_WINDOWS_1255 CP = 1255  // ANSI Hebrew; Hebrew (Windows).
	CP_WINDOWS_1256 CP = 1256  // ANSI Arabic; Arabic (Windows).
	CP_WINDOWS_1257 CP = 1257  // ANSI Baltic; Baltic (Windows).
	CP_WINDOWS_1258 CP = 1258  // ANSI/OEM Vietnamese; Vietnamese (Windows).
	CP_JOHAB        CP = 1361  // Korean (Johab).
	CP_MACINTOSH    CP = 10000 // MAC Roman; Western European (Mac).

	CP_UTF7 CP = 65000 // Unicode (UTF-7).
	CP_UTF8 CP = 65001 // Unicode (UTF-8).
)

// [CreateProcess] dwCreationFlags.
//
// [CreateProcess]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessw
type CREATE uint32

const (
	CREATE_NONE CREATE = 0

	CREATE_BREAKAWAY_FROM_JOB        CREATE = 0x0100_0000
	CREATE_DEFAULT_ERROR_MODE        CREATE = 0x0400_0000
	CREATE_NEW_CONSOLE               CREATE = 0x0000_0010
	CREATE_NEW_PROCESS_GROUP         CREATE = 0x0000_0200
	CREATE_NO_WINDOW                 CREATE = 0x0800_0000
	CREATE_PROTECTED_PROCESS         CREATE = 0x0004_0000
	CREATE_PRESERVE_CODE_AUTHZ_LEVEL CREATE = 0x0200_0000
	CREATE_SECURE_PROCESS            CREATE = 0x0040_0000
	CREATE_SEPARATE_WOW_VDM          CREATE = 0x0000_0800
	CREATE_SHARED_WOW_VDM            CREATE = 0x0000_1000
	CREATE_SUSPENDED                 CREATE = 0x0000_0004
	CREATE_UNICODE_ENVIRONMENT       CREATE = 0x0000_0400

	CREATE_DEBUG_ONLY_THIS_PROCESS      CREATE = 0x0000_0002
	CREATE_DEBUG_PROCESS                CREATE = 0x0000_0001
	CREATE_DETACHED_PROCESS             CREATE = 0x0000_0008
	CREATE_EXTENDED_STARTUPINFO_PRESENT CREATE = 0x0008_0000
	CREATE_INHERIT_PARENT_AFFINITY      CREATE = 0x0001_0000
)

// [CreateFile] dwCreationDisposition. Originally without prefix.
//
// [CreateFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type DISPOSITION uint32

const (
	DISPOSITION_CREATE_ALWAYS     DISPOSITION = 2
	DISPOSITION_CREATE_NEW        DISPOSITION = 1
	DISPOSITION_OPEN_ALWAYS       DISPOSITION = 4
	DISPOSITION_OPEN_EXISTING     DISPOSITION = 3
	DISPOSITION_TRUNCATE_EXISTING DISPOSITION = 5
)

// [SetConsoleMode] mode.
//
// [SetConsoleMode]: https://learn.microsoft.com/en-us/windows/console/setconsolemode
type ENABLE uint32

const (
	ENABLE_ECHO_INPUT             ENABLE = 0x0004
	ENABLE_INSERT_MODE            ENABLE = 0x0020
	ENABLE_LINE_INPUT             ENABLE = 0x0002
	ENABLE_MOUSE_INPUT            ENABLE = 0x0010
	ENABLE_PROCESSED_INPUT        ENABLE = 0x0001
	ENABLE_QUICK_EDIT_MODE        ENABLE = 0x0040
	ENABLE_WINDOW_INPUT           ENABLE = 0x0008
	ENABLE_VIRTUAL_TERMINAL_INPUT ENABLE = 0x0200

	ENABLE_PROCESSED_OUTPUT            ENABLE = 0x0001
	ENABLE_WRAP_AT_EOL_OUTPUT          ENABLE = 0x0002
	ENABLE_VIRTUAL_TERMINAL_PROCESSING ENABLE = 0x0004
	ENABLE_DISABLE_NEWLINE_AUTO_RETURN ENABLE = 0x0008
	ENABLE_LVB_GRID_WORLDWIDE          ENABLE = 0x0010
)

// [WM_ENDSESSION] event.
//
// [WM_ENDSESSION]: https://learn.microsoft.com/en-us/windows/win32/shutdown/wm-endsession
type ENDSESSION uint32

const (
	ENDSESSION_RESTARTORSHUTDOWN ENDSESSION = 0
	ENDSESSION_CLOSEAPP          ENDSESSION = 0x0000_0001
	ENDSESSION_CRITICAL          ENDSESSION = 0x4000_0000
	ENDSESSION_LOGOFF            ENDSESSION = 0x8000_0000
)

// [CreateFile] dwFlagsAndAttributes.
//
// [CreateFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type FILE_ATTRIBUTE uint32

const (
	FILE_ATTRIBUTE_INVALID               FILE_ATTRIBUTE = 0xffff_ffff // -1
	FILE_ATTRIBUTE_READONLY              FILE_ATTRIBUTE = 0x0000_0001
	FILE_ATTRIBUTE_HIDDEN                FILE_ATTRIBUTE = 0x0000_0002
	FILE_ATTRIBUTE_SYSTEM                FILE_ATTRIBUTE = 0x0000_0004
	FILE_ATTRIBUTE_DIRECTORY             FILE_ATTRIBUTE = 0x0000_0010
	FILE_ATTRIBUTE_ARCHIVE               FILE_ATTRIBUTE = 0x0000_0020
	FILE_ATTRIBUTE_DEVICE                FILE_ATTRIBUTE = 0x0000_0040
	FILE_ATTRIBUTE_NORMAL                FILE_ATTRIBUTE = 0x0000_0080
	FILE_ATTRIBUTE_TEMPORARY             FILE_ATTRIBUTE = 0x0000_0100
	FILE_ATTRIBUTE_SPARSE_FILE           FILE_ATTRIBUTE = 0x0000_0200
	FILE_ATTRIBUTE_REPARSE_POINT         FILE_ATTRIBUTE = 0x0000_0400
	FILE_ATTRIBUTE_COMPRESSED            FILE_ATTRIBUTE = 0x0000_0800
	FILE_ATTRIBUTE_OFFLINE               FILE_ATTRIBUTE = 0x0000_1000
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED   FILE_ATTRIBUTE = 0x0000_2000
	FILE_ATTRIBUTE_ENCRYPTED             FILE_ATTRIBUTE = 0x0000_4000
	FILE_ATTRIBUTE_INTEGRITY_STREAM      FILE_ATTRIBUTE = 0x0000_8000
	FILE_ATTRIBUTE_VIRTUAL               FILE_ATTRIBUTE = 0x0001_0000
	FILE_ATTRIBUTE_NO_SCRUB_DATA         FILE_ATTRIBUTE = 0x0002_0000
	FILE_ATTRIBUTE_EA                    FILE_ATTRIBUTE = 0x0004_0000
	FILE_ATTRIBUTE_PINNED                FILE_ATTRIBUTE = 0x0008_0000
	FILE_ATTRIBUTE_UNPINNED              FILE_ATTRIBUTE = 0x0010_0000
	FILE_ATTRIBUTE_RECALL_ON_OPEN        FILE_ATTRIBUTE = 0x0004_0000
	FILE_ATTRIBUTE_RECALL_ON_DATA_ACCESS FILE_ATTRIBUTE = 0x0040_0000
)

// [CreateFile] dwFlagsAndAttributes.
//
// [CreateFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type FILE_FLAG uint32

const (
	FILE_FLAG_NONE                  FILE_FLAG = 0
	FILE_FLAG_WRITE_THROUGH         FILE_FLAG = 0x8000_0000
	FILE_FLAG_OVERLAPPED            FILE_FLAG = 0x4000_0000
	FILE_FLAG_NO_BUFFERING          FILE_FLAG = 0x2000_0000
	FILE_FLAG_RANDOM_ACCESS         FILE_FLAG = 0x1000_0000
	FILE_FLAG_SEQUENTIAL_SCAN       FILE_FLAG = 0x0800_0000
	FILE_FLAG_DELETE_ON_CLOSE       FILE_FLAG = 0x0400_0000
	FILE_FLAG_BACKUP_SEMANTICS      FILE_FLAG = 0x0200_0000
	FILE_FLAG_POSIX_SEMANTICS       FILE_FLAG = 0x0100_0000
	FILE_FLAG_SESSION_AWARE         FILE_FLAG = 0x0080_0000
	FILE_FLAG_OPEN_REPARSE_POINT    FILE_FLAG = 0x0020_0000
	FILE_FLAG_OPEN_NO_RECALL        FILE_FLAG = 0x0010_0000
	FILE_FLAG_FIRST_PIPE_INSTANCE   FILE_FLAG = 0x0008_0000
	FILE_FLAG_OPEN_REQUIRING_OPLOCK FILE_FLAG = 0x0004_0000
)

// [SetFilePointerEx] dwMoveMethod. Originally with FILE prefix.
//
// [SetFilePointerEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-setfilepointerex
type FILE_FROM uint32

const (
	FILE_FROM_BEGIN   FILE_FROM = 0
	FILE_FROM_CURRENT FILE_FROM = 1
	FILE_FROM_END     FILE_FROM = 2
)

// [MapViewOfFile] dwDesiredAccess.
//
// [MapViewOfFile]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-mapviewoffile
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
	FILE_MAP_COPY            FILE_MAP = 0x0000_0001
	FILE_MAP_RESERVE         FILE_MAP = 0x8000_0000
	FILE_MAP_TARGETS_INVALID FILE_MAP = 0x4000_0000
	FILE_MAP_LARGE_PAGES     FILE_MAP = 0x2000_0000
)

// FileOpen() and FileMappedOpen() desired access.
type FILE_OPEN uint8

const (
	FILE_OPEN_READ_EXISTING     FILE_OPEN = iota // Open an existing file for read only.
	FILE_OPEN_RW_EXISTING                        // Open an existing file for read and write.
	FILE_OPEN_RW_OPEN_OR_CREATE                  // Open a file or create if it doesn't exist, for read and write.
)

// [CreateFile] dwShareMode.
//
// [CreateFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type FILE_SHARE uint32

const (
	FILE_SHARE_NONE   FILE_SHARE = 0
	FILE_SHARE_READ   FILE_SHARE = 0x0000_0001
	FILE_SHARE_WRITE  FILE_SHARE = 0x0000_0002
	FILE_SHARE_DELETE FILE_SHARE = 0x0000_0004
)

// [GetVolumeInformation] flags.
//
// [GetVolumeInformation]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-getvolumeinformationw
type FILE_VOL uint32

const (
	FILE_VOL_CASE_PRESERVED_NAMES         FILE_VOL = 0x0000_0002
	FILE_VOL_CASE_SENSITIVE_SEARCH        FILE_VOL = 0x0000_0001
	FILE_VOL_DAX_VOLUME                   FILE_VOL = 0x2000_0000
	FILE_VOL_FILE_COMPRESSION             FILE_VOL = 0x0000_0010
	FILE_VOL_NAMED_STREAMS                FILE_VOL = 0x0004_0000
	FILE_VOL_PERSISTENT_ACLS              FILE_VOL = 0x0000_0008
	FILE_VOL_READ_ONLY_VOLUME             FILE_VOL = 0x0008_0000
	FILE_VOL_SEQUENTIAL_WRITE_ONCE        FILE_VOL = 0x0010_0000
	FILE_VOL_SUPPORTS_ENCRYPTION          FILE_VOL = 0x0002_0000
	FILE_VOL_SUPPORTS_EXTENDED_ATTRIBUTES FILE_VOL = 0x0080_0000
	FILE_VOL_SUPPORTS_HARD_LINKS          FILE_VOL = 0x0040_0000
	FILE_VOL_SUPPORTS_OBJECT_IDS          FILE_VOL = 0x0001_0000
	FILE_VOL_SUPPORTS_OPEN_BY_FILE_ID     FILE_VOL = 0x0100_0000
	FILE_VOL_SUPPORTS_REPARSE_POINTS      FILE_VOL = 0x0000_0080
	FILE_VOL_SUPPORTS_SPARSE_FILES        FILE_VOL = 0x0000_0040
	FILE_VOL_SUPPORTS_TRANSACTIONS        FILE_VOL = 0x0020_0000
	FILE_VOL_SUPPORTS_USN_JOURNAL         FILE_VOL = 0x0200_0000
	FILE_VOL_UNICODE_ON_DISK              FILE_VOL = 0x0000_0004
	FILE_VOL_VOLUME_IS_COMPRESSED         FILE_VOL = 0x0000_8000
	FILE_VOL_VOLUME_QUOTAS                FILE_VOL = 0x0000_0020
	FILE_VOL_SUPPORTS_BLOCK_REFCOUNTING   FILE_VOL = 0x0800_0000
)

// Generic access [rights].
//
// [rights]: https://learn.microsoft.com/en-us/windows/win32/secauthz/generic-access-rights
type GENERIC uint32

const (
	GENERIC_READ    GENERIC = 0x8000_0000
	GENERIC_WRITE   GENERIC = 0x4000_0000
	GENERIC_EXECUTE GENERIC = 0x2000_0000
	GENERIC_ALL     GENERIC = 0x1000_0000
)

// [GlobalAlloc] uFlags.
//
// [GlobalAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-globalalloc
type GMEM uint32

const (
	GMEM_FIXED    GMEM = 0x0000
	GMEM_MOVEABLE GMEM = 0x0002
	GMEM_ZEROINIT GMEM = 0x0040
	GMEM_MODIFY   GMEM = 0x0080
	GMEM_GHND     GMEM = GMEM_MOVEABLE | GMEM_ZEROINIT
	GMEM_GPTR     GMEM = GMEM_FIXED | GMEM_ZEROINIT
)

// [HeapAlloc] flags.
//
// [HeapAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapalloc
type HEAP_ALLOC uint32

const (
	HEAP_ALLOC_GENERATE_EXCEPTIONS HEAP_ALLOC = 0x0000_0004
	HEAP_ALLOC_NO_SERIALIZE        HEAP_ALLOC = 0x0000_0001
	HEAP_ALLOC_ZERO_MEMORY         HEAP_ALLOC = 0x0000_0008
)

// [HeapSetInformation] class.
//
// [HeapSetInformation]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapsetinformation
type HEAP_CLASS uint32

const (
	HEAP_CLASS_CompatibilityInformation      HEAP_CLASS = 0
	HEAP_CLASS_EnableTerminationOnCorruption HEAP_CLASS = 1
	HEAP_CLASS_OptimizeResources             HEAP_CLASS = 3
)

// [HeapCreate] options.
//
// [HeapCreate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapcreate
type HEAP_CREATE uint32

const (
	HEAP_CREATE_ENABLE_EXECUTE      HEAP_CREATE = 0x0004_0000
	HEAP_CREATE_GENERATE_EXCEPTIONS HEAP_CREATE = 0x0000_0004
	HEAP_CREATE_NO_SERIALIZE        HEAP_CREATE = 0x0000_0001
)

// [HeapFree], [HeapSize] and [HeapValidate] flags.
//
// [HeapFree]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapfree
// [HeapSize]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapsize
// [HeapValidate]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heapvalidate
type HEAP_NS uint32

const (
	HEAP_SER_NO_SERIALIZE HEAP_NS = 0x0000_0001
)

// [HeapReAlloc] flags.
//
// [HeapReAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/heapapi/nf-heapapi-heaprealloc
type HEAP_REALLOC uint32

const (
	HEAP_REALLOC_GENERATE_EXCEPTIONS   HEAP_REALLOC = 0x0000_0004
	HEAP_REALLOC_NO_SERIALIZE          HEAP_REALLOC = 0x0000_0001
	HEAP_REALLOC_REALLOC_IN_PLACE_ONLY HEAP_REALLOC = 0x0000_0010
	HEAP_REALLOC_ZERO_MEMORY           HEAP_REALLOC = 0x0000_0008
)

// [Language] identifier.
//
// [Language]: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifier-constants-and-strings
type LANG uint16

const (
	LANG_NEUTRAL             LANG = 0x00
	LANG_INVARIANT           LANG = 0x7f
	LANG_AFRIKAANS           LANG = 0x36
	LANG_ALBANIAN            LANG = 0x1c
	LANG_ALSATIAN            LANG = 0x84
	LANG_AMHARIC             LANG = 0x5e
	LANG_ARABIC              LANG = 0x01
	LANG_ARMENIAN            LANG = 0x2b
	LANG_ASSAMESE            LANG = 0x4d
	LANG_AZERI               LANG = 0x2c
	LANG_AZERBAIJANI         LANG = 0x2c
	LANG_BANGLA              LANG = 0x45
	LANG_BASHKIR             LANG = 0x6d
	LANG_BASQUE              LANG = 0x2d
	LANG_BELARUSIAN          LANG = 0x23
	LANG_BENGALI             LANG = 0x45
	LANG_BRETON              LANG = 0x7e
	LANG_BOSNIAN             LANG = 0x1a
	LANG_BOSNIAN_NEUTRAL     LANG = 0x781a
	LANG_BULGARIAN           LANG = 0x02
	LANG_CATALAN             LANG = 0x03
	LANG_CENTRAL_KURDISH     LANG = 0x92
	LANG_CHEROKEE            LANG = 0x5c
	LANG_CHINESE             LANG = 0x04
	LANG_CHINESE_SIMPLIFIED  LANG = 0x04
	LANG_CHINESE_TRADITIONAL LANG = 0x7c04
	LANG_CORSICAN            LANG = 0x83
	LANG_CROATIAN            LANG = 0x1a
	LANG_CZECH               LANG = 0x05
	LANG_DANISH              LANG = 0x06
	LANG_DARI                LANG = 0x8c
	LANG_DIVEHI              LANG = 0x65
	LANG_DUTCH               LANG = 0x13
	LANG_ENGLISH             LANG = 0x09
	LANG_ESTONIAN            LANG = 0x25
	LANG_FAEROESE            LANG = 0x38
	LANG_FARSI               LANG = 0x29
	LANG_FILIPINO            LANG = 0x64
	LANG_FINNISH             LANG = 0x0b
	LANG_FRENCH              LANG = 0x0c
	LANG_FRISIAN             LANG = 0x62
	LANG_FULAH               LANG = 0x67
	LANG_GALICIAN            LANG = 0x56
	LANG_GEORGIAN            LANG = 0x37
	LANG_GERMAN              LANG = 0x07
	LANG_GREEK               LANG = 0x08
	LANG_GREENLANDIC         LANG = 0x6f
	LANG_GUJARATI            LANG = 0x47
	LANG_HAUSA               LANG = 0x68
	LANG_HAWAIIAN            LANG = 0x75
	LANG_HEBREW              LANG = 0x0d
	LANG_HINDI               LANG = 0x39
	LANG_HUNGARIAN           LANG = 0x0e
	LANG_ICELANDIC           LANG = 0x0f
	LANG_IGBO                LANG = 0x70
	LANG_INDONESIAN          LANG = 0x21
	LANG_INUKTITUT           LANG = 0x5d
	LANG_IRISH               LANG = 0x3c
	LANG_ITALIAN             LANG = 0x10
	LANG_JAPANESE            LANG = 0x11
	LANG_KANNADA             LANG = 0x4b
	LANG_KASHMIRI            LANG = 0x60
	LANG_KAZAK               LANG = 0x3f
	LANG_KHMER               LANG = 0x53
	LANG_KICHE               LANG = 0x86
	LANG_KINYARWANDA         LANG = 0x87
	LANG_KONKANI             LANG = 0x57
	LANG_KOREAN              LANG = 0x12
	LANG_KYRGYZ              LANG = 0x40
	LANG_LAO                 LANG = 0x54
	LANG_LATVIAN             LANG = 0x26
	LANG_LITHUANIAN          LANG = 0x27
	LANG_LOWER_SORBIAN       LANG = 0x2e
	LANG_LUXEMBOURGISH       LANG = 0x6e
	LANG_MACEDONIAN          LANG = 0x2f
	LANG_MALAY               LANG = 0x3e
	LANG_MALAYALAM           LANG = 0x4c
	LANG_MALTESE             LANG = 0x3a
	LANG_MANIPURI            LANG = 0x58
	LANG_MAORI               LANG = 0x81
	LANG_MAPUDUNGUN          LANG = 0x7a
	LANG_MARATHI             LANG = 0x4e
	LANG_MOHAWK              LANG = 0x7c
	LANG_MONGOLIAN           LANG = 0x50
	LANG_NEPALI              LANG = 0x61
	LANG_NORWEGIAN           LANG = 0x14
	LANG_OCCITAN             LANG = 0x82
	LANG_ODIA                LANG = 0x48
	LANG_ORIYA               LANG = 0x48
	LANG_PASHTO              LANG = 0x63
	LANG_PERSIAN             LANG = 0x29
	LANG_POLISH              LANG = 0x15
	LANG_PORTUGUESE          LANG = 0x16
	LANG_PULAR               LANG = 0x67
	LANG_PUNJABI             LANG = 0x46
	LANG_QUECHUA             LANG = 0x6b
	LANG_ROMANIAN            LANG = 0x18
	LANG_ROMANSH             LANG = 0x17
	LANG_RUSSIAN             LANG = 0x19
	LANG_SAKHA               LANG = 0x85
	LANG_SAMI                LANG = 0x3b
	LANG_SANSKRIT            LANG = 0x4f
	LANG_SCOTTISH_GAELIC     LANG = 0x91
	LANG_SERBIAN             LANG = 0x1a
	LANG_SERBIAN_NEUTRAL     LANG = 0x7c1a
	LANG_SINDHI              LANG = 0x59
	LANG_SINHALESE           LANG = 0x5b
	LANG_SLOVAK              LANG = 0x1b
	LANG_SLOVENIAN           LANG = 0x24
	LANG_SOTHO               LANG = 0x6c
	LANG_SPANISH             LANG = 0x0a
	LANG_SWAHILI             LANG = 0x41
	LANG_SWEDISH             LANG = 0x1d
	LANG_SYRIAC              LANG = 0x5a
	LANG_TAJIK               LANG = 0x28
	LANG_TAMAZIGHT           LANG = 0x5f
	LANG_TAMIL               LANG = 0x49
	LANG_TATAR               LANG = 0x44
	LANG_TELUGU              LANG = 0x4a
	LANG_THAI                LANG = 0x1e
	LANG_TIBETAN             LANG = 0x51
	LANG_TIGRIGNA            LANG = 0x73
	LANG_TIGRINYA            LANG = 0x73
	LANG_TSWANA              LANG = 0x32
	LANG_TURKISH             LANG = 0x1f
	LANG_TURKMEN             LANG = 0x42
	LANG_UIGHUR              LANG = 0x80
	LANG_UKRAINIAN           LANG = 0x22
	LANG_UPPER_SORBIAN       LANG = 0x2e
	LANG_URDU                LANG = 0x20
	LANG_UZBEK               LANG = 0x43
	LANG_VALENCIAN           LANG = 0x03
	LANG_VIETNAMESE          LANG = 0x2a
	LANG_WELSH               LANG = 0x52
	LANG_WOLOF               LANG = 0x88
	LANG_XHOSA               LANG = 0x34
	LANG_YAKUT               LANG = 0x85
	LANG_YI                  LANG = 0x78
	LANG_YORUBA              LANG = 0x6a
	LANG_ZULU                LANG = 0x35
)

// [LocalAlloc] uFlags.
//
// [LocalAlloc]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-localalloc
type LMEM uint32

const (
	LMEM_FIXED    LMEM = 0x0000
	LMEM_MOVEABLE LMEM = 0x0002
	LMEM_ZEROINIT LMEM = 0x0040
	LMEM_MODIFY   LMEM = 0x0080
	LMEM_GHND     LMEM = LMEM_MOVEABLE | LMEM_ZEROINIT
	LMEM_GPTR     LMEM = LMEM_FIXED | LMEM_ZEROINIT
)

// [LockFileEx] dwFlags.
//
// [LockFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-lockfileex
type LOCKFILE uint32

const (
	LOCKFILE_NONE             LOCKFILE = 0
	LOCKFILE_FAIL_IMMEDIATELY LOCKFILE = 0x0000_0001
	LOCKFILE_EXCLUSIVE_LOCK   LOCKFILE = 0x0000_0002
)

// [MoveFileEx] dwFlags.
//
// [MoveFileEx]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-movefileexw
type MOVEFILE uint32

const (
	MOVEFILE_COPY_ALLOWED          MOVEFILE = 0x2
	MOVEFILE_CREATE_HARDLINK       MOVEFILE = 0x10
	MOVEFILE_DELAY_UNTIL_REBOOT    MOVEFILE = 0x4
	MOVEFILE_FAIL_IF_NOT_TRACKABLE MOVEFILE = 0x20
	MOVEFILE_REPLACE_EXISTING      MOVEFILE = 0x1
	MOVEFILE_WRITE_THROUGH         MOVEFILE = 0x8
)

// [CreateFileMapping] flProtect.
//
// [CreateFileMapping]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type PAGE uint32

const (
	PAGE_NONE                   PAGE = 0
	PAGE_NOACCESS               PAGE = 0x01
	PAGE_READONLY               PAGE = 0x02
	PAGE_READWRITE              PAGE = 0x04
	PAGE_WRITECOPY              PAGE = 0x08
	PAGE_EXECUTE                PAGE = 0x10
	PAGE_EXECUTE_READ           PAGE = 0x20
	PAGE_EXECUTE_READWRITE      PAGE = 0x40
	PAGE_EXECUTE_WRITECOPY      PAGE = 0x80
	PAGE_GUARD                  PAGE = 0x100
	PAGE_NOCACHE                PAGE = 0x200
	PAGE_WRITECOMBINE           PAGE = 0x400
	PAGE_ENCLAVE_THREAD_CONTROL PAGE = 0x8000_0000
	PAGE_REVERT_TO_FILE_MAP     PAGE = 0x8000_0000
	PAGE_TARGETS_NO_UPDATE      PAGE = 0x4000_0000
	PAGE_TARGETS_INVALID        PAGE = 0x4000_0000
	PAGE_ENCLAVE_UNVALIDATED    PAGE = 0x2000_0000
	PAGE_ENCLAVE_DECOMMIT       PAGE = 0x1000_0000
)

// [WM_POWERBROADCAST] event.
//
// [WM_POWERBROADCAST]: https://learn.microsoft.com/en-us/windows/win32/power/wm-powerbroadcast
type PBT uint32

const (
	PBT_APMQUERYSUSPEND       PBT = 0x0000
	PBT_APMQUERYSTANDBY       PBT = 0x0001
	PBT_APMQUERYSUSPENDFAILED PBT = 0x0002
	PBT_APMQUERYSTANDBYFAILED PBT = 0x0003
	PBT_APMSUSPEND            PBT = 0x0004
	PBT_APMSTANDBY            PBT = 0x0005
	PBT_APMRESUMECRITICAL     PBT = 0x0006
	PBT_APMRESUMESUSPEND      PBT = 0x0007
	PBT_APMRESUMESTANDBY      PBT = 0x0008
	PBT_APMBATTERYLOW         PBT = 0x0009
	PBT_APMPOWERSTATUSCHANGE  PBT = 0x000a
	PBT_APMOEMEVENT           PBT = 0x000b
	PBT_APMRESUMEAUTOMATIC    PBT = 0x0012
	PBT_POWERSETTINGCHANGE    PBT = 0x8013
)

// [CreateNamedPipe] dwPipeMode.
//
// [CreateNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-createnamedpipew
type PIPE uint32

const (
	PIPE_WAIT                  PIPE = 0x0000_0000
	PIPE_NOWAIT                PIPE = 0x0000_0001
	PIPE_READMODE_BYTE         PIPE = 0x0000_0000
	PIPE_READMODE_MESSAGE      PIPE = 0x0000_0002
	PIPE_TYPE_BYTE             PIPE = 0x0000_0000
	PIPE_TYPE_MESSAGE          PIPE = 0x0000_0004
	PIPE_ACCEPT_REMOTE_CLIENTS PIPE = 0x0000_0000
	PIPE_REJECT_REMOTE_CLIENTS PIPE = 0x0000_0008
)

// [CreateNamedPipe] dwOpenMode.
//
// [CreateNamedPipe]: https://learn.microsoft.com/en-us/windows/win32/api/namedpipeapi/nf-namedpipeapi-createnamedpipew
type PIPE_ACCESS uint32

const (
	PIPE_ACCESS_INBOUND  PIPE_ACCESS = 0x0000_0001
	PIPE_ACCESS_OUTBOUND PIPE_ACCESS = 0x0000_0002
	PIPE_ACCESS_DUPLEX   PIPE_ACCESS = 0x0000_0003
)

// Process [access rights].
//
// [access rights]: https://learn.microsoft.com/en-us/windows/win32/procthread/process-security-and-access-rights
type PROCESS uint32

const (
	PROCESS_ALL_ACCESS                PROCESS = PROCESS(STANDARD_RIGHTS_REQUIRED|STANDARD_RIGHTS_SYNCHRONIZE) | 0xffff
	PROCESS_CREATE_PROCESS            PROCESS = 0x0080
	PROCESS_CREATE_THREAD             PROCESS = 0x0002
	PROCESS_DUP_HANDLE                PROCESS = 0x0040
	PROCESS_QUERY_INFORMATION         PROCESS = 0x0400
	PROCESS_QUERY_LIMITED_INFORMATION PROCESS = 0x1000
	PROCESS_SET_LIMITED_INFORMATION   PROCESS = 0x2000
	PROCESS_SET_INFORMATION           PROCESS = 0x0200
	PROCESS_SET_QUOTA                 PROCESS = 0x0100
	PROCESS_SET_SESSIONID             PROCESS = 0x0004
	PROCESS_SUSPEND_RESUME            PROCESS = 0x0800
	PROCESS_TERMINATE                 PROCESS = 0x0001
	PROCESS_VM_OPERATION              PROCESS = 0x0008
	PROCESS_VM_READ                   PROCESS = 0x0010
	PROCESS_VM_WRITE                  PROCESS = 0x0020
	PROCESS_SYNCHRONIZE               PROCESS = PROCESS(STANDARD_RIGHTS_SYNCHRONIZE)
)

// [SYSTEM_INFO] dwProcessorType.
//
// [SYSTEM_INFO]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type PROCESSOR uint32

const (
	PROCESSOR_INTEL_386     PROCESSOR = 386
	PROCESSOR_INTEL_486     PROCESSOR = 486
	PROCESSOR_INTEL_PENTIUM PROCESSOR = 586
	PROCESSOR_INTEL_IA64    PROCESSOR = 2200
	PROCESSOR_AMD_X8664     PROCESSOR = 8664
	PROCESSOR_MIPS_R4000    PROCESSOR = 4000
	PROCESSOR_ALPHA_21064   PROCESSOR = 21064
	PROCESSOR_PPC_601       PROCESSOR = 601
	PROCESSOR_PPC_603       PROCESSOR = 603
	PROCESSOR_PPC_604       PROCESSOR = 604
	PROCESSOR_PPC_620       PROCESSOR = 620
	PROCESSOR_HITACHI_SH3   PROCESSOR = 10003
	PROCESSOR_HITACHI_SH3E  PROCESSOR = 10004
	PROCESSOR_HITACHI_SH4   PROCESSOR = 10005
	PROCESSOR_MOTOROLA_821  PROCESSOR = 821
	PROCESSOR_SHx_SH3       PROCESSOR = 103
	PROCESSOR_SHx_SH4       PROCESSOR = 104
	PROCESSOR_STRONGARM     PROCESSOR = 2577
	PROCESSOR_ARM720        PROCESSOR = 1824
	PROCESSOR_ARM820        PROCESSOR = 2080
	PROCESSOR_ARM920        PROCESSOR = 2336
	PROCESSOR_ARM_7TDMI     PROCESSOR = 70001
	PROCESSOR_OPTIL         PROCESSOR = 0x494f
)

// [SYSTEM_INFO] wProcessorArchitecture.
//
// [SYSTEM_INFO]: https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info
type PROCESSOR_ARCHITECTURE uint16

const (
	PROCESSOR_ARCHITECTURE_INTEL          PROCESSOR_ARCHITECTURE = 0
	PROCESSOR_ARCHITECTURE_MIPS           PROCESSOR_ARCHITECTURE = 1
	PROCESSOR_ARCHITECTURE_ALPHA          PROCESSOR_ARCHITECTURE = 2
	PROCESSOR_ARCHITECTURE_PPC            PROCESSOR_ARCHITECTURE = 3
	PROCESSOR_ARCHITECTURE_SHX            PROCESSOR_ARCHITECTURE = 4
	PROCESSOR_ARCHITECTURE_ARM            PROCESSOR_ARCHITECTURE = 5
	PROCESSOR_ARCHITECTURE_IA64           PROCESSOR_ARCHITECTURE = 6
	PROCESSOR_ARCHITECTURE_ALPHA64        PROCESSOR_ARCHITECTURE = 7
	PROCESSOR_ARCHITECTURE_MSIL           PROCESSOR_ARCHITECTURE = 8
	PROCESSOR_ARCHITECTURE_AMD64          PROCESSOR_ARCHITECTURE = 9
	PROCESSOR_ARCHITECTURE_IA32_ON_WIN64  PROCESSOR_ARCHITECTURE = 10
	PROCESSOR_ARCHITECTURE_NEUTRAL        PROCESSOR_ARCHITECTURE = 11
	PROCESSOR_ARCHITECTURE_ARM64          PROCESSOR_ARCHITECTURE = 12
	PROCESSOR_ARCHITECTURE_ARM32_ON_WIN64 PROCESSOR_ARCHITECTURE = 13
	PROCESSOR_ARCHITECTURE_IA32_ON_ARM64  PROCESSOR_ARCHITECTURE = 14
	PROCESSOR_ARCHITECTURE_UNKNOWN        PROCESSOR_ARCHITECTURE = 0xffff
)

// [ReplaceFile] dwReplaceFlags.
//
// [ReplaceFile]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-replacefilew
type REPLACEFILE uint32

const (
	REPLACEFILE_NONE                REPLACEFILE = 0
	REPLACEFILE_WRITE_THROUGH       REPLACEFILE = 0x0000_0001
	REPLACEFILE_IGNORE_MERGE_ERRORS REPLACEFILE = 0x0000_0002
	REPLACEFILE_IGNORE_ACL_ERRORS   REPLACEFILE = 0x0000_0004
)

// Predefined [resource types].
//
// [resource types]: https://learn.microsoft.com/en-us/windows/win32/menurc/resource-types
type RT uint16

const (
	RT_ACCELERATOR  RT = 9
	RT_ANICURSOR    RT = 21
	RT_ANIICON      RT = 22
	RT_BITMAP       RT = 2
	RT_CURSOR       RT = 1
	RT_DIALOG       RT = 5
	RT_DLGINCLUDE   RT = 17
	RT_FONT         RT = 8
	RT_FONTDIR      RT = 7
	RT_GROUP_CURSOR RT = 12
	RT_GROUP_ICON   RT = 14
	RT_HTML         RT = 23
	RT_ICON         RT = 3
	RT_MANIFEST     RT = 24
	RT_MENU         RT = 4
	RT_MESSAGETABLE RT = 11
	RT_PLUGPLAY     RT = 19
	RT_RCDATA       RT = 10
	RT_STRING       RT = 6
	RT_VERSION      RT = 16
	RT_VXD          RT = 20
)

// [CreateFileMapping] flProtect.
//
// [CreateFileMapping]: https://learn.microsoft.com/en-us/windows/win32/api/memoryapi/nf-memoryapi-createfilemappingw
type SEC uint32

const (
	SEC_NONE                   SEC = 0
	SEC_PARTITION_OWNER_HANDLE SEC = 0x0004_0000
	SEC_64K_PAGES              SEC = 0x0008_0000
	SEC_FILE                   SEC = 0x0080_0000
	SEC_IMAGE                  SEC = 0x0100_0000
	SEC_PROTECTED_IMAGE        SEC = 0x0200_0000
	SEC_RESERVE                SEC = 0x0400_0000
	SEC_COMMIT                 SEC = 0x0800_0000
	SEC_NOCACHE                SEC = 0x1000_0000
	SEC_WRITECOMBINE           SEC = 0x4000_0000
	SEC_LARGE_PAGES            SEC = 0x8000_0000
	SEC_IMAGE_NO_EXECUTE       SEC = SEC_IMAGE | SEC_NOCACHE
)

// [CreateFile] dwFlagsAndAttributes.
//
// [CreateFile]: https://learn.microsoft.com/en-us/windows/win32/api/fileapi/nf-fileapi-createfilew
type SECURITY uint32

const (
	SECURITY_NONE             SECURITY = 0
	SECURITY_ANONYMOUS        SECURITY = 0 << 16
	SECURITY_IDENTIFICATION   SECURITY = 1 << 16
	SECURITY_IMPERSONATION    SECURITY = 2 << 16
	SECURITY_DELEGATION       SECURITY = 3 << 16
	SECURITY_CONTEXT_TRACKING SECURITY = 0x0004_0000
	SECURITY_EFFECTIVE_ONLY   SECURITY = 0x0008_0000
)

// Sort order [identifier] for locales.
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/sort-order-identifiers
type SORT uint16

const (
	SORT_DEFAULT                SORT = 0x0
	SORT_INVARIANT_MATH         SORT = 0x1
	SORT_JAPANESE_XJIS          SORT = 0x0
	SORT_JAPANESE_UNICODE       SORT = 0x1
	SORT_JAPANESE_RADICALSTROKE SORT = 0x4
	SORT_CHINESE_BIG5           SORT = 0x0
	SORT_CHINESE_PRCP           SORT = 0x0
	SORT_CHINESE_UNICODE        SORT = 0x1
	SORT_CHINESE_PRC            SORT = 0x2
	SORT_CHINESE_BOPOMOFO       SORT = 0x3
	SORT_CHINESE_RADICALSTROKE  SORT = 0x4
	SORT_KOREAN_KSC             SORT = 0x0
	SORT_KOREAN_UNICODE         SORT = 0x1
	SORT_GERMAN_PHONE_BOOK      SORT = 0x1
	SORT_HUNGARIAN_DEFAULT      SORT = 0x0
	SORT_HUNGARIAN_TECHNICAL    SORT = 0x1
	SORT_GEORGIAN_TRADITIONAL   SORT = 0x0
	SORT_GEORGIAN_MODERN        SORT = 0x1
)

// Standard [access rights]. These are generic and compose other access right
// types. Also includes unprefixed and SPECIFIC_RIGHT prefix.
//
// [access rights]: https://learn.microsoft.com/en-us/windows/win32/secauthz/standard-access-rights
type STANDARD_RIGHTS uint32

const (
	STANDARD_RIGHTS_NONE STANDARD_RIGHTS = 0

	STANDARD_RIGHTS_DELETE       STANDARD_RIGHTS = 0x0001_0000
	STANDARD_RIGHTS_READ_CONTROL STANDARD_RIGHTS = 0x0002_0000
	STANDARD_RIGHTS_SYNCHRONIZE  STANDARD_RIGHTS = 0x0010_0000
	STANDARD_RIGHTS_WRITE_DAC    STANDARD_RIGHTS = 0x0004_0000
	STANDARD_RIGHTS_WRITE_OWNER  STANDARD_RIGHTS = 0x0008_0000

	STANDARD_RIGHTS_ALL      STANDARD_RIGHTS = 0x001f_0000
	STANDARD_RIGHTS_EXECUTE  STANDARD_RIGHTS = STANDARD_RIGHTS_READ_CONTROL
	STANDARD_RIGHTS_READ     STANDARD_RIGHTS = STANDARD_RIGHTS_READ_CONTROL
	STANDARD_RIGHTS_REQUIRED STANDARD_RIGHTS = 0x000f_0000
	STANDARD_RIGHTS_WRITE    STANDARD_RIGHTS = STANDARD_RIGHTS_READ_CONTROL
)

// [STARTUPINFO] dwFlags.
//
// [STARTUPINFO]: https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/ns-processthreadsapi-startupinfow
type STARTF uint32

const (
	STARTF_FORCEONFEEDBACK  STARTF = 0x0000_0040
	STARTF_FORCEOFFFEEDBACK STARTF = 0x0000_0080
	STARTF_PREVENTPINNING   STARTF = 0x0000_2000
	STARTF_RUNFULLSCREEN    STARTF = 0x0000_0020
	STARTF_TITLEISAPPID     STARTF = 0x0000_1000
	STARTF_TITLEISLINKNAME  STARTF = 0x0000_0800
	STARTF_UNTRUSTEDSOURCE  STARTF = 0x0000_8000
	STARTF_USECOUNTCHARS    STARTF = 0x0000_0008
	STARTF_USEFILLATTRIBUTE STARTF = 0x0000_0010
	STARTF_USEHOTKEY        STARTF = 0x0000_0200
	STARTF_USEPOSITION      STARTF = 0x0000_0004
	STARTF_USESHOWWINDOW    STARTF = 0x0000_0001
	STARTF_USESIZE          STARTF = 0x0000_0002
	STARTF_USESTDHANDLES    STARTF = 0x0000_0100
)

// Standard [devices].
//
// [devices]: https://learn.microsoft.com/en-us/windows/console/getstdhandle
type STD int32

const (
	STD_INPUT_HANDLE  STD = -10
	STD_OUTPUT_HANDLE STD = -11
	STD_ERROR_HANDLE  STD = -12
)

// Sub-language [identifier].
//
// [identifier]: https://learn.microsoft.com/en-us/windows/win32/intl/language-identifier-constants-and-strings
type SUBLANG uint16

const (
	SUBLANG_NEUTRAL                             SUBLANG = 0x00
	SUBLANG_DEFAULT                             SUBLANG = 0x01
	SUBLANG_SYS_DEFAULT                         SUBLANG = 0x02
	SUBLANG_CUSTOM_DEFAULT                      SUBLANG = 0x03
	SUBLANG_CUSTOM_UNSPECIFIED                  SUBLANG = 0x04
	SUBLANG_UI_CUSTOM_DEFAULT                   SUBLANG = 0x05
	SUBLANG_AFRIKAANS_SOUTH_AFRICA              SUBLANG = 0x01
	SUBLANG_ALBANIAN_ALBANIA                    SUBLANG = 0x01
	SUBLANG_ALSATIAN_FRANCE                     SUBLANG = 0x01
	SUBLANG_AMHARIC_ETHIOPIA                    SUBLANG = 0x01
	SUBLANG_ARABIC_SAUDI_ARABIA                 SUBLANG = 0x01
	SUBLANG_ARABIC_IRAQ                         SUBLANG = 0x02
	SUBLANG_ARABIC_EGYPT                        SUBLANG = 0x03
	SUBLANG_ARABIC_LIBYA                        SUBLANG = 0x04
	SUBLANG_ARABIC_ALGERIA                      SUBLANG = 0x05
	SUBLANG_ARABIC_MOROCCO                      SUBLANG = 0x06
	SUBLANG_ARABIC_TUNISIA                      SUBLANG = 0x07
	SUBLANG_ARABIC_OMAN                         SUBLANG = 0x08
	SUBLANG_ARABIC_YEMEN                        SUBLANG = 0x09
	SUBLANG_ARABIC_SYRIA                        SUBLANG = 0x0a
	SUBLANG_ARABIC_JORDAN                       SUBLANG = 0x0b
	SUBLANG_ARABIC_LEBANON                      SUBLANG = 0x0c
	SUBLANG_ARABIC_KUWAIT                       SUBLANG = 0x0d
	SUBLANG_ARABIC_UAE                          SUBLANG = 0x0e
	SUBLANG_ARABIC_BAHRAIN                      SUBLANG = 0x0f
	SUBLANG_ARABIC_QATAR                        SUBLANG = 0x10
	SUBLANG_ARMENIAN_ARMENIA                    SUBLANG = 0x01
	SUBLANG_ASSAMESE_INDIA                      SUBLANG = 0x01
	SUBLANG_AZERI_LATIN                         SUBLANG = 0x01
	SUBLANG_AZERI_CYRILLIC                      SUBLANG = 0x02
	SUBLANG_AZERBAIJANI_AZERBAIJAN_LATIN        SUBLANG = 0x01
	SUBLANG_AZERBAIJANI_AZERBAIJAN_CYRILLIC     SUBLANG = 0x02
	SUBLANG_BANGLA_INDIA                        SUBLANG = 0x01
	SUBLANG_BANGLA_BANGLADESH                   SUBLANG = 0x02
	SUBLANG_BASHKIR_RUSSIA                      SUBLANG = 0x01
	SUBLANG_BASQUE_BASQUE                       SUBLANG = 0x01
	SUBLANG_BELARUSIAN_BELARUS                  SUBLANG = 0x01
	SUBLANG_BENGALI_INDIA                       SUBLANG = 0x01
	SUBLANG_BENGALI_BANGLADESH                  SUBLANG = 0x02
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_LATIN    SUBLANG = 0x05
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_CYRILLIC SUBLANG = 0x08
	SUBLANG_BRETON_FRANCE                       SUBLANG = 0x01
	SUBLANG_BULGARIAN_BULGARIA                  SUBLANG = 0x01
	SUBLANG_CATALAN_CATALAN                     SUBLANG = 0x01
	SUBLANG_CENTRAL_KURDISH_IRAQ                SUBLANG = 0x01
	SUBLANG_CHEROKEE_CHEROKEE                   SUBLANG = 0x01
	SUBLANG_CHINESE_TRADITIONAL                 SUBLANG = 0x01
	SUBLANG_CHINESE_SIMPLIFIED                  SUBLANG = 0x02
	SUBLANG_CHINESE_HONGKONG                    SUBLANG = 0x03
	SUBLANG_CHINESE_SINGAPORE                   SUBLANG = 0x04
	SUBLANG_CHINESE_MACAU                       SUBLANG = 0x05
	SUBLANG_CORSICAN_FRANCE                     SUBLANG = 0x01
	SUBLANG_CZECH_CZECH_REPUBLIC                SUBLANG = 0x01
	SUBLANG_CROATIAN_CROATIA                    SUBLANG = 0x01
	SUBLANG_CROATIAN_BOSNIA_HERZEGOVINA_LATIN   SUBLANG = 0x04
	SUBLANG_DANISH_DENMARK                      SUBLANG = 0x01
	SUBLANG_DARI_AFGHANISTAN                    SUBLANG = 0x01
	SUBLANG_DIVEHI_MALDIVES                     SUBLANG = 0x01
	SUBLANG_DUTCH                               SUBLANG = 0x01
	SUBLANG_DUTCH_BELGIAN                       SUBLANG = 0x02
	SUBLANG_ENGLISH_US                          SUBLANG = 0x01
	SUBLANG_ENGLISH_UK                          SUBLANG = 0x02
	SUBLANG_ENGLISH_AUS                         SUBLANG = 0x03
	SUBLANG_ENGLISH_CAN                         SUBLANG = 0x04
	SUBLANG_ENGLISH_NZ                          SUBLANG = 0x05
	SUBLANG_ENGLISH_EIRE                        SUBLANG = 0x06
	SUBLANG_ENGLISH_SOUTH_AFRICA                SUBLANG = 0x07
	SUBLANG_ENGLISH_JAMAICA                     SUBLANG = 0x08
	SUBLANG_ENGLISH_CARIBBEAN                   SUBLANG = 0x09
	SUBLANG_ENGLISH_BELIZE                      SUBLANG = 0x0a
	SUBLANG_ENGLISH_TRINIDAD                    SUBLANG = 0x0b
	SUBLANG_ENGLISH_ZIMBABWE                    SUBLANG = 0x0c
	SUBLANG_ENGLISH_PHILIPPINES                 SUBLANG = 0x0d
	SUBLANG_ENGLISH_INDIA                       SUBLANG = 0x10
	SUBLANG_ENGLISH_MALAYSIA                    SUBLANG = 0x11
	SUBLANG_ENGLISH_SINGAPORE                   SUBLANG = 0x12
	SUBLANG_ESTONIAN_ESTONIA                    SUBLANG = 0x01
	SUBLANG_FAEROESE_FAROE_ISLANDS              SUBLANG = 0x01
	SUBLANG_FILIPINO_PHILIPPINES                SUBLANG = 0x01
	SUBLANG_FINNISH_FINLAND                     SUBLANG = 0x01
	SUBLANG_FRENCH                              SUBLANG = 0x01
	SUBLANG_FRENCH_BELGIAN                      SUBLANG = 0x02
	SUBLANG_FRENCH_CANADIAN                     SUBLANG = 0x03
	SUBLANG_FRENCH_SWISS                        SUBLANG = 0x04
	SUBLANG_FRENCH_LUXEMBOURG                   SUBLANG = 0x05
	SUBLANG_FRENCH_MONACO                       SUBLANG = 0x06
	SUBLANG_FRISIAN_NETHERLANDS                 SUBLANG = 0x01
	SUBLANG_FULAH_SENEGAL                       SUBLANG = 0x02
	SUBLANG_GALICIAN_GALICIAN                   SUBLANG = 0x01
	SUBLANG_GEORGIAN_GEORGIA                    SUBLANG = 0x01
	SUBLANG_GERMAN                              SUBLANG = 0x01
	SUBLANG_GERMAN_SWISS                        SUBLANG = 0x02
	SUBLANG_GERMAN_AUSTRIAN                     SUBLANG = 0x03
	SUBLANG_GERMAN_LUXEMBOURG                   SUBLANG = 0x04
	SUBLANG_GERMAN_LIECHTENSTEIN                SUBLANG = 0x05
	SUBLANG_GREEK_GREECE                        SUBLANG = 0x01
	SUBLANG_GREENLANDIC_GREENLAND               SUBLANG = 0x01
	SUBLANG_GUJARATI_INDIA                      SUBLANG = 0x01
	SUBLANG_HAUSA_NIGERIA_LATIN                 SUBLANG = 0x01
	SUBLANG_HAWAIIAN_US                         SUBLANG = 0x01
	SUBLANG_HEBREW_ISRAEL                       SUBLANG = 0x01
	SUBLANG_HINDI_INDIA                         SUBLANG = 0x01
	SUBLANG_HUNGARIAN_HUNGARY                   SUBLANG = 0x01
	SUBLANG_ICELANDIC_ICELAND                   SUBLANG = 0x01
	SUBLANG_IGBO_NIGERIA                        SUBLANG = 0x01
	SUBLANG_INDONESIAN_INDONESIA                SUBLANG = 0x01
	SUBLANG_INUKTITUT_CANADA                    SUBLANG = 0x01
	SUBLANG_INUKTITUT_CANADA_LATIN              SUBLANG = 0x02
	SUBLANG_IRISH_IRELAND                       SUBLANG = 0x02
	SUBLANG_ITALIAN                             SUBLANG = 0x01
	SUBLANG_ITALIAN_SWISS                       SUBLANG = 0x02
	SUBLANG_JAPANESE_JAPAN                      SUBLANG = 0x01
	SUBLANG_KANNADA_INDIA                       SUBLANG = 0x01
	SUBLANG_KASHMIRI_SASIA                      SUBLANG = 0x02
	SUBLANG_KASHMIRI_INDIA                      SUBLANG = 0x02
	SUBLANG_KAZAK_KAZAKHSTAN                    SUBLANG = 0x01
	SUBLANG_KHMER_CAMBODIA                      SUBLANG = 0x01
	SUBLANG_KICHE_GUATEMALA                     SUBLANG = 0x01
	SUBLANG_KINYARWANDA_RWANDA                  SUBLANG = 0x01
	SUBLANG_KONKANI_INDIA                       SUBLANG = 0x01
	SUBLANG_KOREAN                              SUBLANG = 0x01
	SUBLANG_KYRGYZ_KYRGYZSTAN                   SUBLANG = 0x01
	SUBLANG_LAO_LAO                             SUBLANG = 0x01
	SUBLANG_LATVIAN_LATVIA                      SUBLANG = 0x01
	SUBLANG_LITHUANIAN                          SUBLANG = 0x01
	SUBLANG_LOWER_SORBIAN_GERMANY               SUBLANG = 0x02
	SUBLANG_LUXEMBOURGISH_LUXEMBOURG            SUBLANG = 0x01
	SUBLANG_MACEDONIAN_MACEDONIA                SUBLANG = 0x01
	SUBLANG_MALAY_MALAYSIA                      SUBLANG = 0x01
	SUBLANG_MALAY_BRUNEI_DARUSSALAM             SUBLANG = 0x02
	SUBLANG_MALAYALAM_INDIA                     SUBLANG = 0x01
	SUBLANG_MALTESE_MALTA                       SUBLANG = 0x01
	SUBLANG_MAORI_NEW_ZEALAND                   SUBLANG = 0x01
	SUBLANG_MAPUDUNGUN_CHILE                    SUBLANG = 0x01
	SUBLANG_MARATHI_INDIA                       SUBLANG = 0x01
	SUBLANG_MOHAWK_MOHAWK                       SUBLANG = 0x01
	SUBLANG_MONGOLIAN_CYRILLIC_MONGOLIA         SUBLANG = 0x01
	SUBLANG_MONGOLIAN_PRC                       SUBLANG = 0x02
	SUBLANG_NEPALI_INDIA                        SUBLANG = 0x02
	SUBLANG_NEPALI_NEPAL                        SUBLANG = 0x01
	SUBLANG_NORWEGIAN_BOKMAL                    SUBLANG = 0x01
	SUBLANG_NORWEGIAN_NYNORSK                   SUBLANG = 0x02
	SUBLANG_OCCITAN_FRANCE                      SUBLANG = 0x01
	SUBLANG_ODIA_INDIA                          SUBLANG = 0x01
	SUBLANG_ORIYA_INDIA                         SUBLANG = 0x01
	SUBLANG_PASHTO_AFGHANISTAN                  SUBLANG = 0x01
	SUBLANG_PERSIAN_IRAN                        SUBLANG = 0x01
	SUBLANG_POLISH_POLAND                       SUBLANG = 0x01
	SUBLANG_PORTUGUESE                          SUBLANG = 0x02
	SUBLANG_PORTUGUESE_BRAZILIAN                SUBLANG = 0x01
	SUBLANG_PULAR_SENEGAL                       SUBLANG = 0x02
	SUBLANG_PUNJABI_INDIA                       SUBLANG = 0x01
	SUBLANG_PUNJABI_PAKISTAN                    SUBLANG = 0x02
	SUBLANG_QUECHUA_BOLIVIA                     SUBLANG = 0x01
	SUBLANG_QUECHUA_ECUADOR                     SUBLANG = 0x02
	SUBLANG_QUECHUA_PERU                        SUBLANG = 0x03
	SUBLANG_ROMANIAN_ROMANIA                    SUBLANG = 0x01
	SUBLANG_ROMANSH_SWITZERLAND                 SUBLANG = 0x01
	SUBLANG_RUSSIAN_RUSSIA                      SUBLANG = 0x01
	SUBLANG_SAKHA_RUSSIA                        SUBLANG = 0x01
	SUBLANG_SAMI_NORTHERN_NORWAY                SUBLANG = 0x01
	SUBLANG_SAMI_NORTHERN_SWEDEN                SUBLANG = 0x02
	SUBLANG_SAMI_NORTHERN_FINLAND               SUBLANG = 0x03
	SUBLANG_SAMI_LULE_NORWAY                    SUBLANG = 0x04
	SUBLANG_SAMI_LULE_SWEDEN                    SUBLANG = 0x05
	SUBLANG_SAMI_SOUTHERN_NORWAY                SUBLANG = 0x06
	SUBLANG_SAMI_SOUTHERN_SWEDEN                SUBLANG = 0x07
	SUBLANG_SAMI_SKOLT_FINLAND                  SUBLANG = 0x08
	SUBLANG_SAMI_INARI_FINLAND                  SUBLANG = 0x09
	SUBLANG_SANSKRIT_INDIA                      SUBLANG = 0x01
	SUBLANG_SCOTTISH_GAELIC                     SUBLANG = 0x01
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_LATIN    SUBLANG = 0x06
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_CYRILLIC SUBLANG = 0x07
	SUBLANG_SERBIAN_MONTENEGRO_LATIN            SUBLANG = 0x0b
	SUBLANG_SERBIAN_MONTENEGRO_CYRILLIC         SUBLANG = 0x0c
	SUBLANG_SERBIAN_SERBIA_LATIN                SUBLANG = 0x09
	SUBLANG_SERBIAN_SERBIA_CYRILLIC             SUBLANG = 0x0a
	SUBLANG_SERBIAN_CROATIA                     SUBLANG = 0x01
	SUBLANG_SERBIAN_LATIN                       SUBLANG = 0x02
	SUBLANG_SERBIAN_CYRILLIC                    SUBLANG = 0x03
	SUBLANG_SINDHI_INDIA                        SUBLANG = 0x01
	SUBLANG_SINDHI_PAKISTAN                     SUBLANG = 0x02
	SUBLANG_SINDHI_AFGHANISTAN                  SUBLANG = 0x02
	SUBLANG_SINHALESE_SRI_LANKA                 SUBLANG = 0x01
	SUBLANG_SOTHO_NORTHERN_SOUTH_AFRICA         SUBLANG = 0x01
	SUBLANG_SLOVAK_SLOVAKIA                     SUBLANG = 0x01
	SUBLANG_SLOVENIAN_SLOVENIA                  SUBLANG = 0x01
	SUBLANG_SPANISH                             SUBLANG = 0x01
	SUBLANG_SPANISH_MEXICAN                     SUBLANG = 0x02
	SUBLANG_SPANISH_MODERN                      SUBLANG = 0x03
	SUBLANG_SPANISH_GUATEMALA                   SUBLANG = 0x04
	SUBLANG_SPANISH_COSTA_RICA                  SUBLANG = 0x05
	SUBLANG_SPANISH_PANAMA                      SUBLANG = 0x06
	SUBLANG_SPANISH_DOMINICAN_REPUBLIC          SUBLANG = 0x07
	SUBLANG_SPANISH_VENEZUELA                   SUBLANG = 0x08
	SUBLANG_SPANISH_COLOMBIA                    SUBLANG = 0x09
	SUBLANG_SPANISH_PERU                        SUBLANG = 0x0a
	SUBLANG_SPANISH_ARGENTINA                   SUBLANG = 0x0b
	SUBLANG_SPANISH_ECUADOR                     SUBLANG = 0x0c
	SUBLANG_SPANISH_CHILE                       SUBLANG = 0x0d
	SUBLANG_SPANISH_URUGUAY                     SUBLANG = 0x0e
	SUBLANG_SPANISH_PARAGUAY                    SUBLANG = 0x0f
	SUBLANG_SPANISH_BOLIVIA                     SUBLANG = 0x10
	SUBLANG_SPANISH_EL_SALVADOR                 SUBLANG = 0x11
	SUBLANG_SPANISH_HONDURAS                    SUBLANG = 0x12
	SUBLANG_SPANISH_NICARAGUA                   SUBLANG = 0x13
	SUBLANG_SPANISH_PUERTO_RICO                 SUBLANG = 0x14
	SUBLANG_SPANISH_US                          SUBLANG = 0x15
	SUBLANG_SWAHILI_KENYA                       SUBLANG = 0x01
	SUBLANG_SWEDISH                             SUBLANG = 0x01
	SUBLANG_SWEDISH_FINLAND                     SUBLANG = 0x02
	SUBLANG_SYRIAC_SYRIA                        SUBLANG = 0x01
	SUBLANG_TAJIK_TAJIKISTAN                    SUBLANG = 0x01
	SUBLANG_TAMAZIGHT_ALGERIA_LATIN             SUBLANG = 0x02
	SUBLANG_TAMAZIGHT_MOROCCO_TIFINAGH          SUBLANG = 0x04
	SUBLANG_TAMIL_INDIA                         SUBLANG = 0x01
	SUBLANG_TAMIL_SRI_LANKA                     SUBLANG = 0x02
	SUBLANG_TATAR_RUSSIA                        SUBLANG = 0x01
	SUBLANG_TELUGU_INDIA                        SUBLANG = 0x01
	SUBLANG_THAI_THAILAND                       SUBLANG = 0x01
	SUBLANG_TIBETAN_PRC                         SUBLANG = 0x01
	SUBLANG_TIGRIGNA_ERITREA                    SUBLANG = 0x02
	SUBLANG_TIGRINYA_ERITREA                    SUBLANG = 0x02
	SUBLANG_TIGRINYA_ETHIOPIA                   SUBLANG = 0x01
	SUBLANG_TSWANA_BOTSWANA                     SUBLANG = 0x02
	SUBLANG_TSWANA_SOUTH_AFRICA                 SUBLANG = 0x01
	SUBLANG_TURKISH_TURKEY                      SUBLANG = 0x01
	SUBLANG_TURKMEN_TURKMENISTAN                SUBLANG = 0x01
	SUBLANG_UIGHUR_PRC                          SUBLANG = 0x01
	SUBLANG_UKRAINIAN_UKRAINE                   SUBLANG = 0x01
	SUBLANG_UPPER_SORBIAN_GERMANY               SUBLANG = 0x01
	SUBLANG_URDU_PAKISTAN                       SUBLANG = 0x01
	SUBLANG_URDU_INDIA                          SUBLANG = 0x02
	SUBLANG_UZBEK_LATIN                         SUBLANG = 0x01
	SUBLANG_UZBEK_CYRILLIC                      SUBLANG = 0x02
	SUBLANG_VALENCIAN_VALENCIA                  SUBLANG = 0x02
	SUBLANG_VIETNAMESE_VIETNAM                  SUBLANG = 0x01
	SUBLANG_WELSH_UNITED_KINGDOM                SUBLANG = 0x01
	SUBLANG_WOLOF_SENEGAL                       SUBLANG = 0x01
	SUBLANG_XHOSA_SOUTH_AFRICA                  SUBLANG = 0x01
	SUBLANG_YAKUT_RUSSIA                        SUBLANG = 0x01
	SUBLANG_YI_PRC                              SUBLANG = 0x01
	SUBLANG_YORUBA_NIGERIA                      SUBLANG = 0x01
	SUBLANG_ZULU_SOUTH_AFRICA                   SUBLANG = 0x01
)

// [CreateToolhelp32Snapshot] dwFlags.
//
// [CreateToolhelp32Snapshot]: https://learn.microsoft.com/en-us/windows/win32/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
type TH32CS uint32

const (
	TH32CS_SNAPHEAPLIST TH32CS = 0x0000_0001
	TH32CS_SNAPPROCESS  TH32CS = 0x0000_0002
	TH32CS_SNAPTHREAD   TH32CS = 0x0000_0004
	TH32CS_SNAPMODULE   TH32CS = 0x0000_0008
	TH32CS_SNAPMODULE32 TH32CS = 0x0000_0010
	TH32CS_SNAPALL      TH32CS = (TH32CS_SNAPHEAPLIST | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD | TH32CS_SNAPMODULE)
	TH32CS_INHERIT      TH32CS = 0x8000_0000
)

// [GetTimeZoneInformation] return value.
//
// [GetTimeZoneInformation]: https://learn.microsoft.com/en-us/windows/win32/api/timezoneapi/nf-timezoneapi-gettimezoneinformation
type TIME_ZONE_ID uint32

const (
	TIME_ZONE_ID_UNKNOWN  TIME_ZONE_ID = 0
	TIME_ZONE_ID_STANDARD TIME_ZONE_ID = 1
	TIME_ZONE_ID_DAYLIGHT TIME_ZONE_ID = 2
)

// [VerifyVersionInfo] dwTypeMask.
//
// [VerifyVersionInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
type VER uint32

const (
	VER_BUILDNUMBER      VER = 0x000_0004
	VER_MAJORVERSION     VER = 0x000_0002
	VER_MINORVERSION     VER = 0x000_0001
	VER_PLATFORMID       VER = 0x000_0008
	VER_PRODUCT_TYPE     VER = 0x000_0080
	VER_SERVICEPACKMAJOR VER = 0x000_0020
	VER_SERVICEPACKMINOR VER = 0x000_0010
	VER_SUITENAME        VER = 0x000_0040
)

// [VerifyVersionInfo] dwlConditionMask.
//
// [VerifyVersionInfo]: https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-verifyversioninfow
type VER_COND uint8

const (
	VER_COND_EQUAL         VER_COND = 1
	VER_COND_GREATER       VER_COND = 2
	VER_COND_GREATER_EQUAL VER_COND = 3
	VER_COND_LESS          VER_COND = 4
	VER_COND_LESS_EQUAL    VER_COND = 5

	VER_COND_AND VER_COND = 6
	VER_COND_OR  VER_COND = 7
)

// [OSVERSIONINFOEX] WSuiteMask. Includes values with VER_NT prefix.
//
// [OSVERSIONINFOEX]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-osversioninfoexw
type VER_SUITE uint16

const (
	VER_SUITE_BACKOFFICE               VER_SUITE = 0x0000_0004
	VER_SUITE_BLADE                    VER_SUITE = 0x0000_0400
	VER_SUITE_COMPUTE_SERVER           VER_SUITE = 0x0000_4000
	VER_SUITE_DATACENTER               VER_SUITE = 0x0000_0080
	VER_SUITE_ENTERPRISE               VER_SUITE = 0x0000_0002
	VER_SUITE_EMBEDDEDNT               VER_SUITE = 0x0000_0040
	VER_SUITE_PERSONAL                 VER_SUITE = 0x0000_0200
	VER_SUITE_SINGLEUSERTS             VER_SUITE = 0x0000_0100
	VER_SUITE_SMALLBUSINESS            VER_SUITE = 0x0000_0001
	VER_SUITE_SMALLBUSINESS_RESTRICTED VER_SUITE = 0x0000_0020
	VER_SUITE_STORAGE_SERVER           VER_SUITE = 0x0000_2000
	VER_SUITE_TERMINAL                 VER_SUITE = 0x0000_0010
	VER_SUITE_WH_SERVER                VER_SUITE = 0x0000_8000

	VER_SUITE_NT_DOMAIN_CONTROLLER VER_SUITE = 0x000_0002
	VER_SUITE_NT_SERVER            VER_SUITE = 0x000_0003
	VER_SUITE_NT_WORKSTATION       VER_SUITE = 0x000_0001
)

// [WaitForSingleObject] return value.
//
// [WaitForSingleObject]: https://learn.microsoft.com/en-us/windows/win32/api/synchapi/nf-synchapi-waitforsingleobject
type WAIT uint32

const (
	WAIT_ABANDONED WAIT = 0x0000_0080
	WAIT_OBJECT_0  WAIT = 0x0000_0000
	WAIT_TIMEOUT   WAIT = 0x0000_0102
	WAIT_FAILED    WAIT = 0xffff_ffff
)

// [IsWindowsVersionOrGreater] values; originally _WIN32_WINNT.
//
// [IsWindowsVersionOrGreater]: https://learn.microsoft.com/en-us/windows/win32/winprog/using-the-windows-headers
type WIN32_WINNT uint16

const (
	WIN32_WINNT_NT4          WIN32_WINNT = 0x0400
	WIN32_WINNT_WIN2K        WIN32_WINNT = 0x0500
	WIN32_WINNT_WINXP        WIN32_WINNT = 0x0501
	WIN32_WINNT_WS03         WIN32_WINNT = 0x0502
	WIN32_WINNT_WIN6         WIN32_WINNT = 0x0600
	WIN32_WINNT_VISTA        WIN32_WINNT = 0x0600
	WIN32_WINNT_WS08         WIN32_WINNT = 0x0600
	WIN32_WINNT_LONGHORN     WIN32_WINNT = 0x0600
	WIN32_WINNT_WIN7         WIN32_WINNT = 0x0601
	WIN32_WINNT_WIN8         WIN32_WINNT = 0x0602
	WIN32_WINNT_WINBLUE      WIN32_WINNT = 0x0603
	WIN32_WINNT_WINTHRESHOLD WIN32_WINNT = 0x0a00
	WIN32_WINNT_WIN10        WIN32_WINNT = 0x0a00
)
