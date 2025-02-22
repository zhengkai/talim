package upload

import (
	"fmt"
	"project/util"

	jp "github.com/buger/jsonparser"
)

const keyTypeUser = `User`
const keyTypeTweet = `Tweet`

type scan struct {
}

func startScan(ab []byte) {
	s := scan{}
	s.scanObj(ab)
}

func (s *scan) check(ab []byte) {
	t, _ := jp.GetString(ab, `__typename`)
	if t == `` {
		return
	}
	file := fmt.Sprintf(`demo/%s.json`, t)
	if !util.FileExists(file) {
		util.WriteFile(file, ab)
	}

	if t != keyTypeUser && t != keyTypeTweet {
		return
	}

	legacy, vt, _, err := jp.Get(ab, `legacy`)
	if err != nil || vt != jp.Object {
		if t == keyTypeTweet && len(ab) < 100 {
			return
		}
		util.WriteFile(`debug/no-legacy.json`, ab)
		return
	}

	switch t {
	case `Tweet`:
		importTweet(legacy)
	case `User`:
		importUser(legacy)
	}
}

func (s *scan) scanObj(ab []byte) {

	s.check(ab)

	jp.ObjectEach(ab, func(k []byte, v []byte, t jp.ValueType, _ int) error {
		switch t {
		case jp.Object:
			s.scanObj(v)
		case jp.Array:
			s.scanArray(v)
		}
		return nil
	})
}

func (s *scan) scanArray(ab []byte) {
	jp.ArrayEach(ab, func(v []byte, t jp.ValueType, _ int, _ error) {
		if t == jp.Object {
			s.scanObj(v)
		}
	})
}
