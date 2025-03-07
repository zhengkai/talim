package view

import (
	"errors"
	"project/db"
	"project/pb"
	"time"

	"github.com/google/uuid"
	"github.com/zhengkai/coral/v2"
)

type View struct {
	uuserial  uint64
	userCache coral.Cache[serialUser, *pb.User]
}

var viewCache = coral.NewLRU(viewCoral, 1000, 100)

func New(u uuid.UUID) *View {
	v, _ := viewCache.Get(u)
	return v
}

func viewCoral(u uuid.UUID) (*View, *time.Time, error) {

	uuserial := db.SerialLoad(u)
	if uuserial == 0 {
		return nil, nil, errors.New(`load serial fail`)
	}

	v := &View{
		uuserial:  uuserial,
		userCache: coral.NewLRU(userCoral, 5000, 100),
	}
	return v, nil, nil
}
