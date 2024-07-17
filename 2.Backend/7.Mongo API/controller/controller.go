package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/varnit-ta/mongo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://varnitsingh:123@mernapp.uqtnbbo.mongodb.net/?retryWrites=true&w=majority&appName=MERNapp"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// Connect with MongoDB ("init" method :- called before main method)
func init() {
	//CLient option
	clientOption := options.Client().ApplyURI(connectionString)

	//Connect to MongoDB
	//Context package is used to set deadline, timeout and cancleation of external process
	//context.TODO() returns empty context
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	//Collection instance
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance created!")
}

// MongoDB helpers - file
func insertOneMovie(movie model.Netflix) {
	//InsertOne() method returns mongo.InsertOneResult
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted movie with ID: ", inserted.InsertedID)
}

func insertManyMovies(movies []interface{}) {
	//InsertMany() method returns mongo.InsertManyResult
	inserted, err := collection.InsertMany(context.Background(), movies)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple movies with IDs: ", inserted.InsertedIDs)
}

func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	//"bson.M" is used to update the document
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated movie with ID: ", result.ModifiedCount)
}
