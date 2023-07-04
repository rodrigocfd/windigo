//go:build windows

package dshowvt

import (
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/com/comvt"
)

// [IBaseFilter] virtual table.
//
// [IBaseFilter]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
type IBaseFilter struct {
	IMediaFilter
	EnumPins        uintptr
	FindPin         uintptr
	QueryFilterInfo uintptr
	JoinFilterGraph uintptr
	QueryVendorInfo uintptr
}

// [IBasicAudio] virtual table.
//
// [IBasicAudio]: https://learn.microsoft.com/en-us/windows/win32/api/control/nn-control-ibasicaudio
type IBasicAudio struct {
	automvt.IDispatch
	PutVolume  uintptr
	GetVolume  uintptr
	PutBalance uintptr
	GetBalance uintptr
}

// [IEnumFilters] virtual table.
//
// [IEnumFilters]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumfilters
type IEnumFilters struct {
	comvt.IUnknown
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// [IEnumMediaTypes] virtual table.
//
// [IEnumMediaTypes]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienummediatypes
type IEnumMediaTypes struct {
	comvt.IUnknown
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// [IEnumPins] virtual table.
//
// [IEnumPins]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumpins
type IEnumPins struct {
	comvt.IUnknown
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// [IFileSinkFilter] virtual table.
//
// [IFileSinkFilter]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter
type IFileSinkFilter struct {
	comvt.IUnknown
	SetFileName uintptr
	GetCurFile  uintptr
}

// [IFileSinkFilter2] virtual table.
//
// [IFileSinkFilter2]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter2
type IFileSinkFilter2 struct {
	IFileSinkFilter
	SetMode uintptr
	GetMode uintptr
}

// [IFileSourceFilter] virtual table.
//
// [IFileSourceFilter]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesourcefilter
type IFileSourceFilter struct {
	comvt.IUnknown
	Load       uintptr
	GetCurFile uintptr
}

// [IFilterGraph] virtual table.
//
// [IFilterGraph]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
type IFilterGraph struct {
	comvt.IUnknown
	AddFilter            uintptr
	RemoveFilter         uintptr
	EnumFilters          uintptr
	FindFilterByName     uintptr
	ConnectDirect        uintptr
	Reconnect            uintptr
	Disconnect           uintptr
	SetDefaultSyncSource uintptr
}

// [IGraphBuilder] virtual table.
//
// [IGraphBuilder]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
type IGraphBuilder struct {
	IFilterGraph
	Connect                 uintptr
	Render                  uintptr
	RenderFile              uintptr
	AddSourceFilter         uintptr
	SetLogFile              uintptr
	Abort                   uintptr
	ShouldOperationContinue uintptr
}

// [IMediaControl] virtual table.
//
// [IMediaControl]: https://learn.microsoft.com/en-us/windows/win32/api/control/nn-control-imediacontrol
type IMediaControl struct {
	automvt.IDispatch
	Run                    uintptr
	Pause                  uintptr
	Stop                   uintptr
	GetState               uintptr
	RenderFile             uintptr
	AddSourceFilter        uintptr
	GetFilterCollection    uintptr
	GetRegFilterCollection uintptr
	StopWhenReady          uintptr
}

// [IMediaFilter] virtual table.
//
// [IMediaFilter]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
type IMediaFilter struct {
	comvt.IPersist
	Stop          uintptr
	Pause         uintptr
	Run           uintptr
	GetState      uintptr
	SetSyncSource uintptr
	GetSyncSource uintptr
}

// [IMediaSeeking] virtual table.
//
// [IMediaSeeking]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediaseeking
type IMediaSeeking struct {
	comvt.IUnknown
	GetCapabilities      uintptr
	CheckCapabilities    uintptr
	IsFormatSupported    uintptr
	QueryPreferredFormat uintptr
	GetTimeFormat        uintptr
	IsUsingTimeFormat    uintptr
	SetTimeFormat        uintptr
	GetDuration          uintptr
	GetStopPosition      uintptr
	GetCurrentPosition   uintptr
	ConvertTimeFormat    uintptr
	SetPositions         uintptr
	GetPositions         uintptr
	GetAvailable         uintptr
	SetRate              uintptr
	GetRate              uintptr
	GetPreroll           uintptr
}

// [IMFGetService] virtual table.
//
// [IMFGetService]: https://learn.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService struct {
	comvt.IUnknown
	GetService uintptr
}

// [IMFVideoDisplayControl] virtual table.
//
// [IMFVideoDisplayControl]: https://learn.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
type IMFVideoDisplayControl struct {
	comvt.IUnknown
	GetNativeVideoSize uintptr
	GetIdealVideoSize  uintptr
	SetVideoPosition   uintptr
	GetVideoPosition   uintptr
	SetAspectRatioMode uintptr
	GetAspectRatioMode uintptr
	SetVideoWindow     uintptr
	GetVideoWindow     uintptr
	RepaintVideo       uintptr
	GetCurrentImage    uintptr
	SetBorderColor     uintptr
	GetBorderColor     uintptr
	SetRenderingPrefs  uintptr
	GetRenderingPrefs  uintptr
	SetFullscreen      uintptr
	GetFullscreen      uintptr
}

// [IPin] virtual table.
//
// [IPin]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ipin
type IPin struct {
	comvt.IUnknown
	Connect                  uintptr
	ReceiveConnection        uintptr
	Disconnect               uintptr
	ConnectedTo              uintptr
	ConnectionMediaType      uintptr
	QueryPinInfo             uintptr
	QueryDirection           uintptr
	QueryId                  uintptr
	QueryAccept              uintptr
	EnumMediaTypes           uintptr
	QueryInternalConnections uintptr
	EndOfStream              uintptr
	BeginFlush               uintptr
	EndFlush                 uintptr
	NewSegment               uintptr
}
