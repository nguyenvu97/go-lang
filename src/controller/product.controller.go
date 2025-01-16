package controller

import (
	"awesomeProject/pkg/response"
	"awesomeProject/src/cache"
	responseDto "awesomeProject/src/dto/response"
	"awesomeProject/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)
var Product = new(ProductController)

type ProductController struct {
}


func(pc * ProductController) SaveProduct(ctx * gin.Context){
	var productDto responseDto.ProductDto
		fmt.Print(productDto)
	 err := ctx.ShouldBindJSON(&productDto);
	if err != nil {
		response.SuccessResponse(ctx,7001,"")
		return
	}
	data := service.AdminProduct().CreateProduct(ctx,productDto)
	if data >=1 {
		response.SuccessResponse(ctx, 20001, productDto)
		return
	}
	response.ErrorResponse(ctx, 5001)
	return
}
func(pc * ProductController) FindByProduct(ctx * gin.Context){
	var productDto *responseDto.ProductDto
	productIDStr := ctx.Param("id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		response.SuccessResponse(ctx,7001,"")
		return
	}
	productDto =  cache.GetProductDefaultCacheVip(productID,ctx)
	if productDto.ID >=0 {
		response.SuccessResponse(ctx, 20001, productDto)
		return
	}
	data := service.AdminProduct().GetInfoProductId(ctx,productID)
	if data.ID != 0 {
		cache.GetProductDefaultCacheVip(productID,ctx)
		response.SuccessResponse(ctx, 20001, productDto)
		return
	}
	response.ErrorResponse(ctx, 5001)
	return
}