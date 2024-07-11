package models

type UserRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Id 		 string `json:"id"`
	Password string `json:"password"`
}