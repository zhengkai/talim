package db

import (
	"github.com/zhengkai/zu"
)

func UserSave(serial, uid, bid uint64) {
	sql := `INSERT INTO user SET serial = ?, uid = ?, bid = ?, ts_update = ? ON DUPLICATE KEY UPDATE bid = ?, ts_update = ?`
	d.Exec(sql, serial, uid, bid, zu.TS(), bid, zu.TS())

	sql = `INSERT INTO user_history SET serial = ?, uid = ?, bid = ?, ts_create = ?`
	d.Exec(sql, serial, uid, bid, zu.TS())
}
