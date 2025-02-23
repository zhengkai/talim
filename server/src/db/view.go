package db

import (
	"database/sql"
	"project/pb"
	"project/zj"
)

type TweetRow struct {
	Tid uint64
	Uid uint64
	Bid uint64
}

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

func TweetRecent(serial, tid uint64) ([]*TweetRow, error) {

	var rs *sql.Rows
	var err error

	if tid == 0 {
		sql := "SELECT tid, uid, bid FROM tweet WHERE serial = ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, serial)
	} else {
		sql := "SELECT tid, uid, bid FROM tweet WHERE serial = ? AND tid < ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, serial, tid)
	}

	if err != nil {
		zj.W(err)
		return nil, err
	}
	defer rs.Close()

	return listTweet(rs)
}

func TweetList(serial, uid, tid uint64) ([]*TweetRow, error) {

	var rs *sql.Rows
	var err error

	if tid == 0 {
		sql := "SELECT tid, uid, bid FROM tweet WHERE serial = ? AND uid = ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, serial, uid)
	} else {
		sql := "SELECT tid, uid, bid FROM tweet WHERE serial = ? AND uid = ? AND tid < ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, serial, uid, tid)
	}

	if err != nil {
		zj.W(err)
		return nil, err
	}
	defer rs.Close()

	return listTweet(rs)
}

func listTweet(rs *sql.Rows) (li []*TweetRow, err error) {
	for rs.Next() {
		r := &TweetRow{}
		err := rs.Scan(&r.Tid, &r.Uid, &r.Bid)
		if err != nil {
			continue
		}
		li = append(li, r)
	}
	return
}
