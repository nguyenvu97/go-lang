package dto

type TicketDto struct {
	Id             int     `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	//StockAvailable int       `json:"stockAvailable"`
	//StockInitial   int       `json:"stockInitial"`
	//IsStockPrepared bool     `json:"isStockPrepared"`
	//PriceOriginal  int64     `json:"priceOriginal"`
	//PriceFlash     int64     `json:"priceFlash"`
	//SaleStartTime  time.Time `json:"saleStartTime"`
	//SaleEndTime    time.Time `json:"saleEndTime"`
	//Status         int       `json:"status"`
	//ActivityId     int64     `json:"activityId"`
	//UpdatedAt      time.Time `json:"updatedAt"`
	//CreatedAt      time.Time `json:"createdAt"`
}

func Testdata () []TicketDto{
	return []TicketDto{
		{
			Id:              1,
			Name:            "Ticket 1",
			Description:     "Description for Ticket 1",
			//StockAvailable:  100,
			//StockInitial:    150,
			//IsStockPrepared: true,
			//PriceOriginal:   1000,
			//PriceFlash:      800,
			//SaleStartTime:   time.Now(),
			//SaleEndTime:     time.Now().Add(time.Hour * 24),
			//Status:          1,
			//ActivityId:      101,
			//UpdatedAt:       time.Now(),
			//CreatedAt:       time.Now(),
		},
		{
			Id:              2,
			Name:            "Ticket 2",
			Description:     "Description for Ticket 2",
			//StockAvailable:  50,
			//StockInitial:    100,
			//IsStockPrepared: false,
			//PriceOriginal:   1500,
			//PriceFlash:      1200,
			//SaleStartTime:   time.Now(),
			//SaleEndTime:     time.Now().Add(time.Hour * 48),
			//Status:          2,
			//ActivityId:      102,
			//UpdatedAt:       time.Now(),
			//CreatedAt:       time.Now(),
		},
		{
			Id:              3,
			Name:            "Ticket 3",
			Description:     "Description for Ticket 3",
			//StockAvailable:  50,
			//StockInitial:    100,
			//IsStockPrepared: false,
			//PriceOriginal:   1500,
			//PriceFlash:      1200,
			//SaleStartTime:   time.Now(),
			//SaleEndTime:     time.Now().Add(time.Hour * 48),
			//Status:          2,
			//ActivityId:      102,
			//UpdatedAt:       time.Now(),
			//CreatedAt:       time.Now(),
		},

	}
}