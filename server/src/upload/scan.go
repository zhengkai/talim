package upload

import (
	"fmt"
	"project/util"

	jp "github.com/buger/jsonparser"
)

const keyTypeUser = `User`
const keyTypeTweet = `Tweet`

func (u *upload) check(ab []byte) {

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
		u.importTweet(legacy)
	case `User`:
		uid := util.JSONStr2Uint(ab, `rest_id`)
		u.importUser(legacy, uid)
	}
}

func (u *upload) scanObj(ab []byte) {

	u.check(ab)

	jp.ObjectEach(ab, func(k []byte, v []byte, t jp.ValueType, _ int) error {
		switch t {
		case jp.Object:
			u.scanObj(v)
		case jp.Array:
			u.scanArray(v)
		}
		return nil
	})
}

func (u *upload) scanArray(ab []byte) {
	jp.ArrayEach(ab, func(v []byte, t jp.ValueType, _ int, _ error) {
		if t == jp.Object {
			u.scanObj(v)
		}
	})
}
