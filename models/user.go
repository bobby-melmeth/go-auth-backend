package models

type User struct {
	Id uint `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}