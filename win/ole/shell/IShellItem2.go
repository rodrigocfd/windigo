//go:build windows

package shell

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// [IShellItem2] COM interface.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	ish, _ := shell.SHCreateItemFromParsingName[shell.IShellItem2](
//		rel, "C:\\Temp\\foo.txt")
//
// [IShellItem2]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-ishellitem2
type IShellItem2 struct{ IShellItem }

// Returns the unique COM interface identifier.
func (*IShellItem2) IID() co.IID {
	return co.IID_IShellItem2
}
