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

	exists, err :=  checkIfUserExists(db, user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("USER ALREADY EXISTS")
	}

	err = db.Create(user).Error
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

func checkIfUserExists(db *gorm.DB, email string) (bool, error) {
	var user tables.User

	err := db.Where("email = ?", email).First(&user).Error
	if err == nil {
		return true, err
	}

	return false, nil
}

