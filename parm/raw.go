package parm

import (
	"winffi/api"
	c "winffi/consts"
)

type Raw struct {
	Msg    c.WM
	WParam api.WPARAM
	LParam api.LPARAM
}
