package media

import (
	"encoding/json"
	"project/dl"
	"project/pb"

	jp "github.com/buger/jsonparser"
)

func Media(bin []byte, r *pb.TweetRow) {

	if r == nil {
		r = &pb.TweetRow{}
	}

	jp.ArrayEach(bin, func(v []byte, _ jp.ValueType, _ int, _ error) {
		o := &pb.TwitterMedia{}
		json.Unmarshal(v, o)
		if o.MediaUrlHttps == `` {
			return
		}

		m := &pb.TweetMedia{
			Img: o.MediaUrlHttps,
		}
		dl.Insert(m.Img)
		r.Media = append(r.Media, m)

		vli := o.GetVideoInfo().GetVariants()
		if len(vli) > 0 {
			vo := vli[0]
			for _, vr := range vli[1:] {
				if vr.Bitrate > vo.Bitrate {
					vo = vr
				}
			}
			m.Video = vo.Url
			dl.Insert(m.Video)
			m.ContentType = vo.ContentType
			m.DurMS = uint32(o.GetVideoInfo().GetDurationMillis())
		}
	}, `entities`, `media`)
}
