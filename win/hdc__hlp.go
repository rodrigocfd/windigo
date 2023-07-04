//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// [AtlHiMetricToPixel] function. Converts HIMETRIC units to pixels.
//
// [AtlHiMetricToPixel]: https://learn.microsoft.com/en-us/cpp/atl/reference/pixel-himetric-conversion-global-functions?view=msvc-170#atlhimetrictopixel
func (hdc HDC) HiMetricToPixel(
	himetricX, himetricY int32) (pixelX, pixelY int32) {

	// http://www.verycomputer.com/5_5f2f75dc2d090ee8_1.htm
	// https://forums.codeguru.com/showthread.php?109554-Unresizable-activeX-control

	pixelX = int32(
		(int64(himetricX) * int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSX))) /
			int64(_HIMETRIC_PER_INCH),
	)
	pixelY = int32(
		(int64(himetricY) * int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSY))) /
			int64(_HIMETRIC_PER_INCH),
	)
	return
}

// [AtlPixelToHiMetric] function. Converts pixels to HIMETRIC units.
//
// [AtlPixelToHiMetric]: https://learn.microsoft.com/en-us/cpp/atl/reference/pixel-himetric-conversion-global-functions?view=msvc-170#atlpixeltohimetric
func (hdc HDC) PixelToHiMetric(
	pixelX, pixelY int32) (himetricX, himetricY int32) {

	himetricX = int32(
		(int64(pixelX) * int64(_HIMETRIC_PER_INCH)) /
			int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSX)),
	)
	himetricY = int32(
		(int64(pixelY) * int64(_HIMETRIC_PER_INCH)) /
			int64(hdc.GetDeviceCaps(co.GDC_LOGPIXELSY)),
	)
	return
}
