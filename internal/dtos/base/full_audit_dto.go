package base

import "time"

type FullAuditDto struct {
	AuditDto
	DeletionTime time.Time `json:"deletion_time"`
}
