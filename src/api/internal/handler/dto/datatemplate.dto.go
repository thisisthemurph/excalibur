// Package dto for handling Data Transfer Objects for API requests and responses
package dto

import "excalibur/internal/handler/hateoas"

// DataTemplateDTO represents the a data template structure
//
// swagger:model
type DataTemplateDTO struct {
	// the name of the DataTemplate
	//
	// required: true
	// example: Registered vehicles data table
	Name string `json:"name"`

	// the list of columns associated with the DataTemplate
	//
	// required: true
	Columns []DataTemplateColumnDTO `json:"columns"`
}

// DataTemplateWithHateoasDTO for returning a result with resource pointer
type DataTemplateWithHateoasDTO struct {
	DataTemplateID string `json:"id"`
	hateoas.H
}
