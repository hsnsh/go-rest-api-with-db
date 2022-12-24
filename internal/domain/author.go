package domain

import (
	"go-rest-api-with-db/internal/domain/base"
)

const AuthorTableName = "authors"

type Author struct {
	base.FullAuditEntity
	Title string `gorm:"column:title;not null;size:250;"`
}

func (Author) TableName() string {
	return AuthorTableName
}
