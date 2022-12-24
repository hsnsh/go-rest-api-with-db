package dtos

import (
	"go-rest-api-with-db/internal/dtos/base"
)

type ProductDto struct {
	base.FullAuditDto
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ProductCreateDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float32 `json:"price" validate:"min=0"`
}

type ProductUpdateDto struct {
	Name  string  `json:"name" validate:"required"`
	Price float32 `json:"price" validate:"min=0"`
}
