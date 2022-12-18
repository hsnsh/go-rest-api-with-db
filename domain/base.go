package domain

import (
	"github.com/google/uuid"
	"time"
)

type BaseEntity struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
