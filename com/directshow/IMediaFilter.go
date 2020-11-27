/**
 * Part of Windigo - Win32 API layer for Go
 * https://github.com/rodrigocfd/windigo
 * This library is released under the MIT license.
 */

package directshow

type (
	// IMediaFilter > IPersist > IUnknown.
	//
	// https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
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
