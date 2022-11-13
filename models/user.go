package models



type User struct {
    Id string `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}