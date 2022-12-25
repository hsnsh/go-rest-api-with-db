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

	product, err := aas.dal.AuthorRepository().GetById(id)
	if err != nil {
		return AuthorDto{}, err
	}

	return entityToDto(product), nil
}

func (aas *authorAppService) CreateAuthor(input AuthorCreateDto) (AuthorDto, error) {

	createdProduct := Author{
		Name: input.Name,
	}

	err := aas.dal.AuthorRepository().Add(&createdProduct)
	if err != nil {
		return AuthorDto{}, err
	}

	return entityToDto(createdProduct), nil
}

func (aas *authorAppService) UpdateAuthor(id guid.UUID, input AuthorUpdateDto) (AuthorDto, error) {

	updated := Author{
		//BaseEntityWithSoftDeletion: BaseEntityWithSoftDeletion{ID: id},
		Name: input.Name,
	}

	err := aas.dal.AuthorRepository().Update(&updated)
	if err != nil {
		return AuthorDto{}, err
	}

	return entityToDto(updated), nil
}

func (aas *authorAppService) DeleteAuthor(id guid.UUID) error {

	err := aas.dal.AuthorRepository().Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func entityToDto(e Author) AuthorDto {
	return AuthorDto{
		FullAuditDto: helpers.MapFullAuditEntityToFullAuditDto(e.FullAuditEntity),
		Name:         e.Name,
	}
}
