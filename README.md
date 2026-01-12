[![Go Reference](https://pkg.go.dev/badge/github.com/rodrigocfd/windigo.svg)](https://pkg.go.dev/github.com/rodrigocfd/windigo)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/rodrigocfd/windigo?style=flat-square&color=03a7ed)](https://github.com/rodrigocfd/windigo)
[![Go Report Card](https://goreportcard.com/badge/github.com/rodrigocfd/windigo?style=flat-square)](https://goreportcard.com/report/github.com/rodrigocfd/windigo)
[![MIT License](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square)](https://github.com/rodrigocfd/windigo/blob/master/LICENSE.md)
[![Lines of code](https://img.shields.io/badge/LoC-51.5k-blue?style=flat-square)](https://github.com/rodrigocfd/windigo)

# Windigo

Win32 API and GUI in idiomatic Go.

Windigo is designed to be familiar to C/C++ Win32 programmers, using the same concepts, and an API as close as possible to the original Win32 API. This allows most C/C++ Win32 tutorials and examples to be translated to Go.

Notably, Windigo is written 100% in pure Go – CGo is **not** used, just native syscalls. 

## Install

```
go get -u github.com/rodrigocfd/windigo
```

## Current status

Native FFI items currently implemented:

| Native FFI item | Count |
| - | -: |
| Handles | 31 |
| Functions | 632 |
| COM interfaces | 47 |
| COM methods | 294 |
| Constants | 10,869 |

## Examples

*In the examples below, error checking is ommited for brevity.*

<details>
<summary>GUI window</summary>

### GUI window

The example below creates a window programmatically, and handles the button click. Also, it uses the `minimal.syso` provided in the [_resources](_resources/) folder.

![Screen capture](example.gif)

```go
package main

import (
	"fmt"
	"runtime"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/ui"
)

func main() {
	runtime.LockOSThread() // important: Windows GUI is single-threaded

	ShowMainWindow()
}

// This struct represents our main window.
type MyWindow struct {
	wnd     *ui.Main
	lblName *ui.Static
	txtName *ui.Edit
	btnShow *ui.Button
}

// Displays the main window, blocking until it is closed.
func ShowMainWindow() int {
	wnd := ui.NewMain( // create the main window
		ui.OptsMain().
			Title("Hello you").
			Size(ui.Dpi(340, 80)).
			ClassIconId(101), // ID of icon resource, see _resources folder
	)

	lblName := ui.NewStatic( // create the child controls
		wnd,
		ui.OptsStatic().
			Text("Your name").
			Position(ui.Dpi(10, 22)),
	)
	txtName := ui.NewEdit(
		wnd,
		ui.OptsEdit().
			Position(ui.Dpi(80, 20)).
			Width(ui.DpiX(150)),
	)
	btnShow := ui.NewButton(
		wnd,
		ui.OptsButton().
			Text("&Show").
			Position(ui.Dpi(240, 19)),
	)

	me := &MyWindow{wnd, lblName, txtName, btnShow}
	me.events()
	return wnd.RunAsMain()
}

func (me *MyWindow) events() {
	me.btnShow.On().BnClicked(func() {
		msg := fmt.Sprintf("Hello, %s!", me.txtName.Text())
		me.wnd.Hwnd().MessageBox(msg, "Saying hello", co.MB_ICONINFORMATION)
	})
}
```

To compile the final `.exe` file, run the command:

```
go build -ldflags "-s -w -H=windowsgui"
```
</details>

<details>
<summary>Registry access</summary>

### Registry access

```go
package main

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

func main() {
	// Open a registry key

	hKey, _ := win.HKEY_CURRENT_USER.RegOpenKeyEx(
		"Control Panel\\Mouse",
		co.REG_OPTION_NONE,
		co.KEY_READ) // open key as read-only
	defer hKey.RegCloseKey()

	// Read a single value from this key

	regVal, _ := hKey.RegQueryValueEx("Beep") // data can be string, uint32, etc.

	if strVal, ok := regVal.Sz(); ok { // try to extract a string value
		println("Beep is", strVal)
	}

	// Enumerate all values under this key

	namesVals, _ := hKey.RegEnumValue()
	for _, nameVal := range namesVals {
		if str, ok := nameVal.Val.Sz(); ok { // does it contain a string?
			println("Value str", nameVal.Name, str)
		} else if num, ok := nameVal.Val.Dword(); ok { // does it contain an uint32?
			println("Value int", nameVal.Name, num)
		} else {
			println("Value other", nameVal.Name)
		}
	}
}
```
</details>

<details>
<summary>Enumerating running processes</summary>

### Enumerating running processes

The example below takes a [process snapshot](https://learn.microsoft.com/en-us/windows/win32/toolhelp/taking-a-snapshot-and-viewing-processes) to list the running processes:

```go
package main

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

func main() {
	hSnap, _ := win.CreateToolhelp32Snapshot(co.TH32CS_SNAPPROCESS, 0)
	defer hSnap.CloseHandle()

	processes, _ := hSnap.EnumProcesses()
	for _, nfo := range processes {
		println("PID:", nfo.Th32ProcessID, "name:", nfo.SzExeFile())
	}

	println(len(processes), "found")
}
```
</details>

<details>
<summary>Taking a screenshot</summary>

### Taking a screenshot

This complex example takes a screenshot using [GDI](https://learn.microsoft.com/en-us/windows/win32/gdi/windows-gdi) and saves it to a BMP file.

```go
package main

import (
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

func main() {
	cxScreen := win.GetSystemMetrics(co.SM_CXSCREEN)
	cyScreen := win.GetSystemMetrics(co.SM_CYSCREEN)

	hdcScreen, _ := win.HWND(0).GetDC()
	defer win.HWND(0).ReleaseDC(hdcScreen)

	hBmp, _ := hdcScreen.CreateCompatibleBitmap(int(cxScreen), int(cyScreen))
	defer hBmp.DeleteObject()

	hdcMem, _ := hdcScreen.CreateCompatibleDC()
	defer hdcMem.DeleteDC()

	hBmpOld, _ := hdcMem.SelectObjectBmp(hBmp)
	defer hdcMem.SelectObjectBmp(hBmpOld)

	_ = hdcMem.BitBlt(
		win.POINT{X: 0, Y: 0},
		win.SIZE{Cx: cxScreen, Cy: cyScreen},
		hdcScreen,
		win.POINT{X: 0, Y: 0},
		co.ROP_SRCCOPY,
	)

	bi := win.BITMAPINFO{
		BmiHeader: win.BITMAPINFOHEADER{
			Width:       cxScreen,
			Height:      cyScreen,
			Planes:      1,
			BitCount:    32,
			Compression: co.BI_RGB,
		},
	}
	bi.BmiHeader.SetSize()

	bmpObj, _ := hBmp.GetObject()
	bmpSize := bmpObj.CalcBitmapSize(bi.BmiHeader.BitCount)

	rawMem, _ := win.GlobalAlloc(co.GMEM_FIXED|co.GMEM_ZEROINIT, bmpSize)
	defer rawMem.GlobalFree()

	bmpSlice, _ := rawMem.GlobalLockSlice()
	defer rawMem.GlobalUnlock()

	_, _ = hdcScreen.GetDIBits(
		hBmp,
		0,
		int(cyScreen),
		bmpSlice,
		&bi,
		co.DIB_COLORS_RGB,
	)

	var bfh win.BITMAPFILEHEADER
	bfh.SetBfType()
	bfh.SetBfOffBits(uint32(unsafe.Sizeof(bfh) + unsafe.Sizeof(bi.BmiHeader)))
	bfh.SetBfSize(bfh.BfOffBits() + uint32(bmpSize))

	fout, _ := win.FileOpen(
		"C:\\Temp\\screenshot.bmp",
		co.FOPEN_RW_OPEN_OR_CREATE,
	)
	defer fout.Close()

	_, _ = fout.Write(bfh.Serialize())
	_, _ = fout.Write(bi.BmiHeader.Serialize())
	_, _ = fout.Write(bmpSlice)
}
```
</details>

<details>
<summary>Component Object Model (COM)</summary>

### Component Object Model (COM)

Windigo has full support for C++ [COM](https://learn.microsoft.com/en-us/windows/win32/com/component-object-model--com--portal) objects. The cleanup is performed by a `win.OleReleaser` object, which calls [`Release`](https://learn.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release) on multiple COM objects at once, much like an arena allocator. Every function which produces a COM object requires a `win.OleReleaser` to take care of its lifetime.

The example below uses COM objects to display the system native [Open File](https://learn.microsoft.com/en-us/windows/win32/learnwin32/example--the-open-dialog-box) window:

```go
package main

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

func main() {
	runtime.LockOSThread() // important: Windows GUI is single-threaded

	_, _ := win.CoInitializeEx(
		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
	defer win.CoUninitialize()

	releaser := win.NewOleReleaser() // will release all COM objects created here
	defer releaser.Release()

	var fod *win.IFileOpenDialog
	_ = win.CoCreateInstance(
		releaser,
		co.CLSID_FileOpenDialog,
		nil,
		co.CLSCTX_INPROC_SERVER,
		&fod,
	)

	defOpts, _ := fod.GetOptions()
	_ = fod.SetOptions(defOpts |
		co.FOS_FORCEFILESYSTEM |
		co.FOS_FILEMUSTEXIST,
	)

	_ = fod.SetFileTypes([]win.COMDLG_FILTERSPEC{
		{Name: "Text files", Spec: "*.txt"},
		{Name: "All files", Spec: "*.*"},
	})
	_ = fod.SetFileTypeIndex(1)

	if ok, _ := fod.Show(win.HWND(0)); ok { // in real applications, pass the parent HWND
		item, _ := fod.GetResult(releaser)
		fileName, _ := item.GetDisplayName(co.SIGDN_FILESYSPATH)
		println(fileName)
	}
}
```
</details>

<details>
<summary>COM Automation</summary>

### COM Automation

Windigo has bindings to [`IDispatch`](https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nn-oaidl-idispatch) COM interface and [`VARIANT`](https://learn.microsoft.com/en-us/windows/win32/api/oaidl/ns-oaidl-variant) parameters, allowing you to [invoke](https://learn.microsoft.com/en-us/windows/win32/api/oaidl/nf-oaidl-idispatch-invoke) Automation methods.

The example below manipulates an Excel spreadsheet, saving a copy of it:

```go
package main

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/win"
)

func main() {
	_, _ = win.CoInitializeEx(
		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
	defer win.CoUninitialize()

	rel := win.NewOleReleaser()
	defer rel.Release()

	clsId, _ := win.CLSIDFromProgID("Excel.Application")

	var excel *win.IDispatch
	_ = win.CoCreateInstance(
		rel,
		clsId,
		nil,
		co.CLSCTX_LOCAL_SERVER,
		&excel,
	)

	books, _ := excel.InvokeGetIDispatch(rel, "Workbooks")
	file, _ := books.InvokeMethodIDispatch(rel, "Open", "C:\\Temp\\foo.xlsx")
	_, _ = file.InvokeMethod(rel, "SaveAs", "C:\\Temp\\foo copy.xlsx")
	_, _ = file.InvokeMethod(rel, "Close")
}
```
</details>

## Architecture

The library is divided in four packages:

| Package | Description |
| - | - |
| [`co`](https://pkg.go.dev/github.com/rodrigocfd/windigo/co) | Native Win32 constants, all typed. |
| [`ui`](https://pkg.go.dev/github.com/rodrigocfd/windigo/ui) | High-level UI windows and controls. |
| [`win`](https://pkg.go.dev/github.com/rodrigocfd/windigo/win) | Native Win32 structs, handles and functions. |
| [`wstr`](https://pkg.go.dev/github.com/rodrigocfd/windigo/wstr) | String and UTF-16 wide string management. |

Package dependency:

```mermaid
flowchart BT
    internal/utl([internal/utl]) --> co
    ui --> win
    win --> internal/dll([internal/dll])
    win --> internal/utl
    win --> wstr
```

## License

Licensed under [MIT license](https://opensource.org/licenses/MIT), see [LICENSE.md](LICENSE.md) for details.
