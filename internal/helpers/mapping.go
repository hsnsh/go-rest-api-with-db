package helpers

import (
	e "go-rest-api-with-db/internal/domain/base"
	d "go-rest-api-with-db/internal/dtos/base"
)

func MapBaseEntityToBaseDto(entity e.Entity) d.Base {
	return d.Base{ID: entity.ID}
}

func MapAuditEntityToAuditDto(entity e.AuditEntity) d.AuditDto {
	return d.AuditDto{
		Base:             MapBaseEntityToBaseDto(entity.Entity),
		CreationTime:     entity.CreationTime,
		ModificationTime: entity.ModificationTime,
	}
}

func MapFullAuditEntityToFullAuditDto(entity e.FullAuditEntity) d.FullAuditDto {
	return d.FullAuditDto{
		AuditDto:     MapAuditEntityToAuditDto(entity.AuditEntity),
		DeletionTime: entity.DeletionTime.Time,
	}
}
