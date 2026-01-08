//go:build windows

package wstr_test

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/wstr"
)

func ExampleDecodeArrPtr() {
	wideStr := []uint16{'a', 'b', 0, 'c', 'd', 0, 0}
	goStr := wstr.DecodeArrPtr(unsafe.SliceData(wideStr))
	fmt.Println(goStr)
	// Output: [ab cd]
}

func ExampleDecodePtr() {
	wideStr := []uint16{'a', 'b', 0}
	goStr := wstr.DecodePtr(unsafe.SliceData(wideStr))
	fmt.Println(goStr)
	// Output: ab
}

func ExampleDecodeSlice() {
	wideStr := []uint16{'a', 'b', 0}
	goStr := wstr.DecodeSlice(wideStr)
	fmt.Println(goStr)
	// Output: ab
}
