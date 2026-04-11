package entity

import "github.com/google/uuid"

type Subject struct {
	Id               uuid.UUID `gorm:"type:varchar(36); primaryKey"`
	Name             string    `gorm:"type:varchar(100); not null"`
	Subject_Grade_Id uuid.UUID `gorm:"type:varchar(36)"`
}
