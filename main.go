package main

import (
	"os"
	"server/common"
	"server/models"
	"server/routers"
	_ "server/routers"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Printf("Knove's Hacker World")

	// INIT LOGRUS

	log.SetFormatter(&log.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// INIT WEBSOCKET
	initSocket()
	log.Printf("InitSocket ...")

	// INIT GORM
	flag := common.GetInstance().Init()
	if !flag {
		log.Error("init database failure...")
		os.Exit(1)
	}
	log.Printf("INIT GORM ...")

	// INIT ROUTER
	routers.Init()
	log.Printf("Init Router ...")
}

func initSocket() {
	go models.ClientManagerHandler.Start()
}
