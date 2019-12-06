package controllers

import (
	"github.com/astaxie/beego"
)

// UsersController operations for Users
type UsersController struct {
	beego.Controller
}

/*
GetAllUser 获取全部用户数据

*/
func (c *UsersController) GetAllUser() {
	info := c.GetString("name")
	beego.Info("hello", info)
}
