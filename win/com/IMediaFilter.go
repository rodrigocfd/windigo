/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// IMediaFilter > IPersist > IUnknown.
type IMediaFilter struct {
	IPersist
}

type iMediaFilterVtbl struct {
	iPersistVtbl
	Stop          uintptr
	Pause         uintptr
	Run           uintptr
	GetState      uintptr
	SetSyncSource uintptr
	GetSyncSource uintptr
}
