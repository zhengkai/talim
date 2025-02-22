package upload

import (
	"project/db"
	"project/zj"

	"github.com/google/uuid"
)

type upload struct {
	uuid   uuid.UUID
	body   []byte
	serial uint64
}

func New(u uuid.UUID, body []byte) {
	upload := upload{
		uuid: u,
		body: body,
	}
	zj.J(`new run`)
	upload.Run()
}

func (u *upload) Serial() uint64 {
	if u.serial == 0 {
		u.serial = db.Serial(u.uuid)
		zj.J(u.serial, u.uuid.String())
	}
	return u.serial
}
