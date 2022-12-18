package dtos

import (
	"github.com/google/uuid"
	"time"
)

type BaseDto struct {
	ID           uuid.UUID `json:"id"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}
