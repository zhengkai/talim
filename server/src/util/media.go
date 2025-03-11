package util

import (
	"encoding/json"
	"project/pb"

	jp "github.com/buger/jsonparser"
)

func Media(bin []byte, r *pb.TweetRow) {

	jp.ArrayEach(bin, func(v []byte, _ jp.ValueType, _ int, _ error) {
		o := &pb.TwitterMedia{}
		json.Unmarshal(v, o)
		if o.MediaUrlHttps == `` {
			return
		}

		m := &pb.TweetMedia{
			Img: o.MediaUrlHttps,
		}
		r.Media = append(r.Media, m)

		vli := o.GetVideoInfo().GetVariants()
		if len(vli) > 0 {
			vo := vli[0]
			for _, vr := range vli {
				if vr.Bitrate > vo.Bitrate {
					vo = vr
				}
			}
			m.Video = vo.Url
			m.ContentType = vo.ContentType
			m.DurMS = uint32(o.GetVideoInfo().GetDurationMillis())
		}
	}, `entities`, `media`)
}
