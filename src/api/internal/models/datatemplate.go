// Package models describes the structure of the database models
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// DataTemplateColumn defines the structure of a column
type DataTemplateColumn struct {
	OriginalName string `bson:"originalName" json:"originalName"`
	PrettyName   string `bson:"prettyName" json:"prettyName"`
	DataType     string `bson:"dataType" json:"dataType"`
}

// DataTemplate defines the structure of a data table
type DataTemplate struct {
	// TODO: Update _id to id for json
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Name    string               `bson:"name" json:"name"`
	Columns []DataTemplateColumn `bson:"columns" json:"columns"`
	Files   []FileMetadata       `bson:"files,omitempty"`
}
