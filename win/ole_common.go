//go:build windows

package win

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
)

// A [COM] object whose lifetime can be managed by an [OleReleaser], automating
// the cleanup.
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
type OleResource interface {
	release()
}

// A [COM] object, derived from [IUnknown].
//
// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type OleObj interface {
	OleResource

	// Returns the unique [COM] [interface ID].
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	// [interface ID]: https://learn.microsoft.com/en-us/office/client-developer/outlook/mapi/iid
	IID() co.IID

	// Returns the [COM] virtual table pointer.
	//
	// This is a low-level method, used internally by the library. Incorrect usage
	// may lead to segmentation faults.
	//
	// [COM]: https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal
	Ppvt() **_IUnknownVt
}

// Returns the virtual table pointer, performing a nil check.
func ppvtOrNil(obj OleObj) unsafe.Pointer {
	if !utl.IsNil(obj) {
		return unsafe.Pointer(obj.Ppvt())
	}
	return nil
}
