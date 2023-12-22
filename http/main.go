package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27018/"

func getClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		fmt.Println(err)
	}
	var result bson.M
	if err = client.Database("admin").RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Pinged your deployment.")
	return client

}

var client *mongo.Client = getClient()
var col *mongo.Collection = client.Database("test").Collection("users")

func countDocOG(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	count, _ := client.Database("test").Collection("users").CountDocuments(ctx, bson.D{})
	fmt.Println(count)
	fmt.Fprintf(w, "%v", client)
}

func countDoc(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	count, _ := col.CountDocuments(ctx, bson.D{})
	fmt.Println(count)
	fmt.Fprintf(w, "%v", client)
}

func main() {

	fmt.Println("server started")

	http.HandleFunc("/count", countDoc)
	http.HandleFunc("/countDocOG", countDocOG)

	log.Fatal(http.ListenAndServe(":8082", nil))

	fmt.Println("server ended")
}
