package dtos

type ProductDto struct {
	BaseDto
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ProductCreateDto struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type ProductUpdateDto struct {
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
