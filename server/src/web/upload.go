package web

import (
	"io"
	"net/http"
	"project/upload"
	"project/util"
	"strconv"

	"github.com/google/uuid"
)

var okJSON = []byte(`{"ok":true}`)

func uploadHandle(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(`Content-Type`, `application/json`)
	w.Header().Set(`Content-Length`, strconv.Itoa(len(okJSON)))
	w.Write(okJSON)
	if r.Method != http.MethodPost {
		return
	}

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

	go upload.New(u, body)
}

func corsWrite(w http.ResponseWriter, r *http.Request) (stop bool) {

	// 写了一顿才知道 GM_xmlhttpRequest 不受 cors 限制，用不上了

	w.Header().Set(`Access-Control-Allow-Origin`, `https://x.com`)
	w.Header().Set(`Access-Control-Allow-Methods`, `GET, POST, OPTIONS`)
	w.Header().Set(`Access-Control-Allow-Headers`, `Content-Type`)

	if r.Method == http.MethodOptions {
		w.Header().Set(`Access-Control-Max-Age`, `864000`)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.Header().Set(`Access-Control-Expose-Headers`, `Content-Length`)
		w.Header().Set(`Content-Type`, `application/json`)
		w.Write(okJSON)
	}

	return r.Method != http.MethodPost // 只有 POST 会继续
}
