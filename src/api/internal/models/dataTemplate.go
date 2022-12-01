// Package models describes the structure of the database models
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// DataTemplate defines the structure of a data table
type DataTemplate struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string               `bson:"name" json:"name"`
	Columns []DataTemplateColumn `bson:"columns" json:"columns"`
}
