//go:build windows

package wstr_test

import (
	"fmt"

	"github.com/rodrigocfd/windigo/wstr"
)

func ExampleDecodeArrPtr() {
	wideStr := []uint16{'a', 'b', 0, 'c', 'd', 0, 0}
	goStr := wstr.DecodeArrPtr(&wideStr[0])
	fmt.Println(goStr)
	// Output: [ab cd]
}

func ExampleDecodePtr() {
	wideStr := []uint16{'a', 'b', 0}
	goStr := wstr.DecodePtr(&wideStr[0])
	fmt.Println(goStr)
	// Output: ab
}

func ExampleDecodeSlice() {
	wideStr := []uint16{'a', 'b', 0}
	goStr := wstr.DecodeSlice(wideStr)
	fmt.Println(goStr)
	// Output: ab
}
