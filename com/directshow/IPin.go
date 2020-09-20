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
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
	//
	// IPin > IUnknown.
	IPin struct{ win.IUnknown }

	IPinVtbl struct {
		win.IUnknownVtbl
		Connect                  uintptr
		ReceiveConnection        uintptr
		Disconnect               uintptr
		ConnectedTo              uintptr
		ConnectionMediaType      uintptr
		QueryPinInfo             uintptr
		QueryDirection           uintptr
		QueryId                  uintptr
		QueryAccept              uintptr
		EnumMediaTypes           uintptr
		QueryInternalConnections uintptr
		EndOfStream              uintptr
		BeginFlush               uintptr
		EndFlush                 uintptr
		NewSegment               uintptr
	}
)
