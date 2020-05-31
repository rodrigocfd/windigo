/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package co

type PRF uint32 // WM_PRINT

const (
	PRF_CHECKVISIBLE PRF = 0x00000001
	PRF_NONCLIENT    PRF = 0x00000002
	PRF_CLIENT       PRF = 0x00000004
	PRF_ERASEBKGND   PRF = 0x00000008
	PRF_CHILDREN     PRF = 0x00000010
	PRF_OWNED        PRF = 0x00000020
)

type PT uint8 // PolyDraw

const (
	PT_CLOSEFIGURE PT = 0x01
	PT_LINETO      PT = 0x02
	PT_BEZIERTO    PT = 0x04
	PT_MOVETO      PT = 0x06
)
