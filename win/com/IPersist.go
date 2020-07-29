/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

type (
	baseIPersist struct{ baseIUnknown }

	// IPersist > IUnknown.
	IPersist struct{ baseIPersist }

	vtbIPersist struct {
		vtbIUnknown
		GetClassID uintptr
	}
)
