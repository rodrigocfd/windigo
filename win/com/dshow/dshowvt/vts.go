package dshowvt

import (
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/com/autom/automvt"
	"github.com/rodrigocfd/windigo/win/com/idl/idlvt"
)

// IBaseFilter virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ibasefilter
type IBaseFilter struct {
	IMediaFilter
	EnumPins        uintptr
	FindPin         uintptr
	QueryFilterInfo uintptr
	JoinFilterGraph uintptr
	QueryVendorInfo uintptr
}

// IBasicAudio virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-ibasicaudio
type IBasicAudio struct {
	automvt.IDispatch
	PutVolume  uintptr
	GetVolume  uintptr
	PutBalance uintptr
	GetBalance uintptr
}

// IEnumFilters virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumfilters
type IEnumFilters struct {
	win.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// IEnumMediaTypes virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienummediatypes
type IEnumMediaTypes struct {
	win.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// IEnumPins virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ienumpins
type IEnumPins struct {
	win.IUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}

// IFileSinkFilter virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter
type IFileSinkFilter struct {
	win.IUnknownVtbl
	SetFileName uintptr
	GetCurFile  uintptr
}

// IFileSinkFilter2 virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesinkfilter2
type IFileSinkFilter2 struct {
	IFileSinkFilter
	SetMode uintptr
	GetMode uintptr
}

// IFileSourceFilter virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifilesourcefilter
type IFileSourceFilter struct {
	win.IUnknownVtbl
	Load       uintptr
	GetCurFile uintptr
}

// IFilterGraph virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-ifiltergraph
type IFilterGraph struct {
	win.IUnknownVtbl
	AddFilter            uintptr
	RemoveFilter         uintptr
	EnumFilters          uintptr
	FindFilterByName     uintptr
	ConnectDirect        uintptr
	Reconnect            uintptr
	Disconnect           uintptr
	SetDefaultSyncSource uintptr
}

// IGraphBuilder virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-igraphbuilder
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

// IMediaControl virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-imediacontrol
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

// IMediaFilter virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediafilter
type IMediaFilter struct {
	idlvt.IPersist
	Stop          uintptr
	Pause         uintptr
	Run           uintptr
	GetState      uintptr
	SetSyncSource uintptr
	GetSyncSource uintptr
}

// IMediaSeeking virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nn-strmif-imediaseeking
type IMediaSeeking struct {
	win.IUnknownVtbl
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

// IMFGetService virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/mfidl/nn-mfidl-imfgetservice
type IMFGetService struct {
	win.IUnknownVtbl
	GetService uintptr
}

// IMFVideoDisplayControl virtual table.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/nn-evr-imfvideodisplaycontrol
type IMFVideoDisplayControl struct {
	win.IUnknownVtbl
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
