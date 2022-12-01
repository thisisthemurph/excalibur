// Package dto for handling Data Transfer Objects for API requests and responses
package dto

import "excalibur/internal/handlers/hateoas"

type baseDataTableDTO struct {
	Name string `json:"name"`
}

// NewDataTemplateDTO for creating a new DT
type NewDataTemplateDTO struct {
	baseDataTableDTO
	columns []baseDataTemplateColumnDTO
}

// UpdateDataTemplateDTO for updating an existing DT
type UpdateDataTemplateDTO struct {
	baseDataTableDTO
}

// DataTemplateWithHateoasDTO for returning a result with resource pointer
type DataTemplateWithHateoasDTO struct {
	DataTemplateID string `json:"id"`
	hateoas.H
}
