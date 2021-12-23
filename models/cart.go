package models

type Cart struct {
	ID        int     `gorm:"primary_key" json:"id"`
	UserID    int     `json:"user_id"`
	ProductID int     `json:"product_id"`
	StoreID   int     `json:"store_id"`
	User      User    `json:"-"`
	Product   Product `json:"-"`
	Store     Store   `json:"-"`
}
