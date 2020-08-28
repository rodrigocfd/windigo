/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package win

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#wparam
type WPARAM uintptr

func (wp WPARAM) LoWord() uint16      { return LoWord(uint32(wp)) }
func (wp WPARAM) HiWord() uint16      { return HiWord(uint32(wp)) }
func (wp WPARAM) LoByteLoWord() uint8 { return LoByte(wp.LoWord()) }
func (wp WPARAM) HiByteLoWord() uint8 { return HiByte(wp.LoWord()) }
func (wp WPARAM) LoByteHiWord() uint8 { return LoByte(wp.HiWord()) }
func (wp WPARAM) HiByteHiWord() uint8 { return HiByte(wp.HiWord()) }

// https://docs.microsoft.com/en-us/windows/win32/winprog/windows-data-types#lparam
type LPARAM uintptr

func (lp LPARAM) LoWord() uint16      { return LoWord(uint32(lp)) }
func (lp LPARAM) HiWord() uint16      { return HiWord(uint32(lp)) }
func (lp LPARAM) LoByteLoWord() uint8 { return LoByte(lp.LoWord()) }
func (lp LPARAM) HiByteLoWord() uint8 { return HiByte(lp.LoWord()) }
func (lp LPARAM) LoByteHiWord() uint8 { return LoByte(lp.HiWord()) }
func (lp LPARAM) HiByteHiWord() uint8 { return HiByte(lp.HiWord()) }

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
