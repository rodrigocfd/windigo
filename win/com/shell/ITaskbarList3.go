package shell

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/shell/shellco"
	"github.com/rodrigocfd/windigo/win/com/shell/shellvt"
	"github.com/rodrigocfd/windigo/win/errco"
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nn-shobjidl_core-itaskbarlist3
type ITaskbarList3 struct{ ITaskbarList2 }

// Constructs a COM object from a pointer to its COM virtual table.
//
// ⚠️ You must defer ITaskbarList3.Release().
//
// Example:
//
//  taskbl3 := shell.NewITaskbarList3(
//      win.CoCreateInstance(
//          shellco.CLSID_TaskbarList, nil,
//          co.CLSCTX_INPROC_SERVER,
//          shellco.IID_ITaskbarList3),
//  )
//  defer taskbl3.Release()
func NewITaskbarList3(ptr win.IUnknownPtr) ITaskbarList3 {
	return ITaskbarList3{
		ITaskbarList2: NewITaskbarList2(ptr),
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-registertab
func (me *ITaskbarList3) RegisterTab(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).RegisterTab, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setoverlayicon
func (me *ITaskbarList3) SetOverlayIcon(
	hWnd win.HWND, hIcon win.HICON, description string) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetOverlayIcon, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(hIcon),
		uintptr(unsafe.Pointer(win.Str.ToNativePtr(description))), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressstate
func (me *ITaskbarList3) SetProgressState(hWnd win.HWND, flags shellco.TBPF) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetProgressState, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(flags))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setprogressvalue
func (me *ITaskbarList3) SetProgressValue(
	hWnd win.HWND, completed, total uint64) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetProgressValue, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(completed), uintptr(total), 0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settabactive
func (me *ITaskbarList3) SetTabActive(hwndTab, hwndMDI win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetTabActive, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndMDI))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-settaborder
func (me *ITaskbarList3) SetTabOrder(hwndTab, hwndInsertBefore win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetTabOrder, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), uintptr(hwndInsertBefore))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailclip
func (me *ITaskbarList3) SetThumbnailClip(hWnd win.HWND, rcClip *win.RECT) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetThumbnailClip, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(unsafe.Pointer(rcClip)))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-setthumbnailtooltip
func (me *ITaskbarList3) SetThumbnailTooltip(hwnd win.HWND, tip string) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).SetThumbnailTooltip, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwnd), uintptr(unsafe.Pointer(win.Str.ToNativePtr(tip))))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbaraddbuttons
func (me *ITaskbarList3) ThumbBarAddButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarAddButtons, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarsetimagelist
func (me *ITaskbarList3) ThumbBarSetImageList(
	hWnd win.HWND, hImgl win.HIMAGELIST) {

	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarSetImageList, 3,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(hImgl))

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-thumbbarupdatebuttons
func (me *ITaskbarList3) ThumbBarUpdateButtons(
	hWnd win.HWND, buttons []THUMBBUTTON) {

	ret, _, _ := syscall.Syscall6(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).ThumbBarUpdateButtons, 4,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hWnd), uintptr(len(buttons)), uintptr(unsafe.Pointer(&buttons[0])),
		0, 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/shobjidl_core/nf-shobjidl_core-itaskbarlist3-unregistertab
func (me *ITaskbarList3) UnregisterTab(hwndTab win.HWND) {
	ret, _, _ := syscall.Syscall(
		(*shellvt.ITaskbarList3)(unsafe.Pointer(*me.Ptr())).UnregisterTab, 2,
		uintptr(unsafe.Pointer(me.Ptr())),
		uintptr(hwndTab), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
