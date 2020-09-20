/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package directshow

import (
	"windigo/win"
)

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch
	//
	// IDispatch > IUnknown.
	IDispatch struct{ win.IUnknown }

	_IDispatchVtbl struct {
		win.IUnknownVtbl
		GetTypeInfoCount uintptr
		GetTypeInfo      uintptr
		GetIDsOfNames    uintptr
		Invoke           uintptr
	}
)
