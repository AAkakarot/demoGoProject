package controllers

import (
	"context"
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/mongoDb"
	"example.com/Users/akashkumar/go/demo_project/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"time"
)

func SearchBooks(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("content-type", "application/json")
	log.Println("getting related books")
	params := mux.Vars(request)
	name := params["name"]
	fmt.Println(name)
	var books []models.Book
	client,_ := mongoDb.GetMongoClient()
	collection := client.Database(mongoDb.DB).Collection(mongoDb.BOOKCOLL)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//opts := options.Find().SetSort(bson.D{{"Name", 1}})
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	cursor,err := collection.Find(ctx, filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}
	log.Println("fin")

	json.NewEncoder(response).Encode(books)
}

