package clinicDB

import (
	"database/sql"
	"time"
)
// status: (0, 1, 2) = (none, doctor, admin)
type Account struct {
	ID       int    `gorm:"column:id;AUTO_INCREMENT"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Type     int    `gorm:"column:type"`
}

type Doctor struct {
	ID 			 int    		`gorm:"column:id"`
	Username 	 string 		`gorm:"column:username"`
	Name 		 string 		`gorm:"column:name"`
	Sex 		 string 		`gorm:"column:sex"`
	Age 		 int 			`gorm:"column:age"`
	Department   string   		`gorm:"column:department"`
	Avatar       sql.NullString `gorm:"column:avatar"`
	Introduction sql.NullString `gorm:"column:introduction"`
}

type Admin struct {
	ID 		 int    `gorm:"column:id"`
	Username string `gorm:"column:username"`
	Name 	 string `gorm:"column:name"`
}

type Medicine struct {
	ID    int    `gorm:"column:id;AUTO_INCREMENT"`
	Name  string `gorm:"column:name"`
	Count int `gorm:"column:count"`
}

type Prescription struct {
	ID 	   	       int        `gorm:"column:id;AUTO_INCREMENT'"`
	Department     string     `gorm:"column:department"`
	DoctorUsername string     `gorm:"column:doctor_username"`
	DoctorName     string     `gorm:"column:doctor_name"`
	PatientID      int        `gorm:"column:patient_id"`
	PatientName    string     `gorm:"column:patient_name"`
	Age 	       int        `gorm:"column:age"`
	Sex		       string     `gorm:"column:sex"`
	Medicines      string     `gorm:"column:medicines"`
	CreateTime     time.Time  `gorm:"column:created_at"`
}
