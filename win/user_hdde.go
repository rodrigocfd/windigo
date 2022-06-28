//go:build windows

package win

import (
	"sync"
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/errco"
)

// A handle to a DDE instance. Actually this handle does not exist, it's just a
// number identifying the instance.
//
// This handle is returned by win.DdeInitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeinitializew
type HDDE uint32

// Note that this function is intended to be called only once. If you call it
// more than once, you'll overwrite the callback function.
//
// ⚠️ You must defer HDDE.DdeUninitialize().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeinitializew
func DdeInitialize(
	callback func(
		wType co.XTYP, wFmt uint32, hConv HCONV,
		hsz1, hsz2 HSZ, hData, dwData1, dwData2 uintptr) uintptr,
	afCmd co.AFCMD) (HDDE, error) {

	var idInst uint32

	pPack := &_DdeInitializePack{f: callback}
	_globalDdeInitizeMutex.Lock()
	_globalDdeInitializeFunc = pPack // store pointer
	_globalDdeInitizeMutex.Unlock()

	ret, _, _ := syscall.Syscall6(proc.DdeInitialize.Addr(), 4,
		uintptr(unsafe.Pointer(&idInst)), _globalDdeInitializeCallback,
		uintptr(afCmd), 0,
		0, 0)

	if dmlErr := errco.DMLERR(ret); dmlErr != errco.DMLERR_NO_ERROR {
		return 0, dmlErr
	} else {
		return HDDE(idInst), nil
	}
}

type _DdeInitializePack struct {
	f func(wType co.XTYP, wFmt uint32, hConv HCONV,
		hsz1, hsz2 HSZ, hData, dwData1, dwData2 uintptr) uintptr
}

var (
	_globalDdeInitializeCallback uintptr = syscall.NewCallback(_DdeInitializeCallback)
	_globalDdeInitializeFunc     *_DdeInitializePack
	_globalDdeInitizeMutex       = sync.Mutex{}
)

func _DdeInitializeCallback(
	wType, wFmt uint32, hConv HCONV,
	hsz1, hsz2 HSZ, hData, dwData1, dwData2 uintptr) uintptr {

	return _globalDdeInitializeFunc.f(
		co.XTYP(wType), wFmt, hConv, hsz1, hsz2, hData, dwData1, dwData2)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddegetlasterror
func (hDde HDDE) DdeGetLastError() errco.DMLERR {
	ret, _, _ := syscall.Syscall(proc.DdeGetLastError.Addr(), 1,
		uintptr(hDde), 0, 0)
	return errco.DMLERR(ret)
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeuninitialize
func (hDde HDDE) DdeUninitialize() error {
	ret, _, _ := syscall.Syscall(proc.DdeUninitialize.Addr(), 1,
		uintptr(hDde), 0, 0)

	if ret == 0 {
		return errco.DMLERR_SYS_ERROR // no return error is actually specified
	} else {
		return nil
	}
}

//------------------------------------------------------------------------------

// DDE conversation handle.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeconnect
type HCONV HANDLE

// ⚠️ You must defer HDDE.DdeDisconnect().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeconnect
func (hDde HDDE) DdeConnect(
	serviceName, topic StrOpt, cc *CONVCONTEXT) (HCONV, error) {

	var serviceNameHsz HSZ
	if s, ok := serviceName.Str(); ok {
		hsz, err := hDde.DdeCreateStringHandle(s)
		if err != nil {
			return HCONV(0), err
		}
		serviceNameHsz = hsz
		defer hDde.DdeFreeStringHandle(serviceNameHsz)
	}

	var topicHsz HSZ
	if s, ok := topic.Str(); ok {
		hsz, err := hDde.DdeCreateStringHandle(s)
		if err != nil {
			return HCONV(0), err
		}
		topicHsz = hsz
		defer hDde.DdeFreeStringHandle(topicHsz)
	}

	ret, _, _ := syscall.Syscall6(proc.DdeConnect.Addr(), 4,
		uintptr(hDde), uintptr(serviceNameHsz), uintptr(topicHsz),
		uintptr(unsafe.Pointer(cc)),
		0, 0)
	if ret == 0 {
		return HCONV(0), hDde.DdeGetLastError()
	}
	return HCONV(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddedisconnect
func (hDde HDDE) DdeDisconnect(hConv HCONV) error {
	ret, _, _ := syscall.Syscall(proc.DdeDisconnect.Addr(), 1,
		uintptr(hConv), 0, 0)
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

//------------------------------------------------------------------------------

// DDE data handle.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeclienttransaction
type HDDEDATA HANDLE

// For an async operation, pass -1 to timeout.
//
// ⚠️ You must defer HDDE.DdeFreeDataHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeclienttransaction
func (hDde HDDE) DdeClientTransaction(
	data []byte, hConv HCONV, item StrOpt,
	fmt co.CF, xType co.XTYP, timeout int) (HDDEDATA, error) {

	timeout32 := uint32(_TIMEOUT_ASYNC)
	if timeout != -1 {
		timeout32 = uint32(timeout)
	}

	var itemHsz HSZ
	if s, ok := item.Str(); ok {
		hsz, err := hDde.DdeCreateStringHandle(s)
		if err != nil {
			return HDDEDATA(0), err
		}
		itemHsz = hsz
		defer hDde.DdeFreeStringHandle(itemHsz)
	}

	ret, _, _ := syscall.Syscall9(proc.DdeClientTransaction.Addr(), 8,
		uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)), uintptr(hConv),
		uintptr(itemHsz), uintptr(fmt), uintptr(xType), uintptr(timeout32), 0, 0)
	if ret == 0 {
		return HDDEDATA(0), hDde.DdeGetLastError()
	}
	return HDDEDATA(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddefreedatahandle
func (hDde HDDE) DdeFreeDataHandle(hData HDDEDATA) error {
	ret, _, _ := syscall.Syscall(proc.DdeFreeDataHandle.Addr(), 1,
		uintptr(hData), 0, 0)
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

//------------------------------------------------------------------------------

// DDE string handle.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddecreatestringhandlew
type HSZ HANDLE

// ⚠️ You must defer HDDE.DdeFreeStringHandle().
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddecreatestringhandlew
func (hDde HDDE) DdeCreateStringHandle(text string) (HSZ, error) {
	ret, _, _ := syscall.Syscall(proc.DdeCreateStringHandle.Addr(), 3,
		uintptr(hDde), uintptr(unsafe.Pointer(Str.ToNativePtr(text))),
		_CP_WINUNICODE)
	if ret == 0 {
		return HSZ(0), hDde.DdeGetLastError()
	}
	return HSZ(ret), nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddefreestringhandle
func (hDde HDDE) DdeFreeStringHandle(hsz HSZ) error {
	ret, _, _ := syscall.Syscall(proc.DdeFreeStringHandle.Addr(), 2,
		uintptr(hDde), uintptr(hsz), 0)
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// ⚠️ You must defer HDDE.DdeFreeStringHandle() on the hsz, because a clone of
// it has been made.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddekeepstringhandle
func (hDde HDDE) DdeKeepStringHandle(hsz HSZ) error {
	ret, _, _ := syscall.Syscall(proc.DdeKeepStringHandle.Addr(), 2,
		uintptr(hDde), uintptr(hsz), 0)
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddequerystringw
func (hDde HDDE) DdeQueryString(hsz HSZ) (string, error) {
	strLen, _, _ := syscall.Syscall6(proc.DdeQueryString.Addr(), 5,
		uintptr(hDde), uintptr(hsz), 0, 0, _CP_WINUNICODE, 0)
	if strLen == 0 {
		return "", hDde.DdeGetLastError()
	}

	buf := make([]uint16, strLen+1)
	ret, _, _ := syscall.Syscall6(proc.DdeQueryString.Addr(), 5,
		uintptr(hDde), uintptr(hsz), uintptr(unsafe.Pointer(&buf[0])),
		strLen+1, _CP_WINUNICODE, 0)
	if ret == 0 {
		return "", hDde.DdeGetLastError()
	}

	return Str.FromNativeSlice(buf), nil
}
