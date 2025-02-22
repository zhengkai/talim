package view

import (
	"encoding/json"
	"project/db"
	"project/pb"
	"project/util"
	"project/zj"
	"time"

	jp "github.com/buger/jsonparser"
)

func (v *View) TweetList(uid uint64) *pb.TweetList {

	o := db.ViewUserTweet(v.uuserial, uid)
	if len(o.GetTweet()) == 0 {
		return nil
	}

	for _, v := range o.Tweet {
		TweetFill(v, v.Uid)
		v.Uid = uid
	}

	v.TweetFillUser(o)

	return o
}

func (v *View) TweetFillUser(t *pb.TweetList) {
	ul := make(map[uint64]*pb.User)

	check := func(uid uint64) {
		if uid == 0 {
			return
		}
		_, ok := ul[uid]
		if !ok {
			ul[uid] = &pb.User{
				Uid:        uid,
				ScreenName: v.UserName(uid),
			}
		}
	}

	for _, r := range t.Tweet {
		check(r.Uid)
		check(r.Reply.GetUid())
	}
	for _, r := range ul {
		if r.ScreenName == `` {
			continue
		}
		t.User = append(t.User, r)
	}
}

func TweetFill(r *pb.TweetRow, bid uint64) {

	bin := db.BinLoad(bid)

	// text
	r.Text, _ = jp.GetString(bin, `full_text`)

	// ts
	st, err := jp.GetString(bin, `created_at`)
	if err == nil {
		t, err := time.Parse(time.RubyDate, st)
		if err == nil {
			r.Ts = uint32(t.Unix())
		}
	}

	// media
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

		zj.J(r.Tid)
		util.DumpJSON(o)
	}, `entities`, `media`)

	// reply
	replyTID := util.JSONStr2Uint(bin, `in_reply_to_status_id_str`)
	replyUID := util.JSONStr2Uint(bin, `in_reply_to_user_id_str`)
	if replyTID > 0 || replyUID > 0 {
		r.Reply = &pb.TweetReply{
			Tid: replyTID,
			Uid: replyUID,
		}
	}
}
