//go:build windows

package wstr_test

import (
	"fmt"
	"unsafe"

	"github.com/rodrigocfd/windigo/wstr"
)

func ExampleBufDecoder() {
	var wBuf wstr.BufDecoder
	wBuf.Alloc(20)

	wBuf.HotSlice()[0] = 'a'
	wBuf.HotSlice()[1] = 'b'

	goStr := wBuf.String()
	fmt.Println(goStr)
	// Output: ab
}

func ExampleBufEncoder() {
	var wBuf wstr.BufEncoder
	p := wBuf.AllowEmpty("ab")

	wideStr := unsafe.Slice((*uint16)(p), 2)
	fmt.Println(string(rune(wideStr[0])), string(rune(wideStr[1])))
	// Output: a b
}

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

func ExampleCapitalize() {
	s := wstr.Capitalize("abc")
	fmt.Println(s)
	// Output: Abc
}

func ExampleCmp() {
	compare1 := wstr.Cmp("aa", "bb")
	compare2 := wstr.Cmp("bb", "aa")
	compare3 := wstr.Cmp("aa", "aa")
	fmt.Println(compare1, compare2, compare3)
	// Output: -1 1 0
}

func ExampleCmpI() {
	compare1 := wstr.CmpI("aa", "BB")
	compare2 := wstr.CmpI("bb", "AA")
	compare3 := wstr.CmpI("aa", "AA")
	fmt.Println(compare1, compare2, compare3)
	// Output: -1 1 0
}

func ExampleCountRunes() {
	c1 := wstr.CountRunes("foo")
	c2 := wstr.CountRunes("🙂")
	fmt.Println(c1, c2)
	// Output: 3 1
}

func ExampleCountUtf16Len() {
	c1 := wstr.CountUtf16Len("foo")
	c2 := wstr.CountUtf16Len("🙂")
	fmt.Println(c1, c2)
	// Output: 3 2
}

func ExampleFmtBytes() {
	s := wstr.FmtBytes(2 * 1024 * 1024)
	fmt.Println(s)
	// Output: 2.00 MB
}

func ExampleFmtThousands() {
	s := wstr.FmtThousands(2000)
	fmt.Println(s)
	// Output: 2,000
}

func ExampleRemoveDiacritics() {
	s := wstr.RemoveDiacritics("Éçãos")
	fmt.Println(s)
	// Output: Ecaos
}

func ExampleSplitLines() {
	lines := wstr.SplitLines("ab\ncd")
	fmt.Println(lines)
	// Output: [ab cd]
}

func ExampleSubstrRunes() {
	s := wstr.SubstrRunes("ab🙂cd", 1, 3)
	fmt.Println(s)
	// Output: b🙂c
}
