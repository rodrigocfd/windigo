package main

import (
	"fmt"
	"time"

	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/dshow"
)

type Picture struct {
	wnd ui.WindowControl

	graphBuilder  dshow.IGraphBuilder
	vmr           dshow.IBaseFilter
	controllerEvr dshow.IMFVideoDisplayControl
	mediaCtrl     dshow.IMediaControl
	mediaSeek     dshow.IMediaSeeking
	basicAudio    dshow.IBasicAudio
}

func NewPicture(parent ui.AnyParent, pos win.POINT, sz win.SIZE) *Picture {
	wnd := ui.NewWindowControlOpts(parent, ui.WindowControlOpts{
		Position: pos,
		Size:     sz,
	})

	me := Picture{
		wnd: wnd,
	}

	me.events()
	return &me
}

func (me *Picture) Free() {
	if me.mediaCtrl.Ppv != nil {
		me.mediaCtrl.Stop()
	}

	me.basicAudio.Release()
	me.mediaSeek.Release()
	me.mediaCtrl.Release()
	me.controllerEvr.Release()
	me.vmr.Release()
	me.graphBuilder.Release()
}

func (me *Picture) events() {
	me.wnd.On().WmPaint(func() {
		ps := win.PAINTSTRUCT{}
		me.wnd.Hwnd().BeginPaint(&ps)
		defer me.wnd.Hwnd().EndPaint(&ps)

		if me.controllerEvr.Ppv != nil {
			me.controllerEvr.RepaintVideo()
		}
	})

	me.wnd.On().WmLButtonDown(func(_ wm.Mouse) {
		// state, _ := me.mediaCtrl.GetState(-1)
		// if state == co.FILTER_STATE_State_Running {
		// 	me.mediaCtrl.Pause()
		// } else {
		// 	me.mediaCtrl.Run()
		// }

		if me.mediaCtrl.Ppv != nil {
			me.mediaSeek.SetPositions(
				5*time.Second, co.SEEKING_FLAGS_RelativePositioning,
				0, co.SEEKING_FLAGS_NoPositioning)
		}
	})
}

func (me *Picture) CurrentTime() string {
	if me.mediaSeek.Ppv == nil {
		return ""
	}

	stNow := win.SYSTEMTIME{}
	stNow.FromDuration(me.mediaSeek.GetCurrentPosition())

	stTot := win.SYSTEMTIME{}
	stTot.FromDuration(me.mediaSeek.GetDuration())

	return fmt.Sprintf("%d:%02d:%02d of %d:%02d:%02d",
		stNow.WHour, stNow.WMinute, stNow.WSecond,
		stTot.WHour, stTot.WMinute, stTot.WSecond)
}

func (me *Picture) StartPlayback(vidPath string) {
	me.Free()

	me.graphBuilder, _ = dshow.CoCreateIGraphBuilder(co.CLSCTX_INPROC_SERVER)
	me.vmr, _ = dshow.CoCreateEnhancedVideoRenderer(co.CLSCTX_INPROC_SERVER)
	me.graphBuilder.AddFilter(&me.vmr, "EVR")

	getSvc, _ := me.vmr.QueryIMFGetService()
	defer getSvc.Release()

	me.controllerEvr, _ = getSvc.GetServiceIMFVideoDisplayControl()
	if e := me.controllerEvr.SetVideoWindow(me.wnd.Hwnd()); e != nil {
		panic(e)
	}
	if e := me.controllerEvr.SetAspectRatioMode(co.MFVideoARMode_PreservePicture); e != nil {
		panic(e)
	}

	me.mediaCtrl, _ = me.graphBuilder.QueryIMediaControl()
	me.mediaSeek, _ = me.graphBuilder.QueryIMediaSeeking()
	me.basicAudio, _ = me.graphBuilder.QueryIBasicAudio()

	if e := me.graphBuilder.RenderFile(vidPath); e != nil {
		panic(e)
	}

	rc := me.wnd.Hwnd().GetWindowRect()
	me.wnd.Hwnd().ScreenToClientRc(&rc)
	me.controllerEvr.SetVideoPosition(nil, &rc)

	me.mediaCtrl.Run()
}