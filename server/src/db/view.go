package db

import (
	"database/sql"
	"project/zj"
)

type TweetRow struct {
	Tid uint64
	Uid uint64
	Bid uint64
}

type UserRow struct {
	Uid        uint64
	TweetCount uint64
}

func ViewUserCount() []*UserRow {

	sql := "SELECT uid, COUNT(uid) as cnt FROM tweet GROUP BY uid"
	rs, err := d.Query(sql)
	if err != nil {
		return nil
	}

	var li []*UserRow
	defer rs.Close()
	for rs.Next() {
		ur := &UserRow{}
		err := rs.Scan(&ur.Uid, &ur.TweetCount)
		if err != nil {
			continue
		}
		li = append(li, ur)
	}
	return li
}

func TweetRecent(tid uint64) ([]*TweetRow, error) {

	var rs *sql.Rows
	var err error

	if tid == 0 {
		sql := "SELECT tid, uid, bid FROM tweet ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql)
	} else {
		sql := "SELECT tid, uid, bid FROM tweet WHERE tid < ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, tid)
	}

	if err != nil {
		zj.W(err)
		return nil, err
	}
	defer rs.Close()

	return listTweet(rs)
}

func TweetList(uid, tid uint64) ([]*TweetRow, error) {

	var rs *sql.Rows
	var err error

	if tid == 0 {
		sql := "SELECT tid, uid, bid FROM tweet WHERE uid = ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, uid)
	} else {
		sql := "SELECT tid, uid, bid FROM tweet WHERE uid = ? AND tid < ? ORDER BY tid DESC LIMIT 5000"
		rs, err = d.Query(sql, uid, tid)
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
