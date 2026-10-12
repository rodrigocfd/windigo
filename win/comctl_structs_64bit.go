//go:build windows && (amd64 || arm64)

package win

import (
	"github.com/rodrigocfd/windigo/co"
)

// [TBBUTTON] struct, with C memory layout.
//
// [TBBUTTON]: https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-tbbutton
type TBBUTTON struct {
	IBitmap   int32 // With multiple image lists, HIWORD is the image list index.
	IdCommand int32
	FsState   co.TBSTATE
	FsStyle   co.BTNS
	bReserved [6]uint8 // This padding is 6 in 64-bit environments.
	DwData    uintptr
	IString   *uint16 // Convert to/from string with [wstr.DecodePtr] and [wstr.EncodeToPtr]; can also be the index in the string list.
}
