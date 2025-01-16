package user

import (
	"awesomeProject/src/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// Public router

	userController, err := wire.InitUserRouterHanldel()
	if err != nil {
		panic("Failed to initialize UserController: " + err.Error())
	}

	userRouterPublic := Router.Group("/user")

	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.GET("/getUser", userController.GetUser)
		userRouterPublic.POST("/login",userController.Login)
		userRouterPublic.GET("/decode",userController.DecodeToken)
		//userRouterPublic.POST("/send/Otp")
		//userRouterPublic.POST("/detai/:id")
	}
	//private router

	//userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Auth())
	//userRouterPrivate.Use(role)
	//{
	//userRouterPrivate.GET("/get_info")

	//}
}
