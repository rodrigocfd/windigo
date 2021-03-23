[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/rodrigocfd/windigo)](https://github.com/rodrigocfd/windigo)
[![Lines of code](https://tokei.rs/b1/github/rodrigocfd/windigo)](https://github.com/rodrigocfd/windigo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# Windigo

Win32 API and GUI in idiomatic Go.

## Overview

The library is divided in the following packages:

| Package | Description |
| - | - |
| `ui` | High-level GUI wrappers for windows and controls. |
| `ui/wm` | High-level message parameters. |
| `win` | Native Win32 structs, handles and functions. |
| `win/co` | Native Win32 constants, all typed. |
| `win/com/dshow` | Native Win32 DirectShow COM interfaces. |
| `win/com/shell` | Native Win32 Shell COM interfaces. |
| `win/err` | Native Win32 error codes, all typed as `err.ERROR`. |

Windigo is designed to be familiar to Win32 programmers, using the same concepts, so most C/C++ Win32 tutorials should be applicable.

Windows and controls can be created in two ways:

* programmatically, by specifying the parameters used in the underlying [CreateWindowEx](https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-createwindowexw);
* by loading resources from a `.rc` file.

CGo is **not** used, just syscalls.

## Error treatment

The native Win32 functions deal with errors in two ways:

* Recoverable errors will return an `err.ERROR` value, which implements the [`error`](https://golang.org/pkg/builtin/#error) interface;

* Unrecoverable errors will simply **panic**. This avoids the excess of `if err != nil` with errors that cannot be recovered anyway, like internal Windows errors.

## Example

```go
package main

import (
    "github.com/rodrigocfd/windigo/win"
    "github.com/rodrigocfd/windigo/win/co"
)

func main() {
    hwnd := win.GetDesktopWindow()
    hwnd.MessageBox("Hello world", "Hello", co.MB_INFORMATION | co.MB_OK)
}
```

## License

Licensed under [MIT license](https://opensource.org/licenses/MIT), see [LICENSE.txt](LICENSE.txt) for details.
