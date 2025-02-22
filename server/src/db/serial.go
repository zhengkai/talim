package db

import (
	"github.com/google/uuid"
	"github.com/zhengkai/zu"
)

func Serial(u uuid.UUID) uint64 {

	sql := `INSERT INTO uuserial SET uuid = ?, ts_create = ?`
	re, err := d.Exec(sql, u[:], zu.TS())
	if err == nil {
		id, _ := re.LastInsertId()
		return uint64(id)
	}

	sql = `SELECT serial FROM uuserial  WHERE uuid = ?`
	row := d.QueryRow(sql, u[:])
	var id uint64
	row.Scan(&id)
	return id
}
