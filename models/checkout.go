package models

import "time"

type Checkout struct {
	ID          int       `json:"id"`
	StoreID     int       `json:"store_id"`
	OrderitemID int       `json:"orderitem_id"`
	CreatedAt   time.Time `json:"created_at"`
	Store       Store     `json:"-"`
	Orderitem   Orderitem `json:"-"`
}
