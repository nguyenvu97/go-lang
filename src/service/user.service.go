package service

import (
	"awesomeProject/src/dto/request"
	"awesomeProject/src/dto/response"
)

type IUserService interface {
	Register(user response.UserDto) int
	FindByEmail(email string) response.UserDto
	Login(auth request.Authen) response.ResponseJwt
	DecodeToken(token string) response.DecodeToken
}
