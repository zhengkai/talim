package view

import (
	"project/pb"

	"github.com/zhengkai/coral/v2"
)

type View struct {
	userCache coral.Cache[uint64, *pb.User]
}

var TheView = &View{
	userCache: coral.NewLRU(userCoral, 5000, 100),
}
