package controllers

import (
	"context"
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/mongoDb"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func DeleteoneBook(res http.ResponseWriter, request *http.Request) {
	res.Header().Set("content-type", "application/json")
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println("deleting book")
	params := mux.Vars(request)
	name := params["name"]
	fmt.Println(name)
	client,_ := mongoDb.GetMongoClient()
	collection := client.Database(mongoDb.DB).Collection(mongoDb.BOOKCOLL)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	//filter := bson.D{primitive.E{Key: "Name", Value: name}}
	result, err := collection.DeleteOne(ctx, bson.M{"name": name})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "` + err.Error() + `" }`))
		return
	}

	fmt.Printf("DeleteOne removed %v document(s) of %v\n", result.DeletedCount, name)
	json.NewEncoder(res).Encode(result)

}

