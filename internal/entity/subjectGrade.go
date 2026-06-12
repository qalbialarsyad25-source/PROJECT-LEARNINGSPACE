package entity

import "github.com/google/uuid"

type SubjectGrade struct {
	Id           uuid.UUID `gorm:"type:varchar(26); primaryKey"`
	Grade        float64   `gorm:"type:decimal(10, 2); not null"`
	ReportCardId uuid.UUID `gorm:"type:char(36)"`
	SubjectId    uuid.UUID `gorm:"type:char(36)"`
}
