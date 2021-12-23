package models

import "time"

type Inventory struct {
	ProductID int       `json:"product_id"`
	StoreID   int       `json:"store_id"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	Product   Product   `json:"-"`
	Store     Store     `json:"-"`
}
