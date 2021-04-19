package co

const (
	CLSID_EnhancedVideoRenderer CLSID = "fa10746c-9b63-4b6c-bc49-fc300ea5f256"
	CLSID_FilterGraph           CLSID = "e436ebb3-524f-11ce-9f53-0020af0ba770"
	CLSID_VideoMixingRenderer9  CLSID = "51b4abf3-748f-4e3b-a276-c828330e926a"

	IID_IBaseFilter            IID = "56a86895-0ad4-11ce-b03a-0020af0ba770"
	IID_IBasicAudio            IID = "56a868b3-0ad4-11ce-b03a-0020af0ba770"
	IID_IFilterGraph           IID = "56a8689f-0ad4-11ce-b03a-0020af0ba770"
	IID_IGraphBuilder          IID = "56a868a9-0ad4-11ce-b03a-0020af0ba770"
	IID_IMediaControl          IID = "56a868b1-0ad4-11ce-b03a-0020af0ba770"
	IID_IMediaSeeking          IID = "36b73880-c2c8-11cf-8b46-00805f6cef60"
	IID_IMFGetService          IID = "fa993888-4383-415a-a930-dd472a8cf6f7"
	IID_IMFVideoDisplayControl IID = "a490b1e4-ab84-4d31-a1b2-181e03b1077a"

	MR_VideoRenderService CLSID = "1092a86c-ab1a-459a-a336-831fbc4d11ff"
)

// Originally AM_SEEKING_SeekingCapabilities enum.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ne-strmif-SEEKING_FLAGS_capabilities
type SEEKING_CAPABILITIES uint32

const (
	SEEKING_CAPABILITIES_CanSeekAbsolute  SEEKING_CAPABILITIES = 0x1
	SEEKING_CAPABILITIES_CanSeekForwards  SEEKING_CAPABILITIES = 0x2
	SEEKING_CAPABILITIES_CanSeekBackwards SEEKING_CAPABILITIES = 0x4
	SEEKING_CAPABILITIES_CanGetCurrentPos SEEKING_CAPABILITIES = 0x8
	SEEKING_CAPABILITIES_CanGetStopPos    SEEKING_CAPABILITIES = 0x10
	SEEKING_CAPABILITIES_CanGetDuration   SEEKING_CAPABILITIES = 0x20
	SEEKING_CAPABILITIES_CanPlayBackwards SEEKING_CAPABILITIES = 0x40
	SEEKING_CAPABILITIES_CanDoSegments    SEEKING_CAPABILITIES = 0x80
	SEEKING_CAPABILITIES_Source           SEEKING_CAPABILITIES = 0x100
)

// IMediaSeeking.SetPositions() flags. Originally AM_SEEKING_SeekingFlags enum.
//
// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setpositions
type SEEKING_FLAGS uint32

const (
	SEEKING_FLAGS_NoPositioning          SEEKING_FLAGS = 0x0
	SEEKING_FLAGS_AbsolutePositioning    SEEKING_FLAGS = 0x1
	SEEKING_FLAGS_RelativePositioning    SEEKING_FLAGS = 0x2
	SEEKING_FLAGS_IncrementalPositioning SEEKING_FLAGS = 0x3
	SEEKING_FLAGS_SeekToKeyFrame         SEEKING_FLAGS = 0x4
	SEEKING_FLAGS_ReturnTime             SEEKING_FLAGS = 0x8
	SEEKING_FLAGS_Segment                SEEKING_FLAGS = 0x10
	SEEKING_FLAGS_NoFlush                SEEKING_FLAGS = 0x20
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ne-strmif-filter_state
type FILTER_STATE int

const (
	FILTER_STATE_State_Stopped FILTER_STATE = 0
	FILTER_STATE_State_Paused  FILTER_STATE = 1
	FILTER_STATE_State_Running FILTER_STATE = 2
)

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/evr/ne-evr-mfvideoaspectratiomode
type MFVideoARMode uint32

const (
	MFVideoARMode_None             MFVideoARMode = 0   // Do not maintain the aspect ratio of the video. Stretch the video to fit the output rectangle.
	MFVideoARMode_PreservePicture  MFVideoARMode = 0x1 // Preserve the aspect ratio of the video by letterboxing or within the output rectangle.
	MFVideoARMode_PreservePixel    MFVideoARMode = 0x2 // Currently the EVR ignores this flag.
	MFVideoARMode_NonLinearStretch MFVideoARMode = 0x4 // Apply a non-linear horizontal stretch if the aspect ratio of the destination rectangle does not match the aspect ratio of the source rectangle.
)
