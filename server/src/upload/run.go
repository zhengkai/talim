package upload

import (
	"fmt"
	"project/util"
	"project/zj"
)

func (u *upload) Run() {
	zj.J(`run`, u.uuid.String(), len(u.body))

	size := len(u.body)
	if size > 10000 {
		file := fmt.Sprintf(`upload/upload-%d.json`, size)
		if !util.FileExists(file) {
			util.WriteFile(file, u.body)
		}
	}
}
