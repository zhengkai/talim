package project

import (
	"project/build"
	"project/db"
	"project/web"
	"project/zj"
)

func run() {

	build.DumpBuildInfo()

	zj.Init()

	db.WaitConn()

	go web.Server()
}

func afterRun() {
}
