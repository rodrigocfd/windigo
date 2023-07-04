//go:build windows

package errco

import (
	"fmt"
)

// DDE [error codes].
//
// These error codes are unrelated to the ordinary system error codes,
// represented by the errco.ERROR type.
//
// [error codes]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddegetlasterror
type DMLERR uint32

// Implements error interface.
func (err DMLERR) Error() string {
	return err.String()
}

// Implements fmt.Stringer.
func (err DMLERR) String() string {
	descr := map[DMLERR]string{
		DMLERR_NO_ERROR:            "NO_ERROR",
		DMLERR_ADVACKTIMEOUT:       "ADVACKTIMEOUT",
		DMLERR_BUSY:                "BUSY",
		DMLERR_DATAACKTIMEOUT:      "DATAACKTIMEOUT",
		DMLERR_DLL_NOT_INITIALIZED: "DLL_NOT_INITIALIZED",
		DMLERR_DLL_USAGE:           "DLL_USAGE",
		DMLERR_EXECACKTIMEOUT:      "EXECACKTIMEOUT",
		DMLERR_INVALIDPARAMETER:    "INVALIDPARAMETER",
		DMLERR_LOW_MEMORY:          "LOW_MEMORY",
		DMLERR_MEMORY_ERROR:        "MEMORY_ERROR",
		DMLERR_NOTPROCESSED:        "NOTPROCESSED",
		DMLERR_NO_CONV_ESTABLISHED: "NO_CONV_ESTABLISHED",
		DMLERR_POKEACKTIMEOUT:      "POKEACKTIMEOUT",
		DMLERR_POSTMSG_FAILED:      "POSTMSG_FAILED",
		DMLERR_REENTRANCY:          "REENTRANCY",
		DMLERR_SERVER_DIED:         "SERVER_DIED",
		DMLERR_SYS_ERROR:           "SYS_ERROR",
		DMLERR_UNADVACKTIMEOUT:     "UNADVACKTIMEOUT",
		DMLERR_UNFOUND_QUEUE_ID:    "UNFOUND_QUEUE_ID",
	}
	return fmt.Sprintf("[%d 0x%02x] %s", uint32(err), uint32(err), descr[err])
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
