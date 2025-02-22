package upload

import "github.com/google/uuid"

type upload struct {
	uuid uuid.UUID
	body []byte
}

func New(u uuid.UUID, body []byte) {
	upload := upload{
		uuid: u,
		body: body,
	}
	upload.Run()
}
