//go:build windows

package shell

import (
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/ole"
)

// [IShellFolder] COM interface.
//
// [IShellFolder]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellfolder
type IShellFolder struct{ ole.IUnknown }

// Returns the unique COM interface identifier.
func (*IShellFolder) IID() co.IID {
	return co.IID_IShellFolder
}
