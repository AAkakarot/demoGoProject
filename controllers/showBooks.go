package controllers

import (
	"context"
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/mongoDb"
	"example.com/Users/akashkumar/go/demo_project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func GetBooks(res http.ResponseWriter, request *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "GET")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var books []models.Book
	log.Println(" getting books")
	client,_ := mongoDb.GetMongoClient()
	collection := client.Database(mongoDb.DB).Collection(mongoDb.BOOKCOLL)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Find().SetSort(bson.D{{"name", 1}})
	filter := bson.D{}

	cursor,err := collection.Find(ctx, filter, opts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	defer cursor.Close(ctx)

	log.Println("books displayed successfully")

	json.NewEncoder(res).Encode(books)
}

