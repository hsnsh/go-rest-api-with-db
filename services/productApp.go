package services

import (
	"github.com/google/uuid"
	. "go-rest-api-with-db/domain"
	. "go-rest-api-with-db/dtos"
	. "go-rest-api-with-db/repositories"
)

type IProductAppService interface {
	GetProductList() ([]ProductDto, error)
	GetProductById(id uuid.UUID) (ProductDto, error)
	CreateProduct(input ProductCreateDto) (ProductDto, error)
	UpdateProduct(id uuid.UUID, input ProductUpdateDto) (ProductDto, error)
	DeleteProduct(id uuid.UUID) error
}

type productAppService struct {
	_productRepository IProductRepository
}

func (pc productAppService) GetProductList() ([]ProductDto, error) {

	products := pc._productRepository.GetList()

	var productDtos []ProductDto
	for _, product := range products {
		productDtos = append(productDtos, entityToDto(product))
	}

	return productDtos, nil
}

func (pc productAppService) GetProductById(id uuid.UUID) (ProductDto, error) {

	product := pc._productRepository.GetById(id.String())

	return entityToDto(product), nil
}

func (pc productAppService) CreateProduct(input ProductCreateDto) (ProductDto, error) {

	createdProduct := pc._productRepository.Add(Product{
		Name:  input.Name,
		Price: input.Price,
	})

	return entityToDto(createdProduct), nil
}

func (pc productAppService) UpdateProduct(id uuid.UUID, input ProductUpdateDto) (ProductDto, error) {

	updatedProduct := pc._productRepository.Update(Product{
		BaseEntity: BaseEntity{ID: id},
		Name:       input.Name,
		Price:      input.Price,
	})

	return entityToDto(updatedProduct), nil
}

func (pc productAppService) DeleteProduct(id uuid.UUID) error {

	pc._productRepository.Delete(id.String())

	return nil
}

func entityToDto(product Product) ProductDto {
	return ProductDto{
		BaseDto: BaseDto{
			ID:           product.ID,
			CreationTime: product.CreatedAt,
			UpdateTime:   product.UpdatedAt,
		},
		Name:  product.Name,
		Price: product.Price,
	}
}

func NewProductAppService() IProductAppService {
	return productAppService{_productRepository: NewProductRepository()}
}
