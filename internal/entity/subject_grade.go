package entity

import "github.com/google/uuid"

type Subject_Grade struct {
	Id             uuid.UUID `gorm:"type:varchar(35); primaryKey"`
	Grade          float64   `gorm:"type:decimal(10, 2); not null"`
	Report_Card_Id uuid.UUID `gorm:"type:char(36)"`
	Subject        []Subject `gorm:"foreignKey;Subject_Grade_Id"`
}
