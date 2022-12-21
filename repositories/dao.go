package repositories

import "gorm.io/gorm"

type DAO interface {
	NewBookRepository() IBookRepository
	NewProductRepository() IProductRepository
}

type dao struct {
	db *gorm.DB
}

func (d *dao) NewBookRepository() IBookRepository {
	return NewBookRepository(d.db)
}

func (d *dao) NewProductRepository() IProductRepository {
	return &productRepository{}
}

func NewDAO(db *gorm.DB) DAO {
	return &dao{db: db}
}
