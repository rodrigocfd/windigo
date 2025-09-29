//go:build windows

package wstr_test

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/wstr"
)

func ExampleEncodeArrToBuf() {
	destBuf := make([]uint16, 7)
	goStrs := []string{"ab", "cd"}
	wstr.EncodeArrToBuf(destBuf, goStrs...)
	fmt.Println(destBuf)
	// Output: [97 98 0 99 100 0 0]
}

func ExampleEncodeArrToPtr() {
	goStrs := []string{"ab", "cd"}
	pBuf := wstr.EncodeArrToPtr(goStrs...)
	sliceBuf := unsafe.Slice(pBuf, 7)
	fmt.Println(sliceBuf)
	// Output: [97 98 0 99 100 0 0]
}

func ExampleEncodeArrToSlice() {
	goStrs := []string{"ab", "cd"}
	buf := wstr.EncodeArrToSlice(goStrs...)
	fmt.Println(buf)
	// Output: [97 98 0 99 100 0 0]
}

func ExampleEncodeToBuf() {
	destBuf := make([]uint16, 3)
	goStr := "ab"
	wstr.EncodeToBuf(destBuf, goStr)
	fmt.Println(destBuf)
	// Output: [97 98 0]
}

func ExampleEncodeToPtr() {
	goStr := "ab"
	pBuf := wstr.EncodeArrToPtr(goStr)
	sliceBuf := unsafe.Slice(pBuf, 3)
	fmt.Println(sliceBuf)
	// Output: [97 98 0]
}

func ExampleEncodeToSlice() {
	goStr := "ab"
	buf := wstr.EncodeToSlice(goStr)
	fmt.Println(buf)
	// Output: [97 98 0]
}
