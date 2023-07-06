//go:build windows

package d2d1co

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// Direct2D COM IIDs.
const (
	IID_ID2D1Factory          co.IID = "06152247-6f50-465a-9245-118bfd3b6007"
	IID_ID2D1HwndRenderTarget co.IID = "2cd90698-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1Layer            co.IID = "2cd9069b-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1Mesh             co.IID = "2cd906c2-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1RenderTarget     co.IID = "2cd90694-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1Resource         co.IID = "2cd90691-12e2-11dc-9fed-001143a055f9"
	IID_ID2D1TessellationSink co.IID = "2cd906c1-12e2-11dc-9fed-001143a055f9"
)
