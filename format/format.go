package format

import (
	"github.com/pickledgator/vdk/av/avutil"
	"github.com/pickledgator/vdk/format/aac"
	"github.com/pickledgator/vdk/format/flv"
	"github.com/pickledgator/vdk/format/mp4"
	"github.com/pickledgator/vdk/format/rtmp"
	"github.com/pickledgator/vdk/format/rtsp"
	"github.com/pickledgator/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
