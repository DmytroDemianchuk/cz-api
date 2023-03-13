package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dmytrodemianchuk/cz-api/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

// alternative connection to MongoDB Atlas
// func init() {
// 	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
// 	clientOptions := options.Client().
// 		ApplyURI("mongodb+srv://Dmytro:<password>@mongodb.nuim2dr.mongodb.net/?retryWrites=true&w=majority").
// 		SetServerAPIOptions(serverAPIOptions)
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	collection = client.Database(dbName).Collection(colName)
// }

const connectionString = "mongodb+srv://Dmytro:moqi7e123@mongodb.nuim2dr.mongodb.net/?retryWrites=true&w=majority"
const dbName = "people"
const colName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connection to MongoDB Atlas
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	//collection instance
	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

// insert 1 record

func insertOneName(name model.People) {
	inserted, err := collection.InsertOne(context.Background(), name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 name in db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneName(nameID string) {
	id, _ := primitive.ObjectIDFromHex(nameID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count:", result.ModifiedCount)
}

// delete 1 record
func deleteOneName(nameId string) {
	id, _ := primitive.ObjectIDFromHex(nameId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name got deleted with delete count:", deleteCount)
}

// delete all record from  mongoDB
func deleteAllName() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies delete:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all names form database
func getAllNames() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var names []primitive.M

	for cur.Next(context.Background()) {
		var name bson.M
		err := cur.Decode(&name)

		if err != nil {
			log.Fatal(err)
		}

		names = append(names, name)
	}

	defer cur.Close(context.Background())
	return names
}

// Actual controller - file
func GetMyAllNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allNames := getAllNames()
	json.NewEncoder(w).Encode(allNames)
}

func CreateName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var name model.People
	_ = json.NewDecoder(r.Body).Decode(&name)
	insertOneName(name)
	json.NewEncoder(w).Encode(name)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneName(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneName(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllName()
	json.NewEncoder(w).Encode(count)
}
