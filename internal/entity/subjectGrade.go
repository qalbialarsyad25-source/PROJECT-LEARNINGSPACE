package entity

import "github.com/google/uuid"

type SubjectGrade struct {
	Id           uuid.UUID `gorm:"type:varchar(35); primaryKey"`
	Grade        float64   `gorm:"type:decimal(10, 2); not null"`
	ReportCardId uuid.UUID `gorm:"type:char(36)"`
	Subject      []Subject `gorm:"foreignKey;SubjectGradeId"`
}
