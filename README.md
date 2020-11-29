# Windigo

A thin Go layer over the Win32 API.

## Overview

Windigo aims to provide a solid foundation to build fast, native and scalable Win32 applications in Go. The windows can be built using raw API calls or loading dialog resources.

The library is composed of 4 packages:

* `co` – typed native Win32 constants;
* `com` – subpackages with native Win32 COM interfaces;
* `win` – native Win32 structs, handles and functions;
* `ui` – Windigo high level wrappers.

It does **not** use CGo anywhere.

Windigo is designed to be familiar to Win32 programmers, using the same concepts, so most C/C++ Win32 tutorials should be applicable. The `ui` package is heavily based on [WinLamb](https://github.com/rodrigocfd/winlamb) C++ library.

Since raw Win32 API is exposed, there are no limits: you can do everything. But you can also shoot yourself in the foot, so please always refer to the [official Win32 documentation](https://docs.microsoft.com/en-us/windows/win32/).

## Example

```go
package main

import (
    "windigo/co"
    "windigo/ui"
    "windigo/win"
)

func main() {
    myWnd := NewMyWindow()
    myWnd.Run()
}

// Struct of our main window.
type MyWindow struct {
    wnd      *ui.WindowMain
    btnHello *ui.Button
}

// Constructor of our main window.
func NewMyWindow() *MyWindow {
    wnd := ui.NewWindowMain(
        &ui.WindowMainOpts{
            Title:     "Hello world",
            StylesAdd: co.WS_MINIMIZEBOX,
        },
    )

    return &MyMain{
        wnd:      wnd,
        btnHello: ui.NewButton(wnd),
    }
}

// Runs our main window. Returns only after the window is closed.
func (me *MyWindow) Run() int {
    me.events()
    return me.wnd.RunAsMain()
}

func (me *MyWindow) events() {
    me.wnd.On().WmCreate(func(_ *win.CREATESTRUCT) int {
        me.btnHello.Create("Click me", ui.Pos{X: 10, Y: 10}, 90, co.BS_DEFPUSHBUTTON)
        return 0
    })

    me.btnHello.On().BnClicked(func() {
        ui.SysDlg.MsgBox(me.wnd, "Hi", "Hello world!", co.MB_ICONINFORMATION)
    })
}
```

## Win32 error handling

Native Win32 API calls may return a `win.WinError` type, which implements `error` interface.

However, in Windigo, most Win32 functions do **not** return errors. That's because most low-level errors are **unrecoverable**, in the sense that if such an error happens in an application, there's really nothing you can do. Unrecoverable errors very rare, occurring in conditions like low memory or internal Windows crashes.

So, in order to keep the API simple, instead of returning lots of errors, unrecoverable errors will simply panic.

### Built-in panic treatment

Windigo will recover all panics at the top of the stack, displaying the error and the stack trace in a [MessageBox](https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw), in order to make debugging easier.

However, if the panic happens in a non-GUI thread, it can't be recovered, and no MessageBox will be displayed.