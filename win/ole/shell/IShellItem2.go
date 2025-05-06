//go:build windows

package shell

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// [IShellItem2] COM interface.
//
// / Usually created with [SHCreateItemFromParsingName].
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
// [SHCreateItemFromParsingName]: https://learn.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-shcreateitemfromparsingname
type IShellItem2 struct{ IShellItem }

// Returns the unique COM interface identifier.
func (*IShellItem2) IID() co.IID {
	return co.IID_IShellItem2
}
