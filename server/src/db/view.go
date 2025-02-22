package db

import "project/pb"

func ViewUserCount(serial uint64) *pb.UserList {

	sql := "SELECT uid, COUNT(uid) as cnt FROM tweet WHERE serial = ? GROUP BY uid"
	rs, err := d.Query(sql, serial)
	if err != nil {
		panic(err)
	}

	o := &pb.UserList{}
	defer rs.Close()
	for rs.Next() {
		var uid uint64
		var cnt uint64
		err := rs.Scan(&uid, &cnt)
		if err != nil {
			continue
		}
		o.List = append(o.List, &pb.UserRow{
			Uid:        uid,
			TweetCount: cnt,
		})
	}
	return o
}
