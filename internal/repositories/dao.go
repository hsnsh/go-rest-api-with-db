package repositories

import "gorm.io/gorm"

type IDataAccessLayer interface {
	AuthorRepository() IAuthorRepository
}

type dataAccessLayer struct {
	db *gorm.DB
}

func NewDataAccessLayer(db *gorm.DB) IDataAccessLayer {
	return &dataAccessLayer{db: db}
}

func (dal *dataAccessLayer) AuthorRepository() IAuthorRepository {
	return NewAuthorRepository(dal.db)
}
