package impl

import (
	"awesomeProject/src/database"
	"context"
)

type implUserServiceProduct struct {
	db * database.Queries
}

func (i *implUserServiceProduct) SetString(key string, value string) {
	//TODO implement me
	panic("implement me")
}

func (i *implUserServiceProduct) SetObject(key string, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (i *implUserServiceProduct) GetObject(key string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (i *implUserServiceProduct) DeleteKey(key string) {
	//TODO implement me
	panic("implement me")
}

func NewImplUserServiceProduct(db * database.Queries) *implUserServiceProduct  {
	return &implUserServiceProduct{
		db: db,
	}
}
func (i *implUserServiceProduct) GetInfoProductId(c context.Context) error {
	return nil
}

func (i *implUserServiceProduct) GetAllProduct(c context.Context) error {
	return nil
}
