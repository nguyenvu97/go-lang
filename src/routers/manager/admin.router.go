package manager

import (
	"awesomeProject/src/controller"
	"awesomeProject/src/middleweres"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// Public router

	adminRouterPublic := Router.Group("/admin")
	adminRouterPublic.Use(middleweres.AuthenMiddleware())
	adminRouterPublic.Use(middleweres.NewRateLimiter().PrivateRateLimiter())

	{
		adminRouterPublic.POST("/add-product",controller.Product.SaveProduct)
		adminRouterPublic.GET("/findByProduct/:id",controller.Product.FindByProduct)
	}
	//private router
	//AdminRouterPrivate := Router.Group("/admin")
	//userRouterPrivate.Use(limiter())
	//userRouterPrivate.Use(Auth())
	//userRouterPrivate.Use(role)
	//{
	//	AdminRouterPrivate.GET("/get_all_user")
	//
	//}
}
