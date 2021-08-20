package co

// ACCELL fVirt.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-accel
type ACCELF uint8

const (
	ACCELF_NONE    ACCELF = 0
	ACCELF_VIRTKEY ACCELF = 1
	ACCELF_SHIFT   ACCELF = 0x04
	ACCELF_CONTROL ACCELF = 0x08
	ACCELF_ALT     ACCELF = 0x10
)

// Standard access rights. Also includes STANDARD_RIGHTS, and SPECIFIC_RIGHT prefixes.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/secauthz/standard-access-rights
type ACCESS_RIGHTS uint32

const (
	ACCESS_RIGHTS_NONE ACCESS_RIGHTS = 0

	ACCESS_RIGHTS_DELETE       ACCESS_RIGHTS = 0x00010000
	ACCESS_RIGHTS_READ_CONTROL ACCESS_RIGHTS = 0x00020000
	ACCESS_RIGHTS_WRITE_DAC    ACCESS_RIGHTS = 0x00040000
	ACCESS_RIGHTS_WRITE_OWNER  ACCESS_RIGHTS = 0x00080000
	ACCESS_RIGHTS_SYNCHRONIZE  ACCESS_RIGHTS = 0x00100000

	ACCESS_RIGHTS_STANDARD_REQUIRED ACCESS_RIGHTS = 0x000f0000
	ACCESS_RIGHTS_STANDARD_READ     ACCESS_RIGHTS = ACCESS_RIGHTS_READ_CONTROL
	ACCESS_RIGHTS_STANDARD_WRITE    ACCESS_RIGHTS = ACCESS_RIGHTS_READ_CONTROL
	ACCESS_RIGHTS_STANDARD_EXECUTE  ACCESS_RIGHTS = ACCESS_RIGHTS_READ_CONTROL
	ACCESS_RIGHTS_STANDARD_ALL      ACCESS_RIGHTS = 0x001f0000

	ACCESS_RIGHTS_SPECIFIC_ALL ACCESS_RIGHTS = 0x0000ffff
)

// NMTVASYNCDRAW dwRetFlags, don't seem to be defined anywhere, values are unconfirmed.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type ADRF uint32

const (
	ADRF_DRAWSYNC     ADRF = 0
	ADRF_DRAWNOTHING  ADRF = 1
	ADRF_DRAWFALLBACK ADRF = 2
	ADRF_DRAWIMAGE    ADRF = 3
)

// WM_APPCOMMAND command.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/inputdev/wm-appcommand
type APPCOMMAND int16

const (
	APPCOMMAND_BROWSER_BACKWARD                  APPCOMMAND = 1
	APPCOMMAND_BROWSER_FORWARD                   APPCOMMAND = 2
	APPCOMMAND_BROWSER_REFRESH                   APPCOMMAND = 3
	APPCOMMAND_BROWSER_STOP                      APPCOMMAND = 4
	APPCOMMAND_BROWSER_SEARCH                    APPCOMMAND = 5
	APPCOMMAND_BROWSER_FAVORITES                 APPCOMMAND = 6
	APPCOMMAND_BROWSER_HOME                      APPCOMMAND = 7
	APPCOMMAND_VOLUME_MUTE                       APPCOMMAND = 8
	APPCOMMAND_VOLUME_DOWN                       APPCOMMAND = 9
	APPCOMMAND_VOLUME_UP                         APPCOMMAND = 10
	APPCOMMAND_MEDIA_NEXTTRACK                   APPCOMMAND = 11
	APPCOMMAND_MEDIA_PREVIOUSTRACK               APPCOMMAND = 12
	APPCOMMAND_MEDIA_STOP                        APPCOMMAND = 13
	APPCOMMAND_MEDIA_PLAY_PAUSE                  APPCOMMAND = 14
	APPCOMMAND_LAUNCH_MAIL                       APPCOMMAND = 15
	APPCOMMAND_LAUNCH_MEDIA_SELECT               APPCOMMAND = 16
	APPCOMMAND_LAUNCH_APP1                       APPCOMMAND = 17
	APPCOMMAND_LAUNCH_APP2                       APPCOMMAND = 18
	APPCOMMAND_BASS_DOWN                         APPCOMMAND = 19
	APPCOMMAND_BASS_BOOST                        APPCOMMAND = 20
	APPCOMMAND_BASS_UP                           APPCOMMAND = 21
	APPCOMMAND_TREBLE_DOWN                       APPCOMMAND = 22
	APPCOMMAND_TREBLE_UP                         APPCOMMAND = 23
	APPCOMMAND_MICROPHONE_VOLUME_MUTE            APPCOMMAND = 24
	APPCOMMAND_MICROPHONE_VOLUME_DOWN            APPCOMMAND = 25
	APPCOMMAND_MICROPHONE_VOLUME_UP              APPCOMMAND = 26
	APPCOMMAND_HELP                              APPCOMMAND = 27
	APPCOMMAND_FIND                              APPCOMMAND = 28
	APPCOMMAND_NEW                               APPCOMMAND = 29
	APPCOMMAND_OPEN                              APPCOMMAND = 30
	APPCOMMAND_CLOSE                             APPCOMMAND = 31
	APPCOMMAND_SAVE                              APPCOMMAND = 32
	APPCOMMAND_PRINT                             APPCOMMAND = 33
	APPCOMMAND_UNDO                              APPCOMMAND = 34
	APPCOMMAND_REDO                              APPCOMMAND = 35
	APPCOMMAND_COPY                              APPCOMMAND = 36
	APPCOMMAND_CUT                               APPCOMMAND = 37
	APPCOMMAND_PASTE                             APPCOMMAND = 38
	APPCOMMAND_REPLY_TO_MAIL                     APPCOMMAND = 39
	APPCOMMAND_FORWARD_MAIL                      APPCOMMAND = 40
	APPCOMMAND_SEND_MAIL                         APPCOMMAND = 41
	APPCOMMAND_SPELL_CHECK                       APPCOMMAND = 42
	APPCOMMAND_DICTATE_OR_COMMAND_CONTROL_TOGGLE APPCOMMAND = 43
	APPCOMMAND_MIC_ON_OFF_TOGGLE                 APPCOMMAND = 44
	APPCOMMAND_CORRECTION_LIST                   APPCOMMAND = 45
	APPCOMMAND_MEDIA_PLAY                        APPCOMMAND = 46
	APPCOMMAND_MEDIA_PAUSE                       APPCOMMAND = 47
	APPCOMMAND_MEDIA_RECORD                      APPCOMMAND = 48
	APPCOMMAND_MEDIA_FAST_FORWARD                APPCOMMAND = 49
	APPCOMMAND_MEDIA_REWIND                      APPCOMMAND = 50
	APPCOMMAND_MEDIA_CHANNEL_UP                  APPCOMMAND = 51
	APPCOMMAND_MEDIA_CHANNEL_DOWN                APPCOMMAND = 52
	APPCOMMAND_DELETE                            APPCOMMAND = 53
	APPCOMMAND_DWM_FLIP3D                        APPCOMMAND = 54
)

// Button control notifications, sent via WM_NOTIFY.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
const (
	_BCN_FIRST NM = -1250

	BCN_HOTITEMCHANGE NM = _BCN_FIRST + 0x0001
	BCN_DROPDOWN      NM = _BCN_FIRST + 0x0002
)

// BITMAPINFOHEADER biCompression.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BI uint32

const (
	BI_RGB       BI = 0
	BI_RLE8      BI = 1
	BI_RLE4      BI = 2
	BI_BITFIELDS BI = 3
	BI_JPEG      BI = 4
	BI_PNG       BI = 5
)

// SetBkMode() mode. Originally has no prefix.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-setbkmode
type BKMODE int32

const (
	BKMODE_TRANSPARENT BKMODE = 1
	BKMODE_OPAQUE      BKMODE = 2
)

// Button control notifications, sent via WM_COMMAND.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/bumper-button-control-reference-notifications
const (
	BN_CLICKED       CMD = 0
	BN_PAINT         CMD = 1
	BN_HILITE        CMD = 2
	BN_UNHILITE      CMD = 3
	BN_DISABLE       CMD = 4
	BN_DOUBLECLICKED CMD = 5
	BN_PUSHED        CMD = BN_HILITE
	BN_UNPUSHED      CMD = BN_UNHILITE
	BN_DBLCLK        CMD = BN_DOUBLECLICKED
	BN_SETFOCUS      CMD = 6
	BN_KILLFOCUS     CMD = 7
)

// Button control styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/button-styles
type BS WS

const (
	BS_PUSHBUTTON      BS = 0x00000000
	BS_DEFPUSHBUTTON   BS = 0x00000001
	BS_CHECKBOX        BS = 0x00000002
	BS_AUTOCHECKBOX    BS = 0x00000003
	BS_RADIOBUTTON     BS = 0x00000004
	BS_3STATE          BS = 0x00000005
	BS_AUTO3STATE      BS = 0x00000006
	BS_GROUPBOX        BS = 0x00000007
	BS_USERBUTTON      BS = 0x00000008
	BS_AUTORADIOBUTTON BS = 0x00000009
	BS_PUSHBOX         BS = 0x0000000a
	BS_OWNERDRAW       BS = 0x0000000b
	BS_TYPEMASK        BS = 0x0000000f
	BS_LEFTTEXT        BS = 0x00000020
	BS_TEXT            BS = 0x00000000
	BS_ICON            BS = 0x00000040
	BS_BITMAP          BS = 0x00000080
	BS_LEFT            BS = 0x00000100
	BS_RIGHT           BS = 0x00000200
	BS_CENTER          BS = 0x00000300
	BS_TOP             BS = 0x00000400
	BS_BOTTOM          BS = 0x00000800
	BS_VCENTER         BS = 0x00000c00
	BS_PUSHLIKE        BS = 0x00001000
	BS_MULTILINE       BS = 0x00002000
	BS_NOTIFY          BS = 0x00004000 // Button will send BN_DISABLE, BN_PUSHED, BN_KILLFOCUS, BN_PAINT, BN_SETFOCUS, and BN_UNPUSHED notifications.
	BS_FLAT            BS = 0x00008000
	BS_RIGHTBUTTON     BS = BS_LEFTTEXT
)

// IsDlgButtonChecked() return value, among others.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-isdlgbuttonchecked
type BST uint32

const (
	BST_UNCHECKED     BST = 0x0000
	BST_CHECKED       BST = 0x0001
	BST_INDETERMINATE BST = 0x0002
	BST_PUSHED        BST = 0x0004
	BST_FOCUS         BST = 0x0008
)

// Toolbar button styles.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/controls/toolbar-control-and-button-styles
type BTNS uint8

const (
	BTNS_BUTTON        BTNS = BTNS(TBSTYLE_BUTTON)
	BTNS_SEP           BTNS = BTNS(TBSTYLE_SEP)
	BTNS_CHECK         BTNS = BTNS(TBSTYLE_CHECK)
	BTNS_GROUP         BTNS = BTNS(TBSTYLE_GROUP)
	BTNS_CHECKGROUP    BTNS = BTNS(TBSTYLE_CHECKGROUP)
	BTNS_DROPDOWN      BTNS = BTNS(TBSTYLE_DROPDOWN)
	BTNS_AUTOSIZE      BTNS = BTNS(TBSTYLE_AUTOSIZE)
	BTNS_NOPREFIX      BTNS = BTNS(TBSTYLE_NOPREFIX)
	BTNS_SHOWTEXT      BTNS = 0x0040
	BTNS_WHOLEDROPDOWN BTNS = 0x0080
)
