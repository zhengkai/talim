package project

import (
	"project/build"
	"project/db"
	"project/util"
	"project/view"
	"project/web"
	"project/zj"
)

func run() {

	build.DumpBuildInfo()

	zj.Init()

	db.WaitConn()

	// upload.Test()

	view.New(util.DefaultUUID)

	go web.Server()
}

func afterRun() {
}
