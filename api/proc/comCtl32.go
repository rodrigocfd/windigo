/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * Copyright 2020-present Rodrigo Cesar de Freitas Dias
 * This library is released under the MIT license
 */

package proc

import (
	"syscall"
)

var (
	dllComCtl32 = syscall.NewLazyDLL("comctl32.dll")

	InitCommonControls = dllComCtl32.NewProc("InitCommonControls")
)
