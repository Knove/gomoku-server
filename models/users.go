package models

import (
	"github.com/jinzhu/gorm"

	// Register some standard stuff
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"server/common"
)

/*
User User

*/
type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Score         string `json:"score"`
	Email         string `json:"email"`
	LastLoginDate string `json:"lastLoginDate"`
}

/*
UserDTO UserDTO

*/
type UserDTO struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Score         string `json:"score"`
	Email         string `json:"email"`
	LastLoginDate string `json:"lastLoginDate"`
}

/*
GetAllUser GetAllUser

*/
func GetAllUser() ([]*UserDTO, error) {
	db := common.GetDB()

	var users []*UserDTO
	err := db.Table("users").Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

/*
GetUserByName GetUserByName

*/
func GetUserByName(username string) (*User, error) {
	db := common.GetDB()

	var user User
	err := db.Where("username = ?", username).Table("users").Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
