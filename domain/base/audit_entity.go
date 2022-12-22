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

func (u *AuditEntity) BeforeSave(tx *gorm.DB) (err error) {
	err = u.Entity.BeforeSave(tx)
	u.ModificationTime = time.Now().UTC()
	return err
}
func (u *AuditEntity) BeforeCreate(tx *gorm.DB) (err error) {
	err = u.Entity.BeforeCreate(tx)
	u.CreationTime = time.Now().UTC()
	u.ModificationTime = time.Now().UTC()
	return err
}
func (u *AuditEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	err = u.Entity.BeforeUpdate(tx)
	u.ModificationTime = time.Now().UTC()
	return err
}

func (u *AuditEntity) AfterCreate(tx *gorm.DB) (err error) {
	return u.Entity.AfterCreate(tx)
}
func (u *AuditEntity) AfterUpdate(tx *gorm.DB) (err error) {
	return u.Entity.AfterUpdate(tx)
}
func (u *AuditEntity) AfterSave(tx *gorm.DB) (err error) {
	return u.Entity.AfterSave(tx)
}

func (u *AuditEntity) BeforeDelete(tx *gorm.DB) (err error) {
	return u.Entity.BeforeDelete(tx)
}
func (u *AuditEntity) AfterDelete(tx *gorm.DB) (err error) {
	return u.Entity.AfterDelete(tx)
}

func (u *AuditEntity) AfterFind(tx *gorm.DB) (err error) {
	return u.Entity.AfterFind(tx)
}
