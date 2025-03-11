package upload

type upload struct {
	body   []byte
	serial uint64
}

func New(body []byte) {
	u := upload{
		body: body,
	}
	u.Run()
}
