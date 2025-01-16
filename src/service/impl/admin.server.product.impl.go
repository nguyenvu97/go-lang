package impl

import (
	"awesomeProject/global"
	"awesomeProject/src/database"
	"awesomeProject/src/dto/response"
	"awesomeProject/src/mapper"
	"context"
	"fmt"
)

type implAdminProduct struct {
	db *database.Queries
}


func NewImplAdminProduct(db *database.Queries) *implAdminProduct {
	return &implAdminProduct{
		db: db,
	}
}
func (i *implAdminProduct) CreateProduct(ctx context.Context,product response.ProductDto) int {
	err := i.db.CreateProduct(ctx,database.CreateProductParams{
		ProductCategory: product.ProductCategory,
		ProductName: product.ProductName,
		Price: product.Price,
		ProductUrl: product.ProductURL,
		Quantity: product.Quantity,
		Description: product.Description,
	})



	if err != nil {
		global.LoggerElk.Error(fmt.Sprintf("product ,%p",err))
		return -1
	}
	global.LoggerElk.Info(fmt.Sprintf("product ,%p",product))
	_= global.Logger.SendLogToLogstash(product)
	return 1
}

func (i *implAdminProduct) GetInfoProductId(c context.Context, productId int) response.ProductDto {
	product ,err := i.db.GetProduct(c,productId)
	if err != nil {
		global.LoggerElk.Error(fmt.Sprintf("product ,%p",err))
		return response.ProductDto{}
	}
	global.LoggerElk.Info(fmt.Sprintf("product ,%p",product))
	return mapper.ProductToDto(product)

}


func (i *implAdminProduct) UpdateProduct(c context.Context) error {
	return  nil
}

func (i *implAdminProduct) DeleteProduct(c context.Context) error {
	return  nil
}

