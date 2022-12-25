package domain

import (
	. "go-rest-api-with-db/internal/domain/base"
)

type Author struct {
	FullAuditEntity
	Name string `gorm:"column:name;not null;size:250;"`
}

func (a *Author) TableName() string {
	return "authors"
}
