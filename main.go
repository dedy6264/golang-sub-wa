package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sub/modelWA"
	"sub/usecase"

	"github.com/go-redis/redis/v8"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})
var ctx = context.Background()

func main() {
	subscriber := redisClient.Subscribe(ctx, "send-user-dat")
	user := modelWA.User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			panic(err)
		}

		if user.Index == "1" {
			fmt.Println("Received message from " + msg.Channel + " channel.")
			fmt.Printf("%+v\n", user)
			usecase.Sending(user)
		}
	}

}
