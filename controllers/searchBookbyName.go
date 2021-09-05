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

func SearchoneBook(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Entered search")
	response.Header().Set("content-type", "application/json")
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "POST")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	response.Header().Set("content-type", "application/json")
	fmt.Println("getting  book")
	params := mux.Vars(request)
	name := params["name"]
	log.Println(name)
	var book models.Book
	client,_ := mongoDb.GetMongoClient()
	collection := client.Database(mongoDb.DB).Collection(mongoDb.BOOKCOLL)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	err := collection.FindOne(ctx, filter).Decode(&book)
	if err != nil {
		log.Println(err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Println("got book")
	json.NewEncoder(response).Encode(book)
	//err := collection.FindOne(ctx, Book{Name: name}).Decode(&book)

}
