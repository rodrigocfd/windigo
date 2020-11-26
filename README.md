# Windigo

A thin Go layer over the Win32 API.

## Overview

Windigo aims to provide a solid foundation to build fast, native and scalable Win32 applications in Go. The windows can be built using raw API calls or loading dialog resources.

The library is composed of three packages:

* `co` – typed Win32 constants;
* `win` – Win32 structs, handles and functions;
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

func NewMyWindow() *MyWindow {
    opts := ui.NewOptsWindowMain()
    opts.Title = "Hello world"
    opts.Styles |= co.WS_MINIMIZEBOX

    wnd := ui.NewWindowMain(opts)

    return &MyMain{
        wnd:      wnd,
        btnHello: ui.NewButton(wnd),
    }
}

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