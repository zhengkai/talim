package view

import (
	"project/db"
	"project/pb"
)

func (v *View) Index() *pb.UserList {

	li := db.ViewUserCount(v.uuserial)
	if len(li) == 0 {
		return nil
	}

	o := &pb.UserList{}
	for _, du := range li {
		pu := v.GetUser(du.Uid)
		if pu == nil {
			continue
		}
		o.List = append(o.List, &pb.UserRow{
			User:       pu,
			TweetCount: du.TweetCount,
		})
	}
	return o
}
