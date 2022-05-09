//go:build windows

package autom

// Identifiers a member in a type description.
//
// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
type MEMBERID int32

// Indicates an "unknown" name.
//
// 📑 https://docs.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
const MEMBERID_NIL = MEMBERID(-1)
