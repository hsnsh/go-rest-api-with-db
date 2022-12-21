package domain

const ProductTableName = "products"

type Product struct {
	BaseEntity
	Name  string
	Price float32
}
