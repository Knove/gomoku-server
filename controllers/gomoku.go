package controllers

import (
	"net/http"
	"server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type GomokuController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{}

func (c *GomokuController) Get() {
	ws, err := websocket.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		beego.Info("ERROR")
		return
	}
	clients[ws] = true
	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		beego.Info(string(p))
		msg := models.Message{Message: string(p) + time.Now().Format("2006-01-02 15:04:05")}
		broadcast <- msg
	}

}
