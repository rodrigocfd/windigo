//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/co"
	"github.com/rodrigocfd/windigo/wstr"
)

// [DXGI_ADAPTER_DESC] struct.
//
// [DXGI_ADAPTER_DESC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_adapter_desc
type DXGI_ADAPTER_DESC struct {
	description           [128]uint16
	VendorId              uint32
	DeviceId              uint32
	SubSysId              uint32
	Revision              uint32
	DedicatedVideoMemory  uint
	DedicatedSystemMemory uint
	SharedSystemMemory    uint
	AdapterLuid           co.LUID
}

func (ad *DXGI_ADAPTER_DESC) Description() string {
	return wstr.DecodeSlice(ad.description[:])
}
func (ad *DXGI_ADAPTER_DESC) SetDescription(val string) {
	wstr.EncodeToBuf(ad.description[:], val)
}

// [DXGI_ADAPTER_DESC1] struct.
//
// [DXGI_ADAPTER_DESC1]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_adapter_desc1
type DXGI_ADAPTER_DESC1 struct {
	description           [128]uint16
	VendorId              uint32
	DeviceId              uint32
	SubSysId              uint32
	Revision              uint32
	DedicatedVideoMemory  uint
	DedicatedSystemMemory uint
	SharedSystemMemory    uint
	AdapterLuid           co.LUID
	Flags                 co.DXGI_ADAPTER_FLAG
}

func (ad *DXGI_ADAPTER_DESC1) Description() string {
	return wstr.DecodeSlice(ad.description[:])
}
func (ad *DXGI_ADAPTER_DESC1) SetDescription(val string) {
	wstr.EncodeToBuf(ad.description[:], val)
}

// [DXGI_FRAME_STATISTICS] struct.
//
// [DXGI_FRAME_STATISTICS]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_frame_statistics)
type DXGI_FRAME_STATISTICS struct {
	PresentCount        uint32
	PresentRefreshCount uint32
	SyncRefreshCount    uint32
	SyncQPCTime         int64
	SyncGPUTime         int64
}

// [DXGI_GAMMA_CONTROL] struct.
//
// [DXGI_GAMMA_CONTROL]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/bb173061(v=vs.85)
type DXGI_GAMMA_CONTROL struct {
	Scale      DXGI_RGB
	Offset     DXGI_RGB
	GammaCurve [1025]DXGI_RGB
}

// [DXGI_GAMMA_CONTROL_CAPABILITIES] struct.
//
// [DXGI_GAMMA_CONTROL_CAPABILITIES]: https://learn.microsoft.com/en-us/windows-hardware/drivers/ddi/dxgitype/ns-dxgitype-dxgi_gamma_control_capabilities
type DXGI_GAMMA_CONTROL_CAPABILITIES struct {
	ScaleAndOffsetSupported BOOL
	MaxConvertedValue       float32
	MinConvertedValue       float32
	NumGammaControlPoints   uint32
	ControlPointPositions   [1025]float32
}

// [DXGI_MAPPED_RECT] struct.
//
// [DXGI_MAPPED_RECT]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_mapped_rect
type DXGI_MAPPED_RECT struct {
	Pitch int32
	PBits *byte
}

// [DXGI_MODE_DESC] struct.
//
// [DXGI_MODE_DESC]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/bb173064(v=vs.85)
type DXGI_MODE_DESC struct {
	Width            uint32
	Height           uint32
	RefreshRate      DXGI_RATIONAL
	Format           co.DXGI_FORMAT
	ScanlineOrdering co.DXGI_MODE_SCANLINE_ORDER
	Scaling          co.DXGI_MODE_SCALING
}

// [DXGI_OUTPUT_DESC] struct.
//
// [DXGI_OUTPUT_DESC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_output_desc
type DXGI_OUTPUT_DESC struct {
	deviceName         [32]uint16
	DesktopCoordinates RECT
	AttachedToDesktop  BOOL
	Rotation           co.DXGI_MODE_ROTATION
	Monitor            HMONITOR
}

func (od *DXGI_OUTPUT_DESC) DeviceName() string {
	return wstr.DecodeSlice(od.deviceName[:])
}
func (od *DXGI_OUTPUT_DESC) SetDeviceName(val string) {
	wstr.EncodeToBuf(od.deviceName[:], val)
}

// [DXGI_RATIONAL] struct.
//
// [DXGI_RATIONAL]: https://learn.microsoft.com/en-us/windows/win32/api/dxgicommon/ns-dxgicommon-dxgi_rational
type DXGI_RATIONAL struct {
	Numerator   uint32
	Denominator uint32
}

// [DXGI_RGB] struct.
//
// [DXGI_RGB]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/bb173071(v=vs.85)
type DXGI_RGB struct {
	Red   float32
	Green float32
	Blue  float32
}

// [DXGI_SAMPLE_DESC] struct.
//
// [DXGI_SAMPLE_DESC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgicommon/ns-dxgicommon-dxgi_sample_desc
type DXGI_SAMPLE_DESC struct {
	Count   uint32
	Quality uint32
}

// [DXGI_SURFACE_DESC] struct.
//
// [DXGI_SURFACE_DESC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_surface_desc
type DXGI_SURFACE_DESC struct {
	Width      uint32
	Height     uint32
	Format     co.DXGI_FORMAT
	SampleDesc DXGI_SAMPLE_DESC
}

// [DXGI_SWAP_CHAIN_DESC] struct.
//
// [DXGI_SWAP_CHAIN_DESC]: https://learn.microsoft.com/en-us/windows/win32/api/dxgi/ns-dxgi-dxgi_swap_chain_desc
type DXGI_SWAP_CHAIN_DESC struct {
	BufferDesc   DXGI_MODE_DESC
	SampleDesc   DXGI_SAMPLE_DESC
	BufferUsage  co.DXGI_USAGE
	BufferCount  uint32
	OutputWindow HWND
	Windowed     BOOL
	SwapEffect   co.DXGI_SWAP_EFFECT
	Flags        co.DXGI_SWAP_CHAIN_FLAG
}
