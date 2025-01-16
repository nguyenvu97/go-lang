package impl

import (
	"awesomeProject/global"
	"awesomeProject/src/dto/request"
	"awesomeProject/src/dto/response"
	"awesomeProject/src/jwt"
	"awesomeProject/src/mapper"
	"awesomeProject/src/repository"
	"awesomeProject/src/service"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"log"
)

type UserServiceImpl struct {
	UserRepo repository.IUserRepo

}

func NewUserServiceImpl(userRepo repository.IUserRepo) service.IUserService {
	return &UserServiceImpl{UserRepo: userRepo}
}
func (u UserServiceImpl) Register(user response.UserDto) int {
	data,err := u.UserRepo.GetByEmail(user.Email)
	if err != nil {
		if data != nil {
			return -1
		}
		u.UserRepo.Save(user)
		return 1
	}
	return -1
}
func (u UserServiceImpl) FindByEmail(email string) response.UserDto {
	data,err := u.UserRepo.GetByEmail(email)


	global.LoggerElk.Info(fmt.Sprintf("email is %d",email))
	_ = global.Logger.SendLogToLogstash(email)
	if err != nil {
		return response.UserDto{}
	}
	return mapper.UserToDto(*data)
}
func (u UserServiceImpl) Login(auth request.Authen) response.ResponseJwt {
	k := global.Config.JwtData

	// Lấy thông tin người dùng từ cơ sở dữ liệu
	data, err := u.UserRepo.GetByEmail(auth.Email)
	if err != nil {
		log.Println("Error getting user by email:", err)
		return response.ResponseJwt{}
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.PasswordHash), []byte(auth.PassWord))
	if err != nil {
		log.Println("Password does not match:", err)
		return response.ResponseJwt{}
	}

	// Tạo token JWT
	token, err := jwt.CreateToken(k.Expiration)
	if err != nil {
		log.Println("Error creating token:", err)
		return response.ResponseJwt{}
	}

	refresh, err := jwt.CreateToken(k.RefreshTokenExpiration)
	if err != nil {
		log.Println("Error creating refresh token:", err)
		return response.ResponseJwt{}
	}

	return response.ResponseJwt{
		Token:        token,
		RefreshToken: refresh,
	}
}


func (u UserServiceImpl) DecodeToken(token string) response.DecodeToken {
	data,err := jwt.ExtractClaim(token)
	if err != nil {
		return response.DecodeToken{}
	}
	return response.DecodeToken{
		Subject: data.Subject,
		ID: data.ID,
		ExpiresAt: convetTime(data.ExpiresAt.Time),
		IssuedAt: convetTime(data.IssuedAt.Time),
	}
}

func convetTime(time  time.Time)int64 {
	data := time.Unix()
	return data
}
