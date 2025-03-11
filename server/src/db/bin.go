package db

import (
	"crypto/md5"
	"errors"
	"project/zj"

	"github.com/zhengkai/zu"
)

var ErrBinFail = errors.New("save bin fail")

func BinSave(v []byte) uint64 {

	h := md5.Sum(v)

	sql := `INSERT INTO bin SET hash = ?, content = ?, ts_create = ?`
	re, err := d.Exec(sql, h[:], v, zu.TS())
	if err == nil {
		id, _ := re.LastInsertId()
		return uint64(id)
	}

	sql = `SELECT bid FROM bin WHERE hash = ?`
	row := d.QueryRow(sql, h[:])
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		zj.W(err)
		return 0
	}
	return id
}

func BinLoad(bid uint64) (v []byte) {
	sql := `SELECT content FROM bin WHERE bid = ?`
	row := d.QueryRow(sql, bid)
	err := row.Scan(&v)
	if err != nil {
		zj.W(err)
		return nil
	}
	return
}
