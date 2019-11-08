package controllers

import (
	"net/http"
	"server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

/*
WebsocketController websocket

*/
type WebsocketController struct {
	beego.Controller
}

/*
Get 建立链接

*/
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

	models.ClientManagerHandler.Register <- client

	// for index := 0; index < 15; index++ {
	// 	time.Sleep(time.Duration(1) * time.Second)
	// 	client.SendMsg([]byte(strconv.Itoa(index)))
	// }

}
