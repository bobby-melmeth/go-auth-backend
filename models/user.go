package models

import "github.com/google/uuid"






type User struct {
    Id uuid.UUID `gorm:"type:string;default:gen_random_uuid()"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}