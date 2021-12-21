package models

import "time"

type Inventory struct {
	Product_ID int       `json:"product_id"`
	Store_ID   int       `json:"store_id"`
	Stock      int       `json:"stock"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}
