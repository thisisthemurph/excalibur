// Package handler classification of Excalibur API
package handler

import (
	"excalibur/internal/handler/dto"
	"excalibur/internal/handler/hateoas"
)

// swagger:parameters getDataTemplateByID updateDataTemplate addNewColumnToDataTemplate
type dataTemplateIDParamsWrapper struct {
	// The ID of the datatemplate
	//
	// in: path
	// required: true
	ID string `json:"id"`
}

// Data structure representing a single DataTemplate
// swagger:response dataTemplateResponse
type dataTemplateResponseWrapper struct {
	// A DataTemplate object
	// in: body
	Body dto.DataTemplateDTO
}

// A list of DataTemplate objects
// swagger:response dataTemplateListResponse
type dataTemplateListResponseWrapper struct {
	// A list of DataTemplate objects
	// in: body
	Body []dto.DataTemplateDTO
}

// Hateoas model containing the ID of the affected object
type hateoasWithIDResponse struct {
	// the ID of the associated object
	//
	// required: true
	ID string `json:"id"`

	// the list of Hateoas link objects
	//
	// required: true
	hateoas.H
}

// A response detailing the ID of the affected object, with appropriate Hateoas links
// swagger:response hateoasWithIdResponse
type hateoasWithIDResponseWrapper struct {
	// A list of DataTemplate objects
	// in: body
	Body hateoasWithIDResponse
}

// swagger:parameters createDataTemplate
type newDataTemplateParamsWrapper struct {
	Body dto.DataTemplateDTO
}

// swagger:parameters addNewColumnToDataTemplate
type addNewColumnToDataTemplateWrapper struct {
	Body dto.DataTemplateColumnDTO
}
