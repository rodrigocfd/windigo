//go:build windows

// This package provides support to convert between Go strings and native,
// null-terminated Windows UTF-16 strings. It's mainly used within the library,
// but it's available if you need this kind of encoding/decoding.
package wstr
