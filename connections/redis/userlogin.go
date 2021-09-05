package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

type User struct {
	Email string `json:"mail"`
	Password string `json:"pass"`
}

func SaveUser(ctx context.Context, user User){
	log.Println("connecting ---to----redis----")

	client := getClient(Ctx)
	err := client.Set(ctx,"abc@gmail.com", "bca", 0).Err()
	if err != nil {
		panic(err)
	}
	err2 := client.Set(ctx,user.Email, user.Password, 0).Err()
	if err2 != nil {
		panic(err2)
	}

}

func CheckLogin(ctx context.Context, user User, res http.ResponseWriter) bool{

	mail := user.Email
	client := getClient(Ctx)
	pass, err := client.Get(ctx, mail).Result()
	if err == redis.Nil {
		fmt.Fprintf(res, mail + " does not exist\n")
		return false
	} else if err != nil {
		panic(err)
		return false
	} else if pass != user.Password{
		fmt.Fprintf(res,"Invalid Password\n")
		return false
	} else{
		fmt.Fprintf(res,"login Successful\n")
		return true
	}

}