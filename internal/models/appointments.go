package models

import (
	"time"
)

type BookingStatus string

const (
	Booked 		BookingStatus = "booked"
	CancelledBooking	BookingStatus = "cancelled" 
	Completed	BookingStatus = "compledted"
)

type Appointment struct {
    AppointmentID string         `json:"appointment_id" gorm:"primaryKey;autoIncrement:false"`
    UserID        string         `json:"user_id"`
    User          User           `json:"user" gorm:"foreignKey:UserID"`
    ServiceID     string         `json:"service_id"`
    Service       Service        `json:"service" gorm:"foreignKey:ServiceID"`
    Date          time.Time      `json:"date"`
    Time          time.Time      `json:"time"`
    Status        string         `json:"status"`
    Payments      []Payment      `json:"payments" gorm:"foreignKey:AppointmentID"`
}

// type Appointment struct {
//     AppointmentID string         `json:"appointment_id" gorm:"primaryKey;autoIncrement:false"`
//     UserID        string         `json:"user_id"`
//     User          User           `json:"user" gorm:"foreignKey:UserID"`
//     ServiceID     string         `json:"service_id"`
//     Service       Service        `json:"service" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
//     Date          time.Time      `json:"date"`
//     Time          time.Time      `json:"time"`
//     Status        BookingStatus  `json:"status"`
//     Payments      []Payment      `json:"payments" gorm:"foreignKey:AppointmentID"`
// }


