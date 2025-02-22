package web

import (
	"io"
	"net/http"
	"project/upload"
	"project/util"

	"github.com/google/uuid"
)

var ok = []byte(`{"ok":true}`)

func uploadHandle(w http.ResponseWriter, r *http.Request) {

	w.Write(ok)

	body, err := io.ReadAll(io.LimitReader(r.Body, 1024*1024))
	if err != nil || len(body) < 1000 {
		return
	}
	if body[0] != '{' || body[len(body)-1] != '}' {
		// not json
		return
	}

	u, err := uuid.Parse(r.URL.Query().Get(`uuid`))
	if err != nil {
		u = util.DefaultUUID
	}

	upload.New(u, body)
}
