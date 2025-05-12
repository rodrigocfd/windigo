//go:build windows

package utl

// Syntactic sugar; converts bool to 0 or 1.
func BoolToInt32(b bool) int32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUint32(b bool) uint32 {
	if b {
		return 1
	}
	return 0
}

// Syntactic sugar; converts bool to 0 or 1.
func BoolToUintptr(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}
