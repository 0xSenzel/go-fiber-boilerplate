package service

import (
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, userRequestDto models.UserRequestDto) (*tables.User, error) {
	user := &tables.User{
		Name:  userRequestDto.Name,
		Email: userRequestDto.Email,
	}

	err := db.Create(user).Error
	if err!= nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(db *gorm.DB, id int) (*tables.User, error) {
	var user tables.User

	err := db.First(&user, id).Error
	if err!= nil {
		return nil, err
	}

	return &user, nil
}

func ValidatePassword(db *gorm.DB, password string) (bool, error) {
return true, nil
}

