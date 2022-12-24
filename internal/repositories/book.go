package repositories

import (
	. "go-rest-api-with-db/internal/domain"
	"gorm.io/gorm"
)

type IBookRepository interface {
	GetList() ([]Book, error)
	GetById(id uint) (Book, error)
	Add(input *Book) error
	Update(input *Book) error
	Delete(id uint) error
}

type bookRepository struct {
	db *gorm.DB
}

func (p *bookRepository) GetList() ([]Book, error) {

	var entities []Book

	// Get all records
	result := p.db.Find(&entities)
	// SELECT * FROM users;

	//result.RowsAffected // returns found records count, equals `len(users)`
	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (p *bookRepository) GetById(id uint) (Book, error) {

	var entity Book

	result := p.db.First(&entity, id)
	// SELECT * FROM users WHERE id = 10;

	//result :=p.db.First(&entity, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	if result.Error != nil {
		return Book{}, result.Error
	}

	return entity, nil
}

func (p *bookRepository) Add(input *Book) error {

	result := p.db.Create(&input) // pass pointer of data to Create

	//user.ID             // returns inserted data's primary key
	//result.Error        // returns error
	//result.RowsAffected // returns inserted records count

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *bookRepository) Update(input *Book) error {

	result := p.db.Save(&input)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *bookRepository) Delete(id uint) error {

	// If your model includes a gorm.DeletedAt field (which is included in gorm.Model), it will get soft delete ability automatically!
	// When calling Delete, the record WON’T be removed from the database, but GORM will set the DeletedAt s value to the current time,
	// and the data is not findable with normal Query methods anymore.

	result := p.db.Delete(&Book{}, id)
	// DELETE FROM users WHERE id = 10;

	//result :=p.db.Delete(&entity{}, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// DELETE FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	if result.Error != nil {
		return result.Error
	}

	// Find soft deleted records
	//db.Unscoped().Where("age = 20").Find(&users)
	// SELECT * FROM users WHERE age = 20;

	//Delete permanently
	//db.Unscoped().Delete(&order)
	// DELETE FROM orders WHERE id=10;

	return nil
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &bookRepository{db: db}
}
