//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/internal/dll"
	"github.com/rodrigocfd/windigo/internal/utl"
	"github.com/rodrigocfd/windigo/wstr"
)

// [CLSIDFromProgID] function.
//
// Used to retrieve class IDs to create COM Automation objects. If the progId is
// invalid, returns [co.HRESULT_CO_E_CLASSSTRING].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	clsId, _ := win.CLSIDFromProgID("Excel.Application")
//
//	var excel *win.IDispatch
//	_ = win.CoCreateInstance(
//		rel,
//		clsId,
//		nil,
//		co.CLSCTX_LOCAL_SERVER,
//		&excel,
//	)
//
// [CLSIDFromProgID]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-clsidfromprogid
func CLSIDFromProgID(progId string) (co.CLSID, error) {
	var wProgId wstr.BufEncoder
	var guid GUID

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_CLSIDFromProgID, "CLSIDFromProgID"),
		uintptr(wProgId.AllowEmpty(progId)),
		uintptr(unsafe.Pointer(&guid)))

	if hr := co.HRESULT(ret); hr == co.HRESULT_S_OK {
		return co.CLSID(guid.String()), nil
	} else {
		return "", hr
	}
}

var _ole_CLSIDFromProgID *syscall.Proc

// [CoCreateInstance] function.
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
//	rel := win.NewOleReleaser()
//	defer rel.Release()
//
//	var taskbl *win.ITaskbarList
//	_ = win.CoCreateInstance(
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
	iid := com_validateAndRelease(ppOut, releaser)
	var ppvtQueried **_IUnknownVt
	guidClsid := GuidFrom(rclsid)
	guidIid := GuidFrom(iid)

	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_CoCreateInstance, "CoCreateInstance"),
		uintptr(unsafe.Pointer(&guidClsid)),
		uintptr(com_ppvtOrNil(unkOuter)),
		uintptr(dwClsContext),
		uintptr(unsafe.Pointer(&guidIid)),
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retHres(ret, ppOut, ppvtQueried, releaser)
}

var _ole_CoCreateInstance *syscall.Proc

// [CoInitializeEx] function.
//
// ⚠️ You must defer [CoUninitialize].
//
// Example:
//
//	_, _ = win.CoInitializeEx(
//		co.COINIT_APARTMENTTHREADED | co.COINIT_DISABLE_OLE1DDE)
//	defer win.CoUninitialize()
//
// [CoInitializeEx]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-coinitializeex
func CoInitializeEx(coInit co.COINIT) (alreadyInitialized bool, hr error) {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_CoInitializeEx, "CoInitializeEx"),
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

var _ole_CoInitializeEx *syscall.Proc

// [CoUninitialize] function.
//
// Paired [CoInitializeEx].
//
// [CoUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/combaseapi/nf-combaseapi-couninitialize
func CoUninitialize() {
	syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_CoUninitialize, "CoUninitialize"))
}

var _ole_CoUninitialize *syscall.Proc

// [CreateBindCtx] function.
//
// Example:
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
		dll.Load(dll.OLE32, &_ole_CreateBindCtx, "CreateBindCtx"),
		0,
		uintptr(unsafe.Pointer(&ppvtQueried)))
	return com_buildObj_retObjHres[*IBindCtx](ret, ppvtQueried, releaser)
}

var _ole_CreateBindCtx *syscall.Proc

// [OleInitialize] function.
//
// ⚠️ You must defer [OleUninitialize].
//
// Example:
//
//	_ = win.OleInitialize()
//	defer win.OleUninitialize()
//
// [OleInitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleinitialize
func OleInitialize() error {
	ret, _, _ := syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_OleInitialize, "OleInitialize"),
		0)
	return utl.HresultToError(ret)
}

var _ole_OleInitialize *syscall.Proc

// [OleUninitialize] function.
//
// Paired with [OleInitialize].
//
// [OleUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-oleuninitialize
func OleUninitialize() {
	syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_OleUninitialize, "OleUninitialize"))
}

var _ole_OleUninitialize *syscall.Proc

// [ReleaseStgMedium] function.
//
// Paired with [IDataObject.GetData].
//
// [ReleaseStgMedium]: https://learn.microsoft.com/en-us/windows/win32/api/ole/nf-ole-releasestgmedium
func ReleaseStgMedium(stg *STGMEDIUM) {
	syscall.SyscallN(
		dll.Load(dll.OLE32, &_ole_ReleaseStgMedium, "ReleaseStgMedium"),
		uintptr(unsafe.Pointer(stg)))
}

var _ole_ReleaseStgMedium *syscall.Proc
