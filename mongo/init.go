package mongo

import (
	"context"
	"log"

	mg "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mg.Client
var errors error
var contextForDb context.Context
var Database *mg.Database

func MongoInit() {
	SourceURL := "URL"
	Db := "DEMODB"
	c, errors = mg.NewClient(options.Client().ApplyURI(SourceURL))
	if errors != nil {
		log.Println("Mongo client connection failed")
		log.Println(errors)
	}
	contextForDb = context.Background()
	errors = c.Connect(contextForDb)
	if errors != nil {
		log.Println("mongo context setup failed")
		log.Println(errors)
	}

	Database = c.Database(Db)
}
