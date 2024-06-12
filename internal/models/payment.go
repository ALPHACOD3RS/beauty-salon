package models



type PaymentStatus string

const (
	Pending 		 PaymentStatus = "Pending"
	Paid			 PaymentStatus = "paid" 
	CancelledPayment PaymentStatus = "cancelled"
)
type Payment struct {
    PaymentID     string `gorm:"primaryKey;autoIncrement:false" json:"payment_id"`
    UserID        string `json:"user_id"`
    AppointmentID string `json:"appointment_id"`
    Amount        float64 `json:"amount"`
    Status        string `json:"status"`
}

// type Payment struct{
	
//     PaymentID     string        	`gorm:"primaryKey;autoIncrement:false" json:"payment_id"`
//     UserID        string        	`json:"user_id"`
//     User          User        	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user"`
//     AppointmentID string        	`json:"appointment_id"`
//     Appointment   Appointment 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"appointment"`
//     Amount        float64     	`json:"amount"`
//     Status        PaymentStatus `json:"status"` 
// }

