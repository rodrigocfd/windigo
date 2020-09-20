/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package directshow

type (
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
	//
	// IMediaFilter > IPersist > IUnknown.
	IMediaFilter struct{ IPersist }

	IMediaFilterVtbl struct {
		IPersistVtbl
		Stop          uintptr
		Pause         uintptr
		Run           uintptr
		GetState      uintptr
		SetSyncSource uintptr
		GetSyncSource uintptr
	}
)
