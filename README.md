# Windigo

A thin Go layer over the Win32 API.

## Overview

Windigo is composed of three packages:

* `co` – typed Win32 constants;
* `win` – Win32 structs, handles and functions;
* `ui` – high level wrappers.

Windigo aims to provide a solid foundation to build fast, native and scalable Win32 applications in Go.

It is designed to be familiar to Win32 programmers, using the same concepts, so most C/C++ Win32 tutorials should be applicable. The `ui` package is heavily based on [WinLamb](https://github.com/rodrigocfd/winlamb) C++ library.

Since raw Win32 API is exposed, there are no limits: you can do everything. But you can also shoot yourself in the foot, so please always refer to the [official Win32 documentation](https://docs.microsoft.com/en-us/windows/win32/).

## Example

```go
package main

import (
    "windigo/co"
    "windigo/ui"
)

func main() {
    w := MyWindow{}
    w.RunThisThing()
}

// We implement our window as a struct, which contains a ui.WindowMain member,
// responsible by window creation and management.
// We also have a button, which we will create during WM_CREATE event.
type MyWindow struct {
    wnd      ui.WindowMain
    btnHello ui.Button
}

const (
    // Here we define a constant to identify our button.
    ID_BTN_HELLO int32 = iota + 1000
)

func (me *MyWindow) RunThisThing() {
    // Here we define some initial parameters of our window.
    // There are many others, and they're all optional.
    me.wnd.Setup().Title = "This is the title"
    me.wnd.Setup().Style |= co.WS_MINIMIZEBOX

    // WM_CREATE event is handled with a closure.
    // https://docs.microsoft.com/en-us/windows/win32/winmsg/wm-create
    me.wnd.On().WmCreate(func(p *win.CREATESTRUCT) int32 {
        // Physically create the button.
        // The last 3 arguments are: left position, top position and width.
        me.btnHello.CreateSimpleDef(&me.wnd, ID_BTN_HELLO, 10, 10, 90)
        return 0
    })

    // The button click is handled in the WM_COMMAND event.
    // https://docs.microsoft.com/en-us/windows/win32/menurc/wm-command
    me.wnd.On().WmCommand(ID_BTN_HELLO, func(p ui.WmCommand) {
        // This is the action we execute: show a popup message box.
        // The Hwnd() method returns the HWND handle of our window, which gives
        // us access to all Win32 functions executed on HWNDs.
        me.wnd.Hwnd().MessageBox("Hello world", "Hi", co.MB_ICONINFORMATION)
    })

    // Finally run our main window.
    // This method will block until the window is closed.
    return me.wnd.RunAsMain()
}
```