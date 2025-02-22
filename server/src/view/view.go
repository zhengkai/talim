package view

import (
	"errors"
	"project/db"
	"time"

	"github.com/google/uuid"
	"github.com/zhengkai/coral/v2"
)

type View struct {
	uuserial  uint64
	nameCache coral.Cache[serialUser, string]
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
		nameCache: coral.NewLRU(userNameCoral, 5000, 100),
	}
	return v, nil, nil
}
