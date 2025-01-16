package user

import "github.com/gin-gonic/gin"

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// Public router
	//productRouterPublic := Router.Group("/products")
	//{
	//	productRouterPublic.GET("/search")
	//	productRouterPublic.POST("/detai/:id")
	//}
	//private router
}
