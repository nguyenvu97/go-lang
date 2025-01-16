package impl

import "awesomeProject/src/database"

type implProduct struct {

	db * database.Queries
}

func NewProductImpl(db *database.Queries) *implProduct{
	return &implProduct{
		db : db,
	}
}

func (s * implProduct) GetInfoProductId() error{

	return  nil
}
func (s * implProduct) GetAllProduct()error {

	return nil
}