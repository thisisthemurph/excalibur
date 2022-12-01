package services

import (
	"excalibur/internal/handlers/errorhandler"
	"excalibur/internal/models"
	"excalibur/internal/repository"
	"log"
)

// DataTemplateService functionality
type DataTemplateService interface {
	GetAllDataTemplates() ([]models.DataTemplate, error)
	GetDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error)
	CreateDataTemplate(dt models.DataTemplate) (string, error)
	UpdateDataTemplateByID(dataTemplateID string, name string) error
	DeleteDataTemplateByID(id string) (*models.DataTemplate, error)
}

type dataTemplateService struct {
	log               log.Logger
	dataTemplateQuery repository.DataTemplateQuery
}

func newDataTemplateService(q repository.DataTemplateQuery, l log.Logger) DataTemplateService {
	return &dataTemplateService{
		log:               l,
		dataTemplateQuery: q,
	}
}

func (s *dataTemplateService) GetAllDataTemplates() ([]models.DataTemplate, error) {
	templates, err := s.dataTemplateQuery.GetAllDataTemplates()
	if err != nil {
		return nil, err
	}

	return templates, nil
}

func (s *dataTemplateService) GetDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error) {
	template, err := s.dataTemplateQuery.GetDataTemplateByID(dataTemplateID)
	if err != nil {
		return nil, err
	}

	return template, nil
}

func (s *dataTemplateService) CreateDataTemplate(dt models.DataTemplate) (string, error) {
	oid, err := s.dataTemplateQuery.CreateDataTemplate(dt)

	return oid, err
}

func (s *dataTemplateService) UpdateDataTemplateByID(dataTemplateID string, name string) error {
	updatedID, err := s.dataTemplateQuery.UpdateDataTemplate(dataTemplateID, name)

	if err == nil && updatedID == nil {
		return errorhandler.ErrorNotFound
	}

	return nil
}

func (s *dataTemplateService) DeleteDataTemplateByID(id string) (*models.DataTemplate, error) {
	dt, err := s.dataTemplateQuery.DeleteDataTemplateByID(id)

	if err == nil && dt == nil {
		return nil, errorhandler.ErrorNotFound
	}

	return dt, err
}
