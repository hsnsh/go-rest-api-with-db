package base

import "time"

type AuditDto struct {
	Base
	CreationTime     time.Time `json:"creation_time"`
	ModificationTime time.Time `json:"modification_time"`
}
