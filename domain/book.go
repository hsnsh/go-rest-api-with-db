package domain

import "go-rest-api-with-db/domain/base"

const BookTableName = "books"

type Book struct {
	base.AuditEntity
	Code  string
	Price uint
}
