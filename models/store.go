package models

import (
	"final/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Store struct {
	ID        int         `json:"id"`
	Username  string      `json:"username"`
	Password  string      `json:"password"`
	City      string      `json:"city"`
	CreatedAt time.Time   `json:"created_at"`
	UpdateAt  time.Time   `json:"update_at"`
	Cart      []Cart      `json:"-"`
	Inventory []Inventory `json:"-"`
	Checkout  []Checkout  `json:"-"`
}

func VerifyPasswordstore(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheckstore(username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Store{}

	err = db.Model(Store{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Store) SaveUserstore(db *gorm.DB) (*Store, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Store{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Create(&u).Error
	if err != nil {
		return &Store{}, err
	}
	return u, nil
}

func (u *Store) UpdateUserstore(db *gorm.DB) (*Store, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Store{}, errPassword
	}
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	var err error = db.Updates(&u).Error
	if err != nil {
		return &Store{}, err
	}
	return u, nil
}
