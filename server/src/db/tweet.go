package db

import (
	"github.com/zhengkai/zu"
)

func TweetSave(tid, uid, bid uint64) {
	sql := `INSERT INTO tweet SET tid = ?, uid = ?, bid = ?, ts_update = ? ON DUPLICATE KEY UPDATE uid = ?, bid = ?, ts_update = ?`
	d.Exec(sql, tid, uid, bid, zu.TS(), uid, bid, zu.TS())

	sql = `INSERT IGNORE INTO tweet_history SET tid = ?, uid = ?, bid = ?, ts_create = ?`
	d.Exec(sql, tid, uid, bid, zu.TS())
}
