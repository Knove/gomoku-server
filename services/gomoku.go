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
func (gomoku *Gomoku) Ping(client *models.Client, request *models.Request, message []byte) (code uint32, data interface{}) {

	code = common.OK
	beego.Info("webSocket request ping 接口 =>", client.Addr, request.TraceID, message)

	data = "pong"

	return
}

/*
Join 加入指定房间号

*/
func (gomoku *Gomoku) Join(client *models.Client, request *models.Request, message []byte) (code uint32, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	traceID := request.TraceID
	userID := request.UserID

	beego.Info("webSocket request Join 接口 =>", client.Addr, traceID, message)

	dataMap := make(map[string]string)
	if err := json.Unmarshal(message, &dataMap); err != nil {
		code = common.ParameterIllegal
		beego.Error("解析数据失败", traceID, err)

		return
	}

	roomName := dataMap["roomName"]

	// 信息封装
	client.Login(userID, currentTime)

	userClient := &models.UserClient{
		UserID:  userID,
		Client:  client,
		Request: request,
	}

	backData := models.Join{Type: "Join", UserID: userID, RoomName: roomName}

	if models.ClientManagerHandler.InRoom(roomName) {
		room := models.ClientManagerHandler.Rooms[roomName]

		userNum := len(room.Users)
		if userNum >= 2 {
			beego.Info("webSocket 用户加入房间 失败 人满", traceID, "userID", userID, "RoomName", roomName, "BackType", 1)
			backData.BackType = -1
		} else {
			userClient.Client.RoomName = roomName
			room.Register <- userClient
			beego.Info("webSocket 用户加入房间", traceID, "userID", userID, "RoomName", roomName, "BackType", 1)
			backData.BackType = 1
		}
	} else {
		room := models.NewRoom(roomName, "gomoku")
		go room.Start()

		userClient.Client.RoomName = roomName
		room.Register <- userClient
		models.ClientManagerHandler.AddRoom <- room
		beego.Info("webSocket 用户创建并加入房间", traceID, "userID", userID, "RoomName", roomName, "BackType", 0)
		backData.BackType = 0
	}

	data = backData

	return
}

/*
Leave 离开某个房间

*/
