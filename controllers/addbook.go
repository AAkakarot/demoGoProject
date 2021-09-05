package controllers

import (
	"context"
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/mongoDb"
	"example.com/Users/akashkumar/go/demo_project/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var book models.Book
	fmt.Println("\n adding book")
	_ = json.NewDecoder(req.Body).Decode(&book)
	fmt.Println(book)
	book.ID = primitive.NewObjectID()
	fmt.Print("Id is :- ")
	fmt.Println(book.ID)
	client,_ := mongoDb.GetMongoClient()
	collection := client.Database(mongoDb.DB).Collection(mongoDb.BOOKCOLL)
	result, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Println("added book")
	json.NewEncoder(res).Encode(result)
}
