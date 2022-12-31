package domain

import (
	"errors"
	guid "github.com/satori/go.uuid"
	. "go-rest-api-with-db/internal/domain/base"
	"go-rest-api-with-db/internal/helpers"
	"time"
)

var (
	ErrNameIsInvalid = errors.New("name is invalid")
)

type Author struct {
	FullAuditEntity
	name string `gorm:"column:name;not null;size:250;"`
}

func NewAuthor(name string) (*Author, error) {

	instance := Author{}
	instance.ID = guid.NewV4()
	instance.CreationTime = time.Now().UTC()

	if err := instance.SetName(name); err != nil {
		return &Author{}, err
	}

	return &instance, nil
}

func (a *Author) TableName() string {
	return "authors"
}

func (a *Author) SetName(name string) error {

	isRequired, maxLength := helpers.GetGormTagDetails(Author{}, "name")
	if isRequired == true && len(name) < 1 {
		return ErrNameIsInvalid
	}
	if maxLength > 0 && len(name) > maxLength {
		return ErrNameIsInvalid
	}

	a.name = name
	return nil
}

func (a *Author) GetName() string {
	return a.name
}
