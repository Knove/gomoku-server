package models

import (
	"fmt"
	"sync"
)

/*
ClientManager 客户端管理

*/
type ClientManager struct {
	Clients     map[*Client]bool   // 全部的连接
	ClientsLock sync.RWMutex       // 读写锁
	Users       map[string]*Client // 登录的用户 // appId+uuid
	UserLock    sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
	Broadcast   chan []byte        // 广播 向全部成员发送数据
}

/*
NewClientManager 新建客户端管理实例

*/
func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
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
func (clientManager *ClientManager) Start() {
	for {
		select {
		case conn := <-clientManager.Register:
			// 建立连接事件
			clientManager.EventRegister(conn)

		case conn := <-clientManager.Unregister:
			// 断开连接事件
			clientManager.EventUnregister(conn)

		}
	}
}

/*
EventRegister 用户建立连接

*/
func (clientManager *ClientManager) EventRegister(client *Client) {
	clientManager.AddClients(client)

	fmt.Println("EventRegister 用户建立连接", client.Addr)

}

/*
EventUnregister 用户断开连接

*/
func (clientManager *ClientManager) EventUnregister(client *Client) {
	clientManager.DelClients(client)

	fmt.Println("EventUnregister 用户断开连接", client.Addr)
}

/*
AddClients 添加客户端

*/
func (clientManager *ClientManager) AddClients(client *Client) {
	clientManager.ClientsLock.Lock()
	defer clientManager.ClientsLock.Unlock()

	clientManager.Clients[client] = true
}

/*
DelClients 删除客户端

*/
func (clientManager *ClientManager) DelClients(client *Client) {
	clientManager.ClientsLock.Lock()
	defer clientManager.ClientsLock.Unlock()

	if _, ok := clientManager.Clients[client]; ok {
		delete(clientManager.Clients, client)
	}
}
