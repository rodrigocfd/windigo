package win

import (
	"runtime"
	"testing"
)

func nuuup() string {
	n0 := "foo bar"
	n1 := n0 + "lenses"
	return n1
}

func BenchmarkFoo(b *testing.B) {
	runtime.LockOSThread()
	for n := 0; n < b.N; n++ {
		HINSTANCE(0).GetModuleFileName()
		// ll := SYSTEMTIME{}
		// GetSystemTime(&ll)
		// nuuup()
	}
}
