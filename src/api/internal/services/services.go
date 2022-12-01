package services

import (
	"excalibur/internal/repository"
	"log"
)

type ServiceCollection struct {
	DataTemplate DataTemplateService
}

func NewServiceCollection(dao repository.DAO, l log.Logger) ServiceCollection {
	dataTemplateQuery := dao.NewDataTemplateQuery(l)

	return ServiceCollection{
		DataTemplate: newDataTemplateService(dataTemplateQuery, l),
	}
}
