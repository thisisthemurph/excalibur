// Package handlers for the web API
package handlers

import (
	"excalibur/internal/services"
	"log"
)

// HandlerCollection a struct for organizing handlers
type HandlerCollection struct {
	DataTemplateHandler DataTemplate
}

// NewHandlerCollection instantiates a new handlers collection
func NewHandlerCollection(sc services.ServiceCollection, l log.Logger) HandlerCollection {
	dataTemplateHandler := NewDataTemplateHandler(sc.DataTemplate, l)

	return HandlerCollection{
		DataTemplateHandler: dataTemplateHandler,
	}
}
