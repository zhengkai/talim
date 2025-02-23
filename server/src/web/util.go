package web

import (
	"net/http"
	"strconv"
)

func getNum(r *http.Request, key string) uint64 {
	s := r.URL.Query().Get(key)
	if s == `` {
		return 0
	}
	n, _ := strconv.ParseUint(s, 10, 64)
	return n
}
