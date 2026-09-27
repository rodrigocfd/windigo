//go:build windows && (amd64 || arm64)

package ui

import (
	"github.com/rodrigocfd/windigo/wstr"
)

// Calls [wstr.FmtBytes], whose parameter is int on 64-bit.
func fmtBytes(numBytes uint64) string {
	return wstr.FmtBytes(int(numBytes))
}
