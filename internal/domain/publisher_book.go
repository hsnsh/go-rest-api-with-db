package domain

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const PublisherBookTableName = "publisher_books"

type PublisherBook struct {
	PublisherID uuid.UUID `gorm:"primary_key;type:uuid;column:publisher_id;"`
	BookID      uuid.UUID `gorm:"primary_key;type:uuid;column:book_id;"`
	OldPrice    float32   `gorm:"column:old_price;"`
	Price       float32   `gorm:"column:price;not null;"`
	Currency    string    `gorm:"column:currency;not null;size:5;"`
	InStock     uint      `gorm:"column:in_stock;not null;"`
	SalesCount  uint      `gorm:"column:sales_count;not null;"`
	OnSale      bool      `gorm:"column:on_sale;not null;default:false"`
}

func (*PublisherBook) TableName() string {
	return PublisherBookTableName
}

func (u *PublisherBook) BeforeSave(tx *gorm.DB) (err error) {
	if u.PublisherID == uuid.Nil {
		err = errors.New(u.PublisherID.String() + " invalid PublisherID")
	}
	if u.BookID == uuid.Nil {
		err = errors.New(u.BookID.String() + " invalid BookID")
	}
	return err
}

func (u *PublisherBook) BeforeCreate(tx *gorm.DB) (err error) {
	if u.OnSale == true {
		err = errors.New("OnSale must be False on creation")
	}
	return err
}
