package controllers

import (
	"github.com/astaxie/beego"
)

// UsersController operations for Users
type UsersController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsersController) URLMapping() {
	c.Mapping("getAllUser", c.GetAllUser)
}

/*
GetAllUser 获取全部用户数据

@router /user/getAllUser [post]
*/
func (c *UsersController) GetAllUser() {
	beego.Info("hello")
}
