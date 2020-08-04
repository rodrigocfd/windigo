/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	_IMediaFilter struct{ _IPersist }

	// IMediaFilter > IPersist > IUnknown.
	IMediaFilter struct{ _IMediaFilter }

	_IMediaFilterVtbl struct {
		_IPersistVtbl
		Stop          uintptr
		Pause         uintptr
		Run           uintptr
		GetState      uintptr
		SetSyncSource uintptr
		GetSyncSource uintptr
	}
)
