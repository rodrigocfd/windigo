/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type IPersist struct {
	IUnknown
}

type iPersistVtbl struct {
	iUnknownVtbl
	GetClassID uintptr
}
