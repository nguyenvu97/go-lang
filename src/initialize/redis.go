package initialize

import (
	"awesomeProject/global"
	"context"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

// InitRedis khởi tạo Redis Client và Redis Lock Client
func InitRedis() {
	r := global.Config.Redis

	// Khởi tạo Redis Client
	rdb := goredislib.NewClient(&goredislib.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: "", // nếu có password, thay đổi
		DB:       0,  // chọn DB nếu cần
		PoolSize: r.Pool_size,
	})

	// Kiểm tra kết nối Redis
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Info("Redis not connected", zap.Error(err))
		return
	}
	// Lưu đối tượng Redis Client vào global.Rdb
	global.Rdb = rdb
	pool := goredis.NewPool(rdb)
	// Khởi tạo Redis Lock Client với Redis Client đã tạo
	rs := redsync.New(pool)
	global.RBD = rs

	global.Logger.Info("Redis initialized successfully")
}
