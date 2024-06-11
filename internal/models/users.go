package model

import "gorm.io/gorm"

type UserRole string

const (
	AdminRole 		UserRole = 	"admin"
	StafRole  		UserRole =  "staff"
	CustomerRole 	UserRole =  "customer"
)


type User  struct {
	gorm.Model

	Name 		string		`json:"name" gorm:"unique, not null"`
	Email 		string		`json:"email" gorm:"not null, unique"`		
	Phone 		string		`json:"phone" gorm:"not null, unique"` 
	Password 	string	`json:"-" gorm:"not null"`
	Role 		UserRole		`json:"role" gorm:"not null"`
}