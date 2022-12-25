package base

import (
	"gorm.io/gorm"
	"time"
)

type AuditEntity struct {
	Entity
	CreationTime     time.Time `gorm:"column:created_at;not null;"`
	ModificationTime time.Time `gorm:"column:updated_at"`
}

func (ae *AuditEntity) BeforeSave(tx *gorm.DB) (err error) {
	err = ae.Entity.BeforeSave(tx)
	ae.ModificationTime = time.Now().UTC()
	return err
}
func (ae *AuditEntity) BeforeCreate(tx *gorm.DB) (err error) {
	err = ae.Entity.BeforeCreate(tx)
	ae.CreationTime = time.Now().UTC()
	ae.ModificationTime = time.Now().UTC()
	return err
}
func (ae *AuditEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	err = ae.Entity.BeforeUpdate(tx)
	ae.ModificationTime = time.Now().UTC()
	return err
}

func (ae *AuditEntity) AfterCreate(tx *gorm.DB) (err error) {
	return ae.Entity.AfterCreate(tx)
}
func (ae *AuditEntity) AfterUpdate(tx *gorm.DB) (err error) {
	return ae.Entity.AfterUpdate(tx)
}
func (ae *AuditEntity) AfterSave(tx *gorm.DB) (err error) {
	return ae.Entity.AfterSave(tx)
}

func (ae *AuditEntity) BeforeDelete(tx *gorm.DB) (err error) {
	return ae.Entity.BeforeDelete(tx)
}
func (ae *AuditEntity) AfterDelete(tx *gorm.DB) (err error) {
	return ae.Entity.AfterDelete(tx)
}

func (ae *AuditEntity) AfterFind(tx *gorm.DB) (err error) {
	return ae.Entity.AfterFind(tx)
}
