package models

import (
	"fmt"
	"sync"
)

/*
Room 房间

*/
type Room struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

/*
NewRoom 新建客户端管理实例

*/
func NewRoom() (room *Room) {
	room = &Room{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
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
		case conn := <-room.Register:
			// 建立连接事件
			room.EventRegister(conn)

		case conn := <-room.Unregister:
			// 断开连接事件
			room.EventUnregister(conn)

		}
	}
}

/*
EventRegister 用户加入房间

*/
func (room *Room) EventRegister(client *Client) {
	room.AddClients(client)

	fmt.Println("EventRegister 用户加入房间", client.Addr)

}

/*
EventUnregister 用户离开房间

*/
func (room *Room) EventUnregister(client *Client) {
	room.DelClients(client)

	fmt.Println("EventUnregister 用户离开房间", client.Addr)
}

/*
AddClients 添加客户端

*/
func (room *Room) AddClients(client *Client) {
	room.ClientsLock.Lock()
	defer room.ClientsLock.Unlock()

	room.Clients[client] = true
}

/*
DelClients 删除客户端

*/
func (room *Room) DelClients(client *Client) {
	room.ClientsLock.Lock()
	defer room.ClientsLock.Unlock()

	if _, ok := room.Clients[client]; ok {
		delete(room.Clients, client)
	}
}
