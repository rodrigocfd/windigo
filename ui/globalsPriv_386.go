//go:build windows

package ui

import (
	"github.com/rodrigocfd/windigo/wstr"
)

// Calls [wstr.FmtBytes], whose parameter is uint64 on 386.
func fmtBytes(numBytes uint64) string {
	return wstr.FmtBytes(numBytes)
}
