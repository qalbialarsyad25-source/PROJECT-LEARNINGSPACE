package entity

import (
	"github.com/google/uuid"
)

type Information struct {
	Id      uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Name    string    `gorm:"type:varchar(100); not null"`
	Title   string    `gorm:"type:varchar(225); not null"`
	Content string    `gorm:"type:longtext; not null"`
}
