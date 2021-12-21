package models

type Status_user struct {
	ID          int    `gorm:"primary_key" json:"id"`
	Status_user string `json:"status_user"`
	Users       []User `json:"-"`
}
