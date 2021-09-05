package controllers

import (
	"context"
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/mongoDb"
	"example.com/Users/akashkumar/go/demo_project/models"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func UpdateBook(res http.ResponseWriter, request *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("updating book")
	params := mux.Vars(request)
	name := params["name"]
	fmt.Println(name)
	book:=models.Book{}
	_ = json.NewDecoder(request.Body).Decode(&book)

	fmt.Println("start")

	//if(jsn == Book{})
	//err := json.Unmarshal([]byte(res.Body), &data)
	client,_ := mongoDb.GetMongoClient()
	collection := client.Database(mongoDb.DB).Collection(mongoDb.BOOKCOLL)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{"name": name}
	opts := bson.D{
		{"$set", bson.M{"name": "Martial Peak"}},
	}
	result,err := collection.UpdateOne(ctx, filter,opts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	fmt.Printf("updated one %v document(s) of %v\n", result.ModifiedCount, name)
	fmt.Println("updated book")
	json.NewEncoder(res).Encode(result)
}
