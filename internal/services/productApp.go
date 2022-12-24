package services

import (
	"github.com/HsnCorp/go-hsn-library/logger"
	"github.com/satori/go.uuid"
	. "go-rest-api-with-db/internal/domain"
	. "go-rest-api-with-db/internal/dtos"
	. "go-rest-api-with-db/internal/dtos/base"
	. "go-rest-api-with-db/internal/repositories"
)

type IProductAppService interface {
	GetProductList() ([]ProductDto, error)
	GetProductById(id uuid.UUID) (ProductDto, error)
	CreateProduct(input ProductCreateDto) (ProductDto, error)
	UpdateProduct(id uuid.UUID, input ProductUpdateDto) (ProductDto, error)
	DeleteProduct(id uuid.UUID) error
}

type productAppService struct {
	_logger logger.IFileLogger
	_dao    DAO
}

func (pc *productAppService) GetProductList() ([]ProductDto, error) {

	products, err := pc._dao.NewProductRepository().GetList()
	if err != nil {
		return nil, err
	}

	var productDtos []ProductDto
	for _, product := range products {
		productDtos = append(productDtos, entityToDto(product))
	}

	return productDtos, nil
}

func (pc *productAppService) GetProductById(id uuid.UUID) (ProductDto, error) {

	product, err := pc._dao.NewProductRepository().GetById(id.String())
	if err != nil {
		return ProductDto{}, err
	}

	return entityToDto(product), nil
}

func (pc *productAppService) CreateProduct(input ProductCreateDto) (ProductDto, error) {

	createdProduct := Product{
		Name:  input.Name,
		Price: input.Price,
	}

	err := pc._dao.NewProductRepository().Add(&createdProduct)
	if err != nil {
		return ProductDto{}, err
	}

	return entityToDto(createdProduct), nil
}

func (pc *productAppService) UpdateProduct(id uuid.UUID, input ProductUpdateDto) (ProductDto, error) {

	updatedProduct := Product{
		//BaseEntityWithSoftDeletion: BaseEntityWithSoftDeletion{ID: id},
		Name:  input.Name,
		Price: input.Price,
	}

	err := pc._dao.NewProductRepository().Update(&updatedProduct)
	if err != nil {
		return ProductDto{}, err
	}

	return entityToDto(updatedProduct), nil
}

func (pc *productAppService) DeleteProduct(id uuid.UUID) error {

	err := pc._dao.NewProductRepository().Delete(id.String())
	if err != nil {
		return err
	}

	return nil
}

func entityToDto(product Product) ProductDto {
	return ProductDto{
		FullAuditDto: FullAuditDto{
			AuditDto: AuditDto{
				Dto:              Dto{ID: product.ID.String()},
				CreationTime:     product.CreationTime,
				ModificationTime: product.ModificationTime,
			},
			DeletionTime: product.DeletionTime.Time,
		},
		Name:  product.Name,
		Price: product.Price,
	}
}

func NewProductAppService(logger logger.IFileLogger, dao DAO) IProductAppService {
	return &productAppService{
		_logger: logger,
		_dao:    dao,
	}
}
