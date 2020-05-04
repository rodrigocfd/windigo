package api

import (
	"log"
	"syscall"
)

func toUtf16Ptr(s string) *uint16 {
	// We won't return an uintptr right away because it has no pointer semantics,
	// it's just a number, so pointed memory can be garbage-collected.
	// https://stackoverflow.com/a/51188315
	pstr, err := syscall.UTF16PtrFromString(s)
	if err != nil {
		log.Panicf("toUtf16Ptr failed \"%s\": %s\n", s, err)
	}
	return pstr
}

func toUtf16PtrBlankIsNil(s string) *uint16 {
	if s != "" {
		return toUtf16Ptr(s)
	}
	return nil
}

func boolToUintptr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}
