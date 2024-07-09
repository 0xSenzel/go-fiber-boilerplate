package auth

import (
	"log"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/middlewares"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, userRequestDto models.UserRequestDto, password string) (string, error) {
	var user tables.User
	if err := db.Where(&tables.User{Email: userRequestDto.Email, Name: userRequestDto.Name}).First(&user).Error; 
	err != nil {
		return "", err
	}

	token, err := middlewares.GenerateToken(&userRequestDto)
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return token, nil
}