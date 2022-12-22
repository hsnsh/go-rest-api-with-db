package base

import (
	"gorm.io/gorm"
)

type FullAuditEntity struct {
	AuditEntity
	DeletionTime gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (u *FullAuditEntity) BeforeSave(tx *gorm.DB) (err error) {
	return u.AuditEntity.BeforeSave(tx)
}
func (u *FullAuditEntity) BeforeCreate(tx *gorm.DB) (err error) {
	return u.AuditEntity.BeforeCreate(tx)
}
func (u *FullAuditEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	return u.AuditEntity.BeforeUpdate(tx)
}

func (u *FullAuditEntity) AfterCreate(tx *gorm.DB) (err error) {
	return u.AuditEntity.AfterCreate(tx)
}
func (u *FullAuditEntity) AfterUpdate(tx *gorm.DB) (err error) {
	return u.AuditEntity.AfterUpdate(tx)
}
func (u *FullAuditEntity) AfterSave(tx *gorm.DB) (err error) {
	return u.AuditEntity.AfterSave(tx)
}

func (u *FullAuditEntity) BeforeDelete(tx *gorm.DB) (err error) {
	return u.AuditEntity.BeforeDelete(tx)
}
func (u *FullAuditEntity) AfterDelete(tx *gorm.DB) (err error) {
	return u.AuditEntity.AfterDelete(tx)
}

func (u *FullAuditEntity) AfterFind(tx *gorm.DB) (err error) {
	return u.AuditEntity.AfterFind(tx)
}
