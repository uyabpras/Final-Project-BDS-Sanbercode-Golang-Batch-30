package models

import "time"

type Product struct {
	ID        int         `gorm:"primary_key" json:"id"`
	Name      string      `json:"name"`
	Price     int         `json:"price"`
	CreatedAt time.Time   `json:"created_at"`
	UpdateAt  time.Time   `json:"update_at"`
	Cart      []Cart      `json:"-"`
	Inventory []Inventory `json:"-"`
	Orderitem []Orderitem `json:"-"`
}
