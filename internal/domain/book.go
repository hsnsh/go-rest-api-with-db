package domain

import (
	guid "github.com/satori/go.uuid"
	. "go-rest-api-with-db/internal/domain/base"
)

type Book struct {
	FullAuditEntity
	AuthorID       guid.UUID `gorm:"type:uuid;not null;"`
	ISBN           string    `gorm:"column:isbn;not null;size:50;"`
	Name           string    `gorm:"column:name;not null;size:250;"`
	CoverImagePath string    `gorm:"column:cover_image_path;"`
	OnSale         bool      `gorm:"column:on_sale;not null;default:false"`
}

func (b *Book) TableName() string {
	return "books"
}
