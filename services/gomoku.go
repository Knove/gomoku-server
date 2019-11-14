package services

import (
	"encoding/json"
	"server/common"
	"server/models"
	"time"

	"github.com/astaxie/beego"
)

/*
Gomoku 五子棋服务层

*/
type Gomoku struct {
}

/*
Ping 测试链接

*/
func (gomoku *Gomoku) Ping(client *models.Client, traceID string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	beego.Info("webSocket request ping 接口 =>", client.Addr, traceID, message)

	data = "pong"

	return
}

/*
Join 加入房间

*/
func (gomoku *Gomoku) Join(client *models.Client, traceID string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	request := &models.Request{}
	beego.Info("webSocket request ping 接口 =>", client.Addr, traceID, message)

	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		beego.Error("解析数据失败", traceID, err)

		return
	}

	// 信息封装
	client.Login(request.UserID, currentTime)

	userClient := &models.UserClient{
		UserID: request.UserID,
		Client: client,
	}

	if models.ClientManagerHandler.InRoom("test") {
		models.ClientManagerHandler.Rooms["test"].Register <- userClient
		beego.Info("webSocket 用户加入房间", traceID, "userID", request.UserID, "RoomName", "test")
	} else {
		room := models.NewRoom("test")
		go room.Start()

		room.Register <- userClient
		models.ClientManagerHandler.AddRoom <- room
		beego.Info("webSocket 用户创建并加入房间", traceID, "userID", request.UserID, "RoomName", "test")
	}

	return
}
