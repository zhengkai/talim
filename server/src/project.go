package project

import (
	"project/build"
	"project/db"
	"project/pb"
	"project/web"
	"project/zj"

	"github.com/zhengkai/zu"
)

func run() {

	build.DumpBuildInfo()

	zj.Init()

	zj.J(zu.JSON(&pb.Demo{
		ID:   43,
		Name: `rpg`,
	}))

	db.WaitConn()

	go web.Server()
}

func afterRun() {
}
