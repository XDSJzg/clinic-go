package clinicDB

import (
	"database/sql"
)

type Account struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	ID       int    `gorm:"column:id"`
	Type     int    `gorm:"column:type"`
}

type Doctor struct {
	ID 			 int    		`gorm:"column:id;AUTO_INCREMENT"`
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