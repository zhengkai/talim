package db

import (
	"project/zj"

	"github.com/zhengkai/zu"
)

func UserSave(uid uint64, bin []byte) error {

	bid := BinSave(bin)
	if bid == 0 {
		return ErrBinFail
	}

	sql := `INSERT INTO user SET uid = ?, bid = ?, ts_update = ? ON DUPLICATE KEY UPDATE bid = ?, ts_update = ?`
	_, err := d.Exec(sql, uid, bid, zu.TS(), bid, zu.TS())
	if err != nil {
		zj.W(`UserSave`, err)
		return err
	}

	sql = `INSERT IGNORE INTO user_history SET uid = ?, bid = ?, ts_create = ?`
	_, err = d.Exec(sql, uid, bid, zu.TS())
	if err != nil {
		zj.W(`UserSave`, err)
		return err
	}
	return nil
}

func UserLoad(uid uint64) []byte {

	sql := `SELECT bid FROM user WHERE uid = ?`
	row := d.QueryRow(sql, uid)
	var bid uint64
	err := row.Scan(&bid)
	if err != nil {
		zj.W(`UserLoad`, err)
	}
	if bid == 0 {
		return nil
	}

	return BinLoad(bid)
}
