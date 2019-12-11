package services

import (
	"server/models"
)

func GetAll() ([]*models.UserDTO, error) {

	users, err := models.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}
