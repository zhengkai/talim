package util

import (
	"strconv"

	jp "github.com/buger/jsonparser"
)

func JSONStr2Uint(ab []byte, k string) uint64 {
	s, err := jp.GetString(ab, `id_str`)
	if err != nil {
		return 0
	}

	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
