package models

// import "gorm.io/gorm"

type UserRoles string

const (
	AdminRole    UserRoles = "admin"
	StaffRole    UserRoles = "staff"
	CustomerRole UserRoles = "customer"
)


type User struct {
    UserID       string       `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
    Name         string       `json:"name"`
    Email        string       `json:"email"`
    Phone        string       `json:"phone"`
	Password     string       `json:"-" gorm:"not null"`
    Role         string       `json:"role"`
    Appointments []Appointment `json:"appointments" gorm:"foreignKey:UserID"`
    Payments     []Payment     `json:"payments" gorm:"foreignKey:UserID"`
}

// type User struct {
// 	UserID       string       `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
// 	Name         string       `json:"name" gorm:"not null"`
// 	Email        string       `json:"email" gorm:"not null;unique"`
// 	Phone        string       `json:"phone" gorm:"not null;unique"`
// 	Password     string       `json:"-" gorm:"not null"`
// 	Role         UserRoles     `json:"role" gorm:"default:'customer'"`
// 	Appointments []Appointment `gorm:"foreignKey:UserID;references:UserID" json:"appointments"`
// 	Payments     []Payment     `gorm:"foreignKey:UserID;references:UserID" json:"payments"`
// }
