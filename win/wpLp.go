/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

func (wp WPARAM) LoWord() uint16      { return loWord(uint32(wp)) }
func (wp WPARAM) HiWord() uint16      { return hiWord(uint32(wp)) }
func (wp WPARAM) LoByteLoWord() uint8 { return loByte(wp.LoWord()) }
func (wp WPARAM) HiByteLoWord() uint8 { return hiByte(wp.LoWord()) }
func (wp WPARAM) LoByteHiWord() uint8 { return loByte(wp.HiWord()) }
func (wp WPARAM) HiByteHiWord() uint8 { return hiByte(wp.HiWord()) }

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

func (lp LPARAM) LoWord() uint16      { return loWord(uint32(lp)) }
func (lp LPARAM) HiWord() uint16      { return hiWord(uint32(lp)) }
func (lp LPARAM) LoByteLoWord() uint8 { return loByte(lp.LoWord()) }
func (lp LPARAM) HiByteLoWord() uint8 { return hiByte(lp.LoWord()) }
func (lp LPARAM) LoByteHiWord() uint8 { return loByte(lp.HiWord()) }
func (lp LPARAM) HiByteHiWord() uint8 { return hiByte(lp.HiWord()) }
func (lp LPARAM) MakePoint() POINT {
	return POINT{
		X: int32(lp.LoWord()),
		Y: int32(lp.HiWord()),
	}
}
func (lp LPARAM) MakeSize() SIZE {
	return SIZE{
		Cx: int32(lp.LoWord()),
		Cy: int32(lp.HiWord()),
	}
}
