package domain

import (
	"errors"
	guid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BookContent struct {
	BookID  guid.UUID `gorm:"primary_key;type:uuid;column:book_id;"`
	Content string    `gorm:"column:currency;not null;"`
}

func (bc *BookContent) TableName() string {
	return "book_contents"
}

func (bc *BookContent) BeforeSave(tx *gorm.DB) (err error) {
	if bc.BookID == guid.Nil {
		err = errors.New(bc.BookID.String() + " invalid BookID")
	}
	return err
}
