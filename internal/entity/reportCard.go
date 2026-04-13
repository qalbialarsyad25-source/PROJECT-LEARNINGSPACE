package entity

import "github.com/google/uuid"

type ReportCard struct {
	Id           uuid.UUID      `gorm:"varchar(36); primaryKey"`
	Average      float64        `gorm:"type:decimal(10, 2); not null"`
	UserId       uuid.UUID      `gorm:"type:char(36)"`
	SubjectGrade []SubjectGrade `gorm:"foreignKey:ReportCardId"`
}
