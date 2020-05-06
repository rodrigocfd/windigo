package parm

import (
	"gowinui/api"
	c "gowinui/consts"
)

type Raw struct {
	Msg    c.WM
	WParam api.WPARAM
	LParam api.LPARAM
}
