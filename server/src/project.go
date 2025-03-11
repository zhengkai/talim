package project

import (
	"project/build"
	"project/config"
	"project/db"
	"project/web"
	"project/zj"
)

func run() {

	build.DumpBuildInfo()

	zj.Init()

	db.WaitConn()

	// upload.Test()

	zj.J(`token`, config.Token[:13], `...`, config.Token[28:])

	go web.Server()
}

func afterRun() {
}
