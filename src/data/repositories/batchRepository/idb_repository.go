package batchRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type IdbRepository struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}
