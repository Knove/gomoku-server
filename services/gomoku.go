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
	currentTime := uint64(time.Now().Unix())
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
func (gomoku *Gomoku) Leave(client *models.Client, request *models.Request, message []byte) (code uint32, data interface{}) {
	code = common.OK

	traceID := request.TraceID
	userID := client.UserID

	// 返回数据
	backData := &models.Leave{
		Type:   "Leave",
		UserID: userID,
	}

	if userID == "" {
		backData.BackType = -2
		beego.Info("webSocket request Leave 接口 => 用户未登录", traceID)
		return
	}

	beego.Info("webSocket request Leave 接口 =>", client.Addr, traceID, message)

	dataMap := make(map[string]string)
	if err := json.Unmarshal(message, &dataMap); err != nil {
		code = common.ParameterIllegal
		beego.Error("解析数据失败", traceID, err)

		return
	}

	roomName := dataMap["roomName"]

	// 判断是否还是存在此房间号
	if models.ClientManagerHandler.InRoom(roomName) {
		room := models.ClientManagerHandler.Rooms[roomName]
		// 保证房间内确实有这个人
		if _, ok := room.Users[userID]; ok {
			if len(room.Users) == 1 {
				// 直接删除房间
				delete(models.ClientManagerHandler.Rooms, roomName)
				room.Exit <- client
			} else {
				room.Unregister <- client
			}
			backData.BackType = 0
		}
	} else {
		backData.BackType = -1
	}

	backData.RoomName = roomName

	data = backData

	return
}

/*
Say 房间内说话

*/
func (gomoku *Gomoku) Say(client *models.Client, request *models.Request, message []byte) (code uint32, data interface{}) {
	code = common.OK

	traceID := request.TraceID
	userID := client.UserID

	// 返回数据
	backData := &models.Say{
		Type:   "Say",
		UserID: userID,
	}

	if userID == "" {
		backData.BackType = -2
		beego.Info("webSocket request Say 接口 => 用户未登录", traceID)
		return
	}

	beego.Info("webSocket request Say 接口 =>", client.Addr, traceID, message)

	dataMap := make(map[string]string)
	if err := json.Unmarshal(message, &dataMap); err != nil {
		code = common.ParameterIllegal
		beego.Error("解析数据失败", traceID, err)

		return
	}

	roomName := dataMap["roomName"]
	content := dataMap["content"]

	// 判断是否还是存在此房间号
	if models.ClientManagerHandler.InRoom(roomName) {
		room := models.ClientManagerHandler.Rooms[roomName]
		// 保证房间内确实有这个人
		if _, ok := room.Users[userID]; ok {
			// 给房间内另外的人广播聊天
			backMsgData := &models.RoomSay{
				Type:    "SayRoom",
				UserID:  client.UserID,
				Content: content,
			}
			request := &models.Request{
				TraceID:  "",            // 此为自发回调，没有此 TraceID
				InfoType: room.RoomType, // 这种情况 InfoType 等于 房间类型 Type
			}

			msgByte := models.DataHandle(common.OK, backMsgData, request)

			room.SendAll(msgByte, client)
			backData.BackType = 0
		} else {
			backData.BackType = -1
		}
	} else {
		backData.BackType = -1
	}

	backData.RoomName = roomName
	data = backData

	return
}
