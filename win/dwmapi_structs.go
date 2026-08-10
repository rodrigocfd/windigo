//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/co"
)

// [DWM_BLURBEHIND] struct, with C memory layout.
//
// [DWM_BLURBEHIND]: https://learn.microsoft.com/en-us/windows/win32/api/dwmapi/ns-dwmapi-dwm_blurbehind
type DWM_BLURBEHIND struct {
	DwFlags                co.DWM_BB
	FEnable                BOOL
	HrgnBlur               HRGN
	FTransitionOnMaximized BOOL
}
