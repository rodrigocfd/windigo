//go:build windows

package ole

import (
	"errors"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/util"
	"github.com/rodrigocfd/windigo/internal/vt"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/wstr"
)

// [CLSIDFromProgID] function.
//
// Used to retrieve class IDs to create COM Automation objects. If the progId is
// invalid, returns errco.CO_E_CLASSSTRING.
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
//	excel, _ := ole.CoCreateInstance[ole.IDispatch](
//		rel, clsId, co.CLSCTX_LOCAL_SERVER)
//
// [CLSIDFromProgID]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	progId16 := wstr.NewBufWith[wstr.Stack20](progId, wstr.ALLOW_EMPTY)
	var guid win.GUID

	ret, _, _ := syscall.SyscallN(_CLSIDFromProgID.Addr(),
		uintptr(progId16.UnsafePtr()), uintptr(unsafe.Pointer(&guid)))

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
//	taskbl, _ := ole.CoCreateInstance[shell.ITaskbarList](
//		rel,
//		co.CLSID_TaskbarList,
//		co.CLSCTX_INPROC_SERVER,
//	)
//
// [CoCreateInstance]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-cocreateinstance
func CoCreateInstance[T any, P ComCtor[T]](
	releaser *Releaser,
	rclsid co.CLSID,
	dwClsContext co.CLSCTX,
) (*T, error) {
	pObj := P(new(T)) // https://stackoverflow.com/a/69575720/6923555
	var ppvtQueried **vt.IUnknown
	guidClsid := win.GuidFrom(rclsid)
	guidIid := win.GuidFrom(pObj.IID())

	ret, _, _ := syscall.SyscallN(_CoCreateInstance.Addr(),
		uintptr(unsafe.Pointer(&guidClsid)), 0, // don't query pUnkOuter
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj.Set(ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _CoCreateInstance = dll.Ole32.NewProc("CoCreateInstance")

// [CoInitializeEx] function.
//
// ⚠️ You must defer CoUninitialize().
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
		0, uintptr(coInit))

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
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
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
	var ppvtQueried **vt.IUnknown
	ret, _, _ := syscall.SyscallN(_CreateBindCtx.Addr(),
		0, uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := vt.NewObj[IBindCtx](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _CreateBindCtx = dll.Ole32.NewProc("CreateBindCtx")

// [OleLoadPicture] function.
//
// Pass size = 0 to read all the bytes from the stream.
//
// The bytes are copied, so IStream can be released after this function returns.
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
//	pic, _ := OleLoadPicture(rel, &stream, 0, true)
//
// [OleLoadPicture]: https://learn.microsoft.com/en-us/windows/win32/api/olectl/nf-olectl-oleloadpicture
func OleLoadPicture(
	releaser *Releaser,
	stream *IStream,
	size uint,
	keepOriginalFormat bool,
) (*IPicture, error) {
	var ppvtQueried **vt.IUnknown
	guid := win.GuidFrom(co.IID_IPicture)

	ret, _, _ := syscall.SyscallN(_OleLoadPicture.Addr(),
		uintptr(unsafe.Pointer(stream.Ppvt())),
		uintptr(size),
		util.BoolToUintptr(!keepOriginalFormat), // note: reversed
		uintptr(unsafe.Pointer(&guid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := vt.NewObj[IPicture](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _OleLoadPicture = dll.Oleaut32.NewProc("OleLoadPicture")

// [OleInitialize] function.
//
// ⚠️ You must defer OleUninitialize().
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
	return util.ErrorAsHResult(ret)
}

var _OleInitialize = dll.Ole32.NewProc("OleInitialize")

// [OleLoadPicturePath] function.
//
// The picture must be in the following formats:
//   - BMP (bitmap)
//   - JPEG
//   - WMF (metafile)
//   - ICO (icon)
//   - GIF
//
// [OleLoadPicturePath]: https://learn.microsoft.com/en-us/windows/win32/api/olectl/nf-olectl-oleloadpicturepath
func OleLoadPicturePath(
	releaser *Releaser,
	path string,
	transparentColor win.COLORREF,
) (*IPicture, error) {
	path16 := wstr.NewBufWith[wstr.Stack20](path, wstr.EMPTY_IS_NIL)
	var ppvtQueried **vt.IUnknown
	guid := win.GuidFrom(co.IID_IPicture)

	ret, _, _ := syscall.SyscallN(_OleLoadPicturePath.Addr(),
		uintptr(path16.UnsafePtr()), 0, 0, uintptr(transparentColor),
		uintptr(unsafe.Pointer(&guid)), uintptr(unsafe.Pointer(&ppvtQueried)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		pObj := vt.NewObj[IPicture](ppvtQueried)
		releaser.Add(pObj)
		return pObj, nil
	} else {
		return nil, hr
	}
}

var _OleLoadPicturePath = dll.Oleaut32.NewProc("OleLoadPicturePath")

// [OleUninitialize] function.
//
// [OleUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleuninitialize
func OleUninitialize() {
	syscall.SyscallN(_OleUninitialize.Addr())
}

var _OleUninitialize = dll.Ole32.NewProc("OleUninitialize")

// [RegisterDragDrop] function.
//
// [RegisterDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole2/nf-ole2-registerdragdrop
func RegisterDragDrop(hWnd win.HWND, dropTarget *IDropTarget) error {
	exStyle, _ := hWnd.ExStyle()
	if (exStyle & co.WS_EX_ACCEPTFILES) != 0 {
		return errors.New("do not use WS_EX_ACCEPTFILES with RegisterDragDrop")
	}

	ret, _, _ := syscall.SyscallN(_RegisterDragDrop.Addr(),
		uintptr(hWnd), uintptr(unsafe.Pointer(dropTarget.Ppvt())))
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
// [ReleaseStgMedium]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-releasestgmedium
func ReleaseStgMedium(stg *STGMEDIUM) {
	syscall.SyscallN(_ReleaseStgMedium.Addr(),
		uintptr(unsafe.Pointer(stg)))
}

var _ReleaseStgMedium = dll.Ole32.NewProc("ReleaseStgMedium")

// [RevokeDragDrop] function.
//
// [RevokeDragDrop]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-revokedragdrop
func RevokeDragDrop(hWnd win.HWND) error {
	ret, _, _ := syscall.SyscallN(_RevokeDragDrop.Addr(),
		uintptr(hWnd))
	return util.ErrorAsHResult(ret)
}

var _RevokeDragDrop = dll.Ole32.NewProc("RevokeDragDrop")

// [SHCreateMemStream] function.
//
// Creates an IStream projection over a slice, which must remain valid in
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
		uintptr(unsafe.Pointer(&src[0])), uintptr(len(src)))
	if ret == 0 {
		return nil, co.HRESULT_E_OUTOFMEMORY
	}

	pObj := vt.NewObj[IStream]((**vt.IUnknown)(unsafe.Pointer(ret)))
	releaser.Add(pObj)
	return pObj, nil
}

var _SHCreateMemStream = dll.Shlwapi.NewProc("SHCreateMemStream")
