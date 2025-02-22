package db

import (
	"project/pb"
	"project/zj"
)

func ViewUserCount(serial uint64) *pb.UserList {

	sql := "SELECT uid, COUNT(uid) as cnt FROM tweet WHERE serial = ? GROUP BY uid"
	rs, err := d.Query(sql, serial)
	if err != nil {
		panic(err)
	}

	o := &pb.UserList{}
	defer rs.Close()
	for rs.Next() {
		var uid uint64
		var cnt uint64
		err := rs.Scan(&uid, &cnt)
		if err != nil {
			continue
		}
		o.List = append(o.List, &pb.UserRow{
			Uid:        uid,
			TweetCount: cnt,
		})
	}
	return o
}

func ViewUserTweet(serial, uid uint64) *pb.TweetList {

	sql := "SELECT tid, bid FROM tweet WHERE serial = ? AND uid = ? ORDER BY tid DESC LIMIT 5000"
	rs, err := d.Query(sql, serial, uid)
	if err != nil {
		zj.W(err)
		return nil
	}
	defer rs.Close()

	o := &pb.TweetList{}
	for rs.Next() {
		r := &pb.TweetRow{}
		err := rs.Scan(&r.Tid, &r.Uid)
		if err != nil {
			continue
		}
		o.Tweet = append(o.Tweet, r)
	}
	return o
}
