package view

import (
	"errors"
	"project/db"
	"time"

	jp "github.com/buger/jsonparser"
)

type serialUser struct {
	uuserial uint64
	uid      uint64
}

func userNameCoral(su serialUser) (string, *time.Time, error) {
	bin := db.UserLoad(su.uuserial, su.uid)

	if bin == nil {
		// zj.J(`load user fail`, su.uid)
		return ``, nil, errors.New(`load user fail`)
	}

	name, _ := jp.GetString(bin, `screen_name`)
	return name, nil, nil
}

func (v *View) UserName(uid uint64) string {
	q := serialUser{
		uuserial: v.uuserial,
		uid:      uid,
	}
	name, _ := v.nameCache.Get(q)
	return name
}
