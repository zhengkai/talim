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

	upload.Test()

	db.WaitConn()

	go web.Server()
}

func afterRun() {
}
