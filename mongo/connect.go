package mongodriver

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() Storage {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("connect database failed")
	}
	c := client.Database("database").Collection("collection")
	s := &Mongox{
		c,
	}

	return s
}

type Mongox struct {
	collect *mongo.Collection
}
