package global

import (
	"awesomeProject/pkg/logger"
	"awesomeProject/pkg/setting"
	"database/sql"
	r "github.com/go-redsync/redsync/v4"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	LoggerElk *logger.ElasticsearchCore
	Mdbcc *sql.DB
	Rdb * redis.Client
	RBD * r.Redsync
)
