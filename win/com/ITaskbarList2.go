/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// ITaskbarList2 > ITaskbarList > IUnknown.
type ITaskbarList2 struct {
	ITaskbarList
}

type iTaskbarList2Vtbl struct {
	iTaskbarListVtbl
	MarkFullscreenWindow uintptr
}
