package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"main/modules/database"
	"net/http"
)

func getPasswordSuggestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://pwnme:password123@localhost:27017"))

	client.Connect(database.GetMongodbContext())

	collection := client.Database("pwn-pass").Collection("passwords")

	findOptions := options.Find()
	findOptions.SetLimit(25)

	cursor, _ := collection.Find(database.GetMongodbContext(), bson.M{"password": bson.M{"$regex": "^" + r.URL.Query().Get("q")}}, findOptions)

	var results []bson.M

	if err = cursor.All(database.GetMongodbContext(), &results); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(results)
}

func handleRequests() {
	muxRouter := mux.NewRouter().StrictSlash(true)

	muxRouter.Use(mux.CORSMethodMiddleware(muxRouter))
	muxRouter.HandleFunc("/api", getPasswordSuggestion).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", muxRouter))
}

func main() {
	handleRequests()
}
