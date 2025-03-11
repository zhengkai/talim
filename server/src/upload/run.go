package upload

import (
	"fmt"
	"project/util"
)

func (u *upload) Run() {

	size := len(u.body)
	if size < 10000 {
		return
	}
	file := fmt.Sprintf(`upload/upload-%d.json`, size)
	if !util.FileExists(file) {
		util.WriteFile(file, u.body)
	}

	u.scanObj(u.body)
}
