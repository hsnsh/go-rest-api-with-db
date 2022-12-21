package dtos

import (
	"time"
)

type BaseDto struct {
	ID           string    `json:"id"`
	CreationTime time.Time `json:"creation_time"`
	UpdateTime   time.Time `json:"update_time"`
}
