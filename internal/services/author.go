package services

import (
	"github.com/HsnCorp/go-hsn-library/logger"
	guid "github.com/satori/go.uuid"
	. "go-rest-api-with-db/internal/domain"
	. "go-rest-api-with-db/internal/dtos"
	"go-rest-api-with-db/internal/helpers"
	rep "go-rest-api-with-db/internal/repositories"
)

type IAuthorAppService interface {
	GetAuthorList() ([]AuthorDto, error)
	GetAuthorById(id guid.UUID) (AuthorDto, error)
	CreateAuthor(input AuthorCreateDto) (AuthorDto, error)
	UpdateAuthor(id guid.UUID, input AuthorUpdateDto) (AuthorDto, error)
	DeleteAuthor(id guid.UUID) error
}

type authorAppService struct {
	logger logger.IFileLogger
	dal    rep.IDataAccessLayer
}

func NewAuthorAppService(logger logger.IFileLogger, dal rep.IDataAccessLayer) IAuthorAppService {
	return &authorAppService{
		logger: logger,
		dal:    dal,
	}
}

func (aas *authorAppService) GetAuthorList() ([]AuthorDto, error) {

	entities, err := aas.dal.AuthorRepository().GetList()
	if err != nil {
		return nil, err
	}

	var dtos []AuthorDto
	for _, entity := range entities {
		dtos = append(dtos, entityToDto(entity))
	}

	return dtos, nil
}

func (aas *authorAppService) GetAuthorById(id guid.UUID) (AuthorDto, error) {

	author, err := aas.dal.AuthorRepository().GetById(id)
	if err != nil {
		return AuthorDto{}, err
	}

	return entityToDto(author), nil
}

func (aas *authorAppService) CreateAuthor(input AuthorCreateDto) (AuthorDto, error) {

	created := Author{
		Name: input.Name,
	}

	err := aas.dal.AuthorRepository().Add(&created)
	if err != nil {
		return AuthorDto{}, err
	}

	return entityToDto(created), nil
}

func (aas *authorAppService) UpdateAuthor(id guid.UUID, input AuthorUpdateDto) (AuthorDto, error) {

	author, errFind := aas.dal.AuthorRepository().GetById(id)
	if errFind != nil {
		return AuthorDto{}, errFind
	}

	author.Name = input.Name

	errUpdate := aas.dal.AuthorRepository().Update(&author)
	if errUpdate != nil {
		return AuthorDto{}, errUpdate
	}

	return entityToDto(author), nil
}

func (aas *authorAppService) DeleteAuthor(id guid.UUID) error {

	_, errFind := aas.dal.AuthorRepository().GetById(id)
	if errFind != nil {
		return errFind
	}

	err := aas.dal.AuthorRepository().Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func entityToDto(e Author) AuthorDto {
	return AuthorDto{
		Base: helpers.MapBaseEntityToBaseDto(e.Entity),
		Name: e.Name,
	}
}
