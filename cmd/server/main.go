package main

import (
	"awesomeProject/src/initialize"
	 "github.com/swaggo/files"
	 "github.com/swaggo/gin-swagger"
	_ "awesomeProject/cmd/swag/docs"
)	

// @title          Test Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8082
// @BasePath  /api/v1
// @schemas  http
func main() {

	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.Run("localhost:8082")
}
