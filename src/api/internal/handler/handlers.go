// Package handler for the web API
package handler

import (
	"excalibur/internal/services"
	"log"
)

// Collection a struct for organizing handlers
type Collection struct {
	DataTemplateHandler DataTemplate
	FileHandler         FileHandler
}

// NewHandlerCollection instantiates a new handlers collection
func NewHandlerCollection(sc services.ServiceCollection, l log.Logger) Collection {
	dataTemplateHandler := NewDataTemplateHandler(sc.DataTemplate, l)
	fileHandler := NewFileHandler(sc.File, l)

	return Collection{
		DataTemplateHandler: dataTemplateHandler,
		FileHandler:         fileHandler,
	}
}
