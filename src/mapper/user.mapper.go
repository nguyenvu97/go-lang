package mapper

import (
	"awesomeProject/src/database"
	"awesomeProject/src/dto/response"
	"fmt"
	"github.com/jinzhu/copier"
)

func UserToDto(user database.User) response.UserDto {
	var userDTO response.UserDto
	var createdAt, updatedAt string
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time.Format("2006-01-02 15:04:05")
	}
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time.Format("2006-01-02 15:04:05")
	}
	userDTO.UpdatedAt = createdAt
	userDTO.UpdatedAt = updatedAt
	err := copier.Copy(&userDTO, &user)
	if err != nil {
		fmt.Println("Error structmapping:", err)

	}
	fmt.Println(userDTO)
	return userDTO


}
