package domain

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const BookContentTableName = "book_contents"

type BookContent struct {
	BookID  uuid.UUID `gorm:"primary_key;type:uuid;column:book_id;"`
	Content string    `gorm:"column:currency;not null;"`
}

func (*BookContent) TableName() string {
	return BookContentTableName
}

func (u *BookContent) BeforeSave(tx *gorm.DB) (err error) {
	if u.BookID == uuid.Nil {
		err = errors.New(u.BookID.String() + " invalid BookID")
	}
	return err
}
