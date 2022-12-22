package domain

import (
	uuid "github.com/satori/go.uuid"
	"go-rest-api-with-db/domain/base"
)

const ProductTableName = "products"

type Product struct {
	base.FullAuditEntity
	Name  string  `gorm:"column:name;not null;size:250;"`
	Price float32 `gorm:"column:price;not null;"`
}

func (Product) TableName() string {
	return ProductTableName
}

const ProductLanguageTableName = "product_languages"

type ProductLanguage struct {
	ProductID  uuid.UUID `gorm:"primary_key;type:uuid;column:product_id;"`
	LanguageID uuid.UUID `gorm:"primary_key;type:uuid;column:language_id;"`
	Code       string
	Name       string
}

func (ProductLanguage) TableName() string {
	return ProductLanguageTableName
}

const CategoryTypeTableName = "category_types"

type CategoryType struct {
	CategoryID uint64 `gorm:"primaryKey;autoIncrement:false;column:category_id;"`
	TypeID     uint64 `gorm:"primaryKey;autoIncrement:false;column:type_id;"`
}

func (CategoryType) TableName() string {
	return CategoryTypeTableName
}
