package batchRepository

import (
	"context"
	"fmt"
	"github.com/azambranapr/rmocs_collector_api/src/data/models/batch"
	"github.com/azambranapr/rmocs_collector_api/src/data/models/total"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func NewRepository(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) *IdbRepository {
	validate = validator.New()
	return &IdbRepository{client, ctx, cancel}
}

func (p *IdbRepository) Create(model *batch.Model) error {

	return nil
}

func (p *IdbRepository) Update(model *batch.Model, id string) (int64, error) {
	return 0, nil
}

func (p *IdbRepository) GetAll() (batch.Models, error) {
	collection := p.client.Database("oilcollector_db").Collection("batch")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Printf("Error in query: %v", err)
		return nil, err
	}
	var results []bson.D
	if err := cursor.All(p.ctx, &results); err != nil {
		fmt.Printf("Error in cursor: %v", err)
		return nil, err
	}
	ms := make(batch.Models, 0)
	for _, doc := range results {
		//fmt.Println(doc)
		var s batch.Model
		bsonBytes, _ := bson.Marshal(doc)

		if err = bson.Unmarshal(bsonBytes, &s); err != nil {
			fmt.Printf("Error :%v", err)
		}

		ms = append(ms, &s)
	}

	if err = cursor.Close(p.ctx); err != nil {
		fmt.Printf("Error in close cursor:%v", err)
	}
	fmt.Printf("cursor.Close(%v)\r\n", p.ctx)
	return ms, nil
}

func (p *IdbRepository) GetById(id int) (batch.Models, error) {
	var filter, option interface{}
	filter = bson.D{
		{"IDBATCH", id},
	}

	collection := p.client.Database("oilcollector_db").Collection("batch")
	cursor, err := collection.Find(p.ctx, filter, options.Find().SetProjection(option))
	if err != nil {
		fmt.Printf("Error in query: %v", err)
		return nil, err
	}
	var results []bson.D
	if err := cursor.All(p.ctx, &results); err != nil {
		// handle the error
		fmt.Printf("Error in cursor: %v", err)
		return nil, err
	}
	ms := make(batch.Models, 0)
	for _, doc := range results {
		var s batch.Model
		bsonBytes, _ := bson.Marshal(doc)
		err = bson.Unmarshal(bsonBytes, &s)
		if err != nil {
			fmt.Printf("Error:%v", err)
		}
		ms = append(ms, &s)
	}
	if err = cursor.Close(p.ctx); err != nil {
		fmt.Printf("Error in close cursor:%v", err)
	}
	fmt.Printf("cursor.Close(%v)\r\n", p.ctx)
	return ms, nil
}

func (p *IdbRepository) GetByClaimNum(id string) (batch.Models, error) {
	return nil, nil
}

func (p *IdbRepository) Delete(id string) (int64, error) {
	return 0, nil
}

func (p *IdbRepository) DeleteAll() (int64, error) {
	return 0, nil
}

func (p *IdbRepository) GetTotal() (total.Models, error) {
	return nil, nil
}

func validateStruct(m *batch.Model) (err error) {

	err = validate.Struct(m)

	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		return
	}
	return nil
}
