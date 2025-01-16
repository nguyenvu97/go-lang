package initialize

import "github.com/gin-gonic/gin"

func Run() * gin.Engine {
	loadConfig()
	InitLogger()
	//InitElasticsearch()
	InitMysql()
	InitRedis()
	InitKafka()
	InitServerInterface()


	r := InitRouter()
	return r
}
