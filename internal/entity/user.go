package entity

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `gorm:"type:char(36);primaryKey"`
	Nama     string    `gorm:"type:varchar(100); not null"`
	Password string    `gorm:"type:varchar(100); not null`
}
