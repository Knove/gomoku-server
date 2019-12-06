package main

import (
	"os"
	_ "server/routers"

	"github.com/astaxie/beego"

	"server/common"
	"server/models"
)

func main() {

	beego.Info("Knove's Hacker World")

	// INIT WEBSOCKET
	initSocket()

	// INIT GORM
	flag := common.GetInstance().Init()
	if !flag {
		beego.Error("init database failure...")
		os.Exit(1)
	}

	beego.Run()
}

func initSocket() {
	go models.ClientManagerHandler.Start()
}
