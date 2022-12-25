package domain

import (
	. "go-rest-api-with-db/internal/domain/base"
)

type Publisher struct {
	FullAuditEntity
	Title string `gorm:"column:title;not null;size:250;"`
}

func (p *Publisher) TableName() string {
	return "publishers"
}
