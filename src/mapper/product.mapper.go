package mapper

import (
	"awesomeProject/src/database"
	"awesomeProject/src/dto/response"

	"fmt"
	"github.com/jinzhu/copier"
)

func ProductToDto(product database.Product) response.ProductDto  {
	var productDTO response.ProductDto
	err := copier.Copy(&productDTO, &product)
	if err != nil {
		fmt.Println("Error structmapping:", err)

	}

	return productDTO
}
func ProductToEntity(productDto response.ProductDto) database.Product {
	var product database.Product
	err := copier.Copy(&product,&productDto)
	if err != nil {
		fmt.Println("Error structmapping:", err)
	}
	return product

}
