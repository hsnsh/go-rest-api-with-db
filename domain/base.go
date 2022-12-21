package domain

import (
	"time"
)

type BaseEntity struct {
	ID        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
