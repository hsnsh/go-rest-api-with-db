package dtos

import (
	. "go-rest-api-with-db/internal/dtos/base"
)

type AuthorDto struct {
	FullAuditDto
	Name string `json:"name"`
}

type AuthorCreateDto struct {
	Name string `json:"name" validate:"required"`
}

type AuthorUpdateDto struct {
	Name string `json:"name" validate:"required"`
}
