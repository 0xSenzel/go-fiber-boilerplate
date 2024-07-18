package models

type UserRequestDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	ID 		 uint 	`json:"id"`
	Password string `json:"password" copier:"-"`
}