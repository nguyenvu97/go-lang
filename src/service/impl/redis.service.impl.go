package impl

import (
	"awesomeProject/src/util"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"

	"time"
)

type implRedisService struct {
	redis *redis.Client
}


func (i implRedisService) SetString(key string, value string,ctx context.Context,time time.Duration) {
	if util.StringIsNil(key)== "" {
		fmt.Println("Key không hợp lệ, không thể lưu vào Redis")
		return
	}
	err :=i.redis.Set(ctx,key,value,0).Err()
	if err != nil {
		fmt.Printf("Lỗi khi lưu vào Redis: %v\n", err)
		return
	}
	fmt.Printf("Đã lưu key: %s với value: %s vào Redis thành công với TTL: %v\n", key, value, time)

}

func (i implRedisService) SetObject(key string, value interface{},ctx context.Context,time time.Duration) {
	if util.StringIsNil(key)== "" {
		fmt.Println("Key không hợp lệ, không thể lưu vào Redis")
		return
	}
	err :=i.redis.Set(ctx,key,value,0).Err()
	if err != nil {
		fmt.Printf("Lỗi khi lưu vào Redis: %v\n", err)
		fmt.Printf("dta luu vao redis %s",value)
		return
	}
	fmt.Printf("Đã lưu key: %s với value: %s vào Redis thành công với TTL: %v\n", key, value, time)
}

func (i implRedisService) GetObject(key string,ctx context.Context) (interface{},error) {
	if i.redis == nil {
		fmt.Println("Redis client is not initialized")
		return nil, fmt.Errorf("redis client is not initialized")
	}
	if util.StringIsNil(key)== "" {
		fmt.Println("Key không hợp lệ, không thể lưu vào Redis")
		return nil,nil
	}
	data,err := i.redis.Get(ctx,key).Result()
	if err  != nil{
		fmt.Printf("data is nil  %s\n",data)
		return nil,nil
	}
	object := util.ConvectDataRedis(data)

	return object,nil
}

func (i implRedisService) DeleteKey(key string,ctx context.Context) {
	if util.StringIsNil(key)== "" {
		fmt.Println("Key không hợp lệ, không thể lưu vào Redis")
		return
	}
	err := i.redis.Del(ctx,key).Err()
	if err != nil {
		fmt.Printf("Lỗi khi delete vào Redis: %v\n", err)
		return
	}

}

func NewImplRedisService(redis * redis.Client)  *implRedisService {
	return &implRedisService{
		redis: redis,
	}
}

