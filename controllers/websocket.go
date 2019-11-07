package controllers

import (
	"net/http"
	"server/models"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebsocketController struct {
	beego.Controller
}

func (c *WebsocketController) Get() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}

	beego.Info("webSocket 建立连接:", ws.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := models.NewClient(ws.RemoteAddr().String(), ws, currentTime)

	go client.Read()
	go client.Write()

	for index := 0; index < 15; index++ {
		time.Sleep(time.Duration(1) * time.Second)
		client.SendMsg([]byte(strconv.Itoa(index)))
	}
}
