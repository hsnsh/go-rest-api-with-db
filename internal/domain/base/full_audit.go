package base

import (
	"gorm.io/gorm"
)

type FullAuditEntity struct {
	AuditEntity
	DeletionTime gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (fae *FullAuditEntity) BeforeSave(tx *gorm.DB) (err error) {
	return fae.AuditEntity.BeforeSave(tx)
}
func (fae *FullAuditEntity) BeforeCreate(tx *gorm.DB) (err error) {
	return fae.AuditEntity.BeforeCreate(tx)
}
func (fae *FullAuditEntity) BeforeUpdate(tx *gorm.DB) (err error) {
	return fae.AuditEntity.BeforeUpdate(tx)
}

func (fae *FullAuditEntity) AfterCreate(tx *gorm.DB) (err error) {
	return fae.AuditEntity.AfterCreate(tx)
}
func (fae *FullAuditEntity) AfterUpdate(tx *gorm.DB) (err error) {
	return fae.AuditEntity.AfterUpdate(tx)
}
func (fae *FullAuditEntity) AfterSave(tx *gorm.DB) (err error) {
	return fae.AuditEntity.AfterSave(tx)
}

func (fae *FullAuditEntity) BeforeDelete(tx *gorm.DB) (err error) {
	return fae.AuditEntity.BeforeDelete(tx)
}
func (fae *FullAuditEntity) AfterDelete(tx *gorm.DB) (err error) {
	return fae.AuditEntity.AfterDelete(tx)
}

func (fae *FullAuditEntity) AfterFind(tx *gorm.DB) (err error) {
	return fae.AuditEntity.AfterFind(tx)
}
