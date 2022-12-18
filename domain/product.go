package domain

type Product struct {
	BaseEntity
	Name  string  `json:"name" validate:"required"`
	Price float32 `json:"price" validate:"min=0"`
}
