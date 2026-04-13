package entity

import (
	"time"

	"github.com/google/uuid"
)

type Presence struct {
	Id        uuid.UUID `gorm:"type:varchar(36); primaryKey"`
	Presence  bool      `gorm:"default:false"`
	Content   string    `gorm:"type:varchar(255)"`
	Title     string    `gorm:"type:varchar(100); not null"`
	CreatedAt time.Time `gorm:"type:timeStamp; not null; autoCreatedTime"`
	UserId    uuid.UUID `gorm:"type:char(36)"`
}
