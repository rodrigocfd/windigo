//go:build windows

package utl

import (
	"syscall"

	"github.com/rodrigocfd/windigo/co"
)

type IDispatchVt struct {
	IUnknownVt
	GetTypeInfoCount uintptr
	GetTypeInfo      uintptr
	GetIDsOfNames    uintptr
	Invoke           uintptr
}

type ISequentialStreamVt struct {
	IUnknownVt
	Read  uintptr
	Write uintptr
}

type IStreamVt struct {
	ISequentialStreamVt
	Seek         uintptr
	SetSize      uintptr
	CopyTo       uintptr
	Commit       uintptr
	Revert       uintptr
	LockRegion   uintptr
	UnlockRegion uintptr
	Stat         uintptr
	Clone        uintptr
}

type IUnknownVt struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

// IUnknown.QueryInterface method for custom-implemented interfaces.
var _queryInterfaceImpl uintptr

func OleQueryInterfaceImpl() uintptr {
	if _queryInterfaceImpl == 0 {
		_queryInterfaceImpl = syscall.NewCallback(
			func(_p uintptr, _riid uintptr, ppv ***IUnknownVt) uintptr {
				*ppv = nil
				return uintptr(co.HRESULT_E_NOTIMPL)
			},
		)
	}
	return _queryInterfaceImpl
}
