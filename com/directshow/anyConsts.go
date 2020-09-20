/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package directshow

// MFVideoAspectRatioMode.
type MFVideoARMode uint32

const (
	MFVideoARMode_None             MFVideoARMode = 0
	MFVideoARMode_PreservePicture  MFVideoARMode = 0x1
	MFVideoARMode_PreservePixel    MFVideoARMode = 0x2
	MFVideoARMode_NonLinearStretch MFVideoARMode = 0x4
	MFVideoARMode_Mask             MFVideoARMode = 0x7
)
