package wire

import (
	"awesomeProject/global"
	"awesomeProject/src/controller"
	impl2 "awesomeProject/src/repository/impl"
	"awesomeProject/src/service/impl"

	"database/sql"
	"github.com/google/wire"
)

func ProvideDB1() *sql.DB {
	return global.Mdbcc
}

func InitUserRouterHanldel1() (*controller.UserController, error) {
	wire.Build(
		//ProvideDB,                    // Đảm bảo rằng provider này được sử dụng
		impl2.NewUserRepoImpl,         // Khởi tạo repository
		impl.NewUserServiceImpl,   // Khởi tạo service
		controller.NewUserController, // Khởi tạo controller
	)
	return new(controller.UserController), nil
}
