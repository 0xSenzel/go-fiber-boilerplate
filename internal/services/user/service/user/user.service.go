package service

import (
	"errors"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, userRequestDto models.UserRequestDto) (*tables.User, error) {
	user := &tables.User{
		Name:  userRequestDto.Name,
		Email: userRequestDto.Email,
	}

	existingUser := db.Where(user).First(&tables.User{}).Error
	if existingUser == nil {
		return nil, errors.New("User already exists with email: " + userRequestDto.Email + " and name: " + userRequestDto.Name)
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

	// TODO: map user to dto
	// use copier?
	return &user, nil
}

func ValidatePassword(db *gorm.DB, password string) (bool, error) {
return true, nil
}

