// Package dto for handling Data Transfer Objects for API requests and responses
package dto

import "excalibur/internal/handler/hateoas"

// DataTemplateColumnDTO represents a column of a DataTemplate
type DataTemplateColumnDTO struct {
	OriginalName string `json:"originalName"`
	PrettyName   string `json:"prettyName"`
	DataType     string `json:"dataType"`
}

// DataTemplateDTO represents the a data template structure
type DataTemplateDTO struct {
	Name    string                  `json:"name"`
	Columns []DataTemplateColumnDTO `json:"columns"`
}

// DataTemplateWithHateoasDTO for returning a result with resource pointer
type DataTemplateWithHateoasDTO struct {
	DataTemplateID string `json:"id"`
	hateoas.H
}
