package models

import "time"

type User struct {
	ID            int         `gorm:"primary_key" json:"id"`
	Username      string      `json:"username"`
	Password      string      `json:"password"`
	Statususer_ID int         `json:"statususer_id"`
	CreatedAt     time.Time   `json:"created_at"`
	UpdateAt      time.Time   `json:"update_at"`
	Status_user   Status_user `json:"-"`
}
