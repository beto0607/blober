package blob_slice

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Database
var coll *mongo.Collection

func InitBlobSliceDB(handler *mongo.Database) {
	db = handler
	err := handler.CreateCollection(context.TODO(), "blobs")
	if err != nil {
		log.Panicln("Couldn't create collection 'blobs'")
	}
	log.Println("Collection 'blobs' created")
	coll = handler.Collection("blobs")
}

func CreateBlobEntity() (*BlobModel, error) {
	newBlob := NewBlobModel()
	result, err := coll.InsertOne(context.TODO(), newBlob)

	if err != nil {
		return nil, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		newBlob.Id = oid
		return &newBlob, nil
	}
	return nil, errors.New("Couldn't retrieve InsertedID")
}

func SaveBlobEntity(blob *BlobModel) (*BlobModel, error) {
	blob.UpdatedAt = time.Now().UTC().String()
	filter := bson.M{"_id": blob.Id}
	update := bson.M{"$set": blob}
	_, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Panicln(err.Error())
		return nil, err
	}

	return blob, nil
}

func FindBlobEntity(id string) (*BlobModel, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	blob := BlobModel{}

	err = coll.FindOne(context.TODO(), bson.M{"_id": objectId, "deleted_at": ""}).Decode(&blob)

	if err != nil {
		return nil, err
	}
	return &blob, nil
}

func DeleteBlobEntity(blob *BlobModel, hardDelete bool) error {
	if !hardDelete {
		blob.DeletedAt = time.Now().UTC().String()
		_, err := SaveBlobEntity(blob)
		return err
	}
	filter := bson.M{"_id": blob.Id}
	_, err := coll.DeleteOne(context.TODO(), filter)
	return err
}
