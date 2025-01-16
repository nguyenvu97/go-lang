package service

import (
	"awesomeProject/src/dto/response"
	"context"
)

type (
	IProduct interface {
		CreateProduct(c  context.Context,product response.ProductDto) int
		UpdateProduct( c  context.Context) error
		DeleteProduct( c  context.Context) error
		GetInfoProductId(c  context.Context,productId int) response.ProductDto
	}

	IProductInfo interface {
		GetInfoProductId(c  context.Context) error
		GetAllProduct(c  context.Context) error
	}

)
var (
	localAdminProduct IProduct
	localUserProduct  IProductInfo
)

func UserProduct() IProductInfo {
	if localUserProduct == nil {
		panic("extend IProductInfo == nil")
		
	}
	return localUserProduct
}
func InitUserProductInfo(i IProductInfo) {
	localUserProduct = i
}
func AdminProduct() IProduct  {
	if localAdminProduct == nil {
		panic("extend IProduct == nil")

	}
	return localAdminProduct
}

func InitAdminProduct(i IProduct) {
	localAdminProduct = i
	
}
