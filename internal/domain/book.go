package domain

import (
	uuid "github.com/satori/go.uuid"
	"go-rest-api-with-db/internal/domain/base"
)

const BookTableName = "books"

type Book struct {
	base.FullAuditEntity
	AuthorID       uuid.UUID `gorm:"type:uuid;not null;"`
	ISBN           string    `gorm:"column:isbn;not null;size:50;"`
	Name           string    `gorm:"column:name;not null;size:250;"`
	CoverImagePath string    `gorm:"column:cover_image_path;"`
	OnSale         bool      `gorm:"column:on_sale;not null;default:false"`
}

func (Book) TableName() string {
	return BookTableName
}
