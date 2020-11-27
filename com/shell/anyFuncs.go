/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package shell

// Syntactic sugar; converts bool to 0 or 1.
func _BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}
