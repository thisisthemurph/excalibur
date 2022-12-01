package repository

import (
	"context"
	"excalibur/internal/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// DataTemplateQuery queries for DataTemplate objects
type DataTemplateQuery interface {
	GetAllDataTemplates() ([]models.DataTemplate, error)
	GetDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error)
	CreateDataTemplate(dt models.DataTemplate) (string, error)
	UpdateDataTemplate(dataTemplateID string, name string) (*string, error)
	DeleteDataTemplateByID(dataTemplateID string) (*models.DataTemplate, error)
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
