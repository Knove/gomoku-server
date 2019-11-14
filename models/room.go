package models

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego"
)

/*
Room 房间

*/
type Room struct {
	Name       string             // 房间号
	Users      map[string]*Client // 登录的用户
	UserLock   sync.RWMutex       // 读写锁
	Register   chan *UserClient   // 连接连接处理
	Unregister chan *Client       // 断开连接处理程序
	Broadcast  chan []byte        // 广播 向全部成员发送数据
}

/*
NewRoom 新建客户端管理实例

*/
func NewRoom(name string) (room *Room) {
	room = &Room{
		Name:       name,
		Users:      make(map[string]*Client),
		Register:   make(chan *UserClient, 1000),
		Unregister: make(chan *Client, 1000),
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
	}
	room.sendAll("加入了房间！", client)
	beego.Info("EventRegister 用户加入房间", client.Addr)

}

/*
EventUnregister 用户离开房间

*/
func (room *Room) EventUnregister(client *Client) {

	room.DelClients(client.UserID)

	fmt.Println("EventUnregister 用户离开房间", client.Addr)
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
sendAll 房间内聊天广播

*/
func (room *Room) sendAll(data string, client *Client) {
	beego.Info("全员广播", client.UserID, data)

	clients := room.GetUserClients()
	for _, conn := range clients {
		if conn != client {
			conn.SendMsg([]byte(data))
		}
	}
}
