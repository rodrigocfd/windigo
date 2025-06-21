//go:build windows

package win

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [CLSIDFromProgID] function.
//
// Used to retrieve class IDs to create COM Automation objects. If the progId is
// invalid, returns [co.HRESULT_CO_E_CLASSSTRING].
//
// # Example
//
//	win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//	var excel *win.IDispatch
//	win.CoCreateInstance(rel, clsId, nil, co.CLSCTX_LOCAL_SERVER, &excel)
//
// [CLSIDFromProgID]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	wbuf := wstr.NewBufConverter()
	defer wbuf.Free()
	pProgId := wbuf.PtrAllowEmpty(progId)

	var guid GUID

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_CLSIDFromProgID, "CLSIDFromProgID"),
		uintptr(pProgId),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

var _CLSIDFromProgID *syscall.Proc

// [CoCreateInstance] function.
//
// # Example
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var taskbl *win.ITaskbarList
//	win.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(
	releaser *OleReleaser,
	rclsid co.CLSID,
	unkOuter *IUnknown,
	dwClsContext co.CLSCTX,
	ppOut interface{},
) error {
	pOut := utl.OleValidateObj(ppOut).(OleObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **_IUnknownVt
	guidClsid := GuidFrom(rclsid)
	guidIid := GuidFrom(pOut.IID())

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_CoCreateInstance, "CoCreateInstance"),
		uintptr(unsafe.Pointer(&guidClsid)),
		uintptr(ppvtOrNil(unkOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.OleCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(OleObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

var _CoCreateInstance *syscall.Proc

// [CoInitializeEx] function.
//
// ⚠️ You must defer [CoUninitialize].
//
// # Example
//
//	win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit co.COINIT) (alreadyInitialized bool, hr error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_CoInitializeEx, "CoInitializeEx"),
		0,
		uintptr(coInit))

	if hr = co.HRESULT(ret); hr == co.HRESULT_S_OK {
		alreadyInitialized, hr = false, nil
	} else if hr == co.HRESULT_S_FALSE {
		alreadyInitialized, hr = true, nil
	} else {
		alreadyInitialized = false
	}
	return
}

var _CoInitializeEx *syscall.Proc

// [CoUninitialize] function.
//
// Paired [CoInitializeEx].
//
// [CoUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.SyscallN(
		dll.Load(dll.OLE32, &_CoUninitialize, "CoUninitialize"))
}

var _CoUninitialize *syscall.Proc

// [CreateBindCtx] function.
//
// # Example
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	bindCtx, _ := win.CreateBindCtx(rel)
//
// [CreateBindCtx]: https://learn.microsoft.com/en-us/windows/win32/api/objbase/nf-objbase-createbindctx
func CreateBindCtx(releaser *OleReleaser) (*IBindCtx, error) {
	var ppvtQueried **_IUnknownVt
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_CreateBindCtx, "CreateBindCtx"),
		0,
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := &IBindCtx{IUnknown{ppvtQueried}}
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _CreateBindCtx *syscall.Proc

// [OleInitialize] function.
//
// ⚠️ You must defer [OleUninitialize].
//
// # Example
//
//	win.OleInitialize()
//	defer win.OleUninitialize()
//
// [OleInitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleinitialize
func OleInitialize() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_OleInitialize, "OleInitialize"),
		0)
	return utl.ErrorAsHResult(ret)
}

var _OleInitialize *syscall.Proc

// [OleUninitialize] function.
//
// Paired with [OleInitialize].
//
// [OleUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleuninitialize
func OleUninitialize() {
	syscall.SyscallN(
		dll.Load(dll.OLE32, &_OleUninitialize, "OleUninitialize"))
}

var _OleUninitialize *syscall.Proc

// [RegisterDragDrop] function.
//
// Paired with [RevokeDragDrop].
//
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-registerdragdrop
func RegisterDragDrop(hWnd HWND, dropTarget *IDropTarget) error {
	exStyle, _ := hWnd.ExStyle()
	if (exStyle & co.WS_EX_ACCEPTFILES) != 0 {
		return errors.New("do not use WS_EX_ACCEPTFILES with RegisterDragDrop")
	}

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_RegisterDragDrop, "RegisterDragDrop"),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(dropTarget.Ppvt())))
	if hr := co.HRESULT(ret); hr != co.HRESULT_S_OK {
		if hr == co.HRESULT_E_OUTOFMEMORY {
			return errors.New("RegisterDragDrop failed, did you call OleInitialize?")
		}
		return hr
	}
	return nil
}

var _RegisterDragDrop *syscall.Proc

// [ReleaseStgMedium] function.
//
// Paired with [IDataObject.GetData].
//
// [ReleaseStgMedium]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-releasestgmedium
func ReleaseStgMedium(stg *STGMEDIUM) {
	syscall.SyscallN(
		dll.Load(dll.OLE32, &_ReleaseStgMedium, "ReleaseStgMedium"),
		uintptr(unsafe.Pointer(stg)))
}

var _ReleaseStgMedium *syscall.Proc

// [RevokeDragDrop] function.
//
// Paired with [RegisterDragDrop].
//
// [RevokeDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-revokedragdrop
func RevokeDragDrop(hWnd HWND) error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_RevokeDragDrop, "RevokeDragDrop"),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

var _RevokeDragDrop *syscall.Proc

// [SHCreateMemStream] function.
//
// Creates an [IStream] projection over a slice, which must remain valid in
// memory throughout IStream's lifetime.
//
// # Example
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	data := []byte{0x10, 0x11, 0x12}
//	defer runtime.KeepAlive(data)
//
//	stream, _ := win.SHCreateMemStream(rel, data)
//
// [SHCreateMemStream]: https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func SHCreateMemStream(releaser *OleReleaser, src []byte) (*IStream, error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.SHLWAPI, &_SHCreateMemStream, "SHCreateMemStream"),
		uintptr(unsafe.Pointer(&src[0])),
		uintptr(uint32(len(src))))
	if ret == 0 {
		return nil, co.HRESULT_E_OUTOFMEMORY
	}

	ppvt := (**_IUnknownVt)(unsafe.Pointer(ret))
	pObj := &IStream{ISequentialStream{IUnknown{ppvt}}}
	releaser.Add(pObj)
	return pObj, nil
}

var _SHCreateMemStream *syscall.Proc
