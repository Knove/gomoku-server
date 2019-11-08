package models

import (
	"fmt"
	"runtime/debug"

	"github.com/gorilla/websocket"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

/*
Client 客户端 链接

*/
type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	UserID        string          // 用户Id，用户登录以后才有
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
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		fmt.Println("读取客户端数据 关闭send", client)
		close(client.Send)
	}()

	for {
		_, message, err := client.Socket.ReadMessage()
		if err != nil {
			fmt.Println("读取客户端数据 错误", client.Addr, err)

			return
		}

		// 处理程序
		fmt.Println("读取客户端数据 处理:", string(message))
		ProcessData(client, message)
	}
}

/*
Write 写数据

*/
func (client *Client) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)

		}
	}()

	defer func() {
		// clientManager.Unregister <- c
		client.Socket.Close()
		fmt.Println("Client发送数据 defer", client)
	}()

	for {
		select {
		case message, ok := <-client.Send:
			if !ok {
				// 发送数据错误 关闭连接
				fmt.Println("Client发送数据 关闭连接", client.Addr, "ok", ok)

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
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
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
