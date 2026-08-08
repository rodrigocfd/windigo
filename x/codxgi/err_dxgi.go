//go:build windows

package codxgi

import (
	"github.com/rodrigocfd/windigo/co"
)

const (
	HRESULT_DXGI_ERROR_INVALID_CALL                  co.HRESULT = 0x887a_0001 // The application made a call that is invalid. Either the parameters of the call or the state of some object was incorrect. Enable the D3D debug layer in order to see details via debug messages.
	HRESULT_DXGI_ERROR_NOT_FOUND                     co.HRESULT = 0x887a_0002 // The object was not found. If calling IDXGIFactory::EnumAdaptes, there is no adapter with the specified ordinal.
	HRESULT_DXGI_ERROR_MORE_DATA                     co.HRESULT = 0x887a_0003 // The caller did not supply a sufficiently large buffer.
	HRESULT_DXGI_ERROR_UNSUPPORTED                   co.HRESULT = 0x887a_0004 // The specified device interface or feature level is not supported on this system.
	HRESULT_DXGI_ERROR_DEVICE_REMOVED                co.HRESULT = 0x887a_0005 // The GPU device instance has been suspended. Use GetDeviceRemovedReason to determine the appropriate action.
	HRESULT_DXGI_ERROR_DEVICE_HUNG                   co.HRESULT = 0x887a_0006 // The GPU will not respond to more commands, most likely because of an invalid command passed by the calling application.
	HRESULT_DXGI_ERROR_DEVICE_RESET                  co.HRESULT = 0x887a_0007 // The GPU will not respond to more commands, most likely because some other application submitted invalid commands. The calling application should re-create the device and continue.
	HRESULT_DXGI_ERROR_WAS_STILL_DRAWING             co.HRESULT = 0x887a_000a // The GPU was busy at the moment when the call was made, and the call was neither executed nor scheduled.
	HRESULT_DXGI_ERROR_FRAME_STATISTICS_DISJOINT     co.HRESULT = 0x887a_000b // An event (such as power cycle) interrupted the gathering of presentation statistics. Any previous statistics should be considered invalid.
	HRESULT_DXGI_ERROR_GRAPHICS_VIDPN_SOURCE_IN_USE  co.HRESULT = 0x887a_000c // Fullscreen mode could not be achieved because the specified output was already in use.
	HRESULT_DXGI_ERROR_DRIVER_INTERNAL_ERROR         co.HRESULT = 0x887a_0020 // An internal issue prevented the driver from carrying out the specified operation. The driver's state is probably suspect, and the application should not continue.
	HRESULT_DXGI_ERROR_NONEXCLUSIVE                  co.HRESULT = 0x887a_0021 // A global counter resource was in use, and the specified counter cannot be used by this Direct3D device at this time.
	HRESULT_DXGI_ERROR_NOT_CURRENTLY_AVAILABLE       co.HRESULT = 0x887a_0022 // A resource is not available at the time of the call, but may become available later.
	HRESULT_DXGI_ERROR_REMOTE_CLIENT_DISCONNECTED    co.HRESULT = 0x887a_0023 // The application's remote device has been removed due to session disconnect or network disconnect. The application should call IDXGIFactory1::IsCurrent to find out when the remote device becomes available again.
	HRESULT_DXGI_ERROR_REMOTE_OUTOFMEMORY            co.HRESULT = 0x887a_0024 // The device has been removed during a remote session because the remote computer ran out of memory.
	HRESULT_DXGI_ERROR_ACCESS_LOST                   co.HRESULT = 0x887a_0026 // The keyed mutex was abandoned.
	HRESULT_DXGI_ERROR_WAIT_TIMEOUT                  co.HRESULT = 0x887a_0027 // The timeout value has elapsed and the resource is not yet available.
	HRESULT_DXGI_ERROR_SESSION_DISCONNECTED          co.HRESULT = 0x887a_0028 // The output duplication has been turned off because the Windows session ended or was disconnected. This happens when a remote user disconnects, or when "switch user" is used locally.
	HRESULT_DXGI_ERROR_RESTRICT_TO_OUTPUT_STALE      co.HRESULT = 0x887a_0029 // The DXGI output (monitor) to which the swapchain content was restricted, has been disconnected or changed.
	HRESULT_DXGI_ERROR_CANNOT_PROTECT_CONTENT        co.HRESULT = 0x887a_002a // DXGI is unable to provide content protection on the swapchain. This is typically caused by an older driver, or by the application using a swapchain that is incompatible with content protection.
	HRESULT_DXGI_ERROR_ACCESS_DENIED                 co.HRESULT = 0x887a_002b // The application is trying to use a resource to which it does not have the required access privileges. This is most commonly caused by writing to a shared resource with read-only access.
	HRESULT_DXGI_ERROR_NAME_ALREADY_EXISTS           co.HRESULT = 0x887a_002c // The application is trying to create a shared handle using a name that is already associated with some other resource.
	HRESULT_DXGI_ERROR_SDK_COMPONENT_MISSING         co.HRESULT = 0x887a_002d // The application requested an operation that depends on an SDK component that is missing or mismatched.
	HRESULT_DXGI_ERROR_NOT_CURRENT                   co.HRESULT = 0x887a_002e // The DXGI objects that the application has created are no longer current & need to be recreated for this operation to be performed.
	HRESULT_DXGI_ERROR_HW_PROTECTION_OUTOFMEMORY     co.HRESULT = 0x887a_0030 // Insufficient HW protected memory exits for proper function.
	HRESULT_DXGI_ERROR_DYNAMIC_CODE_POLICY_VIOLATION co.HRESULT = 0x887a_0031 // Creating this device would violate the process's dynamic code policy.
	HRESULT_DXGI_ERROR_NON_COMPOSITED_UI             co.HRESULT = 0x887a_0032 // The operation failed because the compositor is not in control of the output.
	HRESULT_DXGI_ERROR_MODE_CHANGE_IN_PROGRESS       co.HRESULT = 0x887a_0025 // An on-going mode change prevented completion of the call. The call may succeed if attempted later.
	HRESULT_DXGI_ERROR_CACHE_CORRUPT                 co.HRESULT = 0x887a_0033 // The cache is corrupt and either could not be opened or could not be reset.
	HRESULT_DXGI_ERROR_CACHE_FULL                    co.HRESULT = 0x887a_0034 // This entry would cause the cache to exceed its quota. On a load operation, this may indicate exceeding the maximum in-memory size.
	HRESULT_DXGI_ERROR_CACHE_HASH_COLLISION          co.HRESULT = 0x887a_0035 // A cache entry was found, but the key provided does not match the key stored in the entry.
	HRESULT_DXGI_ERROR_ALREADY_EXISTS                co.HRESULT = 0x887a_0036 // The desired element already exists.
	HRESULT_DXGI_ERROR_MPO_UNPINNED                  co.HRESULT = 0x887a_0064 // The allocation of the MPO plane has been unpinned.
	HRESULT_DXGI_ERROR_SETDISPLAYMODE_REQUIRED       co.HRESULT = 0x887a_0065 // SetDisplayMode is required before present can succeed.
)
