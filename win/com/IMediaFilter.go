/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	baseIMediaFilter struct{ baseIPersist }

	// IMediaFilter > IPersist > IUnknown.
	IMediaFilter struct{ baseIMediaFilter }

	vtbIMediaFilter struct {
		vtbIPersist
		Stop          uintptr
		Pause         uintptr
		Run           uintptr
		GetState      uintptr
		SetSyncSource uintptr
		GetSyncSource uintptr
	}
)
