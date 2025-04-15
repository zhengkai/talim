package dl

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
	"project/util"
	"regexp"
	"strings"
)

var pExt = regexp.MustCompile(`[a-z0-9]{1,10}`)

func download(url string) {

	hash := md5.Sum([]byte(url))

	s := strings.Split(url, `?`)[0]
	ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(s), `.`))
	if ext == `.mp4` {
		return
	}
	if !pExt.MatchString(ext) {
		return
	}

	file := fmt.Sprintf(`file/%x/%x/%x.%s`, hash[:1], hash[1:2], hash[2:], ext)

	util.Download(url, file)
}
