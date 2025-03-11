package upload

import "project/util"

func Test() {
	ab, err := util.ReadFile(`upload/upload-315035.json`)
	if err == nil {
		New(ab)
	}
}
