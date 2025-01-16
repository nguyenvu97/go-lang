package initialize

import (
	"awesomeProject/global"
	"awesomeProject/src/database"
	"awesomeProject/src/service"
	"awesomeProject/src/service/impl"
)

func InitServerInterface()	  {
	query := database.New(global.Mdbcc)

	service.InitAdminProduct(impl.NewImplAdminProduct(query))
	service.InitUserProductInfo(impl.NewImplUserServiceProduct(query))
	service.InitRedisService(impl.NewImplRedisService(global.Rdb))
}
