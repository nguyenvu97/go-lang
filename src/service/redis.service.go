package service

import (
	"context"
	"time"
)

type (

	RedisProduct interface {
		SetString(key string , value string,ctx context.Context,time time.Duration)
		SetObject(key string, value interface{},ctx context.Context,time time.Duration)
		GetObject(key string,ctx context.Context )	(interface{},error)
		DeleteKey(key string,ctx context.Context)
	}

)

var (
	RedisService RedisProduct
)

func RedisConvet() RedisProduct {
	if RedisService == nil {
		panic("extend RedisService == nil")
	}
	return RedisService
}
func InitRedisService(i RedisProduct) {
	RedisService = i
}