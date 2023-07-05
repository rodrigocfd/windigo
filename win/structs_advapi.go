//go:build windows

package win

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// [TOKEN_ELEVATION] struct.
//
// [TOKEN_ELEVATION]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-token_elevation
type TOKEN_ELEVATION struct {
	tokenIsElevated uint32
}

func (te *TOKEN_ELEVATION) TokenIsElevated() bool { return te.tokenIsElevated != 0 }

// [TOKEN_LINKED_TOKEN] struct.
//
// [TOKEN_LINKED_TOKEN]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-token_linked_token
type TOKEN_LINKED_TOKEN struct {
	LinkedToken HACCESSTOKEN
}

// [TOKEN_MANDATORY_POLICY] struct.
//
// [TOKEN_MANDATORY_POLICY]: https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-token_mandatory_policy
type TOKEN_MANDATORY_POLICY struct {
	Policy co.TOKEN_POLICY
}
