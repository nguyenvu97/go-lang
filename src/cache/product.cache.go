package cache

import (
	"awesomeProject/global"
	"awesomeProject/src/dto/response"
	"awesomeProject/src/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"time"
)

func genEventItemKey(productId int) string {
	key := fmt.Sprintf("key%s",productId)
	return key
}
func GetProductDefaultCacheVip(productId int,ctx context.Context) *response.ProductDto {

	var productDto response.ProductDto
	productDtoData, err := service.RedisService.GetObject(genEventItemKey(productId),ctx)
	if err == nil && productDtoData != nil {
		productDto = productDtoData.(response.ProductDto)
		return &productDto

	}

	key := fmt.Sprintf("PRO_LOCK_KEY_ITEM%s",productId)
	mutex := global.RBD.NewMutex(key, redsync.WithExpiry(5 * time.Second))

	// lock lai torng 5 giay
	productDtoData1, err1 := service.RedisService.GetObject(genEventItemKey(productId),ctx)
	// chekc lai trong redis co ko
	if err1 == nil && productDtoData1 != nil {
		productDto = productDtoData1.(response.ProductDto)
		return &productDto
	}
	// ko co goi xuogn db lay len set vao cache
	productDto = service.AdminProduct().GetInfoProductId(ctx,productId)

	if productDto.ID <=0 {
		service.RedisService.SetObject(key,nil,ctx,400)
		return nil
	}
	data, err := json.Marshal(productDto)
	if err != nil {
		fmt.Printf("convect json nill %s",err)
	}
	global.Logger.Info(fmt.Sprintf("product data",data))
	service.RedisService.SetObject(key,data,ctx,400)
	defer mutex.Unlock()
	return &productDto

}