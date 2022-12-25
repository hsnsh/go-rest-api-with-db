package domain

import (
	"errors"
	guid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PublisherBook struct {
	PublisherID guid.UUID `gorm:"primary_key;type:uuid;column:publisher_id;"`
	BookID      guid.UUID `gorm:"primary_key;type:uuid;column:book_id;"`
	OldPrice    float32   `gorm:"column:old_price;"`
	Price       float32   `gorm:"column:price;not null;"`
	Currency    string    `gorm:"column:currency;not null;size:5;"`
	InStock     uint      `gorm:"column:in_stock;not null;"`
	SalesCount  uint      `gorm:"column:sales_count;not null;"`
	OnSale      bool      `gorm:"column:on_sale;not null;default:false"`
}

func (pb *PublisherBook) TableName() string {
	return "publisher_books"
}

func (pb *PublisherBook) BeforeSave(tx *gorm.DB) (err error) {
	if pb.PublisherID == guid.Nil {
		err = errors.New(pb.PublisherID.String() + " invalid PublisherID")
	}
	if pb.BookID == guid.Nil {
		err = errors.New(pb.BookID.String() + " invalid BookID")
	}
	return err
}

func (pb *PublisherBook) BeforeCreate(tx *gorm.DB) (err error) {
	if pb.OnSale == true {
		err = errors.New("OnSale must be False on creation")
	}
	return err
}
