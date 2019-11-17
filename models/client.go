package models

import (
	"runtime/debug"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

/*
UserClient 用户信息

*/
type UserClient struct {
	UserID  string   // 用户ID
	Client  *Client  // 用户连接实例
	Request *Request // 用户参与游戏的类型
}

/*
GetKey 获取用户的 key

*/
func (l *UserClient) GetKey() (key string) {
	key = l.UserID

	return
}

/*
Client 客户端 链接

*/
type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	UserID        string          // 用户Id，用户登录以后才有
	RoomName      string          // 房间名，加入房间才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

/*
NewClient 创建一个链接

*/
func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 100),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}

	return
}

/*
Read 读数据

*/
func (client *Client) Read() {
	defer func() {
		if r := recover(); r != nil {
			beego.Info("read stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		beego.Info("读取客户端数据 关闭send", client)
		close(client.Send)
	}()

	for {
		_, message, err := client.Socket.ReadMessage()
		if err != nil {
			beego.Error("读取客户端数据 错误", client.Addr, err)

			return
		}

		// 处理程序
		beego.Info("读取客户端数据 处理:", string(message))
		ProcessData(client, message)
	}
}

/*
Write 写数据

*/
func (client *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			beego.Info("write stop", string(debug.Stack()), r)

		}
	}()

	defer func() {
		ClientManagerHandler.Unregister <- client

		client.Socket.Close()
		beego.Info("Client发送数据 defer", client)
	}()

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				// 发送数据错误 关闭连接
				beego.Error("Client发送数据 关闭连接", client.Addr, "ok", ok)

				return
			}

			client.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

/*
SendMsg 发送数据

*/
func (client *Client) SendMsg(msg []byte) {

	if client == nil {

		return
	}

	defer func() {
		if r := recover(); r != nil {
			beego.Info("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	client.Send <- msg
}

/*
Close 关闭链接

*/
func (client *Client) Close() {
	close(client.Send)
}

/*
Login 登录

*/
func (client *Client) Login(userID string, loginTime uint64) {
	client.UserID = userID
	client.LoginTime = loginTime
	// 心跳
	client.Heartbeat(loginTime)
}

/*
Heartbeat 心跳

*/
func (client *Client) Heartbeat(currentTime uint64) {
	client.HeartbeatTime = currentTime

	return
}
