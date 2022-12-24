package domain

import (
	"go-rest-api-with-db/internal/domain/base"
)

const PublisherTableName = "publishers"

type Publisher struct {
	base.FullAuditEntity
	Title string `gorm:"column:title;not null;size:250;"`
}

func (*Publisher) TableName() string {
	return PublisherTableName
}
