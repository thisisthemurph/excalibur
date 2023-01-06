package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// FileQuery queries for uploaded files
type FileQuery interface {
	GetStatus(dataTemplateID string, fileID string) (string, error)
}

type fileQuery struct {
	log                    log.Logger
	dataTemplateQuery      DataTemplateQuery
	dataTemplateCollection *mongo.Collection
}

// NewFileQuery instantiates the object
func (d *dao) NewFileQuery(logger log.Logger) FileQuery {
	c := d.database.Collection("datatemplates")
	return &fileQuery{log: logger, dataTemplateCollection: c}
}

func (q *fileQuery) GetStatus(dataTemplateID string, fileID string) (string, error) {
	return "", nil
}
