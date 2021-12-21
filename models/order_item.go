package models

import "time"

type Order_item struct {
	ID          int       `gorm:"primary_key" json:"id"`
	Product_ID  int       `json:"product_id"`
	User_ID     int       `json:"user_id"`
	Price_total int       `json:"price_total"`
	CreatedAt   time.Time `json:"created_at"`
	Order_ID    int       `json:"order_id"`
}
