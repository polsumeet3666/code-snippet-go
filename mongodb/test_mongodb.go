package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func main() {

	// set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to monogb")

	// collection
	collection := client.Database("test").Collection("trainers")

	/* inserted code commented

	ash := Trainer{"ash", 10, "pallet town"}
	misty := Trainer{"misty", 10, "cerelean city"}
	brock := Trainer{"brock", 20, "pew city"}

	// insert single doc
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("inserted a single doc", insertResult.InsertedID)

	// insert multiple doc
	trainers := []interface{}{misty, brock}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted multiple documents", insertManyResult.InsertedIDs)

	*/

	/*

		// update document
		filter := bson.D{{"name", "ash"}}

		update := bson.D{
			{"$inc", bson.D{
				{"age", 1},
			}},
		}

		updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("matched %v docs and updated %v docs\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	*/

	/*
		// find doc

		filter := bson.D{{"name", "ash"}}
		var result Trainer
		err = collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("found doc %v\n", result)

	*/

	/*
		// find multiple docs

		// pass these options to find method
		findOptions := options.Find()
		findOptions.SetLimit(3)

		// array which can store decoded docs
		var results []*Trainer

		// passing bson.D{{}} as filter matches all docs in collection
		cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
		if err != nil {
			log.Fatal(err)
		}

		// finding multiple documents returns a cursor
		for cur.Next(context.TODO()) {
			// create a value into which the single doc can be decoded
			var elem Trainer
			err := cur.Decode(&elem)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, &elem)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		// close cursor
		cur.Close(context.TODO())

		fmt.Printf("found multiple documents %+v\n", results)

	*/

	// delete document
	deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("deleted %v docs\n", deleteResult.DeletedCount)

	// collection disconnect
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to mongodb closed")
}