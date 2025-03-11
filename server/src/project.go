package project

import (
	"project/build"
	"project/config"
	"project/db"
	"project/dl"
	"project/web"
	"project/zj"
)

func run() {

	build.DumpBuildInfo()

	zj.Init()

	db.WaitConn()

	// upload.Test()

	zj.J(`token`, config.Token[:13], `...`, config.Token[28:])

	if config.Prod {
		go dl.Loop()
	}

	go web.Server()
}

func afterRun() {
}
