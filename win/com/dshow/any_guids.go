package dshow

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// DirectShow COM CLSIDs.
var CLSID = struct {
	EnhancedVideoRenderer co.CLSID
	FilterGraph           co.CLSID
	MR_VideoRenderService co.CLSID
	VideoMixingRenderer9  co.CLSID
}{
	EnhancedVideoRenderer: "fa10746c-9b63-4b6c-bc49-fc300ea5f256",
	FilterGraph:           "e436ebb3-524f-11ce-9f53-0020af0ba770",
	MR_VideoRenderService: "1092a86c-ab1a-459a-a336-831fbc4d11ff",
	VideoMixingRenderer9:  "51b4abf3-748f-4e3b-a276-c828330e926a",
}

// DirectShow COM IIDs.
var IID = struct {
	IBaseFilter            co.IID
	IBasicAudio            co.IID
	IEnumFilters           co.IID
	IEnumMediaTypes        co.IID
	IFilterGraph           co.IID
	IGraphBuilder          co.IID
	IMediaControl          co.IID
	IMediaFilter           co.IID
	IMediaSeeking          co.IID
	IMFGetService          co.IID
	IMFVideoDisplayControl co.IID
	IPersist               co.IID
	IPin                   co.IID
}{
	IBaseFilter:            "56a86895-0ad4-11ce-b03a-0020af0ba770",
	IBasicAudio:            "56a868b3-0ad4-11ce-b03a-0020af0ba770",
	IEnumFilters:           "56a86893-0ad4-11ce-b03a-0020af0ba770",
	IEnumMediaTypes:        "89c31040-846b-11ce-97d3-00aa0055595a",
	IFilterGraph:           "56a8689f-0ad4-11ce-b03a-0020af0ba770",
	IGraphBuilder:          "56a868a9-0ad4-11ce-b03a-0020af0ba770",
	IMediaControl:          "56a868b1-0ad4-11ce-b03a-0020af0ba770",
	IMediaFilter:           "56a86899-0ad4-11ce-b03a-0020af0ba770",
	IMediaSeeking:          "36b73880-c2c8-11cf-8b46-00805f6cef60",
	IMFGetService:          "fa993888-4383-415a-a930-dd472a8cf6f7",
	IMFVideoDisplayControl: "a490b1e4-ab84-4d31-a1b2-181e03b1077a",
	IPersist:               "0000010c-0000-0000-c000-000000000046",
	IPin:                   "56a86891-0ad4-11ce-b03a-0020af0ba770",
}
