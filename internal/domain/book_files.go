package domain

import (
	"errors"
	guid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BookFile struct {
	BookID       guid.UUID `gorm:"primary_key;type:uuid;column:book_id;"`
	FilePath     string    `gorm:"column:file_path;"`
	OrderNumber  uint8     `gorm:"column:order_number;not null;default:0"`
	DocumentType uint8     `gorm:"column:document_type;not null;"`
	// TODO: Enum DocumentType => FullImage, ThumbImage, PDF ...
}

func (bf *BookFile) TableName() string {
	return "book_files"
}

func (bf *BookFile) BeforeSave(tx *gorm.DB) (err error) {
	if bf.BookID == guid.Nil {
		err = errors.New(bf.BookID.String() + " invalid BookID")
	}
	return err
}
