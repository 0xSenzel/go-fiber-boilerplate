package auth

import (
	"errors"
	"log"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/middlewares"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	service "github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/service/user"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"gorm.io/gorm"
)

func LoginUser(db *gorm.DB, userRequestDto models.UserRequestDto) (string, error) {
	var user tables.User

	if err := db.Where("email = ?", userRequestDto.Email).First(&user).Error; 
	err != nil {
		return "", errors.New("USER NOT FOUND")
	}
	
	isValid := service.ValidatePassword(user.Password, userRequestDto.Password)
	if ( isValid != true) {
		return "", errors.New("INVALID PASSWORD")
	}

	token, err := middlewares.GenerateToken(&user)
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return token, nil
}