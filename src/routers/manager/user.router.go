package manager

import "github.com/gin-gonic/gin"

type UserAdminRouter struct {
}

func (ur *UserAdminRouter) InitUserRouter(Router *gin.RouterGroup) {

	//private router

		userRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Auth())
	//userRouterPrivate.Use(role)
	{
		userRouterPrivate.GET("/active_user")

	}
}
