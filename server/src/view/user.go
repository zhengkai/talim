package view

import (
	"errors"
	"project/db"
	"project/pb"
	"time"

	jp "github.com/buger/jsonparser"
)

type serialUser struct {
	uuserial uint64
	uid      uint64
}

func userCoral(uid uint64) (*pb.User, *time.Time, error) {

	bin := db.UserLoad(uid)

	if len(bin) < 10 {
		// zj.J(`load user fail`, su.uid)
		return nil, nil, errors.New(`load user fail`)
	}

	u := &pb.User{
		Uid: uid,
	}
	u.Name, _ = jp.GetString(bin, `name`)
	u.ScreenName, _ = jp.GetString(bin, `screen_name`)
	u.Avatar, _ = jp.GetString(bin, `profile_image_url_https`)
	u.Banner, _ = jp.GetString(bin, `profile_banner_url`)
	u.Bio, _ = jp.GetString(bin, `description`)
	u.Location, _ = jp.GetString(bin, `location`)
	st, err := jp.GetString(bin, `created_at`)
	if err == nil {
		t, err := time.Parse(time.RubyDate, st)
		if err == nil {
			u.TsCreate = uint32(t.Unix())
		}
	}

	return u, nil, nil
}

func (v *View) GetUser(uid uint64) *pb.User {
	u, _ := v.userCache.Get(uid)
	return u
}
