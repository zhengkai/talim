package view

import (
	"project/db"
	"project/pb"
	"project/util"
	"time"

	jp "github.com/buger/jsonparser"
)

func (v *View) TweetList(uid uint64, tid uint64) *pb.TweetList {

	li, _ := db.TweetList(uid, tid)
	if len(li) == 0 {
		return nil
	}

	return v.fillTweetList(li)
}

func (v *View) TweetRecent(tid uint64) *pb.TweetList {

	li, _ := db.TweetRecent(tid)
	if len(li) == 0 {
		return nil
	}

	return v.fillTweetList(li)
}

func (v *View) fillTweetList(li []*db.TweetRow) *pb.TweetList {

	o := &pb.TweetList{}

	for _, v := range li {
		o.Tweet = append(o.Tweet, TweetRow(v))
	}

	// v.TweetFillUser(o)
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
			ul[uid] = v.GetUser(uid)
		}
	}

	for _, r := range t.Tweet {
		check(r.Uid)
		check(r.Reply.GetUid())
	}
	for _, r := range ul {
		if r == nil {
			continue
		}
		t.User = append(t.User, r)
	}
}

func TweetRow(dr *db.TweetRow) *pb.TweetRow {

	pr := &pb.TweetRow{
		Tid: dr.Tid,
		Uid: dr.Uid,
	}

	bin := db.BinLoad(dr.Bid)

	// text
	pr.Text, _ = jp.GetString(bin, `full_text`)

	// ts
	st, err := jp.GetString(bin, `created_at`)
	if err == nil {
		t, err := time.Parse(time.RubyDate, st)
		if err == nil {
			pr.Ts = uint32(t.Unix())
		}
	}

	// media
	util.Media(bin, pr)

	// reply
	replyTID := util.JSONStr2Uint(bin, `in_reply_to_status_id_str`)
	replyUID := util.JSONStr2Uint(bin, `in_reply_to_user_id_str`)
	if replyTID > 0 || replyUID > 0 {
		pr.Reply = &pb.TweetReply{
			Tid: replyTID,
			Uid: replyUID,
		}
	}
	return pr
}
