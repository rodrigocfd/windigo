//go:build windows

package wstr_test

import (
	"fmt"

	"github.com/rodrigocfd/windigo/wstr"
)

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
