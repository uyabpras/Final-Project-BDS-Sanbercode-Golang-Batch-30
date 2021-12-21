package models

type Cart struct {
	ID         int `gorm:"primary_key" json:"id"`
	User_ID    int `json:"user_id"`
	Product_ID int `json:"product_id"`
	Store_ID   int `json:"store_id"`
}
