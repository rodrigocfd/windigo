//go:build windows

package errco

import (
	"fmt"
)

// DDE error codes.
//
// These error codes are unrelated to the ordinary system error codes,
// represented by the errco.ERROR type.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddegetlasterror
type DMLERR uint32

// Implements error interface.
func (err DMLERR) Error() string {
	return err.String()
}

// Implements fmt.Stringer.
func (err DMLERR) String() string {
	return fmt.Sprintf("[%d 0x%02x]", uint32(err), uint32(err))
}

const (
	DMLERR_NO_ERROR            DMLERR = 0
	DMLERR_ADVACKTIMEOUT       DMLERR = 0x4000
	DMLERR_BUSY                DMLERR = 0x4001
	DMLERR_DATAACKTIMEOUT      DMLERR = 0x4002
	DMLERR_DLL_NOT_INITIALIZED DMLERR = 0x4003
	DMLERR_DLL_USAGE           DMLERR = 0x4004
	DMLERR_EXECACKTIMEOUT      DMLERR = 0x4005
	DMLERR_INVALIDPARAMETER    DMLERR = 0x4006
	DMLERR_LOW_MEMORY          DMLERR = 0x4007
	DMLERR_MEMORY_ERROR        DMLERR = 0x4008
	DMLERR_NOTPROCESSED        DMLERR = 0x4009
	DMLERR_NO_CONV_ESTABLISHED DMLERR = 0x400a
	DMLERR_POKEACKTIMEOUT      DMLERR = 0x400b
	DMLERR_POSTMSG_FAILED      DMLERR = 0x400c
	DMLERR_REENTRANCY          DMLERR = 0x400d
	DMLERR_SERVER_DIED         DMLERR = 0x400e
	DMLERR_SYS_ERROR           DMLERR = 0x400f
	DMLERR_UNADVACKTIMEOUT     DMLERR = 0x4010
	DMLERR_UNFOUND_QUEUE_ID    DMLERR = 0x4011
)
