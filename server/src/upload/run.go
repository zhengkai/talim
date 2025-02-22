package upload

import "project/zj"

func (u *upload) Run() {
	zj.J(`run`, u.uuid.String(), len(u.body))
}
