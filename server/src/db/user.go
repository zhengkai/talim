package db

import (
	"github.com/zhengkai/zu"
)

func UserSave(serial, uid uint64, bin []byte) error {

	bid := BinSave(bin)
	if bid == 0 {
		return ErrBinFail
	}

	sql := `INSERT INTO user SET serial = ?, uid = ?, bid = ?, ts_update = ? ON DUPLICATE KEY UPDATE bid = ?, ts_update = ?`
	_, err := d.Exec(sql, serial, uid, bid, zu.TS(), bid, zu.TS())
	if err != nil {
		return err
	}

	sql = `INSERT IGNORE INTO user_history SET serial = ?, uid = ?, bid = ?, ts_create = ?`
	d.Exec(sql, serial, uid, bid, zu.TS())
	return nil
}

func UserLoad(serial, uid uint64) []byte {

	sql := `SELECT bid FROM user WHERE serial = ? AND uid = ?`
	row := d.QueryRow(sql, serial, uid)
	var bid uint64
	row.Scan(&bid)
	if bid == 0 {
		return nil
	}

	return BinLoad(bid)
}
