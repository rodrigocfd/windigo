//go:build windows

package ole

import (
	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win/co"
)

// A COM pointer, rooted in [IUnknown].
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
type ComPtr interface {
	// Returns the [IUnknown] virtual table.
	//
	// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	Ppvt() **vt.IUnknown

	// Calls [Release], then sets a new [IUnknown] virtual table.
	//
	// If you pass nil, you effectively release the object; the owning
	// ole.Releaser will simply do nothing.
	//
	// [Release]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
	// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
	Set(ppvt **vt.IUnknown)
}

// A constructible COM pointer, rooted in [IUnknown].
//
// Used in functions that instantiate COM pointers, like [CoCreateInstance] and
// [QueryInterface].
//
// [IUnknown]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nn-unknwn-iunknown
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
// [QueryInterface]: https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
type ComCtor[T any] interface {
	*T
	ComPtr
	IID() co.IID
}
