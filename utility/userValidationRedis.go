package utility

import (
	"encoding/json"
	"example.com/Users/akashkumar/go/demo_project/connections/redis"
	"fmt"
	"log"
	"net/http"
)

func Register(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "registering..\n")
	fmt.Print("connecting")
	var user redis.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	redis.SaveUser(redis.Ctx, user)
	fmt.Fprintf(res, "User " + user.Email +" registered successfully")
}

func Login(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "logging in..\n")
	var user redis.User
	_ = json.NewDecoder(req.Body).Decode(&user)
	if redis.CheckLogin(redis.Ctx, user, res){
		fmt.Fprintf(res, "User "+ user.Email +" logged in")

		//-- creating token
		validToken, err := GenerateJWT(user.Email)
		if err != nil {
			fmt.Println("Failed to generate token")
		}
		log.Println("Token generated \n"+ validToken)
		req.Header.Set("Token", validToken)

	}else{
		fmt.Fprintf(res, "User "+ user.Email +" login failed")
	}
}


