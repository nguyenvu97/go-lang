package repository

import (
	"awesomeProject/src/database"
	"awesomeProject/src/dto/response"
)

type IUserRepo interface {
	Save(user response.UserDto) int
	GetByEmail(email string) (*database.User, error)
	//Login(email string, password string) bool
}
