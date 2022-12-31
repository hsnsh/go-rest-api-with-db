package repositories

import (
	"errors"
	guid "github.com/satori/go.uuid"
	. "go-rest-api-with-db/internal/domain"
	"gorm.io/gorm"
)

var (
	ErrAuthorNotFound     = errors.New("author was not found")
	ErrAuthorAlreadyExist = errors.New("author already exist")
	ErrAuthorCreation     = errors.New("author creation failed")
)

type IAuthorRepository interface {
	GetList() ([]Author, error)
	GetById(id guid.UUID) (Author, error)
	Add(input *Author) error
	Update(input *Author) error
	Delete(id guid.UUID) error
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) IAuthorRepository {
	return &authorRepository{db: db}
}

func (ar *authorRepository) GetList() ([]Author, error) {

	var entities []Author

	// Get all records
	result := ar.db.Find(&entities)
	// SELECT * FROM users;

	//result.RowsAffected // returns found records count, equals `len(users)`
	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}

func (ar *authorRepository) GetById(id guid.UUID) (Author, error) {

	var entity Author

	//result := ar.db.First(&entity, id)
	// SELECT * FROM users WHERE id = 10;

	result := ar.db.First(&entity, "id = ?", id.String())
	// SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	if result.Error != nil {
		return Author{}, result.Error
	}

	return entity, nil
}

func (ar *authorRepository) Add(input *Author) error {

	result := ar.db.Create(&input) // pass pointer of data to Create

	//user.ID             // returns inserted data's primary key
	//result.Error        // returns error
	//result.RowsAffected // returns inserted records count

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (ar *authorRepository) Update(input *Author) error {

	result := ar.db.Save(&input)
	// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

	if result.Error != nil {
		return result.Error
	}

	return nil

	//var updateAuthor Author
	//result := ar.db.Model(&updateAuthor).Where("id = ?", input.ID.String()).Updates(input)
	//if result.RowsAffected == 0 {
	//	return &Author{}, errors.New("payment data not update")
	//}
	//return updateAuthor, nil
}

func (ar *authorRepository) Delete(id guid.UUID) error {

	// If your model includes a gorm.DeletedAt field (which is included in gorm.Model), it will get soft delete ability automatically!
	// When calling Delete, the record WON’T be removed from the database, but GORM will set the DeletedAt s value to the current time,
	// and the data is not findable with normal Query methods anymore.

	//result := p.db.Delete(&Author{}, id)
	// DELETE FROM users WHERE id = 10;

	result := ar.db.Delete(&Author{}, "id = ?", id.String())
	// DELETE FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no affected rows")
	}

	// Find soft deleted records
	//db.Unscoped().Where("age = 20").Find(&users)
	// SELECT * FROM users WHERE age = 20;

	//Delete permanently
	//db.Unscoped().Delete(&order)
	// DELETE FROM orders WHERE id=10;

	return nil

	//var deletedPayment models.Payment
	//result := ar.db.Where("id = ?", id).Delete(&deletedPayment)
	//if result.RowsAffected == 0 {
	//	return 0, errors.New("payment data not update")
	//}
	//return result.RowsAffected, nil
}
