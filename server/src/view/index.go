package view

import (
	"project/db"
	"project/pb"
)

func (v *View) Index() *pb.UserList {

	o := db.ViewUserCount(v.uuserial)
	if o == nil {
		return nil
	}

	for _, u := range o.List {
		u.ScreenName = v.UserName(u.Uid)
	}
	return o
}
