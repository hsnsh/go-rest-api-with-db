package repositories

import (
	"github.com/google/uuid"
	. "go-rest-api-with-db/domain"
	"strconv"
	"time"
)

type IProductRepository interface {
	GetList() []Product
	GetById(id string) Product
	Add(input Product) Product
	Update(input Product) Product
	Delete(id string)
}

type productRepository struct {
	productStore map[string]Product
}

func (p productRepository) GetList() []Product {

	var products []Product
	for _, v := range p.productStore {
		products = append(products, v)
	}

	return products
}

func (p productRepository) GetById(id string) Product {
	productEntity, exist := p.productStore[id]
	if !exist {
		panic(id + " not found")
	}

	return productEntity
}

func (p productRepository) Add(input Product) Product {

	if input.ID == uuid.Nil {
		input.ID = uuid.New()
	}
	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Time{}

	// Add store entity to store
	p.productStore[input.ID.String()] = input

	return input
}

func (p productRepository) Update(input Product) Product {
	id := input.ID.String()
	_, exist := p.productStore[id]
	if !exist {
		panic(id + " not found")
	}
	input.UpdatedAt = time.Now()

	p.productStore[id] = input

	return input
}

func (p productRepository) Delete(id string) {
	_, exist := p.productStore[id]
	if !exist {
		panic(id + " not found")
	}

	delete(p.productStore, id)
}

func NewProductRepository() IProductRepository {

	instance := productRepository{}
	instance.productStore = make(map[string]Product)

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

	return instance
}
