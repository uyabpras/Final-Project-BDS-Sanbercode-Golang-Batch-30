package models

import "time"

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`

	CreatedAt time.Time   `json:"created_at"`
	UpdateAt  time.Time   `json:"update_at"`
	Cart      []Cart      `json:"-"`
	Orderitem []Orderitem `json:"-"`
	Addres    []Addres    `json:"-"`
}
