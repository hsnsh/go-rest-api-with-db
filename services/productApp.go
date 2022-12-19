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

func (pc *productAppService) GetProductList() ([]ProductDto, error) {

	products, err := pc._productRepository.GetList()
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

	product, err := pc._productRepository.GetById(id.String())
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

	err := pc._productRepository.Add(&createdProduct)
	if err != nil {
		return ProductDto{}, err
	}

	return entityToDto(createdProduct), nil
}

func (pc *productAppService) UpdateProduct(id uuid.UUID, input ProductUpdateDto) (ProductDto, error) {

	updatedProduct := Product{
		BaseEntity: BaseEntity{ID: id},
		Name:       input.Name,
		Price:      input.Price,
	}

	err := pc._productRepository.Update(&updatedProduct)
	if err != nil {
		return ProductDto{}, err
	}

	return entityToDto(updatedProduct), nil
}

func (pc *productAppService) DeleteProduct(id uuid.UUID) error {

	err := pc._productRepository.Delete(id.String())
	if err != nil {
		return err
	}

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
	return &productAppService{_productRepository: NewProductRepository()}
}
