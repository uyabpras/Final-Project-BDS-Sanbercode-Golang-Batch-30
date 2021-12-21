package models

import "time"

type Addres struct {
	ID          int       `gorm:"primary_key" json:"id"`
	User_ID     int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"update_at"`
	Addres_line string    `json:"addres_line"`
	Postal_code int       `json:"postal_code"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	User        User      `json:"-"`
}
