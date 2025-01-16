package initialize

import (
	"awesomeProject/global"
	"awesomeProject/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
	global.LoggerElk = logger.NewLoggerELk(global.Config.Logger)
}

