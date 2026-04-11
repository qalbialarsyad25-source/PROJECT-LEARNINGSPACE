package entity

import "github.com/google/uuid"

type Report_Card struct {
	Id            uuid.UUID       `gorm:"varchar(36); primaryKey"`
	Average       float64         `gorm:"type:decimal(10, 2); not null"`
	User_Id       uuid.UUID       `gorm:"type:char(36)"`
	Subject_Grade []Subject_Grade `gorm:"foreignKey:Report_Card_Id"`
}
