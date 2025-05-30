//go:build windows

package ole

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/win"
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
//	ole.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer ole.CoUninitialize()
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	clsId, _ := ole.CLSIDFromProgID("Excel.Application")
//
//	var dispExcel *oleaut.IDispatch
//	ole.CoCreateInstance(
//		rel, clsId, nil, co.CLSCTX_LOCAL_SERVER, &dispExcel)
//
// [CLSIDFromProgID]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	progId16 := wstr.NewBufWith[wstr.Stack20](progId, wstr.ALLOW_EMPTY)
	var guid win.GUID

	ret, _, _ := syscall.SyscallN(_CLSIDFromProgID.Addr(),
		uintptr(progId16.UnsafePtr()),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

var _CLSIDFromProgID = dll.Ole32.NewProc("CLSIDFromProgID")

// [CoCreateInstance] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	var taskbl *shell.ITaskbarList
//	ole.CoCreateInstance(
//		rel,
//		co.CLSID_TaskbarList,
//		nil,
//		co.CLSCTX_INPROC_SERVER,
//		&taskbl,
//	)
//
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance(
	releaser *Releaser,
	rclsid co.CLSID,
	unkOuter *IUnknown,
	dwClsContext co.CLSCTX,
	ppOut interface{},
) error {
	pOut := utl.ComValidateObj(ppOut).(ComObj)
	releaser.ReleaseNow(pOut)

	var ppvtQueried **IUnknownVt
	guidClsid := win.GuidFrom(rclsid)
	guidIid := win.GuidFrom(pOut.IID())

	var pUnkOuter **IUnknownVt
	if unkOuter != nil {
		pUnkOuter = unkOuter.Ppvt()
	}

	ret, _, _ := syscall.SyscallN(_CoCreateInstance.Addr(),
		uintptr(unsafe.Pointer(&guidClsid)),
		uintptr(unsafe.Pointer(pUnkOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pOut = utl.ComCreateObj(ppOut, unsafe.Pointer(ppvtQueried)).(ComObj)
		releaser.Add(pOut)
		return nil
	} else {
		return hr
	}
}

var _CoCreateInstance = dll.Ole32.NewProc("CoCreateInstance")

// [CoInitializeEx] function.
//
// ⚠️ You must defer [CoUninitialize].
//
// # Example
//
//	ole.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer ole.CoUninitialize()
//
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit co.COINIT) (alreadyInitialized bool, hr error) {
	ret, _, _ := syscall.SyscallN(_CoInitializeEx.Addr(),
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

var _CoInitializeEx = dll.Ole32.NewProc("CoInitializeEx")

// [CoUninitialize] function.
//
// Paired [CoInitializeEx].
//
// [CoUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.SyscallN(_CoUninitialize.Addr())
}

var _CoUninitialize = dll.Ole32.NewProc("CoUninitialize")

// [CreateBindCtx] function.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	bindCtx, _ := ole.CreateBindCtx(rel)
//
// [CreateBindCtx]: https://learn.microsoft.com/en-us/windows/win32/api/objbase/nf-objbase-createbindctx
func CreateBindCtx(releaser *Releaser) (*IBindCtx, error) {
	var ppvtQueried **IUnknownVt
	ret, _, _ := syscall.SyscallN(_CreateBindCtx.Addr(),
		0,
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		var pObj *IBindCtx
		utl.ComCreateObj(&pObj, unsafe.Pointer(ppvtQueried))
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _CreateBindCtx = dll.Ole32.NewProc("CreateBindCtx")

// [OleInitialize] function.
//
// ⚠️ You must defer [OleUninitialize].
//
// # Example
//
//	ole.OleInitialize()
//	defer ole.OleUninitialize()
//
// [OleInitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleinitialize
func OleInitialize() error {
	ret, _, _ := syscall.SyscallN(_OleInitialize.Addr(),
		0)
	return utl.ErrorAsHResult(ret)
}

var _OleInitialize = dll.Ole32.NewProc("OleInitialize")

// [OleUninitialize] function.
//
// Paired with [OleInitialize].
//
// [OleUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleuninitialize
func OleUninitialize() {
	syscall.SyscallN(_OleUninitialize.Addr())
}

var _OleUninitialize = dll.Ole32.NewProc("OleUninitialize")

// [RegisterDragDrop] function.
//
// Paired with [RevokeDragDrop].
//
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-registerdragdrop
func RegisterDragDrop(hWnd win.HWND, dropTarget *IDropTarget) error {
	exStyle, _ := hWnd.ExStyle()
	if (exStyle & co.WS_EX_ACCEPTFILES) != 0 {
		return errors.New("do not use WS_EX_ACCEPTFILES with RegisterDragDrop")
	}

	ret, _, _ := syscall.SyscallN(_RegisterDragDrop.Addr(),
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

var _RegisterDragDrop = dll.Ole32.NewProc("RegisterDragDrop")

// [ReleaseStgMedium] function.
//
// Paired with [IDataObject.GetData].
//
// [ReleaseStgMedium]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-releasestgmedium
func ReleaseStgMedium(stg *STGMEDIUM) {
	syscall.SyscallN(_ReleaseStgMedium.Addr(),
		uintptr(unsafe.Pointer(stg)))
}

var _ReleaseStgMedium = dll.Ole32.NewProc("ReleaseStgMedium")

// [RevokeDragDrop] function.
//
// Paired with [RegisterDragDrop].
//
// [RevokeDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-revokedragdrop
func RevokeDragDrop(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(_RevokeDragDrop.Addr(),
		uintptr(hWnd))
	return utl.ErrorAsHResult(ret)
}

var _RevokeDragDrop = dll.Ole32.NewProc("RevokeDragDrop")

// [SHCreateMemStream] function.
//
// Creates an [IStream] projection over a slice, which must remain valid in
// memory throughout IStream's lifetime.
//
// # Example
//
//	rel := ole.NewReleaser()
//	defer rel.Release()
//
//	data := []byte{0x10, 0x11, 0x12}
//	defer runtime.KeepAlive(data)
//
//	stream, _ := SHCreateMemStream(rel, data)
//
// [SHCreateMemStream]: https://learn.microsoft.com/en-us/windows/win32/api/shlwapi/nf-shlwapi-shcreatememstream
func SHCreateMemStream(releaser *Releaser, src []byte) (*IStream, error) {
	ret, _, _ := syscall.SyscallN(_SHCreateMemStream.Addr(),
		uintptr(unsafe.Pointer(&src[0])),
		uintptr(uint32(len(src))))
	if ret == 0 {
		return nil, co.HRESULT_E_OUTOFMEMORY
	}

	ppvt := (**IUnknownVt)(unsafe.Pointer(ret))
	var pObj *IStream
	utl.ComCreateObj(&pObj, unsafe.Pointer(ppvt))
	releaser.Add(pObj)
	return pObj, nil
}

var _SHCreateMemStream = dll.Shlwapi.NewProc("SHCreateMemStream")
