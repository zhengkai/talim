package upload

import (
	"encoding/json"
	"project/db"
	"project/util"
	"project/zj"
	"strconv"
	"strings"

	jp "github.com/buger/jsonparser"
)

type user struct {
	uid uint64
	bin []byte
}

type tweet struct {
	tid uint64
	uid uint64
	bin []byte
}

var cleanUserKey = []string{
	`can_dm`,
	`following`,
	`possibly_sensitive`,
	`is_translator`,
	`has_custom_timelines`,
	`can_media_tag`,
}
var cleanTweetKey = []string{
	`possibly_sensitive`,
	`possibly_sensitive_editable`,
	`display_text_range`,
	`bookmarked`,
	`favorited`,
	`retweeted`,
	`is_quote_status`,
}

func (u *upload) importUser(ab []byte, uid uint64) *user {
	var o map[string]any
	err := json.Unmarshal(ab, &o)
	if err != nil {
		util.WriteFile(`debug/user-unmarshal-fail.json`, ab)
		return nil
	}
	if len(o) < 10 {
		return nil
	}
	for _, k := range cleanUserKey {
		delete(o, k)
	}
	for k := range o {
		if strings.HasSuffix(k, `_count`) {
			delete(o, k)
			continue
		}
	}
	if uid == 0 {
		uid = getUIDFromBanner(ab)
		if uid == 0 {
			util.WriteFile(`debug/user-no-uid.json`, ab)
			return nil
		}
	}
	re, err := json.Marshal(o)
	if err != nil {
		zj.W(err)
		return nil
	}

	db.UserSave(u.Serial(), uid, re)

	return &user{
		uid: uid,
		bin: re,
	}
}

func getUIDFromBanner(ab []byte) uint64 {

	s, err := jp.GetString(ab, `profile_banner_url`)
	if err != nil {
		return 0
	}

	if !strings.HasPrefix(s, `https://pbs.twimg.com/profile_banners/`) {
		return 0
	}
	li := strings.Split(s, `/`)
	i, err := strconv.ParseUint(li[4], 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func (u *upload) importTweet(ab []byte) *tweet {

	uid := util.JSONStr2Uint(ab, `user_id_str`)
	if uid == 0 {
		util.WriteFile(`debug/tweet-no-uid.json`, ab)
		return nil
	}
	tid := util.JSONStr2Uint(ab, `id_str`)
	if tid == 0 {
		util.WriteFile(`debug/tweet-no-tid.json`, ab)
		return nil
	}

	var o map[string]any
	err := json.Unmarshal(ab, &o)
	if err != nil {
		util.WriteFile(`debug/tweet-unmarshal-fail.json`, ab)
		return nil
	}
	if len(o) < 10 {
		return nil
	}
	for _, k := range cleanTweetKey {
		delete(o, k)
	}
	for k := range o {
		if strings.HasSuffix(k, `_count`) {
			delete(o, k)
			continue
		}
	}

	re, err := json.Marshal(o)
	if err != nil {
		return nil
	}

	bid := db.BinSave(re)
	if bid == 0 {
		return nil
	}

	db.TweetSave(u.Serial(), tid, uid, bid)

	return &tweet{
		uid: uid,
		tid: tid,
		bin: re,
	}
}
