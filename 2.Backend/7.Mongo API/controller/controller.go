package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func deleteOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted movie with ID: ", result.DeletedCount)
}

func deleteAllMovie() int64 {
	result, err := collection.DeleteMany(context.Background(), bson.M{}, nil)

	if err != nil {
		log.Fatal(err)
	}

	return result.DeletedCount
}

func getAllMovies() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	//Close the cursor once the work is done
	defer curr.Close(context.Background())

	//Decode() method is used to decode the document
	for curr.Next(context.Background()) {
		var movie bson.M

		err := curr.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

// Controllers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "qpplication/x-www-form-urlencoded")

	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
