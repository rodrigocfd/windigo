//go:build windows

package dshowco

// [IFileSinkFilter2] modes.
//
// [IFileSinkFilter2]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-ifilesinkfilter2-setmode
type AM_FILE uint32

const (
	AM_FILE_NONE      AM_FILE = 0
	AM_FILE_OVERWRITE AM_FILE = 0x1
)

// [FILTER_STATE] enumeration.
//
// [FILTER_STATE]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/ne-strmif-filter_state
type FILTER_STATE uint32

const (
	FILTER_STATE_State_Stopped FILTER_STATE = iota
	FILTER_STATE_State_Paused
	FILTER_STATE_State_Running
)

// [MFVideoAspectRatioMode] enumeration.
//
// [MFVideoAspectRatioMode]: https://learn.microsoft.com/en-us/windows/win32/api/evr/ne-evr-mfvideoaspectratiomode
type MFVideoARMode uint32

const (
	MFVideoARMode_None             MFVideoARMode = 0   // Do not maintain the aspect ratio of the video. Stretch the video to fit the output rectangle.
	MFVideoARMode_PreservePicture  MFVideoARMode = 0x1 // Preserve the aspect ratio of the video by letterboxing or within the output rectangle.
	MFVideoARMode_PreservePixel    MFVideoARMode = 0x2 // Currently the EVR ignores this flag.
	MFVideoARMode_NonLinearStretch MFVideoARMode = 0x4 // Apply a non-linear horizontal stretch if the aspect ratio of the destination rectangle does not match the aspect ratio of the source rectangle.
)

// [PIN_DIRECTION] enumeration.
//
// [PIN_DIRECTION]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/ne-strmif-pin_direction
type PIN_DIRECTION uint32

const (
	PIN_DIRECTION_INPUT PIN_DIRECTION = iota
	PIN_DIRECTION_OUTOUT
)

// [AM_SEEKING_SEEKING_CAPABILITIES] enumeration.
//
// [AM_SEEKING_SEEKING_CAPABILITIES]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/ne-strmif-am_seeking_seeking_capabilities
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

// [IMediaSeeking.SetPositions] flags.
//
// [IMediaSeeking.SetPositions]: https://learn.microsoft.com/en-us/windows/win32/api/strmif/nf-strmif-imediaseeking-setpositions
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
