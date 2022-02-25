package mongo

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func Update(collection string, filter interface{}, customUpdateStruct interface{}, upsert bool) (int64, error) {
	updateOptions := options.Update()
	updateOptions = updateOptions.SetUpsert(upsert)
	var modifiedCount int64
	collectionName := Database.Collection(collection)
	result, insertErr := collectionName.UpdateOne(contextForDb, filter, customUpdateStruct, updateOptions)
	if insertErr != nil {
		log.Println("Mongo Update Failed : ", insertErr)
		return modifiedCount, insertErr
	}
	log.Printf("[MONGO UPDATE] to Collection: %s", collection)
	return result.ModifiedCount, nil
}
