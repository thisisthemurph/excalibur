// Package services contains the available services
package services

import (
	"excalibur/internal/repository"
	"log"
)

// ServiceCollection a collection of available services
type ServiceCollection struct {
	DataTemplate DataTemplateService
	File         FileService
}

// NewServiceCollection creates a collection of available services
func NewServiceCollection(dao repository.DAO, l log.Logger) ServiceCollection {
	dataTemplateQuery := dao.NewDataTemplateQuery(l)

	return ServiceCollection{
		DataTemplate: newDataTemplateService(dataTemplateQuery, l),
		File:         newFileService(l),
	}
}
