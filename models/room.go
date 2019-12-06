package models

import (
	"server/common"
	"sync"

	log "github.com/sirupsen/logrus"
)

/*
Room 房间

*/
type Room struct {
	Name       string             // 房间号
	RoomType   string             // 房间类型 gomoku
	Users      map[string]*Client // 登录的用户
	UserLock   sync.RWMutex       // 读写锁
	Register   chan *UserClient   // 连接连接处理
	Unregister chan *Client       // 断开连接处理程序
	Exit       chan *Client       // 关闭房间
	Broadcast  chan []byte        // 广播 向全部成员发送数据
}

/*
NewRoom 新建客户端管理实例

*/
func NewRoom(name string, roomType string) (room *Room) {
	room = &Room{
		Name:       name,
		RoomType:   roomType,
		Users:      make(map[string]*Client),
		Register:   make(chan *UserClient, 1000),
		Unregister: make(chan *Client, 1000),
		Exit:       make(chan *Client, 1000),
		Broadcast:  make(chan []byte, 1000),
	}

	return
}

/*
Start 管道处理

*/
func (room *Room) Start() {
	for {
		select {
		case UserClient := <-room.Register:
			// 进入房间事件
			room.EventRegister(UserClient)

		case conn := <-room.Unregister:
			// 离开房间事件
			room.EventUnregister(conn)

		case conn := <-room.Exit:
			// 离开房间事件
			log.Printf("EventRegister 房间关闭", "用户房间名：", conn.RoomName, "真实房间名", room.Name)
			return
		}
	}
}

/*
EventRegister 用户加入房间

*/
func (room *Room) EventRegister(userClient *UserClient) {
	client := userClient.Client
	// 连接存在，在添加
	if ClientManagerHandler.InClient(client) {
		key := userClient.GetKey()
		room.AddClients(key, client)
		backData := &RoomSay{
			Type:    "JoinRoom",
			UserID:  userClient.UserID,
			Content: userClient.UserID + "加入了房间",
		}
		msgByte := DataHandle(common.OK, backData, userClient.Request)
		room.SendAll(msgByte, client)
		log.Printf("EventRegister 用户加入房间", client.Addr)
	} else {
		log.Printf("EventRegister 用户加入房间失败，没有此连接", client.Addr)
	}
}

/*
EventUnregister 用户离开房间

*/
func (room *Room) EventUnregister(client *Client) {

	room.DelClients(client.UserID)

	// 给房间内另外的人广播有人离开事件
	backData := &RoomSay{
		Type:    "LeaveRoom",
		UserID:  client.UserID,
		Content: client.UserID + "离开了房间",
	}
	request := &Request{
		TraceID:  "",            // 此为自发回调，没有此 TraceID
		InfoType: room.RoomType, // 这种情况 InfoType 等于 房间类型 Type
	}

	msgByte := DataHandle(common.OK, backData, request)
	room.SendAll(msgByte, client)
	log.Printf("EventUnregister 用户离开房间", client.Addr)
}

/*
AddClients 添加用户

*/
func (room *Room) AddClients(key string, client *Client) {
	room.UserLock.Lock()
	defer room.UserLock.Unlock()

	room.Users[key] = client
}

/*
DelClients 删除用户

*/
func (room *Room) DelClients(key string) {
	room.UserLock.Lock()
	defer room.UserLock.Unlock()

	if _, ok := room.Users[key]; ok {
		delete(room.Users, key)
	}
}

/*
GetUserClient 获取用户所在的 Client

*/
func (room *Room) GetUserClient(userID string) (client *Client) {
	room.UserLock.RLock()
	defer room.UserLock.RUnlock()

	if value, ok := room.Users[userID]; ok {
		client = value
	}

	return
}

/*
GetUserClients 获取所有用户的 Client

*/
func (room *Room) GetUserClients() (clients []*Client) {

	clients = make([]*Client, 0)
	room.UserLock.RLock()
	defer room.UserLock.RUnlock()
	for _, v := range room.Users {
		clients = append(clients, v)
	}

	return
}

/*
SendAll 房间内聊天广播

*/
func (room *Room) SendAll(msg []byte, client *Client) {
	log.Printf("全员广播", client.UserID, msg)

	clients := room.GetUserClients()
	for _, conn := range clients {
		if conn != client {
			conn.SendMsg(msg)
		}
	}
}
