//go:build windows

package errco

import (
	"fmt"
)

// Common dialog box [error codes].
//
// These error codes are unrelated to the ordinary system error codes,
// represented by the errco.ERROR type.
//
// [error codes]: https://learn.microsoft.com/en-us/windows/win32/api/commdlg/nf-commdlg-commdlgextendederror
type CDERR uint32

// Implements error interface.
func (err CDERR) Error() string {
	return err.String()
}

// Implements fmt.Stringer.
func (err CDERR) String() string {
	return fmt.Sprintf("[%d 0x%02x]", uint32(err), uint32(err))
}

const (
	CDERR_OK                  CDERR = 0
	CDERR_DIALOGFAILURE       CDERR = 0xffff
	CDERR_FINDRESFAILURE      CDERR = 0x0006
	CDERR_INITIALIZATION      CDERR = 0x0002
	CDERR_LOADRESFAILURE      CDERR = 0x0007
	CDERR_LOADSTRFAILURE      CDERR = 0x0005
	CDERR_LOCKRESFAILURE      CDERR = 0x0008
	CDERR_MEMALLOCFAILURE     CDERR = 0x0009
	CDERR_MEMLOCKFAILURE      CDERR = 0x000a
	CDERR_NOHINSTANCE         CDERR = 0x0004
	CDERR_NOHOOK              CDERR = 0x000b
	CDERR_NOTEMPLATE          CDERR = 0x0003
	CDERR_REGISTERMSGFAIL     CDERR = 0x000c
	CDERR_STRUCTSIZE          CDERR = 0x0001
	CDERR_PD_CREATEICFAILURE  CDERR = 0x100a
	CDERR_PD_DEFAULTDIFFERENT CDERR = 0x100c
	CDERR_PD_DNDMMISMATCH     CDERR = 0x1009
	CDERR_PD_GETDEVMODEFAIL   CDERR = 0x1005
	CDERR_PD_INITFAILURE      CDERR = 0x1006
	CDERR_PD_LOADDRVFAILURE   CDERR = 0x1004
	CDERR_PD_NODEFAULTPRN     CDERR = 0x1008
	CDERR_PD_NODEVICES        CDERR = 0x1007
	CDERR_PD_PARSEFAILURE     CDERR = 0x1002
	CDERR_PD_PRINTERNOTFOUND  CDERR = 0x100b
	CDERR_PD_RETDEFFAILURE    CDERR = 0x1003
	CDERR_PD_SETUPFAILURE     CDERR = 0x1001
	CDERR_CF_MAXLESSTHANMIN   CDERR = 0x2002
	CDERR_CF_NOFONTS          CDERR = 0x2001
	CDERR_FN_BUFFERTOOSMALL   CDERR = 0x3003
	CDERR_FN_INVALIDFILENAME  CDERR = 0x3002
	CDERR_FN_SUBCLASSFAILURE  CDERR = 0x3001
	CDERR_FR_BUFFERLENGTHZERO CDERR = 0x4001
)
