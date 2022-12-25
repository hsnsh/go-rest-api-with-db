package base

import (
	"errors"
	guid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Entity struct {
	ID guid.UUID `gorm:"primary_key;type:uuid;"`
}

func (e *Entity) BeforeSave(tx *gorm.DB) (err error) {
	return
}
func (e *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	if e.ID == guid.Nil {
		e.ID = guid.NewV4()
	}
	return
}
func (e *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	if e.ID == guid.Nil {
		err = errors.New(e.ID.String() + " invalid ID")
	}
	return err
}

func (e *Entity) AfterCreate(tx *gorm.DB) (err error) {
	return
}
func (e *Entity) AfterUpdate(tx *gorm.DB) (err error) {
	return
}
func (e *Entity) AfterSave(tx *gorm.DB) (err error) {
	return
}

func (e *Entity) BeforeDelete(tx *gorm.DB) (err error) {
	return
}
func (e *Entity) AfterDelete(tx *gorm.DB) (err error) {
	return
}

func (e *Entity) AfterFind(tx *gorm.DB) (err error) {
	return
}
