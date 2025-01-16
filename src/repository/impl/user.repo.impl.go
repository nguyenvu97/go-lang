package impl

import (
	"awesomeProject/global"
	"awesomeProject/src/database"
	"awesomeProject/src/dto/response"
	"awesomeProject/src/repository"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type UserRepoImpl struct {
	sqlc *database.Queries
}
func NewUserRepoImpl() repository.IUserRepo {  // Trả về IUserRepo, nhưng thực tế là một instance của UserRepoImpl
	return &UserRepoImpl{
		sqlc: database.New(global.Mdbcc),
	}
}

	func (u *UserRepoImpl) Save(user response.UserDto) int {

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			// Nếu có lỗi khi mã hóa mật khẩu, trả về -1
			return -1
		}
		_, err1 :=  u.sqlc.CreateUser(context.Background(), database.CreateUserParams{
			Username:     user.Username,
			Email:        user.Email,
			PasswordHash: string(hash),
		})

		if err1 != nil {
			return -1
		}

		return 1
	}
func (u *UserRepoImpl) GetByEmail(email string) (*database.User, error) {
	user, err := u.sqlc.GetUserByEmailSQLC(context.Background(), email)
	if err != nil {
		return nil, err
	}

	return &user, nil // Trả về user tìm được
}
