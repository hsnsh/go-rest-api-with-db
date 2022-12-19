package repositories

import (
	"errors"
	"github.com/google/uuid"
	. "go-rest-api-with-db/domain"
	"strconv"
	"time"
)

type IProductRepository interface {
	GetList() ([]Product, error)
	GetById(id string) (Product, error)
	Add(input *Product) error
	Update(input *Product) error
	Delete(id string) error
}

type productRepository struct {
	productStore map[string]Product
}

func (p *productRepository) GetList() ([]Product, error) {

	var products []Product
	for _, v := range p.productStore {
		products = append(products, v)
	}

	return products, nil
}

func (p *productRepository) GetById(id string) (Product, error) {
	productEntity, exist := p.productStore[id]
	if !exist {
		return Product{}, errors.New(id + " not found")
	}

	return productEntity, nil
}

func (p *productRepository) Add(input *Product) error {

	if input.ID == uuid.Nil {
		input.ID = uuid.New()
	}
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Time{}

	// Add store entity to store
	p.productStore[input.ID.String()] = *input

	return nil
}

func (p *productRepository) Update(input *Product) error {
	id := input.ID.String()
	_, exist := p.productStore[id]
	if !exist {
		return errors.New(id + " not found")
	}
	input.UpdatedAt = time.Now()

	p.productStore[id] = *input

	return nil
}

func (p *productRepository) Delete(id string) error {
	_, exist := p.productStore[id]
	if !exist {
		return errors.New(id + " not found")
	}

	delete(p.productStore, id)

	return nil
}

func NewProductRepository() IProductRepository {

	instance := productRepository{productStore: make(map[string]Product)}

	// insert fake data
	for i := 0; i < 10; i++ {
		id := uuid.New()
		instance.productStore[id.String()] = Product{
			BaseEntity: BaseEntity{
				ID:        id,
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
			Name:  "Product" + strconv.Itoa(i),
			Price: 39.99 + float32(i),
		}
	}

	return &instance
}
