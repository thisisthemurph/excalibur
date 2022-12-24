package services

import (
	"excalibur/internal/handler/errorhandler"
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
	AddNewColumn(dataTemplateID string, column models.DataTemplateColumn) (*models.DataTemplate, error)
	DeleteDataTemplateByID(id string) (*models.DataTemplate, error)
	AddFileMetadata(dataTemplateID string, file models.FileMetadata) (*models.FileMetadata, error)
	UpdateFileStatus(dataTemplateID string, fileID string, status models.FileUploadStatus) error
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

func (s *dataTemplateService) AddNewColumn(dataTemplateID string, column models.DataTemplateColumn) (*models.DataTemplate, error) {
	dt, err := s.dataTemplateQuery.AddNewColumn(dataTemplateID, column)

	if err != nil || (err == nil && dt == nil) {
		return nil, errorhandler.ErrorNotFound
	}

	return dt, nil
}

func (s *dataTemplateService) DeleteDataTemplateByID(id string) (*models.DataTemplate, error) {
	dt, err := s.dataTemplateQuery.DeleteDataTemplateByID(id)

	if err == nil && dt == nil {
		return nil, errorhandler.ErrorNotFound
	}

	return dt, err
}

func (s *dataTemplateService) AddFileMetadata(dataTemplateID string, file models.FileMetadata) (*models.FileMetadata, error) {
	metadata, err := s.dataTemplateQuery.AddFileMetadata(dataTemplateID, file)
	if err != nil {
		return metadata, errorhandler.ErrorInternalServer
	}

	return metadata, nil
}

func (s *dataTemplateService) UpdateFileStatus(dataTemplateID string, fileID string, status models.FileUploadStatus) error {
	return s.dataTemplateQuery.UpdateFileStatus(dataTemplateID, fileID, status)
}
