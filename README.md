[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/rodrigocfd/windigo.svg)](https://github.com/rodrigocfd/windigo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Lines of code](https://tokei.rs/b1/github/rodrigocfd/windigo)](https://github.com/rodrigocfd/windigo)

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
    "fmt"

    "github.com/rodrigocfd/windigo/ui"
    "github.com/rodrigocfd/windigo/win"
    "github.com/rodrigocfd/windigo/win/co"
)

func main() {
    myWindow := NewMyWindow()
    myWindow.wnd.RunAsMain()
}

// This struct represents our main window.
type MyWindow struct {
    wnd     ui.WindowMain
    lblName ui.Static
    txtName ui.Edit
    btnShow ui.Button
}

// Creates a new instance of our main window.
func NewMyWindow() *MyWindow {
    wnd := ui.NewWindowMainOpts(ui.WindowMainOpts{
        Title:          "Hello world",
        ClientAreaSize: win.SIZE{Cx: 340, Cy: 80},
    })

    me := MyWindow{
        wnd: wnd,
        lblName: ui.NewStaticOpts(wnd, ui.StaticOpts{
            Text:     "Your name",
            Position: win.POINT{X: 10, Y: 21},
        }),
        txtName: ui.NewEditOpts(wnd, ui.EditOpts{
            Position: win.POINT{X: 80, Y: 20},
            Size:     win.SIZE{Cx: 150},
        }),
        btnShow: ui.NewButtonOpts(wnd, ui.ButtonOpts{
            Text:     "&Show",
            Position: win.POINT{X: 240, Y: 19},
        }),
    }

    me.btnShow.On().BnClicked(func() {
        me.wnd.Hwnd().MessageBox(
            fmt.Sprintf("Hello, %s!", me.txtName.Text()),
            "Saying hello", co.MB_ICONINFORMATION)
    })

    return &me
}
```

## License

Licensed under [MIT license](https://opensource.org/licenses/MIT), see [LICENSE.txt](LICENSE.txt) for details.
