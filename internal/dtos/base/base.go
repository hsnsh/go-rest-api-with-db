package base

import guid "github.com/satori/go.uuid"

type Base struct {
	ID guid.UUID `json:"id"`
}
