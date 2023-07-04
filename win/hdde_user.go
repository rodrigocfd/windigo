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

// A handle to a [DDE instance]. Actually this handle does not exist, it's just
// a number identifying the instance.
//
// This handle is returned by win.DdeInitialize().
//
// [DDE instance]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeinitializew
type HDDE uint32

// [DdeInitialize] function.
//
// Note that this function is intended to be called only once. If you call it
// more than once, you'll overwrite the callback function.
//
// ⚠️ You must defer HDDE.DdeUninitialize().
//
// [DdeInitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeinitializew
func DdeInitialize(
	callback func(
		wType co.XTYP,
		wFmt uint32,
		hConv HCONV,
		hsz1, hsz2 HSZ,
		hData, dwData1, dwData2 uintptr) uintptr,
	afCmd co.AFCMD) (HDDE, error) {

	var idInst uint32

	pPack := &_DdeInitializePack{f: callback}
	_globalDdeInitizeMutex.Lock()
	_globalDdeInitializeFunc = pPack // store pointer
	_globalDdeInitizeMutex.Unlock()

	ret, _, _ := syscall.SyscallN(proc.DdeInitialize.Addr(),
		uintptr(unsafe.Pointer(&idInst)), _globalDdeInitializeCallback,
		uintptr(afCmd), 0)

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
	_globalDdeInitializeFunc     *_DdeInitializePack // DdeInitialize() should be called only once
	_globalDdeInitizeMutex       = sync.Mutex{}
	_globalDdeInitializeCallback = syscall.NewCallback(
		func(wType, wFmt uint32, hConv HCONV,
			hsz1, hsz2 HSZ, hData, dwData1, dwData2 uintptr) uintptr {

			return _globalDdeInitializeFunc.f(
				co.XTYP(wType), wFmt, hConv, hsz1, hsz2, hData, dwData1, dwData2)
		})
)

// [DdeGetLastError] function.
//
// [DdeGetLastError]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddegetlasterror
func (hDde HDDE) DdeGetLastError() errco.DMLERR {
	ret, _, _ := syscall.SyscallN(proc.DdeGetLastError.Addr(),
		uintptr(hDde))
	return errco.DMLERR(ret)
}

// [DdeNameService] function.
//
// [DdeNameService]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddenameservice
func (hDde HDDE) DdeNameService(serviceName StrOpt, opts co.DDENS) error {
	serviceNameHsz, err := hDde._strOptToHsz(serviceName)
	if err != nil {
		return err
	}
	defer hDde.DdeFreeStringHandle(serviceNameHsz)

	ret, _, _ := syscall.SyscallN(proc.DdeNameService.Addr(),
		uintptr(hDde), uintptr(serviceNameHsz), 0, uintptr(opts))
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// [DdePostAdvise] function.
//
// [DdePostAdvise]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddepostadvise
func (hDde HDDE) DdePostAdvise(topic, item StrOpt) error {
	topicHsz, err := hDde._strOptToHsz(topic)
	if err != nil {
		return err
	}
	defer hDde.DdeFreeStringHandle(topicHsz)

	itemHsz, err := hDde._strOptToHsz(item)
	if err != nil {
		return err
	}
	defer hDde.DdeFreeStringHandle(itemHsz)

	ret, _, _ := syscall.SyscallN(proc.DdePostAdvise.Addr(),
		uintptr(hDde), uintptr(topicHsz), uintptr(itemHsz))
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// [DdeUninitialize] function.
//
// [DdeUninitialize]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeuninitialize
func (hDde HDDE) DdeUninitialize() error {
	ret, _, _ := syscall.SyscallN(proc.DdeUninitialize.Addr(),
		uintptr(hDde))

	if ret == 0 {
		return errco.DMLERR_SYS_ERROR // no return error is actually specified
	} else {
		return nil
	}
}

//------------------------------------------------------------------------------

// DDE conversation [handle].
//
// [handle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeconnect
type HCONV HANDLE

// [DdeConnect] function.
//
// ⚠️ You must defer HDDE.DdeDisconnect().
//
// [DdeConnect]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeconnect
func (hDde HDDE) DdeConnect(
	serviceName,
	topic StrOpt,
	cc *CONVCONTEXT) (HCONV, error) {

	serviceNameHsz, err := hDde._strOptToHsz(serviceName)
	if err != nil {
		return HCONV(0), err
	}
	defer hDde.DdeFreeStringHandle(serviceNameHsz)

	topicHsz, err := hDde._strOptToHsz(topic)
	if err != nil {
		return HCONV(0), err
	}
	defer hDde.DdeFreeStringHandle(topicHsz)

	ret, _, _ := syscall.SyscallN(proc.DdeConnect.Addr(),
		uintptr(hDde), uintptr(serviceNameHsz), uintptr(topicHsz),
		uintptr(unsafe.Pointer(cc)))
	if ret == 0 {
		return HCONV(0), hDde.DdeGetLastError()
	}
	return HCONV(ret), nil
}

// [DdeDisconnect] function.
//
// [DdeDisconnect]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddedisconnect
func (hDde HDDE) DdeDisconnect(hConv HCONV) error {
	ret, _, _ := syscall.SyscallN(proc.DdeDisconnect.Addr(),
		uintptr(hConv))
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

//------------------------------------------------------------------------------

// DDE data [handle].
//
// [handle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeclienttransaction
type HDDEDATA HANDLE

// [DdeAddData] function.
//
// ⚠️ You must defer HDDE.DdeFreeDataHandle().
//
// [DdeAddData]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeadddata
func (hDde HDDE) DdeAddData(
	hData HDDEDATA,
	data []byte,
	offset int) (HDDEDATA, error) {

	ret, _, _ := syscall.SyscallN(proc.DdeAddData.Addr(),
		uintptr(hData), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)),
		uintptr(offset))
	if ret == 0 {
		return HDDEDATA(0), hDde.DdeGetLastError()
	}
	return HDDEDATA(ret), nil
}

// [DdeClientTransaction] function.
//
// For an async operation, pass -1 to timeout.
//
// ⚠️ You must defer HDDE.DdeFreeDataHandle().
//
// [DdeClientTransaction]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddeclienttransaction
func (hDde HDDE) DdeClientTransaction(
	data []byte,
	hConv HCONV,
	item StrOpt,
	fmt co.CF,
	xType co.XTYP,
	msTimeout int) (HDDEDATA, error) {

	var pData unsafe.Pointer
	if data != nil {
		pData = unsafe.Pointer(&data[0])
	}
	var szData int
	if data != nil {
		szData = len(data)
	}

	itemHsz, err := hDde._strOptToHsz(item)
	if err != nil {
		return HDDEDATA(0), err
	}
	defer hDde.DdeFreeStringHandle(itemHsz)

	timeout32 := uint32(_TIMEOUT_ASYNC)
	if msTimeout != -1 {
		timeout32 = uint32(msTimeout)
	}

	ret, _, _ := syscall.SyscallN(proc.DdeClientTransaction.Addr(),
		uintptr(pData), uintptr(szData), uintptr(hConv), uintptr(itemHsz),
		uintptr(fmt), uintptr(xType), uintptr(timeout32), 0)
	if ret == 0 {
		return HDDEDATA(0), hDde.DdeGetLastError()
	}
	return HDDEDATA(ret), nil
}

// [DdeFreeDataHandle] function.
//
// [DdeFreeDataHandle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddefreedatahandle
func (hDde HDDE) DdeFreeDataHandle(hData HDDEDATA) error {
	ret, _, _ := syscall.SyscallN(proc.DdeFreeDataHandle.Addr(),
		uintptr(hData))
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// [DdeGetData] function.
//
// The buffer size is automatically determined.
//
// [DdeGetData]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddegetdata
func (hDde HDDE) DdeGetData(hData HDDEDATA, offset int) ([]byte, error) {
	ret, _, _ := syscall.SyscallN(proc.DdeGetData.Addr(),
		uintptr(hData), 0, 0, uintptr(offset))
	if ret == 0 {
		return nil, hDde.DdeGetLastError()
	}

	numBytes := int(ret)
	retBuf := make([]byte, numBytes)

	ret, _, _ = syscall.SyscallN(proc.DdeGetData.Addr(),
		uintptr(hData), uintptr(unsafe.Pointer(&retBuf[0])),
		uintptr(numBytes), uintptr(offset))
	if ret == 0 {
		return nil, hDde.DdeGetLastError()
	}

	return retBuf, nil
}

//------------------------------------------------------------------------------

// DDE string [handle].
//
// [handle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddecreatestringhandlew
type HSZ HANDLE

// [DdeCmpStringHandles] function.
//
// [DdeCmpStringHandles]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddecmpstringhandles
func (hDde HDDE) DdeCmpStringHandles(hsz1, hsz2 HSZ) int {
	ret, _, _ := syscall.SyscallN(proc.DdeCmpStringHandles.Addr(),
		uintptr(hsz1), uintptr(hsz2))
	return int(ret)
}

// [DdeCreateStringHandle] function.
//
// ⚠️ You must defer HDDE.DdeFreeStringHandle().
//
// [DdeCreateStringHandle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddecreatestringhandlew
func (hDde HDDE) DdeCreateStringHandle(text string) (HSZ, error) {
	ret, _, _ := syscall.SyscallN(proc.DdeCreateStringHandle.Addr(),
		uintptr(hDde), uintptr(unsafe.Pointer(Str.ToNativePtr(text))),
		_CP_WINUNICODE)
	if ret == 0 {
		return HSZ(0), hDde.DdeGetLastError()
	}
	return HSZ(ret), nil
}

// [DdeFreeStringHandle] function.
//
// [DdeFreeStringHandle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddefreestringhandle
func (hDde HDDE) DdeFreeStringHandle(hsz HSZ) error {
	ret, _, _ := syscall.SyscallN(proc.DdeFreeStringHandle.Addr(),
		uintptr(hDde), uintptr(hsz))
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// [DdeKeepStringHandle] function.
//
// ⚠️ You must defer HDDE.DdeFreeStringHandle() on the hsz, because a clone of
// it has been made.
//
// [DdeKeepStringHandle]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddekeepstringhandle
func (hDde HDDE) DdeKeepStringHandle(hsz HSZ) error {
	ret, _, _ := syscall.SyscallN(proc.DdeKeepStringHandle.Addr(),
		uintptr(hDde), uintptr(hsz))
	if ret == 0 {
		return hDde.DdeGetLastError()
	}
	return nil
}

// [DdeQueryString] function.
//
// [DdeQueryString]: https://learn.microsoft.com/en-us/windows/win32/api/ddeml/nf-ddeml-ddequerystringw
func (hDde HDDE) DdeQueryString(hsz HSZ) (string, error) {
	strLen, _, _ := syscall.SyscallN(proc.DdeQueryString.Addr(),
		uintptr(hDde), uintptr(hsz), 0, 0, _CP_WINUNICODE)
	if strLen == 0 {
		return "", hDde.DdeGetLastError()
	}

	buf := make([]uint16, strLen+1)
	ret, _, _ := syscall.SyscallN(proc.DdeQueryString.Addr(),
		uintptr(hDde), uintptr(hsz), uintptr(unsafe.Pointer(&buf[0])),
		strLen+1, _CP_WINUNICODE)
	if ret == 0 {
		return "", hDde.DdeGetLastError()
	}

	return Str.FromNativeSlice(buf), nil
}

//------------------------------------------------------------------------------

func (hDde HDDE) _strOptToHsz(s StrOpt) (HSZ, error) {
	var hszVal HSZ
	if strVal, ok := s.Str(); ok {
		hsz, err := hDde.DdeCreateStringHandle(strVal)
		if err != nil {
			return HSZ(0), err
		}
		hszVal = hsz
	}
	return hszVal, nil
}
