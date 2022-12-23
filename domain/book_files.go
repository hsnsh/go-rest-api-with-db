package domain

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const BookFileTableName = "book_files"

type BookFile struct {
	BookID       uuid.UUID `gorm:"primary_key;type:uuid;column:book_id;"`
	FilePath     string    `gorm:"column:file_path;"`
	OrderNumber  uint8     `gorm:"column:order_number;not null;default:0"`
	DocumentType uint8     `gorm:"column:document_type;not null;"`
	// TODO: Enum DocumentType => FullImage, ThumbImage, PDF ...
}

func (*BookFile) TableName() string {
	return BookFileTableName
}

func (u *BookFile) BeforeSave(tx *gorm.DB) (err error) {
	if u.BookID == uuid.Nil {
		err = errors.New(u.BookID.String() + " invalid BookID")
	}
	return err
}
