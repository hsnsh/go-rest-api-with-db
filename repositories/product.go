package repositories

import (
	"errors"
	"github.com/satori/go.uuid"
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
	productStore map[string]*Product
}

func (p *productRepository) GetList() ([]Product, error) {

	var products []Product
	for _, v := range p.productStore {
		products = append(products, *v)
	}

	return products, nil
}

func (p *productRepository) GetById(id string) (Product, error) {

	searchId, idErr := uuid.FromString(id)
	if idErr != nil {
		return Product{}, errors.New(id + " invalid ID")
	}

	if searchId == uuid.Nil {
		return Product{}, errors.New(id + " invalid ID")
	}

	productEntity, exist := p.productStore[id]
	if !exist {
		return Product{}, errors.New(id + " not found")
	}

	return *productEntity, nil
}

func (p *productRepository) Add(input *Product) error {

	if input.ID == uuid.Nil {
		input.ID = uuid.NewV4()
	}

	input.CreationTime = time.Now()
	input.ModificationTime = time.Time{}

	// Add store entity to store
	p.productStore[input.ID.String()] = input

	return nil
}

func (p *productRepository) Update(input *Product) error {

	if input.ID == uuid.Nil {
		return errors.New(input.ID.String() + " invalid ID")
	}

	id := input.ID.String()
	_, exist := p.productStore[id]
	if !exist {
		return errors.New(id + " not found")
	}
	input.ModificationTime = time.Now()

	p.productStore[id] = input

	return nil
}

func (p *productRepository) Delete(id string) error {

	deleteId, idErr := uuid.FromString(id)
	if idErr != nil {
		return errors.New(id + " invalid ID")
	}

	if deleteId == uuid.Nil {
		return errors.New(id + " invalid ID")
	}

	_, exist := p.productStore[id]
	if !exist {
		return errors.New(id + " not found")
	}

	delete(p.productStore, id)

	return nil
}

func NewProductRepository() IProductRepository {

	instance := productRepository{productStore: map[string]*Product{}}

	// insert fake data
	for i := 0; i < 10; i++ {
		id := uuid.NewV4()
		instance.productStore[id.String()] = &Product{
			Name:  "Product" + strconv.Itoa(i),
			Price: 39.99 + float32(i),
		}
	}

	return &instance
}
