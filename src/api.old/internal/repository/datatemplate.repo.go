package repository

import (
	"context"
	"errors"
	"excalibur/internal/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DataTemplateQuery queries for DataTemplate objects
type DataTemplateQuery interface {
	GetAllDataTemplates() ([]models.DataTemplate, error)
	GetDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error)
	CreateDataTemplate(dt models.DataTemplate) (string, error)
	UpdateDataTemplate(dataTemplateID string, name string) (*string, error)
	AddNewColumn(dataTemplateID string, column models.DataTemplateColumn) (*models.DataTemplate, error)
	DeleteDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error)
	AddFileMetadata(dataTemplateID string, file models.FileMetadata) (*models.FileMetadata, error)
	UpdateFileStatus(dataTemplateID string, fileID string, status models.FileUploadStatus) error
}

type dataTemplateQuery struct {
	log                    log.Logger
	dataTemplateCollection *mongo.Collection
}

func (d *dao) NewDataTemplateQuery(logger log.Logger) DataTemplateQuery {
	c := d.database.Collection("datatemplates")

	return &dataTemplateQuery{
		log:                    logger,
		dataTemplateCollection: c,
	}
}

func (q *dataTemplateQuery) GetAllDataTemplates() ([]models.DataTemplate, error) {
	ctx := context.Background()
	cursor, err := q.dataTemplateCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	// Use primitive.M if there are any issues with bson
	var templates []models.DataTemplate

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var template models.DataTemplate
		if err = cursor.Decode(&template); err != nil {
			return nil, err
		}

		templates = append(templates, template)
	}

	return templates, nil
}

func (q *dataTemplateQuery) GetDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error) {
	oid, err := primitive.ObjectIDFromHex(dataTemplateID)
	if err != nil {
		log.Println("Unknown ID, could not convert " + dataTemplateID + " into a mongo primitive.ObjectID")
		return nil, err
	}

	filter := bson.M{"_id": oid}

	var result models.DataTemplate
	err = q.dataTemplateCollection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (q *dataTemplateQuery) CreateDataTemplate(dt models.DataTemplate) (string, error) {
	ctx := context.Background()

	inserted, err := q.dataTemplateCollection.InsertOne(ctx, dt)
	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (q *dataTemplateQuery) UpdateDataTemplate(dataTemplateID string, name string) (*string, error) {
	id, err := primitive.ObjectIDFromHex(dataTemplateID)
	if err != nil {
		fmt.Println("Unknown ID, could not convert " + dataTemplateID + " into a mongo primitive.ObjectID")
		log.Fatal(err)
		return nil, err
	}

	// bson.D should be used if the order of the elements matters
	// bson.M  should be used otherwise
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"name": name}}

	_, err = q.dataTemplateCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &dataTemplateID, nil
}

func (q *dataTemplateQuery) AddNewColumn(dataTemplateID string, column models.DataTemplateColumn) (*models.DataTemplate, error) {
	id, err := primitive.ObjectIDFromHex(dataTemplateID)
	if err != nil {
		fmt.Println("Unknown ID, could not convert " + dataTemplateID + " into a mongo primitive.ObjectID")
		log.Fatal(err)
		return nil, err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$push": bson.M{
		"columns": bson.M{
			"originalName": column.OriginalName,
			"prettyName":   column.PrettyName,
			"dataType":     column.DataType,
		},
	}}

	var result models.DataTemplate
	err = q.dataTemplateCollection.FindOneAndUpdate(
		context.TODO(),
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("repo - could not find document", dataTemplateID)
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (q *dataTemplateQuery) DeleteDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error) {
	id, err := primitive.ObjectIDFromHex(dataTemplateID)
	if err != nil {
		fmt.Println("Unknown ID, could not convert " + dataTemplateID + " into a mongo primitive ID")
		log.Println(err)
		return nil, err
	}

	filter := bson.M{"_id": id}

	var result models.DataTemplate
	err = q.dataTemplateCollection.FindOneAndDelete(context.TODO(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (q *dataTemplateQuery) AddFileMetadata(dataTemplateID string, file models.FileMetadata) (*models.FileMetadata, error) {
	var _, err = q.GetDataTemplateByID(dataTemplateID)
	if err != nil {
		return nil, err
	}

	oid, err := primitive.ObjectIDFromHex(dataTemplateID)
	if err != nil {
		return nil, errors.New("could not convert ID from string to ObjectId")
	}

	filter := bson.M{"_id": oid}
	file.ID = primitive.NewObjectID()
	update := bson.M{"$push": bson.M{"files": file}}

	result, err := q.dataTemplateCollection.UpdateOne(context.TODO(), filter, update)
	if result.MatchedCount == 0 {
		return nil, errors.New("file metadata was not updated")
	}

	template, err := q.GetDataTemplateByID(dataTemplateID)
	if err != nil {
		return nil, err
	}

	return &template.Files[len(template.Files)-1], nil
}

func (q *dataTemplateQuery) UpdateFileStatus(dataTemplateID string, fileID string, status models.FileUploadStatus) error {
	dtid, err := primitive.ObjectIDFromHex(dataTemplateID)
	if err != nil {
		return errors.New("could not convert dataTemplateID from string to ObjectId")
	}

	fid, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return errors.New("could not convert fileID from string to ObjectId")
	}

	filter := bson.M{"_id": dtid, "files._id": fid}
	update := bson.M{"$set": bson.M{"files.$.status": status.String()}}
	r, err := q.dataTemplateCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return err
	}

	if r.ModifiedCount < 1 {
		return errors.New("the status of the file was not updated")
	}

	return nil
}
