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

