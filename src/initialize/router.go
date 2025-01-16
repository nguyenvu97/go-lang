package initialize

import (
	"awesomeProject/global"
	"awesomeProject/src/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Model == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()

	}



	// limiter := rate.NewLimiter(rate.Every(time.Second), 2)
	user_router := routers.RouterGroupApp.User

	MainGroup := r.Group("api/v1")
	// MainGroup.Use(middleweres.ReteLimitRequest(limiter))

	{
		user_router.InitUserRouter(MainGroup)
	}
	manager_router := routers.RouterGroupApp.Admin
	{
		manager_router.InitAdminRouter(MainGroup)

	}
	return r
}
