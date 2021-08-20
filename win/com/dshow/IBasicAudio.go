package dshow

import (
	"syscall"
	"unsafe"

	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/errco"
)

type _IBasicAudioVtbl struct {
	win.IDispatchVtbl
	PutVolume  uintptr
	GetVolume  uintptr
	PutBalance uintptr
	GetBalance uintptr
}

//------------------------------------------------------------------------------

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nn-control-ibasicaudio
type IBasicAudio struct {
	win.IDispatch // Base IDispatch > IUnknown.
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-get_balance
func (me *IBasicAudio) GetBalance() int {
	balance := int32(0)
	ret, _, _ := syscall.Syscall(
		(*_IBasicAudioVtbl)(unsafe.Pointer(*me.Ppv)).GetBalance, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&balance)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(balance)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-get_volume
func (me *IBasicAudio) GetVolume() int {
	volume := int32(0)
	ret, _, _ := syscall.Syscall(
		(*_IBasicAudioVtbl)(unsafe.Pointer(*me.Ppv)).GetVolume, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(unsafe.Pointer(&volume)), 0)

	if hr := errco.ERROR(ret); hr == errco.S_OK {
		return int(volume)
	} else {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-put_balance
func (me *IBasicAudio) PutBalance(balance int) {
	ret, _, _ := syscall.Syscall(
		(*_IBasicAudioVtbl)(unsafe.Pointer(*me.Ppv)).PutBalance, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(int32(balance)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}

// 📑 https://docs.microsoft.com/en-us/windows/win32/api/control/nf-control-ibasicaudio-put_volume
func (me *IBasicAudio) PutVolume(volume int) {
	ret, _, _ := syscall.Syscall(
		(*_IBasicAudioVtbl)(unsafe.Pointer(*me.Ppv)).PutVolume, 2,
		uintptr(unsafe.Pointer(me.Ppv)),
		uintptr(int32(volume)), 0)

	if hr := errco.ERROR(ret); hr != errco.S_OK {
		panic(hr)
	}
}
