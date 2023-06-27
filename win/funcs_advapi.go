//go:build windows

package win

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/internal/proc"
	"github.com/rodrigocfd/windigo/win/errco"
)

// [GetManagedApplications] function.
//
// [GetManagedApplications]: https://learn.microsoft.com/en-us/windows/win32/api/appmgmt/nf-appmgmt-getmanagedapplications
func GetManagedApplications(category *GUID) ([]MANAGEDAPPLICATION, error) {
	flags := _MANAGED_APPS_USERAPPLICATIONS
	if category != nil {
		flags = _MANAGED_APPS_FROMCATEGORY
	}

	level := _MANAGED_APPS_INFOLEVEL_DEFAULT

	var count uint32
	var pMa *MANAGEDAPPLICATION
	defer HLOCAL(unsafe.Pointer(pMa)).LocalFree()

	ret, _, _ := syscall.SyscallN(proc.GetManagedApplications.Addr(),
		uintptr(unsafe.Pointer(category)), uintptr(flags), uintptr(level),
		uintptr(unsafe.Pointer(&count)), uintptr(unsafe.Pointer(&pMa)))

	if wErr := errco.ERROR(ret); ret != uintptr(errco.SUCCESS) {
		return nil, wErr
	}

	localSlice := unsafe.Slice(pMa, count)
	retSlice := make([]MANAGEDAPPLICATION, count)
	for idx := range localSlice {
		retSlice[idx] = localSlice[idx]
	}
	return retSlice, nil
}
