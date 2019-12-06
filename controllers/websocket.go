package controllers

import (
	"net/http"
	"server/models"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

/*
Websocket 建立链接

*/
func Websocket(c *gin.Context) {
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(c.Writer, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Error("Cannot setup WebSocket connection:", err)
		return
	}

	log.Printf("webSocket 建立连接:", ws.RemoteAddr().String())

	currentTime := uint64(time.Now().Unix())
	client := models.NewClient(ws.RemoteAddr().String(), ws, currentTime)

	go client.Read()
	go client.Write()

	models.ClientManagerHandler.Register <- client

}
