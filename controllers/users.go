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
	c.Mapping("getAllUser", c.getAllUser)
}

// @router /getAllUser [post]
func (c *UsersController) getAllUser() {
	beego.Info("hello")
}
