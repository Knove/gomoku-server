package services

import (
	"server/models"
)

/*
GetAllUser 获取全部用户数据

*/
func GetAllUser() ([]*models.UserDTO, error) {

	users, err := models.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

/*
CheckUser 检测用户是否合法

*/
func CheckUser(user *models.User) (bool, error) {

	realUser, err := models.GetUserByName(user.Username)
	if err != nil {
		return false, err
	}

	if realUser.Password == user.Password && realUser.Password != "" {
		return true, nil
	}

	return false, nil

}
