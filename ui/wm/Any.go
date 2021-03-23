package wm

import (
	"github.com/rodrigocfd/windigo/win"
)

// Raw message parameters to any message: WPARAM and LPARAM.
type Any struct {
	WParam win.WPARAM
	LParam win.LPARAM
}
