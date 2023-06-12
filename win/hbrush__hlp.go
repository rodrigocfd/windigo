//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// This helper method performs the tricky conversion to create a brush from a
// system color, particularly used when registering a window class.
func CreateSysColorBrush(sysColor co.COLOR) HBRUSH {
	return HBRUSH(sysColor + 1)
}
