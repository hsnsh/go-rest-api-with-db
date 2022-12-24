package base

import "time"

type AuditDto struct {
	Dto
	CreationTime     time.Time `json:"creation_time"`
	ModificationTime time.Time `json:"modification_time"`
}
