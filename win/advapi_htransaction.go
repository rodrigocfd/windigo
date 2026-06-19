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

// Handle to a [transaction].
//
// [transaction]: https://learn.microsoft.com/en-us/windows/win32/ktm/ktm-security-and-access-rights
type HTRANSACTION HANDLE

// [CreateTransaction] function.
//
// ⚠️ You must defer [HTRANSACTION.CloseHandle].
//
// [CreateTransaction]: https://learn.microsoft.com/en-us/windows/win32/api/ktmw32/nf-ktmw32-createtransaction
func (hTrans HTRANSACTION) CreateTransaction(
	pSecurityAttributes *SECURITY_ATTRIBUTES,
	options co.TRANSACTION_OPT,
	timeout Timeout,
	description string,
) (HTRANSACTION, error) {
	var wDescription wstr.BufEncoder
	ret, _, err := syscall.SyscallN(
		dll.Ktmw.Load(&_ktmw_CreateTransaction, "CreateTransaction"),
		uintptr(unsafe.Pointer(pSecurityAttributes)),
		0,
		uintptr(options),
		0, 0,
		uintptr(timeout.raw()),
		uintptr(wDescription.EmptyIsNil(description)))
	if ret == 0 && int(ret) == utl.INVALID_HANDLE_VALUE {
		return HTRANSACTION(0), co.ERROR(err)
	}
	return HTRANSACTION(ret), nil
}

var _ktmw_CreateTransaction *syscall.Proc

// [OpenTransaction] function.
//
// ⚠️ You must defer [HTRANSACTION.CloseHandle].
//
// [OpenTransaction]: https://learn.microsoft.com/en-us/windows/win32/api/ktmw32/nf-ktmw32-opentransaction
func OpenTransaction(desiredAccess co.TRANSACTION, pGuid *co.GUID) (HTRANSACTION, error) {
	ret, _, err := syscall.SyscallN(
		dll.Ktmw.Load(&_ktmw_OpenTransaction, "OpenTransaction"),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(pGuid)))
	if ret == 0 && int(ret) == utl.INVALID_HANDLE_VALUE {
		return HTRANSACTION(0), co.ERROR(err)
	}
	return HTRANSACTION(ret), nil
}

var _ktmw_OpenTransaction *syscall.Proc

// [CloseHandle] function.
//
// [CloseHandle]: https://learn.microsoft.com/en-us/windows/win32/api/handleapi/nf-handleapi-closehandle
func (hTrans HTRANSACTION) CloseHandle() error {
	return HANDLE(hTrans).CloseHandle()
}

// [CommitTransaction] function.
//
// [CommitTransaction]: https://learn.microsoft.com/en-us/windows/win32/api/ktmw32/nf-ktmw32-committransaction
func (hTrans HTRANSACTION) CommitTransaction() error {
	ret, _, err := syscall.SyscallN(
		dll.Ktmw.Load(&_ktmw_CommitTransaction, "CommitTransaction"),
		uintptr(hTrans))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ktmw_CommitTransaction *syscall.Proc

// [GetTransactionId] function.
//
// [GetTransactionId]: https://learn.microsoft.com/en-us/windows/win32/api/ktmw32/nf-ktmw32-gettransactionid
func (hTrans HTRANSACTION) GetTransactionId() (co.GUID, error) {
	var guid co.GUID
	ret, _, err := syscall.SyscallN(
		dll.Ktmw.Load(&_ktmw_GetTransactionId, "GetTransactionId"),
		uintptr(hTrans),
		uintptr(unsafe.Pointer(&guid)))
	if ret == 0 {
		return co.GUID{}, co.ERROR(err)
	}
	return guid, nil
}

var _ktmw_GetTransactionId *syscall.Proc

// [RollbackTransaction] function.
//
// [RollbackTransaction]: https://learn.microsoft.com/en-us/windows/win32/api/ktmw32/nf-ktmw32-rollbacktransaction
func (hTrans HTRANSACTION) RollbackTransaction() error {
	ret, _, err := syscall.SyscallN(
		dll.Ktmw.Load(&_ktmw_RollbackTransaction, "RollbackTransaction"),
		uintptr(hTrans))
	return utl.ZeroAsGetLastError(ret, err)
}

var _ktmw_RollbackTransaction *syscall.Proc
