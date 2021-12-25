package models

import "time"

type Orderitem struct {
	ID        int        `gorm:"primary_key" json:"id"`
	ProductID int        `json:"product_id"`
	UserID    int        `json:"user_id"`
	Price     int        `json:"price_total"`
	CreatedAt time.Time  `json:"created_at"`
	Addres_id int        `json:"addres_id"`
	User      User       `json:"-"`
	Product   Product    `json:"-"`
	Addres    Addres     `json:"-"`
	Checkout  []Checkout `json:"-"`
}
