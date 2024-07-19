package service

import (
	"errors"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, userRequestDto models.UserRequestDto) (*models.UserResponseDto, error) {
	user := &tables.User{
		Name:  userRequestDto.Name,
		Email: userRequestDto.Email,
	}

	existingUser := db.Where(user).First(&tables.User{}).Error
	if existingUser == nil {
		return nil, errors.New("USER ALREADY EXISTS")
	}

	hash, err := hashPassword(userRequestDto.Password)
	if err != nil {
		return nil, errors.New("Unable to encrypt password with error:" + err.Error())
	}

	user.Password = hash

	err = db.Create(user).Error
	if err!= nil {
		return nil, err
	}

	var userResponseDto models.UserResponseDto
	err = copier.Copy(&userResponseDto, user)
	if err != nil {
		return nil, err
	}

	return &userResponseDto, nil
}

func GetUserById(db *gorm.DB, id int) (*models.UserResponseDto, error) {
	var user tables.User

	err := db.First(&user, id).Error
	if err!= nil {
		return nil, err
	}

	var userResponseDto models.UserResponseDto
	err = copier.CopyWithOption(&userResponseDto, &user, copier.Option{
		DeepCopy: true,
	})
	if err != nil {
		return &userResponseDto, errors.New("Failed to copy object with error:" + err.Error())
	}

	userResponseDto.ID = user.ID

	return &userResponseDto, nil
}

func ValidatePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}