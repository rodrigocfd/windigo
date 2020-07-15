/**
 * Part of Wingows - Win32 API layer for Go
 * https://github.com/rodrigocfd/wingows
 * This library is released under the MIT license.
 */

package com

// IGraphBuilder > IFilterGraph > IUnknown.
type IGraphBuilder struct {
	IFilterGraph
}

type iGraphBuilderVtbl struct {
	iFilterGraphVtbl
	Connect                 uintptr
	Render                  uintptr
	RenderFile              uintptr
	AddSourceFilter         uintptr
	SetLogFile              uintptr
	Abort                   uintptr
	ShouldOperationContinue uintptr
}
