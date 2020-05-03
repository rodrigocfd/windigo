package consts

type GWLP int32

const (
	GWLP_EXSTYLE   GWLP = -20
	GWLP_HINSTANCE GWLP = -6
	GWLP_ID        GWLP = -12
	GWLP_STYLE     GWLP = -16
	GWLP_USERDATA  GWLP = -21
	GWLP_WNDPROC   GWLP = -4
)

type ID uint16

const (
	IDABORT    ID = 3
	IDCANCEL   ID = 2
	IDCONTINUE ID = 11
	IDIGNORE   ID = 5
	IDNO       ID = 7
	IDOK       ID = 1
	IDRETRY    ID = 4
	IDTRYAGAIN ID = 10
	IDYES      ID = 6
)

const LF_FACESIZE = 32

type LVP uint8

const (
	LVP_LISTITEM         LVP = 1
	LVP_LISTGROUP        LVP = 2
	LVP_LISTDETAIL       LVP = 3
	LVP_LISTSORTEDDETAIL LVP = 4
	LVP_EMPTYTEXT        LVP = 5
)
