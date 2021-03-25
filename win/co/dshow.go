package co

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/strmif/ne-strmif-am_seeking_seeking_capabilities
type AM_SEEKING uint32

const (
	AM_SEEKING_CanSeekAbsolute  AM_SEEKING = 0x1
	AM_SEEKING_CanSeekForwards  AM_SEEKING = 0x2
	AM_SEEKING_CanSeekBackwards AM_SEEKING = 0x4
	AM_SEEKING_CanGetCurrentPos AM_SEEKING = 0x8
	AM_SEEKING_CanGetStopPos    AM_SEEKING = 0x10
	AM_SEEKING_CanGetDuration   AM_SEEKING = 0x20
	AM_SEEKING_CanPlayBackwards AM_SEEKING = 0x40
	AM_SEEKING_CanDoSegments    AM_SEEKING = 0x80
	AM_SEEKING_Source           AM_SEEKING = 0x100
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
