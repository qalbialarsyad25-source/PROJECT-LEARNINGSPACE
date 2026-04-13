package entity

import "github.com/google/uuid"

type User struct {
	Id         uuid.UUID    `gorm:"type:char(36);primaryKey"`
	Name       string       `gorm:"type:varchar(100); not null"`
	Email      string       `gorm:"type:varchar(255); not null"`
	Password   string       `gorm:"type:varchar(100); not null"`
	Role       string       `grom:"type:varchar(20); not null"`
	Presence   []Presence   `gorm:"foreignKey:UserId"`
	ReportCard []ReportCard `gorm:"foreignKey:UserId"`
}
