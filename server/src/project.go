package project

import (
	"project/build"
	"project/db"
	"project/upload"
	"project/web"
	"project/zj"
)

func run() {

	build.DumpBuildInfo()

	zj.Init()

	db.WaitConn()

	upload.Test()

	go web.Server()
}

func afterRun() {
}
