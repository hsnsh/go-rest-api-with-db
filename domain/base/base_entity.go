package base

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Entity struct {
	ID uuid.UUID `gorm:"primary_key;type:uuid;"`
}

func (u *Entity) BeforeSave(tx *gorm.DB) (err error) {
	return
}
func (u *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.NewV4()
	}
	return
}
func (u *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		err = errors.New(u.ID.String() + " invalid ID")
	}
	return err
}

func (u *Entity) AfterCreate(tx *gorm.DB) (err error) {
	return
}
func (u *Entity) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (u *Entity) AfterSave(tx *gorm.DB) (err error) {
	return
}

func (u *Entity) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
func (u *Entity) AfterDelete(tx *gorm.DB) (err error) {
	return
}

func (u *Entity) AfterFind(tx *gorm.DB) (err error) {
	return
}
