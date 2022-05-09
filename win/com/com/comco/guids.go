//go:build windows

package comco

import (
	"github.com/rodrigocfd/windigo/win/co"
)

// IDL COM IIDs.
const (
	IID_IBindCtx          co.IID = "0000000e-0000-0000-c000-000000000046"
	IID_IPersist          co.IID = "0000010c-0000-0000-c000-000000000046"
	IID_IPicture          co.IID = "7bf80980-bf32-101a-8bbb-00aa00300cab"
	IID_ISequentialStream co.IID = "0c733a30-2a1c-11ce-ade5-00aa0044773d"
	IID_IStream           co.IID = "0000000c-0000-0000-c000-000000000046"
	IID_IUnknown          co.IID = "00000000-0000-0000-c000-000000000046"
	IID_NULL              co.IID = "00000000-0000-0000-0000-000000000000"
)
