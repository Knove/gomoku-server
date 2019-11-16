/**
 * @Desc: Gomoku
 * @Author: Knove
 * @createTime: 2019/11/16 22:17
 * @Email: knove@qq.com
 */

package models

/*
Join 返回格式

*/
type Join struct {
	Type     string `json:"type"`     // 返回类型 Join
	UserID   string `json:"userID"`   // 用户ID
	RoomName string `json:"roomName"` // 房间名
	BackType int32  `json:"backType"` // 返回类型 0 创建房间并加入 1 加入已经存在的房间
}

/*
RoomSay 在房间内的广播

*/
type RoomSay struct {
	Type    string `json:"type"`    // 返回类型 Say 聊天 JoinRoom 加入房间 LeaveRoom 离开房间
	UserID  string `json:"userID"`  // 用户ID
	Content string `json:"content"` // 说的话
}
