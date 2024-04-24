package main

import (
	"fmt"
	"time"

	"github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/ui/wm"
	"github.com/rodrigocfd/windigo/win"
	"github.com/rodrigocfd/windigo/win/co"
	"github.com/rodrigocfd/windigo/win/com/com"
	"github.com/rodrigocfd/windigo/win/com/com/comco"
	"github.com/rodrigocfd/windigo/win/com/dshow"
	"github.com/rodrigocfd/windigo/win/com/dshow/dshowco"
)

// Child window which renders the video.
type Picture struct {
	wnd ui.WindowControl

	graphBuilder  dshow.IGraphBuilder
	vmr           dshow.IBaseFilter
	controllerEvr dshow.IMFVideoDisplayControl
	mediaCtrl     dshow.IMediaControl
	mediaSeek     dshow.IMediaSeeking
	basicAudio    dshow.IBasicAudio
}

func NewPicture(
	parent ui.AnyParent, pos win.POINT, sz win.SIZE,
	horz ui.HORZ, vert ui.VERT) *Picture {

	wnd := ui.NewWindowControl(
		parent,
		ui.WindowControlOpts().
			WndExStyles(co.WS_EX_NONE).
			Position(pos).
			Size(sz).
			Horz(horz).
			Vert(vert),
	)

	me := &Picture{
		wnd: wnd,
	}

	me.events()
	return me
}

func (me *Picture) FreeComObjs() {
	if com.IsObj(me.mediaCtrl) {
		me.mediaCtrl.Stop()

		me.basicAudio.Release()
		me.mediaSeek.Release()
		me.mediaCtrl.Release()
		me.controllerEvr.Release()
		me.vmr.Release()
		me.graphBuilder.Release()
	}
}

func (me *Picture) StartPlayback(vidPath string) {
	me.FreeComObjs()

	me.graphBuilder = dshow.NewIGraphBuilder(
		com.CoCreateInstance(
			dshowco.CLSID_FilterGraph, nil,
			comco.CLSCTX_INPROC_SERVER,
			dshowco.IID_IGraphBuilder),
	)
	me.vmr = dshow.NewIBaseFilter(
		com.CoCreateInstance(
			dshowco.CLSID_EnhancedVideoRenderer, nil,
			comco.CLSCTX_INPROC_SERVER,
			dshowco.IID_IBaseFilter),
	)
	me.graphBuilder.AddFilter(me.vmr, "EVR")

	getSvc := dshow.NewIMFGetService(
		me.vmr.QueryInterface(dshowco.IID_IMFGetService),
	)
	defer getSvc.Release()

	me.controllerEvr = dshow.NewIMFVideoDisplayControl(
		getSvc.GetService(
			win.GuidFromClsid(dshowco.CLSID_MR_VideoRenderService),
			win.GuidFromIid(dshowco.IID_IMFVideoDisplayControl),
		),
	)

	if err := me.controllerEvr.SetVideoWindow(me.wnd.Hwnd()); err != nil {
		panic(err)
	}
	if err := me.controllerEvr.SetAspectRatioMode(dshowco.MFVideoARMode_PreservePicture); err != nil {
		panic(err)
	}

	me.mediaCtrl = dshow.NewIMediaControl(
		me.graphBuilder.QueryInterface(dshowco.IID_IMediaControl),
	)
	me.mediaSeek = dshow.NewIMediaSeeking(
		me.graphBuilder.QueryInterface(dshowco.IID_IMediaSeeking),
	)
	me.basicAudio = dshow.NewIBasicAudio(
		me.graphBuilder.QueryInterface(dshowco.IID_IBasicAudio),
	)

	if err := me.graphBuilder.RenderFile(vidPath); err != nil {
		panic(err)
	}

	rc := me.wnd.Hwnd().GetWindowRect()
	me.wnd.Hwnd().ScreenToClientRc(&rc)
	me.controllerEvr.SetVideoPosition(nil, &rc)

	me.mediaCtrl.Run()
}

func (me *Picture) Pause() {
	if com.IsObj(me.mediaCtrl) {
		state, _ := me.mediaCtrl.GetState(win.NumInfInfinite())
		if state == dshowco.FILTER_STATE_State_Running {
			me.mediaCtrl.Pause()
		}
	}
}

func (me *Picture) TogglePlayPause() {
	if com.IsObj(me.mediaCtrl) {
		state, _ := me.mediaCtrl.GetState(win.NumInfInfinite())
		if state == dshowco.FILTER_STATE_State_Running {
			me.mediaCtrl.Pause()
		} else {
			me.mediaCtrl.Run()
		}
	}
}

func (me *Picture) Duration() (secs int) {
	if !com.IsObj(me.mediaSeek) {
		return 0
	} else {
		return int(me.mediaSeek.GetDuration() / time.Second)
	}
}

func (me *Picture) SetCurrentPos(secs int) {
	if com.IsObj(me.mediaSeek) {
		me.mediaSeek.SetPositions(
			time.Duration(secs)*time.Second, dshowco.SEEKING_FLAGS_AbsolutePositioning,
			0, dshowco.SEEKING_FLAGS_NoPositioning)
	}
}

func (me *Picture) CurrentPos() (secs int) {
	if !com.IsObj(me.mediaSeek) {
		return 0
	} else {
		return int(me.mediaSeek.GetCurrentPosition() / time.Second)
	}
}

func (me *Picture) CurrentPosDurFmt() string {
	if !com.IsObj(me.mediaSeek) {
		return "NO VIDEO"

	} else {
		stCurPos := win.SYSTEMTIME{}
		stCurPos.FromDuration(me.mediaSeek.GetCurrentPosition())

		stDur := win.SYSTEMTIME{}
		stDur.FromDuration(me.mediaSeek.GetDuration())

		return fmt.Sprintf("%d:%02d:%02d of %d:%02d:%02d",
			stCurPos.WHour, stCurPos.WMinute, stCurPos.WSecond,
			stDur.WHour, stDur.WMinute, stDur.WSecond)
	}
}

func (me *Picture) ForwardSecs(secs int) {
	if com.IsObj(me.mediaSeek) {
		newSecs := me.CurrentPos() + secs
		duration := me.Duration()
		if newSecs >= duration {
			newSecs = duration - 1 // max pos
		}
		me.SetCurrentPos(newSecs)
	}
}

func (me *Picture) BackwardSecs(secs int) {
	if com.IsObj(me.mediaSeek) {
		newSecs := me.CurrentPos() - secs
		if newSecs < 0 {
			newSecs = 0 // min pos
		}
		me.SetCurrentPos(newSecs)
	}
}

func (me *Picture) events() {
	me.wnd.On().WmPaint(func() {
		ps := win.PAINTSTRUCT{}
		me.wnd.Hwnd().BeginPaint(&ps)
		defer me.wnd.Hwnd().EndPaint(&ps)

		if com.IsObj(me.controllerEvr) {
			me.controllerEvr.RepaintVideo()
		}
	})

	me.wnd.On().WmSize(func(p wm.Size) {
		if com.IsObj(me.controllerEvr) {
			rc := me.wnd.Hwnd().GetWindowRect()
			me.wnd.Hwnd().ScreenToClientRc(&rc)
			me.controllerEvr.SetVideoPosition(nil, &rc)
		}
	})

	me.wnd.On().WmLButtonDown(func(_ wm.Mouse) {
		me.wnd.Hwnd().SetFocus()
		me.TogglePlayPause()
	})

	me.wnd.On().WmKeyDown(func(p wm.Key) {
		if p.VirtualKeyCode() == co.VK_SPACE {
			me.TogglePlayPause()
		}
	})

	me.wnd.On().WmGetDlgCode(func(p wm.GetDlgCode) co.DLGC {
		if p.VirtualKeyCode() == co.VK_LEFT {
			me.BackwardSecs(10)
			return co.DLGC_WANTARROWS
		} else if p.VirtualKeyCode() == co.VK_RIGHT {
			me.ForwardSecs(10)
			return co.DLGC_WANTARROWS
		}
		return co.DLGC_NONE
	})
}
