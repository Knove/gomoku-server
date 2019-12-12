/**
 * @Desc: go_server
 * @Author: Knove
 * @createTime: 2019/11/8 18:04
 * @Email: knove@qq.com
 */

package main

import (
	"os"
	"server/common"
	"server/models"
	"server/routers"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Printf("Knove's Hacker World")

	// INIT LOGRUS
	log.SetFormatter(&log.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	log.Printf("Init Logrus ...")

	// INIT WEBSOCKET
	initSocket()
	log.Printf("Init Socket ...")

	// INIT GORM
	flag := common.GetInstance().Init()
	if !flag {
		log.Fatal("init database failure...")
		os.Exit(1)
	}
	log.Printf("INIT Gorm ...")

	// INIT ROUTER
	log.Printf("Init Router ...")
	routers.Init()

}

func initSocket() {
	go models.ClientManagerHandler.Start()
}
