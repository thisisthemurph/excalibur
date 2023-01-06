// Package repository container package for database operations and queries
package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// DAO methods to fetch data
type DAO interface {
	NewDataTemplateQuery(log log.Logger) DataTemplateQuery
}

type dao struct {
	log       log.Logger
	database  *mongo.Database
	dbContext context.Context
}

// NewDAO creates a new data access object
func NewDAO(logger log.Logger) DAO {
	db, ctx := connect()

	return &dao{
		log:       logger,
		database:  db,
		dbContext: ctx,
	}
}
