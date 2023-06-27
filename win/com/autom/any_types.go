//go:build windows

package autom

// [MEMBERID] identifiers a member in a type description.
//
// [MEMBERID]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
type MEMBERID int32

// [MEMBERID_NIL] indicates an "unknown" name.
//
// [MEMBERID_NIL]: https://learn.microsoft.com/en-us/previous-versions/windows/desktop/automat/memberid
const MEMBERID_NIL = MEMBERID(-1)
