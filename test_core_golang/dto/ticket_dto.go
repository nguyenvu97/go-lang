package dto

import "time"

type TicketDto struct {
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	StockAvailable int       `json:"stockAvailable"`
	StockInitial   int       `json:"stockInitial"`
	IsStockPrepared bool     `json:"isStockPrepared"`
	PriceOriginal  int64     `json:"priceOriginal"`
	PriceFlash     int64     `json:"priceFlash"`
	SaleStartTime  time.Time `json:"saleStartTime"`
	SaleEndTime    time.Time `json:"saleEndTime"`
	Status         int       `json:"status"`
	ActivityId     int64     `json:"activityId"`
	UpdatedAt      time.Time `json:"updatedAt"`
	CreatedAt      time.Time `json:"createdAt"`
}