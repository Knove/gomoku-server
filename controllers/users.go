package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"server/common"
	"server/models"
	"server/services"
)

/*
Login 登录

*/
func Login(c *gin.Context) {

	log.WithFields(log.Fields{
		"action": "Login",
	}).Info("Login 接口调用")

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, models.ErrorResponse(common.ParameterIllegal))
		return
	}

	isExist, err := services.CheckUser(&user)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, models.ErrorResponse(common.ServerError))
		return
	}

	if !isExist {
		log.Error(err)
		c.JSON(http.StatusOK, models.ErrorResponse(common.UserNotFindOrPasswordIsWrong))
		return
	}

	token, err := common.GenerateToken(user.Username, user.Password)

	response := models.NewAPIResponse(map[string]string{
		"token": token,
	}, "Success", http.StatusOK)

	c.JSON(http.StatusOK, response)
}

/*
GetAllUser 获取全部用户数据

*/
func GetAllUser(c *gin.Context) {

	log.WithFields(log.Fields{
		"action": "GetAllUser",
	}).Info("GetAllUser 接口调用")

	users, err := services.GetAllUser()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, models.ErrorResponse(common.ServerError))
	}
	response := models.NewAPIResponse(users, "Success", http.StatusOK)

	c.JSON(http.StatusOK, response)
}
