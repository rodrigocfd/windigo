//go:build windows

package co

// [CHOOSECOLOR] Flags.
//
// [CHOOSECOLOR]: https://learn.microsoft.com/en-us/windows/win32/api/commdlg/ns-commdlg-choosecolorw-r1
type CC uint32

const (
	CC_RGBINIT              CC = 0x0000_0001
	CC_FULLOPEN             CC = 0x0000_0002
	CC_PREVENTFULLOPEN      CC = 0x0000_0004
	CC_SHOWHELP             CC = 0x0000_0008
	CC_ENABLEHOOK           CC = 0x0000_0010
	CC_ENABLETEMPLATE       CC = 0x0000_0020
	CC_ENABLETEMPLATEHANDLE CC = 0x0000_0040
	CC_SOLIDCOLOR           CC = 0x0000_0080
	CC_ANYCOLOR             CC = 0x0000_0100
)
