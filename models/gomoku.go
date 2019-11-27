/**
 * @Desc: Gomoku
 * @Author: Knove
 * @createTime: 2019/11/16 22:17
 * @Email: knove@qq.com
 */

package models

/*
Join 加入 返回格式

*/
type Join struct {
	Type     string `json:"type"`     // 返回类型 Join
	UserID   string `json:"userID"`   // 用户ID
	RoomName string `json:"roomName"` // 房间名
	BackType int32  `json:"backType"` // 返回类型 -1 房间人满2人 0 创建房间并加入 1 加入已经存在的房间
}

/*
Leave 离开 返回格式

*/
type Leave struct {
	Type     string `json:"type"`     // 返回类型 Leave
	UserID   string `json:"userID"`   // 用户ID
	RoomName string `json:"roomName"` // 房间名
	BackType int32  `json:"backType"` // 返回类型 -1 失败，房间不存在 -2 用户未登录 0 离开房间
}

/*
Say 房间内讲话 返回格式

*/
type Say struct {
	Type     string `json:"type"`     // 返回类型 Say
	UserID   string `json:"userID"`   // 用户ID
	RoomName string `json:"roomName"` // 房间名
	BackType int32  `json:"backType"` // 返回类型 -1 失败 -2 用户未登录 0 成功
}

/*
RoomSay 在房间内的广播

*/
type RoomSay struct {
	Type    string `json:"type"`    // 返回类型 SayRoom 聊天 JoinRoom 加入房间 LeaveRoom 离开房间
	UserID  string `json:"userID"`  // 用户ID
	Content string `json:"content"` // 说的话
}
