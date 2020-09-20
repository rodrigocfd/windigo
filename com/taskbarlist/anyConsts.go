/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package taskbarlist

// ITaskbarList3.SetProgressState() tbpFlags.
type TBPF uint32

const (
	TBPF_NOPROGRESS    TBPF = 0
	TBPF_INDETERMINATE TBPF = 0x1
	TBPF_NORMAL        TBPF = 0x2
	TBPF_ERROR         TBPF = 0x4
	TBPF_PAUSED        TBPF = 0x8
)

// THUMBBUTTONMASK.
type THB uint32

const (
	THB_BITMAP  THB = 0x1
	THB_ICON    THB = 0x2
	THB_TOOLTIP THB = 0x4
	THB_FLAGS   THB = 0x8
)

// THUMBBUTTONFLAGS.
type THBF uint32

const (
	THBF_ENABLED        THBF = 0
	THBF_DISABLED       THBF = 0x1
	THBF_DISMISSONCLICK THBF = 0x2
	THBF_NOBACKGROUND   THBF = 0x4
	THBF_HIDDEN         THBF = 0x8
	THBF_NONINTERACTIVE THBF = 0x10
)
