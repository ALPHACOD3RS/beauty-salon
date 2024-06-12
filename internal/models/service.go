package models



// type Service struct{

// 	ServiceID		string			`json:"service_id" gorm:"primaryKey;autoIncrement:false"`
// 	Name			string			`json:"name"`
// 	Description		string			`json:"description"`
// 	Price			float64			`json:"price"`
// 	Duration		int				`json:"duration"`
//     Appointments []Appointment 		`json:"appointments" gorm:"foreignKey:ServiceID"`
// }

type Service struct {
    ServiceID    string        `json:"service_id" gorm:"primaryKey;autoIncrement:false"`
    Name         string        `json:"name"`
    Description  string        `json:"description"`
    Price        float64       `json:"price"`
    Duration     int           `json:"duration"`
    Appointments []Appointment `json:"appointments" gorm:"foreignKey:ServiceID"`
}

