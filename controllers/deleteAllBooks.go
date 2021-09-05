package controllers

import (
	"context"
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/mongoDb"
	"example.com/Users/akashkumar/go/demo_project/utility"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func DeleteBooks(res http.ResponseWriter, request *http.Request) {
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
	result, err := collection.DeleteMany(ctx, bson.M{"Name": name})
	utility.ErrorHandler(err,"could not delete the objects!")

	fmt.Printf("DeleteOne removed %v document(s)\n of %v", result.DeletedCount, name)
	json.NewEncoder(res).Encode(result)

}
